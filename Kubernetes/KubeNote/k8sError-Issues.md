1. Image Retrieval & Configuration Errors
These occur before the container actually starts, often due to manifest typos or registry issues. 
ImagePullBackOff / ErrImagePull: Kubernetes cannot download the container image.
Causes: Incorrect image name/tag, private registry authentication failure (missing imagePullSecrets), or network connectivity issues to the registry.
CreateContainerConfigError: The pod configuration is invalid.
Causes: Referencing a ConfigMap or Secret that does not exist, or using an invalid key within them.
InvalidYAML / Schema Errors: Kubernetes rejects the manifest before it is even applied.
Causes: Indentation errors, mixing tabs and spaces, or using unsupported API versions (e.g., extensions/v1beta1 in newer clusters). 
2. Pod Scheduling & Resource Errors
These occur when the cluster cannot find a suitable place to run your application. 
FailedScheduling / Pod Stuck in Pending: The pod remains "Pending" because it cannot be assigned to a node.
Causes: Insufficient CPU or memory in the cluster, or restrictive node selectors, taints, and affinities.
OOMKilled (Exit Code 137): The container exceeded its memory limit and was terminated by the system.
Causes: Application memory leaks or setting resource limits too low for the workload.
NodeNotReady: The selected node is unhealthy and cannot host pods.
Causes: Kubelet crashes, network partitions, or extreme disk/memory pressure on the host. 
3. Startup & Runtime Errors
These happen after the container has been successfully scheduled and the image pulled. 
CrashLoopBackOff: A container starts, crashes, and restarts repeatedly.
Causes: Application bugs (e.g., uncaught exceptions), missing environment variables, or incorrect startup commands.
CreateContainerError / StartError: The container fails to be created at the runtime level.
Causes: Invalid entrypoint commands, insufficient permissions for the container user, or inaccessible volume mounts.
Liveness / Readiness Probe Failures: Kubernetes detects the app is not "healthy" or "ready".
Causes: The app takes too long to start (timeout), or the health check endpoint is misconfigured/broken. 
Quick Troubleshooting Commands 
To diagnose these errors, use the following official Kubernetes tools:
kubectl describe pod <pod-name>: Best for checking events, scheduling errors, and pod lifecycle history.
kubectl logs <pod-name> --previous: Essential for seeing why a container crashed in a CrashLoopBackOff.
kubectl apply --dry-run=client -f <file>.yaml: Validates YAML syntax before deployment to prevent configuration errors
