If a **Linux partition shows used space**, but the **mount point appears empty**, **yet no data is visible**, the issue is almost always **one of the following**:

---

## 🔍 Common Cause: **Mounting Over Existing Data**

When a partition is mounted on a directory that already had data, **that original data becomes hidden** — it’s not deleted, just hidden **under the mount point**.

### 🧠 Example:

Let’s say:

* `/mnt/data` has 2 GB of files.
* Then you mount a new partition on `/mnt/data`:

  ```bash
  mount /dev/sdb1 /mnt/data
  ```

Now, when you do:

```bash
ls /mnt/data
```

→ You’ll only see the **contents of the mounted partition** (`/dev/sdb1`), not the original 2 GB of data, which is still on the root filesystem — but now **inaccessible via `/mnt/data`**.

### ✅ How to Check:

1. **Unmount the partition:**

   ```bash
   sudo umount /mnt/data
   ```

2. **Check what’s really in the directory:**

   ```bash
   ls -la /mnt/data
   du -sh /mnt/data
   ```

If you now see the missing files, **that was the problem**.

---

## 🧤 Another Possibility: Deleted Files Still in Use

As mentioned earlier, files deleted but still **held open by running processes** can still consume space, but are **not visible in the filesystem**.

### ✅ How to Check:

```bash
sudo lsof | grep '(deleted)'
```

If you find files listed here, they're still consuming space until the owning process is restarted or the file is closed.

---

## 📦 Hidden Files/Directories

Files that start with a dot (`.`) are hidden by default. Try:

```bash
ls -la /mount/point
```

Also check for nested subdirectories or symlinks.

---

## 🧰 Summary of Checks

| Check                           | Command                                       |                    |
| ------------------------------- | --------------------------------------------- | ------------------ |
| Unmount and check original data | `sudo umount /mnt/point && ls -la /mnt/point` |                    |
| Check disk usage                | `df -h`                                       |                    |
| Check space used in dir         | `du -sh /mnt/point`                           |                    |
| Check for deleted files in use  | \`sudo lsof                                   | grep '(deleted)'\` |
| Check for hidden files          | `ls -la /mnt/point`                           |                    |

---

Let me know if you want help walking through this live on your system.
