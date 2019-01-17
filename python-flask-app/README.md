# Instrumented Python Flask application

In this sandbox, a Prometheus instance scrapes metrics from a simple Python [Flask](http://flask.pocoo.org/) web application.

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
`myapp` | A simple Python Flask web application exports a [`LATENCY_HISTOGRAM`](./myapp/server.py#L6) and a [`REQUEST_COUNTER`](./myapp/server.py#L7)

Create some example web traffic to the app:

```shell
for n in {1..100}; do
    curl http://localhost:5000/test-endpoint
done
```

Once the sandbox is up and running, navigate to http://localhost:9090/graph in your browser and enter `myapp_request_count_total` into the expression bar to see the most up-to-date value of the counter.

You can also see this metrics by "scraping" the Flask web server's `/metrics` endpoint:

```bash
curl localhost:5000/metrics | grep myapp_request_count_total
# HELP myapp_request_count_total myapp HTTP request count
# TYPE myapp_request_count_total counter
myapp_request_count_total{endpoint="/metrics",method="GET",status="200"} 17.0
```

Here are some example metrics to explore in the Prometheus [expression browser](https://prometheus.io/docs/visualization/browser):

* [`myapp_request_count_total{endpoint="/test-endpoint"}`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=myapp_request_count_total%7Bendpoint%3D%22%2Ftest-endpoint%22%7D&g0)
* [`myapp_request_latency_seconds_bucket{endpoint="/test-endpoint",le="0.1"}`](http://localhost:9090/graph?g0.range_input=1h&g0.expr=myapp_request_latency_seconds_bucket%7Bendpoint%3D%22%2Ftest-endpoint%22%2Cle%3D%220.1%22%7D&g0)
