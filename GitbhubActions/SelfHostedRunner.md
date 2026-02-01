**full steps to configure Helm, log in to GHCR, and install the Actions Runner Controller so you can deploy runners**. Let’s go **step by step**, clean and tested.

---

## **Step 1️⃣ — Create a GitHub Personal Access Token (PAT)**

1. Go to GitHub → **Settings → Developer settings → Personal access tokens → Tokens (classic) → Generate new token**
2. Give it **repo** and **admin:org** (or **workflow** + **repo**) permissions depending on whether it’s org-level or repo-level runner.
3. Copy the token — you’ll use it as your password.

---

## **Step 2️⃣ — Helm login to GHCR**

The chart is stored in **GitHub Container Registry (GHCR)**, which requires authentication:

```bash
helm registry login ghcr.io \
  -u <github-username> \
  -p <github-token>
```

* `<github-username>` → your GitHub username
* `<github-token>` → the PAT you just created

✅ You are now authenticated and can pull OCI Helm charts from GHCR.

> **To logout:**
>
> ```bash
> helm registry logout ghcr.io
> ```

---

## **Step 3️⃣ — Create a Kubernetes namespace**

```bash
kubectl create namespace arc-systems
```

* This is where the controller and runners will live.

---

## **Step 4️⃣ — Install the Actions Runner Controller via Helm**

```bash
helm install arc \
  --namespace arc-systems \
  --create-namespace \
  oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-controller
```

Then follow: https://github.com/Shehbab-Kakkar/github-actions-runners-k8s
