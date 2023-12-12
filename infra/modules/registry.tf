resource "google_artifact_registry_repository" "docker" {
  project       = var.common.project
  location      = var.common.region
  repository_id = "${var.common.prefix}-docker-${var.common.env}"
  format        = "DOCKER"
}
