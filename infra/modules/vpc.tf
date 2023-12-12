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

resource "google_compute_global_address" "static" {
  project = var.common.project
  name    = "${var.common.prefix}-static-ip-${var.common.env}"
}

resource "google_dns_managed_zone" "dns" {
  name        = "${var.common.prefix}-dns-${var.common.env}"
  dns_name    = "${var.common.domain}."
  description = "Managed by Terraform"
}

resource "google_dns_record_set" "dns" {
  name         = "${var.common.domain}."
  type         = "A"
  ttl          = 300
  managed_zone = google_dns_managed_zone.dns.name
  rrdatas      = [google_compute_global_address.static.address]
}
