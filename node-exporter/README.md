# Node Exporter sandbox

In this sandbox, [Prometheus](https://prometheus.io) scrapes Linux host metrics from a [Node Exporter](https://prometheus.io/docs/guides/node-exporter).

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up a `node_exporter` container and a `prometheus` container.

> To kill the sandbox, run `make kill` (alias for `docker-compose kill`).

Open up `http://localhost:9090/graph` to access the Prometheus [expression browser](https://prometheus.io/docs/visualization/browser). Here are some example metrics to explore:

* [`rate(node_cpu_seconds_total{mode="system"}[1m])`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=rate(node_cpu_seconds_total%7Bmode%3D%22system%22%7D%5B1m%5D)&g0.tab=1)
* [`node_filesystem_avail_bytes`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=node_filesystem_avail_bytes&g0.tab=1)
* [`rate(node_network_receive_bytes_total[1m])`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=rate(node_network_receive_bytes_total%5B1m%5D)&g0.tab=1)

## Assets

Folder | Assets
:------|:------
[`prometheus`](./prometheus) | A [`prometheus.yml`](./prometheus/prometheus.yml) configuration file for Prometheus