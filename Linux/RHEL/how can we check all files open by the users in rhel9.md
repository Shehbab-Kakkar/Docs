Ah! You want to see **all open files per user** on RHEL 9. On Linux, this is typically done with the **`lsof` command**. Let me break it down **in detail**, including examples and tips for root cause analysis.

---

## 1️⃣ Using `lsof` to list open files

`lsof` = **LiSt Open Files**.
On Linux, everything is a file — regular files, directories, sockets, pipes, devices, network connections, etc.

### Basic usage

```bash
lsof
```

This lists **all open files on the system**.
Not very readable? Narrow it down.

---

## 2️⃣ List open files by a specific user

```bash
lsof -u username
```

Example:

```bash
lsof -u alice
```

Output columns:

| Column   | Meaning                                                 |
| -------- | ------------------------------------------------------- |
| COMMAND  | Process name                                            |
| PID      | Process ID                                              |
| USER     | Owner                                                   |
| FD       | File descriptor (e.g., `cwd`, `txt`, `mem`, `1u`)       |
| TYPE     | File type (`REG`=regular, `DIR`, `CHR`, `FIFO`, `SOCK`) |
| SIZE/OFF | File size / offset                                      |
| NODE     | Inode                                                   |
| NAME     | Path or network info                                    |

---

## 3️⃣ Count open files per user

Sometimes you want to **see if a user is hitting file limits**:

```bash
lsof -u alice | wc -l
```

---

## 4️⃣ List open files for all users **grouped by user**

```bash
lsof -F -u $(cut -f1 -d: /etc/passwd) | awk -F: '{print $2}' | sort | uniq -c
```

Or simpler:

```bash
lsof | awk '{print $3}' | sort | uniq -c | sort -nr
```

* Shows **which users have the most open files**
* Great for debugging **“too many open files” issues**

---

## 5️⃣ Check open files for a specific process

```bash
lsof -p PID
```

Example:

```bash
lsof -p 1234
```

* Useful if you know a process is **leaking file descriptors**

---

## 6️⃣ Check network connections opened by a user

```bash
lsof -u alice -i
```

Example output:

```
COMMAND  PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
ssh      3456 alice  3u  IPv4  23456      0t0  TCP 192.168.1.2:22->192.168.1.10:54321 (ESTABLISHED)
```

---

## 7️⃣ Check all deleted files still held open (can cause disk full)

Sometimes a process keeps a deleted file open, eating disk space:

```bash
lsof +L1
```

Filter by user:

```bash
lsof +L1 -u alice
```

---

## 8️⃣ Check limits (for root cause)

Even if `lsof` shows lots of files, the **user might hit limits**:

```bash
ulimit -n   # max open files for shell
cat /etc/security/limits.conf
```

---

### ✅ Summary of commands

| Goal                     | Command               |        |
| ------------------------ | --------------------- | ------ |
| All open files           | `lsof`                |        |
| Open files by user       | `lsof -u username`    |        |
| Count open files         | `lsof -u username     | wc -l` |
| Open files by process    | `lsof -p PID`         |        |
| Open network connections | `lsof -u username -i` |        |
| Deleted files still open | `lsof +L1`            |        |

---

If you want, I can give you a **ready-to-run one-liner** to **list all users and all their open files, sorted by count**, which is extremely useful for spotting runaway processes.

Do you want me to do that?
