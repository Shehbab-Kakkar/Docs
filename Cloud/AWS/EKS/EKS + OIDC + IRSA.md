# 🎯 EKS + OIDC + IRSA: Fine-Grained IAM for Kubernetes Pods

This guide walks you through creating an EKS cluster, enabling OIDC, and using **IAM Roles for Service Accounts (IRSA)** so your pods can securely access AWS resources with least privilege.

---

## 🛡️ Why Use OIDC and IRSA in EKS?

- **By default:** Pods inherit the IAM role of the node — too much access!
- **With IRSA:** Each pod can use a Kubernetes ServiceAccount mapped to a specific IAM role via OIDC, granting only the permissions it needs.

---

## 🛠️ Step-by-Step: EKS + OIDC + IRSA

### 1️⃣ Create the EKS Cluster (with OIDC)

The simplest way is with **eksctl** (installs OIDC provider automatically):

```bash
eksctl create cluster \
  --name my-eks-cluster \
  --region us-west-2 \
  --with-oidc \
  --nodes 2 \
  --node-type t3.medium \
  --managed
```

- `--with-oidc`: Enables the OIDC identity provider for the cluster (required for IRSA).
- ⏱️ Takes about 10–15 minutes.

---

### 2️⃣ Create an IAM Policy (e.g., S3 Access)

Define only the permissions your pods need. Example: Read-only access to a bucket.

**s3-access-policy.json:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["s3:GetObject"],
      "Resource": ["arn:aws:s3:::my-app-bucket/*"]
    }
  ]
}
```

Create the policy:

```bash
aws iam create-policy \
  --policy-name S3ReadOnlyPolicyForPods \
  --policy-document file://s3-access-policy.json
```

---

### 3️⃣ Create the IRSA Role and Associate with a Service Account

Create a service account (Kubernetes) linked to the IAM role (AWS):

```bash
eksctl create iamserviceaccount \
  --name s3-reader-sa \
  --namespace default \
  --cluster my-eks-cluster \
  --attach-policy-arn arn:aws:iam::<account-id>:policy/S3ReadOnlyPolicyForPods \
  --approve
```

This will:
- Create a Kubernetes ServiceAccount (`s3-reader-sa`)
- Create an IAM role with a trust relationship to the EKS OIDC provider
- Link the IAM role to the ServiceAccount

---

### 4️⃣ Deploy a Pod Using the Service Account

Sample deployment (`s3-reader-deployment.yaml`):

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: s3-reader
spec:
  replicas: 1
  selector:
    matchLabels:
      app: s3-reader
  template:
    metadata:
      labels:
        app: s3-reader
    spec:
      serviceAccountName: s3-reader-sa
      containers:
      - name: s3-reader
        image: amazonlinux
        command: ["/bin/sh", "-c"]
        args: ["yum -y install awscli; aws s3 ls s3://my-app-bucket"]
```

Apply it:

```bash
kubectl apply -f s3-reader-deployment.yaml
```

---

## 🔎 What’s Happening Behind the Scenes?

1. The pod starts with the `s3-reader-sa` ServiceAccount.
2. Kubernetes issues an OIDC token for the pod.
3. AWS STS validates the OIDC token and the IAM trust policy.
4. AWS returns temporary credentials scoped to the pod’s IAM role.
5. The pod uses only the permissions in the attached IAM policy.

---

## ✅ Summary Table

| Step | Action                               | Purpose                                    |
|------|--------------------------------------|--------------------------------------------|
| 1    | Create EKS with `--with-oidc`        | Sets up cluster and OIDC provider          |
| 2    | Create IAM policy                    | Define what permissions pods need          |
| 3    | Create IRSA with eksctl              | Associate IAM role to Kubernetes SA        |
| 4    | Deploy pod using that SA             | Pod assumes role via OIDC—fine-grained IAM |

---

## 📚 References

- [IAM Roles for Service Accounts (IRSA)](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
- [eksctl IRSA docs](https://eksctl.io/usage/iamserviceaccounts/)
- [AWS Blog: Fine-grained IAM roles for Kubernetes](https://aws.amazon.com/blogs/opensource/introducing-fine-grained-iam-roles-service-accounts/)

---

**Security Tip:**  
Never use the node role for pod permissions—use IRSA for least-privilege access!
