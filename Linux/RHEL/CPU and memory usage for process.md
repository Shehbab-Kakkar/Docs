how Linux (RHEL 9 in your case) accounts **CPU and memory usage** for a process. Let’s break it down clearly.

---

## **1. CPU usage by a process**

CPU usage is **time slices the process spends running on the CPU**.

* Linux tracks CPU in **jiffies** (small ticks of the kernel timer).
* Tools like `top`, `htop`, and `ps` calculate **%CPU** as the percentage of CPU time used in the last measurement interval.

**Key formulas/concepts:**

* **%CPU** ≈ `(process_cpu_time / elapsed_time) * 100`
* **process_cpu_time** = user time + system time
* If you have multiple cores, CPU usage can go over 100% (e.g., 4 cores → max 400%).

**Command examples:**

```bash
# Show top CPU processes
top

# Show CPU % for all processes
ps -eo pid,comm,%cpu,%mem --sort=-%cpu
```

---

## **2. Memory usage by a process**

Linux memory tracking is more nuanced. A process can “use” different types of memory:

| Metric                          | What it means                                           | Command             |
| ------------------------------- | ------------------------------------------------------- | ------------------- |
| **RSS** (Resident Set Size)     | Actual RAM the process occupies                         | `ps -o pid,rss,cmd` |
| **VSZ** (Virtual Memory Size)   | Total virtual address space used                        | `ps -o pid,vsz,cmd` |
| **PSS** (Proportional Set Size) | RAM shared with other processes, counted proportionally | `smem` tool         |
| **Shared / Private**            | Shared libraries vs private memory                      | `pmap <pid>`        |

Example with `ps`:

```bash
ps -eo pid,comm,rss,vsz,%mem --sort=-%mem
```

* `%MEM` = `(RSS / total_physical_memory) * 100`

**Key point:**

* Zombies **don’t use CPU or RSS memory** because they are dead; only a tiny process table entry exists.

---

## **3. How a process “gets” CPU & memory**

* **CPU:** Scheduled by the kernel. If runnable, it is put in the CPU queue. Linux uses **CFS (Completely Fair Scheduler)** in RHEL 9.
* **Memory:** Allocated via `malloc()` or mapped files.

  * Kernel tracks **virtual memory**, **resident pages**, **swap usage**, etc.
* Process memory may include:

  * Code/text
  * Data (heap)
  * Stack
  * Shared libraries

---

If you want, I can make a **visual example showing a live process on RHEL 9 and exactly how CPU and memory are allocated**—it’s surprisingly easy to track.

Do you want me to do that?
