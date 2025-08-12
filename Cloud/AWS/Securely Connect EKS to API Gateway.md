# Securely Connect EKS to API Gateway via Private NLB and VPC Link

This guide walks you through securely exposing an EKS service through AWS API Gateway using a **private Network Load Balancer (NLB) and VPC Link**—all with explicit AWS CLI and Kubernetes commands, and optionally with ACK for Kubernetes-native API Gateway management.

---

## 📋 Architecture Overview

- **EKS**: Runs your Kubernetes workloads.
- **Network Load Balancer (NLB)**: Exposes your EKS service internally.
- **API Gateway**: Provides a managed API endpoint, integrating with the NLB via VPC Link for secure, private connectivity.
- **VPC Link**: Connects API Gateway privately to your VPC resources (the NLB).

---

## 🛠️ Step-by-Step Instructions

### 1. **Create EKS Cluster & Deploy Sample Service**

```bash
export AWS_REGION=ap-south-1
export CLUSTER_NAME=eks-demo

eksctl create cluster \
  --name $CLUSTER_NAME \
  --region $AWS_REGION \
  --managed
```

**Deploy a sample app (echoserver) and expose it via NLB:**
```bash
kubectl apply -f https://raw.githubusercontent.com/aws-samples/amazon-apigateway-ingress-controller-blog/Mainline/apigw-ingress-controller-blog/echoserver.yml
```
> This creates a Kubernetes Service of type `LoadBalancer` (provisions an NLB). Get the NLB DNS/ARN for API Gateway integration.

---

### 2. **Install AWS Load Balancer Controller (for NLB support)**

```bash
eksctl utils associate-iam-oidc-provider \
  --region $AWS_REGION \
  --cluster $CLUSTER_NAME \
  --approve

curl -S https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.2.0/docs/install/iam_policy.json -o iam-policy.json

aws iam create-policy \
  --policy-name AWSLoadBalancerControllerPolicy \
  --policy-document file://iam-policy.json

eksctl create iamserviceaccount \
  --cluster=$CLUSTER_NAME \
  --region=$AWS_REGION \
  --namespace=kube-system \
  --name=aws-load-balancer-controller \
  --attach-policy-arn=arn:aws:iam::$(aws sts get-caller-identity --query Account --output text):policy/AWSLoadBalancerControllerPolicy \
  --approve

helm repo add eks https://aws.github.io/eks-charts && helm repo update

VPC_ID=$(aws eks describe-cluster --name $CLUSTER_NAME --region $AWS_REGION --query "cluster.resourcesVpcConfig.vpcId" --output text)

helm install aws-load-balancer-controller \
  eks/aws-load-balancer-controller \
  --namespace kube-system \
  --set clusterName=$CLUSTER_NAME \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller \
  --set vpcId=$VPC_ID \
  --set region=$AWS_REGION
```
> This enables Kubernetes Services of type LoadBalancer to provision and manage AWS NLBs.

---

### 3. **Create a VPC Link in API Gateway**

#### **A. Using ACK (Kubernetes-native):**

`vpclink.yaml`
```yaml
apiVersion: apigatewayv2.services.k8s.aws/v1alpha1
kind: VPCLink
metadata:
  name: nlb-internal
spec:
  name: nlb-internal
  securityGroupIDs: ["<your-sg-id>"]
  subnetIDs:
    - "<subnet-id-1>"
    - "<subnet-id-2>"
```
```bash
kubectl apply -f vpclink.yaml
```

#### **B. Using AWS CLI Directly:**

```bash
aws apigatewayv2 create-vpc-link \
  --name my-vpclink \
  --subnet-ids subnet-abcde123 subnet-fghij456 \
  --security-group-ids sg-0123456789abcdef0
```
> This creates a VPC Link resource for API Gateway to access your NLB privately.

---

### 4. **Set Up API Gateway Integration to NLB**

#### **A. Using ACK (Kubernetes-native):**

`apigw-api.yaml`
```yaml
apiVersion: apigatewayv2.services.k8s.aws/v1alpha1
kind: API
metadata:
  name: eks-api
spec:
  body: '{
    "openapi":"3.0.1",
    "paths": {
      "/": {
        "x-amazon-apigateway-any-method": {
          "isDefaultRoute": true,
          "x-amazon-apigateway-integration": {
            "type":"HTTP_PROXY",
            "connectionType":"VPC_LINK",
            "connectionId":"$(kubectl get vpclinks.apigatewayv2.services.k8s.aws nlb-internal -o jsonpath="{.status.vpcLinkID}")",
            "httpMethod":"GET",
            "uri":"<YOUR_NLB_LISTENER_ARN>"
          }
        }
      }
    }
  }'
```
```bash
kubectl apply -f apigw-api.yaml
```

**Deploy a stage:**
```bash
cat <<EOF | kubectl apply -f -
apiVersion: apigatewayv2.services.k8s.aws/v1alpha1
kind: Stage
metadata:
  name: prod-stage
spec:
  apiID: $(kubectl get apis.apigatewayv2.services.k8s.aws eks-api -o=jsonpath='{.status.apiID}')
  stageName: prod
  autoDeploy: true
EOF
```

#### **B. Using AWS CLI:**

```bash
aws apigatewayv2 create-api \
  --name eks-proxy \
  --protocol-type HTTP \
  --target-type VPC_LINK \
  --connection-type VPC_LINK \
  --connection-id <VPC_LINK_ID> \
  --route-key "GET /" \
  --target "<NLB_LISTENER_ARN>"

aws apigatewayv2 create-deployment --api-id <API_ID> --stage-name prod
```
> This sets up API Gateway methods backed by your EKS service via NLB and VPC Link.

---

### 5. **Test the Integration**

```bash
curl https://<api-id>.execute-api.${AWS_REGION}.amazonaws.com/prod/
```
> You should see the response from your EKS service (echo server).

---

## ✅ Summary Checklist

| Step | Action                                         |
|------|------------------------------------------------|
| 1    | Create EKS cluster and deploy service with NLB |
| 2    | Install AWS Load Balancer Controller           |
| 3    | Create VPC Link targeting NLB                  |
| 4    | Setup API Gateway integration (ACK or CLI)     |
| 5    | Deploy and test the API                        |

---

## 🔒 Security Notes

- The NLB and API Gateway communicate privately via VPC Link—no public exposure of the EKS service.
- Ensure proper security group and subnet configuration for least privilege.

---

## 🏷️ References

- [AWS EKS Documentation](https://docs.aws.amazon.com/eks/)
- [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/)
- [AWS API Gateway VPC Link](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-vpc-links.html)
- [AWS ACK Controller](https://aws-controllers-k8s.github.io/community/)

---
