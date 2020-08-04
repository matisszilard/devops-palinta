# Collect log data using Elasticsearch, Logstash and Kibana

Elasticseacrh: http://elk.apps.okd.codespring.ro or http://okd-5mthh-worker-tb667.apps.okd.codespring.ro:30029
Kibana: http://elk-kibana.apps.okd.codespring.ro

## Step 1: Setup Kibana

Go to address: http://elk-kibana.apps.okd.codespring.ro .

Create a new space for user. Create unique index, dashboard.

## Step 2: Log directly into Elasticsearch

Create / update a microservice to send logs into Elasticsearch. The microservice should run on OC.

Elasticsearch host: http://elk.apps.okd.codespring.ro, port: 80 or http://okd-5mthh-worker-tb667.apps.okd.codespring.ro:30029 .

In order to able to view the logs follow the steps: select owned space, open hamburger menu, open `Discover` under the `Kibana` section.

> Testing: Trigger logs by calling the microservice API with curl.

## Step 3: Process and save logs with Logstash using HTTP API as input

Create / update a microservice to send logs into Logstash. The microservice can run locally.

Run Logstash `locally` using docker. Create configuration files for accessing the Elasticsearch and create pipelines to send logs to it.

> Testing: Trigger logs by calling the microservice API with curl. Log messages should be shown in the Logstash consol and Kibana's discover page.

## Step 4: Process and save logs with Logstash using a logfile as input

Create / update a microservice to send logs into a file. The microservice can run locally.

Setup Logstash to capture the given logfile as an input.

> Testing: Trigger logs by calling the microservice API with curl. Log messages should be shown in the Logstash consol and Kibana's discover page.

## Step 5: Optional: Process and save logs with Logstash using Filebeats as input

Create / update a microservice to send logs into a file. The microservice can run locally.

Setup Filebeats locally. Collect data from a given logfile. Send it to Logstash.
Configure Logstash to handle data from Filebeats.

## Step 6: Try out Kibana and Elasticsearch

Create different visualizations based on the logged data.

## Step 7: Have a beer, have a kitkat! :tada:
