# Firehose Logging Destination Setup

In this Terraform code, we are defining an IAM Role named `fleet_role` in our AWS Account, that will be assumed by the Fleet application we are hosting. We are only allowing this specific IAM Role (identified by its ARN) to perform certain actions on the Firehose service, such as `DescribeDeliveryStream`, `PutRecord`, and `PutRecordBatch`.

The reason we need a local IAM role in your account is so that we can assume role into it, and you have full control over the permissions it has. The associated IAM policy in the same file specifies the minimum allowed permissions.

The Firehose service is KMS encrypted, so the IAM Role we assume into needs permission to the KMS key that is being used to encrypt the data going into Firehose. Additionally, if the data is being delivered to S3, it will also be encrypted with KMS using the AWS S3 KMS key that is managed by AWS. This is because only customer managed keys can be shared across accounts, and the Firehose delivery stream is actually the one writing to S3.

This code sets up a secure and controlled environment for the Fleet application to perform its necessary actions on the Firehose service within your AWS Account.

## Requirements

| Name                                   | Version   |
| -------------------------------------- | --------- |
| [terraform](./#requirement\_terraform) | >= 1.3.7  |
| [aws](./#requirement\_aws)             | >= 5.29.0 |

## Providers

| Name                    | Version   |
| ----------------------- | --------- |
| [aws](./#provider\_aws) | >= 5.29.0 |

## Modules

No modules.

## Resources

| Name                                                                                                                                                                                                    | Type        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [aws\_iam\_policy.firehose](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                                                                    | resource    |
| [aws\_iam\_policy.fleet\_firehose](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                                                             | resource    |
| [aws\_iam\_role.firehose](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role)                                                                                        | resource    |
| [aws\_iam\_role.fleet\_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role)                                                                                     | resource    |
| [aws\_iam\_role\_policy\_attachment.firehose](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)                                                | resource    |
| [aws\_iam\_role\_policy\_attachment.fleet\_firehose](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)                                         | resource    |
| [aws\_kinesis\_firehose\_delivery\_stream.fleet\_log\_destinations](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kinesis\_firehose\_delivery\_stream)                    | resource    |
| [aws\_kms\_key.firehose\_key](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms\_key)                                                                                     | resource    |
| [aws\_s3\_bucket.destination](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket)                                                                                   | resource    |
| [aws\_s3\_bucket\_public\_access\_block.destination](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_public\_access\_block)                                     | resource    |
| [aws\_s3\_bucket\_server\_side\_encryption\_configuration.destination](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket\_server\_side\_encryption\_configuration) | resource    |
| [aws\_caller\_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller\_identity)                                                                        | data source |
| [aws\_iam\_policy\_document.assume\_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                                         | data source |
| [aws\_iam\_policy\_document.firehose](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                                             | data source |
| [aws\_iam\_policy\_document.firehose\_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                                     | data source |
| [aws\_iam\_policy\_document.osquery\_firehose\_assume\_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                                      | data source |
| [aws\_kms\_alias.s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/kms\_alias)                                                                                         | data source |
| [aws\_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region)                                                                                            | data source |

## Inputs

| Name                                                                                                 | Description                                                                                                               | Type                                                                                                                                                                                                                                                          | Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Required |
| ---------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| [fleet\_iam\_role\_arn](./#input\_fleet\_iam\_role\_arn)                                             | the arn of the fleet role that firehose will assume to write data to your bucket                                          | `string`                                                                                                                                                                                                                                                      | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |    yes   |
| [kms\_key\_arn](./#input\_kms\_key\_arn)                                                             | An optional KMS key ARN for server-side encryption. If not provided and encryption is enabled, a new key will be created. | `string`                                                                                                                                                                                                                                                      | `""`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |    no    |
| [log\_destinations](./#input\_log\_destinations)                                                     | A map of configurations for Firehose delivery streams.                                                                    | <pre><code>map(object({
    name                  = string
    prefix                = string
    error_output_prefix   = string
    buffering_size        = number
    buffering_interval    = number
    compression_format    = string
  }))
</code></pre> | <pre><code>{
  "audit": {
    "buffering_interval": 120,
    "buffering_size": 20,
    "compression_format": "UNCOMPRESSED",
    "error_output_prefix": "audit/error/error=!{firehose:error-output-type}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/",
    "name": "fleet_audit",
    "prefix": "audit/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/"
  },
  "results": {
    "buffering_interval": 120,
    "buffering_size": 20,
    "compression_format": "UNCOMPRESSED",
    "error_output_prefix": "results/error/error=!{firehose:error-output-type}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/",
    "name": "osquery_results",
    "prefix": "results/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/"
  },
  "status": {
    "buffering_interval": 120,
    "buffering_size": 20,
    "compression_format": "UNCOMPRESSED",
    "error_output_prefix": "status/error/error=!{firehose:error-output-type}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/",
    "name": "osquery_status",
    "prefix": "status/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/"
  }
}
</code></pre> |    no    |
| [osquery\_logging\_destination\_bucket\_name](./#input\_osquery\_logging\_destination\_bucket\_name) | name of the bucket to store osquery results & status logs                                                                 | `string`                                                                                                                                                                                                                                                      | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |    yes   |
| [server\_side\_encryption\_enabled](./#input\_server\_side\_encryption\_enabled)                     | A boolean flag to enable/disable server-side encryption. Defaults to true (enabled).                                      | `bool`                                                                                                                                                                                                                                                        | `true`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    no    |

## Outputs

| Name                                                  | Description                              |
| ----------------------------------------------------- | ---------------------------------------- |
| [firehose\_iam\_role](./#output\_firehose\_iam\_role) | n/a                                      |
| [log\_destinations](./#output\_log\_destinations)     | Map of Firehose delivery streams' names. |
| [s3\_destination](./#output\_s3\_destination)         | n/a                                      |
