pipeline {
  agent any
  stages {
    stage('Preparation') {
      steps {
        node {
          label 'scm'
          checkout scm
        }
      }
    }
    stage('Build') {
      steps {
        sh "docker build -t 'ruediger' ."
      }
    }
    stage('Run') {
      steps {
        sh "docker run -itd --name ruediger --publish 8090:80 ruediger"
      }
    }
  }
}
