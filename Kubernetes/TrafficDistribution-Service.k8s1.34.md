Here’s how you can configure Kubernetes **v1.34** to get behavior like *PreferSameNode* and *PreferSameZone*, by using the newer `trafficDistribution` field on Services. I’ll include a sample YAML and also notes on what needs to be enabled / checked.

---

## What’s new in v1.34

* Kubernetes 1.34 has graduated the `PreferSameTrafficDistribution` feature gate: the values `PreferSameZone` and `PreferSameNode` in `spec.trafficDistribution` are **beta and enabled by default**. ([Kubernetes][1])
* `PreferSameZone` is effectively what used to be called `PreferClose`. ([Kubernetes][2])

---

## Requirements / Preconditions

To use this successfully:

1. Kubernetes cluster must be version **1.34** (or ≥1.33 with the feature gate).
2. The feature gate `PreferSameTrafficDistribution` must be enabled if you are prior to v1.34. From v1.34, it’s on by default. ([Kubernetes][2])
3. EndpointSlice / kube-proxy must support using trafficDistribution / topology aware hints.
4. Nodes must be labeled with standard topology labels: `topology.kubernetes.io/zone`, etc.
5. If you use `internalTrafficPolicy` = `Local` or `externalTrafficPolicy = Local`, those may interact / disable some behavior of traffic distribution. (I.e., trafficPolicy local pulls only “local” endpoints.)

---

## Sample YAML configurations

Here are two example Service definitions: one using **PreferSameNode**, another **PreferSameZone**.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: svc-prefer-same-node
  namespace: my-app
spec:
  selector:
    app: my-app
  ports:
    - port: 80
      targetPort: 8080
  type: ClusterIP
  # This is the new field in v1.34
  trafficDistribution: PreferSameNode
  # (optional) internalTrafficPolicy or external may affect behavior
  # internalTrafficPolicy: Cluster
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: svc-prefer-same-zone
  namespace: my-app
spec:
  selector:
    app: my-app
  ports:
    - port: 80
      targetPort: 8080
  type: ClusterIP
  trafficDistribution: PreferSameZone
  # internalTrafficPolicy: Cluster
```

---

## What these do

* `PreferSameNode`: tries to route traffic from a client to a pod on the **same node** if possible. If that is not possible (no endpoint on same node), it falls back to other nodes (possibly same zone, or cross-zone depending on availability).
* `PreferSameZone`: tries to route traffic to pods in the **same zone** of the client first; if none available, then across zones.

---

## Caveats / Behavior Details

* `trafficDistribution` only *prefers*; it does *not guarantee* that traffic stays in same node/zone, when no endpoints meet the preference.
* If the number of endpoints is small, or endpoints are unevenly spread, Kubernetes may fall back to cross-zone or cross-node endpoints.
* `trafficDistribution` interacts with topology-aware hints (EndpointSlice) and kube-proxy. Hints are included in EndpointSlices (in `hints` field) which help kube-proxy route accordingly. ([kubespec.dev][3])
* If you set `internalTrafficPolicy: Local`, traffic must go to endpoints on the same node; but this may restrict availability.

---

If you want, I can give you a *full deployment + service + testing setup* to verify the behavior in your cluster. Do you want that?

[1]: https://kubernetes.io/blog/2025/07/28/kubernetes-v1-34-sneak-peek/?utm_source=chatgpt.com "Kubernetes v1.34 Sneak Peek | Kubernetes"
[2]: https://kubernetes.io/blog/2025/08/27/kubernetes-v1-34-release/?utm_source=chatgpt.com "Kubernetes v1.34: Of Wind & Will (O' WaW) | Kubernetes"
[3]: https://kubespec.dev/kubernetes/discovery.k8s.io/v1/EndpointSlice?utm_source=chatgpt.com "Kubernetes v1.34 Spec: EndpointSlice"
