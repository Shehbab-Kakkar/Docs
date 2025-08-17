# 🛡️ Creating IAM Roles for Amazon EKS (Elastic Kubernetes Service)

This guide explains the **three core types of IAM roles for EKS** and provides example trust policies and recommended AWS-managed policies for each use case.

---

## 🧰 Common EKS IAM Role Types

| Role Type                                      | Purpose                                         |
|------------------------------------------------|-------------------------------------------------|
| EKS Cluster Role                               | Control plane: create/manage EKS cluster        |
| Node Instance Role                             | EC2 worker nodes (or Fargate)                   |
| IAM Roles for Service Accounts (IRSA)          | Pod-level, fine-grained permissions             |

---

## 📌 1. EKS Cluster Role (Control Plane)

**Purpose:**  
Allows the EKS service to create and manage cluster resources.

**Trust Policy:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "eks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

**Recommended IAM Policy to Attach:**
- `AmazonEKSClusterPolicy` (AWS managed)

**How to create:**
```bash
aws iam create-role --role-name EKSClusterRole \
  --assume-role-policy-document file://cluster-trust.json

aws iam attach-role-policy --role-name EKSClusterRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterPolicy
```

---

## 📌 2. EKS Node Instance Role (EC2 Worker Nodes)

**Purpose:**  
Allows EC2 instances to join the EKS cluster and interact with AWS services.

**Trust Policy:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

**Recommended IAM Policies to Attach:**
- `AmazonEKSWorkerNodePolicy`
- `AmazonEC2ContainerRegistryReadOnly`
- `AmazonEKS_CNI_Policy`

**How to create:**
```bash
aws iam create-role --role-name EKSNodeInstanceRole \
  --assume-role-policy-document file://node-trust.json

aws iam attach-role-policy --role-name EKSNodeInstanceRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy

aws iam attach-role-policy --role-name EKSNodeInstanceRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly

aws iam attach-role-policy --role-name EKSNodeInstanceRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy
```

---

## 📌 3. IAM Role for Service Account (IRSA)

**Purpose:**  
Assigns IAM permissions to Kubernetes pods via service accounts (fine-grained, pod-level access).

**Trust Policy Example:**  
Replace `<ACCOUNT_ID>`, `<OIDC_PROVIDER_URL>`, `<namespace>`, `<serviceaccount-name>` as appropriate.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::<ACCOUNT_ID>:oidc-provider/<OIDC_PROVIDER_URL>"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "<OIDC_PROVIDER_URL>:sub": "system:serviceaccount:<namespace>:<serviceaccount-name>"
        }
      }
    }
  ]
}
```

**Attach:**  
- Only the fine-grained IAM policy your pods need (e.g., S3 access).

**How to create:**
```bash
aws iam create-role --role-name MyServiceAccountRole \
  --assume-role-policy-document file://irsa-trust.json

aws iam attach-role-policy --role-name MyServiceAccountRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
```

---

## 📝 Which One Do You Need?

- To create/manage EKS clusters: **EKS Cluster Role**
- For EC2 worker nodes: **Node Instance Role**
- For fine-grained pod permissions: **IRSA**
- For CI/CD or other use cases: tailor the trust and managed policies accordingly

---

## 📚 References

- [EKS IAM Documentation](https://docs.aws.amazon.com/eks/latest/userguide/security-iam.html)
- [IAM Roles for Service Accounts (IRSA)](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
- [EKS Best Practices](https://aws.github.io/aws-eks-best-practices/security/docs/iam/)

---

**Tip:**  
Always follow the principle of least privilege when attaching policies to your IAM roles.
