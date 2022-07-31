pipeline {
    agent none

    stages {
        stage('Prepare env') {
            agent {
                docker {
                    image 'golang:latest'
                    args '-v /var/run/docker.sock:/var/run/docker.sock'
                    args '-v /usr/bin/docker:/usr/bin/docker'
                }
            }
            steps {
                git branch: 'main', url: "https://github.com/ionnnnnnn/CRD.git"
                sh 'apt -y install make'
                sh 'go env -w GO111MODULE=auto'
                sh 'go env -w GOPROXY=https://goproxy.cn'
                sh 'make docker-build IMG=harbor.edu.cn/nju12/rolling-monitor:${BUILD_ID}'
                sh 'docker login harbor.edu.cn -u nju12 -p nju122022'
                sh 'make docker-push IMG=harbor.edu.cn/nju12/rolling-monitor:${BUILD_ID}'
            }
        }
    }
}
