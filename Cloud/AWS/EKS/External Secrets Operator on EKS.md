# 🛡️ External Secrets Operator on EKS: AWS Secrets Manager Integration for Django

This guide will walk you through an end-to-end setup to expose AWS Secrets Manager secrets to your Django pods running in EKS using [External Secrets Operator (ESO)](https://external-secrets.io/).

---

## 🔧 **Prerequisites**

- An **EKS** cluster up and running
- **kubectl** configured for your EKS cluster
- **Helm** installed
- AWS CLI configured with permissions to access EKS and Secrets Manager
- IAM permissions to create roles and policies

---

## 1️⃣ **Install External Secrets Operator (ESO)**

```sh
helm repo add external-secrets https://charts.external-secrets.io
helm repo update

kubectl create namespace external-secrets

helm upgrade --install external-secrets external-secrets/external-secrets \
  -n external-secrets
```

---

## 2️⃣ **Create an AWS IAM Policy for ESO**

Save as `eso-secretsmanager-policy.json`:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "secretsmanager:GetSecretValue",
        "secretsmanager:DescribeSecret",
        "secretsmanager:ListSecrets"
      ],
      "Resource": "*"
    }
  ]
}
```

**Create the policy:**

```sh
aws iam create-policy \
  --policy-name ESOSecretsManagerAccess \
  --policy-document file://eso-secretsmanager-policy.json
```

---

## 3️⃣ **Create an IAM Role for ServiceAccount (IRSA) and Attach Policy**

### a. **Create an OIDC provider for your cluster (if not already done):**

```sh
eksctl utils associate-iam-oidc-provider --region <region> --cluster <cluster-name> --approve
```

### b. **Create the IAM role and ServiceAccount:**

```sh
eksctl create iamserviceaccount \
  --name eso-sa \
  --namespace external-secrets \
  --cluster <cluster-name> \
  --attach-policy-arn arn:aws:iam::<account-id>:policy/ESOSecretsManagerAccess \
  --approve
```

- Replace `<cluster-name>` and `<account-id>` as appropriate.

---

## 4️⃣ **Create an AWS Secret in Secrets Manager**

Example: Create a secret for Django DB credentials.

```json
{
  "DJANGO_DB_NAME": "yourdbname",
  "DJANGO_DB_USER": "youruser",
  "DJANGO_DB_PASSWORD": "yourpassword",
  "DJANGO_DB_HOST": "yourdbhost",
  "DJANGO_DB_PORT": "5432"
}
```

```sh
aws secretsmanager create-secret \
  --name prod/django/db \
  --secret-string file://my-django-db-secret.json
```

---

## 5️⃣ **Configure External Secrets Operator in Kubernetes**

### a. **Create a SecretStore referencing the IRSA ServiceAccount**

```yaml
apiVersion: external-secrets.io/v1beta1
kind: SecretStore
metadata:
  name: aws-secretsmanager
  namespace: default
spec:
  provider:
    aws:
      service: SecretsManager
      region: <region>
      auth:
        jwt:
          serviceAccountRef:
            name: eso-sa
            namespace: external-secrets
```

Apply:

```sh
kubectl apply -f secretstore.yaml
```

---

### b. **Create an ExternalSecret Resource**

```yaml
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: django-db-secrets
  namespace: default
spec:
  refreshInterval: 1h
  secretStoreRef:
    name: aws-secretsmanager
    kind: SecretStore
  target:
    name: django-db
    creationPolicy: Owner
  dataFrom:
    - extract:
        key: prod/django/db
```

Apply:

```sh
kubectl apply -f externalsecret.yaml
```

---

### c. **Result**

- A Kubernetes secret named `django-db` will be created in the `default` namespace, containing your DB credentials.

---

## 6️⃣ **Use the Secret in Your Django Deployment**

Example snippet for your Django deployment:

```yaml
envFrom:
  - secretRef:
      name: django-db
```

Or for individual keys:

```yaml
env:
  - name: DJANGO_DB_NAME
    valueFrom:
      secretKeyRef:
        name: django-db
        key: DJANGO_DB_NAME
  # ... repeat for other keys
```

---

## 📝 **Summary of Steps**

1. **Install ESO via Helm**
2. **Create IAM policy and IRSA role/serviceaccount**
3. **Create AWS Secret in Secrets Manager**
4. **Define SecretStore and ExternalSecret in K8s**
5. **Use generated K8s secret in Django pod**

---

## 🛠️ **Troubleshooting**

- Check ESO logs:  
  `kubectl logs -n external-secrets deploy/external-secrets`
- Ensure IAM permissions are correct.
- Make sure ServiceAccount is referenced correctly in SecretStore.

---

## 📚 **References**

- [External Secrets Operator Docs](https://external-secrets.io/)
- [AWS EKS IRSA Docs](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
- [ESO AWS Provider Docs](https://external-secrets.io/providers/aws/)

---
