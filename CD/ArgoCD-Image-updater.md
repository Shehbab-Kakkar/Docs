# 🚀 Argo CD Image Updater with Kustomize & Helm

This README explains how to use [Argo CD Image Updater](https://argocd-image-updater.readthedocs.io/) to automatically update container image tags in your Kubernetes manifests. This works for both **Kustomize** and **Helm** applications managed by Argo CD.

---

## 📦 What is Argo CD Image Updater?

Argo CD Image Updater is a tool that automatically updates the image tags in your GitOps manifests (Kustomize or Helm values) based on new image versions found in your container registry. It then commits and pushes these changes to your Git repository, triggering Argo CD to deploy the new version.

---

## 🛠️ How It Works

1. **Build Phase**  
   - You build and push new images (e.g., via your CI/CD pipeline).
2. **Image Updater**  
   - Argo CD Image Updater scans your manifests (Kustomize or Helm values), checks for new versions, updates image tags, and commits the changes.
3. **Argo CD Sync**  
   - Argo CD detects the manifest changes and syncs your cluster with the updated images.

---

## 🚦 Example 1: Kustomize Application

**Kustomization.yaml:**
```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - deployment.yaml

images:
  - name: myrepo/myapp
    newTag: "1.0.0"      # <-- Will be updated by Argo CD Image Updater
```

**Argo CD Application Annotation:**
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-kustomize-app
  annotations:
    argocd-image-updater.argoproj.io/image-list: myrepo/myapp
    argocd-image-updater.argoproj.io/myrepo.myapp.update-strategy: latest
spec:
  source:
    repoURL: https://github.com/myorg/myrepo.git
    path: kustomize
    targetRevision: main
    kustomize:
      namePrefix: myapp-
```

**How it works:**
- Argo CD Image Updater checks your image registry for new tags of `myrepo/myapp`.
- When it finds a new tag (e.g., `1.0.1`), it updates `newTag: "1.0.1"` in `Kustomization.yaml` and commits the change to Git.
- Argo CD sees the commit and deploys the new image.

---

## 🚦 Example 2: Helm Application

**values.yaml:**
```yaml
image:
  repository: myrepo/myapp
  tag: "1.0.0"      # <-- Will be updated by Argo CD Image Updater
```

**Argo CD Application Annotation:**
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-helm-app
  annotations:
    argocd-image-updater.argoproj.io/image-list: myrepo/myapp
    argocd-image-updater.argoproj.io/myrepo.myapp.update-strategy: latest
    argocd-image-updater.argoproj.io/write-back-method: git
spec:
  source:
    repoURL: https://github.com/myorg/myrepo.git
    path: helm
    targetRevision: main
    helm:
      valueFiles:
        - values.yaml
```

**How it works:**
- Argo CD Image Updater finds the latest tag for `myrepo/myapp`.
- It updates `tag: "1.0.1"` in `values.yaml` and commits to your repo.
- Argo CD syncs and applies the new image.

---

## 📝 Key Annotations and Settings

- `argocd-image-updater.argoproj.io/image-list`: Comma-separated list of images to track.
- `argocd-image-updater.argoproj.io/<image>.update-strategy`: How to select new tags (e.g., latest, semver).
- `argocd-image-updater.argoproj.io/write-back-method`: How to update manifests (default is `git`).

See [full annotation docs](https://argocd-image-updater.readthedocs.io/en/stable/configuration/applications/#application-annotations).

---

## 🏗️ How to Install Argo CD Image Updater

### 1. **Prerequisites**
- A running Argo CD instance (v1.8+).
- kubectl access to your cluster.
- Git repository access credentials ready (for private repos).

### 2. **Install using Helm** (recommended)

**Add the Helm repo and install:**
```sh
helm repo add argo https://argoproj.github.io/argo-helm
helm repo update

helm install argocd-image-updater argo/argocd-image-updater \
  --namespace argocd \
  --set argocdImageUpdater.config.gitWriteBack=true
```
- This deploys the image updater in the `argocd` namespace.
- Adjust values as needed (see [Helm values](https://github.com/argoproj/argo-helm/tree/main/charts/argocd-image-updater)).

### 3. **Install using kubectl (YAML Manifest)**

**Apply the official manifest:**
```sh
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj-labs/argocd-image-updater/master/manifests/install.yaml
```

### 4. **Configure Access to Your Git Repository**
- For private repos, create a Kubernetes secret (see [docs](https://argocd-image-updater.readthedocs.io/en/stable/configuration/access-credentials/)).
- Example for SSH key:
  ```sh
  kubectl create secret generic git-creds \
    --from-file=sshPrivateKey=<path-to-private-key> \
    -n argocd
  ```

### 5. **Configure Registry Credentials (if needed)**
- For private registries, see [image credentials setup](https://argocd-image-updater.readthedocs.io/en/stable/configuration/image-registries/).

---

## 💡 Notes

- **Argo CD Image Updater** must be running with permissions to update your repo.
- For **Helm**, the updater will patch `values.yaml` or a custom values file.
- For **Kustomize**, it modifies the `images:` section in `Kustomization.yaml`.
- You can use advanced filters and strategies for image selection.

---

## 📚 References

- [Argo CD Image Updater Docs](https://argocd-image-updater.readthedocs.io/)
- [Argo CD Application CRD](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/)
- [Argo CD Image Updater Annotations](https://argocd-image-updater.readthedocs.io/en/stable/configuration/applications/#application-annotations)
- [Helm Chart for Argo CD Image Updater](https://github.com/argoproj/argo-helm/tree/main/charts/argocd-image-updater)

---
