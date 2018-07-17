podTemplate {
    def image="zerospam/kubernetes-replicator"
    def tag = "1.3_${env.BUILD_NUMBER}"
    def repoArtifactory = "hub.docker.com"
    def imageArtifactory = "dev/build/ibp/kubernetes-replicator"
    def image = null

    node {
        gitInfo = checkout scm
        
        stage('docker build') {
           image = docker.build("${image}:${env.BUILD_ID}")
        }
    
        if (gitInfo.GIT_BRANCH.equals('master')) {
            // master branch release
            stage('Push docker image to Docker Hub') {
                image.push('latest')
                image.push('${tag}')
                image.push('${gitInfo.GIT_COMMIT}')
            } // stage
        } // if master branch
    } // stages
} //pipeline