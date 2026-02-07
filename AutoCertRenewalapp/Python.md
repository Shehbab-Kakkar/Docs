I'll help you develop an automatic certificate renewal system. Here's a comprehensive solution using Python:

## Automatic Certificate Renewal System

### 1. Core Renewal Script (`certificate_renewer.py`)

```python
import os
import json
import logging
from datetime import datetime, timedelta
from cryptography import x509
from cryptography.hazmat.primitives import hashes, serialization
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.x509.oid import NameOID
import subprocess
import smtplib
from email.mime.text import MimeText
from email.mime.multipart import MimeMultipart

class CertificateRenewer:
    def __init__(self, config_file="renewal_config.json"):
        self.config = self.load_config(config_file)
        self.setup_logging()
    
    def load_config(self, config_file):
        """Load configuration from JSON file"""
        with open(config_file, 'r') as f:
            return json.load(f)
    
    def setup_logging(self):
        """Setup logging for the renewal process"""
        logging.basicConfig(
            level=logging.INFO,
            format='%(asctime)s - %(levelname)s - %(message)s',
            handlers=[
                logging.FileHandler('certificate_renewal.log'),
                logging.StreamHandler()
            ]
        )
        self.logger = logging.getLogger(__name__)
    
    def check_certificate_expiry(self, cert_path):
        """Check if certificate is about to expire"""
        with open(cert_path, 'rb') as f:
            cert_data = f.read()
        
        cert = x509.load_pem_x509_certificate(cert_data)
        expiry_date = cert.not_valid_after
        days_until_expiry = (expiry_date - datetime.now()).days
        
        return {
            'expires_in_days': days_until_expiry,
            'expiry_date': expiry_date,
            'should_renew': days_until_expiry <= self.config['renewal_threshold_days']
        }
    
    def generate_new_certificate(self, cert_info):
        """Generate a new certificate"""
        # Generate private key
        private_key = rsa.generate_private_key(
            public_exponent=65537,
            key_size=2048
        )
        
        # Create certificate
        subject = issuer = x509.Name([
            x509.NameAttribute(NameOID.COUNTRY_NAME, cert_info.get('country', 'US')),
            x509.NameAttribute(NameOID.STATE_OR_PROVINCE_NAME, cert_info.get('state', 'CA')),
            x509.NameAttribute(NameOID.LOCALITY_NAME, cert_info.get('locality', 'San Francisco')),
            x509.NameAttribute(NameOID.ORGANIZATION_NAME, cert_info.get('organization', 'MyOrg')),
            x509.NameAttribute(NameOID.COMMON_NAME, cert_info['common_name']),
        ])
        
        cert = x509.CertificateBuilder().subject_name(
            subject
        ).issuer_name(
            issuer
        ).public_key(
            private_key.public_key()
        ).serial_number(
            x509.random_serial_number()
        ).not_valid_before(
            datetime.utcnow()
        ).not_valid_after(
            datetime.utcnow() + timedelta(days=cert_info.get('validity_days', 365))
        ).add_extension(
            x509.SubjectAlternativeName([x509.DNSName(cert_info['common_name'])]),
            critical=False,
        ).sign(private_key, hashes.SHA256())
        
        return cert, private_key
    
    def renew_certificate(self, cert_name, cert_info):
        """Renew a specific certificate"""
        try:
            # Generate new certificate and private key
            cert, private_key = self.generate_new_certificate(cert_info)
            
            # Write new certificate
            cert_path = os.path.join(self.config['certificates_dir'], f"{cert_name}.crt")
            with open(cert_path, 'wb') as f:
                f.write(cert.public_bytes(serialization.Encoding.PEM))
            
            # Write new private key
            key_path = os.path.join(self.config['certificates_dir'], f"{cert_name}.key")
            with open(key_path, 'wb') as f:
                f.write(private_key.private_bytes(
                    encoding=serialization.Encoding.PEM,
                    format=serialization.PrivateFormat.TraditionalOpenSSL,
                    encryption_algorithm=serialization.NoEncryption()
                ))
            
            self.logger.info(f"Certificate {cert_name} renewed successfully")
            
            # Backup old certificate if needed
            self.backup_old_certificate(cert_name)
            
            # Restart services if configured
            self.restart_services(cert_info.get('services_to_restart', []))
            
            return True
            
        except Exception as e:
            self.logger.error(f"Failed to renew certificate {cert_name}: {str(e)}")
            return False
    
    def backup_old_certificate(self, cert_name):
        """Backup old certificate files"""
        timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
        old_cert = os.path.join(self.config['certificates_dir'], f"{cert_name}.crt")
        old_key = os.path.join(self.config['certificates_dir'], f"{cert_name}.key")
        
        if os.path.exists(old_cert):
            backup_cert = os.path.join(self.config['backup_dir'], f"{cert_name}_old_{timestamp}.crt")
            os.rename(old_cert, backup_cert)
        
        if os.path.exists(old_key):
            backup_key = os.path.join(self.config['backup_dir'], f"{cert_name}_old_{timestamp}.key")
            os.rename(old_key, backup_key)
    
    def restart_services(self, services):
        """Restart specified services after certificate renewal"""
        for service in services:
            try:
                subprocess.run(['sudo', 'systemctl', 'restart', service], check=True)
                self.logger.info(f"Service {service} restarted")
            except subprocess.CalledProcessError as e:
                self.logger.error(f"Failed to restart service {service}: {str(e)}")
    
    def send_notification(self, cert_name, success=True):
        """Send notification about renewal status"""
        if not self.config.get('email_notifications', False):
            return
        
        try:
            msg = MimeMultipart()
            msg['Subject'] = f"Certificate Renewal {'Success' if success else 'Failed'}: {cert_name}"
            msg['From'] = self.config['smtp']['from_email']
            msg['To'] = ', '.join(self.config['smtp']['to_emails'])
            
            body = f"Certificate {cert_name} renewal {'completed successfully' if success else 'failed'}"
            msg.attach(MimeText(body, 'plain'))
            
            server = smtplib.SMTP(self.config['smtp']['server'], self.config['smtp']['port'])
            server.starttls()
            server.login(self.config['smtp']['username'], self.config['smtp']['password'])
            server.send_message(msg)
            server.quit()
            
        except Exception as e:
            self.logger.error(f"Failed to send notification: {str(e)}")
    
    def run_renewal_check(self):
        """Main function to check and renew certificates"""
        self.logger.info("Starting certificate renewal check")
        
        for cert_name, cert_info in self.config['certificates'].items():
            cert_path = os.path.join(self.config['certificates_dir'], f"{cert_name}.crt")
            
            if not os.path.exists(cert_path):
                self.logger.warning(f"Certificate file does not exist: {cert_path}")
                continue
            
            expiry_info = self.check_certificate_expiry(cert_path)
            
            self.logger.info(f"Certificate {cert_name}: expires in {expiry_info['expires_in_days']} days")
            
            if expiry_info['should_renew']:
                self.logger.info(f"Renewing certificate {cert_name}")
                success = self.renew_certificate(cert_name, cert_info)
                self.send_notification(cert_name, success)
            else:
                self.logger.info(f"Certificate {cert_name} does not need renewal yet")

if __name__ == "__main__":
    renewer = CertificateRenewer()
    renewer.run_renewal_check()
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

### 3. Docker Container Version (`Dockerfile`)

```dockerfile
FROM python:3.9-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

# Set up cron job for automatic renewal
RUN apt-get update && apt-get install -y cron
COPY crontab /etc/cron.d/cert-renewal
RUN chmod 0644 /etc/cron.d/cert-renewal
RUN crontab /etc/cron.d/cert-renewal

CMD ["cron", "-f"]
```

### 4. Requirements File (`requirements.txt`)

```
cryptography>=3.4.8
requests>=2.25.1
```

### 5. Cron Job Setup (`crontab`)

```
# Check and renew certificates daily at 2 AM
0 2 * * * cd /app && python certificate_renewer.py >> /var/log/cert_renewal.log 2>&1
```

### 6. Kubernetes Deployment Example (`k8s-deployment.yaml`)

```yaml
apiVersion: batch/v2alpha1
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
```

### 7. Monitoring Script (`monitor.py`)

```python
#!/usr/bin/env python3
import requests
import json
from datetime import datetime, timedelta

def monitor_ssl_certificates():
    """Monitor SSL certificates via API endpoint"""
    config = json.load(open('renewal_config.json'))
    
    for cert_name, cert_info in config['certificates'].items():
        domain = cert_info['common_name']
        
        try:
            response = requests.get(f"https://{domain}", timeout=10)
            cert_expiry = response.raw._fp.fp._sock.getpeercert()['notAfter']
            expiry_date = datetime.strptime(cert_expiry, '%b %d %H:%M:%S %Y %Z')
            days_until_expiry = (expiry_date - datetime.now()).days
            
            print(f"Certificate for {domain}: {days_until_expiry} days until expiry")
            
        except Exception as e:
            print(f"Error checking certificate for {domain}: {str(e)}")

if __name__ == "__main__":
    monitor_ssl_certificates()
```

### Installation and Usage:

1. **Install dependencies:**
   ```bash
   pip install -r requirements.txt
   ```

2. **Configure the system:**
   - Update `renewal_config.json` with your certificate details
   - Set up proper directory permissions
   - Configure SMTP settings for notifications

3. **Run manually:**
   ```bash
   python certificate_renewer.py
   ```

4. **Set up automated execution:**
   - Add to cron: `crontab -e` and add the cron job
   - Or use the Docker/Kubernetes deployment

This system provides:
- Automated certificate generation and renewal
- Email notifications
- Service restart capabilities
- Backup of old certificates
- Comprehensive logging
- Configurable renewal thresholds
- Support for multiple certificates

Would you like me to customize this further for your specific use case or explain any particular component in more detail?
