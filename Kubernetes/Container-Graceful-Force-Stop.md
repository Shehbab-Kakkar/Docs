When you **kill** a Docker container or **stop** it gracefully, Docker sends **specific Unix signals** to the container’s main process (PID 1 inside the container). The signal issued depends on the command used.

---

## 🛑 1. When you run: `docker stop <container>`

* Docker sends: **`SIGTERM`** (signal number **15**) to the container’s **PID 1**.
* After a **10-second timeout** (by default), if the process hasn't exited, Docker sends: **`SIGKILL`** (signal number **9**), which **immediately terminates** the process.

### ⏱ You can adjust the timeout:

```bash
docker stop -t 30 <container>
```

This gives the process 30 seconds to exit gracefully.

---

## 🔪 2. When you run: `docker kill <container>`

* Docker sends: **`SIGKILL`** (by default) immediately to the main process — **no chance to clean up**.
* You can optionally specify a different signal:

```bash
docker kill --signal=SIGTERM <container>
docker kill --signal=SIGHUP <container>
```

---

## 💡 Summary Table

| Command               | Signal Sent                       | Graceful? | Notes                             |
| --------------------- | --------------------------------- | --------- | --------------------------------- |
| `docker stop`         | SIGTERM → SIGKILL (after timeout) | ✅ Yes     | Default timeout is 10s            |
| `docker kill`         | SIGKILL (default)                 | ❌ No      | Can be overridden with `--signal` |
| `kill -s SIGTERM PID` | SIGTERM                           | ✅ Yes     | Used inside or outside container  |

---

## 📦 Kubernetes Context (Bonus)

In **Kubernetes**, when a **Pod is terminated**, the kubelet sends a **`SIGTERM`** to the container's main process, and after a **grace period** (`terminationGracePeriodSeconds`), it sends **`SIGKILL`**.

So Docker and Kubernetes behave similarly in this regard.

---

Would you like to see how to handle `SIGTERM` properly in your containerized app (e.g., in Python, Node.js, etc.)?
