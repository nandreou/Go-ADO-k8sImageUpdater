# Introduction 
This is a tool which runs in AzureDevOps Pipelines. Pushes an image to AzureContainerRegistry and then updates the k8s files in order your CI/CD (e.g. Argo) deploy the new app image on k8s cluster. 

# Getting Started
1. Go to azure DevOps Create Pipeline using the .yml in Folder /Pipelines/deployPipeline.yml (Or which ever Folder you put the pipeline file into).
2. Edit the Pipeline according to the paths that your files exist.
3. Generate PAT key in order for the pipeline has the credentials to push to repository.
3. Run Pipeline.
4. Check that everything was updated correctly.
5. Go to your CI/CD (e.g Argo) check that it tracked the change.
6. See your new Image to be deployed in the cluster. Happy Deploying !!!

# PS
In case You do not want to use azure pipelines. Just get the script code and use it how ever you want
