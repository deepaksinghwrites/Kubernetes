# Import the existing IAM role for the EKS Cluster
data "aws_iam_role" "demo" {
  name = "eks-cluster-demo"  # Correct argument name
}


## Attach AmazonEKSClusterPolicy to the imported IAM role
#resource "aws_iam_role_policy_attachment" "demo_amazon_eks_cluster_policy" {
#  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
#  role       = data.aws_iam_role.demo.name  # Correct reference to the imported IAM role
#}

# Create the EKS Cluster and associate the IAM role
resource "aws_eks_cluster" "demo" {
  name     = "demo"
  role_arn = data.aws_iam_role.demo.arn  # Correct reference to the imported IAM role

  vpc_config {
    subnet_ids = [
      aws_subnet.private_us_east_2a.id,
      aws_subnet.private_us_east_2b.id,
      aws_subnet.public_us_east_2a.id,
      aws_subnet.public_us_east_2b.id
    ]
  }
}
