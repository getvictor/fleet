# byo-db

### Requirements

No requirements.

### Providers

| Name                    | Version |
| ----------------------- | ------- |
| [aws](./#provider\_aws) | n/a     |

### Modules

| Name                          | Source                        | Version |
| ----------------------------- | ----------------------------- | ------- |
| [alb](./#module\_alb)         | terraform-aws-modules/alb/aws | 8.2.1   |
| [cluster](./#module\_cluster) | terraform-aws-modules/ecs/aws | 4.1.2   |
| [ecs](./#module\_ecs)         | ./byo-ecs                     | n/a     |

### Resources

| Name                                                                                                                    | Type     |
| ----------------------------------------------------------------------------------------------------------------------- | -------- |
| [aws\_security\_group.alb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security\_group) | resource |

### Inputs

| Name                                             | Description                                                                                                                           | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | Required |
| ------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| [alb\_config](./#input\_alb\_config)             | n/a                                                                                                                                   | <pre><code>object({
    name                 = optional(string, "fleet")
    subnets              = list(string)
    security_groups      = optional(list(string), [])
    access_logs          = optional(map(string), {})
    certificate_arn      = string
    allowed_cidrs        = optional(list(string), ["0.0.0.0/0"])
    allowed_ipv6_cidrs   = optional(list(string), ["::/0"])
    egress_cidrs         = optional(list(string), ["0.0.0.0/0"])
    egress_ipv6_cidrs    = optional(list(string), ["::/0"])
    extra_target_groups  = optional(any, [])
    https_listener_rules = optional(any, [])
    tls_policy           = optional(string, "ELBSecurityPolicy-TLS-1-2-2017-01")
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    yes   |
| [ecs\_cluster](./#input\_ecs\_cluster)           | The config for the terraform-aws-modules/ecs/aws module                                                                               | <pre><code>object({
    autoscaling_capacity_providers = optional(any, {})
    cluster_configuration = optional(any, {
      execute_command_configuration = {
        logging = "OVERRIDE"
        log_configuration = {
          cloud_watch_log_group_name = "/aws/ecs/aws-ec2"
        }
      }
    })
    cluster_name = optional(string, "fleet")
    cluster_settings = optional(map(string), {
      "name" : "containerInsights",
      "value" : "enabled",
    })
    create                                = optional(bool, true)
    default_capacity_provider_use_fargate = optional(bool, true)
    fargate_capacity_providers = optional(any, {
      FARGATE = {
        default_capacity_provider_strategy = {
          weight = 100
        }
      }
      FARGATE_SPOT = {
        default_capacity_provider_strategy = {
          weight = 0
        }
      }
    })
    tags = optional(map(string))
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | <pre><code>{
  "autoscaling_capacity_providers": {},
  "cluster_configuration": {
    "execute_command_configuration": {
      "log_configuration": {
        "cloud_watch_log_group_name": "/aws/ecs/aws-ec2"
      },
      "logging": "OVERRIDE"
    }
  },
  "cluster_name": "fleet",
  "cluster_settings": {
    "name": "containerInsights",
    "value": "enabled"
  },
  "create": true,
  "default_capacity_provider_use_fargate": true,
  "fargate_capacity_providers": {
    "FARGATE": {
      "default_capacity_provider_strategy": {
        "weight": 100
      }
    },
    "FARGATE_SPOT": {
      "default_capacity_provider_strategy": {
        "weight": 0
      }
    }
  },
  "tags": {}
}
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |    no    |
| [fleet\_config](./#input\_fleet\_config)         | The configuration object for Fleet itself. Fields that default to null will have their respective resources created if not specified. | <pre><code>object({
    mem                          = optional(number, 4096)
    cpu                          = optional(number, 512)
    image                        = optional(string, "fleetdm/fleet:v4.39.0")
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
    database = optional(object({
      password_secret_arn = string
      user                = string
      database            = string
      address             = string
      rr_address          = optional(string, null)
      }), {
      password_secret_arn = null
      user                = null
      database            = null
      address             = null
      rr_address          = null
    })
    redis = optional(object({
      address = string
      use_tls = optional(bool, true)
      }), {
      address = null
      use_tls = true
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
    loadbalancer = optional(object({
      arn = string
      }), {
      arn = null
    })
    extra_load_balancers = optional(list(any), [])
    networking = optional(object({
      subnets         = list(string)
      security_groups = optional(list(string), null)
      }), {
      subnets         = null
      security_groups = null
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
  "extra_load_balancers": [],
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
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <pre><code>{
  "cpu": 1024,
  "mem": 2048
}
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |    no    |
| [vpc\_id](./#input\_vpc\_id)                     | n/a                                                                                                                                   | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    yes   |

### Outputs

| Name                          | Description |
| ----------------------------- | ----------- |
| [alb](./#output\_alb)         | n/a         |
| [byo-ecs](./#output\_byo-ecs) | n/a         |
| [cluster](./#output\_cluster) | n/a         |
