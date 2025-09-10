Amazon EC2 G6f instances are the latest GPU-based offerings from AWS, designed to deliver flexible, cost-efficient GPU acceleration, especially for users who do not always need a full GPU for their workloads.

Key Features
GPU Partitioning: G6f instances allow you to provision as little as 1/8 of an NVIDIA L4 Tensor Core GPU (3GB GPU memory), right-sizing resources and decreasing costs for workloads that don't require the power of a full GPU. You can select from half, quarter, and one-eighth GPU instance sizes.
Backed by Latest Hardware: Paired with third-generation AMD EPYC processors, G6f supports up to 16 vCPUs and 12GB GPU memory, delivering a strong balance between CPU and GPU resources for graphics-heavy and ML workloads.
Strong Use Cases: Ideal for:
Machine learning inference and experimentation.
Graphics-intensive applications (e.g., rendering, simulations, CAD/CAE).
Media & Entertainment virtual workstations.
Gaming, video encoding, and application streaming.
Environments where GPU needs are small or highly variable (e.g., training, education, shared/tenant platforms).
Performance & Regions:
Up to 10Gbps networking for smaller instances, higher on larger types.
Storage includes fast NVMe SSDs, optimized for high-throughput and low-latency local data access.
Available in major AWS regions, including US East/West, Europe, APAC, Canada, and South America.
Availability & Pricing:
Purchase options: On-Demand, Spot, or Savings Plans.
Supported by NVIDIA GRID drivers for fractional instances; eligible for fast desktop access via Amazon DCV remote desktop from anywhere.
Example pricing: g6f.large (2 vCPUs, 8GB RAM, 1/8 GPU) from ~$0.22/hour or ~$147/month depending on region and plan.
Benefits compared to previous generations:
2x the inference and graphics performance of G4dn.
Very granular scaling and rightsizing for GPU resources to better match application need — better cost efficiency for teams doing development, research, or multi-tenant cloud environments.
Example G6f Instance Types
Instance Type	vCPUs	RAM (GiB)	GPU (L4) Fraction	GPU Memory (GiB)	Main Use
g6f.large	2	8	1/8 (0.125)	2.8–3	ML inference, light graphics
g6f.xlarge	4	16	1/4 (0.25)	6	Medium graphics, dev/test
g6f.2xlarge	8	32	1/2 (0.5)	12	Heavier graphics, rendering
g6f.4xlarge	16	64	1 (full)	24	Demanding ML/graphics
Summary:
G6f instances are a cutting-edge AWS GPU solution offering unmatched flexibility and cost efficiency thanks to fractional GPU allocation, ideal for a wide range of graphics and ML workloads — especially where workloads don’t fully utilize an entire GPU.
