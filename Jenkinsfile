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
      environment {
        GO111MODULE = 'on'
        GOPATH = '/root/Go-Projects'
        PATH = '$GOPATH/bin'
      }
      steps {
        sh 'go get github.com/hashicorp/terraform'
        sh 'make build'
      }
    }
  }
}