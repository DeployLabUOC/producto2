pipeline {
  agent any
  environment {
    APP     = 'goapp'
    ENTORNO = "${env.BRANCH_NAME == 'main' ? 'produccion' : 'pruebas'}"
    PUERTO  = "${env.BRANCH_NAME == 'main' ? '8081' : '8082'}"
  }
  stages {
    stage('Build') {
      steps { sh 'docker build -t $APP:$BRANCH_NAME-$BUILD_NUMBER .' }
    }
    stage('Deploy') {
      when { anyOf { branch 'main'; branch 'develop' } }
      steps {
        sh '''
          docker rm -f $APP-$ENTORNO || true
          docker run -d --name $APP-$ENTORNO -p $PUERTO:8080 \
            -e ENTORNO=$ENTORNO $APP:$BRANCH_NAME-$BUILD_NUMBER
        '''
      }
    }
  }
}