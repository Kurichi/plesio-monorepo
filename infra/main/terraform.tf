provider "google" {
  project = local.common.project
  region  = local.common.region
}

terraform {
  backend "gcs" {
    bucket = "tf-state-plesio"
    prefix = "terraform/state"
  }
}
