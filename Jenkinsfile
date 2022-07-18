pipeline {
    agent any
    stages {
         stage('Compile') {
            steps {
               sh 'go build'
            }
         }
        stage('Test') {
            steps {
                sh 'go test ./unittest'
            }
        }
    }
}
