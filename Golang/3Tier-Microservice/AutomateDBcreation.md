Ah! I see exactly what you want 😄 — you don’t want to manually exec into MySQL every time. You want **full automation**: database, user, permissions, and table setup — all ready when MySQL starts.

The best way to do this in Kubernetes (or even plain Docker) is to use **MySQL initialization scripts**. MySQL official images automatically run any `.sql` or `.sh` scripts in `/docker-entrypoint-initdb.d/` on first startup.

---

## ✅ Automated MySQL Setup

1. **Create an SQL init file**:

Call it `init-appdb.sql`:

```sql
-- Create database
CREATE DATABASE IF NOT EXISTS appdb;

-- Create users and grant privileges
CREATE USER IF NOT EXISTS 'appuser'@'%' IDENTIFIED BY 'apppass';
GRANT ALL PRIVILEGES ON appdb.* TO 'appuser'@'%';

CREATE USER IF NOT EXISTS 'appuser'@'localhost' IDENTIFIED BY 'apppass';
GRANT ALL PRIVILEGES ON appdb.* TO 'appuser'@'localhost';

FLUSH PRIVILEGES;

-- Create table and initialize
USE appdb;

CREATE TABLE IF NOT EXISTS counter (
  id INT PRIMARY KEY,
  visits INT
);

INSERT INTO counter (id, visits)
VALUES (1, 0)
ON DUPLICATE KEY UPDATE id=id;
```

2. **Mount this file in your MySQL Pod / container**:

### Docker (local dev)

If you’re testing locally:

```bash
docker run -d \
  --name mysql \
  -e MYSQL_ROOT_PASSWORD=rootpass \
  -v $(pwd)/init-appdb.sql:/docker-entrypoint-initdb.d/init-appdb.sql:ro \
  mysql:8
```

* MySQL will automatically run this script on **first container startup**.
* No manual `kubectl exec` needed.

---

### Kubernetes (Helm / Deployment)

If using Kubernetes, create a **ConfigMap**:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql-initdb
data:
  init-appdb.sql: |
    CREATE DATABASE IF NOT EXISTS appdb;
    CREATE USER IF NOT EXISTS 'appuser'@'%' IDENTIFIED BY 'apppass';
    GRANT ALL PRIVILEGES ON appdb.* TO 'appuser'@'%';
    CREATE USER IF NOT EXISTS 'appuser'@'localhost' IDENTIFIED BY 'apppass';
    GRANT ALL PRIVILEGES ON appdb.* TO 'appuser'@'localhost';
    FLUSH PRIVILEGES;
    USE appdb;
    CREATE TABLE IF NOT EXISTS counter (
      id INT PRIMARY KEY,
      visits INT
    );
    INSERT INTO counter (id, visits)
    VALUES (1, 0)
    ON DUPLICATE KEY UPDATE id=id;
```

Then mount it in the MySQL Pod:

```yaml
        volumeMounts:
        - name: initdb
          mountPath: /docker-entrypoint-initdb.d
          readOnly: true
      volumes:
      - name: initdb
        configMap:
          name: mysql-initdb
```

✅ Kubernetes MySQL will run the script automatically on the **first startup** of the pod.

---

If you want, I can write a **full automated Kubernetes Deployment + Service for MySQL** that:

* Sets root password
* Runs your init script automatically
* Is ready for your Go app to connect immediately

This way, **no manual `kubectl exec` ever again**.

Do you want me to do that?
