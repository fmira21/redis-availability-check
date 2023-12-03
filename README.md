# Redis availability checker

This simple tool allows to test your Redis setup.

While creating it, I intended to learn more about Redis and its Go library - and check my own Redis cluster for availability corner cases.

Primarily, this tool is made to test clusters, but also can be used to load single-endpoint setups.

## Usage

All required configuration is done via `config.yml`.

Parameter | Default value | Description
--- | --- | ---
database | `0` | Redis database to operate with
password | `""` | Redis password
nonClusterAddress | `""` | Redis non-cluster endpoint. It works if `cluster.enabled` is false
verbose | `false` | Enable `DEBUG` log entries: iterations over the Redis DB
stopIfUnavailable | `false` | Stop execution if Redis becomes unavailable
cluster.enabled | `false` | Enable the cluster mode.
cluster.randomRouting | `false` | Connect randomly to endpoints from the `clusterAddresses` list. If `false` connection by latency is used
cluster.clusterAddresses | `[]` | List of Redis cluster endpoints

To set a config field to its default value, just remove it.

Specify your config as the first argument for the binary.

## Run and build

To run the source with `config.yml`, execute `make run`.
To build the binary, execute `make build`.