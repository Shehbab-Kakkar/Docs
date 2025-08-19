# 🆘 Recover a Corrupt VM in Google Cloud Platform (GCP) Using `gce-rescue`

If your GCP Compute Engine VM is unbootable or corrupt (e.g., boot, disk, or OS issues), you can use the `gce-rescue` tool to recover it.  
**gce-rescue** mounts the broken disk to a rescue VM so you can fix configuration, remove problematic files, reset passwords, etc.

---

## 🚑 Steps to Recover a Corrupt GCP VM Using `gce-rescue`

### 1. **Install or Get `gce-rescue`**

- The `gce-rescue` tool is available at: https://github.com/GoogleCloudPlatform/compute-image-tools/tree/master/tools/gce-rescue
- You can run it using Docker or install Go and build from source.

#### **Run via Docker** (recommended):

```sh
docker run --rm -ti \
  -v ~/.config/gcloud:/root/.config/gcloud \
  gcr.io/compute-image-tools/gce-rescue -h
```

---

### 2. **Find Your Broken VM's Details**

- **INSTANCE_NAME**: The broken VM's name (e.g., `my-vm`)
- **ZONE**: The zone of your VM (e.g., `us-central1-a`)
- **PROJECT**: Your GCP project ID

---

### 3. **Run `gce-rescue` to Rescue the VM**

This command will:

- Stop the broken VM
- Detach its boot disk
- Create a rescue VM and attach the broken disk to it
- SSH you into the rescue VM

```sh
docker run --rm -ti \
  -v ~/.config/gcloud:/root/.config/gcloud \
  gcr.io/compute-image-tools/gce-rescue \
  --project=<PROJECT> \
  --zone=<ZONE> \
  --instance=<INSTANCE_NAME>
```

**Replace:**
- `<PROJECT>` with your project ID
- `<ZONE>` with your VM's zone
- `<INSTANCE_NAME>` with your VM name

---

### 4. **Fix Issues Inside the Rescue VM**

- Once SSH’d into the rescue VM, your broken disk will be mounted (typically at `/mnt/broken-disk`).
- Now you can:
  - Edit files
  - Remove bad software/configs
  - Reset passwords
  - Recover data
- When done, exit the SSH session.

---

### 5. **Restore the Original VM**

The `gce-rescue` tool will prompt you to restore the disk to the original VM.  
You can also do it manually if needed.

---

## 📝 Example Full Command

```sh
docker run --rm -ti \
  -v ~/.config/gcloud:/root/.config/gcloud \
  gcr.io/compute-image-tools/gce-rescue \
  --project=my-gcp-project \
  --zone=us-central1-a \
  --instance=my-corrupt-vm
```

---

## 📚 References

- [gce-rescue GitHub](https://github.com/GoogleCloudPlatform/compute-image-tools/tree/master/tools/gce-rescue)
- [GCP Docs: Troubleshooting boot issues](https://cloud.google.com/compute/docs/troubleshooting/troubleshooting-vm-boot)
- [GCP Community: VM Rescue Guide](https://cloud.google.com/blog/products/compute/introducing-gce-rescue-tool)

---

**Tip:** Always take a snapshot of your VM disk before performing recovery operations!
