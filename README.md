# **Golang Microservices Project**
A proof-of-concept (POC) for deploying Golang microservices using Docker, Kubernetes, and Azure DevOps.

## **Project Overview**
This project showcases a complete microservices workflow from local development to deployment on Kubernetes, using Azure DevOps for CI/CD.

### **Home Lab Setup**
- Running on **Proxmox**, with a **Kubernetes cluster (2 nodes)**.
- **SQL Server 2022** is deployed inside a container on an **Ubuntu VM**.
- With minor adjustments, this setup can be deployed on any cloud platform.

---

## **Microservices Overview**
The project consists of **three microservices**, each developed in Golang:

### **1️⃣ DB Service**
- Initializes the `UserDB` database and a `Users` table.
- Inserts sample data for testing.
- Deployed as a **Kubernetes Job** (runs only once).

### **2️⃣ Web Service**
- Reads data from `UserDB`.
- Runs an internal web server on **port 8081**.
- Exposed as a **Kubernetes Service**.

### **3️⃣ Frontend Service**
- Fetches data from the Web Service.
- Uses **Golang HTML templates** to render data in a browser.
- Runs on **port 8082** and is exposed as a **Kubernetes Service**.

### **4️⃣ Config Loader**
- Stores database connection details.
- Used by both `db-service` and `web-service` to load configurations.

---

## **Containerization**

### **Dockerfiles**
- Each microservice has its own `Dockerfile` for containerization.

### **Docker Compose**
- Used for local testing before deploying to Kubernetes.

---

## **Kubernetes Deployment**

### **K8s Folder Structure**
- Contains deployment YAML files for all microservices.
- **`db-service`** is a **Kubernetes Job** (not a persistent service).
- **`web-service`** and **`frontend-service`** are **Kubernetes Services**.

### **Service Type**
- Since the cluster runs **on-prem (Proxmox)** and not in a cloud environment, services are exposed as **NodePort**.
- **Web Service** → Exposed on **port 30002**.
- **Frontend Service** → Exposed on **port 30003**.

---

## **Azure DevOps Pipeline**

### **Setup**
- Created an **Azure DevOps project** and connected it to **GitHub**.
- Connected the **Proxmox-based Kubernetes cluster** to **Azure DevOps**.
- Uploaded the **Kubeconfig** file to the **Azure DevOps Pipeline Library**.
- Created an **Azure Container Registry (ACR)** to store container images.

### **CI/CD Pipelines**

#### **1️⃣ Build Pipeline**
- Builds a Docker container for each service.
- Pushes images to **Azure Container Registry (ACR)**.
- Tags each image for version control.

#### **2️⃣ Deploy Pipeline**
- **Step 1**: Deploys `db-service` job to populate the database.
- **Step 2**: Deploys `web-service`.
- **Step 3**: Deploys `frontend-service`.

---

## **Next Steps**
- Implement automated database migrations.
- Add monitoring (Prometheus & Grafana).
- Improve security (RBAC & secrets management).
