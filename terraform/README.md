# terraform

This module provides a basic Fleet setup. This assumes that you bring nothing to the installation. If you want to bring your own VPC/database/cache nodes/ECS cluster, then use one of the submodules provided.

To quickly list all available module versions you can run:

```shell
git tag |grep '^tf'
```

The following is the module layout, so you can navigate to the module that you want:

* Root module (use this to get a Fleet instance ASAP with minimal setup)
  * BYO-VPC (use this if you want to install Fleet inside an existing VPC)
    * BYO-database (use this if you want to use an existing database and cache node)
      * BYO-ECS (use this if you want to bring your own everything but Fleet ECS services)

## Migrating from existing Dogfood code

The below code describes how to migrate from existing Dogfood code

```hcl
moved {
  from = module.vpc
  to   = module.main.module.vpc
}

moved {
  from = module.aurora_mysql
  to = module.main.module.byo-vpc.module.rds
}

moved {
  from = aws_elasticache_replication_group.default
  to = module.main.module.byo-vpc.module.redis.aws_elasticache_replication_group.default
}
```

This focuses on the resources that are "heavy" or store data. Note that the ALB cannot be moved like this because Dogfood uses the `aws_alb` resource and the module uses the `aws_lb` resource. The resources are aliases of eachother, but Terraform can't recognize that.

## How to improve this module

If this module somehow doesn't fit your needs, feel free to contact us by opening a ticket, or contacting your contact at Fleet. Our goal is to make this module fit all needs within AWS, so we will try to find a solution so that this module fits your needs.

If you want to make the changes yourself, simply make a PR into main with your additions. We would ask that you make sure that variables are defined as null if there is no default that makes sense and that variable changes are reflected all the way up the stack.

## How to update this readme

Edit .header.md and run `terraform-docs markdown . > README.md`

### Requirements

| Name                                   | Version  |
| -------------------------------------- | -------- |
| [terraform](./#requirement\_terraform) | >= 1.3.8 |

### Providers

No providers.

### Modules

| Name                          | Source                        | Version |
| ----------------------------- | ----------------------------- | ------- |
| [byo-vpc](./#module\_byo-vpc) | ./byo-vpc                     | n/a     |
| [vpc](./#module\_vpc)         | terraform-aws-modules/vpc/aws | 5.1.2   |

### Resources

No resources.

### Inputs

| Name                                             | Description                                                                                                                           | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | Required |
| ------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| [alb\_config](./#input\_alb\_config)             | n/a                                                                                                                                   | <pre><code>object({
    name                 = optional(string, "fleet")
    security_groups      = optional(list(string), [])
    access_logs          = optional(map(string), {})
    allowed_cidrs        = optional(list(string), ["0.0.0.0/0"])
    allowed_ipv6_cidrs   = optional(list(string), ["::/0"])
    egress_cidrs         = optional(list(string), ["0.0.0.0/0"])
    egress_ipv6_cidrs    = optional(list(string), ["::/0"])
    extra_target_groups  = optional(any, [])
    https_listener_rules = optional(any, [])
    tls_policy           = optional(string, "ELBSecurityPolicy-TLS-1-2-2017-01")
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `{}`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |    no    |
| [certificate\_arn](./#input\_certificate\_arn)   | n/a                                                                                                                                   | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | n/a                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |    yes   |
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
    image                        = optional(string, "fleetdm/fleet:v4.40.0")
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
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | <pre><code>{
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
  "snapshot_identifier": null,
  "subnets": []
}
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |    no    |
| [redis\_config](./#input\_redis\_config)         | n/a                                                                                                                                   | <pre><code>object({
    name                          = optional(string, "fleet")
    replication_group_id          = optional(string)
    elasticache_subnet_group_name = optional(string)
    allowed_security_group_ids    = optional(list(string), [])
    subnets                       = optional(list(string))
    availability_zones            = optional(list(string))
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
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <pre><code>{
  "allowed_security_group_ids": [],
  "apply_immediately": true,
  "at_rest_encryption_enabled": true,
  "automatic_failover_enabled": false,
  "availability_zones": null,
  "cluster_size": 3,
  "elasticache_subnet_group_name": null,
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
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |    no    |
| [vpc](./#input\_vpc)                             | n/a                                                                                                                                   | <pre><code>object({
    name                = optional(string, "fleet")
    cidr                = optional(string, "10.10.0.0/16")
    azs                 = optional(list(string), ["us-east-2a", "us-east-2b", "us-east-2c"])
    private_subnets     = optional(list(string), ["10.10.1.0/24", "10.10.2.0/24", "10.10.3.0/24"])
    public_subnets      = optional(list(string), ["10.10.11.0/24", "10.10.12.0/24", "10.10.13.0/24"])
    database_subnets    = optional(list(string), ["10.10.21.0/24", "10.10.22.0/24", "10.10.23.0/24"])
    elasticache_subnets = optional(list(string), ["10.10.31.0/24", "10.10.32.0/24", "10.10.33.0/24"])

    create_database_subnet_group              = optional(bool, false)
    create_database_subnet_route_table        = optional(bool, true)
    create_elasticache_subnet_group           = optional(bool, true)
    create_elasticache_subnet_route_table     = optional(bool, true)
    enable_vpn_gateway                        = optional(bool, false)
    one_nat_gateway_per_az                    = optional(bool, false)
    single_nat_gateway                        = optional(bool, true)
    enable_nat_gateway                        = optional(bool, true)
    enable_dns_hostnames                      = optional(bool, false)
    enable_dns_support                        = optional(bool, true)
    enable_flow_log                           = optional(bool, false)
    create_flow_log_cloudwatch_log_group      = optional(bool, false)
    create_flow_log_cloudwatch_iam_role       = optional(bool, false)
    flow_log_max_aggregation_interval         = optional(number, 600)
    flow_log_cloudwatch_log_group_name_prefix = optional(string, "/aws/vpc-flow-log/")
    flow_log_cloudwatch_log_group_name_suffix = optional(string, "")
    vpc_flow_log_tags                         = optional(map(string), {})
  })
</code></pre>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | <pre><code>{
  "azs": [
    "us-east-2a",
    "us-east-2b",
    "us-east-2c"
  ],
  "cidr": "10.10.0.0/16",
  "create_database_subnet_group": false,
  "create_database_subnet_route_table": true,
  "create_elasticache_subnet_group": true,
  "create_elasticache_subnet_route_table": true,
  "create_flow_log_cloudwatch_iam_role": false,
  "create_flow_log_cloudwatch_log_group": false,
  "database_subnets": [
    "10.10.21.0/24",
    "10.10.22.0/24",
    "10.10.23.0/24"
  ],
  "elasticache_subnets": [
    "10.10.31.0/24",
    "10.10.32.0/24",
    "10.10.33.0/24"
  ],
  "enable_dns_hostnames": false,
  "enable_dns_support": true,
  "enable_flow_log": false,
  "enable_nat_gateway": true,
  "enable_vpn_gateway": false,
  "flow_log_cloudwatch_log_group_name_prefix": "/aws/vpc-flow-log/",
  "flow_log_cloudwatch_log_group_name_suffix": "",
  "flow_log_max_aggregation_interval": 600,
  "name": "fleet",
  "one_nat_gateway_per_az": false,
  "private_subnets": [
    "10.10.1.0/24",
    "10.10.2.0/24",
    "10.10.3.0/24"
  ],
  "public_subnets": [
    "10.10.11.0/24",
    "10.10.12.0/24",
    "10.10.13.0/24"
  ],
  "single_nat_gateway": true,
  "vpc_flow_log_tags": {}
}
</code></pre>                                                                               |    no    |

### Outputs

| Name                          | Description |
| ----------------------------- | ----------- |
| [byo-vpc](./#output\_byo-vpc) | n/a         |
| [vpc](./#output\_vpc)         | n/a         |
