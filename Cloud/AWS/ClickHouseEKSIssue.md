This issue is a real-world example of hidden infrastructure bottlenecks that can affect cloud-native applications, specifically when running distributed databases like ClickHouse on Kubernetes clusters in AWS.
What Happened?
The user experienced repeated crashes of ClickHouse pods and sometimes the entire nodes, despite there being no signs of problems in query workload, logs, compute, or memory.
Scaling up the node group (adding more EC2 instances) didn't help—new nodes filled up and crashed just as quickly.
The initial suspicion was that some resource limit (CPU, RAM, disk, etc.) was the cause. However, there was a subtler infrastructure limit at play.
Root Causes and Fixes:
Pod IP Address Exhaustion:
In AWS EKS, the number of pods that each node can run is capped by the number of available IP addresses per Elastic Network Interface (ENI).
For certain EC2 instance types, like t3.medium, the default was only 17 pod IPs per node.
Once that limit was reached, new pods couldn’t be scheduled, manifesting as instability and crashes.
Fix 1: Enabling prefix delegation in the AWS VPC CNI plugin increased the pod IP quota per node, solving the immediate IP exhaustion.
ENI Throughput Saturation:
Despite resolving the IP limit, pods continued to crash.
Investigation with Pixie (an observability tool) revealed network saturation: the ENI’s throughput (measured in bandwidth and packets per second) was being maxed out.
The t3.medium instance type has limited network and ENI capacity, insufficient for the high volume of traffic between distributed ClickHouse nodes.
Fix 2: Switching to larger EC2 instances with better network and ENI performance allowed pods to stabilize, scale, and communicate without exceeding underlying networking limits.
Key Lessons:
Cloud-Native Scaling Limits: Kubernetes and cloud services introduce new types of scaling pain points, like pod IP allocation and network throughput, which often aren’t obvious through traditional DB/server monitoring.
Horizontal and Vertical Scaling: Scaling isn’t just about adding more nodes (horizontal)—sometimes you need “vertical scaling” (upgrading node specs) or changes in network config.
Observability is Critical: Tools like Pixie help reveal problems that common monitoring and logs miss, highlighting the need for deep observability in modern systems.
Resilience Engineering: True reliability work goes “beneath the stack,” into the details of cloud networking, kernel limits, and platform integration.
Takeaway: Even when your application code, configurations, and query patterns are flawless, infrastructure limitations—especially in networking—can cause unexplained failures. Understanding and tuning these “hidden” layers is essential for resilient, scalable cloud architectures.
