pipeline {
    docker { image 'golang:latest' }
  stages {
    stage('Preparation') {
      node {
        checkout scm
      }
    }
    stage('Build') {
      sh "docker build -t 'ruediger' ."
    }
    stage('Run') {
      sh "docker run -itd --name ruediger --publish 8090:80 ruediger"
    }
  }
}
