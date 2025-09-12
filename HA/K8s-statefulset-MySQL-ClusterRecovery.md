Mysql 3 node Statefulset set. When happen if the Master mysql goes down. How we can start the mysql statefulset

When the **master MySQL pod** in a 3-node Kubernetes StatefulSet fails, replication and leadership responsibilities must be carefully managed to restore full database functionality. Out-of-the-box, Kubernetes StatefulSets do not autonomously promote a replica to master—this requires MySQL clustering logic or external automation for safe failover[1][2].

## What Happens When Master Fails

- If the master pod (e.g., mysql-0) goes down, writes to the cluster will fail as replicas are set to read-only[2].
- Reads from replicas (mysql-1, mysql-2) are still possible, but no new data gets inserted until a new master is promoted[3][2].
- MySQL itself does not automatically promote a new master in basic master-slave replication mode—external tools or manual intervention are required for failover[2].

## Restoring the MySQL StatefulSet

- To bring back a failed statefulset/mysqld-master pod:
  - The failed pod will automatically restart, and Kubernetes will attempt to schedule it using the same persistent storage due to the PVC (Persistent Volume Claim) association[3].
- If the persistent storage is intact and there is no corruption, the pod should recover as the master and replication resumes as before[3][4].
- If the pod or its storage is lost/damaged, you must manually promote a replica:
  - Update a replica pod's configuration to become the new master.
  - Point remaining replicas to the new master (by changing their replication targets).
  - This can be automated via scripts, operators, or HA solutions, but is not handled by StatefulSet natively[2][5].

## Typical Manual Recovery Steps

- If master pod and its PVC are lost/corrupted:
  - Promote a healthy replica to master using `CHANGE MASTER TO` commands or orchestration.
  - Adjust StatefulSet or orchestration configuration to recognize the new master role[5].
  - Optionally update Service endpoints to redirect writes to the new master.
- If just the pod (not PVC) went down, simply let Kubernetes restart it. Data should persist, and replication will restore states automatically[3][4].

## Enhancing Recovery and Availability

- Use MySQL clustering solutions (e.g. Group Replication, Percona XtraDB Cluster, or external MySQL operators) for automatic failover, seamless promotion, and self-healing capabilities[6][2].
- For basic StatefulSet deployments, manual failover scripts or human intervention are required if the master’s PVC is unrecoverable.

***

In summary: the default MySQL StatefulSet does not handle automatic master failover. If the master pod fails but its PVC is intact, Kubernetes will restart it. For master promotion after data loss/corruption, manual steps or additional tooling/automation are required to promote a replica and update the topology[1][5][2].

Citations:
[1] In a StatefulSet replication, what happens if master goes ... https://www.reddit.com/r/kubernetes/comments/uc4hd7/in_a_statefulset_replication_what_happens_if/
[2] Introduction to Failover for MySQL Replication - the 101 Blog https://severalnines.com/blog/introduction-failover-mysql-replication-101-blog/
[3] Run a Replicated Stateful Application https://kubernetes.io/docs/tasks/run-application/run-replicated-stateful-application/
[4] I have deleted a StatefulSet on accident and need to re- ... https://stackoverflow.com/questions/75384304/i-have-deleted-a-statefulset-on-accident-and-need-to-re-create-it
[5] Database Disaster Recovery in Kubernetes https://www.linkedin.com/pulse/database-disaster-recovery-kubernetes-krishan-shamod
[6] Best practices for 3-node HA server setup with automatic ... https://forums.percona.com/t/best-practices-for-3-node-ha-server-setup-with-automatic-failover-proxysql-cluster-xtradb-cluster-8-0/32700
[7] Stateful set not shifting to a new node - kubernetes https://stackoverflow.com/questions/53783815/stateful-set-not-shifting-to-a-new-node
[8] MySQL Master-Slave Replication When Master is Down https://stackoverflow.com/questions/29411655/mysql-master-slave-replication-when-master-is-down
[9] docker - Kubernetes: Error when creating a StatefulSet with ... https://stackoverflow.com/questions/53189490/kubernetes-error-when-creating-a-statefulset-with-a-mysql-container
[10] MariaDB statefulset with one replica fails to restart due to ... https://www.reddit.com/r/kubernetes/comments/1e99p1m/mariadb_statefulset_with_one_replica_fails_to/
[11] Stateful Workload Upgrade Patterns - Azure https://learn.microsoft.com/en-us/azure/aks/stateful-workload-upgrades
[12] Guide to Kubernetes StatefulSets: Deploying Stateful Apps https://www.plural.sh/blog/kubernetes-statefulset-guide/
[13] Understanding StatefulSets in Kubernetes https://portworx.com/knowledge-hub/understanding-statefulsets-in-kubernetes/
[14] How to run a MySQL pod on OCP using OCS and ... https://www.redhat.com/en/blog/how-run-mysql-pod-ocp-using-ocs-and-statefulsets
[15] Troubleshoot and Fix Kubernetes CrashLoopBackoff Status https://komodor.com/learn/how-to-fix-crashloopbackoff-kubernetes-error/
[16] StatefulSet Basics https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/
[17] Scalable MySQL Server Setup with StatefulSets https://www.mirantis.com/blog/how-to-use-statefulsets-and-create-a-scalable-mysql-server-on-kubernetes/
[18] Kubernetes Application Management: Stateful Services https://www.alibabacloud.com/blog/kubernetes-application-management-stateful-services_595087
[19] Deploying MySQL on Kubernetes: A Comprehensive Guide https://www.plural.sh/blog/kubernetes-mysql-deployment-guide/
[20] Kubernetes StatefulSet - Examples & Best Practices https://www.vcluster.com/blog/kubernetes-statefulset-examples-and-best-practices

---

Give me the sequence of commands to bring the cluster up.

To bring up a MySQL 3-node StatefulSet cluster on Kubernetes after a failure or outage, follow this detailed command sequence. These steps assume that PersistentVolumeClaims (PVCs) are healthy and the StatefulSet and Service YAMLs are already deployed.

## Sequence of Commands

### 1. Check Pod and Storage Status
```shell
kubectl get pods -l app=mysql
kubectl get pvc -l app=mysql
```
Review which pods are not running and ensure PVCs are still bound and healthy[1][2].

### 2. Restart the StatefulSet or Individual Pods
If the entire StatefulSet needs to be restarted:
```shell
kubectl rollout restart statefulset <statefulset-name>
```
If only one or more pods have failed (for example, the master):
```shell
kubectl delete pod <pod-name>
```
Kubernetes will automatically recreate the deleted pod using the same persistent volume[1][2].

### 3. Monitor Pod Recovery
```shell
kubectl get pods -l app=mysql -w
```
Wait until all pods show `STATUS` as `Running` and `READY` is at full for each pod.

### 4. Check MySQL Process Health and Cluster Topology
Check database startup and replication from inside the master and replica pods:
```shell
kubectl logs <pod-name>
kubectl exec -it mysql-0 -- mysql -u root -p -e "SHOW MASTER STATUS\G"
kubectl exec -it mysql-1 -- mysql -u root -p -e "SHOW SLAVE STATUS\G"
kubectl exec -it mysql-2 -- mysql -u root -p -e "SHOW SLAVE STATUS\G"
```
This confirms that the master is running and replicas are connected[1].

### 5. (If Needed) Manually Promote Replica
If the master (`mysql-0`) and its PVC cannot be recovered, promote a replica to be master:
```shell
kubectl exec -it mysql-1 -- mysql -u root -p
# Inside MySQL shell:
STOP SLAVE;
RESET SLAVE ALL;
SET GLOBAL read_only = OFF;
```
Update other replicas to replicate from the new master using the `CHANGE MASTER TO ...` command[1].

***

These commands cover most operational recovery scenarios for a StatefulSet-managed MySQL cluster in Kubernetes, assuming standard replication setups and healthy storage[1][2][3].

Citations:
[1] Run a Replicated Stateful Application https://kubernetes.io/docs/tasks/run-application/run-replicated-stateful-application/
[2] How to gracefully restart the cluster without downtime? https://forums.percona.com/t/how-to-gracefully-restart-the-cluster-without-downtime/23142
[3] Deploy a stateful MySQL cluster on GKE | Kubernetes Engine https://cloud.google.com/kubernetes-engine/docs/tutorials/stateful-workloads/mysql
