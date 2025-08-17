# 🚀 Exposing Golang & Rust Apps in EKS with AWS Load Balancer Controller (ALB & NLB)

This guide walks you through exposing Golang and Rust applications running in private subnets of an EKS cluster using the [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/).  
- **Golang**: Exposed via ALB (Ingress)
- **Rust**: Exposed via NLB (Service of type LoadBalancer)

---

## 🟢 Prerequisites

- **EKS cluster** up and running
- **kubectl** and **helm** installed
- **eksctl** and **AWS CLI** configured
- Kubernetes pods (Golang and Rust) deployed in private subnets

---

## 🔧 1. Install AWS Load Balancer Controller

### a. Add Helm repo and update

```bash
helm repo add eks https://aws.github.io/eks-charts
helm repo update
```

### b. Create namespace

```bash
kubectl create namespace aws-load-balancer-controller
```

### c. Associate OIDC provider (once per cluster)

```bash
eksctl utils associate-iam-oidc-provider --region <your-region> --cluster <your-cluster-name> --approve
```

### d. Create IAM policy

```bash
curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/main/docs/install/iam_policy.json
aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy --policy-document file://iam-policy.json
```

### e. Create service account with IAM policy

```bash
eksctl create iamserviceaccount \
  --cluster <your-cluster-name> \
  --namespace aws-load-balancer-controller \
  --name aws-load-balancer-controller \
  --attach-policy-arn arn:aws:iam::<account-id>:policy/AWSLoadBalancerControllerIAMPolicy \
  --approve
```

### f. Install controller via Helm

```bash
helm install aws-load-balancer-controller eks/aws-load-balancer-controller \
  -n aws-load-balancer-controller \
  --set clusterName=<your-cluster-name> \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller \
  --set region=<your-region> \
  --set vpcId=<your-vpc-id>
```

---

## 🟦 2. Expose Golang App with ALB (Ingress)

Assumes your Golang app listens on port `8080`.

### a. Deployment

```yaml
# golang-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: golang-app
  labels:
    app: golang
spec:
  replicas: 2
  selector:
    matchLabels:
      app: golang
  template:
    metadata:
      labels:
        app: golang
    spec:
      containers:
      - name: golang
        image: <your-golang-image>
        ports:
        - containerPort: 8080
```

### b. Service

```yaml
# golang-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: golang-service
spec:
  selector:
    app: golang
  ports:
    - port: 80
      targetPort: 8080
  type: ClusterIP
```

### c. Ingress

```yaml
# golang-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: golang-ingress
  annotations:
    alb.ingress.kubernetes.io/scheme: internal
    alb.ingress.kubernetes.io/target-type: ip
    alb.ingress.kubernetes.io/subnets: <subnet-ids>  # Comma-separated
    alb.ingress.kubernetes.io/listen-ports: '[{"HTTP": 80}]'
    alb.ingress.kubernetes.io/load-balancer-attributes: deletion_protection.enabled=true
spec:
  ingressClassName: alb
  rules:
    - host: golang.internal.example.com
      http:
        paths:
        - path: /
          pathType: Prefix
          backend:
            service:
              name: golang-service
              port:
                number: 80
```

---

## 🟧 3. Expose Rust App with NLB (Service type=LoadBalancer)

Assumes your Rust app listens on port `9000`.

### a. Deployment

```yaml
# rust-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rust-app
  labels:
    app: rust
spec:
  replicas: 2
  selector:
    matchLabels:
      app: rust
  template:
    metadata:
      labels:
        app: rust
    spec:
      containers:
      - name: rust
        image: <your-rust-image>
        ports:
        - containerPort: 9000
```

### b. Service

```yaml
# rust-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: rust-service
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-type: "external"  # NLB
    service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: "ip"
    service.beta.kubernetes.io/aws-load-balancer-scheme: "internal"
    service.beta.kubernetes.io/aws-load-balancer-subnets: <subnet-ids>
spec:
  type: LoadBalancer
  selector:
    app: rust
  ports:
    - name: rust-port
      port: 9000
      targetPort: 9000
      protocol: TCP
```

---

## 🔒 Internal Access Notes

- Both ALB and NLB are `internal` (not internet-facing).
- Access via:
  - VPC-connected VPN
  - AWS PrivateLink
  - VPC Peering
  - Bastion host (for testing)

_Set `alb.ingress.kubernetes.io/scheme: internet-facing` or `service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing` for public access if required._

---

## 🏷️ DNS & Testing

- **ALB DNS**:  
  Check via  
  ```bash
  kubectl get ingress
  ```
  Look for the `ADDRESS` field.

- **NLB DNS**:  
  Check via  
  ```bash
  kubectl get svc rust-service
  ```
  Look for the `EXTERNAL-IP` field.

- **Route53**:  
  Point internal DNS names to the ALB/NLB endpoints as desired.

---

## 🧑‍💻 Example Workflow

```bash
kubectl apply -f golang-deployment.yaml
kubectl apply -f golang-service.yaml
kubectl apply -f golang-ingress.yaml

kubectl apply -f rust-deployment.yaml
kubectl apply -f rust-service.yaml
```

Check resources:
```bash
kubectl get ingress
kubectl get svc
```

---

## 📚 References

- [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/)
- [ALB Ingress Annotations](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.6/guide/ingress/annotations/)
- [NLB Service Annotations](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.6/guide/service/annotations/)
- [EKS Best Practices](https://aws.github.io/aws-eks-best-practices/networking/load-balancers/)

---

**Need help or a tailored config?**  
Open an issue with your architecture and network/DNS requirements!
