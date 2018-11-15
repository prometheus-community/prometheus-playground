# File-based service discovery sandbox

In this sandbox, a simple [web service](./myservice) running on the `myservice` host is scraped by a Prometheus instance. Prometheus discovers that host via [file-based service discovery](https://www.robustperception.io/using-json-file-service-discovery-with-prometheus).

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up two services:

Service | Description
:-------|:-----------
`prometheus` | A Prometheus instance configured to use the [`targets.json`](./prometheus/targets.json) file for file-based service discovery (see the [`dummy`](./prometheus/prometheus.yml#L2-L5) job configuration)
`myservice1` and `myservice2` | A basic web service that exposes a `/metrics` endpoint for Prometheus metrics and one custom metric: [`myservice_processed_ops_total`](./myservice/main.go#L20-L24). Two instances of this service are run.

> To kill the sandbox, run `make kill` (alias for `docker-compose kill`).

Once the sandbox is up and running, you can verify that the `myservice_processed_ops_total` metric is being scraped—and thus that the `myservice1` and `myservice2` services have been properly discovered by Prometheus—you can see the current value in the [expression browser](http://localhost:9090/graph?g0.range_input=1h&g0.expr=myservice_processed_ops_total&g0.tab=1).