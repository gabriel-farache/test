
resource "google_service_account" "cldrun_sa" {
  account_id   = "${var.app_name}-cldrun-sa"                 # service account id
  display_name = "${var.app_name} cloud run service account" # service account name
  project      = var.gcp_project_id                          # project id
}

resource "google_cloud_run_service" "default" {
  name     = var.app_name
  location = var.gcp_region
  project  = var.gcp_project_id

  metadata {
    annotations = {
      "run.googleapis.com/ingress" = "internal"
    }
  }

  template {
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = var.cloud_run_max_instance_count
        "autoscaling.knative.dev/minScale" = var.cloud_run_min_instance_count
        #
        "run.googleapis.com/cpu-throttling"       = true
        "run.googleapis.com/vpc-access-connector" = var.cloud_run_service_vpc_connector
        "run.googleapis.com/vpc-access-egress"    = "all-traffic"
      }
    }
    spec {
      service_account_name = google_service_account.cldrun_sa.email
      containers {
        name  = var.app_name
        image = var.app_image_tag
        resources {
          limits = {
            cpu    = "1000m"
            memory = "500M"
          }
        }
        startup_probe {
          initial_delay_seconds = 1
          timeout_seconds       = 1
          period_seconds        = 3
          failure_threshold     = 3

          http_get {
            path = "/healthz"
          }
        }
        liveness_probe {
          http_get {
            path = "/livez"
          }
        }

        args = [
          "run",
          "server",
          "--config-file-path",
          "/mnt/app-config/config.yaml",
        ]
        volume_mounts {
          name       = "${var.app_name}-configuration"
          mount_path = "/mnt/app-config"
        }
        # Add the environment variable here
        env {
          name  = "APP_NAME"
          value = var.app_name
        }
      }

      volumes {
        name = "${var.app_name}-configuration"
        secret {
          secret_name = google_secret_manager_secret.app_config_secret.secret_id
          items {
            key  = "latest"
            path = "config.yaml"
          }
        }
      }
    }
  }
  lifecycle {
    ignore_changes = [metadata[0].annotations["run.googleapis.com/operation-id"]]
  }
  autogenerate_revision_name = true
}


resource "google_cloud_run_service_iam_member" "invoker" {
  project  = var.gcp_project_id
  service  = google_cloud_run_service.default.name
  location = google_cloud_run_service.default.location
  count    = length(var.cloud_run_service_invoker)
  role     = "roles/run.invoker"
  member   = var.cloud_run_service_invoker[count.index]
}


data "template_file" "default" {
  template = file("${path.module}/applicationConfig.tftpl")
  vars = {
    # 
    gcp_project_id                  = "",
    app_connection_string_secret_id = ""
    # 
  }
}


resource "google_secret_manager_secret" "app_config_secret" {
  secret_id = "${var.app_name}-application-config"
  project   = var.gcp_project_id
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "app_config_secret" {
  secret      = google_secret_manager_secret.app_config_secret.id
  secret_data = data.template_file.default.rendered
}

resource "google_secret_manager_secret_iam_binding" "app_config_secret" {
  project   = var.gcp_project_id
  secret_id = google_secret_manager_secret.app_config_secret.secret_id
  role      = "roles/secretmanager.secretAccessor"
  members   = ["serviceAccount:${google_service_account.cldrun_sa.email}"]
}