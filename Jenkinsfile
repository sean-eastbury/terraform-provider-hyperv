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
    stage('Test-Code') {
      steps {
        sh 'make test'
      }
    }
    stage('Build') {
      steps {
        sh 'make build'
      }
    }
  }
}