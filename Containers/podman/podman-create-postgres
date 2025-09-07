To create and run a container from your pulled Postgres image using **Podman** (the equivalent of your Docker command), you simply use:

```sh
podman run --name some-postgres -e POSTGRES_PASSWORD=secret -d postgres
```

### **Explanation**
- `podman run`: Start a new container.
- `--name some-postgres`: Name the container "some-postgres".
- `-e POSTGRES_PASSWORD=secret`: Set the environment variable to configure the database password.
- `-d`: Run the container in detached mode (in the background).
- `postgres`: Use the `postgres` image (it will use `latest` by default).

---

### **Extra (if you want to map the port to your host)**
```sh
podman run --name some-postgres -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
```
This will allow you to connect to Postgres via `localhost:5432` on your machine.

---

### **If you want to do this using Postman (the API tool):**
- **Postman is not designed to interact with Podman/Docker directly.**  
- It can only trigger remote APIs (like Docker’s HTTP API, if enabled, or Podman’s REST API if set up).
- **Typical local usage is via CLI, not via Postman.**

If you want to control Podman remotely (e.g., via Podman’s REST API), let me know and I can show how to set up and use Podman’s API endpoints with Postman. But for most local scenarios, you use the CLI command above.

---

To **login to a running Postgres container** (whether created via Podman or Docker) and run SQL commands, follow these steps:

---

## **1. Get the Container Name or ID**

You already named your container (for example, `some-postgres`).  
If you forgot, list containers:

```sh
podman ps
```

---

## **2. Open a Shell Inside the Container**

Open a shell inside the running container:

```sh
podman exec -it some-postgres bash
```
- `-it`: Interactive terminal.
- `some-postgres`: The container name.

---

## **3. Login to the Postgres Database**

Once inside the container shell, use the `psql` client to connect:

```sh
psql -U postgres
```
- `-U postgres`: Connect as the `postgres` user (default superuser).

If you set a different user, replace `postgres` with that username.

---

## **4. Run SQL Commands**

Now you’re at the `psql` prompt. You can run SQL like:

```sql
\l              -- List databases
\dt             -- List tables (in the current database)
CREATE DATABASE testdb;
```

To quit the `psql` prompt, type:

```sql
\q
```
To exit the container shell:

```sh
exit
```

---

## **One-Line Command (without entering bash)**

You can also run a SQL command directly from your host:

```sh
podman exec -it some-postgres psql -U postgres -c "SELECT version();"
```
This will print the Postgres version.

---

## **Summary**

1. `podman exec -it some-postgres bash`  
2. `psql -U postgres`  
3. Run your SQL commands  
4. `\q` to quit psql, `exit` to leave the container shell

Let me know if you want to connect from your host instead of inside the container, or need help with users/passwords!
