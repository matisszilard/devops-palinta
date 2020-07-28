# Palinta devops playground [![Build Status](https://travis-ci.com/matisszilard/devops-palinta.svg?branch=master)](https://travis-ci.com/github/matisszilard/devops-palinta)

Playground project for Kubernetes, Openshift, AWS, etc.

Note: the listed commands, params are heavily specific. To able to run in your
environment please change the kube configs, docker hub specific parameters to your
configuration.

## Build the project

To generate the binaries run the following command:

```sh
make xbuild
```

It is going to generate the macOS and Linux binaries.

## Build the project using docker

Build `demeter`:

```sh
make docker-demeter
```

Build data generator:

```sh
make docker-generator
```

## Upload it to Kubernetes

```sh
make kube-up
```

## Upload it to Openshift

```sh
make oc-up
```

Note: for more make commands please check the `Makefile`.
