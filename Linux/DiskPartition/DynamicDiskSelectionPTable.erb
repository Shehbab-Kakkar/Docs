#Dynamic (do not remove this line)
DISKS=($(lsblk -d -n -o NAME,MODEL | grep 'Mircon' | awk '{print $1}'))
# Now $DISKS[0] and $DISKS[1] are your two Micron disk device names (e.g., sdb sdc)
cat <<EOF > /tmp/diskpart.cfg
bootloader --location=mbr --driveorder=${DISKS[0]},${DISKS[1]}
zerombr
clearpart --all --initlabel --linux

part biosboot --size=1 --ondisk=${DISKS[0]} --fstype=biosboot
part biosboot --size=1 --ondisk=${DISKS[1]} --fstype=biosboot

part /boot --size=500 --ondisk=${DISKS[0]} --fstype=mdmember
part /boot --size=500 --ondisk=${DISKS[1]} --fstype=mdmember

part pv --grow --ondisk=${DISKS[0]} --fstype=mdmember
part pv --grow --ondisk=${DISKS[1]} --fstype=mdmember

raid /boot --device=md0 --level=1 --fstype=ext4 ${DISKS[0]}2 ${DISKS[1]}2
raid pv.raid --device=md1 --level=1 --fstype=lvmpv ${DISKS[0]}3 ${DISKS[1]}3

volgroup vg_root pv.raid
logvol / --percent=100 --fstype=xfs --name=lv_root --vgname=vg_root

EOF
