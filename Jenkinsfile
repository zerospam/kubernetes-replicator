pipeline {
    agent {
        label 'jnlp-with-docker'
    }
    
    enviornment{
        def tag = "0.1_${env.BUILD_NUMBER}"
        def repoArtifactory = "docker.artifactory.a.intuit.com"
        def imageArtifactory = "dev/build/ibp/kubernetes-replicator"
    }

    stages {
        stage('Clone repository') {
            gitInfo = checkout scm
        }
        
        stage('docker build') {
            steps {
                echo "Building image ${image}:${tag}"
                sh "docker build --no-cache -t ${image}:${tag} ."
            }
        }
    
        if (gitInfo.GIT_BRANCH.equals('master')) {
            // master branch release
            stage('Push docker image to Artifactory') {
                docker.withRegistry("https://${repoArtifactory}", 'ibp-artifactory-creds') {
                    // Pushing multiple tags is cheap, as all the layers are reused.
                    sh "docker tag ${image}:${tag} ${repoArtifactory}/${imageArtifactory}:${tag}"
                    sh "docker tag ${image}:${tag} ${repoArtifactory}/${imageArtifactory}:latest"
                    sh "docker tag ${image}:${tag} ${repoArtifactory}/${imageArtifactory}:${gitInfo.GIT_COMMIT}"
                    sh "docker push ${repoArtifactory}/${imageArtifactory}:${tag}"
                    sh "docker push ${repoArtifactory}/${imageArtifactory}:latest"
                    sh "docker push ${repoArtifactory}/${imageArtifactory}:${gitInfo.GIT_COMMIT}"
                }
            } // stage
        } // if master branch
    } // stages
} //pipeline