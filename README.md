# Golang-microservices
Golang microservices project with docker ,  from code to the Azure cloud and kubernetes
This is a simple POC project.
In my home lab I have Proxmox server where I first created a kubernetes cluster with 2 nodes and SQL server 2022  inside a container on Ubuntu VM.
This POC with some little changes can work on any of the Clouds.


First I created 3 microservices using Golang ,
 ## DB-service 
 Just create the database UserDB, and one  table in it and write some data so that I have to work with something.

 ## web-service

 Loads the data from the database, run the web serrver on port 8081 internaly

 ## Front End server

Reads the data from the web-service and show it in the browser on port 8082 using some Golang HTML templates

## Configloader

Containes data about the database whih is used in db-service and web-service

## Dockerfile
in each of microservices folders there is the spacific dockerfile for that service

## docker compose
I tested this solution using docker compose first to see if everything is working as expected

## k8s folder

in this folder I have kubernetes deployments for all microservices.
db-service will be just a job not a service
web-service and front end are kubernetes services
because this is on my Proxmox server and not on the cloud I made the service type for both of them to be NodePort so I can access it from the browser
Web Service will be on port 30002 and Front End service on port 30003

## Azure Devops Pipeline

First I created a new project and connected it with the gihub repo where the code is.
Then I connected the local Proxmox kubernetes cluster with Azure DevOps pipeline.
I uploaded Kubeconfig file in the Pipeline library
On Azure I created Azure Container Registry where all the conatiners will be uploaded

Created two pipelines , build and deploy

## Build pipeline Stage
  Build the docker container for each service and upload it to ACR.
  For each service I have a Tag to better understand which container image was build and uploaded

## Deploy pipeline stage
  First deploys db-service job to populate the database.
  second is web-service and then third is frontend service
  






