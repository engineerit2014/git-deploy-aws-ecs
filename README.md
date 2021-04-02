# Deploy: Dockerized Microservice to AWS ECS (EC2 instance)
This project shows how to deploy a dockerized Golang microservice on AWS ECS using Github Actions. Emphasizing that the same process works for any microservice deployed in another language, because the starting point is to have the microservice dockerized.

Before we start, it is important to list the steps that Github, or we ourselves must execute if we do not use Github Actions to make our microservice available to an EC2 instance.

Workflow:

- Build and push a new container image to Amazon ECR.
- Deploy a new task definition to Amazon ECS.
- Check the service using the public IP of the configured EC2 or ELB.

## Table of Contents

- [Installing Go](#installing-go)
    - [Local Installation](#local-installation)
- [Editors](#editors)
- [Installing the repository](#installing-the-repository)
- [Configuration](#configuration)
    - [Dependencies](#dependencies)
    - [Required files in your repo](#required-files-in-your-repo)
    - [Set-up steps](#set-up-steps)
- [Running](#running)

## Installing Go

### Local Installation

https://www.ardanlabs.com/blog/2016/05/installing-go-and-your-workspace.html

## Editors

**Visual Studio Code**  
https://code.visualstudio.com/Updates  
https://github.com/microsoft/vscode-go

**VIM**  
http://www.vim.org/download.php  
http://farazdagi.com/blog/2015/vim-as-golang-ide/

**Goland**  
https://www.jetbrains.com/go/

## Installing the repository

From a command prompt, issue the following commands:

```sh
mkdir -p $(go env GOPATH)/src/github.com/laironacosta && cd $_
git clone https://github.com/laironacosta/git-deploy-aws-ecs.git
```

*NOTE:* This assumes you have Git installed.  If you don’t, you can find the installation instructions here: https://git-scm.com/

## Configuration

### Dependencies 

The only dependencies to start the configuration of your project are:

- `Dockerfile`: You must have implemented the Dockerfile of your microservice.
- `AWS Console Access`: You must have access to the AWS console.
- `AWS Credentials`: You must have aws credentials configured on your machine to be able to execute commands to AWS.
- `Github`: Your repository must be versioned on Github and you must have access.

### Required files in your repo

You must create two files in order to deploy the microservice on AWS ECS using Github actions:

1. `.github/workflows/aws.yml`: This file contains the entire workflow that will build and push a new container image to Amazon ECR, and then deploy a new task definition to Amazon ECS, when a condition is met in Github such as doing Push to branch develop. (copy the example code from the repo)
2. `task-definition.json`: This file contains the details of your container definition on AWS.

### Set-up steps

#### 1. Create an ECR repository to store your images
- For example: `aws ecr create-repository --repository-name my-ecr-repo --region us-east-2`.
- Replace the value of `ECR_REPOSITORY` in the workflow file (`aws.yml)` with your repository's name.
- Replace the value of `aws-region` in the workflow file (`aws.yml)` with your repository's region.

#### 2. Create an ECS task definition, an ECS cluster, and an ECS service
- For example, follow the Getting Started guide on the ECS console: https://us-east-2.console.aws.amazon.com/ecs/home?region=us-east-2#/firstRun
- Replace the values for `service` and `cluster` in the workflow file (`aws.yml)` with your service and cluster names.

#### 3. Store your ECS task definition as a JSON file in your repository
- The format should follow the output of `aws ecs register-task-definition --generate-cli-skeleton`.
- Replace the value of `task-definition` in the workflow file (`aws.yml)` with your JSON file's name.
- Replace the value of `container-name` in the workflow file (`aws.yml)` with the name of the container in the `containerDefinitions` section of the task definition.

#### 4. Store an IAM user access key in GitHub Actions secrets named `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.
- See the documentation for each action used below for the recommended IAM policies for this IAM user, and best practices on handling the access key credentials.

## Running

After configuring all the above, it is time to push the changes to your Github repository and after the automatic execution of the workflow in Github Actions, you could validate the availability of your app in your EC2 instance.