resource "aws_eks_addon" "csi_driver" {
  cluster_name = aws_eks_cluster.demo.name
  addon_name   = "aws-ebs-csi-driver"

  service_account_role_arn = data.aws_iam_role.eks_ebs_csi_driver.arn  # Reference to the existing IAM role

  # Additional configurations for the addon
}