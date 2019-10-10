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
    stage('Get-Terraform') {
      steps {
        sh 'go get github.com/hashicorp/terraform'
      }
    }
    stage('Build for development') {
      when {
        branch 'development'
      }
      steps {
        sh 'make build'
      }
      post {
        always  {
          archiveArtifacts artifacts: 'dist/terraform-p*'
        }
      }
    }
    stage('Build for production') {
      when {
        branch 'production'
      }
      steps {
        sh 'make build'
      }
      post {
        always  {
          archiveArtifacts artifacts: 'dist/terrafor*'
        }
      }
    }
  }
  environment {
    GO111MODULE = 'on'
    GOPATH = '/home/Go-Projects'
    GOBIN = '/home/Go-Projects/src/github.com/sean-eastbury/terraform-provider-hyperv/dist/'
  }
}
