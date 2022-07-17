pipeline {
    agent any

    stages {
        stage('compile') {
            steps {
                sh 'go build -o server main.go'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}
