In Linux, the **MMU (Memory Management Unit)** itself is **not done by Linux**—it is **hardware**, built into the **CPU**.
Linux’s job is to **control and use** the MMU.

Here’s how the responsibility is split 👇

---

## Who does what with the MMU?

### 1. **MMU = CPU hardware**

* The MMU is a **hardware component inside the processor** (x86, ARM, RISC-V, etc.).
* It does things like:

  * Translate **virtual addresses → physical addresses**
  * Enforce **memory protection** (user vs kernel, read/write/execute)
  * Cache translations using the **TLB (Translation Lookaside Buffer)**

Linux does **not** implement these mechanisms in software.

---

### 2. **Linux kernel manages the MMU**

Linux tells the MMU **how to behave**.

The kernel:

* Creates and manages **page tables**
* Decides which process can access which memory
* Switches memory mappings during **context switches**
* Handles **page faults**
* Sets up memory protection rules

Key kernel parts involved:

* `mm/` directory (memory management code)
* Architecture-specific code (e.g., `arch/x86/mm/`, `arch/arm/mm/`)

---

### 3. **Per-process virtual memory**

* Each process gets its **own virtual address space**
* Linux switches page tables when switching processes
* The MMU makes it look like each process has its own memory

---

## Simple summary

* ✅ **MMU** → **CPU hardware**
* 🧠 **Memory policy & setup** → **Linux kernel**
* 🔁 **Address translation & protection** → **MMU (using kernel-provided tables)**

---

