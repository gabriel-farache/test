# The following determine where Terraform's state file is stored.

#
terraform {
  backend "gcs" {
    bucket = "bkt-tfstate-iac-team-s" # the backend is stored in a bucket typically in the common project
    prefix = "terraform/iac-team/sandbox/go-kcloutie"
  }
}
# 