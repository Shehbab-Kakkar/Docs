For **Google-managed certificates** (the SSL/TLS certificates that Google Cloud provisions and renews for you), there is a documented limit on how many **non-wildcard domains** (exact hostnames) a single certificate can cover:

🔹 A single Google-managed certificate supports **up to 100 non-wildcard domains** (no wildcard `*` entries). ([Google Cloud Documentation][1])

✔️ *Non-wildcard domains* means exact hostnames like `example.com`, `www.example.com`, `app.example.org`, etc.

❗ Wildcard domains (e.g., `*.example.com`) are **not supported** in the traditional GKE ManagedCertificate setup. ([Google Cloud Documentation][1])

So the answer is:

✅ **Up to 100 non-wildcard domains per Google-managed certificate.** ([Google Cloud Documentation][1])

If you need to cover more than 100 hostnames, you’d generally need:

* Multiple managed certificates and attach them to your load balancer/Ingress (subject to other LB limits), **or**
* Use Google Cloud Certificate Manager directly with DNS authorization, which also has limits (e.g., up to 100 domains per certificate when using DNS authorization). ([Google Cloud Documentation][2])

Let me know if you want specifics on combining multiple certificates or limits with load balancers!

[1]: https://docs.cloud.google.com/kubernetes-engine/docs/how-to/managed-certs?utm_source=chatgpt.com "Using Google-managed SSL certificates  |  GKE networking  |  Google Cloud Documentation"
[2]: https://docs.cloud.google.com/certificate-manager/docs/quotas?utm_source=chatgpt.com "Quotas and limits  |  Certificate Manager  |  Google Cloud Documentation"
