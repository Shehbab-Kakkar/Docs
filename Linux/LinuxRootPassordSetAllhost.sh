for i in `cat hosts`; do sshpass -p 'xxxx!' ssh -o StrictHostKeyChecking=no root@$i "echo 'root:xxxx!' | chpasswd"; done
for i in `cat hosts`; do sshpass -p 'xxx' ssh -o StrictHostKeyChecking=no root@$i 'hostname;cat /etc/redhat-release;' ;done
