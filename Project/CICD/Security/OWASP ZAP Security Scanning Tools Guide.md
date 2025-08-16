# 🛡️ Security Scanning Tools Guide

This README provides step-by-step instructions for using popular open source security tools:

- [Trivy](#-trivy)
- [OWASP ZAP](#-owasp-zap)

---

## 🔍 Trivy

_Trivy_ is a universal open source vulnerability and misconfiguration scanner for containers, code, and cloud.

### Installation (Debian/Ubuntu)

```bash
sudo apt-get install wget apt-transport-https gnupg lsb-release
wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | gpg --dearmor | sudo tee /usr/share/keyrings/trivy.gpg > /dev/null
echo "deb [signed-by=/usr/share/keyrings/trivy.gpg] https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main" | sudo tee -a /etc/apt/sources.list.d/trivy.list
sudo apt-get update
sudo apt-get install trivy
```

### Usage Examples

- Scan a Docker image:
  ```bash
  trivy image imagename
  ```
- Scan a directory for vulnerabilities and misconfigurations:
  ```bash
  trivy fs --security-checks vuln,config Folder_name_OR_Path
  ```
- Scan for only HIGH and CRITICAL vulnerabilities:
  ```bash
  trivy image --severity HIGH,CRITICAL image_name
  ```
- Output scan results in JSON:
  ```bash
  trivy image -f json -o results.json image_name
  ```
- Scan a remote git repository:
  ```bash
  trivy repo repo-url
  ```
- Scan a running Kubernetes cluster:
  ```bash
  trivy k8s --report summary cluster
  ```

---

## 🕷️ OWASP ZAP (Zed Attack Proxy)

_OWASP ZAP_ is a free, open source security tool for finding web application vulnerabilities, especially useful for penetration testing and DAST (Dynamic Application Security Testing).

### Installation

#### On Debian/Ubuntu Linux

```bash
sudo apt update
sudo apt install snapd -y
sudo snap install zaproxy --classic
```
Or, download the [latest release from the official site](https://www.zaproxy.org/download/):

```bash
wget https://github.com/zaproxy/zaproxy/releases/download/v2.14.0/ZAP_2.14.0_Linux.tar.gz
tar -xzf ZAP_2.14.0_Linux.tar.gz
cd ZAP_2.14.0
./zap.sh
```

#### On macOS (using Homebrew)

```bash
brew install --cask owasp-zap
```

#### On Windows

Download and install from: [https://www.zaproxy.org/download/](https://www.zaproxy.org/download/)

---

### Usage Examples

#### 1. **Run the ZAP GUI**

```bash
zaproxy
```
or (if installed via tarball)
```bash
./zap.sh
```

#### 2. **Basic Spider and Active Scan (CLI/Headless)**

You can perform an automated scan against a target website:

```bash
zap-baseline.py -t https://example.com
```
Or, for a full active scan (may be more intrusive):

```bash
zap-full-scan.py -t https://example.com
```
> These scripts are located in the ZAP install directory (e.g., `ZAP_2.14.0/`).

#### 3. **REST API Example**

Start ZAP as a daemon (headless, with API):

```bash
./zap.sh -daemon -port 8090
```
Then interact via the [ZAP API](https://www.zaproxy.org/docs/api/).

#### 4. **Docker Usage**

```bash
docker run -u zap -p 8080:8080 -i owasp/zap2docker-stable zap.sh
```

---

### Example: Quick Headless Scan & Report

```bash
zap-baseline.py -t https://yourapp.local -r zap_report.html
```
This will run a passive scan and produce an HTML report.

---

## 📝 References

- [Trivy Documentation](https://aquasecurity.github.io/trivy/)
- [OWASP ZAP Documentation](https://www.zaproxy.org/docs/)
- [OWASP ZAP GitHub](https://github.com/zaproxy/zaproxy)

---

**Security Tip:**  
Integrate these tools into your CI/CD pipelines for continuous security testing and compliance.
