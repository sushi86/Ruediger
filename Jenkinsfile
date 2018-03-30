pipeline {
  agent any
  stages {
    stage('Clean') {
      steps {
        try {
          sh "docker rm -f \$(docker ps -aq --filter name=ruediger)"
        } catch(err) { 
          echo "error during cleanup containers, maybe no earlier container"
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
