# High Availability Web Streaming Application

This project is a **Highly Available Web Streaming Application** built with:
- **Backend:** Golang (API, streaming, orchestration)
- **Frontend:** Next.js (React-based SSR)
- **Container Orchestration:** AWS EKS (Kubernetes)
- **Storage:** AWS S3 (video chunks, thumbnails)
- **Database:** PostgreSQL (metadata)
- **Cache:** Redis (sessions, hot data)
- **CDN:** AWS CloudFront
- **Transcoding:** FFmpeg (K8s job)
- **Ingress:** AWS ALB or Nginx Ingress Controller

---

## Architecture Overview

```
User
  |
  v
[CloudFront CDN]
  |
  v
[ALB / Nginx Ingress (EKS)]
  |
  +------------+
  |            |
  v            v
[Next.js]   [Golang Backend API]
  |            |
  |            v
  |        [Redis, PostgreSQL]
  |            |
  |            v
  |        [FFmpeg Pods (K8s Jobs)]
  |            |
  v            v
[S3 for static, video segments, thumbnails]
```

---

## Main Features

- **User authentication** (JWT)
- **Video upload** via presigned S3 URLs
- **Automated transcoding** (FFmpeg) and HLS packaging
- **Highly available streaming** via CloudFront CDN and S3
- **Horizontal scalability** for all stateless services
- **Metrics and logging** via Prometheus, Grafana, ELK

---

## Quick Start (Development)

### 1. Clone the Repository

```bash
git clone https://github.com/your-org/ha-web-streaming.git
cd ha-web-streaming
```

### 2. Environment Setup

Create a `.env` file for backend:

```env
# Golang Backend .env example
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/stream_db
REDIS_URL=redis://localhost:6379
AWS_REGION=us-east-1
S3_BUCKET=your-bucket
JWT_SECRET=your-secret
```

And for Next.js frontend:

```env
# Next.js .env.local example
NEXT_PUBLIC_API_URL=http://localhost:8080
CLOUDFRONT_URL=https://dxxxx.cloudfront.net
```

---

### 3. Local Development (Docker Compose)

```yaml name=docker-compose.yaml
version: '3.8'
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: stream_db
    ports:
      - 5432:5432

  redis:
    image: redis:7
    ports:
      - 6379:6379

  backend:
    build: ./backend
    env_file: .env
    ports:
      - 8080:8080
    depends_on:
      - postgres
      - redis

  frontend:
    build: ./frontend
    env_file: ./frontend/.env.local
    ports:
      - 3000:3000
    depends_on:
      - backend
```

Start all services:

```bash
docker-compose up --build
```

---

## Example API Usage

### 1. Request Presigned Upload URL

```http
POST /api/videos/upload
Authorization: Bearer <JWT>
{
  "filename": "myvideo.mp4",
  "contentType": "video/mp4"
}
```

**Response:**
```json
{
  "uploadUrl": "https://s3.amazonaws.com/your-bucket/...",
  "videoId": "abc123"
}
```

### 2. Trigger Transcoding (Automatic via S3 Event or Manual)

```http
POST /api/videos/abc123/process
Authorization: Bearer <JWT>
```

---

## Sample K8s Deployments

### Golang Backend

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  replicas: 4
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
      - name: backend
        image: yourrepo/backend:latest
        envFrom:
        - configMapRef:
            name: backend-config
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  selector:
    app: backend
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: ClusterIP
```

### Next.js Frontend

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  replicas: 3
  selector:
    matchLabels:
      app: frontend
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
      - name: frontend
        image: yourrepo/frontend:latest
        envFrom:
        - configMapRef:
            name: frontend-config
        ports:
        - containerPort: 3000
---
apiVersion: v1
kind: Service
metadata:
  name: frontend
spec:
  selector:
    app: frontend
  ports:
  - protocol: TCP
    port: 80
    targetPort: 3000
  type: ClusterIP
```

### FFmpeg Processor (as K8s Job)

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: ffmpeg-job
spec:
  template:
    spec:
      containers:
      - name: ffmpeg
        image: yourrepo/ffmpeg-processor:latest
        env:
        - name: S3_BUCKET
          value: your-bucket
        - name: INPUT_FILE
          value: s3://your-bucket/input/video.mp4
        - name: OUTPUT_PREFIX
          value: s3://your-bucket/output/
      restartPolicy: Never
```

### Ingress (Nginx or AWS ALB)

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: app-ingress
  annotations:
    alb.ingress.kubernetes.io/scheme: internet-facing
spec:
  rules:
  - host: yourapp.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: frontend
            port:
              number: 80
      - path: /api
        pathType: Prefix
        backend:
          service:
            name: backend
            port:
              number: 80
```

---

## Service Communication Table

| Service          | Communicates With      | Protocol | Endpoint Example                              |
|------------------|-----------------------|----------|-----------------------------------------------|
| Next.js Frontend | Golang Backend        | HTTP     | http://backend.default.svc.cluster.local      |
| Golang Backend   | PostgreSQL            | TCP      | host: db.instance.rds.amazonaws.com:5432      |
| Golang Backend   | Redis                 | TCP      | redis.default.svc.cluster.local:6379          |
| Golang Backend   | S3                    | HTTPS    | s3.amazonaws.com/your-bucket                  |
| Golang Backend   | FFmpeg Processor      | REST/K8s | K8s Job API                                   |
| FFmpeg Processor | S3                    | HTTPS    | s3.amazonaws.com/your-bucket                  |
| User Browser     | CloudFront CDN        | HTTPS    | https://dxxxxx.cloudfront.net/video/playlist.m3u8 |

---

## Security Best Practices

- All traffic via HTTPS (TLS termination at ALB)
- IAM roles for K8s service accounts to access S3
- JWT for auth, Redis for token revocation
- S3 bucket policies restrict public access; only CloudFront allowed

---

## Observability

- **Metrics:** Prometheus, Grafana
- **Logs:** ELK stack (Elasticsearch, Logstash, Kibana)
- **Tracing:** Jaeger (optional)

---

## Contribution

1. Fork the repo
2. Create your branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'feat: add foo'`)
4. Push to the branch (`git push origin feature/foo`)
5. Open a pull request

---

## License

MIT

---

## Notes

- For production, use managed PostgreSQL (RDS) and Redis (ElastiCache).
- Ensure S3/CloudFront properly configured for public video delivery via signed URLs if needed.
- For live streaming, extend the architecture with NGINX RTMP or AWS MediaLive.

---

**Happy Streaming!**
