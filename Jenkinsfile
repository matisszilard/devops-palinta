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
    }
}
