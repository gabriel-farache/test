resource "google_service_account_iam_member" "pd303-wif-binding-iac-team" {
  service_account_id = "projects/-123/serviceAccounts/sa-pipeline@-123.iam.gserviceaccount.com"
  role               = "roles/iam.workloadIdentityUser"
  member             = "principal://iam.googleapis.com/projects/123/locations/global/workloadIdentityPools/pd303-5g4zs/subject/system:serviceaccount:iac-team:pipeline"
}