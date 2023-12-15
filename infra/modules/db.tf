resource "google_sql_database_instance" "main" {
  name             = "${var.common.prefix}-db-${var.common.env}"
  database_version = "POSTGRES_15"
  region           = var.common.region
  settings {
    tier = var.db.tier
    ip_configuration {
      ipv4_enabled = true
      # private_network = google_compute_network.vpc.self_link
    }
  }
}

resource "google_sql_database" "main" {
  name     = "main"
  instance = google_sql_database_instance.main.name
}

resource "google_sql_user" "main" {
  for_each = var.db.users
  name     = each.value.name
  instance = google_sql_database_instance.main.name
  password = each.value.password
}

