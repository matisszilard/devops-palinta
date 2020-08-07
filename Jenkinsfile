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
        stage('Build') {
            agent { 
                label 'docker'
            }
            steps {
                sh 'make docker-build'
            }
        }
    }
}
