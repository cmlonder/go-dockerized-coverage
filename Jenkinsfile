def workSpace
pipeline {
    agent any

    stages {

        stage('Test') {
            steps {
                workSpace = env.WORKSPACE
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

        stage("SonarQube analysis") {
            steps {
                withSonarQubeEnv('DEV-SonarQube_PLA_Project_9292') {
                    try {
                        sh "docker run --user=\$(id -u):\$(id -g) --rm -e SONAR_HOST_URL=http://sonar:9292 -v $workSpace:/usr/src sonarsource/sonar-scanner-cli"
                    } finally {
                        sh "docker image rm $imageName"
                    }
                }
            }
        }

        stage("Quality Gate Check"){
            steps {
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
