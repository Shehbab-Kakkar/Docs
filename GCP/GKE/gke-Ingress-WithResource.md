gke internal Ingress example with the resource not service

Kubernetes Ingress can be configured to use a backend that is a resource (an ObjectRef to another Kubernetes resource) instead of a service. This approach is less common but supported in Kubernetes API v1.23 and later, allowing an Ingress to direct traffic to a resource like a custom object or storage bucket rather than just a Service.

### Example of GKE Internal Ingress using resource backend (not service)

A Kubernetes Ingress using a resource backend looks like this example:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-resource-backend
  annotations:
    kubernetes.io/ingress.class: "gce-internal"  # specifies internal GKE load balancer
spec:
  defaultBackend:
    resource:
      apiGroup: k8s.example.com
      kind: StorageBucket
      name: static-assets
  rules:
  - http:
      paths:
      - path: /icons
        pathType: ImplementationSpecific
        backend:
          resource:
            apiGroup: k8s.example.com
            kind: StorageBucket
            name: icon-assets
```

- This example defines an internal GKE ingress (using `kubernetes.io/ingress.class: "gce-internal"` annotation).
- The backend is not a service, but rather a resource identified by `apiGroup`, `kind`, and `name`.
- Such resource backends are mutually exclusive with service backends.
- This is useful when ingressing data to custom storage backends or other Kubernetes API resources rather than Services.

### Important notes for GKE usage:
- GKE internal ingress requires annotation `kubernetes.io/ingress.class: "gce-internal"` for internal load balancer.
- Usually, internal ingress routes traffic to Kubernetes Services, but resource backends can be configured if the backend is a valid namespaced Kubernetes resource.
- Proper setup of ingress class and backend resources is necessary for GKE load balancer provisioning.
- Health checks and other GKE-specific load balancer config happen per backend service normally.
- Resource backend support is more Kubernetes generic; GKE specifics for these might require custom controller support.

This example and details are derived from official Kubernetes docs and GKE internal ingress guidance [1][2][3].

If a specific resource type other than `Service` is desired as backend in GKE Ingress, also ensure your cluster and ingress controller support it or have a custom controller handling it. 

Would you like me to provide a ready-to-use internal GKE ingress manifest example with resource backend tailored to a specific use case?

Citations:
[1] Ingress https://kubernetes.io/docs/concepts/services-networking/ingress/
[2] GKE Ingress for Application Load Balancers https://cloud.google.com/kubernetes-engine/docs/concepts/ingress
[3] gce-internal ingress fails with "Missing one or mo... https://www.googlecloudcommunity.com/gc/Google-Kubernetes-Engine-GKE/gce-internal-ingress-fails-with-quot-Missing-one-or-more/m-p/694885/highlight/true
[4] Configuring Ingress for internal Application Load Balancers https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balance-ingress
[5] Google Kubernetes Engine Networking - Ingress https://partner.cloudskillsboost.google/paths/101/course_templates/34/video/559717
[6] How to Setup Ingress on GKE using GKE Ingress Controller https://devopscube.com/setup-ingress-gke-ingress-controller/
[7] Set up an external Application Load Balancer with Ingress https://cloud.google.com/kubernetes-engine/docs/tutorials/http-balancer
[8] GKE Custom Ingress Routing Made Simple https://developer.cyberark.com/blog/gke-custom-ingress-routing-made-simple/
[9] Kubernetes Ingress with NGINX Ingress Controller Example https://spacelift.io/blog/kubernetes-ingress
[10] Deploying Public & Private Nginx Ingress Controllers with ... https://blog.thecloudside.com/deploying-public-private-nginx-ingress-controllers-with-http-s-loadbalancer-in-gke-dcf894197fb7
[11] Emissary-ingress with GKE https://www.getambassador.io/docs/emissary/latest/topics/running/ambassador-with-gke
[12] Exam Professional Cloud Developer topic 1 question 70 ... https://www.examtopics.com/discussions/google/view/40504-exam-professional-cloud-developer-topic-1-question-70/
[13] Creating a Kubernetes Ingress resource for GCP/GKE by ... https://stackoverflow.com/questions/69018880/creating-a-kubernetes-ingress-resource-for-gcp-gke-by-example
[14] Google Kubernetes Engine Ingress configuration for ... https://engineering.sada.com/google-kubernetes-engine-ingress-configuration-for-production-environments-9677a303166a
[15] Could an Ingress be attached to a Loadbalancer type ... https://stackoverflow.com/questions/65413929/could-an-ingress-be-attached-to-a-loadbalancer-type-service-in-gke
[16] Google Kubernetes Engine Networking - Ingress https://partner.cloudskillsboost.google/paths/89/course_templates/34/video/559717?locale=id
