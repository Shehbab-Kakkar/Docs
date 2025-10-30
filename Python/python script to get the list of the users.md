For example:

* 🧑‍💻 **System users** on a Linux or Windows machine
* 👥 **Users in a web application** (e.g., from a database or API)
* ☁️ **Cloud users** (e.g., AWS IAM, Azure AD, etc.)

Here are examples for the most common cases:

---

### 🧩 **1. List system users (Linux / macOS)**

```python
import pwd

def list_system_users():
    users = [user.pw_name for user in pwd.getpwall()]
    return users

if __name__ == "__main__":
    for user in list_system_users():
        print(user)
```

✅ Works on Unix-like systems (Linux, macOS).
❌ Not for Windows.

---

### 🪟 **2. List local users on Windows**

```python
import subprocess

def list_windows_users():
    result = subprocess.run(["net", "user"], capture_output=True, text=True)
    lines = result.stdout.splitlines()
    users_section = False
    users = []

    for line in lines:
        if "-------" in line:
            if users_section:
                break
            users_section = True
            continue
        if users_section:
            users.extend(line.split())
    return users

if __name__ == "__main__":
    for user in list_windows_users():
        print(user)
```

✅ Works on Windows using the built-in `net user` command.

---

### 🗄️ **3. List users from a database (e.g., MySQL)**

```python
import mysql.connector

def list_db_users():
    conn = mysql.connector.connect(
        host="localhost",
        user="root",
        password="your_password",
        database="your_database"
    )
    cursor = conn.cursor()
    cursor.execute("SELECT username FROM users;")
    for (username,) in cursor.fetchall():
        print(username)
    conn.close()

if __name__ == "__main__":
    list_db_users()
```

✅ Replace `your_database` and table names as needed.

---

