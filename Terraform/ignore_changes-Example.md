# 🚀 Terraform EKS Node Group: `scaling_config[0].desired_size` and Lifecycle Explained

This README explains how scaling and lifecycle management works for EKS node groups in Terraform, focusing on the meaning and best practices around `scaling_config[0].desired_size` and `lifecycle { ignore_changes = [...] }`.

---

## 💡 Original Example Code

```hcl
resource "aws_iam_role" "nodes" {
  name = "${local.env}-${local.eks_name}-eks-nodes"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      }
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "amazon_eks_worker_node_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
  role       = aws_iam_role.nodes.name
}

resource "aws_iam_role_policy_attachment" "amazon_eks_cni_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
  role       = aws_iam_role.nodes.name
}

resource "aws_iam_role_policy_attachment" "amazon_ec2_container_registry_read_only" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
  role       = aws_iam_role.nodes.name
}

resource "aws_eks_node_group" "general" {
  cluster_name    = aws_eks_cluster.eks.name
  version         = local.eks_version
  node_group_name = "general"
  node_role_arn   = aws_iam_role.nodes.arn

  subnet_ids = [
    aws_subnet.private_zone1.id,
    aws_subnet.private_zone2.id
  ]

  capacity_type  = "ON_DEMAND"
  instance_types = ["t3.large"]

  scaling_config {
    desired_size = 1
    max_size     = 10
    min_size     = 0
  }

  update_config {
    max_unavailable = 1
  }

  labels = {
    role = "general"
  }

  depends_on = [
    aws_iam_role_policy_attachment.amazon_eks_worker_node_policy,
    aws_iam_role_policy_attachment.amazon_eks_cni_policy,
    aws_iam_role_policy_attachment.amazon_ec2_container_registry_read_only,
  ]

  # Allow external changes without Terraform plan difference
  lifecycle {
    ignore_changes = [scaling_config[0].desired_size]
  }
}
```

---

## 📦 What is `scaling_config` in EKS Node Groups?

The `scaling_config` block in your `aws_eks_node_group` resource controls how AWS manages the size of your node group (i.e., the number of EC2 instances running as Kubernetes nodes).

**Example:**
```hcl
scaling_config {
  desired_size = 1
  max_size     = 10
  min_size     = 0
}
```

- **desired_size**: How many nodes you want to run right now
- **max_size**: Maximum number of nodes allowed
- **min_size**: Minimum number of nodes allowed

---

## 🧠 What does `scaling_config[0].desired_size` Mean?

In Terraform, complex/nested blocks (like `scaling_config`) are treated as **lists** of objects, even if you only define one.

- `scaling_config[0]` refers to the first (and usually only) scaling_config block.
- `scaling_config[0].desired_size` accesses the `desired_size` inside that block.

**Terraform's internal view:**
```hcl
scaling_config = [
  {
    desired_size = 1
    max_size     = 10
    min_size     = 0
  }
]
```

---

## 🔁 Why Use `lifecycle { ignore_changes = [scaling_config[0].desired_size] }`?

**Purpose:**  
Prevents Terraform from treating manual/automatic scaling changes as "drift" that needs to be reverted.

**Why ignore `desired_size` changes?**

- **Cluster Autoscaler** or manual AWS Console changes often adjust `desired_size` to scale up/down nodes.
- If you don’t ignore changes, any manual/automatic scaling will show up as a diff in `terraform plan` and be reverted on `terraform apply`.
- By ignoring `desired_size`, you let Terraform manage only the _initial_ node count and focus on `min_size`/`max_size` for guardrails.

**Example:**
```hcl
lifecycle {
  ignore_changes = [scaling_config[0].desired_size]
}
```
This tells Terraform:
- "If someone changes the number of nodes (desired_size) outside Terraform, don't worry about it."
- Only manage `min_size` and `max_size` strictly.

---

## ⚖️ Should You Ignore the Whole `scaling_config` Block?

You can, but then **Terraform will stop managing** even `min_size` and `max_size`, which may not be what you want.

```hcl
lifecycle {
  ignore_changes = [scaling_config]
}
```
- Use this only if you want all scaling properties to be managed externally (not typical for most teams).

---

## 📝 Summary Table

| Expression                       | Meaning                                         |
|-----------------------------------|-------------------------------------------------|
| `scaling_config`                  | Block containing scaling settings               |
| `scaling_config[0]`               | The first (only) scaling config object/block    |
| `scaling_config[0].desired_size`  | Desired node count in the node group            |
| `lifecycle.ignore_changes`        | Tells Terraform to ignore changes to attribute  |
| `[scaling_config[0].desired_size]`| Ignore only `desired_size` changes              |
| `[scaling_config]`                | Ignore all scaling config changes               |

---

## 👍 Best Practices

- **Use `ignore_changes = [scaling_config[0].desired_size]`** if you want to allow external/manual scaling without Terraform interference.
- **Let Terraform manage `min_size` and `max_size`** to ensure your cluster doesn’t scale beyond safe limits.
- **Use autoscalers or AWS Console for day-to-day scaling**; use Terraform for infrastructure as code and guardrails.

---

## 📚 References

- [Terraform aws_eks_node_group Docs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eks_node_group)
- [Terraform: ignore_changes](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes)
- [EKS Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler)

---

**Tip:**  
This pattern is widely used in production EKS environments to balance infrastructure-as-code with dynamic Kubernetes scaling!

---
🔧 What is update_config in EKS Node Group?

The update_config block controls how AWS EKS updates the node group when a change is made — such as changing instance types, scaling parameters, or AMI versions.

Here’s your usage:

update_config {
  max_unavailable = 1
}

⚙️ What does max_unavailable = 1 mean?

This setting tells AWS:

“During an update to the node group, take down at most 1 node at a time.”

This is part of rolling updates — where EKS:

Terminates an old EC2 node (based on the old config)

Brings up a new one (with the updated config)

Waits for it to be healthy before moving on to the next

By limiting max_unavailable to 1, you:

Ensure high availability (only 1 node is offline during update)

Avoid overwhelming the cluster with too many changes at once

📘 Full update_config options (as of AWS/Terraform)
Attribute	Description
max_unavailable	(Optional) Number of nodes that can be unavailable during an update.
max_unavailable_percentage	(Optional) Percentage of nodes that can be unavailable during an update (alternative to above).

You can only use one: max_unavailable or max_unavailable_percentage.
