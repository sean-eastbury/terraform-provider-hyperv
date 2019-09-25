pipeline {
  agent {
    docker {
      image 'golang'
      args '-p 3030:3030'
    }

  }
  stages {
    stage('Format') {
      steps {
        sh 'make fmt'
      }
    }
  }
}