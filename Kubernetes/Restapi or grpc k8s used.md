**Short answer:**
Kubernetes primarily uses a **REST API (HTTP/JSON)** for its public, user-facing API, while **gRPC is used internally** between certain components.

---

## How Kubernetes uses REST vs gRPC

### ✅ REST API (Primary & Public)

Kubernetes exposes a **RESTful API** via the **API Server**.

**Used for:**

* `kubectl`
* Client libraries (Go, Python, Java, etc.)
* Controllers & operators
* CI/CD tools
* Dashboard and external integrations

**Characteristics:**

* HTTP/HTTPS + JSON (sometimes Protobuf for efficiency)
* Resource-oriented (`/api/v1/pods`, `/apis/apps/v1/deployments`)
* Declarative (CRUD on resources)
* Stable and versioned

👉 If you interact with Kubernetes as a platform, **you are using REST**.

---

### ⚙️ gRPC (Internal & Component-Level)

Kubernetes uses **gRPC internally** where performance and strict contracts matter.

**Key examples:**

* **CRI (Container Runtime Interface)**

  * `kubelet ↔ containerd / CRI-O`
  * **gRPC**
* **CSI (Container Storage Interface)**

  * Kubernetes ↔ storage providers
  * **gRPC**
* **CNI (Container Networking Interface)**

  * Mostly exec-based, but modern implementations may use gRPC-like APIs internally

**Why gRPC here?**

* High performance
* Strongly typed APIs
* Efficient binary protocol (Protobuf)
* Clear contracts between components

---

## Summary Table

| Aspect         | REST API                  | gRPC                |
| -------------- | ------------------------- | ------------------- |
| Used by        | Users, kubectl, operators | Internal components |
| Protocol       | HTTP + JSON               | HTTP/2 + Protobuf   |
| Public API     | ✅ Yes                     | ❌ No                |
| Performance    | Good                      | Excellent           |
| Human-readable | ✅ Yes                     | ❌ No                |

---

## Which should *you* use?

* **Building tools, operators, or integrations?** → **REST API**
* **Extending Kubernetes internals (runtime, storage, networking)?** → **gRPC**

If you want, I can:

* Show a **kubectl REST call example**
* Explain **why Kubernetes didn’t choose gRPC for its public API**
* Compare **REST vs gRPC for cloud-native systems**
