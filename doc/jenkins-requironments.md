# Create CI/CD configuration with Jenkins

The aim of the playground is to create the CI/CD configuration for a project
using Jenkins.

Jenkins host: http://jenkins.apps.okd.codespring.ro

## Step 1: Create a custom Kubernetes Jenkins agent

Create a docker image containing the Jenkins agent (jenkins/jnlp-slave) and your
project specific dependencies.

Add your image to the Kubernetes configuration as a pod template.

Create a specific job to run on Kubernetes.

## Step 2: Create a local Jenkins agent

Create a Jenkins slave on your local machine. Add specific labels to it: use name
prefixes to avoid Jenkins agent collisions.

The agent should have the following possibility:

- Build project (in case of multi-stage docker builds it is not required)
- Build docker image
- Publish docker image (Docker Hub, Gitlab registry, etc.)
- Deploy micro services to Kubernetes cluster

Create jobs, pipelines (**Step 3**)

### Step 2.1: Optional: Create a Jenkins agent in vagrant

Create a Jenkins agent in vagrant. Different vagrant setup can be created for
each step.

## Step 3: Create a pipeline

The Jenkins pipeline should contain the following stages:

- Test, Build, Docker build, Deploy docker, Deploy kubernetes

> Note: use the Jenkins agents created previously.

## Step 4: Trigger Jenkins builds

Find out how can automatically trigger Jenkins builds.

## Step ?: Have a :beer:, have a kitkat! :tada:
