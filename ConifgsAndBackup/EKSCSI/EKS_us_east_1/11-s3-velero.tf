# S3 Bucket for Velero Backups
resource "aws_s3_bucket" "velero" {
  bucket = var.bucket  # Use the bucket variable here
  tags = {
    Name        = "Velero Backup Bucket"
    Environment = "Production"
  }
}


# S3 Bucket Versioning
resource "aws_s3_bucket_versioning" "velero_versioning" {
  bucket = aws_s3_bucket.velero.id

  versioning_configuration {
    status = "Enabled"
  }
}
