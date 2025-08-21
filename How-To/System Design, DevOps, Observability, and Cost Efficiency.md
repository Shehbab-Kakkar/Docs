
# System Design, DevOps, Observability, and Cost Efficiency — A Practical Guide

This README provides a concise, battle-tested reference for modern backend/platform engineers. It covers scalable microservices, DevOps automation, observability, resilience, and cloud cost optimization with actionable patterns, pitfalls, and checklists.

***

## 🔹 System Design \& Reliability

### 1) Designing a Scalable Microservices Architecture

- Decompose by business domain (DDD): define bounded contexts and API contracts; avoid entity-driven splitting that causes chatty services.
- Choose service granularity pragmatically: start coarser; split only when scaling, team ownership, or change velocity demands it.
- API design: prefer async events (Kafka, SNS/SQS, Pub/Sub) for decoupling; use REST/gRPC for synchronous calls; establish versioning and backward compatibility.
- Data ownership: one service owns its data; use CDC and event sourcing where appropriate; avoid distributed transactions—prefer Saga patterns for long-lived workflows.
- Service discovery and config: service mesh (Istio/Linkerd) or managed (EKS/ECS + CloudMap, GKE + Traffic Director); centralized config (Consul, Spring Cloud Config, AWS AppConfig).
- Reliability patterns: circuit breakers, retries with exponential backoff and jitter, bulkheads, timeouts, idempotency keys, rate limiting.
- Observability: structured logs, metrics (RED/USE), traces; propagate correlation IDs.
- Security: mTLS inside mesh, JWT/OIDC at edges, secrets via KMS/SM/HashiCorp Vault; least-privilege IAM, network segmentation.
- Delivery: blue/green or canary + gradual traffic shifting; infra as code (Terraform), GitOps (Argo CD/Flux).
- Platform foundations: container orchestration (Kubernetes), standard base images, sidecars for mesh/telemetry, golden paths via templates.
- Data access patterns: read replicas, CQRS for read-heavy domains, caching (Redis) with careful TTLs and cache invalidation strategies.

Checklist:

- Clear domain boundaries
- Async-first communication where possible
- Resilience defaults (timeouts, retries, CB)
- Centralized observability and policy
- Zero-trust networking

***

### 2) Capacity Planning Approach

- Establish SLOs and growth assumptions: target p99 latency, error budget, expected QPS/tps, data growth/month.
- Measure baseline: load test per service; derive RPS per pod/instance, CPU/memory profiles, and headroom.
- Model demand: conversions, diurnal/seasonal patterns, marketing events; use percentiles not averages.
- Translate to capacity: replicas = peak_load / per_instance_capacity × safety_factor (typically 1.3–2.0).
- Include dependencies: databases, caches, queues, storage, NAT gateways; ensure downstream can scale.
- Plan buffers: handle regional failover (N+1 regions), AZ failure (N+1 zones), and noisy neighbor effects.
- Reassess continuously: autoscaling with guardrails; monthly reforecast with last 90-day usage and booked events.
- Cost-aware: evaluate instance types vs rightsizing; spot/preemptible for stateless; savings plans/committed use for steady baselines.

Artifacts:

- Demand model spreadsheet
- Load test reports
- Capacity runbooks per service
- Scaling policies and limits

***

### 3) Stateful vs Stateless Services

- Stateless: no user/session state in memory; safe to scale horizontally; ideal for HTTP APIs, workers; state externalized (DB, cache, object store).
- Stateful: holds session or durable state (databases, queues, stateful stream processors, WebSocket hubs); scaling requires sharding/replication and careful placement.
- Operational implications:
    - Stateless: easy blue/green, HPA-friendly, cheap to autoscale.
    - Stateful: needs PVs, quorum, consistency configs, backup/restore, stricter upgrade plans.
- Design guidance: default to stateless; for stateful needs, prefer managed services where possible.

***

### 4) Designing Multi-Region Cloud Deployments

- Drivers: latency, availability (region failure tolerance), data residency.
- Topologies:
    - Active/Passive: primary serves; secondary warm standby with async replication; simple, higher RTO/RPO.
    - Active/Active: traffic steered by geo-DNS/Anycast; requires conflict resolution and idempotent writes.
- Data strategies:
    - Read-heavy: global read replicas; write to home region.
    - Multi-leader/CRDT/event sourcing when multi-write is needed; resolve conflicts deterministically.
    - Compliance: keep PII in-region; tokenize or reference data cross-region.
- Traffic management: Route53/Cloud DNS/Traffic Manager; health checks; failover policies; session affinity minimized.
- Platform:
    - Per-region K8s clusters; image and config promotion; region-scoped secrets and KMS keys.
    - Global services: CDN, WAF, edge auth; per-region runtime configs.
- Failure drills: game days simulating region loss; verify RTO/RPO; practice DB promotion.
- Observability: region-tagged metrics/logs; synthetic checks per region.

***

### 5) Role and Effective Configuration of Load Balancers

- Roles: distribute traffic, health check endpoints, terminate TLS, enforce WAF, session policies, and routing (L7 path/host rules).
- Types:
    - L4 (NLB/TCP) for low latency and non-HTTP protocols.
    - L7 (ALB/HTTP/S) for HTTP routing, header-based rules, JWT validation at edge.
    - Global anycast/edge LB for multi-region steering.
- Best practices:
    - Health checks per path/version; fail-fast with sensible thresholds.
    - Timeouts aligned with upstreams; set connection pooling and keep-alive.
    - Circuit breaking and outlier detection (Envoy/Istio).
    - TLS: modern ciphers, cert rotation, HSTS; mutual TLS internally.
    - Observability: access logs, request IDs, WAF logs; percentiles on backend latency.
    - Security: WAF rules, bot management, rate limiting, IP allow/deny where relevant.

***

## 🔹 DevOps \& Automation

### 6) Enforcing Security in CI/CD Pipelines

- Principles: shift-left, least privilege, reproducible builds, provenance, and separation of duties.
- Controls:
    - Source: branch protections, signed commits (GPG/Sigstore), CODEOWNERS, dependency update bots.
    - Build: SAST, dependency scanning (SBOM, vulnerability scans), secret scanning; build in isolated runners; sign artifacts (SLSA, cosign).
    - Container: minimal base images, rootless, distroless, scan with Trivy/Grype; enforce non-root users.
    - Infra as Code: scan Terraform/K8s manifests with Checkov, tfsec, kube-score, OPA/Gatekeeper.
    - Policy gates: PR checks required; block on critical vulns unless exception approved.
    - Deploy: admission controllers (OPA/Gatekeeper/Kyverno), verify signatures, enforce namespace/network policies.
    - Runtime: eBPF/Falco rules, workload identity, rotate secrets, periodic pen tests.
- Secrets: use Vault/Secret Manager; never in repo; short TTL dynamic creds.
- Provenance: generate SBOMs, attestations; store in artifact registry.

***

### 7) Policy as Code and OPA (Open Policy Agent)

- Concept: encode organizational rules (security, compliance, cost, reliability) as machine-enforced policies across the stack.
- OPA:
    - Engine using Rego policy language; evaluates input JSON against rules to allow/deny or mutate.
    - Integrations: Kubernetes (Gatekeeper), CI (policy checks), API gateways, Terraform (OPA/Terraform Cloud), custom services via sidecar/library.
- Examples:
    - K8s admission: deny images without signature, containers running as root, no resource limits set.
    - Terraform guardrails: prevent public S3 buckets; require tags; restrict instance types.
    - CI gates: block deploys from non-main branches or without approvals.
- Benefits: consistency, auditability, version control, automated enforcement; reduces manual review toil.

***

### 8) Push-Based vs Pull-Based Deployments

- Push-based: CI/CD system pushes artifacts to targets (kubectl apply, SSH, cloud deploy APIs).
    - Pros: simple, fast; Cons: requires target credentials in CI; harder to audit drift; blast radius larger.
- Pull-based (GitOps): cluster-side agents (Argo CD/Flux) continuously reconcile desired state in Git.
    - Pros: strong audit trail, least-privilege (cluster pulls), drift detection and auto-heal, rollbacks via git revert.
    - Cons: added components; eventual consistency; requires repo hygiene and promotion workflows.
- Guidance: prefer GitOps for K8s; use push for immutable serverless or PaaS with minimal secrets in CI.

***

### 9) Zero-Downtime Database Migrations

- Principles: backward-compatible, multi-step, idempotent, with feature flags.
- Expand/contract pattern:

1) Expand: add new columns/tables/indices; backfill asynchronously; deploy app supporting both schemas.
2) Migrate traffic: write to both (dual-write) or read via fallback; verify parity.
3) Cutover: switch reads/writes; monitor; keep old path as fallback.
4) Contract: remove legacy columns/paths after soak period.
- Avoid risky ops: blocking ALTER TABLE on large tables; use online schema change tools (gh-ost/pt-osc for MySQL, pg_repack/PG online DDL approaches).
- Transactions: keep migrations small; use transactional DDL if supported.
- Rollbacks: always safe to roll back app without breaking due to schema incompatibility.
- Data integrity: checksums, row counts, sampling to validate backfills.
- Indexing: create concurrently where supported to avoid locks.

***

### 10) Immutable Infrastructure (with Examples)

- Definition: infrastructure and images are never modified in place; changes occur by replacing with new, versioned artifacts.
- Implementations:
    - Machine images: Packer builds AMIs/VM images; deploy via ASGs/MIGs; old nodes drained and terminated.
    - Containers: build versioned images; deploy via blue/green/canary in Kubernetes; no “ssh fix” allowed.
    - Serverless: new versions/aliases; shift traffic incrementally.
- Benefits: consistency, easy rollbacks, reduced config drift, improved security posture.
- Requirements: image pipelines, configuration baked or injected declaratively, secrets via runtime identity.

***

## 🔹 Monitoring \& Observability

### 11) Implementing Distributed Tracing (Jaeger, OpenTelemetry)

- Instrumentation:
    - Use OpenTelemetry SDKs/auto-instrumentation to emit traces, metrics, logs; propagate W3C tracecontext (traceparent).
    - Tag spans with service.name, version, http.method/url/status, db.system/statement (sanitized), messaging attributes.
- Collection:
    - Run OTel Collector as a sidecar/daemonset/gateway; receive OTLP, batch, tail-sample, export to Jaeger/Tempo/X-Ray/Cloud Trace.
- Storage/Visualization:
    - Jaeger as backend/UI or vendor backends; enforce retention policies and sampling rates (e.g., 1–10% baseline + tail-based triggers).
- Sampling:
    - Head-based for simplicity; tail-based for capturing slow/errors; dynamic sampling rules by endpoint, status, or SLO burn.
- Correlation:
    - Inject trace IDs into logs; link metrics exemplars; create service dependency graphs.
- Governance:
    - Tracing SLAs per team; lint span names; PII redaction; budgets for storage.

***

### 12) Monitoring vs Observability

- Monitoring: collecting known signals and thresholds to detect known failure modes (dashboards, alerts on CPU, latency, errors).
- Observability: ability to answer unknown questions using rich telemetry (high-cardinality logs/metrics/traces) without predefining all issues.
- Practice:
    - Monitoring gives “is it broken?”; observability enables “why is it broken?”
    - Combine SLOs/error budgets (monitoring) with exploratory tools (tracing/log search) for deep diagnostics.

***

### 13) Debugging Latency in a Distributed System

- Start with SLOs and golden signals (latency, errors, saturation, traffic).
- Use distributed traces to locate the critical path: identify slow spans, retries, N+1 calls, or chatty cross-service hops.
- Check per-hop:
    - Client-side: timeouts, retries, connection pooling.
    - Network: DNS latency, TLS handshakes, LB queuing.
    - Service: GC pauses, CPU throttling, lock contention, thread pool saturation.
    - Data: slow queries, missing indices, hot partitions, cache misses, eviction storms.
    - Dependencies: external APIs rate-limiting or slow responses.
- Techniques:
    - Compare p50 vs p99 to separate tail latency; look for coordinated omission.
    - Use heatmaps; analyze concurrency vs latency; verify autoscaler behavior at edges.
    - Introduce caching, batch calls, paginate, precompute; reduce serialization overhead (gRPC + protobuf).
    - Apply backpressure and limit concurrency; tune pools and queues.
- Preventive:
    - Load tests with realistic think time; chaos and latency injection; profile regularly.

***

## 🔹 Resilience \& Cost Efficiency

### 14) Autoscaling Anti-Patterns

- Scaling on the wrong metric (CPU for I/O-bound workloads, or queue length without normalizing by consumers).
- No cooldowns or too-aggressive policies causing thrash.
- Scale-to-zero without warm paths or connection priming, causing cold start storms.
- Ignoring downstream capacity (DB/cache) leading to cascading failures.
- Large pods/instances limiting bin-packing; using maxed-out nodes blocking scale-out.
- Missing limits/requests leading to CPU throttling or OOM kills.
- Not testing HPA policies under load; assuming linear scaling.
- Stateful workloads auto-scaling as if stateless.
- Single-AZ node groups; insufficient pod disruption budgets (PDBs) causing outages during scale events.

***

### 15) Cost Monitoring and Optimization (AWS/Azure/GCP)

- FinOps foundations:
    - Tagging/labels: owner, env, cost-center, app, compliance. Enforce in CI with policy as code.
    - Budgets and alerts: per project/account/team with anomaly detection.
    - Showback/chargeback dashboards to drive accountability.
- Right-sizing:
    - Use recommender tools (AWS Compute Optimizer, Azure Advisor, GCP Recommender).
    - Reduce overprovisioned CPU/memory; tune K8s requests/limits; prefer smaller instance sizes for better bin-packing.
- Commitment discounts:
    - AWS Savings Plans/Reserved Instances; Azure Reservations; GCP Committed Use Discounts.
    - Cover steady-state baseline (60–80%) with commitments; keep headroom on on-demand/spot.
- Storage:
    - Tiering and lifecycle policies (S3/Blob/GCS): move cold data to infrequent access/Archive; delete duplicates; compress.
    - Optimize EBS/Premium/PD types; switch from gp3/premium SSD where IOPS not needed.
- Data transfer:
    - Minimize cross-AZ/region egress; use private links/peering; cache at edge; bundle requests.
- Databases/analytics:
    - Choose serverless autoscaling (Aurora/RDS serverless, BigQuery slots autoscaling) for variable loads.
    - Turn off dev/test at night; use smaller shards; optimize queries and indices; partition/prune scans.
- Kubernetes:
    - Cluster autoscaler with multiple node groups; spot for stateless (with PDBs and disruption tolerance).
    - Use vertical pod autoscaler recommendations; limit over-requesting.
- Governance:
    - Guardrails for expensive instance types, GPU quotas; pre-approval for large resources.
    - Automatic TTL for ephemeral environments and preview deployments.

***

## Appendix: Practical Templates

### Scaling Defaults (Kubernetes)

- Readiness/liveness probes per versioned health endpoint
- Timeouts: client 2–5s, upstream 1–2s, DB 1–3s; retries: 2–3 with jitter
- Resource requests: start from measured p95; limits at 1.5–2× requests
- HPA: target 60–70% on meaningful metric (RPS per pod, queue depth per consumer)


### SLO Starter

- API: 99.9% success over 30d; p99 latency <300ms; error budget policy with freeze at 2× burn
- Batch: completion within window 99%; queue time p95 <1min


### Migration Safety List

- Backward compatible schema
- Online DDL
- Feature flag toggles
- Backfill with progress and checksums
- Rollback plan and data snapshot
- Observability focus on error/latency deltas during cutover

***

License: MIT (adjust as needed)

