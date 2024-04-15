# External Vulnerability Scans addon

This addon creates an additional ECS service that only runs a single task, responsible for vuln processing. It receives no web traffic. We utilize [current instance checks](https://fleetdm.com/docs/configuration/fleet-server-configuration#current-instance-checks) to make this happen. The advantages of this mechanism:

1. dedicating processing power to vuln processing 2. ensures task responsible for vuln processing isn't also trying to serve web traffic
2. caching of vulnerability artifacts/dependencies

Usage is simplified by using the output from the fleet byo-ecs module (../terraform/byo-vpc/byo-db/byo-ecs/README.md)

## Requirements

No requirements.

## Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | 5.11.0  |

## Modules

No modules.

## Resources

| Name                                                                                                                                            | Type        |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [aws\_ecs\_service.fleet](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_service)                             | resource    |
| [aws\_ecs\_task\_definition.vuln-processing](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_task\_definition) | resource    |
| [aws\_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region)                                    | data source |

## Inputs

| Name                                                             | Description                                                                         | Type                                                                                               | Default   | Required |
| ---------------------------------------------------------------- | ----------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | --------- | :------: |
| [awslogs\_config](./#input\_awslogs\_config)                     | n/a                                                                                 | <pre><code>object({
    group  = string
    region = string
    prefix = string
  })
</code></pre> | n/a       |    yes   |
| [customer\_prefix](./#input\_customer\_prefix)                   | n/a                                                                                 | `string`                                                                                           | `"fleet"` |    no    |
| [ecs\_cluster](./#input\_ecs\_cluster)                           | The ecs cluster module that is created by the byo-db module                         | `any`                                                                                              | n/a       |    yes   |
| [execution\_iam\_role\_arn](./#input\_execution\_iam\_role\_arn) | The ARN of the fleet execution role, this is necessary to pass role from ecs events | `any`                                                                                              | n/a       |    yes   |
| [fleet\_config](./#input\_fleet\_config)                         | The root Fleet config object                                                        | `any`                                                                                              | n/a       |    yes   |
| [security\_groups](./#input\_security\_groups)                   | n/a                                                                                 | `list(string)`                                                                                     | n/a       |    yes   |
| [subnets](./#input\_subnets)                                     | n/a                                                                                 | `list(string)`                                                                                     | n/a       |    yes   |
| [task\_role\_arn](./#input\_task\_role\_arn)                     | The ARN of the fleet task role, this is necessary to pass role from ecs events      | `any`                                                                                              | n/a       |    yes   |
| [vuln\_processing\_cpu](./#input\_vuln\_processing\_cpu)         | The amount of CPU to dedicate to the vuln processing command                        | `number`                                                                                           | `1024`    |    no    |
| [vuln\_processing\_memory](./#input\_vuln\_processing\_memory)   | The amount of memory to dedicate to the vuln processing command                     | `number`                                                                                           | `4096`    |    no    |

## Outputs

| Name                                                                      | Description |
| ------------------------------------------------------------------------- | ----------- |
| [extra\_environment\_variables](./#output\_extra\_environment\_variables) | n/a         |
