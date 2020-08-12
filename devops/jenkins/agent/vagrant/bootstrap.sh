#!/usr/bin/env bash

# Install docker
apt-get update
apt-get install apt-transport-https ca-certificates curl software-properties-common -y --no-install-recommends
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"
apt update
apt-cache policy docker-ce
apt-get install docker-ce -y --no-install-recommends
systemctl status docker
usermod -aG docker 'vagrant'
