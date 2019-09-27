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
      environment {
        GO111MODULE = 'on'
        GOPATH = '/root/Go-Projects'
      }
      steps {
        sh 'make build'
        sh 'go get github.com/hashicorp/terraform'
      }
    }
  }
}