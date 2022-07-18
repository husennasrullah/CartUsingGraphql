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
                echo 'Running unit test...'
                sh 'go test ./unittest'
            }
        }
    }
}
