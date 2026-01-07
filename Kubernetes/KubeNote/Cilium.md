To use **Cilium eBPF** in Kubernetes, you need to deploy **Cilium** as the CNI (Container Network Interface) plugin. Cilium leverages eBPF to provide networking, security, and observability for your Kubernetes cluster.

Below are the **key steps to get started**:

---

## 1. **Install Prerequisites**
- **Kubernetes cluster** (on-prem or cloud; most environments supported)
- **Kernel 4.9+** (recommended: 5.x+ for best eBPF support)
- **kubectl** access to your cluster

---

## 2. **Deploy Cilium with eBPF**

### a. **Standard Quick Install (Recommended for New Users)**
Run this command to install Cilium with eBPF datapath:

```bash
kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.15/install/kubernetes/quick-install.yaml
```
(*Check the [latest install URL](https://docs.cilium.io/en/stable/gettingstarted/k8s-install-standard/)*)

### b. **Install with CLI Tool (Advanced, Customizable)**
Install the Cilium CLI:
```bash
curl -L --remote-name-all https://github.com/cilium/cilium-cli/releases/latest/download/cilium-linux-amd64.tar.gz{,.sha256sum}
tar xzvf cilium-linux-amd64.tar.gz
sudo mv cilium /usr/local/bin
```
Then deploy Cilium:
```bash
cilium install
```

---

## 3. **Verify the Installation**
Check that Cilium Pods are running:
```bash
kubectl -n kube-system get pods -l k8s-app=cilium
```
Validate connectivity using Cilium CLI:
```bash
cilium status
```

---

## 4. **(Optional) Enable Advanced eBPF Features**
Edit the CiliumConfig (or DaemonSet) and enable features like:
- eBPF-based kube-proxy replacement
- eBPF-based load balancing or network policies

Example (using Helm values):
```yaml
kubeProxyReplacement: strict
enableCNPNodeStatus: true
```
Install/upgrade using Helm (more customizable):
```bash
helm repo add cilium https://helm.cilium.io/
helm install cilium cilium/cilium --version <version> \
  --namespace kube-system \
  --set kubeProxyReplacement=strict
```

---

## 5. **Use Cilium Features**
- **Network policies** with eBPF acceleration.
- **Visibility** (Hubble observability, also built on eBPF):
  ```bash
  cilium hubble enable
  cilium hubble port-forward & # For UI
  ```
- **Service load balancing, DNS-aware security, etc.**

---

## 6. **Documentation & Resources**
- [Official Cilium Getting Started Guide](https://docs.cilium.io/en/stable/gettingstarted/k8s-install-default/)
- [Cilium CLI documentation](https://docs.cilium.io/en/stable/gettingstarted/cilium-cli/)
- [Cilium eBPF Deep Dive](https://docs.cilium.io/en/stable/concepts/ebpf/)

---

### **Summary Table**

| Step       | Description                                   | Key Command        |
|------------|-----------------------------------------------|--------------------|
| Install    | Deploy Cilium manifest or use Cilium CLI      | `kubectl create -f ...` or `cilium install` |
| Verify     | Check pods/status                             | `kubectl get pods`, `cilium status` |
| Customize  | Enable advanced (eBPF) features (Helm/YAML)   | (Helm install or edit YAML) |
| Observe    | Use Hubble                                    | `cilium hubble enable`     |

---
