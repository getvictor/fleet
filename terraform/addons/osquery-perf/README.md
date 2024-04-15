# osquery-perf addon

This addon adds osquery-perf hosts to the Fleet installation. These are generally used for loadtesting or other testing purposes. See https://github.com/fleetdm/fleet/tree/main/cmd/osquery-perf to learn more about osquery-perf itself.

This addon creates an AWS Secrets Manager secret that will be used to store the enroll secret that the osquery-perf hosts use to enroll into Fleet. This secret will need to have its `SecretString` populated with the enroll secret manually once everything is setup in order for the osquery-perf hosts to connect.

Below is an example implementation of the module:

```
module "osquery_perf" {
  source                     = "github.com/fleetdm/fleet//terraform/addons/osquery-perf?ref=main"
  customer_prefix            = "fleet"
  ecs_cluster                = module.main.byo-vpc.byo-db.byo-ecs.service.cluster
  subnets                    = module.main.byo-vpc.byo-db.byo-ecs.service.network_configuration[0].subnets
  security_groups            = module.main.byo-vpc.byo-db.byo-ecs.service.network_configuration[0].security_groups
  ecs_iam_role_arn           = module.main.byo-vpc.byo-db.byo-ecs.iam_role_arn
  ecs_execution_iam_role_arn = module.main.byo-vpc.byo-db.byo-ecs.execution_iam_role_arn
  server_url                 = "https://${aws_route53_record.main.fqdn}"
  osquery_perf_image         = local.osquery_perf_image
  extra_flags                = ["--os_templates", "mac10.14.6,ubuntu_22.04,windows_11"]
  logging_options            = module.main.byo-vpc.byo-db.byo-ecs.logging_config
}
```

## Requirements

No requirements.

## Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | n/a     |

## Modules

No modules.

## Resources

| Name                                                                                                                                                                  | Type        |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [aws\_ecs\_service.osquery\_perf](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_service)                                           | resource    |
| [aws\_ecs\_task\_definition.osquery\_perf](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_task\_definition)                         | resource    |
| [aws\_kms\_alias.enroll\_secret](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms\_alias)                                              | resource    |
| [aws\_kms\_key.enroll\_secret](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms\_key)                                                  | resource    |
| [aws\_secretsmanager\_secret.enroll\_secret](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager\_secret)                      | resource    |
| [aws\_secretsmanager\_secret\_version.enroll\_secret](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/secretsmanager\_secret\_version) | data source |

## Inputs

| Name                                                                       | Description                                       | Type                                                                                                                                            | Default   | Required |
| -------------------------------------------------------------------------- | ------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | --------- | :------: |
| [customer\_prefix](./#input\_customer\_prefix)                             | customer prefix to use to namespace all resources | `string`                                                                                                                                        | `"fleet"` |    no    |
| [ecs\_cluster](./#input\_ecs\_cluster)                                     | n/a                                               | `string`                                                                                                                                        | n/a       |    yes   |
| [ecs\_execution\_iam\_role\_arn](./#input\_ecs\_execution\_iam\_role\_arn) | n/a                                               | `string`                                                                                                                                        | n/a       |    yes   |
| [ecs\_iam\_role\_arn](./#input\_ecs\_iam\_role\_arn)                       | n/a                                               | `string`                                                                                                                                        | n/a       |    yes   |
| [extra\_flags](./#input\_extra\_flags)                                     | n/a                                               | `list(string)`                                                                                                                                  | `[]`      |    no    |
| [loadtest\_containers](./#input\_loadtest\_containers)                     | n/a                                               | `number`                                                                                                                                        | `1`       |    no    |
| [logging\_options](./#input\_logging\_options)                             | n/a                                               | <pre><code>object({
    awslogs-group         = string
    awslogs-region        = string
    awslogs-stream-prefix = string
  })
</code></pre> | n/a       |    yes   |
| [osquery\_perf\_image](./#input\_osquery\_perf\_image)                     | n/a                                               | `string`                                                                                                                                        | n/a       |    yes   |
| [security\_groups](./#input\_security\_groups)                             | n/a                                               | `list(string)`                                                                                                                                  | n/a       |    yes   |
| [server\_url](./#input\_server\_url)                                       | n/a                                               | `string`                                                                                                                                        | n/a       |    yes   |
| [subnets](./#input\_subnets)                                               | n/a                                               | `list(string)`                                                                                                                                  | n/a       |    yes   |

## Outputs

No outputs.
