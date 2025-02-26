# Golang-microservices
Golang microservices project with docker ,  from code to the Azure cloud and kubernetes
This is a simple POC project.
In my home lab I have Proxmox server where I first created a kubernetes cluster with 2 nodes and SQL server inside a container on Ubuntu VM.
This with some little changes can work on any of the Clouds.


First I created 3 microservices using Golang ,
 ## DB-service 
 Just create the database UserDB, and one  table in it and write some data so that I have to work with something.

 ## web-service

 Loads the data from the database, run the web serrver on port 8081 internaly

 ## Front End server

Reads the data from the web-service and show it in the browser using some Golang templates

## Configloader

Containes data about the database whih is used in db-service and web-service

## Dockerfile
in each of microservices folders there is the spacific dockerfile for that service

## docker compose
I tested this solution using docker compose first to see if everything is working as expected

### Azure Devops Pipeline



