podTemplate(label: 'replicator', containers: [],
  volumes: [
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
  ]) {
    def image="zerospam/kubernetes-replicator"
    def tag = "1.3_${env.BUILD_NUMBER}"
    def repoArtifactory = "hub.docker.com"
    def imageArtifactory = "dev/build/ibp/kubernetes-replicator"
    def builtImage = null

    node {
        gitInfo = checkout scm
        
        stage('docker build') {
           builtImage = docker.build("${image}:${env.BUILD_ID}")
        }
    
        if (gitInfo.GIT_BRANCH.equals('master')) {
            // master branch release
            stage('Push docker image to Docker Hub') {
                builtImage.push('latest')
                builtImage.push('${tag}')
                builtImage.push('${gitInfo.GIT_COMMIT}')
            } // stage
        } // if master branch
    } // stages
} //pipeline