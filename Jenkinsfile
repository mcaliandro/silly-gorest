podTemplate(yaml: '''
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: jenkins-admin
  containers:
    - name: go
      image: golang:1.20-alpine
      command: ["/bin/sh", "-c", "sleep 300"]
      volumeMounts:
        - name: docker-secret
          mountPath: /root/.docker
  volumes:
    - name: docker-secret
      secret:
        secretName: registry-secret
        items:
          - key: .dockerconfigjson
            path: config.json
''') {
    node(POD_LABEL) {
        stage('Build and publish') {
            container('go') {
                steps {
                    echo "Download project dependencies..."
                    sh 'go mod download'
                    echo "Perfrom unit test..."
                    sh 'go test'
                    echo "Download 'ko' builder"
                    sh 'go get github.com/google/ko@latest'
                    echo "Build a container image..."
                    sh 'ko build .'
                }
            }
        }
    }
}
