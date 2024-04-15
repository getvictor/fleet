# WAF for ALB addon

This addon creates and manages WAF attached to an ALB

## Requirements

No requirements.

## Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | n/a     |

## Modules

No modules.

## Resources

| Name                                                                                                                                               | Type     |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| [aws\_wafv2\_ip\_set.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/wafv2\_ip\_set)                             | resource |
| [aws\_wafv2\_rule\_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/wafv2\_rule\_group)                     | resource |
| [aws\_wafv2\_web\_acl.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/wafv2\_web\_acl)                           | resource |
| [aws\_wafv2\_web\_acl\_association.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/wafv2\_web\_acl\_association) | resource |

## Inputs

| Name                                               | Description | Type           | Default                                                                                                                                                     | Required |
| -------------------------------------------------- | ----------- | -------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| [blocked\_addresses](./#input\_blocked\_addresses) | n/a         | `list(string)` | `[]`                                                                                                                                                        |    no    |
| [blocked\_countries](./#input\_blocked\_countries) | n/a         | `list(string)` | <pre><code>[
  "BI",
  "BY",
  "CD",
  "CF",
  "CU",
  "IQ",
  "IR",
  "LB",
  "LY",
  "SD",
  "SO",
  "SS",
  "SY",
  "VE",
  "ZW",
  "RU"
]
</code></pre> |    no    |
| [lb\_arn](./#input\_lb\_arn)                       | n/a         | `any`          | n/a                                                                                                                                                         |    yes   |
| [name](./#input\_name)                             | n/a         | `any`          | n/a                                                                                                                                                         |    yes   |

## Outputs

No outputs.
