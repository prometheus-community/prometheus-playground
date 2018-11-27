# Blackbox prober exporter sandbox

In this sandbox, the [Blackbox prober exporter](https://github.com/prometheus/blackbox_exporter) probes a [simple web service](./web/main.go), while Prometheus scrapes `probe_*` metrics from the Blackbox exporter.

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up three services:

Service | Description
:-------|:-----------
`blackbox` | A Prometheus Blackbox exporter running on port 9115, which runs probes against the `web` service using the [`http_2xx` module](./blackbox/blackbox.yml).
`prometheus` | A Prometheus instance configured to scrape the Blackbox exporter and collect a wide variety of metrics (prefaced by `probe_`).
`web` | A simple web service with two endpoints: `/hello` returns a `{"hello": "world"}` JSON object and `/health` returns an HTTP `200 OK` indicating that the service is running. The `web` service 

Things to check out:

* The Blackbox exporter dashboard at http://localhost:9115. You'll see a **Recent Probes** table listing all probes, including the following information for each probe: module, target, result, and a link to debugging logs.
* The results of the most recent probe of the `web` service at http://localhost:9115/probe?target=web:2112/health&module=http_2xx. If the value of `probe_success` metric is 1, then the service is running properly (otherwise the value will be 0).
* The `probe_success` metric on Prometheus at http://localhost:9090/graph?g0.range_input=1h&g0.expr=probe_success%7Binstance%3D%22web%3A2112%2Fhealth%22%7D&g0.tab=1.

At first, the `web` service will be running normally, which means that the Blackbox probes of that service's `/health` will succeed. This will change if you shut down the service:

```bash
docker-compose stop web
```

Now you can navigate back to the most recent [probe result](http://localhost:9115/probe?target=web:2112/health&module=http_2xx) and see that `probe_success` now has a value of 0. If you navigate back to the Prometheus [expression browser](http://localhost:9090/graph?g0.range_input=1h&g0.expr=probe_success%7Binstance%3D%22web%3A2112%2Fhealth%22%7D&g0.tab=1) you can see that `probe_success` for the `web` service now returns 0 as well (the scrape interval is 5 seconds, so you may need to wait a bit and hit refresh).

If you restart the `web` service, the probes will once again be successful:

```bash
docker-compose start web
```

## Assets

Folder | Assets
:------|:------
[`prometheus`](./prometheus) | A [`prometheus.yml`](./prometheus/prometheus.yml) configuration file for Prometheus
[`blackbox`](./blackbox) | A [`blackbox.yml`](./blackbox/blackbox.yml) configuration file for the Blackbox exporter
[`web`](./web) | Source files for the simple web server (written in Go)