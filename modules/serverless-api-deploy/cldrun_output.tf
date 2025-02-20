# ==============================================================================
# CloudRun Outputs
# ==============================================================================

output "endpoint_url" {
  value       = google_cloud_run_service.default.status[0].url
  description = "The URL of the Cloud Run service."
}

output "service_account_email" {
  value       = google_service_account.cldrun_sa.email
  description = "The service account being used as the service's runtime account."
}
