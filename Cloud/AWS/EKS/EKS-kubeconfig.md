✅ Corrected AWS CLI command to update kubeconfig for EKS:
aws eks update-kubeconfig --region <your-region> --name <your-cluster-name>


--region: The AWS region where your EKS cluster is located (e.g., us-west-2).

--name: The name of your EKS cluster.

This command updates your local ~/.kube/config file with credentials and cluster info from EKS.

Example:

aws eks update-kubeconfig --region us-west-2 --name my-eks-cluster

✅ Corrected kubectl command to check RBAC permissions:
kubectl auth can-i '*' --all-namespaces


This checks if your current identity has permission to do everything (*) in all namespaces.

Use --as or --as-group if you want to simulate another user/group.

Optional alternative:
To check all verbs on all resources:

kubectl auth can-i '*' '*' --all-namespaces
