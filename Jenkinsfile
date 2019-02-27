pipeline {
    agent {
        docker {
            image "golang:1.12-alpine"
            customWorkspace "workspace/${BRANCH_NAME}/go/src/concurrency-9"
        }
    }
    environment {
        XDG_CACHE_HOME = "/tmp/.cache"
        GOPATH = "${WORKSPACE}/../.."
    }
    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh 'echo $GOPATH'
                sh 'echo $GOROOT'
                sh 'ls -a'
                sh 'pwd'
                sh 'CGO_ENABLED=0 go test ./...'
                sh 'go build -v ./...'
                sh 'echo "Passed"'
            }
        }
    }
}
