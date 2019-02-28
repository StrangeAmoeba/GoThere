#!/usr/bin/env groovy

pipeline {
    agent {
        docker {
            image 'golang:1.12-alpine'
            customWorkspace "workspace/${BRANCH_NAME}/go/src/concurrency-9"
        }
    }
    environment {
        XDG_CACHE_HOME = "/tmp/.cache"
        GOPATH = "${WORKSPACE}/../.."
    }
    stages {
        stage('Init') {
            steps {
                sh 'sudo apk add docker'
                sh 'service docker start'
                checkout scm
            }
        }
        stage('Test and Build') {
            steps {
                checkout scm
                sh 'go version'
                sh 'go clean -cache'
                sh 'CGO_ENABLED=0 go test ./...'
                sh 'go build -v ./...'
                sh 'echo "Tests Passed"'
            }
        }
        stage('Build Image') {
            steps {
                script {
                    def image
                    image = docker.build("strangeamoeba/concurrency9")
                }
            }
        }
        stage('Push Image') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                        image.push("${env.BUILD_NUMBER}")
                        image.push("latest")
                    }
                }
            }
        }
    }
}
