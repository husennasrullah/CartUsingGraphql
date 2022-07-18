pipeline {
    agent any
    stages {
        stage('SCM') {
            checkout scm
        }
         stage('Compile') {
            steps {
               sh 'go build'
            }
         }
        stage('Test') {
            steps {
                 echo 'Running unit test...'
                sh 'go test ./unittest'
            }
        }
    }
}
