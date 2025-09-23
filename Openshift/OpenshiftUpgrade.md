Upgrading **Red Hat OpenShift** to the latest version (e.g., from 4.x to 4.y) requires careful planning, pre-checks, and execution steps, especially in production environments. Below are **general steps** to guide you through upgrading OpenShift 4.x clusters to the latest supported version.

> 💡 These steps apply to OpenShift 4.x clusters managed via **OpenShift Container Platform (OCP)** using the **Cluster Version Operator (CVO)**.

---

## ✅ Prerequisites

1. **Check current OpenShift version**:

   ```bash
   oc get clusterversion
   ```

2. **Ensure cluster is healthy**:

   ```bash
   oc get clusteroperators
   oc get nodes
   ```

3. **Check upgrade path compatibility**:
   Use Red Hat’s [Upgrade Graph](https://access.redhat.com/labs/ocpupgradegraph/update_path) or `oc adm upgrade` to verify supported versions.

4. **Backup your cluster**:

   * Backup **ETCD** using `etcdctl` or the OpenShift documentation method.
   * Take note of custom resources, secrets, and configs.

5. **Confirm OCP subscription and entitlements**:

   * Ensure cluster is connected to the Red Hat pull secret registry.
   * Check entitlements with:

     ```bash
     oc get clusterversion
     ```

---

## 🔧 Upgrade Steps

### Step 1: Login as admin

```bash
oc login -u kubeadmin https://<api-server>
```

---

### Step 2: Verify available update versions

```bash
oc adm upgrade
```

This will list available versions the CVO can upgrade to.

You can also check available updates:

```bash
oc get clusterversion version -o json | jq '.status.availableUpdates'
```

---

### Step 3: Start the upgrade

```bash
oc adm upgrade --to=<version> --allow-explicit-upgrade
```

Example:

```bash
oc adm upgrade --to=4.14.10 --allow-explicit-upgrade
```

---

### Step 4: Monitor the upgrade

Check upgrade status:

```bash
oc get clusterversion
```

Watch ClusterOperators status:

```bash
watch oc get co
```

Check upgrade history:

```bash
oc get clusterversion -o json | jq '.status.history'
```

---

## 🚨 Important Notes

* OpenShift upgrades are **automated** and orchestrated by the **Cluster Version Operator (CVO)**.
* Control plane nodes are upgraded first, followed by worker nodes.
* Custom Operators installed via OperatorHub should be verified for version compatibility **before upgrade**.

---

## 🧪 Optional: Run a pre-upgrade check

If using Red Hat Advanced Cluster Management (ACM), or OpenShift Cluster Manager, you can run upgrade **preflight checks**.

CLI example (if available in your setup):

```bash
oc adm upgrade --to=<version> --dry-run
```

---

## 📚 Documentation Links

* [OpenShift Upgrade Guide (Official)](https://docs.openshift.com/container-platform/latest/updating/index.html)
* [OpenShift CLI `oc adm upgrade` Reference](https://docs.openshift.com/container-platform/latest/cli_reference/openshift_cli/administrator-cli-commands.html#oc-adm-upgrade)

---

Would you like me to tailor these steps for a **disconnected / air-gapped** cluster or **OpenShift Dedicated (OSD)**?

Great — since you said **“yes”**, I’ll now tailor the **OpenShift upgrade steps for a disconnected (air-gapped) cluster**.

Disconnected clusters require **extra preparation**, mainly around mirroring container images and configuring a local registry.

---

# 🛠️ OpenShift Upgrade: Disconnected (Air-Gapped) Cluster

This guide assumes you're using OpenShift 4.x and want to upgrade to the latest supported version **without internet access** on the cluster nodes.

---

## ✅ Prerequisites

### 1. **Workstation Requirements (Connected System)**

You need a system with internet access to:

* Pull and mirror container images.
* Download the necessary update payloads.
* Push content to your **disconnected registry**.

Install:

* `oc` CLI
* `podman`
* `oc-mirror` (for OCP 4.11+)
* `jq`

---

### 2. **Disconnected Registry (Local Mirror)**

Prepare a **local registry** that is accessible to your disconnected cluster. Example: `registry.example.com:5000`.

You must trust its certificate in the cluster and your nodes.

---

### 3. **Pull Secret**

Merge your **Red Hat pull secret** with credentials for your local registry:

```bash
oc registry login --registry=registry.redhat.io
podman login registry.example.com:5000
```

Use `jq` to merge pull secrets if needed:

```bash
jq -s '.[0] * .[1]' redhat-pull-secret.json local-registry-secret.json > merged-secret.json
```

---

## 🔄 Step-by-Step Upgrade Process (Disconnected)

---

### **Step 1: Mirror the Upgrade Images (on Connected Machine)**

Use **`oc-mirror`** (recommended for OCP ≥ 4.11):

```bash
oc mirror --config mirror-config.yaml docker://registry.example.com:5000
```

Example `mirror-config.yaml`:

```yaml
apiVersion: mirror.openshift.io/v1alpha2
kind: ImageSetConfiguration
archiveSize: 4
storageConfig:
  local:
    path: ./ocp-mirror
mirror:
  platform:
    channels:
      - name: stable-4.14
        minVersion: 4.13.0
        maxVersion: 4.14.10
        type: ocp
  operators:
    - catalog: registry.redhat.io/redhat/redhat-operator-index:v4.14
```

This command:

* Downloads OCP release images
* Downloads operator catalogs
* Pushes them to `registry.example.com:5000`

---

### **Step 2: Configure ImageContentSourcePolicy (ICSP)**

Once mirroring is complete, `oc-mirror` generates `ImageContentSourcePolicy` and `CatalogSource` YAMLs. Apply them on the disconnected cluster:

```bash
oc apply -f ./ocp-mirror/results-*/imageContentSourcePolicy.yaml
oc apply -f ./ocp-mirror/results-*/catalogSource.yaml
```

---

### **Step 3: Update Pull Secret in the Cluster**

Patch the cluster pull secret to include local registry auth:

```bash
oc set data secret/pull-secret -n openshift-config --from-file=.dockerconfigjson=merged-secret.json
```

---

### **Step 4: Start the Upgrade**

Manually point the CVO to the image:

```bash
oc adm upgrade --to-image=registry.example.com:5000/ocp/release@sha256:<digest> --allow-explicit-upgrade
```

You get the `<digest>` from the mirrored metadata or from:

```bash
oc adm release info --registry-config=<your-auth.json> registry.example.com:5000/ocp/release:4.14.10
```

---

### **Step 5: Monitor the Upgrade**

```bash
oc get clusterversion
watch oc get co
```

The Cluster Version Operator (CVO) will now use your mirrored images to upgrade the cluster components.

---

## 🧪 Optional: Verify Upgrade Path Locally

```bash
oc adm release info --commits registry.example.com:5000/ocp/release:4.14.10
```

---

## 🔐 Security Notes

* Ensure that your local registry uses a valid TLS certificate.
* Add the CA to:

  * `/etc/pki/ca-trust/source/anchors/` on RHCOS nodes
  * `/etc/docker/certs.d/` for container tools

Update CA trust:

```bash
update-ca-trust extract
```

---

## 📘 Official References

* [OpenShift Disconnected Install & Upgrade](https://docs.openshift.com/container-platform/latest/updating/updating-restricted-network-cluster/updating-restricted-network-cluster.html)
* [oc-mirror tool](https://docs.openshift.com/container-platform/latest/cli_tools/oc-mirror.html)

---

Would you like a sample `mirror-config.yaml` file and the exact `oc-mirror` commands for a specific version (e.g., 4.14.10)?

Great — here’s a full **working example** to upgrade a **disconnected OpenShift cluster to version 4.14.10** using the `oc-mirror` tool.

---

# 🧾 Sample: Upgrade Disconnected OpenShift Cluster to 4.14.10

This example assumes:

* Your **local registry** is: `registry.example.com:5000`
* You want to mirror **OCP 4.14.10** and **Red Hat Operator Catalog**
* You’re using **`oc-mirror` (recommended)**

---

## 🔧 1. Create `mirror-config.yaml`

Save the following as `mirror-config.yaml` on a **connected machine**:

```yaml
apiVersion: mirror.openshift.io/v1alpha2
kind: ImageSetConfiguration
archiveSize: 4
storageConfig:
  local:
    path: ./ocp-mirror
mirror:
  platform:
    channels:
      - name: stable-4.14
        type: ocp
        minVersion: 4.14.10
        maxVersion: 4.14.10
  operators:
    - catalog: registry.redhat.io/redhat/redhat-operator-index:v4.14
```

---

## 🧰 2. Run `oc-mirror` to Mirror Content

> Make sure you're authenticated to both `registry.redhat.io` and your `registry.example.com:5000`.

```bash
oc mirror --config=mirror-config.yaml docker://registry.example.com:5000
```

This will:

* Download OCP 4.14.10 release images
* Mirror operator catalog images
* Push to your local registry
* Output manifests to `ocp-mirror/results-*`

---

## 🗂 3. Apply Cluster Manifests (on disconnected cluster)

Transfer the `results-*` folder to your disconnected cluster (e.g., via USB or SCP) and run:

```bash
oc apply -f ./results-*/imageContentSourcePolicy.yaml
oc apply -f ./results-*/catalogSource.yaml
```

This will tell OpenShift to use your local registry instead of Red Hat's.

---

## 🔑 4. Patch Pull Secret

Merge Red Hat & local registry credentials (if not already):

```bash
jq -s '.[0] * .[1]' redhat-pull-secret.json local-registry-secret.json > merged-secret.json
```

Patch into OpenShift:

```bash
oc set data secret/pull-secret -n openshift-config --from-file=.dockerconfigjson=merged-secret.json
```

---

## 🚀 5. Start the Upgrade

First, extract the digest of the mirrored release:

```bash
oc adm release info registry.example.com:5000/ocp/release:4.14.10 --output=json | jq '.metadata.metadata'
```

Copy the value under `"digest"` (e.g., `sha256:abcd1234...`).

Then upgrade:

```bash
oc adm upgrade --to-image=registry.example.com:5000/ocp/release@sha256:<digest> --allow-explicit-upgrade
```

Example:

```bash
oc adm upgrade --to-image=registry.example.com:5000/ocp/release@sha256:abc123... --allow-explicit-upgrade
```

---

## 📈 6. Monitor the Upgrade

```bash
watch oc get clusterversion
watch oc get co
```

Once `clusterversion` shows `4.14.10` and all `ClusterOperators` are `AVAILABLE`, the upgrade is complete.

---

## 📌 Notes

* You **must trust your local registry’s TLS certs** on all cluster nodes.
* Store mirrored images for future use if you're managing multiple clusters.
* You can also use `oc-mirror` with `--to-mirror` for portable archives (air-gap transfers).

---
