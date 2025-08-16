# 🔐 Snyk Security Scanning – Complete Guide

This README provides usage instructions for securing your project using **Snyk**, covering:

- 🐳 Docker image scanning
- ☁️ IaC scanning (Terraform)
- ☕ Java application scanning (Maven/Gradle)

Snyk helps you find and fix vulnerabilities in:

- Application code
- Open-source dependencies
- Container images
- Infrastructure as Code (IaC)

---

## 📦 1. Install Snyk CLI

Install via npm:

```bash
npm install -g snyk
```

Or with Homebrew (macOS):

```bash
brew tap snyk/tap && brew install snyk
```

---

## 🔑 2. Authenticate Snyk CLI

Before running scans, authenticate with your Snyk account:

```bash
snyk auth
```

This opens a browser window to authorize access.

---

## 🐳 3. Scan Docker Images

Snyk can scan Docker images for vulnerabilities in OS packages and application dependencies.

### 🔍 Scan a Docker Image

```bash
snyk container test <image-name> --file=<Dockerfile>
```

**Example:**

```bash
snyk container test my-app:latest --file=Dockerfile
```

### 📤 Monitor Docker Image (Optional)

```bash
snyk container monitor my-app:latest --file=Dockerfile
```

This uploads a snapshot to Snyk for ongoing monitoring.

---

## ☁️ 4. Scan Infrastructure as Code (Terraform)

Snyk can detect security misconfigurations in your Terraform (`.tf`) files.

### 🔍 Scan Terraform Configs

```bash
snyk iac test <path-to-tf-files>
```

**Example:**

```bash
snyk iac test infrastructure/
```

### 📤 Monitor IaC (Optional)

```bash
snyk iac monitor infrastructure/
```

---

## ☕ 5. Scan Java Projects (Maven / Gradle)

Snyk can detect vulnerabilities in dependencies declared in Java project manifests.

### 🔍 Maven Project

```bash
snyk test --file=pom.xml
```

### 🔍 Gradle Project

```bash
snyk test --file=build.gradle
```

### 📤 Monitor Java Project (Optional)

```bash
snyk monitor --file=pom.xml
```

---

## ⚙️ 6. Advanced Options

| Option                      | Description                                         |
|-----------------------------|-----------------------------------------------------|
| `--severity-threshold=high` | Show only high-severity issues                      |
| `--all-projects`            | Scan all supported manifests in the folder          |
| `--json`                    | Output results in JSON format (for CI pipelines)    |
| `--org=<org-name>`          | Scan under a specific Snyk organization             |
| `--sarif`                   | Output in SARIF format (for GitHub code scanning)   |

---

## 🤖 CI/CD Integration (GitHub Actions Example)

```yaml
name: Snyk Security Scan

on: [push]

jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Install Snyk CLI
        run: npm install -g snyk

      - name: Authenticate Snyk
        run: snyk auth ${{ secrets.SNYK_TOKEN }}

      - name: Scan Java
        run: snyk test --file=pom.xml

      - name: Scan Docker Image
        run: |
          docker build -t my-app .
          snyk container test my-app --file=Dockerfile

      - name: Scan Terraform
        run: snyk iac test infrastructure/
```

---

## 📘 References

- [Snyk CLI Documentation](https://docs.snyk.io/snyk-cli)
- [Snyk Container](https://snyk.io/product/container-vulnerability-management/)
- [Snyk Infrastructure as Code](https://snyk.io/product/infrastructure-as-code/)
- [Snyk for Java Projects](https://snyk.io/docs/snyk-for-java/)
- [Snyk Best Practices](https://snyk.io/blog/best-practices/)

---

## 🛡️ Best Practices

- Integrate Snyk into your CI/CD pipeline for continuous monitoring.
- Use `snyk monitor` to track project health in the Snyk dashboard.
- Regularly update dependencies and rebuild Docker images to reduce risk.
- Use `--all-projects` for monorepos or polyglot codebases.

---

## 🧪 Example Project Structure

```
.
├── Dockerfile
├── pom.xml              # Java Maven project
├── infrastructure/      # Terraform IaC code
│   └── main.tf
```

---

## 🧑‍💻 Author

Maintained by your DevSecOps team. Built with ❤️ using Snyk.
