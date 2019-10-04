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
  }
    stage('Get-Terrafom') {
      steps {
        sh 'go get github.com/hashicorp/terraform'
      }
    }
  }
  environment {
    GO111MODULE = 'on'
    GOPATH = '/home/Go-Projects'
    GOBIN = '/home/Go-Projects/src/github.com/sean-eastbury/terraform-provider-hyperv/dist/'
  }
}
