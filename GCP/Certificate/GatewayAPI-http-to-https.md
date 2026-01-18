In the Google Cloud Gateway API (available on GKE), HTTP-to-HTTPS redirection is achieved by defining a Gateway with two listeners (HTTP and HTTPS) and an HTTPRoute that uses a RequestRedirect filter to force the protocol change. 
1. Define the Gateway
The Gateway must have a listener on port 80 for HTTP and port 443 for HTTPS.
yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
spec:
  gatewayClassName: gke-l7-global-external-managed # Standard for 2026
  listeners:
  - name: http
    protocol: HTTP
    port: 80
  - name: https
    protocol: HTTPS
    port: 443
    tls:
      mode: Terminate
      certificateRefs:
      - name: my-ssl-cert # Reference to a Secret or Certificate Manager map
Use code with caution.

2. Create the Redirect Route
This HTTPRoute attaches specifically to the HTTP listener (sectionName: http) and redirects all traffic to the HTTPS scheme with a 301 status code. 
yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: http-to-https-redirect
spec:
  parentRefs:
  - name: my-gateway
    sectionName: http # Only targets the HTTP listener
  rules:
  - filters:
    - type: RequestRedirect
      requestRedirect:
        scheme: https
        statusCode: 301
Use code with caution.

3. Create the Main Application Route
You still need a second HTTPRoute (or additional rules) to handle the actual traffic once it reaches the HTTPS listener.
yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: app-route
spec:
  parentRefs:
  - name: my-gateway
    sectionName: https # Targets the secure listener
  rules:
  - backendRefs:
    - name: my-service
      port: 80
Use code with caution.

Key Differences from Ingress
Decoupled Logic: Unlike Ingress, which uses a single resource and a FrontendConfig annotation, the Gateway API uses separate listeners and routes to explicitly define the traffic flow.
Standardized: The RequestRedirect filter is a standard part of the Kubernetes Gateway API, making your configuration more portable across different cloud providers or implementations.
Status Codes: You can easily choose between 301 (Permanent), 302 (Found), 307, or 308 (preserving HTTP methods) directly in the manifest. 
