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
        sh 'export GOPATH="$HOME/Go-Projects"'
        sh ' export PATH=$PATH:$GOPATH/bin'
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
        archiveArtifacts '/'
      }
    }
  }
}