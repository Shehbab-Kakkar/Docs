---

## Why an in-place upgrade isn’t supported

1. **Official Policy**

   * Rocky’s documentation clearly states that they *do not* support upgrades between major versions. ([Rocky Linux Documentation][2])
   * Their “Upgrade Policy” (on the Wiki) also says major version upgrades are *not generally supported*, though an experimental tool (ELevate) exists. ([Rocky Linux Wiki][3])
2. **Rocky 10 Release Notes**

   * The Rocky Linux 10 release announcement itself recommends a **fresh install** from Rocky 9 or 8. ([Rocky Linux][4])
   * There are also architectural changes: e.g., **x86-64-v2** CPUs are *no longer supported* in Rocky 10 — you need at least x86-64-v3. ([Rocky Linux][4])

---

## What *is* the recommended way to move to Rocky 10

Since in-place isn’t officially supported, the *safest* and *recommended* path is:

1. **Backup Everything**

   * Your data
   * Configuration files (e.g. `/etc/`)
   * List of installed packages
   * Users, custom scripts, databases, etc.

2. **Record Current State**

   * Get a list of installed packages:

     ```bash
     sudo dnf list installed | awk 'NR>1 {print $1}' | sort -u > installed.txt
     ```

     ([Rocky Linux Documentation][1])
   * Get a list of enabled repositories:

     ```bash
     ls -al /etc/yum.repos.d/ > repolist.txt
     ```

     ([Rocky Linux Documentation][1])
   * Export user accounts (optional, but useful):

     ```bash
     sudo getent passwd > userid.txt
     ```

     ([Rocky Linux Documentation][1])

3. **Install Rocky Linux 10 Fresh**

   * Download the Rocky 10 ISO from the official site. ([Rocky Linux][4])
   * Install on the machine (or in a new partition / VM, depending on your setup).
   * After installation, run `dnf upgrade` to catch up to the latest packages. ([Rocky Linux Documentation][1])

4. **Restore**

   * Re-create users based on `userid.txt`. ([Rocky Linux Documentation][1])
   * Re-add repositories you were using (non-Rocky ones) as per `repolist.txt`. ([Rocky Linux Documentation][1])
   * Reinstall packages from `installed.txt`:

     ```bash
     sudo dnf -y install $(cat installed.txt)
     ```

     ([Rocky Linux Documentation][1])
   * Copy back configuration files, data, scripts, and other custom stuff.

---

## Risks of Trying an Unsupported In-Place Upgrade

* Because it's **unsupported**, if things go wrong, you may not be able to revert. ([vInfrastructure][5])
* Dependency issues: packages from 9 may not have exact equivalents in 10.
* Some reported CPU incompatibility (if your CPU is too old) when people tried upgrading. ([Reddit][6])
* RPM database format may change, requiring manual fixes. ([vInfrastructure][5])

---

## Experimental / Community Methods

* **ELevate**: A tool mentioned in the Rocky Wiki for version upgrades. But again, *not officially supported*. ([Rocky Linux Wiki][3])
* **Manual “hacky” method**: Some users on Reddit report trying things like:

  ```bash
  sudo dnf -y install rocky-release-10 ... rocky-gpg-keys-10  
  sudo dnf -y --releasever=10 --allowerasing distro-sync  
  rpm --rebuilddb  
  ```

  But this is **risky** and not guaranteed. ([Reddit][6])

---

## My Recommendation (Given Your Situation)

* If it's a **production system**: Do a **fresh install**. Backup, reinstall, and restore. This is the safest.
* If it's a **test / homelab** machine: You *could* experiment with the unsupported upgrade, but **test thoroughly** after and be ready to revert / reinstall if something breaks.
* Verify whether your **hardware** (especially CPU) supports the new requirements of Rocky 10. ([Rocky Linux][4])

---

If you like, I can check **current community tools** (as of now) for in-place upgrades from Rocky 9 → 10 (maybe there is some better-tested method), and walk you through it. Do you want me to do that?

[1]: https://docs.rockylinux.org/10/guides/update_versions/?utm_source=chatgpt.com "Rocky supported version upgrades - Documentation"
[2]: https://docs.rockylinux.org/10/zh/guides/update_versions/?utm_source=chatgpt.com "Rocky supported version upgrades - Documentation"
[3]: https://wiki.rockylinux.org/rocky/version/?utm_source=chatgpt.com "Rocky Linux Release and Version Guide - Rocky Linux Wiki"
[4]: https://rockylinux.org/news/rocky-linux-10-0-ga-release?utm_source=chatgpt.com "Rocky Linux 10.0 Available Now - Rocky Linux"
[5]: https://vinfrastructure.it/2025/07/upgrade-rocky-linux-from-9-6-to-10/?utm_source=chatgpt.com "Upgrade Rocky Linux from 9.6 to 10 - vInfrastructure Blog"
[6]: https://www.reddit.com//r/RockyLinux/comments/1l9gxdh?utm_source=chatgpt.com "RHEL 10 Upgrade from 9.6?"
