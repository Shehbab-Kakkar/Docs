Give me Example of readiness proble, livenessproble, startprobe in gke

Here are clear YAML examples for readiness, liveness, and startup probes commonly used in Google Kubernetes Engine (GKE). These can be adapted directly for container health checks in GKE Deployments.

### Readiness Probe Example

This probe checks if an application is ready to receive traffic by hitting an HTTP endpoint:

```yaml
readinessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 2
  failureThreshold: 3
```
This probe periodically checks whether `/healthz` on port 8080 responds, marking the pod as ready for service routing if it does[1][2].

### Liveness Probe Example

This probe detects if the application is stuck or not operating correctly, triggering a restart if failures persist:

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 15
  timeoutSeconds: 2
  failureThreshold: 3
```
If the liveness probe fails (no response or error), Kubernetes restarts the container automatically[2][3].

### Startup Probe Example

This probe checks that the application has started successfully before the other probes begin. Ideal for slow-starting apps:

```yaml
startupProbe:
  httpGet:
    path: /start
    port: 8080
  initialDelaySeconds: 30
  periodSeconds: 10
  failureThreshold: 15
```
This setup lets a container take up to 2.5 minutes to fully start before Kubernetes considers it unhealthy[4][5].

### Example Combined in a Pod Spec

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: demo-app
spec:
  containers:
  - name: demo-container
    image: myapp:latest
    ports:
    - containerPort: 8080
    readinessProbe:
      httpGet:
        path: /healthz
        port: 8080
      initialDelaySeconds: 5
      periodSeconds: 10
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
      initialDelaySeconds: 10
      periodSeconds: 15
    startupProbe:
      httpGet:
        path: /start
        port: 8080
      initialDelaySeconds: 30
      periodSeconds: 10
      failureThreshold: 15
```

These patterns work out of the box on GKE for robust application health management[6][1][4].Here are concise examples of readiness, liveness, and startup probes as typically used in Kubernetes on Google Kubernetes Engine (GKE):

### Readiness Probe Example

```yaml
readinessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 2
  failureThreshold: 3
```
This will periodically check the `/healthz` HTTP endpoint. If healthy, the pod is added to service endpoints[1][2].

### Liveness Probe Example

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 15
  timeoutSeconds: 2
  failureThreshold: 3
```
Liveness probes confirm the app is running. Failures trigger container restarts[2][3].

### Startup Probe Example

```yaml
startupProbe:
  httpGet:
    path: /start
    port: 8080
  initialDelaySeconds: 30
  periodSeconds: 10
  failureThreshold: 15
```
Startup probes are useful for apps with slow initialization, allowing up to 2.5 minutes to start[4][5].

You can include all three in a pod spec for comprehensive health checking in GKE[6][1][4].

Citations:
[1] Kubernetes Readiness Probes: Guide & Examples https://www.groundcover.com/blog/kubernetes-readiness-probe
[2] Readiness vs liveliness probes: How to set them up and ... https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-setting-up-health-checks-with-readiness-and-liveness-probes
[3] Guide to Kubernetes Liveness Probes with Examples https://spacelift.io/blog/kubernetes-liveness-probe
[4] The Only Blog You'll Ever Need to Master Kubernetes ... https://www.kubeblogs.com/kubernetes-probes-liveness-readiness-and-startup-probes/
[5] Kubernetes Startup Probes: Examples & Common Pitfalls https://www.vcluster.com/blog/kubernetes-startup-probes-examples-and-common-pitfalls
[6] Configure Liveness, Readiness and Startup Probes https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/
[7] Configure Kubernetes readiness probes (Optional) https://cloud.google.com/service-mesh/docs/gateway/configure-readiness-probes
[8] Liveness, Readiness, and Startup Probes https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/
[9] Configuring Liveness and Readiness Probes https://www.cloudskillsboost.google/focuses/19114?parent=catalog
[10] Kubernetes Liveness Probes: Configuration & Best Practices https://www.groundcover.com/blog/kubernetes-liveness-probe
[11] How to enable startup probe on GKE 1.16? https://stackoverflow.com/questions/61918062/how-to-enable-startup-probe-on-gke-1-16
[12] Kubernetes Liveness Probes: A Complete Guide https://www.qovery.com/blog/kubernetes-liveness-probes-a-complete-guide
[13] Guide to Understanding Your Kubernetes Liveness Probes ... https://www.fairwinds.com/blog/a-guide-to-understanding-kubernetes-liveness-probes-best-practices
[14] Configure container health checks for services | Cloud Run https://cloud.google.com/run/docs/configuring/healthchecks
[15] Kubernetes Health Check | LivenessProbe - K21 Academy https://k21academy.com/docker-kubernetes/kubernetes-readiness-and-livenessprobe/
[16] Configure exec probe timeouts before upgrading to GKE ... https://cloud.google.com/kubernetes-engine/docs/deprecations/exec-probe-timeouts
[17] Google Kubernetes Engine GKE with DevOps 75 Real- ... https://github.com/stacksimplify/google-kubernetes-engine
