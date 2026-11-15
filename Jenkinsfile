pipeline {
  agent any
  environment {
    ENTORNO  = "${env.BRANCH_NAME == 'main' ? 'produccion' : 'pruebas'}"
    NODEPORT = "${env.BRANCH_NAME == 'main' ? '30080' : '30081'}"
    IMAGEN   = "goapp:${env.BRANCH_NAME}-${env.BUILD_NUMBER}"
  }
  stages {
    stage('Build') {
      steps { sh 'docker build -t $IMAGEN .' }
    }
    stage('Cargar imagen en Minikube') {
      steps { sh 'minikube image load $IMAGEN' }
    }
    stage('Deploy en Kubernetes') {
      when { anyOf { branch 'main'; branch 'develop' } }
      steps {
        sh '''
          kubectl get namespace $ENTORNO || kubectl create namespace $ENTORNO
          sed -e "s|IMAGEN_VALOR|$IMAGEN|" -e "s|ENTORNO_VALOR|$ENTORNO|" k8s/deployment.yaml | kubectl apply -n $ENTORNO -f -
          sed -e "s|NODEPORT_VALOR|$NODEPORT|" k8s/service.yaml | kubectl apply -n $ENTORNO -f -
          kubectl rollout status deployment/goapp -n $ENTORNO --timeout=180s
          kubectl get pods -n $ENTORNO -o wide
        '''
      }
    }
  }
}
