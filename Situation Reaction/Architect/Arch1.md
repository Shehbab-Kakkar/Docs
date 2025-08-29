```go
Short answer: Place an external AWS API Gateway at the edge for cross-cutting API concerns, put an internet-facing ALB or NLB as the Kubernetes entry Load Balancer managed by the AWS Load Balancer Controller, and use Kubernetes Ingress inside EKS to route to microservices; for private or hybrid needs, use API Gateway with VPC Link to a private NLB in front of an internal ALB/Ingress in EKS. This pattern cleanly separates edge concerns (auth, rate limiting) from cluster routing (host/path rules) and scales for Java, Go, React, DB, MQ, and ML workloads.
Recommended placement
Edge: Put API Gateway in front to provide a single public entry point, auth (JWT/Lambda authorizers), throttling, API keys, caching, usage plans, and versioning; integrate it with EKS via either HTTP proxy to ALB or Private Integration using VPC Link to an internal NLB for private EKS ingress. This lets edge policies be enforced uniformly across all services.
Cluster ingress: Use the AWS Load Balancer Controller on EKS to provision an ALB for Kubernetes Ingress (Layer 7) or an NLB for Service type LoadBalancer (Layer 4), depending on protocol needs; ALB is typical for HTTP(S) microservices with host/path routing, while NLB is preferred for gRPC/HTTP2, WebSockets at L4, or when you need static IPs.
Ingress routing: Define Kubernetes Ingress resources (optionally grouped) to route based on hostnames (e.g., serviceA.example.com) or paths (e.g., /orders, /users) to per-service ClusterIP Services and Deployments; share a single ALB across many Ingresses using IngressGroup to reduce cost.
Why combine API Gateway + Ingress
One ALB or one API Gateway can expose many microservices via path- or host-based rules, reducing surface area and ops overhead; API Gateway adds edge features (WAF, auth, throttling) while Ingress keeps routing close to pods for fidelity and service discovery.
With Private Integration, keep the EKS ALB internal; API Gateway publishes public endpoints and reaches the cluster over VPC Link and private NLB, tightening exposure while retaining managed edge controls.
High-level architecture diagram (text)
Route 53 → API Gateway (REST/HTTP) with WAF →
Option A (Public): API Gateway HTTP proxy to public ALB (AWS Load Balancer Controller) → Kubernetes Ingress → Services (ClusterIP) → Pods (Java, Go, ML) → internal dependencies (DB, MQ)
Route 53 → API Gateway with WAF →
Option B (Private): API Gateway VPC Link → Private NLB → Internal ALB (Ingress) → Services/Pods; Frontend React can be on S3+CloudFront hitting API Gateway; stateful stores (RDS/Aurora, MSK/SQS, ElastiCache) are in private subnets.
Workload mapping
Frontend React: Host static site on S3 + CloudFront; API calls go to API Gateway domain; for internal-only UIs, target Ingress via private ALB/NLB and corporate network/VPN.
Java/Golang microservices: Expose via Kubernetes Ingress (ALB) with host/path rules; optional direct Service type LoadBalancer (NLB) for non-HTTP protocols; consolidate ALBs using IngressGroup to control costs.
DB: Use managed RDS/Aurora in private subnets; do not expose via Ingress; access via security groups/network policies from EKS only.
Message queue: Use managed MSK/Kafka, Amazon MQ, or SQS/SNS; for Kafka, consider NLB for brokers if client traffic needs L4; prefer private access from pods.
ML services: Containerize models (e.g., REST/gRPC) and expose through the same Ingress; for high-perf gRPC consider NLB or ALB HTTP/2; still front with API Gateway if needing unified auth/throttle.
Security: Attach AWS WAF at API Gateway and/or ALB; blog-guided patterns show WAF protection for EKS web apps and integration nuances.
Patterns to choose
Pattern 1: API Gateway → public ALB (Ingress)
Use when wanting edge governance plus simple public routing; API Gateway proxies to ALB that routes to services.
Pattern 2: API Gateway (Private Integration) → VPC Link → private NLB → internal ALB (Ingress)
Use when EKS must not be internet-exposed; API Gateway is the only public surface.
Pattern 3: Public ALB (Ingress) only
Use when edge features are not required or handled elsewhere; ALB terminates TLS and routes to services; can attach WAF to ALB.
Pattern 4: Public NLB → in-cluster gateway (e.g., Envoy/Gloo) or Ingress
Use for L4 performance or static IP needs; pair with separate WAF strategy since AWS WAF attaches to ALB/API Gateway, not NLB.
Concrete design for the stated stack
DNS: Route 53 zones for app and API domains; A/AAAA/ALIAS to API Gateway or ALB as chosen pattern.
Edge: API Gateway HTTP APIs with JWT authorizers, throttling, usage plans, staged deployments; AWS WAF attached at edge.
Networking: Public subnets for API Gateway edge and internet ALB if used; private subnets for EKS nodes, internal ALB/NLB, RDS/Aurora, and MQ.
EKS: Install AWS Load Balancer Controller; define Ingress per domain/path; use IngressGroup to share ALB; Services are ClusterIP behind Ingress.
Frontend: React on S3 + CloudFront; origin to S3; API calls to API Gateway custom domain.
Observability and protection: WAF rules, Shield Advanced as needed; API Gateway metrics and logs; ALB access logs; pod-level tracing.
Implementation tips
Single ALB for many services: Use host/path routing and IngressGroup to minimize ALB count and cost while keeping team autonomy on Ingress resources.
One edge, many backends: A single API Gateway can front many EKS services via a single domain with routes mapping to Ingress paths or via VPC Link to private NLB targets.
Controller choice: Use AWS Load Balancer Controller to ensure correct ALB (Ingress) vs NLB (Service type LoadBalancer) reconciliation on EKS.
ASCII architecture diagram
Public pattern (Gateway + public ALB + Ingress)
[Internet/WAF] → API Gateway (HTTP) → HTTPS → ALB (internet-facing, managed by AWS LB Controller) → Kubernetes Ingress (host/path) → Services (ClusterIP) → Pods (Java, Go, ML) → DB (RDS/Aurora, private) / MQ (MSK/SQS/SNS)
React (CloudFront → S3 static) → API calls → API Gateway domain
Private pattern (Gateway + VPC Link + private NLB + internal ALB + Ingress)
[Internet/WAF] → API Gateway (HTTP) → VPC Link → Private NLB → Internal ALB (Ingress) → Services (ClusterIP) → Pods → DB/MQ (private)
No direct public access to EKS; only API Gateway is exposed.
When to deviate
Pure ALB Ingress without API Gateway: Simpler and cheaper when edge policies are minimal; attach WAF to ALB if needed.
Microservices API gateway inside cluster (e.g., Envoy/Gloo/NGINX) with public ALB: Use when preferring rich in-cluster routing, but note WAF attachment tradeoffs if NLB is used; ALB supports WAF.
References: integration guides and best practices for API Gateway with EKS, ALB/NLB behavior with AWS Load Balancer Controller, cost/capability trade-offs, and WAF protection patterns.
```
