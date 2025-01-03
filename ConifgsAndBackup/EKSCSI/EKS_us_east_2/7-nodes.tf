## Import the existing IAM role for the EKS Node Group
data "aws_iam_role" "nodes" {
  name = "eks-node-group-nodes"  # Correct argument name
}

## Attach policies to the imported IAM role for the Node Group
#resource "aws_iam_role_policy_attachment" "nodes_amazon_eks_worker_node_policy" {
#  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
#  role       = data.aws_iam_role.nodes.name
#}
#
#resource "aws_iam_role_policy_attachment" "nodes_amazon_eks_cni_policy" {
#  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
#  role       = data.aws_iam_role.nodes.name
#}
#
#resource "aws_iam_role_policy_attachment" "nodes_amazon_ec2_container_registry_read_only" {
#  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
#  role       = data.aws_iam_role.nodes.name
#}

# Create EKS Node Group and associate with IAM role
resource "aws_eks_node_group" "private_nodes" {
  cluster_name    = aws_eks_cluster.demo.name
  node_group_name = "private-nodes"
  node_role_arn   = data.aws_iam_role.nodes.arn  # Correctly referencing the data block

  subnet_ids = [
    aws_subnet.private_us_east_2a.id,
    aws_subnet.private_us_east_2b.id
  ]

  capacity_type  = "ON_DEMAND"
  instance_types = ["t3.medium"]

  scaling_config {
    desired_size = 1
    max_size     = 5
    min_size     = 0
  }

  update_config {
    max_unavailable = 1
  }

  labels = {
    role = "general"
  }

}
