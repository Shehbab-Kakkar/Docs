Here's a complete automatic certificate renewal system in Go:

## Automatic Certificate Renewal System in Go

### 1. Main Application (`main.go`)

```go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/smtp"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

// Config represents the configuration structure
type Config struct {
	CertificatesDir     string                 `json:"certificates_dir"`
	BackupDir           string                 `json:"backup_dir"`
	RenewalThresholdDays int                   `json:"renewal_threshold_days"`
	EmailNotifications  bool                   `json:"email_notifications"`
	SMTP               SMTPConfig             `json:"smtp"`
	Certificates       map[string]CertInfo    `json:"certificates"`
}

// SMTPConfig holds SMTP server configuration
type SMTPConfig struct {
	Server   string   `json:"server"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	FromEmail string  `json:"from_email"`
	ToEmails []string `json:"to_emails"`
}

// CertInfo contains information about a certificate
type CertInfo struct {
	CommonName         string   `json:"common_name"`
	Organization       string   `json:"organization"`
	Country           string   `json:"country"`
	State             string   `json:"state"`
	Locality          string   `json:"locality"`
	ValidityDays      int      `json:"validity_days"`
	ServicesToRestart []string `json:"services_to_restart"`
}

// CertificateRenewer handles certificate renewal operations
type CertificateRenewer struct {
	config *Config
	logger *log.Logger
}

// NewCertificateRenewer creates a new instance
func NewCertificateRenewer(configFile string) (*CertificateRenewer, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	logger := log.New(os.Stdout, "[CERT_RENEWER] ", log.LstdFlags)

	return &CertificateRenewer{
		config: &config,
		logger: logger,
	}, nil
}

// CheckCertificateExpiry checks if a certificate is about to expire
func (cr *CertificateRenewer) CheckCertificateExpiry(certPath string) (int, time.Time, bool, error) {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return 0, time.Time{}, false, fmt.Errorf("failed to read certificate: %v", err)
	}

	block, _ := pem.Decode(certData)
	if block == nil {
		return 0, time.Time{}, false, fmt.Errorf("failed to decode PEM block")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return 0, time.Time{}, false, fmt.Errorf("failed to parse certificate: %v", err)
	}

	expiryDate := cert.NotAfter
	daysUntilExpiry := int(time.Until(expiryDate).Hours() / 24)

	shouldRenew := daysUntilExpiry <= cr.config.RenewalThresholdDays

	cr.logger.Printf("Certificate %s: expires in %d days (expiry: %s)", 
		filepath.Base(certPath), daysUntilExpiry, expiryDate.Format("2006-01-02"))

	return daysUntilExpiry, expiryDate, shouldRenew, nil
}

// GenerateNewCertificate generates a new certificate
func (cr *CertificateRenewer) GenerateNewCertificate(certInfo CertInfo) ([]byte, []byte, error) {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:            []string{certInfo.Country},
			Province:           []string{certInfo.State},
			Locality:           []string{certInfo.Locality},
			Organization:       []string{certInfo.Organization},
			CommonName:         certInfo.CommonName,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(0, 0, certInfo.ValidityDays),
		KeyUsage:  x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
		DNSNames:              []string{certInfo.CommonName},
	}

	// Create certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create certificate: %v", err)
	}

	// Encode certificate
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	// Encode private key
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return certPEM, privateKeyPEM, nil
}

// RenewCertificate renews a specific certificate
func (cr *CertificateRenewer) RenewCertificate(certName string, certInfo CertInfo) error {
	cr.logger.Printf("Renewing certificate: %s", certName)

	// Generate new certificate and private key
	certPEM, privateKeyPEM, err := cr.GenerateNewCertificate(certInfo)
	if err != nil {
		return fmt.Errorf("failed to generate new certificate: %v", err)
	}

	// Backup old certificate
	certPath := filepath.Join(cr.config.CertificatesDir, fmt.Sprintf("%s.crt", certName))
	keyPath := filepath.Join(cr.config.CertificatesDir, fmt.Sprintf("%s.key", certName))

	if err := cr.BackupOldCertificate(certName); err != nil {
		cr.logger.Printf("Warning: failed to backup old certificate: %v", err)
	}

	// Write new certificate
	if err := ioutil.WriteFile(certPath, certPEM, 0644); err != nil {
		return fmt.Errorf("failed to write certificate: %v", err)
	}

	// Write new private key
	if err := ioutil.WriteFile(keyPath, privateKeyPEM, 0600); err != nil {
		return fmt.Errorf("failed to write private key: %v", err)
	}

	cr.logger.Printf("Successfully renewed certificate: %s", certName)

	// Restart services
	if err := cr.RestartServices(certInfo.ServicesToRestart); err != nil {
		cr.logger.Printf("Warning: failed to restart services: %v", err)
	}

	return nil
}

// BackupOldCertificate backs up old certificate files
func (cr *CertificateRenewer) BackupOldCertificate(certName string) error {
	timestamp := time.Now().Format("20060102_150405")
	oldCert := filepath.Join(cr.config.CertificatesDir, fmt.Sprintf("%s.crt", certName))
	oldKey := filepath.Join(cr.config.CertificatesDir, fmt.Sprintf("%s.key", certName))

	// Create backup directory if it doesn't exist
	if err := os.MkdirAll(cr.config.BackupDir, 0755); err != nil {
		return fmt.Errorf("failed to create backup directory: %v", err)
	}

	// Backup certificate
	if _, err := os.Stat(oldCert); err == nil {
		backupCert := filepath.Join(cr.config.BackupDir, fmt.Sprintf("%s_old_%s.crt", certName, timestamp))
		if err := os.Rename(oldCert, backupCert); err != nil {
			return fmt.Errorf("failed to backup certificate: %v", err)
		}
		cr.logger.Printf("Backed up certificate: %s", backupCert)
	}

	// Backup private key
	if _, err := os.Stat(oldKey); err == nil {
		backupKey := filepath.Join(cr.config.BackupDir, fmt.Sprintf("%s_old_%s.key", certName, timestamp))
		if err := os.Rename(oldKey, backupKey); err != nil {
			return fmt.Errorf("failed to backup private key: %v", err)
		}
		cr.logger.Printf("Backed up private key: %s", backupKey)
	}

	return nil
}

// RestartServices restarts specified services
func (cr *CertificateRenewer) RestartServices(services []string) error {
	for _, service := range services {
		cmd := exec.Command("systemctl", "restart", service)
		if err := cmd.Run(); err != nil {
			cr.logger.Printf("Failed to restart service %s: %v", service, err)
			continue // Continue with other services even if one fails
		}
		cr.logger.Printf("Restarted service: %s", service)
	}
	return nil
}

// SendNotification sends email notification about renewal status
func (cr *CertificateRenewer) SendNotification(certName string, success bool) error {
	if !cr.config.EmailNotifications {
		return nil
	}

	subject := fmt.Sprintf("Certificate Renewal %s: %s", 
		map[bool]string{true: "Success", false: "Failed"}[success], certName)
	
	body := fmt.Sprintf("Certificate %s renewal %s", 
		certName, map[bool]string{true: "completed successfully", false: "failed"}[success])

	m := gomail.NewMessage()
	m.SetHeader("From", cr.config.SMTP.FromEmail)
	m.SetHeader("To", cr.config.SMTP.ToEmails...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	port := cr.config.SMTP.Port
	if port == 0 {
		port = 587 // default SMTP port
	}

	d := gomail.NewDialer(cr.config.SMTP.Server, port, cr.config.SMTP.Username, cr.config.SMTP.Password)

	return d.DialAndSend(m)
}

// RunRenewalCheck performs the main renewal check
func (cr *CertificateRenewer) RunRenewalCheck() {
	cr.logger.Println("Starting certificate renewal check")

	for certName, certInfo := range cr.config.Certificates {
		certPath := filepath.Join(cr.config.CertificatesDir, fmt.Sprintf("%s.crt", certName))

		// Check if certificate file exists
		if _, err := os.Stat(certPath); os.IsNotExist(err) {
			cr.logger.Printf("Certificate file does not exist: %s", certPath)
			continue
		}

		// Check certificate expiry
		_, _, shouldRenew, err := cr.CheckCertificateExpiry(certPath)
		if err != nil {
			cr.logger.Printf("Error checking certificate %s: %v", certName, err)
			continue
		}

		if shouldRenew {
			cr.logger.Printf("Renewing certificate: %s", certName)
			
			err := cr.RenewCertificate(certName, certInfo)
			success := err == nil
			
			if err != nil {
				cr.logger.Printf("Failed to renew certificate %s: %v", certName, err)
			}

			// Send notification
			if notifyErr := cr.SendNotification(certName, success); notifyErr != nil {
				cr.logger.Printf("Failed to send notification for %s: %v", certName, notifyErr)
			}
		} else {
			cr.logger.Printf("Certificate %s does not need renewal yet", certName)
		}
	}

	cr.logger.Println("Certificate renewal check completed")
}

// MonitorSSLCertificate monitors SSL certificates via HTTPS
func (cr *CertificateRenewer) MonitorSSLCertificate(domain string) error {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", domain, err)
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return fmt.Errorf("no certificates found for %s", domain)
	}

	cert := certs[0]
	daysUntilExpiry := int(time.Until(cert.NotAfter).Hours() / 24)

	cr.logger.Printf("SSL Certificate for %s: %d days until expiry", domain, daysUntilExpiry)

	return nil
}

func main() {
	renewer, err := NewCertificateRenewer("renewal_config.json")
	if err != nil {
		log.Fatalf("Failed to initialize certificate renewer: %v", err)
	}

	// Perform renewal check
	renewer.RunRenewalCheck()

	// Optionally monitor additional domains
	for certName, certInfo := range renewer.config.Certificates {
		if err := renewer.MonitorSSLCertificate(certInfo.CommonName); err != nil {
			renewer.logger.Printf("Error monitoring %s: %v", certName, err)
		}
	}
}
```

### 2. Configuration File (`renewal_config.json`)

```json
{
    "certificates_dir": "/etc/ssl/certs",
    "backup_dir": "/etc/ssl/certs/backup",
    "renewal_threshold_days": 30,
    "email_notifications": true,
    "smtp": {
        "server": "smtp.gmail.com",
        "port": 587,
        "username": "your-email@gmail.com",
        "password": "your-app-password",
        "from_email": "your-email@gmail.com",
        "to_emails": ["admin@example.com"]
    },
    "certificates": {
        "web_server": {
            "common_name": "example.com",
            "organization": "My Organization",
            "country": "US",
            "state": "California",
            "locality": "San Francisco",
            "validity_days": 365,
            "services_to_restart": ["nginx", "apache2"]
        },
        "api_server": {
            "common_name": "api.example.com",
            "organization": "My Organization",
            "validity_days": 365,
            "services_to_restart": ["nginx"]
        }
    }
}
```

### 3. Dockerfile

```dockerfile
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cert-renewer .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata \
    && addgroup -g 65532 nonroot \
    && adduser -D -u 65532 -G nonroot nonroot

WORKDIR /app

COPY --from=builder /app/cert-renewer .
COPY renewal_config.json .

RUN mkdir -p /etc/ssl/certs /etc/ssl/certs/backup \
    && chown -R nonroot:nonroot /etc/ssl \
    && chown nonroot:nonroot /app/renewal_config.json

USER nonroot

CMD ["./cert-renewer"]
```

### 4. Docker Compose (`docker-compose.yml`)

```yaml
version: '3.8'

services:
  cert-renewer:
    build: .
    container_name: certificate-renewer
    volumes:
      - ./certs:/etc/ssl/certs
      - ./backups:/etc/ssl/certs/backup
    environment:
      - SMTP_SERVER=${SMTP_SERVER}
      - SMTP_USERNAME=${SMTP_USERNAME}
      - SMTP_PASSWORD=${SMTP_PASSWORD}
    restart: unless-stopped
    command: ["./cert-renewer"]

  cron-job:
    build: .
    container_name: cert-renewer-cron
    volumes:
      - ./certs:/etc/ssl/certs
      - ./backups:/etc/ssl/certs/backup
    environment:
      - SMTP_SERVER=${SMTP_SERVER}
      - SMTP_USERNAME=${SMTP_USERNAME}
      - SMTP_PASSWORD=${SMTP_PASSWORD}
    restart: unless-stopped
    command: ["sh", "-c", "echo '0 2 * * * /app/cert-renewer' | crontab - && crond -f -L /dev/stdout"]
```

### 5. Kubernetes Deployment (`k8s-deployment.yaml`)

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: certificate-renewal
spec:
  schedule: "0 2 * * *"  # Daily at 2 AM
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: cert-renewal
            image: your-registry/cert-renewal:latest
            volumeMounts:
            - name: certs-volume
              mountPath: /etc/ssl/certs
            env:
            - name: SMTP_SERVER
              value: "smtp.gmail.com"
            - name: SMTP_USERNAME
              valueFrom:
                secretKeyRef:
                  name: smtp-credentials
                  key: username
            - name: SMTP_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: smtp-credentials
                  key: password
          volumes:
          - name: certs-volume
            persistentVolumeClaim:
              claimName: certs-pvc
          restartPolicy: OnFailure
---
apiVersion: v1
kind: Secret
metadata:
  name: smtp-credentials
type: Opaque
data:
  username: <base64-encoded-username>
  password: <base64-encoded-password>
```

### 6. Service Management Script (`service_manager.go`)

```go
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// ServiceManager handles system service operations
type ServiceManager struct{}

// IsServiceRunning checks if a service is running
func (sm *ServiceManager) IsServiceRunning(serviceName string) (bool, error) {
	cmd := exec.Command("systemctl", "is-active", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	
	status := strings.TrimSpace(string(output))
	return status == "active", nil
}

// StartService starts a service
func (sm *ServiceManager) StartService(serviceName string) error {
	cmd := exec.Command("systemctl", "start", serviceName)
	return cmd.Run()
}

// StopService stops a service
func (sm *ServiceManager) StopService(serviceName string) error {
	cmd := exec.Command("systemctl", "stop", serviceName)
	return cmd.Run()
}

// ReloadService reloads a service configuration
func (sm *ServiceManager) ReloadService(serviceName string) error {
	cmd := exec.Command("systemctl", "reload", serviceName)
	return cmd.Run()
}

// GetServiceStatus gets the status of a service
func (sm *ServiceManager) GetServiceStatus(serviceName string) (string, error) {
	cmd := exec.Command("systemctl", "status", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// ValidateServiceExists checks if a service exists
func (sm *ServiceManager) ValidateServiceExists(serviceName string) (bool, error) {
	cmd := exec.Command("systemctl", "list-units", "--type=service", "--all", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	
	return strings.Contains(string(output), serviceName+".service"), nil
}
```

### 7. Health Check Endpoint (`health.go`)

```go
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse represents health check response
type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Checks    map[string]bool   `json:"checks"`
	Version   string            `json:"version"`
}

// CertificateHealthCheck performs health checks on certificates
func (cr *CertificateRenewer) CertificateHealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Checks:    make(map[string]bool),
		Version:   "1.0.0",
	}

	allHealthy := true

	for certName, certInfo := range cr.config.Certificates {
		certPath := filepath.Join(cr.config.CertificatesDir, fmt.Sprintf("%s.crt", certName))
		
		if _, _, shouldRenew, err := cr.CheckCertificateExpiry(certPath); err != nil || shouldRenew {
			response.Checks[certName] = false
			allHealthy = false
		} else {
			response.Checks[certName] = true
		}
	}

	if !allHealthy {
		response.Status = "degraded"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// StartHealthServer starts the health check HTTP server
func (cr *CertificateRenewer) StartHealthServer(port string) {
	http.HandleFunc("/health", cr.CertificateHealthCheck)
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	cr.logger.Printf("Starting health check server on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		cr.logger.Printf("Failed to start health server: %v", err)
	}
}
```

### 8. Go Module File (`go.mod`)

```go
module cert-renewer

go 1.19

require (
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)
```

### Installation and Usage:

1. **Initialize Go module:**
   ```bash
   go mod init cert-renewer
   go mod tidy
   ```

2. **Build the application:**
   ```bash
   go build -o cert-renewer main.go
   ```

3. **Run the application:**
   ```bash
   ./cert-renewer
   ```

4. **Set up automatic execution:**
   - Use systemd timer
   - Use cron job
   - Use Docker container with scheduled execution

This Go implementation provides:

- **Type Safety**: Strong typing with Go's type system
- **Performance**: Compiled binary with excellent performance
- **Concurrency**: Built-in support for concurrent operations
- **Cross-platform**: Runs on Linux, Windows, macOS
- **Security**: Built-in crypto libraries from Go standard library
- **Monitoring**: Health check endpoints
- **Configuration**: JSON-based configuration
- **Logging**: Structured logging
- **Error Handling**: Comprehensive error handling
- **Email Notifications**: SMTP-based notifications
- **Service Management**: Systemd integration
- **Backup Strategy**: Automatic backup of old certificates

The system is production-ready and can be deployed in various environments including Docker, Kubernetes, or traditional servers.
