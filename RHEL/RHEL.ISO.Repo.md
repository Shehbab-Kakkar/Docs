# RHEL 9.4 ISO PXE Boot & HTTP Repo Setup on Apache2 (Ubuntu)

## 1. Prerequisites

- Ubuntu server with Apache2 installed
- RHEL 9.4 ISO at `/home/ubuntu/isos/rhel-9.4-x86_64-dvd.iso`
- Sufficient disk space

## 2. Install Apache2

```bash
sudo apt update
sudo apt install apache2 -y
```

## 3. Mount the RHEL ISO

Let’s mount the ISO to `/var/www/html/rhel9.4`:

```bash
sudo mkdir -p /var/www/html/rhel9.4
sudo mount -o loop /home/ubuntu/isos/rhel-9.4-x86_64-dvd.iso /var/www/html/rhel9.4
```

> To mount automatically after reboot, add this line to `/etc/fstab`:
>
> ```
> /home/ubuntu/isos/rhel-9.4-x86_64-dvd.iso /var/www/html/rhel9.4 iso9660 loop,ro 0 0
> ```

## 4. Set Permissions

```bash
sudo chown -R www-data:www-data /var/www/html/rhel9.4
sudo chmod -R 755 /var/www/html/rhel9.4
```

## 5. Verify HTTP Access

Open a browser and visit:  
`http://<your-server-ip>/rhel9.4/`  
You should see the repo contents (e.g., `Packages/`, `repodata/`, etc.).

## 6. Configure PXE Boot (TFTP, DHCP, syslinux)

### Install TFTP and syslinux

```bash
sudo apt install tftpd-hpa syslinux pxelinux
```

### Copy Boot Files

Create a tftp root:

```bash
sudo mkdir -p /var/lib/tftpboot
```

Copy PXE bootloader and kernel/initrd from the ISO:

```bash
sudo cp /usr/lib/PXELINUX/pxelinux.0 /var/lib/tftpboot/
sudo cp /var/www/html/rhel9.4/isolinux/vmlinuz /var/lib/tftpboot/
sudo cp /var/www/html/rhel9.4/isolinux/initrd.img /var/lib/tftpboot/
sudo mkdir -p /var/lib/tftpboot/pxelinux.cfg
```

### Create PXE Boot Menu

Create `/var/lib/tftpboot/pxelinux.cfg/default`:

```
DEFAULT rhel9
LABEL rhel9
  KERNEL vmlinuz
  APPEND initrd=initrd.img inst.repo=http://<your-server-ip>/rhel9.4
```

Replace `<your-server-ip>` with your Ubuntu server’s IP.

### Configure TFTP

Edit `/etc/default/tftpd-hpa`:

```
TFTP_USERNAME="tftp"
TFTP_DIRECTORY="/var/lib/tftpboot"
TFTP_ADDRESS=":69"
TFTP_OPTIONS="--secure"
```
Restart TFTP:
```bash
sudo systemctl restart tftpd-hpa
```

### Configure DHCP (optional)

If you control DHCP, add:

```
next-server <your-server-ip>;
filename "pxelinux.0";
```

## 7. Firewall (if needed)

Allow HTTP and TFTP:

```bash
sudo ufw allow 80/tcp
sudo ufw allow 69/udp
```

## 8. Test

- PXE boot a client.  
- It should get PXE menu, boot kernel/initrd, and use the HTTP repo for installation.

---

## Troubleshooting

- Ensure the ISO is always mounted.
- Check Apache logs: `/var/log/apache2/`
- Check TFTP logs: `/var/log/syslog`

---

## References

- [Red Hat Documentation - Installing using PXE](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/9/html/performing_an_advanced_rhel_installation/assembly_installing-using-pxe-performing-an-advanced-rhel-installation)
- [Ubuntu Apache2](https://ubuntu.com/server/docs/web-servers-apache)
