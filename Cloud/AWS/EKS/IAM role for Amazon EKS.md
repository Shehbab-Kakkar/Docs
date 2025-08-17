Creating an IAM role for Amazon EKS (Elastic Kubernetes Service) depends on what the role will be used for. There are a few common use cases:

🧰 Common EKS IAM Role Types

EKS Cluster Role — for the control plane (needed when creating a cluster)

Node Instance Role — for EC2 worker nodes (or Fargate)

IAM Roles for Service Accounts (IRSA) — fine-grained pod-level permissions

🎯 You mentioned: "IAM Role for EKS I want to create"

So let's clarify.

What do you need the role for?
Choose one:

✅ To create and manage an EKS Cluster

✅ To be used by EKS worker nodes (EC2 instances)

✅ To assign IAM permissions to Kubernetes pods (IRSA)

✅ Something else (e.g., CI/CD access to EKS)

If you're not sure, here's a basic example for each below.

📌 1. EKS Cluster Role (Used by EKS Control Plane)

This role is assumed by the EKS service to create and manage the cluster resources.

Trust Policy
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

Recommended IAM Policy to Attach

Use the AWS-managed policy:
AmazonEKSClusterPolicy

📌 2. EKS Node Instance Role (Used by EC2 Worker Nodes)

This role is assumed by EC2 instances running as EKS nodes.

Trust Policy
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

Recommended IAM Policies to Attach

AmazonEKSWorkerNodePolicy

AmazonEC2ContainerRegistryReadOnly

AmazonEKS_CNI_Policy (for networking)

📌 3. IAM Role for Service Account (IRSA)

Used to assign IAM permissions to Kubernetes pods.

Create an IAM Role with trust relationship to OIDC provider of your EKS cluster.

Example Trust Policy:

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


Attach fine-grained policy for your pod’s needs.
