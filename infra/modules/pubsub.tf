locals {
  topics = [
    "give_item",
    "act_to_tree",
  ]
}

resource "google_pubsub_topic" "main" {
  for_each = toset(local.topics)
  name     = "${each.value}-topic"
}

resource "google_pubsub_subscription" "terraform-subscription" {
  for_each = toset(local.topics)
  name     = "${each.value}-subscription"
  topic    = google_pubsub_topic.main[each.value].name

  message_retention_duration = "604800s"
  retain_acked_messages      = true

  ack_deadline_seconds = 10

  enable_message_ordering = false
}
