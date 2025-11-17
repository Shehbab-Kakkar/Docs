for i in `cat hosts`; do sshpass -p 'xxxx!' ssh -o StrictHostKeyChecking=no root@$i "echo 'root:xxxx!' | chpasswd"; done
