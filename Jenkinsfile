pipeline {
    agent any
    environment {
        PORT = '3051'
        DOCKER_TAG = 'Agrobal'
        EXTERNAL_PORT = '8000'
        LOGIN_BASE_PATH = "https://lila.com.ar"
    }
    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'docker build -t $DOCKER_TAG .'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Stopping previous version...'
                sh 'docker stop $DOCKER_TAG || echo No hay nada que detener'
                sh 'docker rm $DOCKER_TAG || echo No hay nada que eliminar'
                echo 'Deploying....'
                sh 'docker run -d -e LOGIN_BASE_PATH -e BASE_PATH -e PORT -p $EXTERNAL_PORT:8080 --name $DOCKER_TAG $DOCKER_TAG'
            }
        }
    }
}
