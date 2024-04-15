# Logging Destination: Firehose

This addon provides a Kinesis Firehose logging destination for Fleet.

## Requirements

No requirements.

## Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | 5.44.0  |

## Modules

No modules.

## Resources

| Name                                                                                                                                                                                                        | Type        |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [aws\_iam\_policy.firehose-logging](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                                                                | resource    |
| [aws\_iam\_policy.firehose-results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                                                                | resource    |
| [aws\_iam\_policy.firehose-status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                                                                 | resource    |
| [aws\_iam\_role.firehose-results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role)                                                                                    | resource    |
| [aws\_iam\_role.firehose-status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role)                                                                                     | resource    |
| [aws\_iam\_role\_policy\_attachment.firehose-results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)                                            | resource    |
| [aws\_iam\_role\_policy\_attachment.firehose-status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)                                             | resource    |
| [aws\_kinesis\_firehose\_delivery\_stream.osquery\_results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kinesis\_firehose\_delivery\_stream)                                | resource    |
| [aws\_kinesis\_firehose\_delivery\_stream.osquery\_status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kinesis\_firehose\_delivery\_stream)                                 | resource    |
| [aws\_s3\_bucket.osquery-results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket)                                                                                   | resource    |
| [aws\_s3\_bucket.osquery-status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket)                                                                                    | resource    |
| [aws\_s3\_bucket\_lifecycle\_configuration.osquery-results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_lifecycle\_configuration)                               | resource    |
| [aws\_s3\_bucket\_lifecycle\_configuration.osquery-status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_lifecycle\_configuration)                                | resource    |
| [aws\_s3\_bucket\_public\_access\_block.osquery-results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_public\_access\_block)                                     | resource    |
| [aws\_s3\_bucket\_public\_access\_block.osquery-status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_public\_access\_block)                                      | resource    |
| [aws\_s3\_bucket\_server\_side\_encryption\_configuration.osquery-results](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_server\_side\_encryption\_configuration) | resource    |
| [aws\_s3\_bucket\_server\_side\_encryption\_configuration.osquery-status](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_server\_side\_encryption\_configuration)  | resource    |
| [aws\_iam\_policy\_document.firehose-logging](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                                         | data source |
| [aws\_iam\_policy\_document.osquery\_firehose\_assume\_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                          | data source |
| [aws\_iam\_policy\_document.osquery\_results\_policy\_doc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                            | data source |
| [aws\_iam\_policy\_document.osquery\_status\_policy\_doc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                             | data source |
| [aws\_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region)                                                                                                | data source |

## Inputs

| Name                                                                   | Description | Type                                                                                                                                               | Default                                                                                     | Required |
| ---------------------------------------------------------------------- | ----------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | :------: |
| [compression\_format](./#input\_compression\_format)                   | n/a         | `string`                                                                                                                                           | `"UNCOMPRESSED"`                                                                            |    no    |
| [osquery\_results\_s3\_bucket](./#input\_osquery\_results\_s3\_bucket) | n/a         | <pre><code>object({
    name         = optional(string, "fleet-osquery-results-archive")
    expires_days = optional(number, 1)
  })
</code></pre> | <pre><code>{
  "expires_days": 1,
  "name": "fleet-osquery-results-archive"
}
</code></pre> |    no    |
| [osquery\_status\_s3\_bucket](./#input\_osquery\_status\_s3\_bucket)   | n/a         | <pre><code>object({
    name         = optional(string, "fleet-osquery-status-archive")
    expires_days = optional(number, 1)
  })
</code></pre>  | <pre><code>{
  "expires_days": 1,
  "name": "fleet-osquery-status-archive"
}
</code></pre>  |    no    |

## Outputs

| Name                                                                                    | Description |
| --------------------------------------------------------------------------------------- | ----------- |
| [fleet\_extra\_environment\_variables](./#output\_fleet\_extra\_environment\_variables) | n/a         |
| [fleet\_extra\_iam\_policies](./#output\_fleet\_extra\_iam\_policies)                   | n/a         |
