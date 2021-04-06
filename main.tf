resource "aws_s3_bucket" "self" {
  bucket = var.bucket_name

  versioning {
    enabled = true
  }
}
