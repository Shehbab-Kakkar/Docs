# 🔐 Okta Federated Read-Only Access to AWS EKS: Step-by-Step Guide

This guide explains how to grant Okta users **read-only Kubernetes access** to your AWS EKS cluster, using OIDC federation, IAM roles, Kubernetes RBAC, and client tools.

---

## 🚦 Architecture Overview

| Layer            | Tool/Config          | Purpose                                        |
|------------------|---------------------|------------------------------------------------|
| Identity         | Okta (OIDC)         | User authentication                            |
| AWS IAM          | Federated IAM Role  | Maps Okta users to AWS permissions             |
| EKS Cluster      | aws-auth ConfigMap   | Maps IAM role to k8s group                     |
| Kubernetes RBAC  | ClusterRole/Binding  | Assign read-only permissions                   |
| Auth Tool        | kubectl + OIDC      | Client login with Okta OIDC token              |

---

## 🛠️ Step 1: Add Okta as OIDC Identity Provider in AWS

1. Go to **IAM > Identity Providers > Add provider**
2. Provider Type: **OpenID Connect**
3. Provider URL: `https://<your-okta-domain>`
4. Audience: `sts.amazonaws.com`
5. Upload Okta's metadata document or URL
6. Save the ARN for use in the IAM role

---

## 🛠️ Step 2: Create an IAM Role for Okta Federated Users

**Trust Policy Example (`okta-eks-readonly-trust-policy.json`):**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::<account-id>:oidc-provider/<okta-domain>"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "<okta-domain>:sub": "okta-user-id" // or match by group claim
        }
      }
    }
  ]
}
```
- To use Okta groups:  
  `"StringEquals": { "<okta-domain>:groups": "developer-rnd" }`

**Create the role:**
```bash
aws iam create-role \
  --role-name OktaEKSReadOnlyRole \
  --assume-role-policy-document file://okta-eks-readonly-trust-policy.json
```

**Attach EKS read-only permissions:**
```bash
aws iam attach-role-policy \
  --role-name OktaEKSReadOnlyRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterAccessPolicy
# Or attach a custom restricted policy
```

---

## 🛠️ Step 3: Map IAM Role to Kubernetes Group in aws-auth ConfigMap

Edit your EKS `aws-auth` ConfigMap (via `kubectl edit -n kube-system configmap/aws-auth`):

```yaml
mapRoles:
  - rolearn: arn:aws:iam::<account-id>:role/OktaEKSReadOnlyRole
    username: okta-user
    groups:
      - developer-rnd
```
_Note: The group name here should match your RBAC binding in Kubernetes._

---

## 🛠️ Step 4: Create Kubernetes RBAC for Read-Only Access

**ClusterRole (`eks-readonly-role.yaml`):**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: eks-readonly-role
rules:
  - apiGroups: [""]
    resources: ["pods", "services", "endpoints", "namespaces"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["apps"]
    resources: ["deployments", "replicasets", "statefulsets", "daemonsets"]
    verbs: ["get", "list", "watch"]
  # Add more as needed
```

**RoleBinding (`eks-readonly-rolebinding.yaml`):**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: developer-rnd-readonly-binding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: eks-readonly-role
subjects:
  - kind: Group
    name: developer-rnd
    apiGroup: rbac.authorization.k8s.io
```

_Apply both:_
```bash
kubectl apply -f eks-readonly-role.yaml
kubectl apply -f eks-readonly-rolebinding.yaml
```

---

## 🛠️ Step 5: Configure kubectl for Okta OIDC Login

### Option 1: aws-iam-authenticator or eksctl

- Use Okta SSO to authenticate and assume the IAM role.
- Generate a kubeconfig that points to the EKS cluster with your assumed IAM identity.

### Option 2: kubelogin (Recommended for OIDC)

- Register an OIDC app in Okta for your users.
- Configure the kubeconfig user section like:
```yaml
users:
- name: okta-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: kubectl
      args:
        - oidc-login
        - get-token
        - --oidc-issuer-url=https://<your-okta-domain>
        - --oidc-client-id=<your-client-id>
        - --oidc-client-secret=<your-client-secret>
        - --oidc-extra-scope=groups
```
- Use [kubelogin](https://github.com/int128/kubelogin) to perform interactive login and obtain tokens.

---

## 📝 Summary Table

| Component         | You Configure                       |
|-------------------|-------------------------------------|
| Okta              | OIDC app with group/email claims     |
| AWS IAM           | Role with trust policy for Okta OIDC |
| EKS               | aws-auth ConfigMap role mapping      |
| Kubernetes RBAC   | ClusterRole & ClusterRoleBinding     |
| Client            | kubectl + kubelogin/oidc-auth        |

---

## 📚 References

- [AWS EKS OIDC Federation](https://docs.aws.amazon.com/eks/latest/userguide/authenticate-oidc-identity-provider.html)
- [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
- [kubelogin for OIDC](https://github.com/int128/kubelogin)
- [Okta OIDC Docs](https://developer.okta.com/docs/guides/implement-grant-type/authcode/main/)

---

**Tip:**  
For production, use Okta group claims to control access, and scope IAM/Kubernetes permissions as tightly as possible!
