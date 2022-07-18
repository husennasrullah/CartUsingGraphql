pipeline {
    agent any
    stages {
         stage('Compile') {
            steps {
               sh 'go build'
            }
         }
        stage('Unit Tests') {
            steps {
                echo 'Running unit test...'
                sh 'go test ./unittest'
            }
        }
    }
}
