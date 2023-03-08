pipeline {
	agent {
		kubernetes {
			yaml '''
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: jenkins-admin
  containers:
    - name: go
      image: golang:1.20-alpine
      command: ["/bin/sh", "-c", "sleep infinity"]
    - name: kaniko
      image: gcr.io/kaniko-project/executor:debug
      command: ["/bin/sh", "-c", "sleep infinity"]
      env:
        - name: CONTAINER_REGISTRY
          valueFrom:
            configMapKeyRef:
              name: registry-info
              key: CONTAINER_REGISTRY
        - name: IMAGE_NAME
          value: "silly-gorest"
        - name: IMAGE_TAG
          value: "latest"
      volumeMounts:
        - name: docker-secret
          mountPath: /kaniko/.docker
  volumes:
    - name: docker-secret
      secret:
        secretName: registry-secret
        items:
          - key: .dockerconfigjson
            path: config.json
'''
		}
	}
    stages {
		// 
        stage('Testing') {
			steps {
				git url: 'https://github.com/mcaliandro/silly-gorest.git', branch: 'main'
            	container('go') {
                    sh '''
                    echo "Download project dependencies..."
                    go mod download
                    echo "Perfrom unit test..."
                    go test
                    '''
                }
            }
        }
		// 
        stage('Publish') {
			steps {
            	container('kaniko') {
                    sh '''
                    /kaniko/executor \
                        --dockerfile=`pwd`/Dockerfile \
						--context=`pwd` \
                        --destination=${CONTAINER_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} \
						--verbosity debug
                    '''
                }
            }
        }
    }
}
