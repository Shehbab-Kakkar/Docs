Thanks for the follow-up — if your PXE boot is **first trying IPv6**, even though your Kickstart has `--noipv6`, it means the **IPv6 behavior is happening *before* the Kickstart is applied** — i.e., during:

* **PXE firmware/bootloader stage**
* **Linux kernel/initramfs stage (before Kickstart takes effect)**

So the key is to **prevent IPv6 before the kernel loads or in the kernel command line**, not just in Kickstart.

---

## ✅ What You Need to Do

To stop the system from even **trying IPv6 before IPv4** during PXE boot and installation:

---

### ✅ 1. Add `ipv6.disable=1` to the Kernel Command Line

This is **essential**. Add this in your PXE boot config:

#### If you're using `pxelinux.cfg/default` (Legacy BIOS):

```cfg
label rocky9
  menu label Install Rocky Linux 9.4
  kernel images/rocky9/vmlinuz
  append initrd=images/rocky9/initrd.img ip=dhcp ipv6.disable=1 inst.repo=http://your.repo/rocky9.4/
```

#### If you're using GRUB (UEFI PXE boot):

```cfg
menuentry 'Install Rocky 9.4' {
    linuxefi /images/rocky9/vmlinuz ip=dhcp ipv6.disable=1 inst.repo=http://your.repo/rocky9.4/
    initrdefi /images/rocky9/initrd.img
}
```

This **disables the IPv6 kernel module**, so it won’t even attempt an IPv6 address during the early boot/initrd stage.

---

### ✅ 2. Optional: Disable IPv6 in the initrd (if necessary)

If you're still seeing IPv6 attempts even after `ipv6.disable=1`, you can **modify the initrd** to remove or blacklist the IPv6 module. This is rarely needed with `ipv6.disable=1`, but here's how to do it just in case:

#### Steps (advanced):

1. Unpack initrd:

   ```bash
   mkdir /tmp/initrd
   cd /tmp/initrd
   zcat /path/to/initrd.img | cpio -idmv
   ```

2. Blacklist IPv6:
   Add a file `/tmp/initrd/etc/modprobe.d/disable-ipv6.conf`:

   ```bash
   install ipv6 /bin/true
   blacklist ipv6
   ```

3. Repack initrd:

   ```bash
   find . | cpio --create --format='newc' | gzip > /path/to/new-initrd.img
   ```

Then point your PXE config to the new initrd.

---

### ✅ 3. UEFI/BIOS PXE Settings (if applicable)

If IPv6 is **still** being used during PXE boot **before kernel loads**, it's coming from the **firmware**. Do this:

* Go to BIOS/UEFI
* Disable "IPv6 PXE Boot" or "Network Stack IPv6" if available
* Ensure only "IPv4 PXE Boot" is enabled

---

## TL;DR — Final Setup Checklist

| Where           | Setting                                               |
| --------------- | ----------------------------------------------------- |
| PXE boot config | Add `ipv6.disable=1` to kernel command line           |
| Kickstart       | Already has `--noipv6` — good                         |
| BIOS/UEFI       | Disable IPv6 PXE boot (if you see it happening early) |
| DHCP server     | Don’t offer DHCPv6 or RA (optional but good)          |

---

