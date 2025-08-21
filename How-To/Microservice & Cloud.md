
***

## 7) Upgrading a Kubernetes Cluster

Upgrades should be safe, staged, and reversible.

### General Strategy

- Read release notes and deprecations.
- Upgrade control plane → node pools → system add-ons → workloads.
- Use surge/rolling strategies; ensure PodDisruptionBudgets (PDBs) are set.
- Back up etcd and app data before starting.


### kubeadm (self-managed) example

1. Plan:
```bash
kubeadm upgrade plan
```

2. Upgrade control plane:
```bash
# On control-plane node
sudo apt-get update && sudo apt-get install -y kubeadm=1.30.3-00
sudo kubeadm upgrade apply v1.30.3
sudo apt-get install -y kubelet=1.30.3-00 kubectl=1.30.3-00
sudo systemctl restart kubelet
```

3. Upgrade workers:
```bash
kubectl drain <node> --ignore-daemonsets --delete-emptydir-data
sudo apt-get update && sudo apt-get install -y kubeadm=1.30.3-00
sudo kubeadm upgrade node
sudo apt-get install -y kubelet=1.30.3-00
sudo systemctl restart kubelet
kubectl uncordon <node>
```


### Managed clusters

- EKS: Update control plane in console/CLI, then update nodegroup AMIs and perform rolling replace:

```bash
aws eks update-cluster-version --name my-eks --kubernetes-version 1.30
eksctl upgrade nodegroup --cluster my-eks --name ng-1
```

- GKE/AKS: Use console/CLI “Surge upgrade”/“Node image auto-upgrade”.

***

## 2) Handling Rollbacks in Kubernetes

### Roll back a Deployment

```bash
kubectl rollout history deployment/my-app
kubectl rollout undo deployment/my-app                 # to previous
kubectl rollout undo deployment/my-app --to-revision=3 # to specific
kubectl rollout status deployment/my-app
```

Ensure deployment has revision history:

```yaml
spec:
  revisionHistoryLimit: 10
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 0
```


### Helm rollback

```bash
helm history my-release
helm rollback my-release 5
helm status my-release
```


### Canary/Blue‑Green tip

- Use separate Service selectors or Ingress routes to switch traffic.
- Keep last known good image tag available.

***

## 10) Terraform State Drift: Detect and Fix

Drift: When real infrastructure changes outside Terraform or by other processes, deviating from state.

### Detect

```bash
terraform plan
terraform plan -detailed-exitcode   # exit 2 indicates changes (drift)
terraform state list
terraform providers schema -json     # to inspect resources if needed
```


### Fix

- If the real infra is correct, update code to match, then `terraform apply`.
- If code is correct, re-apply:

```bash
terraform apply
```

- If a resource was changed/replaced manually, and you want Terraform to adopt/point to it:

```bash
terraform import aws_instance.web i-0123456789abcdef0
```

- If an object should be forgotten:

```bash
terraform state rm aws_s3_bucket.logs
```

- If IDs changed:

```bash
terraform state replace-provider 'registry.terraform.io/-/aws' 'registry.terraform.io/hashicorp/aws'
```

Use drift detection in CI by running `terraform plan -detailed-exitcode` nightly.

***

## 1) Isolating a Network within an AWS VPC

Options: Subnets + route tables + NACLs + SGs; dedicated VPC; VPC Lattice; Dedicated TGW attachments.

### Isolated private subnet pattern

- No route to IGW (Internet Gateway).
- Optional NAT-only egress via NAT Gateway in a public subnet.
- Restrict inbound with Security Groups; restrict subnet-wide with NACLs.

Example:

1. Create subnets:

- Public: 10.0.1.0/24 with route to IGW.
- Private-isolated: 10.0.2.0/24 with no IGW/NAT route (true isolation).

2. Route table change:
```bash
# Associate 10.0.2.0/24 with route table that only has local route 10.0.0.0/16 → local
```

3. Lock down NACLs:

- Inbound: allow from known subnets/ports only.
- Outbound: deny 0.0.0.0/0.

4. SG only allows traffic from specific SGs:
```bash
aws ec2 authorize-security-group-ingress \
  --group-id sg-abc \
  --source-group sg-def \
  --protocol tcp --port 5432
```

For stricter isolation, place sensitive workloads in a separate VPC and connect only via VPC peering/TGW with restrictive routes.

***

## 9) Connecting Two AWS VPCs

### VPC Peering (simple, flat)

- No transitive routing; same/overlapping CIDRs not allowed.

```bash
aws ec2 create-vpc-peering-connection --vpc-id vpc-a --peer-vpc-id vpc-b
aws ec2 accept-vpc-peering-connection --vpc-peering-connection-id pcx-123
# Add routes in both VPC route tables pointing to pcx-123
```


### AWS Transit Gateway (hub-and-spoke)

- Scales to many VPCs, supports transitive routing.

```bash
aws ec2 create-transit-gateway --description core-tgw
aws ec2 create-transit-gateway-vpc-attachment --transit-gateway-id tgw-123 --vpc-id vpc-a --subnet-ids subnet-a1 subnet-a2
aws ec2 create-transit-gateway-vpc-attachment --transit-gateway-id tgw-123 --vpc-id vpc-b --subnet-ids subnet-b1 subnet-b2
# Add VPC routes to TGW, configure TGW route tables/associations
```


### PrivateLink (service exposure)

- One-way, endpoint service to consumers; not general routing.

Choose: Peering for 1:1; TGW for multi-VPC/regions; PrivateLink for exposing a service to many consumers.

***

## 3) Architectural Differences: GCP VPC vs AWS VPC

| Aspect | AWS VPC | GCP VPC |
| :-- | :-- | :-- |
| Scope | Regional (VPC spans a region; subnets AZ-scoped) | Global (VPC is global; subnets are regional) |
| Peering | Non-transitive; inter-region via inter-region peering | Non-transitive; global across regions by default |
| Routing | Route tables per subnet; explicit associations | Global dynamic routing mode; subnet-level routes auto-create; per-NET/RT limited; uses VPC-wide routing modes |
| Services | IGW, NAT GW, TGW, PrivateLink, VPC Lattice | Cloud NAT, Cloud Router, Shared VPC, Private Service Connect |
| Multi-tenancy | Account-level VPCs; Resource Share (RAM) | Shared VPC across projects (host/service projects) |
| Firewall | Security Groups + NACLs (stateful SG, stateless NACLs) | VPC firewall rules are stateful, hierarchical (network/project) |
| Global load balancers | Mostly regional ALBs/NLBs; Global via CloudFront/GA | Native global L7/L4 load balancers |

Implications:

- GCP’s global VPC simplifies multi-region comms; AWS often needs TGW/inter-region peering.
- AWS SGs are powerful for identity/SG-to-SG references; GCP uses tags/service accounts in rules.

***

## 4) What Are Network Policies?

Kubernetes NetworkPolicies define allowed L3/L4 traffic between pods/namespaces. They require a CNI that supports policies (Calico, Cilium, Antrea, etc.).

### Example: Deny all by default in a namespace, allow only app→db 5432

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny
  namespace: prod
spec:
  podSelector: {}
  policyTypes: [Ingress, Egress]

---
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: app-to-db
  namespace: prod
spec:
  podSelector:
    matchLabels:
      role: db
  ingress:
  - from:
    - podSelector:
        matchLabels:
          role: app
    ports:
    - protocol: TCP
      port: 5432
```

Apply:

```bash
kubectl apply -f policy.yaml
```


***

## 13) What Are Admission Controllers?

Kubernetes admission chain components that intercept API requests to validate or mutate objects after authentication/authorization.

- MutatingAdmissionWebhook: can modify requests (e.g., inject sidecars).
- ValidatingAdmissionWebhook: can accept/reject based on policies.
- Built-ins: NamespaceLifecycle, LimitRanger, ResourceQuota, PodSecurity, etc.

Example: Gate all pods to require specific labels via ValidatingAdmissionWebhook (with OPA Gatekeeper or Kyverno).

Kyverno policy to require labels:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-app-label
spec:
  validationFailureAction: Enforce
  rules:
  - name: check-app-label
    match:
      resources:
        kinds: ["Pod"]
    validate:
      message: "app label required"
      pattern:
        metadata:
          labels:
            app: "?*"
```


***

## 12) Kubernetes Operators

Operators extend Kubernetes with app/domain-specific controllers using CustomResourceDefinitions (CRDs) to encode operational knowledge (backups, failover, upgrades).

Examples:

- Prometheus Operator, Cert-Manager, ArgoCD, Elastic Operator, Redis Operator, Kafka (Strimzi), PostgreSQL (Zalando/CloudNativePG).

Basic structure:

- CRD defines API (e.g., PostgresCluster).
- Controller watches CRs and reconciles actual state.

Install example (Prometheus Operator via kube-prometheus-stack):

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install monitoring prometheus-community/kube-prometheus-stack -n monitoring --create-namespace
```


***

## 5) If Credentials Were Pushed to a Remote Repo

Immediate actions:

1. Revoke and rotate the credentials (do not rely on history rewrite).
2. Invalidate sessions/tokens (AWS: deactivate access keys; GitHub: revoke PAT/SSH).
3. Search for any other secrets in repo and org.

AWS rotation example:

```bash
aws iam update-access-key --user-name alice --access-key-id AKIA... --status Inactive
aws iam create-access-key --user-name alice
```

Git history cleanup:

```bash
# Remove file and rewrite history
git filter-repo --path secrets.env --invert-paths
git push --force
```

Repo scanning:

- GitHub Advanced Security or secret scanning.
- Tools: trufflehog, gitleaks:

```bash
gitleaks detect --source .
trufflehog git file://. --since-commit HEAD~200
```

Notify stakeholders, check logs for abuse, rotate dependent systems (e.g., CI secrets, Terraform cloud vars).

***

## 6) Bring an Existing Resource into Terraform State

Steps:

1. Write resource block in code with correct arguments (minus computed fields).
2. Run import with correct ID.
3. Run plan to reconcile differences; adjust code.

Example: Import an AWS Security Group

```hcl
resource "aws_security_group" "web" {
  name        = "web-sg"
  description = "Web SG"
  vpc_id      = "vpc-0123456789abcdef0"
  # add ingress/egress to match actual
}
```

```bash
terraform init
terraform import aws_security_group.web sg-0123456789abcdef0
terraform plan
```

Tip: For complex resources, use `terraform import` then `terraform plan -out tfplan`, and copy missing computed attributes; consider `terraformer` or `tfimport` tools to bootstrap.

***

## 8) Practices to Reduce Compute Costs Organization-wide

- Rightsizing:
    - Use autoscaling, vertical pod autoscaler (VPA) recommendations, and cluster autoscaler.
    - Analyze CPU/memory with Prometheus/CloudWatch and adjust requests/limits.
- Scheduling efficiency:
    - Bin packing with resource quotas and topology spread; use “Guaranteed/Burstable” wisely.
    - Use Spot/Preemptible for stateless workloads; implement PDB and interruption handlers.
- Scale-to-zero:
    - Knative/KEDA for event-driven workloads.
    - Turn off non-prod at night via schedules (Instance Scheduler, Lambda/CloudWatch Events).
- Instance/Node types:
    - Graviton/ARM where supported; Savings Plans/Reserved Instances for steady workloads.
- Storage/network:
    - Right-size volumes (gp3 vs gp2), lifecycle policies for logs/snapshots, S3 Intelligent-Tiering.
    - Egress minimization: colocate services/data; use PrivateLink/VPC endpoints.
- Container images:
    - Multi-arch slim images reduce pull time and storage.
- Governance:
    - Budgets/alerts; chargeback/showback by tags/labels; cost scorecards in CI.

Example: EKS with Spot and scale-down at night

```yaml
# Karpenter provisioner example snippet
spec:
  requirements:
    - key: "karpenter.sh/capacity-type"
      operator: In
      values: ["spot","on-demand"]
```


***

## 11) How to Back Up a Kubernetes Cluster

Back up both cluster state and application data.

### etcd and cluster objects

- Managed clusters: rely on provider snapshots (EKS control plane managed; take backups of manifests in GitOps).
- Self-managed: snapshot etcd:

```bash
ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 \
  --cacert=/etc/etcd/ca.crt --cert=/etc/etcd/etcd.crt --key=/etc/etcd/etcd.key \
  snapshot save /backup/etcd-$(date +%F).db
```


### Manifests/configs

- Store in Git (GitOps). Periodic `kubectl get all -A -o yaml > cluster-export.yaml` as an extra copy.


### Persistent data

- Use Velero for PV backups to object storage:

```bash
velero install --provider aws --plugins velero/velero-plugin-for-aws:v1.8.0 \
  --bucket k8s-backups --backup-location-config region=us-east-1
velero backup create ns-prod-2025-08-21 --include-namespaces prod
velero restore create --from-backup ns-prod-2025-08-21
```

- App-native backups: databases with operator-native backups (e.g., CloudNativePG, Percona XtraDB, MongoDB OpsManager) to external storage.
- CSI snapshots:

```bash
kubectl create -f volumesnapshotclass.yaml
kubectl create -f volumesnapshot.yaml
```

Run scheduled backups and test restores regularly.

***

## 3b) Bonus: AWS VPC Isolation Using NACL vs SG

- NACLs: stateless, subnet-wide, ordered rules; good for coarse boundaries.
- SGs: stateful, attached to ENIs; preferred for workload-level control.
- Combine: Use NACLs to block broad CIDRs; SGs for pod/instance ingress/egress rules.

***

## 4b) Enforcing Network Policies by Namespace

Default deny in each namespace via label selector:

```yaml
kind: NetworkPolicy
metadata:
  name: default-deny-all
  namespace: staging
spec:
  podSelector: {}
  policyTypes: [Ingress, Egress]
```

Allow DNS egress:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-dns
  namespace: staging
spec:
  podSelector: {}
  egress:
  - to:
    - namespaceSelector: {}
    ports:
    - protocol: UDP
      port: 53
  policyTypes: [Egress]
```


***

## 2b) Rollbacks with Argo Rollouts (Canary)

Example canary with manual promotion:

```yaml
strategy:
  canary:
    steps:
    - setWeight: 10
    - pause: {duration: 5m}
    - setWeight: 50
    - pause: {}
```

Promote:

```bash
kubectl argo rollouts promote my-app
```

Abort:

```bash
kubectl argo rollouts abort my-app
```


***

## 9b) Secure Cross-VPC Connectivity

- Use Security Groups for cross-VPC (SG referencing via TGW/Peering is limited; consider SG referencing with VPC Lattice).
- Restrict routes to least privilege prefixes.
- Encrypt in transit: TLS everywhere; optionally IPsec over TGW or use AWS VPN.

***

## 1b) VPC Endpoints for Private Access

Keep private subnets isolated and still access AWS APIs:

```bash
aws ec2 create-vpc-endpoint --vpc-id vpc-123 --service-name com.amazonaws.us-east-1.s3 --vpc-endpoint-type Gateway --route-table-ids rtb-aaa
aws ec2 create-vpc-endpoint --vpc-id vpc-123 --service-name com.amazonaws.us-east-1.ec2 --vpc-endpoint-type Interface --subnet-ids subnet-1 subnet-2 --security-group-ids sg-endpoints
```


***

## 6b) Importing with for_each

When IDs are many, use import blocks (Terraform 1.5+):

```hcl
resource "aws_s3_bucket" "b" {
  for_each = toset(var.bucket_names)
  bucket   = each.key
}

import {
  to = aws_s3_bucket.b["my-logs"]
  id = "my-logs"
}
```

Then:

```bash
terraform plan
terraform apply
```


***

## 8b) Cost Guardrails in CI/CD

- Policy-as-code (OPA/Conftest/Terraform Cloud policies):
    - Block untagged resources, too-large instance types, gp2 volumes, public EIPs.
- Example Conftest policy (deny m5.24xlarge):

```rego
package main
deny[msg] {
  input.resource.type == "aws_instance"
  input.resource.instance_type == "m5.24xlarge"
  msg := "Instance type not allowed"
}
```


***

## 11b) Backup Tools Overview

- Velero: cluster/PV backups to S3/GCS/Azure.
- Restic: file-level backups for non-snapshot-capable volumes (Velero can integrate).
- Stash by AppsCode, Kasten K10: enterprise features, app-aware backups.
- Database operators with native backups: CloudNativePG, Percona, Crunchy, etc.

***

## 5b) Secret Management Post‑Incident

- Move secrets to a vault: AWS Secrets Manager, HashiCorp Vault, SOPS + KMS.
- Rotate app configs to pull at runtime (sidecars or init containers).
- Use commit hooks to prevent secret commits:

```bash
pre-commit install
# .pre-commit-config.yaml
- repo: https://github.com/gitleaks/gitleaks
  rev: v8.18.0
  hooks:
  - id: gitleaks
```


***

## 3c) GCP vs AWS Networking Services Mapping

- NAT: AWS NAT Gateway vs GCP Cloud NAT.
- Central routing: AWS Transit Gateway vs GCP Cloud Router + VPC peering.
- Private service access: AWS PrivateLink vs GCP Private Service Connect.
- Cross-project: AWS RAM vs GCP Shared VPC.

***

## 4c) CNI Considerations for NetworkPolicy

- Calico: rich policies, eBPF data plane, global network sets.
- Cilium: eBPF-based, L7 policies, Hubble observability.
- Azure/CNI, AWS VPC CNI: need Calico/Cilium/Antrea for full policy features.

***

## 2c) Database Rollbacks

- Prefer forward-fixes with feature flags.
- For schema: use migration tooling (Liquibase/Flyway) with reversible migrations, and back up before irreversible changes.

***

## 1c) Truly Air‑gapped Subnet

- No IGW, no NAT, no VPC endpoints.
- Only TGW to a security enclave or dedicated Direct Connect.
- NACL deny all except specific known CIDRs.

***

## 9c) Cross‑Account Connectivity

- Peering supports cross-account within same/other regions.
- TGW with Resource Access Manager to attach other accounts:

```bash
aws ram create-resource-share --name tgw-share
aws ram associate-resource-share --resource-share-arn arn:... --resource-arn arn:aws:ec2:...:transit-gateway/tgw-123
```


***

## 10b) Preventing Drift

- Prohibit console edits; use IAM SCPs to block changes outside Terraform.
- Use AWS Config + Conformance Packs to detect config changes.
- Scheduled CI `terraform plan` across workspaces; alert on non-zero exit code.

***

## 7b) Zero‑Downtime Upgrade Tips

- Set PDBs and readiness gates; use `maxUnavailable=0` for critical workloads.
- Pre-pull new images to reduce rollout time.
- Stagger AZ upgrades to preserve multi-AZ resilience.

***

## Appendix: Commands Reference

- kubectl rollout:

```bash
kubectl rollout status deploy/api
kubectl rollout undo deploy/api --to-revision=2
```

- Terraform drift check in CI:

```bash
terraform init -input=false
terraform validate
terraform plan -detailed-exitcode -out tfplan || code=$?
if [ "$code" -eq 2 ]; then echo "Drift detected"; exit 2; fi
```

- Velero schedule:

```bash
velero create schedule daily --schedule="0 2 * * *" --ttl 168h
```

- Karpenter install (snippet):

```bash
helm repo add karpenter https://charts.karpenter.sh/
helm install karpenter karpenter/karpenter -n karpenter --create-namespace
```


***

## License

MIT. Contributions welcome.

