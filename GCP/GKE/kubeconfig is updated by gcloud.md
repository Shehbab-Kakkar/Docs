Where kubeconfig is updated by gcloud
By default, gcloud writes cluster credentials to the Kubernetes config file at:
Linux/macOS: ~/.kube/config
Windows (PowerShell/CMD): %USERPROFILE%.kube\config
Windows (Git Bash/Cygwin): ~/.kube/config
The exact file path can be overridden with the KUBECONFIG environment variable. If KUBECONFIG points to one or more files (colon-separated on Unix, semicolon-separated on Windows), kubectl will read/merge from those, and gcloud will update the first writable file in that list.
The command you ran:
gcloud container clusters get-credentials my-cluster --zone us-east1a --project my-project
will add/merge a context, user, and cluster entry for “my-cluster” into the kubeconfig and set the current-context to that cluster (unless you specify --internal-ip or other flags that alter behavior).
Tips:
To see which kubeconfig kubectl is using: kubectl config view --minify --raw
To print the current context: kubectl config current-context
To check KUBECONFIG:
Linux/macOS: echo $KUBECONFIG
Windows (PowerShell): $env:KUBECONFIG
To specify a different config file for a one-off command: KUBECONFIG=/path/to/altconfig kubectl get pods (use $env:KUBECONFIG in PowerShell)
