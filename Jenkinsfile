pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                script {
                    def imageName = "localtest" + new Random().nextInt()
                    def containerName = "coverage" + new Random().nextInt()
                    try {
                       sh "docker build -t $imageName ."
                       sh "docker run -dt --rm --name $containerName $imageName"
                       sh "docker cp $containerName:/src/coverage.out ."
                     } finally {
                       sh "docker stop $containerName"
                     }
                }
             }
        }

        stage("SonarQube analysis") {
            steps {
                script {
                    try {
                        sh "docker run --user=\$(id -u):\$(id -g) --rm -e SONAR_HOST_URL=http://f591d5279d19.ngrok.io -v ${WORKSPACE}:/usr/src sonarsource/sonar-scanner-cli"
                    } finally {
                        sh "docker image rm $imageName"
                    }
                }
            }
        }

        stage("Quality Gate Check"){
            steps {
                script {
                    timeout(time: 4, unit: 'MINUTES') {
                        sleep(10)
                        def qg = waitForQualityGate()
                        if (qg.status != 'OK') {
                            error "Pipeline aborted due to quality gate failure: ${qg.status}"
                        }
                    }
                }
            }
        }
    }
}
