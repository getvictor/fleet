# saml-auth-proxy

### Requirements

No requirements.

### Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | 5.17.0  |

### Modules

| Name                                                        | Source                        | Version |
| ----------------------------------------------------------- | ----------------------------- | ------- |
| [saml\_auth\_proxy\_alb](./#module\_saml\_auth\_proxy\_alb) | terraform-aws-modules/alb/aws | 8.2.1   |

### Resources

| Name                                                                                                                                                      | Type        |
| --------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [aws\_cloudwatch\_log\_group.saml\_auth\_proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_log\_group)       | resource    |
| [aws\_ecs\_service.saml\_auth\_proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_service)                           | resource    |
| [aws\_ecs\_task\_definition.saml\_auth\_proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_task\_definition)         | resource    |
| [aws\_iam\_policy.saml\_auth\_proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                             | resource    |
| [aws\_secretsmanager\_secret.saml\_auth\_proxy\_cert](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager\_secret) | resource    |
| [aws\_security\_group.saml\_auth\_proxy\_alb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security\_group)                | resource    |
| [aws\_security\_group.saml\_auth\_proxy\_service](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security\_group)            | resource    |
| [aws\_iam\_policy\_document.saml\_auth\_proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)      | data source |
| [aws\_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region)                                              | data source |

### Inputs

| Name                                                                       | Description                                       | Type                                                                                                                                            | Default                                                                                                 | Required |
| -------------------------------------------------------------------------- | ------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | :------: |
| [alb\_target\_group\_arn](./#input\_alb\_target\_group\_arn)               | n/a                                               | `string`                                                                                                                                        | n/a                                                                                                     |    yes   |
| [base\_url](./#input\_base\_url)                                           | n/a                                               | `string`                                                                                                                                        | n/a                                                                                                     |    yes   |
| [cookie\_max\_age](./#input\_cookie\_max\_age)                             | n/a                                               | `string`                                                                                                                                        | `"1h"`                                                                                                  |    no    |
| [customer\_prefix](./#input\_customer\_prefix)                             | customer prefix to use to namespace all resources | `string`                                                                                                                                        | `"fleet"`                                                                                               |    no    |
| [ecs\_cluster](./#input\_ecs\_cluster)                                     | n/a                                               | `string`                                                                                                                                        | n/a                                                                                                     |    yes   |
| [ecs\_execution\_iam\_role\_arn](./#input\_ecs\_execution\_iam\_role\_arn) | n/a                                               | `string`                                                                                                                                        | n/a                                                                                                     |    yes   |
| [ecs\_iam\_role\_arn](./#input\_ecs\_iam\_role\_arn)                       | n/a                                               | `string`                                                                                                                                        | n/a                                                                                                     |    yes   |
| [idp\_metadata\_url](./#input\_idp\_metadata\_url)                         | n/a                                               | `string`                                                                                                                                        | n/a                                                                                                     |    yes   |
| [logging\_options](./#input\_logging\_options)                             | n/a                                               | <pre><code>object({
    awslogs-group         = string
    awslogs-region        = string
    awslogs-stream-prefix = string
  })
</code></pre> | n/a                                                                                                     |    yes   |
| [proxy\_containers](./#input\_proxy\_containers)                           | n/a                                               | `number`                                                                                                                                        | `1`                                                                                                     |    no    |
| [saml\_auth\_proxy\_image](./#input\_saml\_auth\_proxy\_image)             | n/a                                               | `string`                                                                                                                                        | `"itzg/saml-auth-proxy:1.12.0@sha256:ddff17caa00c1aad64d6c7b2e1d5eb93d97321c34d8ad12a25cfd8ce203db723"` |    no    |
| [security\_groups](./#input\_security\_groups)                             | n/a                                               | `list(string)`                                                                                                                                  | n/a                                                                                                     |    yes   |
| [subnets](./#input\_subnets)                                               | n/a                                               | `list(string)`                                                                                                                                  | n/a                                                                                                     |    yes   |
| [vpc\_id](./#input\_vpc\_id)                                               | n/a                                               | `string`                                                                                                                                        | n/a                                                                                                     |    yes   |

### Outputs

| Name                                                                              | Description                     |
| --------------------------------------------------------------------------------- | ------------------------------- |
| [fleet\_extra\_execution\_policies](./#output\_fleet\_extra\_execution\_policies) | n/a                             |
| [lb](./#output\_lb)                                                               | n/a                             |
| [lb\_target\_group\_arn](./#output\_lb\_target\_group\_arn)                       | Keep for legacy support for now |
| [name](./#output\_name)                                                           | n/a                             |
| [secretsmanager\_secret\_id](./#output\_secretsmanager\_secret\_id)               | n/a                             |
