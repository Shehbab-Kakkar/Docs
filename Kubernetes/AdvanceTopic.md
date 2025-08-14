Kyverno
Cilicum
Karpenter to right-size EC2 instances and consolidate nodes.

➟ They leveraged spot instances to lower compute costs.

➟ They created VPAs automatically using Kyverno for efficient resource scaling.

➟ They set minimal resource requests to handle usage spikes without overprovisioning.

➟ They used kube-downscaler to scale down replicas during off-hours.

➟ They used KEDA for custom metric-based scaling, allowing HPA and VPA to work together.

➟ They applied Kyverno policies to manage PDBs, allowing underutilized nodes to be removed.
