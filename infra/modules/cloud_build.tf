locals {
  services = [
    "bff",
    "item",
    "mission",
    "tree",
    "user"
  ]
}

resource "google_cloudbuild_trigger" "main" {
  for_each = toset(local.services)

  project  = var.common.project
  location = "global"
  name     = "${each.value}-build"

  included_files = [
    "services/${each.value}/**",
  ]
  ignored_files = [
    "*_test.go",
    ".gitignore",
    "*.md",
  ]

  filename = "services/cloudbuild.yaml"
  substitutions = {
    "_SERVICE_NAME" = each.value
  }

  approval_config {
    approval_required = false
  }

  github {
    name  = "plesio-monorepo"
    owner = "Kurichi"

    push {
      branch       = "^main$"
      invert_regex = false
    }
  }
  include_build_logs = "INCLUDE_BUILD_LOGS_WITH_STATUS"

  timeouts {}
}
