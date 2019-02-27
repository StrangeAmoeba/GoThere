pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'go test ./...'
                sh 'go install -v ./...'
            }
        }
    }
}
