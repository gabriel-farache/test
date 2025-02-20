resource "google_service_account_iam_member" "pd303-wif-binding-iac-team" {
  service_account_id = "projects/-a87865c181efbb2660ba1a1f/serviceAccounts/sa-pipeline@-a87865c181efbb2660ba1a1f.iam.gserviceaccount.com"
  role               = "roles/iam.workloadIdentityUser"
  member             = "principal://iam.googleapis.com/projects/123/locations/global/workloadIdentityPools/pd303-5g4zs/subject/system:serviceaccount:iac-team:pipeline"
}