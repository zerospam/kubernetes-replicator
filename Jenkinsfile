def label = "replicator-${UUID.randomUUID().toString()}"
podTemplate(label: label, inheritFrom: 'docker') {
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
                    }
                } // stage
            } // if master branch

        } //container
    } //node
} //pipeline