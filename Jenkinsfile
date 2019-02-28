#!/usr/bin/env groovy

pipeline {
    agent none
    stages {
        def image
        stage('Test and Build') {
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
            agent any
            steps {
                script {
                    image = docker.build("strangeamoeba/concurrency9")
                }
            }
        }
        stage('Push Image') {
            agent any
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
