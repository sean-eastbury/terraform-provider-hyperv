pipeline {
  agent {
    docker {
      image 'golang'
      args '-p 3030:3030'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh 'make fmt'
        sh 'make build'
      }
    }
  }
}