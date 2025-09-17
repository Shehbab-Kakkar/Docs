It sounds like you're asking **how to check for deleted files still being used by the Linux OS**, i.e., **open but deleted files** that are still consuming disk space until the process using them is terminated.

Here’s how to do that:

---

### 🔍 **Check for Deleted Files Still in Use**

Use the `lsof` command to identify deleted files still held open by processes.

```bash
sudo lsof | grep '(deleted)'
```

**Explanation:**

* `lsof`: Lists open files.
* `grep '(deleted)'`: Filters the output to show only files that have been deleted but are still in use.

---

### 📌 Typical Output Example

```
COMMAND    PID  USER   FD   TYPE DEVICE SIZE/OFF    NODE NAME
python3   1234  root  txt    REG  8,1   10240    12345 /tmp/somefile (deleted)
```

This tells you:

* `PID`: The process ID still holding the deleted file open.
* `NAME`: The file path (if available).
* `(deleted)`: Indicates the file is removed from disk directory structure, but still open.

---

### 📉 Reclaiming Disk Space

To **free up the space**, you have two main options:

1. **Restart the process:**

   ```bash
   sudo kill -HUP <PID>
   ```

   Or if it’s safe:

   ```bash
   sudo systemctl restart <service-name>
   ```

2. **Clear specific file descriptors manually (advanced):**
   If you don’t want to kill the process, you can truncate the open file descriptor:

   ```bash
   sudo truncate -s 0 /proc/<PID>/fd/<FD>
   ```

   ⚠️ Be careful — this can affect the running process.

---

### ✅ Useful Tips

* You can list only deleted files with full paths using:

  ```bash
  sudo lsof -nP | grep '(deleted)'
  ```

* To monitor disk space reclaimed, use:

  ```bash
  df -h
  ```

---

Let me know if you want a script to monitor or clean up these files.
