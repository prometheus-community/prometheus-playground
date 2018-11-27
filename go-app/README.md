# Instrumented Go application

In this sandbox, a Prometheus intance scrapes metrics from a simple [Go](https://golang.org) web application.

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
`prometheus` | A Prometheus instance that's [configured](./prometheus/prometheus.yml) to scrape metrics from the Go application running on port 2112
`myapp` | A simple Go web application that pretty much only exports a simple metric (a [`myapp_processed_ops_total`](./myapp/main.go#L20-L24) counter that is [incremented](./myapp/main.go#L27-L34) every 2 seconds)

Once the sandbox is up and running, navigate to http://localhost:9090/graph in your browser and enter `myapp_processed_ops_total` into the expression bar to see the most up-to-date value of the counter.

You can also see this metrics by "scraping" the Go web server's `/metrics` endpoint:

```bash
curl localhost:2112/metrics | grep myapp_processed_ops_total
# HELP myapp_processed_ops_total The total number of processed events
# TYPE myapp_processed_ops_total counter
myapp_processed_ops_total 4
```