# monitoring

## Monitoring addon

This addon enables Cloudwatch monitoring for Fleet.

This includes:

* 5XX Errors on ALB
* ECS Service Monitoring
* RDS Monitoring
* Redis Monitoring
* ACM Certificate Monitoring
* A custom Lambda to check the Fleet DB for Cron runs

## Preparation

Some of the for\_each and counts in this module cannot pre-determine the numbers until the `main` fleet module is applied.

You will need to `terraform apply -target module.main` prior applying monitoring assuming the use of a configuration matching the example at https://github.com/fleetdm/fleet/blob/main/terraform/example/main.tf.

Multiple alb support was added in order to allow monitoring `saml-auth-proxy`. See https://github.com/fleetdm/fleet/tree/main/terraform/addons/saml-auth-proxy

## Example configuration

This assumes your fleet module is `main` and is configured with it's default documentation.

https://github.com/fleetdm/fleet/blob/main/terraform/example/main.tf for details.

```
module "monitoring" {
  source                      = "github.com/fleetdm/fleet//terraform/addons/monitoring?ref=tf-mod-addon-monitoring-v1.1.0"
  customer_prefix             = local.customer
  fleet_ecs_service_name      = module.main.byo-vpc.byo-db.byo-ecs.service.name
  albs                        = [
    {
      name = module.main.byo-vpc.byo-db.alb.lb_dns_name,
      target_group_name = module.main.byo-vpc.byo-db.alb.target_group_names[0]
      target_group_arn_suffix = module.main.byo-vpc.byo-db.alb.target_group_arn_suffixes[0]
      arn_suffix = module.main.byo-vpc.byo-db.alb.lb_arn_suffix
      ecs_service_name = module.main.byo-vpc.byo-db.byo-ecs.service.name
      min_containers = module.main.byo-vpc.byo-db.byo-ecs.appautoscaling_target.min_capacity
    },
  ]
  # Only publish alerts for items in this map
  sns_topic_arns_map = {
    alb_httpcode_5xx = [var.sns_topic_arn]
    cron_monitoring  = [var.sns_topic_arn]
  }
  mysql_cluster_members = module.main.byo-vpc.rds.cluster_members
  # The cloudposse module seems to have a nested list here.
  redis_cluster_members = module.main.byo-vpc.redis.member_clusters[0]
  acm_certificate_arn   = module.acm.acm_certificate_arn
  cron_monitoring = {
    mysql_host                 = module.main.byo-vpc.rds.cluster_reader_endpoint
    mysql_database             = module.main.byo-vpc.rds.cluster_database_name
    mysql_user                 = module.main.byo-vpc.rds.cluster_master_username
    mysql_password_secret_name = module.main.byo-vpc.secrets.secret_ids["${local.customer}-database-password"]
    rds_security_group_id      = module.main.byo-vpc.rds.security_group_id
    subnet_ids                 = module.main.vpc.private_subnets
    vpc_id                     = module.main.vpc.vpc_id
    # Format of https://pkg.go.dev/time#ParseDuration
    delay_tolerance            = "2h"
    # Interval format for: https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html#rate-based
    run_interval               = "1 hour"
  }
}
```

## SNS topic ARNs map

Valid targets for `sns_topic_arns_map`:

* acm\_certificate\_expired
* alb\_helthyhosts
* alb\_httpcode\_5xx
* backend\_response\_time
* cron\_monitoring
* rds\_cpu\_untilizaton\_too\_high
* rds\_db\_event\_subscription
* redis\_cpu\_engine\_utilization
* redis\_cpu\_utilization
* redis\_current\_connections
* redis\_database\_memory\_percentage
* redis\_replication\_lag

If you want to publish to all, use `default_sns_topic_arns` instead and include your notification ARNs there.

### Requirements

No requirements.

### Providers

| Name                            | Version |
| ------------------------------- | ------- |
| [archive](./#provider\_archive) | 2.4.0   |
| [aws](./#provider\_aws)         | 5.22.0  |
| [null](./#provider\_null)       | 3.2.1   |

### Modules

No modules.

### Resources

| Name                                                                                                                                                                              | Type        |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [aws\_cloudwatch\_event\_rule.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_event\_rule)                      | resource    |
| [aws\_cloudwatch\_event\_target.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_event\_target)                  | resource    |
| [aws\_cloudwatch\_log\_group.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_log\_group)                        | resource    |
| [aws\_cloudwatch\_metric\_alarm.acm\_certificate\_expired](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)                 | resource    |
| [aws\_cloudwatch\_metric\_alarm.alb\_healthyhosts](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)                         | resource    |
| [aws\_cloudwatch\_metric\_alarm.cpu\_utilization\_too\_high](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)               | resource    |
| [aws\_cloudwatch\_metric\_alarm.lb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)                                        | resource    |
| [aws\_cloudwatch\_metric\_alarm.redis-current-connections](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)                 | resource    |
| [aws\_cloudwatch\_metric\_alarm.redis-database-memory-percentage](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)          | resource    |
| [aws\_cloudwatch\_metric\_alarm.redis-replication-lag](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)                     | resource    |
| [aws\_cloudwatch\_metric\_alarm.redis\_cpu](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)                                | resource    |
| [aws\_cloudwatch\_metric\_alarm.redis\_cpu\_engine\_utilization](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)           | resource    |
| [aws\_cloudwatch\_metric\_alarm.target\_response\_time](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_metric\_alarm)                    | resource    |
| [aws\_db\_event\_subscription.default](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db\_event\_subscription)                                       | resource    |
| [aws\_iam\_policy.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                              | resource    |
| [aws\_iam\_role.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role)                                                  | resource    |
| [aws\_iam\_role\_policy\_attachment.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)          | resource    |
| [aws\_iam\_role\_policy\_attachment.cron\_monitoring\_lambda\_managed](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment) | resource    |
| [aws\_lambda\_function.cron\_monitoring](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda\_function)                                            | resource    |
| [aws\_lambda\_permission.cron\_monitoring\_cloudwatch](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda\_permission)                            | resource    |
| [aws\_security\_group.cron\_monitoring](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security\_group)                                              | resource    |
| [aws\_security\_group\_rule.cron\_monitoring\_to\_rds](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security\_group\_rule)                         | resource    |
| [null\_resource.cron\_monitoring\_build](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource)                                                   | resource    |
| [archive\_file.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/archive/latest/docs/data-sources/file)                                                 | data source |
| [aws\_caller\_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller\_identity)                                                  | data source |
| [aws\_iam\_policy\_document.cron\_monitoring\_lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                       | data source |
| [aws\_iam\_policy\_document.cron\_monitoring\_lambda\_assume\_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)         | data source |
| [aws\_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region)                                                                      | data source |
| [aws\_secretsmanager\_secret.mysql\_database\_password](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/secretsmanager\_secret)                    | data source |

### Inputs

| Name                                                             | Description | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Default   | Required |
| ---------------------------------------------------------------- | ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------- | :------: |
| [acm\_certificate\_arn](./#input\_acm\_certificate\_arn)         | n/a         | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `null`    |    no    |
| [albs](./#input\_albs)                                           | n/a         | <pre><code>list(object({
    name                    = string
    arn_suffix              = string
    target_group_name       = string
    target_group_arn_suffix = string
    min_containers          = optional(string, 1)
    ecs_service_name        = string
  }))
</code></pre>                                                                                                                                                                                   | `[]`      |    no    |
| [cron\_monitoring](./#input\_cron\_monitoring)                   | n/a         | <pre><code>object({
    mysql_host                 = string
    mysql_database             = string
    mysql_user                 = string
    mysql_password_secret_name = string
    vpc_id                     = string
    subnet_ids                 = list(string)
    rds_security_group_id      = string
    delay_tolerance            = string
    run_interval               = string
    log_retention_in_days      = optional(number, 7)
  })
</code></pre> | `null`    |    no    |
| [customer\_prefix](./#input\_customer\_prefix)                   | n/a         | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `"fleet"` |    no    |
| [default\_sns\_topic\_arns](./#input\_default\_sns\_topic\_arns) | n/a         | `list(string)`                                                                                                                                                                                                                                                                                                                                                                                                                                                            | `[]`      |    no    |
| [fleet\_ecs\_service\_name](./#input\_fleet\_ecs\_service\_name) | n/a         | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `null`    |    no    |
| [mysql\_cluster\_members](./#input\_mysql\_cluster\_members)     | n/a         | `list(string)`                                                                                                                                                                                                                                                                                                                                                                                                                                                            | `[]`      |    no    |
| [redis\_cluster\_members](./#input\_redis\_cluster\_members)     | n/a         | `list(string)`                                                                                                                                                                                                                                                                                                                                                                                                                                                            | `[]`      |    no    |
| [sns\_topic\_arns\_map](./#input\_sns\_topic\_arns\_map)         | n/a         | `map(list(string))`                                                                                                                                                                                                                                                                                                                                                                                                                                                       | `{}`      |    no    |

### Outputs

No outputs.
