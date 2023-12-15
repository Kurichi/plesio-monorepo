# GKE cluster(Autopilot)
resource "google_container_cluster" "gke" {
  name = "${var.common.prefix}-cluster-${var.common.env}"

  enable_autopilot = true
  location         = var.common.region

  release_channel {
    channel = "STABLE"
  }

  networking_mode = "VPC_NATIVE"
  ip_allocation_policy {
    cluster_ipv4_cidr_block  = var.gke.cluster_cidr
    services_ipv4_cidr_block = var.gke.service_cidr
  }

  network    = google_compute_network.vpc.self_link
  subnetwork = google_compute_subnetwork.private.self_link

  maintenance_policy {
    recurring_window {
      start_time = "2006-01-01T17:00:00Z" // 2:00 AM JST
      end_time   = "2006-01-01T21:00:00Z" // 6:00 AM JST
      recurrence = "FREQ=WEEKLY;BYDAY=FR,SA,SU"
    }
  }

  deletion_protection = false

}
