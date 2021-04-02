variable "region" {
  type        = string
  default     = "eu-west-1"
  description = "Region where to deploy"
}

variable "bucket_name" {
  type        = string
  description = "Unique name for the bucket"
  default = "paloth-test-bucket"
}
