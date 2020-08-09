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
        stage('Test') {
            agent {
                label 'golang'
            }
            steps {
                sh 'go test ./...'
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
        stage('Deploy docker images') {
            agent {
                label 'docker'
            }
            steps {
                sh 'make docker-push'
            }
        }
        stage('Deploy to Kubernetes') {
            agent {
                label 'kubernetes'
            }
            steps {
                sh 'echo "TODO :)"'
            }
        }
    }
}
