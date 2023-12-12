# 踏み台 VM が使用するサービスアカウント
resource "google_service_account" "bastion" {
  project = var.common.project

  account_id   = "sa-${var.common.prefix}-${var.common.env}"
  display_name = "Service Account for bastion-${var.common.prefix}-${var.common.env}"
}

# サービスアカウントにロールを付与
resource "google_project_iam_member" "bastion" {
  project = var.common.project

  role   = "roles/container.developer" # Kubernetes Engine デベロッパー ロール
  member = "serviceAccount:${google_service_account.bastion.email}"
}

# 踏み台 VM
resource "google_compute_instance" "bastion" {
  project = var.common.project

  name         = "bastion-${var.common.prefix}-${var.common.env}"
  machine_type = var.bastion.machine_type
  zone         = "${var.common.region}-a"

  tags = ["ssh"] # ネットワークタグ

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11" # OS イメージ
    }
  }

  network_interface {
    subnetwork_project = var.common.project

    network    = google_compute_network.vpc.name        # VPC
    subnetwork = google_compute_subnetwork.private.name # サブネット

    access_config {} # パブリック IP を付与
  }

  metadata = {
    enable-oslogin = "true" # OS Login を有効化
  }

  # 起動スクリプトで kubectl と GKE のプラグインをインストール
  metadata_startup_script = <<EOF
#!/bin/bash
sudo apt update
sudo apt install kubectl
sudo apt install google-cloud-sdk-gke-gcloud-auth-plugin
EOF

  service_account {
    email  = google_service_account.bastion.email
    scopes = ["cloud-platform"]
  }
}

resource "google_compute_firewall" "ssh" {
  project = var.common.project

  name    = "vpc-${var.common.prefix}-${var.common.env}-ssh-allow"
  network = google_compute_network.vpc.name

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = [
    var.bastion.ssh_sourcerange
  ]
  target_tags = ["ssh"]
}
