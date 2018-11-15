# Federation sandbox

In this sandbox, three Prometheus instances are run as a single Prometheus [federation](https://prometheus.io/docs/prometheus/latest/federation). All three instances scrape metrics from a single [Node Exporter](../node-exporter).

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up three Prometheus services (`prometheus1`, `prometheus2`, and `prometheus3`) and a `node_exporter` service.

Service | Description
:-------|:-----------
`prometheus1`, `prometheus2`, `prometheus3` | Standard Prometheus instances [configured](./prometheus/prometheus.yml#L10-L14) to run as a single federation
`node_exporter` | A Node Exporter instance that gathers and exposes Linux host metrics to be scraped by the Prometheus federation

> To kill the sandbox, run `make kill` (alias for `docker-compose kill`).

Open up `http://localhost:9090/graph` to access the Prometheus [expression browser](https://prometheus.io/docs/visualization/browser). The expression browser is actually available via all three Prometheus instances (ports 9090, 9091, and 9092).
