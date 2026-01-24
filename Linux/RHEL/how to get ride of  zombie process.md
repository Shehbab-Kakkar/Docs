Got it! Let’s go step by step—zombie processes in Linux are a special case because they **aren’t actually running**, so they don’t consume CPU or memory in the usual sense. But I’ll explain everything clearly and show you how to find them and deal with them.

---

### **1. What is a zombie process?**

* A **zombie process** is a process that has finished execution but still has an entry in the process table.
* This happens when the **parent process hasn’t read its exit status** (with `wait()` or `waitpid()`).
* Key point: zombies don’t use **CPU** or **memory**, only a **small amount of kernel memory** (process table entry).

---

### **2. How to find zombie processes**

You can use several commands:

**Option A: `ps`**

```bash
ps aux | grep Z
```

* The **`STAT` column** shows `Z` for zombies.

**Option B: `top`**

```bash
top
```

* Press `Shift + M` to sort by memory.
* Zombies show as `<defunct>` in the COMMAND column.

**Option C: `htop`** (if installed)

* Look for **`<defunct>`** processes.

---

### **3. How to clear zombie processes**

Zombies **cannot be killed with `kill -9`**, because they are already dead. They exist only because their **parent hasn’t acknowledged them**.

**Options:**

**A. If parent is alive:**
Send `SIGCHLD` to the parent so it can reap the zombie:

```bash
kill -s SIGCHLD <parent_pid>
```

**B. If parent is stuck or dead:**

1. Kill the parent process:

```bash
kill -9 <parent_pid>
```

2. When the parent dies, zombies are **adopted by `init` (PID 1)**, which automatically reaps them.

**C. Automatic cleanup:**

* Usually, zombies are cleaned up when the parent process exits.
* If you see lots of zombies frequently, it’s a **bug in the parent program**.

---

### **4. How much resources do zombies use?**

* **CPU**: 0 (they’re dead)
* **Memory**: tiny, just the process table entry (~1-2 KB per zombie)
* **Disk/Storage**: 0

You can get a rough count:

```bash
ps -eo stat,pid,ppid,cmd | grep Z | wc -l
```

* This gives the **number of zombies**.
* Memory usage is negligible, but if you want to see all processes’ memory:

```bash
ps aux --sort=-%mem
```

Zombies will appear as `<defunct>` but with **0 RSS**.

---

💡 **TL;DR**

* Zombies are harmless in small numbers.
* Cannot be killed directly.
* Kill or fix the parent process to clean them up.
* Don’t worry about CPU/memory—they barely take any.

---

