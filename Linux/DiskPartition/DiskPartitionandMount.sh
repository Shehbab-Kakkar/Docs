#!/bin/bash
set -e

DEVICE="/dev/nvme9n1"
PARTITION="${DEVICE}p1"
MOUNT_POINT="/mnt/disk8"

# 1. Create a single GPT partition using parted
parted --script $DEVICE mklabel gpt
parted --script $DEVICE mkpart primary xfs 0% 100%

# 2. Wait for partition node
sleep 2

# 3. Format the partition with XFS
mkfs.xfs $PARTITION

# 4. Create mount point
mkdir -p $MOUNT_POINT

# 5. Get UUID and update /etc/fstab for permanent mount
UUID=$(blkid -s UUID -o value $PARTITION)
echo "UUID=${UUID} $MOUNT_POINT xfs defaults,noatime 0 0" >> /etc/fstab

# 6. Mount it
mount $MOUNT_POINT

echo "Done! Full disk partitioned, formatted as XFS, and mounted at $MOUNT_POINT"
