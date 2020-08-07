pipeline {
    agent none

    stages {
        stage('Build') {
            agent { 
                label 'golang'
            }
            steps {
                sh 'make build'
            }
        }
        stage('Build docker images') {
            agent { 
                label 'docker'
            }
            steps {
                sh 'make docker-build'
            }
        }
    }
}
