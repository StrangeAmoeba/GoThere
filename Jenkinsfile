pipeline {
    agent { docker { image 'golang:1.12-alpine' } }
    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh 'echo $GOPATH'
                sh 'echo $GOROOT'
                sh 'ls -a'
                sh 'go test ./...'
                sh 'go install -v ./...'
                sh 'echo "Passed"'
            }
        }
    }
}
