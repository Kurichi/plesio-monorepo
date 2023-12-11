resource "google_compute_network" "vpc" {
  project                 = var.common.project
  name                    = "${var.common.prefix}-vpc-${var.common.env}"
  auto_create_subnetworks = false
  routing_mode            = "REGIONAL"
}

# resource "google_compute_subnetwork" "public" {
#   project       = var.common.project
#   name          = "${var.common.prefix}-public-subnet-${var.common.env}"
#   region        = var.common.region
#   ip_cidr_range = var.vpc.public_subnet_cidr
#   network       = google_compute_network.vpc.id
# }

resource "google_compute_subnetwork" "private" {
  project       = var.common.project
  name          = "${var.common.prefix}-private-subnet-${var.common.env}"
  region        = var.common.region
  ip_cidr_range = var.vpc.private_subnet_cidr
  network       = google_compute_network.vpc.id
}
