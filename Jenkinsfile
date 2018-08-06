def label = "replicator-${UUID.randomUUID().toString()}"
podTemplate(label: label, containers: [
    containerTemplate(name: 'docker', image: 'docker', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'jnlp', image: 'jenkins/jnlp-slave:3.19-1-alpine', args: '${computer.jnlpmac} ${computer.name}')
 ],
  volumes: [
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
  ]) {
    def image="zerospam/kubernetes-replicator"
    def tag = "1.3"
    def builtImage = null

    node (label) {
        gitInfo = checkout scm
         container('docker') {
            stage('docker build') {
               builtImage = docker.build("${image}:${env.BUILD_ID}")
            }

            if (gitInfo.GIT_BRANCH.equals('master')) {
                // master branch release
                stage('Push docker image to Docker Hub') {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub') {
                            builtImage.push('latest')
                            builtImage.push("${tag}")
                            builtImage.push("${gitInfo.GIT_COMMIT}")
                    }
                } // stage
            } // if master branch

        } //container
    } //node
} //pipeline