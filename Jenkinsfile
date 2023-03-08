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
                stage('shell') {
                    sh '''
                    echo "Download project dependencies..." && go mod download
                    echo "Perfrom unit test..." && go test
                    echo "Download 'ko' builder" && go get github.com/google/ko@latest
                    echo "Build a container image..." && ko build .
                    echo "Completed!"
                    '''
                }
            }
        }
    }
}
