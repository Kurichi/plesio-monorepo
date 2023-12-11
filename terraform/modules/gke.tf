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

  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = true
    master_ipv4_cidr_block  = var.gke.master_cidr

    master_global_access_config {
      enabled = false
    }
  }

  master_authorized_networks_config {
  }

  maintenance_policy {
    recurring_window {
      start_time = "2006-01-02T02:00:00+09:00"
      end_time   = "2006-01-02T06:00:00+09:00"
      recurrence = "FREQ=WEEKLY;BYDAY=FR,SA,SU"
    }
  }

  deletion_protection = false

}
