# provider "google" {
#   credentials = file("sapient-terrain-348922-e7900ffd496f.json")
#   project     = "sapient-terrain-348922"
#   region      = "europe-north1"
#   zone        = "europe-north1-a"
# }

resource "google_cloud_run_service" "default" {
  name     = "in-memory-key-value-store"
  location = "europe-north1"

  template {
    spec {
      containers {
        image = "gcr.io/sapient-terrain-348922/in-memory-key-value-store"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.default.location
  project  = google_cloud_run_service.default.project
  service  = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
