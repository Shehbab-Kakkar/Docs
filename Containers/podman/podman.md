Great! If you want to **build and deploy container images using Podman** (instead of Jenkins or Docker) from the **CLI**, here’s a step-by-step guide:

---

# **1. Build a Container Image with Podman**

Assuming you have a `Dockerfile` in your project directory:

```sh
# Build the image (replace myimage:latest with your desired image name and tag)
podman build -t myimage:latest .
```

- `-t myimage:latest`: Tags the image as `myimage` with the `latest` tag.
- `.`: The build context (current directory).

---

# **2. List Your Images**

```sh
podman images
```

---

# **3. Run a Container from Your Image**

```sh
# Run a container from the image (replace myimage:latest as needed)
podman run --name mycontainer -d -p 8080:80 myimage:latest
```

- `--name mycontainer`: Names the running container.
- `-d`: Detached mode (runs in the background).
- `-p 8080:80`: Maps port 8080 on your host to port 80 in the container.

---

# **4. View Running Containers**

```sh
podman ps
```

---

# **5. Push Image to a Registry (Optional)**

If you want to deploy on another machine or share the image:

```sh
# Tag your image for the registry (example: quay.io/youruser/myimage:latest)
podman tag myimage:latest quay.io/youruser/myimage:latest

# Login to the registry
podman login quay.io

# Push the image
podman push quay.io/youruser/myimage:latest
```

---

# **6. Remove a Running Container**

```sh
podman stop mycontainer
podman rm mycontainer
```

---

# **7. Remove an Image**

```sh
podman rmi myimage:latest
```

---

## **Example: Full Workflow**

```sh
# 1. Build the image
podman build -t myapp:1.0 .

# 2. Run the container
podman run --name myapp-test -d -p 8080:80 myapp:1.0

# 3. Stop and remove the container
podman stop myapp-test
podman rm myapp-test

# 4. (Optional) Push to registry
podman tag myapp:1.0 quay.io/shehbabk/myapp:1.0
podman login quay.io
podman push quay.io/shehbabk/myapp:1.0
```

---

## **References**
- [Podman Official Docs](https://docs.podman.io/)
- [Podman Build](https://docs.podman.io/en/latest/markdown/podman-build.1.html)
- [Podman Run](https://docs.podman.io/en/latest/markdown/podman-run.1.html)

---

**If you need an example Dockerfile, or want to run with specific environment variables, volumes, or in rootless mode, just ask!**
