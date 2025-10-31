Upgrading **OpenShift** (OCP) isn’t done purely by a simple Python script — it typically involves using the `oc` CLI, and sometimes Red Hat’s automation tools like **Ansible** or **OpenShift Upgrade Operator**.

However, you *can* automate the process with Python by wrapping `oc` commands or using the **OpenShift REST API**. Below is an example of a **Python script** that:

* Checks the current OpenShift cluster version
* Lists available updates
* Applies an upgrade to a specified version

---

### 🧩 Example: Python Script to Upgrade OpenShift Cluster

```python
#!/usr/bin/env python3
import subprocess
import json
import sys

def run_cmd(cmd):
    """Run a shell command and return its output."""
    try:
        output = subprocess.check_output(cmd, shell=True, stderr=subprocess.STDOUT)
        return output.decode().strip()
    except subprocess.CalledProcessError as e:
        print(f"Command failed: {cmd}\n{e.output.decode()}")
        sys.exit(1)

def get_current_version():
    """Get current OpenShift cluster version."""
    print("Fetching current OpenShift version...")
    cmd = "oc get clusterversion version -o json"
    output = run_cmd(cmd)
    data = json.loads(output)
    return data["status"]["desired"]["version"]

def get_available_updates():
    """List available OpenShift updates."""
    print("Checking available updates...")
    cmd = "oc get clusterversion version -o json"
    output = run_cmd(cmd)
    data = json.loads(output)
    updates = data["status"].get("availableUpdates", [])
    return [u["version"] for u in updates]

def upgrade_cluster(target_version):
    """Apply an OpenShift cluster upgrade."""
    print(f"Upgrading OpenShift to version {target_version}...")
    cmd = f"oc adm upgrade --to={target_version} --allow-explicit-upgrade=true"
    run_cmd(cmd)
    print("Upgrade command applied successfully.")

if __name__ == "__main__":
    print("=== OpenShift Upgrade Script ===")

    current_version = get_current_version()
    print(f"Current version: {current_version}")

    updates = get_available_updates()
    if not updates:
        print("No available updates found.")
        sys.exit(0)

    print("Available updates:")
    for u in updates:
        print(f" - {u}")

    target = input("Enter target version to upgrade to: ").strip()
    if target not in updates:
        print("Invalid target version. Must be one of the available updates.")
        sys.exit(1)

    upgrade_cluster(target)
    print("Upgrade initiated. Monitor progress with:")
    print("  oc get clusterversion")
```

---

### ⚙️ Requirements

* Python 3.x
* Logged into OpenShift with admin privileges (`oc login`)
* Proper `kubeconfig` context set
* `oc` CLI installed

---

### 🧠 Notes

* This script **initiates** an upgrade — it does *not* monitor or validate completion. You can add that by polling `oc get clusterversion -o json` until `status.conditions` show `"Progressing": False` and `"Available": True`.
* For production upgrades, always **back up etcd**, **verify Operators’ compatibility**, and check **Red Hat’s upgrade matrix**.
* This script can be integrated into a CI/CD pipeline or an Ansible playbook.

---

Would you like me to extend this script to **automatically monitor** the upgrade until completion?
