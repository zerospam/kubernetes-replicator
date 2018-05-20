podTemplate
{
    def image="ibp/kubernetes-replicator"
    def tag = "0.1_${env.BUILD_NUMBER}"
    def repoArtifactory = "docker.artifactory.a.intuit.com"
    def imageArtifactory = "dev/build/ibp/kubernetes-replicator"

    node {
        gitInfo = checkout scm
        
        docker.withRegistry("https://${repoArtifactory}", 'ibp-artifactory-creds') {
            echo "Building image ${image}:${tag}"
            def customImage = docker.build("${image}:${tag}")

            /* Push the container to the custom Registry */
            if (gitInfo.GIT_BRANCH.equals('master')) {
                customImage.push("${repoArtifactory}/${imageArtifactory}:${tag}")
            }
        }
    
    } // node
} //pipeline