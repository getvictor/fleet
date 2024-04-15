# byo-vpc

### Requirements

No requirements.

### Providers

| Name                          | Version |
| ----------------------------- | ------- |
| [aws](./#provider\_aws)       | 5.23.1  |
| [random](./#provider\_random) | 3.5.1   |

### Modules

| Name                                              | Source                               | Version |
| ------------------------------------------------- | ------------------------------------ | ------- |
| [byo-db](./#module\_byo-db)                       | ./byo-db                             | n/a     |
| [rds](./#module\_rds)                             | terraform-aws-modules/rds-aurora/aws | 7.6.0   |
| [redis](./#module\_redis)                         | cloudposse/elasticache-redis/aws     | 0.53.0  |
| [secrets-manager-1](./#module\_secrets-manager-1) | lgallard/secrets-manager/aws         | 0.6.1   |

### Resources

| Name                                                                                                                                                   | Type        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------- |
| [aws\_db\_parameter\_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db\_parameter\_group)                     | resource    |
| [aws\_rds\_cluster\_parameter\_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/rds\_cluster\_parameter\_group) | resource    |
| [random\_password.rds](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password)                                        | resource    |
| [aws\_subnet.redis](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/subnet)                                             | data source |

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
    idle_timeout         = optional(number, 60)
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    yes   |
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
    image                        = optional(string, "fleetdm/fleet:v4.48.2")
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
| [rds\_config](./#input\_rds\_config)             | The config for the terraform-aws-modules/rds-aurora/aws module                                                                        | <pre><code>object({
    name                            = optional(string, "fleet")
    engine_version                  = optional(string, "8.0.mysql_aurora.3.02.2")
    instance_class                  = optional(string, "db.t4g.large")
    subnets                         = optional(list(string), [])
    allowed_security_groups         = optional(list(string), [])
    allowed_cidr_blocks             = optional(list(string), [])
    apply_immediately               = optional(bool, true)
    monitoring_interval             = optional(number, 10)
    db_parameter_group_name         = optional(string)
    db_parameters                   = optional(map(string), {})
    db_cluster_parameter_group_name = optional(string)
    db_cluster_parameters           = optional(map(string), {})
    enabled_cloudwatch_logs_exports = optional(list(string), [])
    master_username                 = optional(string, "fleet")
    snapshot_identifier             = optional(string)
    cluster_tags                    = optional(map(string), {})
    preferred_maintenance_window    = optional(string, "thu:23:00-fri:00:00")
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | <pre><code>{
  "allowed_cidr_blocks": [],
  "allowed_security_groups": [],
  "apply_immediately": true,
  "cluster_tags": {},
  "db_cluster_parameter_group_name": null,
  "db_cluster_parameters": {},
  "db_parameter_group_name": null,
  "db_parameters": {},
  "enabled_cloudwatch_logs_exports": [],
  "engine_version": "8.0.mysql_aurora.3.02.2",
  "instance_class": "db.t4g.large",
  "master_username": "fleet",
  "monitoring_interval": 10,
  "name": "fleet",
  "preferred_maintenance_window": "thu:23:00-fri:00:00",
  "snapshot_identifier": null,
  "subnets": []
}
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    no    |
| [redis\_config](./#input\_redis\_config)         | n/a                                                                                                                                   | <pre><code>object({
    name                          = optional(string, "fleet")
    replication_group_id          = optional(string)
    elasticache_subnet_group_name = optional(string, "")
    allowed_security_group_ids    = optional(list(string), [])
    subnets                       = list(string)
    allowed_cidrs                 = list(string)
    availability_zones            = optional(list(string), [])
    cluster_size                  = optional(number, 3)
    instance_type                 = optional(string, "cache.m5.large")
    apply_immediately             = optional(bool, true)
    automatic_failover_enabled    = optional(bool, false)
    engine_version                = optional(string, "6.x")
    family                        = optional(string, "redis6.x")
    at_rest_encryption_enabled    = optional(bool, true)
    transit_encryption_enabled    = optional(bool, true)
    parameter = optional(list(object({
      name  = string
      value = string
    })), [])
    log_delivery_configuration = optional(list(map(any)), [])
    tags                       = optional(map(string), {})
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | <pre><code>{
  "allowed_cidrs": null,
  "allowed_security_group_ids": [],
  "apply_immediately": true,
  "at_rest_encryption_enabled": true,
  "automatic_failover_enabled": false,
  "availability_zones": [],
  "cluster_size": 3,
  "elasticache_subnet_group_name": "",
  "engine_version": "6.x",
  "family": "redis6.x",
  "instance_type": "cache.m5.large",
  "log_delivery_configuration": [],
  "name": "fleet",
  "parameter": [],
  "replication_group_id": null,
  "subnets": null,
  "tags": {},
  "transit_encryption_enabled": true
}
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    no    |
| [vpc\_config](./#input\_vpc\_config)             | n/a                                                                                                                                   | <pre><code>object({
    vpc_id = string
    networking = object({
      subnets = list(string)
    })
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    yes   |

### Outputs

| Name                          | Description |
| ----------------------------- | ----------- |
| [byo-db](./#output\_byo-db)   | n/a         |
| [rds](./#output\_rds)         | n/a         |
| [redis](./#output\_redis)     | n/a         |
| [secrets](./#output\_secrets) | n/a         |
