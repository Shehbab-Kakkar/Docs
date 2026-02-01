Got it! You want to **uninstall both Helm releases**: `arc` in `arc-systems` namespace and `arc-runner-set` in `arc-runners` namespace. You can do this with `helm uninstall`. Here's the exact commands:

```bash
helm uninstall arc -n arc-systems
helm uninstall arc-runner-set -n arc-runners
```

After running these, you can verify they’re gone:

```bash
helm list -A
```

```
[root@rocky ~]# helm list -A
NAME          	NAMESPACE  	REVISION	UPDATED                                	STATUS  	CHART                                 	APP VERSION
arc           	arc-systems	1       	2026-02-01 17:38:41.859558888 +0530 IST	deployed	gha-runner-scale-set-controller-0.13.1	0.13.1     
arc-runner-set	arc-runners	1       	2026-02-01 17:39:54.491781114 +0530 IST	deployed	gha-runner-scale-set-0.13.1           	0.13.1     
[root@rocky ~]# helm uninstall arc -n arc-systems
helm uninstall arc-runner-set -n arc-runners
release "arc" uninstalled
release "arc-runner-set" uninstalled
[root@rocky ~]# helm list -A
NAME	NAMESPACE	REVISION	UPDATED	STATUS	CHART	APP VERSION
[root@rocky ~]# helm registry logout ghcr.io
Removing login credentials for ghcr.io
[root@rocky ~]# 

```
