Below is a **complete, working example** of using **External Secrets Operator (ESO)** with **AWS Secrets Manager** to inject **DB credentials** into Kubernetes and consume them in a **Golang application**.

I’ll assume:

* Kubernetes cluster (EKS or any K8s with AWS access)
* AWS Secrets Manager holds DB credentials
* ESO is already installed (I’ll still show the IAM + setup pieces)
* App uses environment variables

---

## 1. AWS Secrets Manager – Create the Secret

Example secret name: `prod/db/postgres`

**Value (JSON):**

```json
{
  "username": "dbuser",
  "password": "supersecretpassword",
  "host": "postgres.prod.svc.cluster.local",
  "port": "5432",
  "dbname": "appdb"
}
```

---

## 2. IAM Policy for ESO

ESO needs permission to read secrets.

### IAM Policy

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "secretsmanager:GetSecretValue",
        "secretsmanager:DescribeSecret"
      ],
      "Resource": "arn:aws:secretsmanager:REGION:ACCOUNT_ID:secret:prod/db/postgres*"
    }
  ]
}
```

Attach this policy to:

* **IRSA role** (recommended for EKS), or
* Node instance role (not recommended but works)

---

## 3. SecretStore (AWS Provider)

### `secretstore.yaml`

```yaml
apiVersion: external-secrets.io/v1beta1
kind: SecretStore
metadata:
  name: aws-secretsmanager
  namespace: default
spec:
  provider:
    aws:
      service: SecretsManager
      region: us-east-1
      auth:
        jwt:
          serviceAccountRef:
            name: external-secrets-sa
```

> If you’re **not using IRSA**, replace `auth.jwt` with:

```yaml
      auth:
        secretRef:
          accessKeyIDSecretRef:
            name: aws-creds
            key: access-key
          secretAccessKeySecretRef:
            name: aws-creds
            key: secret-key
```

---

## 4. ExternalSecret (Maps AWS → K8s Secret)

### `externalsecret-db.yaml`

```yaml
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: db-credentials
  namespace: default
spec:
  refreshInterval: 1h
  secretStoreRef:
    name: aws-secretsmanager
    kind: SecretStore
  target:
    name: db-credentials
    creationPolicy: Owner
  data:
    - secretKey: DB_USER
      remoteRef:
        key: prod/db/postgres
        property: username
    - secretKey: DB_PASSWORD
      remoteRef:
        key: prod/db/postgres
        property: password
    - secretKey: DB_HOST
      remoteRef:
        key: prod/db/postgres
        property: host
    - secretKey: DB_PORT
      remoteRef:
        key: prod/db/postgres
        property: port
    - secretKey: DB_NAME
      remoteRef:
        key: prod/db/postgres
        property: dbname
```

This creates a Kubernetes Secret:

```bash
kubectl get secret db-credentials -o yaml
```

---

## 5. Kubernetes Deployment (Golang App)

### `deployment.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-db-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-db-app
  template:
    metadata:
      labels:
        app: go-db-app
    spec:
      containers:
        - name: app
          image: yourrepo/go-db-app:latest
          env:
            - name: DB_USER
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_USER
            - name: DB_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_PASSWORD
            - name: DB_HOST
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_HOST
            - name: DB_PORT
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_PORT
            - name: DB_NAME
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_NAME
```

---

## 6. Golang Application Example

### `main.go`

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	log.Println("Successfully connected to the database 🎉")
}
```

---

## 7. Minimal Dockerfile

```dockerfile
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/app .
CMD ["/app/app"]
```

---

## 8. Flow Summary

```
AWS Secrets Manager
        ↓
External Secrets Operator
        ↓
Kubernetes Secret (db-credentials)
        ↓
Env Vars in Pod
        ↓
Golang app uses os.Getenv()
```

---

If you want, I can also:

* Convert this to **RDS + SSL**
* Show **ESO ClusterSecretStore**
* Use **file-based secrets instead of env vars**
* Add **Helm chart version**
* Add **automatic secret rotation handling**

Just tell me 👍
