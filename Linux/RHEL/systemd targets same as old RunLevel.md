Perfect! On **RHEL 9**, the old SysV `runlevel` commands are mostly replaced by **systemd targets**, but the concepts are the same. I’ll break it down step by step — including **checking TTYs**.

---

## 1️⃣ Understanding runlevels in RHEL 9

RHEL 9 uses **systemd**, so “runlevels” are now **targets**:

| Classic Runlevel | systemd Target      | Description                    |
| ---------------- | ------------------- | ------------------------------ |
| 0                | `poweroff.target`   | Shutdown                       |
| 1                | `rescue.target`     | Single-user mode               |
| 3                | `multi-user.target` | CLI (non-graphical) multi-user |
| 5                | `graphical.target`  | GUI multi-user                 |
| 6                | `reboot.target`     | Reboot                         |

> So when you “switch runlevels”, you’re actually **switching targets**.

---

## 2️⃣ Check current runlevel / target

### Using `runlevel`

```bash
runlevel
```

Output example:

```
N 5
```

* `N` = previous runlevel unknown
* `5` = current runlevel (graphical)

### Using `systemctl`

```bash
systemctl get-default
```

Output example:

```
graphical.target
```

---

## 3️⃣ Change the runlevel (target)

### Temporarily (without reboot)

```bash
sudo systemctl isolate multi-user.target
```

* Switches to CLI mode (runlevel 3) **immediately**
* All GUI processes stopped

### Permanently (next boot)

```bash
sudo systemctl set-default multi-user.target
```

* On next reboot, system will boot into CLI mode (runlevel 3)

To switch back to GUI:

```bash
sudo systemctl isolate graphical.target
sudo systemctl set-default graphical.target
```

---

## 4️⃣ Check active TTYs

TTY = terminals. Useful if you want to see **which consoles are active**.

### List all TTYs

```bash
who
```

Example output:

```
root     pts/0        2026-01-24 09:00 (:0)
alice    pts/1        2026-01-24 09:05 (192.168.1.100)
```

* `pts/X` = virtual terminals (SSH or GUI)
* `ttyX` = physical console

### Check all systemd TTY services

```bash
systemctl list-units | grep tty
```

Output example:

```
getty@tty1.service loaded active running Getty on tty1
getty@tty2.service loaded active running Getty on tty2
```

* These `getty` services manage **login prompts** on TTYs

---

## 5️⃣ Switch to a different TTY (keyboard only)

* Press `Ctrl + Alt + F1..F6`
* `F1` to `F6` → virtual terminals
* `F7` → usually the GUI (if running graphical.target)

---

## 6️⃣ Summary Workflow

| Task                        | Command                                   |           |
| --------------------------- | ----------------------------------------- | --------- |
| Check runlevel              | `runlevel` or `systemctl get-default`     |           |
| Switch runlevel temporarily | `systemctl isolate multi-user.target`     |           |
| Switch runlevel permanently | `systemctl set-default multi-user.target` |           |
| List active TTYs            | `who`                                     |           |
| Check TTY services          | `systemctl list-units                     | grep tty` |
| Switch TTY                  | `Ctrl + Alt + F1-F6`                      |           |

---

If you want, I can make a **quick cheat-sheet for RHEL 9 runlevels + TTY commands** that you can keep for sysadmin tasks — it’s super handy in production.

Do you want me to make that?
