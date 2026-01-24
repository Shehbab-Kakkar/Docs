In **RHEL 9**, *all processes stop during shutdown* because the system is managed by **systemd**, which has a very structured and intentional shutdown sequence.

Here’s what’s happening under the hood 👇

---

## High-level flow (RHEL 9 shutdown)

When you run:

```bash
shutdown -h now
# or
systemctl poweroff
```

this is what systemd does:

---

## 1️⃣ systemd enters shutdown mode

* The system switches to a special **shutdown target**
* No new services or user processes are allowed to start

---

## 2️⃣ Services are stopped cleanly (SIGTERM)

Systemd stops services **in reverse order of startup**, based on dependencies.

* Each service gets **SIGTERM**
* This tells the process:
  *“Hey, shut down gracefully”*

Examples:

* Databases flush data
* Web servers stop accepting connections
* Apps save state

You can see this behavior in service unit files:

```ini
TimeoutStopSec=90
KillSignal=SIGTERM
```

---

## 3️⃣ Forced kill if needed (SIGKILL)

If a process **ignores SIGTERM** and doesn’t exit in time:

* systemd sends **SIGKILL**
* SIGKILL **cannot be ignored**
* Process is terminated immediately

This guarantees *nothing survives shutdown*.

---

## 4️⃣ User processes are killed

All user sessions are terminated:

* systemd-logind kills:

  * SSH sessions
  * TTY sessions
  * GUI logins
* Remaining user processes are stopped

Controlled by:

```bash
KillUserProcesses=yes
```

(check with)

```bash
loginctl show-logind | grep KillUserProcesses
```

---

## 5️⃣ Filesystems are synced & unmounted

Before power-off:

* `sync()` is called → writes all buffered data to disk
* Filesystems are cleanly unmounted
* Swap is disabled

---

## 6️⃣ Kernel takes over

Finally:

* systemd tells the kernel to:

  * halt / reboot / power off
* Hardware power is cut

---

## 🔍 Key signals used

| Signal  | Purpose            |
| ------- | ------------------ |
| SIGTERM | Graceful shutdown  |
| SIGKILL | Forced termination |

---

## 🧠 Why this is reliable

* Dependency-aware (services stop in the correct order)
* Graceful first, forceful only if needed
* Prevents data corruption
* Ensures **zero orphan processes**

---

## Want to go deeper?

I can show you:

* How to **delay shutdown** for a service
* How to **prevent a service from being killed**
* How to **debug stuck shutdowns**
* The exact **systemd targets involved**

Just tell me what angle you want 👍
