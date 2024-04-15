# byo-ecs

### Requirements

No requirements.

### Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | 5.14.0  |

### Modules

No modules.

### Resources

| Name                                                                                                                                                              | Type        |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| [aws\_appautoscaling\_policy.ecs\_policy\_cpu](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/appautoscaling\_policy)                | resource    |
| [aws\_appautoscaling\_policy.ecs\_policy\_memory](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/appautoscaling\_policy)             | resource    |
| [aws\_appautoscaling\_target.ecs\_target](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/appautoscaling\_target)                     | resource    |
| [aws\_cloudwatch\_log\_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch\_log\_group)                            | resource    |
| [aws\_ecs\_service.fleet](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_service)                                               | resource    |
| [aws\_ecs\_task\_definition.backend](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs\_task\_definition)                           | resource    |
| [aws\_iam\_policy.execution](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                             | resource    |
| [aws\_iam\_policy.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_policy)                                                  | resource    |
| [aws\_iam\_role.execution](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role)                                                 | resource    |
| [aws\_iam\_role.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role)                                                      | resource    |
| [aws\_iam\_role\_policy\_attachment.execution](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)         | resource    |
| [aws\_iam\_role\_policy\_attachment.execution\_extras](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment) | resource    |
| [aws\_iam\_role\_policy\_attachment.extras](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)            | resource    |
| [aws\_iam\_role\_policy\_attachment.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)              | resource    |
| [aws\_iam\_role\_policy\_attachment.role\_attachment](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam\_role\_policy\_attachment)  | resource    |
| [aws\_security\_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security\_group)                                          | resource    |
| [aws\_iam\_policy\_document.assume\_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                   | data source |
| [aws\_iam\_policy\_document.fleet](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                          | data source |
| [aws\_iam\_policy\_document.fleet-execution](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam\_policy\_document)                | data source |
| [aws\_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region)                                                      | data source |

### Inputs

| Name                                             | Description                                                                                                                           | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | Required |
| ------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| [ecs\_cluster](./#input\_ecs\_cluster)           | The name of the ECS cluster to use                                                                                                    | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    yes   |
| [fleet\_config](./#input\_fleet\_config)         | The configuration object for Fleet itself. Fields that default to null will have their respective resources created if not specified. | <pre><code>object({
    mem                          = optional(number, 4096)
    cpu                          = optional(number, 512)
    image                        = optional(string, "fleetdm/fleet:v4.42.0")
    family                       = optional(string, "fleet")
    sidecars                     = optional(list(any), [])
    depends_on                   = optional(list(any), [])
    mount_points                 = optional(list(any), [])
    volumes                      = optional(list(any), [])
    extra_environment_variables  = optional(map(string), {})
    extra_iam_policies           = optional(list(string), [])
    extra_execution_iam_policies = optional(list(string), [])
    extra_secrets                = optional(map(string), {})
    security_groups              = optional(list(string), null)
    security_group_name          = optional(string, "fleet")
    iam_role_arn                 = optional(string, null)
    service = optional(object({
      name = optional(string, "fleet")
      }), {
      name = "fleet"
    })
    database = object({
      password_secret_arn = string
      user                = string
      database            = string
      address             = string
      rr_address          = optional(string, null)
    })
    redis = object({
      address = string
      use_tls = optional(bool, true)
    })
    awslogs = optional(object({
      name      = optional(string, null)
      region    = optional(string, null)
      create    = optional(bool, true)
      prefix    = optional(string, "fleet")
      retention = optional(number, 5)
      }), {
      name      = null
      region    = null
      prefix    = "fleet"
      retention = 5
    })
    loadbalancer = object({
      arn = string
    })
    extra_load_balancers = optional(list(any), [])
    networking = object({
      subnets         = list(string)
      security_groups = optional(list(string), null)
    })
    autoscaling = optional(object({
      max_capacity                 = optional(number, 5)
      min_capacity                 = optional(number, 1)
      memory_tracking_target_value = optional(number, 80)
      cpu_tracking_target_value    = optional(number, 80)
      }), {
      max_capacity                 = 5
      min_capacity                 = 1
      memory_tracking_target_value = 80
      cpu_tracking_target_value    = 80
    })
    iam = optional(object({
      role = optional(object({
        name        = optional(string, "fleet-role")
        policy_name = optional(string, "fleet-iam-policy")
        }), {
        name        = "fleet-role"
        policy_name = "fleet-iam-policy"
      })
      execution = optional(object({
        name        = optional(string, "fleet-execution-role")
        policy_name = optional(string, "fleet-execution-role")
        }), {
        name        = "fleet-execution-role"
        policy_name = "fleet-iam-policy-execution"
      })
      }), {
      name = "fleetdm-execution-role"
    })
  })
</code></pre> | <pre><code>{
  "autoscaling": {
    "cpu_tracking_target_value": 80,
    "max_capacity": 5,
    "memory_tracking_target_value": 80,
    "min_capacity": 1
  },
  "awslogs": {
    "create": true,
    "name": null,
    "prefix": "fleet",
    "region": null,
    "retention": 5
  },
  "cpu": 256,
  "database": {
    "address": null,
    "database": null,
    "password_secret_arn": null,
    "rr_address": null,
    "user": null
  },
  "depends_on": [],
  "extra_environment_variables": {},
  "extra_execution_iam_policies": [],
  "extra_iam_policies": [],
  "extra_load_balacners": [],
  "extra_secrets": {},
  "family": "fleet",
  "iam": {
    "execution": {
      "name": "fleet-execution-role",
      "policy_name": "fleet-iam-policy-execution"
    },
    "role": {
      "name": "fleet-role",
      "policy_name": "fleet-iam-policy"
    }
  },
  "iam_role_arn": null,
  "image": "fleetdm/fleet:v4.31.1",
  "loadbalancer": {
    "arn": null
  },
  "mem": 512,
  "mount_points": [],
  "networking": {
    "security_groups": null,
    "subnets": null
  },
  "redis": {
    "address": null,
    "use_tls": true
  },
  "security_group_name": "fleet",
  "security_groups": null,
  "service": {
    "name": "fleet"
  },
  "sidecars": [],
  "volumes": []
}
</code></pre> |    no    |
| [migration\_config](./#input\_migration\_config) | The configuration object for Fleet's migration task.                                                                                  | <pre><code>object({
    mem = number
    cpu = number
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | <pre><code>{
  "cpu": 1024,
  "mem": 2048
}
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |    no    |
| [vpc\_id](./#input\_vpc\_id)                     | n/a                                                                                                                                   | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | `null`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |    no    |

### Outputs

| Name                                                              | Description |
| ----------------------------------------------------------------- | ----------- |
| [appautoscaling\_target](./#output\_appautoscaling\_target)       | n/a         |
| [execution\_iam\_role\_arn](./#output\_execution\_iam\_role\_arn) | n/a         |
| [iam\_role\_arn](./#output\_iam\_role\_arn)                       | n/a         |
| [logging\_config](./#output\_logging\_config)                     | n/a         |
| [non\_circular](./#output\_non\_circular)                         | n/a         |
| [service](./#output\_service)                                     | n/a         |
| [task\_definition](./#output\_task\_definition)                   | n/a         |
