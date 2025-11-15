#!/bin/bash

for i in {0..9}; do
  dev="/dev/nvme${i}n1"
  part="${dev}p1"
  mount_point="/mnt/disk${i}"

  # Create partition if it doesn't exist
  if ! ls ${part} >/dev/null 2>&1; then
    echo "Creating GPT and partition on ${dev}..."
    parted -s ${dev} mklabel gpt
    parted -s ${dev} mkpart primary xfs 0% 100%
    sleep 2
  fi

  # Format partition as XFS
  mkfs.xfs -f ${part}

  # Create mount point
  mkdir -p ${mount_point}

  # Get UUID for fstab entry
  uuid=$(blkid -s UUID -o value ${part})

  # Add to /etc/fstab if not already present
  if ! grep -q "${uuid}" /etc/fstab; then
    echo "UUID=${uuid} ${mount_point} xfs defaults 0 0" | tee -a /etc/fstab
  fi

  # Mount disk
  mount ${mount_point}
done
