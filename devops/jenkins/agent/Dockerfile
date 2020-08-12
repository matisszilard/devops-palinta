FROM jenkins/jnlp-slave

RUN apt-get update && apt-get install -y --no-install-recommends \
    openjdk-11-jre-headless \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable" \
    && apt update && apt-cache policy docker-ce \
    && apt-get install docker-ce -y --no-install-recommends \
    && systemctl status docker \
    && usermod -aG docker 'jenkins'

CMD [ "" ]
