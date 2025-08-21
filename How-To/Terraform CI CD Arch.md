Note: The commands and patterns below are proven practices; adapt values to your environment (clusters, regions, accounts, repos).

***

## Architecture of Current Infra: What’s owned vs. managed

- Control plane: Managed (EKS control plane) for HA/patching; own worker nodes, nodegroups, add‑ons, policies, cluster‑addons (VPC CNI, CoreDNS, KubeProxy).
- Networking: Own VPCs, subnets, routing, NAT Gateways, peering/transit, Security Groups, NACLs, Ingress/Egress (ALB/NLB).
- Platform add‑ons: Own cluster‑bootstrap via GitOps/Helm (Ingress Controller, Cert‑Manager, ExternalDNS, Cluster Autoscaler/Karpenter, Metrics Server, Prometheus/Grafana, Loki, Tempo, Jaeger, Kyverno/OPA, EBS/EFS CSI, Fluent Bit).
- Runtime: Own container images, deployments, HPAs, PDBs, PSP replacement (PodSecurity), NetworkPolicies, SSO/RBAC, namespaces, resource quotas/limits, secrets management, service mesh (optional).
- Data: Own RDS/Aurora, ElastiCache, S3, Kafka/MSK (or self‑managed), backups/DR.
- CI/CD: Own pipelines, artifact registries (ECR), environments, rollout strategies, SBOM/supply‑chain (SLSA/COSIGN).
- Observability: Own alert rules, SLOs/Error Budgets, log/metric/trace pipelines, on‑call.
- Security: Own IAM/IAM Roles for Service Accounts (IRSA), KMS, Vault, image scanning, policy as code.
- IaC: Own Terraform modules, state backends, environments, drift detection, change management.

***

## EKS vs. Self‑Managed Kubernetes

- EKS
    - Pros: AWS‑managed control plane upgrades/HA, easy cluster creation, tight IAM/IRSA integration, out‑of‑box CNI (VPC CNI), stability, support SLA.
    - Cons: Less control over control plane flags; cost per cluster; version rollout cadence governed by AWS.
- Self‑Managed
    - Pros: Full control over components/flags/networking; any CNI/CSI, custom API server flags.
    - Cons: Own control plane HA/etcd backups, patching, upgrades, on‑call toil.

Pick EKS unless there’s a hard requirement for bespoke control‑plane config or on‑prem constraints.

***

## CI/CD Rollback Strategy for Monoliths

- Always keep last N artifacts tagged and immutable.
- Blue/Green (keep previous “blue” live until “green” is healthy) or canary with fast fail‑back.
- Database strategy: backward‑compatible migrations (expand → deploy → contract).
- Kubernetes example (Helm):
    - Deploy: helm upgrade --install app charts/app --version 1.2.3
    - Rollback: helm rollback app <REVISION>
- Argo Rollouts:
    - kubectl argo rollouts undo rollout/app
- Plain Deployments:
    - kubectl rollout undo deploy/app --to-revision=12
- DB rollback via migrations:
    - make migrate_down step=1 or liquibase rollbackCount 1

***

## Detecting and Fixing OOMKilled

- Detect:
    - kubectl get pod POD -o jsonpath='{.status.containerStatuses[*].lastState.terminated.reason}'
    - kubectl describe pod POD | grep -i -A3 "State:" | grep -i oom
    - Metrics: container_memory_working_set_bytes spikes near limits.
- Fix:
    - Profile memory, optimize allocations.
    - Set sane requests/limits, avoid limit=0:
        - resources:
requests: { cpu: "250m", memory: "256Mi" }
limits:   { cpu: "1",    memory: "512Mi" }
    - Add liveness only after stability; use HPA on CPU/Custom metrics carefully.
    - Verify JVM/Go memory flags align with cgroup limits.

***

## NAT Gateway vs. VPC Peering

- NAT Gateway: Private subnets’ egress to Internet (managed, per‑AZ billing). One‑way egress, no inbound.
- VPC Peering: Private, non‑transitive connectivity between VPCs. No NAT, no IGW needed. Subject to CIDR non‑overlap, routing/Security Groups. Not a replacement for NAT to Internet; use for service‑to‑service across VPCs.

***

## kubectl exec vs logs vs describe — When to use what

- kubectl logs: Inspect app logs, crash loops, init/sidecars:
    - kubectl logs deploy/app
    - kubectl logs pod/p -c container --previous
- kubectl describe: State and events: scheduling, image pulls, probes, reasons:
    - kubectl describe pod p
    - kubectl describe deploy d
- kubectl exec: Interactive diagnosis inside container for runtime/env issues (not a logging tool):
    - kubectl exec -it p -- sh
    - Prefer ephemeral debug containers when distroless: kubectl debug p -it --image=busybox

***

## ConfigMap vs Secret Use Cases

- ConfigMap: Non‑sensitive config (feature flags, templates, env, JSON/YAML). Mount or envFrom. Not encrypted by default.
- Secret: Sensitive material (passwords, tokens, keys). Use KMS/Vault, enable encryption at rest, mount via projected volume; prefer IRSA over long‑lived creds.

***

## What breaks if readinessProbe is wrong?

- Traffic routing: Pods marked Unready → load balancer/Service won’t send traffic (good) or, if probe succeeds too early, traffic hits an unready app → 5xx, timeouts.
- HPA/Autoscaling: Readiness failures can cause churn if combined with liveness misconfigs.
- Rollouts: Deployments can hang waiting for readiness; canary weight may stick; blue/green cutover stalls.

Checklist:

- Use a lightweight health endpoint that checks downstream deps minimalistically.
- initialDelaySeconds covers warm‑up; successThreshold=1; use timeoutSeconds small (1‑2s), periodSeconds 5‑10s.

***

## One‑liner to find open ports on a Linux node

- ss (modern):
    - sudo ss -tulpn
- netstat (legacy):
    - sudo netstat -tulpn
- nmap from another host:
    - nmap -sT -p- NODE_IP

***

## NetworkPolicy Debugging Steps

- Confirm plugin supports policies (Calico/Cilium/Weave).
- List policies and namespaces:
    - kubectl get netpol -A
    - kubectl describe netpol NAME -n NS
- Visualize effective policy:
    - kubectl exec -it POD -- curl -m2 SVC:PORT
    - From busybox: wget -S -T2 http://IP:PORT
- Use policy audit tools (e.g., cilium monitor, calico flow logs).
- Temporarily allow all egress/ingress to bisect:
    - Create a permissive policy; tighten gradually.
- Verify namespaceSelectors/labels match pods:
    - kubectl get ns --show-labels
    - kubectl get pod -L app,tier

***

## K8s app crashes during image pull — triage

- Describe pod for reasons:
    - kubectl describe pod p | sed -n '/Events/,\$p'
    - Common: ImagePullBackOff, ErrImagePull, Unauthorized.
- Check image ref and tag digest:
    - kubectl get pod p -o jsonpath='{.spec.containers[*].image}'
- Validate pull secrets and permissions:
    - kubectl get secret -n NS
    - kubectl create secret docker-registry regcred --docker-server=... --docker-username=... --docker-password=...
    - Patch SA: kubectl patch sa default -n NS -p '{"imagePullSecrets":[{"name":"regcred"}]}'
- Check registry reachability/DNS:
    - kubectl run tmp --rm -it --image=alpine -- ash -c "apk add curl; curl -v https://registry/v2/"
- Node/image cache issues:
    - kubectl get nodes; kubectl describe node NODE
    - SSH to node: sudo crictl images; sudo crictl rmi IMAGE or systemctl restart containerd (last resort).
- If tag mutable, pin digest:
    - image: repo/app@sha256:...

***

## Debugging Stuck Pods \& Taints/Tolerations

- Pending pods:
    - kubectl describe pod p | sed -n '/Events/,\$p'
    - kubectl get nodes -o custom-columns=NAME:.metadata.name,TAINTS:.spec.taints
- Check taints:
    - NoSchedule/PreferNoSchedule/NoExecute requiring tolerations.
    - Add toleration:
        - tolerations:
            - key: "workload"
operator: "Equal"
value: "gpu"
effect: "NoSchedule"
- Resource fit:
    - kubectl describe node NODE | egrep -i "Allocatable|Capacity"
    - Reduce requests/limits or change nodeSelector/topologySpreadConstraints.
- PVC binding:
    - kubectl get pvc -n NS; kubectl describe pvc PVC
- Affinity conflicts:
    - Remove conflicting nodeAffinity/podAntiAffinity.
- Image pull issues: see prior section.

***

## Multi‑Account Terraform with Backend Isolation

- Separate AWS accounts per env (dev/stage/prod), separate S3 buckets and KMS keys per env/workspace, least‑privilege IAM roles per pipeline.
- Example backend per env:
    - backend "s3" {
bucket         = "tfstate-prod-1234"
key            = "network/primary.tfstate"
region         = "us-east-1"
dynamodb_table = "tfstate-locks-prod"
kms_key_id     = "alias/tfstate-prod"
}
- Assume‑role via OIDC in CI:
    - aws sts assume-role --role-arn arn:aws:iam::ACCOUNT:role/terraform --role-session-name ci
- Folder layout:
    - live/
        - prod/
        - stage/
        - dev/
    - modules/
- Use terraform workspaces only for minor variants; prefer full backend isolation per account.

***

## Staging Cost Exploded due to HPA — RCA

- Symptoms: HPA scale‑out to max replicas due to noisy CPU/latency metrics, traffic spike, or bad probe.
- Steps:
    - Inspect HPA events:
        - kubectl describe hpa app
    - Check metrics source:
        - If custom metrics (Prometheus Adapter), confirm query is correct and series not missing or NaN.
    - Verify requests vs actual load:
        - Low CPU requests + Cluster Autoscaler/Karpenter = many small pods and new nodes.
    - Check load generators/test data and readiness/liveness flapping.
    - Audit recent releases/config changes.
- Fixes:
    - Set sane min/max replicas and cool‑downs:
        - behavior.scaleUp.policies, stabilizationWindowSeconds.
    - Use targetUtilization with realistic requests.
    - Protect staging with cluster‑autoscaler maxNodesTotal, budget caps, and scheduled HPA min=0 off‑hours.
    - Alert on HPA recommendation spikes and node creation rate.

***

## Dockerfile Best Practices for Go Apps

- Multi‑stage, reproducible, small images:
    - Use Go build stage with CGO disabled unless needed.
    - Pin versions and use go mod download caching.
    - Use distroless or alpine runtime as appropriate.
- Example:

```
# syntax=docker/dockerfile:1.6
FROM golang:1.22 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -trimpath -o /out/app ./cmd/app

FROM gcr.io/distroless/static:nonroot
USER nonroot:nonroot
COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
```

- Add SBOM/signing:
    - cosign sign --key cosign.key IMAGE
    - syft packages IMAGE -o json > sbom.json

***

## Blue/Green vs. Canary — When to pick

- Blue/Green:
    - Two environments; switch traffic when green is healthy.
    - Pros: Instant rollback, easy DB migrations with dual‑write/read strategies.
    - Cons: Double infra at cutover; coarse‑grained.
    - Use when: monoliths, large changes, strict uptime, simple traffic switch needed.
- Canary:
    - Gradual traffic shifting (1%→5%→25%→100%), automated metrics gates.
    - Pros: Risk reduction, real‑user testing, auto‑abort.
    - Cons: Requires observability and progressive delivery tooling.
    - Use when: frequent releases, microservices, strong SLOs/metrics gates.

***

## Cross‑Region S3 Replication as Code (Terraform)

- Requirements: Versioning enabled, IAM role for replication, distinct destination bucket/region, KMS keys if encrypted.
- Snippets:

```
resource "aws_s3_bucket" "src" { bucket = "my-src"; region = "us-east-1" }
resource "aws_s3_bucket_versioning" "src" { bucket = aws_s3_bucket.src.id  versioning_configuration { status = "Enabled" } }
resource "aws_s3_bucket" "dst" { bucket = "my-dst"; provider = aws.usw2 }
resource "aws_s3_bucket_versioning" "dst" { bucket = aws_s3_bucket.dst.id  versioning_configuration { status = "Enabled" } }

resource "aws_iam_role" "replication" {
  name = "s3-replication"
  assume_role_policy = data.aws_iam_policy_document.s3_replication_trust.json
}

resource "aws_s3_bucket_replication_configuration" "rep" {
  bucket = aws_s3_bucket.src.id
  role   = aws_iam_role.replication.arn
  rule {
    id     = "crr"
    status = "Enabled"
    destination { bucket = aws_s3_bucket.dst.arn  storage_class = "STANDARD" }
    delete_marker_replication { status = "Enabled" }
  }
}
```

- Remember KMS key grants for both source/dest if SSE‑KMS.

***

## Real War Room: Service Down, No Alerts, Logs Green — What next?

- Verify blast radius:
    - Check SLO dashboards (latency, error rate, saturation), dependency dashboards, upstream/downstream.
- End‑to‑end synthetic checks:
    - Execute scripted transaction from multiple regions (RUM/synthetic).
- Correlate recent changes:
    - git log --since "2h"; kubectl get events -A; cloudtrail for IAM/SG changes; feature flag flips.
- Network/DNS:
    - dig app.domain; nslookup service.cluster.local; curl --resolve to bypass DNS; check TLS expiry.
- Infra signals:
    - Load balancer 5xx/target health; node readiness; throttling; rate limits; WAF.
- Disable canaries/feature flags; roll back last deployment; route around failing AZ/region.
- Open a comms channel, note timestamps, owners; start incident doc.

***

## Vault + GitHub Actions Integration (Secure)

- Use OIDC‑based auth (no static GitHub secrets).
- Steps:
    - Enable JWT auth in Vault; configure trusted GitHub OIDC issuer and claims (repo, branch, env).
    - Map roles to Vault policies with tight constraints (aud, sub, ref, workflow).
    - GitHub Actions job requests Vault token using OIDC JWT, reads short‑lived secrets/dynamic DB creds.
- Example policy and workflow:
    - Vault role (pseudo):
        - bound_subject="repo:org/repo:ref:refs/heads/main"
        - token_ttl=5m
        - policies=["ci-read"]
    - GitHub Actions:

```
permissions:
  id-token: write
  contents: read
steps:
  - uses: hashicorp/vault-action@v2
    with:
      method: jwt
      url: ${{ secrets.VAULT_ADDR }}
      role: github-ci
      secrets: |
        kv/data/ci DOCKER_PASSWORD | DOCKER_PASSWORD ;
        database/creds/readonly username | DB_USER ;
        database/creds/readonly password | DB_PASS
```

- Prefer dynamic secrets engines; audit via Vault.

***

## Pod‑to‑Pod TLS in Kubernetes

- Options:
    - Service Mesh (mTLS): Istio, Linkerd, Consul. Easiest org‑wide encryption + identity + policy.
    - CNI‑level encryption (Cilium WireGuard) for node‑to‑node and pod‑to‑pod.
    - App‑level TLS if mesh not desired.
- Quick start with Linkerd:
    - linkerd install | kubectl apply -f -
    - linkerd check
    - annotate namespace for auto‑inject; mTLS enabled by default between meshed pods.
- Cilium:
    - cilium install --encryption wireguard

***

## Incident Postmortem Format

- Title, severity, dates/times, owners.
- Summary: 1‑2 paragraphs, user impact, duration.
- Timeline: precise UTC timestamps, who/what, automation/manual actions.
- Root Cause: primary and contributing factors; why chain (5 Whys).
- Detection: how it was found, why alerts missed, MTTA/MTTR.
- Remediation: immediate fixes, rollbacks.
- Preventive actions: concrete, owners, due dates.
- Learnings: what worked/failed (people/process/tools).
- Attachments: graphs, logs, PRs, runbooks.

***

## One Terraform Command to Detect Drift

- terraform plan -detailed-exitcode
    - Exit codes: 0 no change, 2 changes present (drift or pending apply).
- For read‑only check in CI:
    - terraform plan -detailed-exitcode -lock=false -input=false

***

## Handling Terraform State in a Remote Team

- Use remote backend (S3+DynamoDB lock or Terraform Cloud) per env.
- Server‑side encryption with KMS; versioning enabled.
- Enforce locking (DynamoDB); least‑privilege IAM for CI.
- PR‑based plans with artifacted plan files; only merges can apply.
- State move/splits via terraform state mv and import with approvals.
- Backups and state retention policies; restricted access.

***

## Simulate a DNS Failure in CoreDNS

- Break config safely in a test namespace/cluster:
    - Edit CoreDNS configmap:
        - kubectl -n kube-system edit configmap coredns
        - Introduce bad upstream or block a domain using the "rewrite" or "block" plugins (if enabled).
- Example to blackhole a domain:
    - In Corefile:

```
.:53 {
  errors
  health
  kubernetes cluster.local in-addr.arpa ip6.arpa {
    pods insecure
    upstream
    fallthrough in-addr.arpa ip6.arpa
  }
  hosts {
    0.0.0.0 bad.example.com
    fallthrough
  }
  forward . 127.0.0.1
  cache 30
  loop
  reload
  loadbalance
}
```

- Test:
    - kubectl run -it dns-test --rm --image=alpine -- ash -c "apk add bind-tools; dig bad.example.com +short"
- Rollback by reverting the ConfigMap and restarting pods:
    - kubectl -n kube-system rollout restart deploy/coredns

***

# Commands \& Snippets Reference

## kubectl CrashLoop/ImagePull triage

- kubectl get pods -n NS
- kubectl describe pod POD -n NS
- kubectl logs POD -n NS --previous
- kubectl get events -n NS --sort-by=.lastTimestamp
- kubectl get sa default -n NS -o yaml | yq '.imagePullSecrets'
- kubectl run netcheck --rm -it --image=appropriate/curl -- curl -kv https://REGISTRY/v2/


## Taints \& Scheduling

- kubectl get nodes -o wide
- kubectl get nodes -o custom-columns=NAME:.metadata.name,TAINTS:.spec.taints
- kubectl taint nodes NODE key=value:NoSchedule
- Add toleration in deployment as shown earlier.


## NetworkPolicy Smoke Tests

- Busybox to test:
    - kubectl run bb --rm -it --image=busybox -- sh
    - wget -S -T2 http://SVC:PORT
- Temporary allow‑all policy:

```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata: { name: allow-all, namespace: NS }
spec:
  podSelector: {}
  ingress: [ {} ]
  egress:  [ {} ]
  policyTypes: [ Ingress, Egress ]
```


## Readiness/Liveness Example

```
livenessProbe:
  httpGet: { path: /healthz, port: 8080 }
  initialDelaySeconds: 20
  timeoutSeconds: 2
  periodSeconds: 10
  failureThreshold: 3

readinessProbe:
  httpGet: { path: /ready, port: 8080 }
  initialDelaySeconds: 10
  timeoutSeconds: 1
  periodSeconds: 5
  successThreshold: 1
```


## HPA with Stabilization

```
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
spec:
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 60
  behavior:
    scaleUp:
      stabilizationWindowSeconds: 120
      policies:
      - type: Percent
        value: 100
        periodSeconds: 60
    scaleDown:
      stabilizationWindowSeconds: 300
      policies:
      - type: Percent
        value: 50
        periodSeconds: 60
```


## Example Deployment with Resources/Probes/IRSA

```
apiVersion: apps/v1
kind: Deployment
metadata: { name: app, labels: { app: app } }
spec:
  replicas: 3
  selector: { matchLabels: { app: app } }
  template:
    metadata:
      labels: { app: app }
      annotations:
        iam.amazonaws.com/role: arn:aws:iam::123456789012:role/app-irsa
    spec:
      serviceAccountName: app-sa
      containers:
      - name: app
        image: 123456789012.dkr.ecr.us-east-1.amazonaws.com/app@sha256:...
        ports: [ { containerPort: 8080 } ]
        resources:
          requests: { cpu: "250m", memory: "256Mi" }
          limits: { cpu: "1", memory: "512Mi" }
        envFrom: [ { configMapRef: { name: app-config } }, { secretRef: { name: app-secrets } } ]
        livenessProbe: { httpGet: { path: /healthz, port: 8080 }, initialDelaySeconds: 20 }
        readinessProbe: { httpGet: { path: /ready, port: 8080 }, initialDelaySeconds: 10 }
      nodeSelector: { workload: general }
      tolerations:
      - key: "workload"
        operator: "Equal"
        value: "general"
        effect: "NoSchedule"
```


## Terraform Drift/Plan in CI

- Init with explicit backend:
    - terraform init -backend-config=env/prod.backend.hcl
- Validate/Format:
    - terraform fmt -check
    - terraform validate
- Drift detect:
    - terraform plan -detailed-exitcode -out=plan.tfplan || code=$?; if [ "$code" -eq 2 ]; then echo "Drift or changes"; exit 2; else exit \$code; fi
- Safe apply via manual approval of plan artifact.

***

# Blue/Green and Canary with Argo Rollouts

- Canary example:

```
apiVersion: argoproj.io/v1alpha1
kind: Rollout
spec:
  strategy:
    canary:
      steps:
      - setWeight: 5
      - pause: { duration: 2m }
      - setWeight: 25
      - pause: { duration: 5m }
      - setWeight: 50
      - pause: {}
      analysis:
        templates:
        - templateName: error-rate-check
```

- Blue/Green example:

```
strategy:
  blueGreen:
    activeService: app-svc
    previewService: app-svc-preview
    autoPromotionEnabled: false
```


***

# S3 CRR with KMS (Important bits)

- Enable SSE‑KMS on both buckets; grant replication role “kms:Encrypt/Decrypt/GenerateDataKey” on both keys.
- Allow destination bucket policy to accept source account’s replication role.

***

# Open Ports One‑Liners

- ss/netstat on node:
    - sudo ss -lntup | sort -k5
- Top 10 listeners:
    - sudo ss -lntup | awk 'NR>1{print \$5}' | cut -d: -f2 | sort | uniq -c | sort -nr | head

***

# Realistic RCA Template for HPA Cost Spike

- What changed: HPA target from 70%→50%; CPU requests lowered; traffic generator left on.
- Impact: Nodes grew from 6→42; cost +6x over 4h.
- Why chain:
    - Lower requests → higher utilization readings per pod → scale out.
    - Metrics adapter returned stale values → oscillation.
    - CA unconstrained → provisioned max nodes.
- Fix:
    - Add CA maxNodesTotal=20 in staging; HPA minReplicas=0 off‑hours; add stabilization windows; alerting on node add rate.

***

# End‑to‑End Encryption: Cilium WireGuard

- cilium install --encryption wireguard
- cilium status
- Verify:
    - cilium config view | grep encryption
    - tcpdump on node shows UDP traffic to WG port.

***

# War‑Room Commands Cheat Sheet

- DNS/TLS:
    - openssl s_client -connect host:443 -servername host -showcerts < /dev/null | openssl x509 -noout -dates
    - dig +trace domain.com
- LB health:
    - aws elbv2 describe-target-health --target-group-arn ...
- K8s health:
    - kubectl get nodes,pods -A -o wide
    - kubectl get events -A --sort-by=.lastTimestamp
    - kubectl top nodes; kubectl top pods -A

***

# Appendix: Quick Answers

- NAT Gateway vs VPC Peering: egress to Internet vs private VPC‑to‑VPC routing.
- kubectl exec/logs/describe: exec for interactive shell, logs for app output, describe for state/events.
- ConfigMap vs Secret: non‑sensitive vs sensitive with encryption/short‑lived access.
- Wrong readinessProbe breaks rollouts/traffic; cause flapping or black‑hole routing.
- One‑liner open ports: sudo ss -tulpn
- NetworkPolicy debug: confirm plugin, list policies, label matches, temporary allow‑all, flow logs.
- Terraform: drift with terraform plan -detailed-exitcode; remote state S3+DDB+KMS; per‑env backends.
- CoreDNS failure simulation: edit ConfigMap to blackhole domain, test with dig, rollout restart.

***

This README is structured for direct use in GitHub. Replace placeholders (namespaces, ARNs, bucket names, domains) before publishing.

