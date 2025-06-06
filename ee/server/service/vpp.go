package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/fleetdm/fleet/v4/server"
	"github.com/fleetdm/fleet/v4/server/authz"
	"github.com/fleetdm/fleet/v4/server/contexts/ctxerr"
	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/fleetdm/fleet/v4/server/mdm/apple/itunes"
	"github.com/fleetdm/fleet/v4/server/mdm/apple/vpp"
	"github.com/go-kit/log/level"
)

// Used for overriding the env var value in testing
var testSetEmptyPrivateKey bool

// getVPPToken returns the base64 encoded VPP token, ready for use in requests to Apple's VPP API.
// It returns an error if the token is expired.
func (svc *Service) getVPPToken(ctx context.Context, teamID *uint) (string, error) {
	token, err := svc.ds.GetVPPTokenByTeamID(ctx, teamID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fleet.NewUserMessageError(errors.New("No available VPP Token"), http.StatusUnprocessableEntity)
		}
		return "", ctxerr.Wrap(ctx, err, "fetching vpp token")
	}

	if time.Now().After(token.RenewDate) {
		return "", fleet.NewUserMessageError(errors.New("Couldn't install. VPP token expired."), http.StatusUnprocessableEntity)
	}

	return token.Token, nil
}

func (svc *Service) BatchAssociateVPPApps(ctx context.Context, teamName string, payloads []fleet.VPPBatchPayload, dryRun bool) ([]fleet.VPPAppResponse, error) {
	if err := svc.authz.Authorize(ctx, &fleet.Team{}, fleet.ActionRead); err != nil {
		return nil, err
	}

	var teamID *uint
	if teamName != "" {
		tm, err := svc.ds.TeamByName(ctx, teamName)
		if err != nil {
			// If this is a dry run, the team may not have been created yet
			if dryRun && fleet.IsNotFound(err) {
				return nil, nil
			}
			return nil, err
		}
		teamID = &tm.ID
	}

	if err := svc.authz.Authorize(ctx, &fleet.SoftwareInstaller{TeamID: teamID}, fleet.ActionWrite); err != nil {
		return nil, ctxerr.Wrap(ctx, err, "validating authorization")
	}

	// Adding VPP apps will add them to all available platforms per decision:
	// https://github.com/fleetdm/fleet/issues/19447#issuecomment-2256598681
	// The code is already here to support individual platforms, so we can easily enable it later.

	payloadsWithPlatform := make([]fleet.VPPBatchPayloadWithPlatform, 0, len(payloads))
	for _, payload := range payloads {
		// Currently only macOS is supported for self-service. Don't
		// import vpp apps as self-service for ios or ipados
		payloadsWithPlatform = append(payloadsWithPlatform, []fleet.VPPBatchPayloadWithPlatform{{
			AppStoreID:       payload.AppStoreID,
			SelfService:      false,
			Platform:         fleet.IOSPlatform,
			LabelsExcludeAny: payload.LabelsExcludeAny,
			LabelsIncludeAny: payload.LabelsIncludeAny,
			Categories:       payload.Categories,
		}, {
			AppStoreID:       payload.AppStoreID,
			SelfService:      false,
			Platform:         fleet.IPadOSPlatform,
			LabelsExcludeAny: payload.LabelsExcludeAny,
			LabelsIncludeAny: payload.LabelsIncludeAny,
			Categories:       payload.Categories,
		}, {
			AppStoreID:         payload.AppStoreID,
			SelfService:        payload.SelfService,
			Platform:           fleet.MacOSPlatform,
			InstallDuringSetup: payload.InstallDuringSetup,
			LabelsExcludeAny:   payload.LabelsExcludeAny,
			LabelsIncludeAny:   payload.LabelsIncludeAny,
			Categories:         payload.Categories,
		}}...)
	}

	if dryRun {
		// On dry runs return early because the VPP token might not exist yet
		// and we don't want to apply the VPP apps.
		return nil, nil
	}

	var vppAppTeams []fleet.VPPAppTeam
	// Don't check for token if we're only disassociating assets
	if len(payloads) > 0 {
		token, err := svc.getVPPToken(ctx, teamID)
		if err != nil {
			return nil, fleet.NewUserMessageError(ctxerr.Wrap(ctx, err, "could not retrieve vpp token"), http.StatusUnprocessableEntity)
		}

		for _, payload := range payloadsWithPlatform {
			if payload.Platform == "" {
				payload.Platform = fleet.MacOSPlatform
			}
			if payload.Platform != fleet.IOSPlatform && payload.Platform != fleet.IPadOSPlatform && payload.Platform != fleet.MacOSPlatform {
				return nil, fleet.NewInvalidArgumentError("app_store_apps.platform",
					fmt.Sprintf("platform must be one of '%s', '%s', or '%s", fleet.IOSPlatform, fleet.IPadOSPlatform, fleet.MacOSPlatform))
			}
			validatedLabels, err := ValidateSoftwareLabels(ctx, svc, payload.LabelsIncludeAny, payload.LabelsExcludeAny)
			if err != nil {
				return nil, ctxerr.Wrap(ctx, err, "validating software labels for batch adding vpp app")
			}

			payload.Categories = server.RemoveDuplicatesFromSlice(payload.Categories)
			catIDs, err := svc.ds.GetSoftwareCategoryIDs(ctx, payload.Categories)
			if err != nil {
				return nil, ctxerr.Wrap(ctx, err, "getting software category ids")
			}

			if len(catIDs) != len(payload.Categories) {
				return nil, &fleet.BadRequestError{
					Message:     "some or all of the categories provided don't exist",
					InternalErr: fmt.Errorf("categories provided: %v", payload.Categories),
				}
			}

			vppAppTeams = append(vppAppTeams, fleet.VPPAppTeam{
				VPPAppID: fleet.VPPAppID{
					AdamID:   payload.AppStoreID,
					Platform: payload.Platform,
				},
				SelfService:        payload.SelfService,
				InstallDuringSetup: payload.InstallDuringSetup,
				ValidatedLabels:    validatedLabels,
				CategoryIDs:        catIDs,
			})
		}

		var missingAssets []string

		assets, err := vpp.GetAssets(token, nil)
		if err != nil {
			return nil, ctxerr.Wrap(ctx, err, "unable to retrieve assets")
		}

		assetMap := map[string]struct{}{}
		for _, asset := range assets {
			assetMap[asset.AdamID] = struct{}{}
		}

		for _, vppAppID := range vppAppTeams {
			if _, ok := assetMap[vppAppID.AdamID]; !ok {
				missingAssets = append(missingAssets, vppAppID.AdamID)
			}
		}

		if len(missingAssets) != 0 {
			reqErr := ctxerr.Errorf(ctx, "requested app not available on vpp account: %s", strings.Join(missingAssets, ","))
			return nil, fleet.NewUserMessageError(reqErr, http.StatusUnprocessableEntity)
		}
	}

	if len(vppAppTeams) > 0 {
		apps, err := getVPPAppsMetadata(ctx, vppAppTeams)
		if err != nil {
			return nil, ctxerr.Wrap(ctx, err, "refreshing VPP app metadata")
		}
		if len(apps) == 0 {
			return nil, fleet.NewInvalidArgumentError("app_store_apps",
				"no valid apps found matching the provided app store IDs and platforms")
		}

		if err := svc.ds.BatchInsertVPPApps(ctx, apps); err != nil {
			return nil, ctxerr.Wrap(ctx, err, "inserting vpp app metadata")
		}
		// Filter out the apps with invalid platforms
		if len(apps) != len(vppAppTeams) {
			vppAppTeams = make([]fleet.VPPAppTeam, 0, len(apps))
			for _, app := range apps {
				vppAppTeams = append(vppAppTeams, app.VPPAppTeam)
			}
		}

	}
	if err := svc.ds.SetTeamVPPApps(ctx, teamID, vppAppTeams); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fleet.NewUserMessageError(ctxerr.Wrap(ctx, err, "no vpp token to set team vpp assets"), http.StatusUnprocessableEntity)
		}
		return nil, ctxerr.Wrap(ctx, err, "set team vpp assets")
	}

	if len(vppAppTeams) == 0 {
		return []fleet.VPPAppResponse{}, nil
	}

	return svc.ds.GetVPPApps(ctx, teamID)
}

func (svc *Service) GetAppStoreApps(ctx context.Context, teamID *uint) ([]*fleet.VPPApp, error) {
	if err := svc.authz.Authorize(ctx, &fleet.VPPApp{TeamID: teamID}, fleet.ActionRead); err != nil {
		return nil, err
	}

	vppToken, err := svc.getVPPToken(ctx, teamID)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "retrieving VPP token")
	}

	assets, err := vpp.GetAssets(vppToken, nil)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "fetching Apple VPP assets")
	}

	if len(assets) == 0 {
		return []*fleet.VPPApp{}, nil
	}

	var adamIDs []string
	for _, a := range assets {
		adamIDs = append(adamIDs, a.AdamID)
	}

	assetMetadata, err := itunes.GetAssetMetadata(adamIDs, &itunes.AssetMetadataFilter{Entity: "software"})
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "fetching VPP asset metadata")
	}

	assignedApps, err := svc.ds.GetAssignedVPPApps(ctx, teamID)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "retrieving assigned VPP apps")
	}

	var apps []*fleet.VPPApp
	var appsToUpdate []*fleet.VPPApp
	for _, a := range assets {
		m, ok := assetMetadata[a.AdamID]
		if !ok {
			// Then this adam_id is not a VPP software entity, so skip it.
			continue
		}

		platforms := getPlatformsFromSupportedDevices(m.SupportedDevices)

		for platform := range platforms {
			vppAppID := fleet.VPPAppID{
				AdamID:   a.AdamID,
				Platform: platform,
			}
			vppAppTeam := fleet.VPPAppTeam{
				VPPAppID: vppAppID,
			}
			app := &fleet.VPPApp{
				VPPAppTeam:       vppAppTeam,
				BundleIdentifier: m.BundleID,
				IconURL:          m.ArtworkURL,
				Name:             m.TrackName,
				LatestVersion:    m.Version,
			}

			if appFleet, ok := assignedApps[vppAppID]; ok {
				// Then this is already assigned, so filter it out.
				app.SelfService = appFleet.SelfService
				appsToUpdate = append(appsToUpdate, app)
				continue
			}

			apps = append(apps, app)
		}
	}

	if len(appsToUpdate) > 0 {
		if err := svc.ds.BatchInsertVPPApps(ctx, appsToUpdate); err != nil {
			return nil, ctxerr.Wrap(ctx, err, "updating existing VPP apps")
		}
	}

	// Sort apps by name and by platform
	sort.Slice(apps, func(i, j int) bool {
		if apps[i].Name != apps[j].Name {
			return apps[i].Name < apps[j].Name
		}
		return apps[i].Platform < apps[j].Platform
	})

	return apps, nil
}

func getPlatformsFromSupportedDevices(supportedDevices []string) map[fleet.AppleDevicePlatform]struct{} {
	platforms := make(map[fleet.AppleDevicePlatform]struct{}, 1)
	if len(supportedDevices) == 0 {
		platforms[fleet.MacOSPlatform] = struct{}{}
		return platforms
	}
	for _, device := range supportedDevices {
		// It is rare that a single app supports all platforms, but it is possible.
		switch {
		case strings.HasPrefix(device, "iPhone"):
			platforms[fleet.IOSPlatform] = struct{}{}
		case strings.HasPrefix(device, "iPad"):
			platforms[fleet.IPadOSPlatform] = struct{}{}
		case strings.HasPrefix(device, "Mac"):
			platforms[fleet.MacOSPlatform] = struct{}{}
		}
	}
	return platforms
}

func (svc *Service) AddAppStoreApp(ctx context.Context, teamID *uint, appID fleet.VPPAppTeam) (uint, error) {
	if err := svc.authz.Authorize(ctx, &fleet.VPPApp{TeamID: teamID}, fleet.ActionWrite); err != nil {
		return 0, err
	}
	if appID.AddAutoInstallPolicy {
		// Currently, same write permissions are applied on software and policies,
		// but leaving this here in case it changes in the future.
		if err := svc.authz.Authorize(ctx, &fleet.Policy{PolicyData: fleet.PolicyData{TeamID: teamID}}, fleet.ActionWrite); err != nil {
			return 0, err
		}
	}

	// Validate platform
	if appID.Platform == "" {
		appID.Platform = fleet.MacOSPlatform
	}
	if appID.Platform != fleet.IOSPlatform && appID.Platform != fleet.IPadOSPlatform && appID.Platform != fleet.MacOSPlatform {
		return 0, fleet.NewInvalidArgumentError("platform",
			fmt.Sprintf("platform must be one of '%s', '%s', or '%s", fleet.IOSPlatform, fleet.IPadOSPlatform, fleet.MacOSPlatform))
	}

	validatedLabels, err := ValidateSoftwareLabels(ctx, svc, appID.LabelsIncludeAny, appID.LabelsExcludeAny)
	if err != nil {
		return 0, ctxerr.Wrap(ctx, err, "validating software labels for adding vpp app")
	}

	var teamName string
	if teamID != nil && *teamID != 0 {
		tm, err := svc.ds.Team(ctx, *teamID)
		if fleet.IsNotFound(err) {
			return 0, fleet.NewInvalidArgumentError("team_id", fmt.Sprintf("team %d does not exist", *teamID)).
				WithStatus(http.StatusNotFound)
		} else if err != nil {
			return 0, ctxerr.Wrap(ctx, err, "checking if team exists")
		}

		teamName = tm.Name
	}

	if appID.SelfService && appID.Platform != fleet.MacOSPlatform {
		return 0, fleet.NewUserMessageError(errors.New("Currently, self-service is only supported on macOS, Windows, and Linux. Please add the app without self_service and manually install it on the Host details page."), http.StatusBadRequest)
	}
	if appID.AddAutoInstallPolicy && appID.Platform != fleet.MacOSPlatform {
		return 0, fleet.NewUserMessageError(errors.New("Currently, automatic install is only supported on macOS, Windows, and Linux. Please add the app without automatic_install and manually install it on the Host details page."), http.StatusBadRequest)
	}

	vppToken, err := svc.getVPPToken(ctx, teamID)
	if err != nil {
		return 0, ctxerr.Wrap(ctx, err, "retrieving VPP token")
	}

	assets, err := vpp.GetAssets(vppToken, &vpp.AssetFilter{AdamID: appID.AdamID})
	if err != nil {
		return 0, ctxerr.Wrap(ctx, err, "retrieving VPP asset")
	}

	if len(assets) == 0 {
		return 0, ctxerr.New(ctx,
			fmt.Sprintf("Error: Couldn't add software. %s isn't available in Apple Business Manager. Please purchase license in Apple Business Manager and try again.",
				appID.AdamID))
	}

	asset := assets[0]

	assetMetadata, err := itunes.GetAssetMetadata([]string{asset.AdamID}, &itunes.AssetMetadataFilter{Entity: "software"})
	if err != nil {
		return 0, ctxerr.Wrap(ctx, err, "fetching VPP asset metadata")
	}

	assetMD := assetMetadata[asset.AdamID]

	platforms := getPlatformsFromSupportedDevices(assetMD.SupportedDevices)
	if _, ok := platforms[appID.Platform]; !ok {
		return 0, fleet.NewInvalidArgumentError("app_store_id", fmt.Sprintf("%s isn't available for %s", assetMD.TrackName, appID.Platform))
	}

	if appID.Platform == fleet.MacOSPlatform {
		// Check if we've already added an installer for this app
		exists, err := svc.ds.UploadedSoftwareExists(ctx, assetMD.BundleID, teamID)
		if err != nil {
			return 0, ctxerr.Wrap(ctx, err, "checking existence of VPP app installer")
		}

		if exists {
			return 0, ctxerr.New(ctx,
				fmt.Sprintf("Error: Couldn't add software. %s already has software available for install on the %s team.",
					assetMD.TrackName, teamName))
		}
	}

	appID.ValidatedLabels = validatedLabels

	appID.Categories = server.RemoveDuplicatesFromSlice(appID.Categories)
	catIDs, err := svc.ds.GetSoftwareCategoryIDs(ctx, appID.Categories)
	if err != nil {
		return 0, ctxerr.Wrap(ctx, err, "getting software category ids")
	}

	if len(catIDs) != len(appID.Categories) {
		return 0, &fleet.BadRequestError{
			Message:     "some or all of the categories provided don't exist",
			InternalErr: fmt.Errorf("categories provided: %v", appID.Categories),
		}
	}
	appID.CategoryIDs = catIDs

	app := &fleet.VPPApp{
		VPPAppTeam:       appID,
		BundleIdentifier: assetMD.BundleID,
		IconURL:          assetMD.ArtworkURL,
		Name:             assetMD.TrackName,
		LatestVersion:    assetMD.Version,
	}

	addedApp, err := svc.ds.InsertVPPAppWithTeam(ctx, app, teamID)
	if err != nil {
		return 0, ctxerr.Wrap(ctx, err, "writing VPP app to db")
	}

	actLabelsIncl, actLabelsExcl := activitySoftwareLabelsFromValidatedLabels(addedApp.ValidatedLabels)

	act := fleet.ActivityAddedAppStoreApp{
		AppStoreID:       app.AdamID,
		Platform:         app.Platform,
		TeamName:         &teamName,
		SoftwareTitle:    app.Name,
		SoftwareTitleId:  addedApp.TitleID,
		TeamID:           teamID,
		SelfService:      app.SelfService,
		LabelsIncludeAny: actLabelsIncl,
		LabelsExcludeAny: actLabelsExcl,
	}
	if err := svc.NewActivity(ctx, authz.UserFromContext(ctx), act); err != nil {
		return 0, ctxerr.Wrap(ctx, err, "create activity for add app store app")
	}

	if appID.AddAutoInstallPolicy && app.AddedAutomaticInstallPolicy != nil {
		policyAct := fleet.ActivityTypeCreatedPolicy{
			ID:   app.AddedAutomaticInstallPolicy.ID,
			Name: app.AddedAutomaticInstallPolicy.Name,
		}

		if err := svc.NewActivity(ctx, authz.UserFromContext(ctx), policyAct); err != nil {
			level.Warn(svc.logger).Log("msg", "failed to create activity for create automatic install policy for app store app", "err", err)
		}

	}

	return addedApp.TitleID, nil
}

func getVPPAppsMetadata(ctx context.Context, ids []fleet.VPPAppTeam) ([]*fleet.VPPApp, error) {
	var apps []*fleet.VPPApp

	// Map of adamID to platform, then to whether it's available as self-service
	// and installed during setup.
	adamIDMap := make(map[string]map[fleet.AppleDevicePlatform]fleet.VPPAppTeam)
	for _, id := range ids {
		if _, ok := adamIDMap[id.AdamID]; !ok {
			adamIDMap[id.AdamID] = make(map[fleet.AppleDevicePlatform]fleet.VPPAppTeam, 1)
			adamIDMap[id.AdamID][id.Platform] = fleet.VPPAppTeam{
				SelfService:        id.SelfService,
				InstallDuringSetup: id.InstallDuringSetup,
				ValidatedLabels:    id.ValidatedLabels,
				AppTeamID:          id.AppTeamID,
				Categories:         id.Categories,
				CategoryIDs:        id.CategoryIDs,
			}
		} else {
			adamIDMap[id.AdamID][id.Platform] = fleet.VPPAppTeam{
				SelfService:        id.SelfService,
				InstallDuringSetup: id.InstallDuringSetup,
				ValidatedLabels:    id.ValidatedLabels,
				AppTeamID:          id.AppTeamID,
				Categories:         id.Categories,
				CategoryIDs:        id.CategoryIDs,
			}
		}
	}

	var adamIDs []string
	for adamID := range adamIDMap {
		adamIDs = append(adamIDs, adamID)
	}
	assetMetatada, err := itunes.GetAssetMetadata(adamIDs, &itunes.AssetMetadataFilter{Entity: "software"})
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "fetching VPP asset metadata")
	}

	for adamID, metadata := range assetMetatada {
		platforms := getPlatformsFromSupportedDevices(metadata.SupportedDevices)
		for platform := range platforms {
			if props, ok := adamIDMap[adamID][platform]; ok {
				app := &fleet.VPPApp{
					VPPAppTeam: fleet.VPPAppTeam{
						VPPAppID: fleet.VPPAppID{
							AdamID:   adamID,
							Platform: platform,
						},
						SelfService:        props.SelfService,
						InstallDuringSetup: props.InstallDuringSetup,
						ValidatedLabels:    props.ValidatedLabels,
						AppTeamID:          props.AppTeamID,
						Categories:         props.Categories,
						CategoryIDs:        props.CategoryIDs,
					},
					BundleIdentifier: metadata.BundleID,
					IconURL:          metadata.ArtworkURL,
					Name:             metadata.TrackName,
					LatestVersion:    metadata.Version,
				}
				apps = append(apps, app)
			} else {
				continue
			}
		}
	}

	return apps, nil
}

func (svc *Service) UpdateAppStoreApp(ctx context.Context, titleID uint, teamID *uint, selfService bool, labelsIncludeAny, labelsExcludeAny, categories []string) (*fleet.VPPAppStoreApp, error) {
	if err := svc.authz.Authorize(ctx, &fleet.VPPApp{TeamID: teamID}, fleet.ActionWrite); err != nil {
		return nil, err
	}

	var teamName string
	if teamID != nil && *teamID != 0 {
		tm, err := svc.ds.Team(ctx, *teamID)
		if fleet.IsNotFound(err) {
			return nil, fleet.NewInvalidArgumentError("team_id", fmt.Sprintf("team %d does not exist", *teamID)).
				WithStatus(http.StatusNotFound)
		} else if err != nil {
			return nil, ctxerr.Wrap(ctx, err, "UpdateAppStoreApp: checking if team exists")
		}

		teamName = tm.Name
	}

	validatedLabels, err := ValidateSoftwareLabels(ctx, svc, labelsIncludeAny, labelsExcludeAny)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "UpdateAppStoreApp: validating software labels")
	}

	meta, err := svc.ds.GetVPPAppMetadataByTeamAndTitleID(ctx, teamID, titleID)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "UpdateAppStoreApp: getting vpp app metadata")
	}

	if selfService && meta.Platform != fleet.MacOSPlatform {
		return nil, fleet.NewUserMessageError(errors.New("Currently, self-service only supports macOS"), http.StatusBadRequest)
	}

	appToWrite := &fleet.VPPApp{
		VPPAppTeam: fleet.VPPAppTeam{
			VPPAppID: fleet.VPPAppID{
				AdamID: meta.AdamID, Platform: meta.Platform,
			},
			SelfService:     selfService,
			ValidatedLabels: validatedLabels,
		},
		TeamID:           teamID,
		TitleID:          titleID,
		BundleIdentifier: meta.BundleIdentifier,
		Name:             meta.Name,
		LatestVersion:    meta.LatestVersion,
	}
	if meta.IconURL != nil {
		appToWrite.IconURL = *meta.IconURL
	}

	categories = server.RemoveDuplicatesFromSlice(categories)
	catIDs, err := svc.ds.GetSoftwareCategoryIDs(ctx, categories)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "getting software category ids")
	}

	if len(catIDs) != len(categories) {
		return nil, &fleet.BadRequestError{
			Message:     "some or all of the categories provided don't exist",
			InternalErr: fmt.Errorf("categories provided: %v", categories),
		}
	}

	appToWrite.CategoryIDs = catIDs

	// check if labels have changed
	var existingLabels fleet.LabelIdentsWithScope
	switch {
	case len(meta.LabelsExcludeAny) > 0:
		existingLabels.LabelScope = fleet.LabelScopeExcludeAny
		existingLabels.ByName = make(map[string]fleet.LabelIdent, len(meta.LabelsExcludeAny))
		for _, l := range meta.LabelsExcludeAny {
			existingLabels.ByName[l.LabelName] = fleet.LabelIdent{LabelName: l.LabelName, LabelID: l.LabelID}
		}

	case len(meta.LabelsIncludeAny) > 0:
		existingLabels.LabelScope = fleet.LabelScopeIncludeAny
		existingLabels.ByName = make(map[string]fleet.LabelIdent, len(meta.LabelsIncludeAny))
		for _, l := range meta.LabelsIncludeAny {
			existingLabels.ByName[l.LabelName] = fleet.LabelIdent{LabelName: l.LabelName, LabelID: l.LabelID}
		}
	}

	labelsChanged := !validatedLabels.Equal(&existingLabels)

	// Get the hosts that are NOT in label scope currently (before the update happens)
	var hostsNotInScope map[uint]struct{}
	if labelsChanged {
		hostsNotInScope, err = svc.ds.GetExcludedHostIDMapForVPPApp(ctx, meta.VPPAppsTeamsID)
		if err != nil {
			return nil, ctxerr.Wrap(ctx, err, "UpdateAppStoreApp: getting hosts not in scope for vpp app")
		}
	}

	// Update the app
	_, err = svc.ds.InsertVPPAppWithTeam(ctx, appToWrite, teamID)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "UpdateAppStoreApp: write app to db")
	}

	if labelsChanged {
		// Get the hosts that are now IN label scope (after the update)
		hostsInScope, err := svc.ds.GetIncludedHostIDMapForVPPApp(ctx, meta.VPPAppsTeamsID)
		if err != nil {
			return nil, ctxerr.Wrap(ctx, err, "UpdateAppStoreApp: getting hosts in scope for vpp app")
		}

		var hostsToClear []uint
		for id := range hostsInScope {
			if _, ok := hostsNotInScope[id]; ok {
				// it was not in scope but now it is, so we should clear policy status
				hostsToClear = append(hostsToClear, id)
			}
		}

		// We clear the policy status here because otherwise the policy automation machinery
		// won't pick this up and the software won't install.
		if err := svc.ds.ClearVPPAppAutoInstallPolicyStatusForHosts(ctx, meta.VPPAppsTeamsID, hostsToClear); err != nil {
			return nil, ctxerr.Wrap(ctx, err, "failed to clear auto install policy status for host")
		}
	}

	actLabelsIncl, actLabelsExcl := activitySoftwareLabelsFromValidatedLabels(validatedLabels)

	act := fleet.ActivityEditedAppStoreApp{
		TeamName:         &teamName,
		TeamID:           teamID,
		SelfService:      selfService,
		SoftwareTitleID:  titleID,
		SoftwareTitle:    meta.Name,
		AppStoreID:       meta.AdamID,
		Platform:         meta.Platform,
		LabelsIncludeAny: actLabelsIncl,
		LabelsExcludeAny: actLabelsExcl,
	}
	if err := svc.NewActivity(ctx, authz.UserFromContext(ctx), act); err != nil {
		return nil, ctxerr.Wrap(ctx, err, "create activity for update app store app")
	}

	updatedAppMeta, err := svc.ds.GetVPPAppMetadataByTeamAndTitleID(ctx, teamID, titleID)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "UpdateAppStoreApp: getting updated app metadata")
	}

	return updatedAppMeta, nil
}

func (svc *Service) UploadVPPToken(ctx context.Context, token io.ReadSeeker) (*fleet.VPPTokenDB, error) {
	if err := svc.authz.Authorize(ctx, &fleet.AppleCSR{}, fleet.ActionWrite); err != nil {
		return nil, err
	}

	privateKey := svc.config.Server.PrivateKey
	if testSetEmptyPrivateKey {
		privateKey = ""
	}

	if len(privateKey) == 0 {
		return nil, ctxerr.New(ctx, "Couldn't add content token. Missing required private key. Learn how to configure the private key here: https://fleetdm.com/learn-more-about/fleet-server-private-key")
	}

	if token == nil {
		return nil, ctxerr.Wrap(ctx, fleet.NewInvalidArgumentError("token", "Invalid token. Please provide a valid content token from Apple Business Manager."))
	}

	tokenBytes, err := io.ReadAll(token)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "reading VPP token")
	}

	locName, err := vpp.GetConfig(string(tokenBytes))
	if err != nil {
		var vppErr *vpp.ErrorResponse
		if errors.As(err, &vppErr) {
			// Per https://developer.apple.com/documentation/devicemanagement/app_and_book_management/app_and_book_management_legacy/interpreting_error_codes
			if vppErr.ErrorNumber == 9622 {
				return nil, ctxerr.Wrap(ctx, fleet.NewInvalidArgumentError("token", "Invalid token. Please provide a valid content token from Apple Business Manager."))
			}
		}
		return nil, ctxerr.Wrap(ctx, err, "validating VPP token with Apple")
	}

	data := fleet.VPPTokenData{
		Token:    string(tokenBytes),
		Location: locName,
	}

	tok, err := svc.ds.InsertVPPToken(ctx, &data)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "writing VPP token to db")
	}

	if err := svc.NewActivity(ctx, authz.UserFromContext(ctx), fleet.ActivityEnabledVPP{
		Location: locName,
	}); err != nil {
		return nil, ctxerr.Wrap(ctx, err, "create activity for upload VPP token")
	}

	return tok, nil
}

func (svc *Service) UpdateVPPToken(ctx context.Context, tokenID uint, token io.ReadSeeker) (*fleet.VPPTokenDB, error) {
	if err := svc.authz.Authorize(ctx, &fleet.AppleCSR{}, fleet.ActionWrite); err != nil {
		return nil, err
	}

	privateKey := svc.config.Server.PrivateKey
	if testSetEmptyPrivateKey {
		privateKey = ""
	}

	if len(privateKey) == 0 {
		return nil, ctxerr.New(ctx, "Couldn't add content token. Missing required private key. Learn how to configure the private key here: https://fleetdm.com/learn-more-about/fleet-server-private-key")
	}

	if token == nil {
		return nil, ctxerr.Wrap(ctx, fleet.NewInvalidArgumentError("token", "Invalid token. Please provide a valid content token from Apple Business Manager."))
	}

	tokenBytes, err := io.ReadAll(token)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "reading VPP token")
	}

	locName, err := vpp.GetConfig(string(tokenBytes))
	if err != nil {
		var vppErr *vpp.ErrorResponse
		if errors.As(err, &vppErr) {
			// Per https://developer.apple.com/documentation/devicemanagement/app_and_book_management/app_and_book_management_legacy/interpreting_error_codes
			if vppErr.ErrorNumber == 9622 {
				return nil, ctxerr.Wrap(ctx, fleet.NewInvalidArgumentError("token", "Invalid token. Please provide a valid content token from Apple Business Manager."))
			}
		}
		return nil, ctxerr.Wrap(ctx, err, "validating VPP token with Apple")
	}

	data := fleet.VPPTokenData{
		Token:    string(tokenBytes),
		Location: locName,
	}

	tok, err := svc.ds.UpdateVPPToken(ctx, tokenID, &data)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "updating vpp token")
	}

	return tok, nil
}

func (svc *Service) UpdateVPPTokenTeams(ctx context.Context, tokenID uint, teamIDs []uint) (*fleet.VPPTokenDB, error) {
	if err := svc.authz.Authorize(ctx, &fleet.AppleCSR{}, fleet.ActionWrite); err != nil {
		return nil, err
	}

	tok, err := svc.ds.UpdateVPPTokenTeams(ctx, tokenID, teamIDs)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "updating vpp token team")
	}

	return tok, nil
}

func (svc *Service) GetVPPTokens(ctx context.Context) ([]*fleet.VPPTokenDB, error) {
	if err := svc.authz.Authorize(ctx, &fleet.AppleCSR{}, fleet.ActionRead); err != nil {
		return nil, err
	}

	return svc.ds.ListVPPTokens(ctx)
}

func (svc *Service) DeleteVPPToken(ctx context.Context, tokenID uint) error {
	if err := svc.authz.Authorize(ctx, &fleet.AppleCSR{}, fleet.ActionWrite); err != nil {
		return err
	}
	tok, err := svc.ds.GetVPPToken(ctx, tokenID)
	if err != nil {
		return ctxerr.Wrap(ctx, err, "getting vpp token")
	}
	if err := svc.NewActivity(ctx, authz.UserFromContext(ctx), fleet.ActivityDisabledVPP{
		Location: tok.Location,
	}); err != nil {
		return ctxerr.Wrap(ctx, err, "create activity for delete VPP token")
	}

	return svc.ds.DeleteVPPToken(ctx, tokenID)
}
