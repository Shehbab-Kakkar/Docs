 EKS cluster, and you want to:

✅ Create IAM users: John and Danny

✅ Add them to a group called developer-rnd

✅ Give them read-only access to EKS Kubernetes resources

✅ Set up the IAM policy, EKS aws-auth mapping, and Kubernetes RBAC

✅ Full Working Setup: EKS Read-Only Access for Developer Group
Prerequisites:

AWS CLI configured with admin credentials

kubectl connected to the EKS cluster

Replace placeholder values (like <region>, <cluster-name>, and <account-id>) with your actual values

🔹 Step 1: Create IAM Group & Users
📜 create-iam-developers.sh
#!/bin/bash

# Variables
GROUP_NAME="developer-rnd"
POLICY_NAME="EKSDeveloperReadOnlyPolicy"
ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)

# Create group
aws iam create-group --group-name $GROUP_NAME

# Create IAM policy
cat > readonly-policy.json <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "EKSDescribeAccess",
      "Effect": "Allow",
      "Action": [
        "eks:DescribeCluster"
      ],
      "Resource": "*"
    },
    {
      "Sid": "AllowSTSIdentity",
      "Effect": "Allow",
      "Action": [
        "sts:GetCallerIdentity"
      ],
      "Resource": "*"
    }
  ]
}
EOF

aws iam create-policy --policy-name $POLICY_NAME --policy-document file://readonly-policy.json

# Attach policy to group
aws iam attach-group-policy \
  --group-name $GROUP_NAME \
  --policy-arn arn:aws:iam::$ACCOUNT_ID:policy/$POLICY_NAME

# Create users and add to group
for user in John Danny; do
  aws iam create-user --user-name $user
  aws iam add-user-to-group --user-name $user --group-name $GROUP_NAME
  aws iam create-login-profile --user-name $user --password "TempPass123!" --password-reset-required
done

echo "Users John and Danny created and added to group $GROUP_NAME"

🔹 Step 2: Update aws-auth ConfigMap in EKS

Fetch your IAM user ARNs:

aws iam get-user --user-name John --query 'User.Arn' --output text
aws iam get-user --user-name Danny --query 'User.Arn' --output text


Example outputs:

arn:aws:iam::123456789012:user/John
arn:aws:iam::123456789012:user/Danny

📜 Edit the ConfigMap:

Run:

kubectl edit configmap aws-auth -n kube-system


Add under mapUsers::

mapUsers:
  - userarn: arn:aws:iam::123456789012:user/John
    username: john
    groups:
      - developer-rnd
  - userarn: arn:aws:iam::123456789012:user/Danny
    username: danny
    groups:
      - developer-rnd

🔹 Step 3: Create Kubernetes Read-Only Role and Bind It
📜 eks-readonly-role.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: eks-readonly-role
rules:
- apiGroups: [""]
  resources: ["pods", "services", "endpoints", "configmaps", "nodes", "namespaces"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["apps"]
  resources: ["deployments", "replicasets", "statefulsets", "daemonsets"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["batch"]
  resources: ["jobs", "cronjobs"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["networking.k8s.io"]
  resources: ["ingresses", "networkpolicies"]
  verbs: ["get", "list", "watch"]

📜 eks-readonly-rolebinding.yaml
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


Apply both:

kubectl apply -f eks-readonly-role.yaml
kubectl apply -f eks-readonly-rolebinding.yaml

✅ Step 4: Test Access

Once John and Danny:

Sign in via AWS CLI

Use this to configure kubeconfig:

aws eks update-kubeconfig --region <region> --name <cluster-name>


Then try:

kubectl get pods --all-namespaces


They should only be able to view resources — not create, delete, or modify them.

🚀 Optional Enhancements

Rotate initial passwords via aws iam create-access-key

Use AWS SSO or IAM Identity Center instead of IAM users

Lock down access to specific namespaces (I can show you how if needed)
