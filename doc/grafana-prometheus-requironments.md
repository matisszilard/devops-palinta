# Monitor micro services and visualize metrics with Prometheus and Grafana

> Goal: Create and deploy a micro service to Openshift / Kubernetes and show different information on a dashboard.

> Note: `Optional` fields are not required.

## Step 1: Create a microservice

Create a small project which contains at least 1 microservice. The service should have at least a Rest API endpoint.

Feel free to use any programming language, library or framework.

> Hint: choose a language, where is available a microservice framework/library.

> Please find out that in that language/framework combination what is the preferred microservice project layout.

- Create a docker image with the microservice and publish it into a registry (ex.: Docker Hub).
- Optional: Try out gRPC instead of HTTP
- Optional: Create a GitHub / GitLab project
- Optional: Build and deploy image in CI (Example: Gitlab CI, Travis CI)

## Step 2: Create metrics endpoint

Extend the created microservice to collect request related information (request count, request latency, etc.). Publish the collected information on endpoint `/metrics`.

> Optional: find out what other information can be collected.

Example prometheus libraries for this purpose:

- https://www.npmjs.com/package/express-prometheus-middleware
- https://medium.com/@dale.bingham_30375/net-core-web-api-metrics-with-prometheus-and-grafana-fe84a52d9843
- https://docs.spring.io/spring-metrics/docs/current/public/prometheus
- https://docs.spring.io/spring-boot/docs/current/reference/html/production-ready-features.html#production-ready-metrics-export-prometheus
- https://github.com/prometheus/client_java
- https://github.com/prometheus/client_golang

## Step 3: Create Kubernetes setup

Create deployment, services, config maps, persistent volumes for Grafana, for Prometheus, and for the microservice:

- Connect services together. The Grafana should automatically use the Prometheus data source and have a default dashboard.
- Setup Prometheus to find the microservice (Service Discovery)
  - Find out how can the Prometheus scrape metrics
- Create Ingress setup for services: Prometheus, Grafana, microservice

## Step 4: Dashboard setup

Find out what can be shown on the Grafana dashboard from the collected Prometheus data.

## Step 5: Have a beer!
