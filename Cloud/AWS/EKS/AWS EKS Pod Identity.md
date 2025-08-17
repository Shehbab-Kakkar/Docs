# 🚀 AWS EKS Pod Identity

AWS EKS Pod Identity is the next-generation approach for securely assigning IAM roles to Kubernetes workloads running in Amazon EKS, without the operational complexity of the older IRSA (IAM Roles for Service Accounts) method. It is the new default for EKS IAM integration as of 2024.

---

## 🧭 What is EKS Pod Identity?

Traditionally, EKS workloads used [IRSA](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html), which required an OIDC provider and annotation-based role mapping for each service account.

**EKS Pod Identity** removes the need for OIDC and simplifies the process:
- No external OIDC provider required.
- No need for per-service-account IAM roles and annotations.
- Managed entirely by EKS and the AWS Pod Identity Agent.

---

## 🏗️ How Does It Work?

1. **Install the Pod Identity Agent:**  
   EKS-managed add-on (`eks-pod-identity-agent`) runs on your cluster nodes.

2. **Create an IAM Role for Pod Identity:**  
   A standard IAM role with a trust policy for the EKS Pod Identity.

3. **Create a Pod Identity Association:**  
   Maps a Kubernetes namespace/service account to the IAM role.

4. **Pod uses the Service Account:**  
   When a pod runs as that service account, EKS automatically injects the IAM role credentials.

---

## 🛠️ Step-by-Step: Enabling EKS Pod Identity

### 1️⃣ Enable the Pod Identity Agent Add-On

```bash
aws eks create-addon \
  --cluster-name my-eks-cluster \
  --addon-name eks-pod-identity-agent
```

### 2️⃣ Create an IAM Role for Pod Identity

**Trust Policy Example:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": { "Service": "pods.eks.amazonaws.com" },
      "Action": "sts:AssumeRole"
    }
  ]
}
```
- Attach the required IAM policies for your workload (e.g., `AmazonS3ReadOnlyAccess`).

### 3️⃣ Create a Pod Identity Association

```bash
aws eks create-pod-identity-association \
    --cluster-name my-eks-cluster \
    --namespace default \
    --service-account my-app-sa \
    --role-arn arn:aws:iam::<account-id>:role/EKSPodIdentityRole
```

### 4️⃣ Deploy Your Pod Using the Service Account

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-app-sa
  namespace: default
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      serviceAccountName: my-app-sa
      containers:
      - name: my-app
        image: <your-app-image>
        # your container spec here
```

---

## 🔎 How is This Different from IRSA?

| Feature                  | IRSA (Old)                  | EKS Pod Identity (New)           |
|--------------------------|-----------------------------|----------------------------------|
| OIDC Provider Needed     | Yes                         | No                               |
| Role Mapping             | Annotations (per SA)        | Pod Identity Association         |
| Complexity               | High (manual setup)         | Low (EKS native)                 |
| Default for New Clusters | No                          | Yes (2024+)                      |
| Security                 | Strong                      | Strong (simpler, less error-prone)|

---

## 📚 References

- [AWS EKS Pod Identity Official Docs](https://docs.aws.amazon.com/eks/latest/userguide/pod-identity.html)
- [EKS Pod Identity Announce Blog](https://aws.amazon.com/blogs/containers/introducing-amazon-eks-pod-identity/)
- [EKS Add-Ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html)

---

**Tip:**  
For most new EKS clusters, use Pod Identity instead of IRSA for easier, safer IAM-to-pod mapping!
