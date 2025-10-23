Upgrading a **Google Kubernetes Engine (GKE)** cluster using **Python** can be done with the **Google Cloud Python client libraries**, specifically the **Container API** (`google-cloud-container`).

Here’s a step-by-step guide and example script 👇

---

### ✅ **Prerequisites**

1. **Install required packages**

   ```bash
   pip install google-cloud-container google-auth
   ```

2. **Authenticate with Google Cloud**

   ```bash
   gcloud auth application-default login
   ```

   This sets up credentials used by the Python client.

3. **Enable the GKE API**

   ```bash
   gcloud services enable container.googleapis.com
   ```

---

### 🧠 **Script Overview**

The process:

1. Use `ClusterManagerClient` from `google.cloud.container_v1`.
2. Call the `update_cluster` method with the new version or configuration.
3. Optionally, wait for the operation to complete.

---

### 🐍 **Python Script Example**

```python
from google.cloud import container_v1
from google.api_core import operation
import google.auth

def upgrade_gke_cluster(project_id, location, cluster_id, new_version):
    """
    Upgrades a GKE cluster to a new version.

    Args:
        project_id (str): Your Google Cloud project ID.
        location (str): The cluster location (zone or region, e.g., "us-central1-c").
        cluster_id (str): The name of the GKE cluster.
        new_version (str): The target Kubernetes version (e.g., "1.29.4-gke.100").
    """

    # Create a client
    client = container_v1.ClusterManagerClient()

    # Build the upgrade request
    cluster_update = container_v1.ClusterUpdate(desired_master_version=new_version)

    request = container_v1.UpdateClusterRequest(
        project_id=project_id,
        zone=location,
        cluster_id=cluster_id,
        update=cluster_update,
    )

    print(f"Starting upgrade of cluster '{cluster_id}' to version {new_version}...")

    # Trigger the upgrade
    operation = client.update_cluster(request=request)

    # Wait for completion
    print("Waiting for operation to complete...")
    operation.result()  # Blocks until done

    print(f"✅ Cluster '{cluster_id}' successfully upgraded to {new_version}.")

if __name__ == "__main__":
    upgrade_gke_cluster(
        project_id="my-gcp-project",
        location="us-central1-c",
        cluster_id="my-gke-cluster",
        new_version="1.29.4-gke.100"
    )
```

---

### ⚙️ **Optional: Upgrade Node Pools**

To upgrade node pools as well:

```python
def upgrade_node_pool(project_id, location, cluster_id, node_pool_id, new_version):
    client = container_v1.ClusterManagerClient()

    request = container_v1.UpdateNodePoolRequest(
        project_id=project_id,
        zone=location,
        cluster_id=cluster_id,
        node_pool_id=node_pool_id,
        node_version=new_version,
    )

    print(f"Upgrading node pool '{node_pool_id}' to {new_version}...")
    operation = client.update_node_pool(request=request)
    operation.result()
    print(f"✅ Node pool '{node_pool_id}' upgraded.")
```

---

### 🧩 **Notes**

* If your cluster is **regional**, use `location` instead of `zone`.
  Example:

  ```python
  parent = f"projects/{project_id}/locations/{location}/clusters/{cluster_id}"
  ```

  And call:

  ```python
  client.update_cluster(name=parent, update=cluster_update)
  ```
* You can check available versions with:

  ```bash
  gcloud container get-server-config --zone=us-central1-c
  ```

---

Would you like me to show you how to **upgrade both the control plane and all node pools automatically** in one script (with version compatibility checks)?
