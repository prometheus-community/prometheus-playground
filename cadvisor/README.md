# cAdvisor sandbox

In this sandbox, [cAdvisor](https://github.com/google/cadvisor) gathers metrics for three running containers:

* Its own container
* A container running [Redis](https://redis.io)
* A container running [Prometheus](https://prometheus.io)

The Prometheus instance scrapes metrics from cAdvisor and makes them accessible via the Prometheus [expression browser](https://prometheus.io/docs/visualization/browser).

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up the three containers mentioned above.

> To kill the sandbox, run `make kill` (alias for `docker-compose kill`).

Open up `http://localhost:9090/graph` to access the Prometheus expression browser. Some example metrics to explore:

* [`rate(container_cpu_usage_seconds_total{name="redis"}[1m])`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=rate(container_cpu_usage_seconds_total%7Bname%3D%22redis%22%7D%5B1m%5D)&g0.tab=1)
* [`container_memory_usage_bytes{name="redis"}`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=container_memory_usage_bytes%7Bname%3D%22redis%22%7D&g0.tab=1)
* [`rate(container_network_transmit_bytes_total[1m])`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=rate(container_network_transmit_bytes_total%5B1m%5D)&g0.tab=1)
* [`rate(container_network_receive_bytes_total[1m])`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=rate(container_network_receive_bytes_total%5B1m%5D)&g0.tab=1)

## Assets

Folder | Assets
:------|:------
[`prometheus`](./prometheus) | A [`prometheus.yml`](./prometheus/prometheus.yml) configuration file for Prometheus