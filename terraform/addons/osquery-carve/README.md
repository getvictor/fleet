# Osquery Carve Bucket Addon

This addon provides a S3 bucket for Osquery Carve results.

## Requirements

No requirements.

## Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | 5.39.1  |

## Modules

No modules.

## Resources

| Name                                                                                                                                                                                             | Type        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------- |
| [aws\_iam\_policy.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                                                                 | resource    |
| [aws\_s3\_bucket.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket)                                                                                   | resource    |
| [aws\_s3\_bucket\_lifecycle\_configuration.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_lifecycle\_configuration)                               | resource    |
| [aws\_s3\_bucket\_public\_access\_block.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_public\_access\_block)                                     | resource    |
| [aws\_s3\_bucket\_server\_side\_encryption\_configuration.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_server\_side\_encryption\_configuration) | resource    |
| [aws\_iam\_policy\_document.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                                          | data source |

## Inputs

| Name                                                               | Description | Type                                                                                                                                               | Default                                                                                     | Required |
| ------------------------------------------------------------------ | ----------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | :------: |
| [osquery\_carve\_s3\_bucket](./#input\_osquery\_carve\_s3\_bucket) | n/a         | <pre><code>object({
    name         = optional(string, "fleet-osquery-results-archive")
    expires_days = optional(number, 1)
  })
</code></pre> | <pre><code>{
  "expires_days": 1,
  "name": "fleet-osquery-results-archive"
}
</code></pre> |    no    |

## Outputs

| Name                                                                                    | Description |
| --------------------------------------------------------------------------------------- | ----------- |
| [fleet\_extra\_environment\_variables](./#output\_fleet\_extra\_environment\_variables) | n/a         |
| [fleet\_extra\_iam\_policies](./#output\_fleet\_extra\_iam\_policies)                   | n/a         |
