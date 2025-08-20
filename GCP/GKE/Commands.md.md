# 🚀 Kubernetes & GKE Hands-on Guide: Commands Explained

This README provides an explanation for each command used in a typical Google Kubernetes Engine (GKE) workflow, including deployment, scaling, configuration, secrets, and cleanup.

---

## 1️⃣ **Set the GCP Project**

```sh
gcloud config set project my-kubernetes-project-304910
```
- Sets your default Google Cloud Project for all `gcloud` commands.

---

## 2️⃣ **Authenticate kubectl with GKE Cluster**

```sh
gcloud container clusters get-credentials my-cluster --zone us-central1-c --project my-kubernetes-project-304910
```
- Fetches Kubernetes credentials from GKE and configures `kubectl` to manage your cluster.

---

## 3️⃣ **Deploy an Application**

```sh
kubectl create deployment hello-world-rest-api --image=in28min/hello-world-rest-api:0.0.1.RELEASE
```
- Creates a deployment named `hello-world-rest-api` using the specified Docker image.

---

## 4️⃣ **Check Deployment Status**

```sh
kubectl get deployment
```
- Lists all deployments and their status.

---

## 5️⃣ **Expose the Deployment as a Service**

```sh
kubectl expose deployment hello-world-rest-api --type=LoadBalancer --port=8080
```
- Exposes your deployment with a public IP using a LoadBalancer on port 8080.

---

## 6️⃣ **Check Services**

```sh
kubectl get services
kubectl get services --watch
```
- Shows all services and their external/internal IPs.
- `--watch` keeps the output updating in real-time.

---

## 7️⃣ **Test the Application**

```sh
curl 35.184.204.214:8080/hello-world
```
- Sends an HTTP request to your app's public IP and port (replace with your actual IP).

---

## 8️⃣ **Scale the Deployment Manually**

```sh
kubectl scale deployment hello-world-rest-api --replicas=3
```
- Changes the number of running pods to 3.

---

## 9️⃣ **Resize GKE Node Pool**

```sh
gcloud container clusters resize my-cluster --node-pool default-pool --num-nodes=2 --zone=us-central1-c
```
- Changes the number of nodes (VMs) in the default node pool to 2.

---

## 🔟 **Enable Horizontal Pod Autoscaling**

```sh
kubectl autoscale deployment hello-world-rest-api --max=4 --cpu-percent=70
```
- Automatically scales pods between 1 and 4 based on average CPU usage (target 70%).

---

## 1️⃣1️⃣ **Check Autoscaler (HPA) Status**

```sh
kubectl get hpa
```
- Displays Horizontal Pod Autoscaler resources.

---

## 1️⃣2️⃣ **Create a ConfigMap**

```sh
kubectl create configmap hello-world-config --from-literal=RDS_DB_NAME=todos
```
- Creates a config map with a key-value pair for app configuration.

---

## 1️⃣3️⃣ **Inspect ConfigMaps**

```sh
kubectl get configmap
kubectl describe configmap hello-world-config
```
- Lists all config maps and shows detailed info about one.

---

## 1️⃣4️⃣ **Create a Secret**

```sh
kubectl create secret generic hello-world-secrets-1 --from-literal=RDS_PASSWORD=dummytodos
```
- Stores sensitive data (like passwords) in a Kubernetes secret.

---

## 1️⃣5️⃣ **Inspect Secrets**

```sh
kubectl get secret
kubectl describe secret hello-world-secrets-1
```
- Lists all secrets and shows details for one (note: values are base64 encoded).

---

## 1️⃣6️⃣ **Apply a Manifest File**

```sh
kubectl apply -f deployment.yaml
```
- Applies a YAML configuration to the cluster (can be used for deployments, services, etc.).

---

## 1️⃣7️⃣ **View Node Pools in the Cluster**

```sh
gcloud container node-pools list --zone=us-central1-c --cluster=my-cluster
```
- Lists all node pools in your GKE cluster.

---

## 1️⃣8️⃣ **Inspect Running Pods**

```sh
kubectl get pods -o wide
```
- Shows all pods with additional details like node assignment and IPs.

---

## 1️⃣9️⃣ **Update the Deployment Image**

```sh
kubectl set image deployment hello-world-rest-api hello-world-rest-api=in28min/hello-world-rest-api:0.0.2.RELEASE
```
- Updates the deployment to use a new Docker image version (rolling update).

---

## 2️⃣0️⃣ **Check Services, ReplicaSets, and Pods**

```sh
kubectl get services
kubectl get replicasets
kubectl get pods
```
- Displays the status of services, ReplicaSets (history of deployments), and pods.

---

## 2️⃣1️⃣ **Delete a Pod**

```sh
kubectl delete pod hello-world-rest-api-58dc9d7fcc-8pv7r
```
- Deletes a specific pod (will be recreated by the deployment controller).

---

## 2️⃣2️⃣ **Scale Down the Deployment**

```sh
kubectl scale deployment hello-world-rest-api --replicas=1
```
- Reduces the number of pods to 1.

---

## 2️⃣3️⃣ **Check ReplicaSets**

```sh
kubectl get replicasets
```
- Lists all ReplicaSets (each deployment update creates a new one).

---

## 2️⃣4️⃣ **List GCP Projects**

```sh
gcloud projects list
```
- Lists all Google Cloud projects you have access to.

---

## 2️⃣5️⃣ **Delete the Service and Deployment**

```sh
kubectl delete service hello-world-rest-api
kubectl delete deployment hello-world-rest-api
```
- Removes the service and deployment from your cluster.

---

## 2️⃣6️⃣ **Delete the GKE Cluster**

```sh
gcloud container clusters delete my-cluster --zone us-central1-c
```
- Deletes your entire Kubernetes cluster (irreversible).

---

# 📝 Summary

These commands cover the **full lifecycle** of a Kubernetes application on GKE:
- Setup & authentication
- Deploying & exposing an app
- Scaling & updating
- Managing configs and secrets
- Autoscaling & infrastructure
- Cleanup

> **Tip:** Always replace placeholder values (project name, cluster name, IPs) with your actual resources.

---
