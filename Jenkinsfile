#!/usr/bin/env groovy

pipeline {
    agent none
    stages {
        stage('Test') {
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
                sh 'go version'
                sh 'go clean -cache'
                sh 'CGO_ENABLED=0 go test ./...'
                sh 'go build -v ./...'
                sh 'echo "Tests Passed"'
            }
        }
        stage('Build and Push Image') {
            agent any
            when {
                branch 'master'
            }
            steps {
                script {
                    def image
                    image = docker.build("strangeamoeba/concurrency9")
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                        image.push("${BRANCH_NAME}-${env.BUILD_NUMBER}")
                        image.push("latest")
                    }
                }   
            }
            post {
                success {
                    sh "curl -X POST 'http://gothere.tk/restart'"
                }
            }
        }
    }
}
