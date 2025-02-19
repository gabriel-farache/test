resource "google_artifact_registry_repository" "artifact-registry-repo" {
  location      = var.artifact_registry_gcp_region
  repository_id = var.artifact_registry_id
  description   = var.artifact_registry_description
  format        = "DOCKER"
  project = var.shared_project_id
  labels = {
    "resource-type" = "application-env-shared"
    "env"           = "prod"
  }
}


# 
# 
resource "google_artifact_registry_repository_iam_member" "registry_write_member_0" {
  project    = google_artifact_registry_repository.artifact-registry-repo.project
  location   = google_artifact_registry_repository.artifact-registry-repo.location
  repository = google_artifact_registry_repository.artifact-registry-repo.repository_id
  # role   = "organizations/99808423033/roles/.artifact.registry.provisioner"
  role = "roles/artifactregistry.writer"
  member = "serviceAccount:sa-pipeline@-123.iam.gserviceaccount.com"
}
# 
# 
resource "google_artifact_registry_repository_iam_member" "registry_write_member_1" {
  project    = google_artifact_registry_repository.artifact-registry-repo.project
  location   = google_artifact_registry_repository.artifact-registry-repo.location
  repository = google_artifact_registry_repository.artifact-registry-repo.repository_id
  # role   = "organizations/99808423033/roles/.artifact.registry.provisioner"
  role = "roles/artifactregistry.writer"
  member = "serviceAccount:sa-pipeline@-a87865c181efbb2660ba1a1f.iam.gserviceaccount.com"
}
# 
# 
resource "google_artifact_registry_repository_iam_member" "registry_write_member_2" {
  project    = google_artifact_registry_repository.artifact-registry-repo.project
  location   = google_artifact_registry_repository.artifact-registry-repo.location
  repository = google_artifact_registry_repository.artifact-registry-repo.repository_id
  # role   = "organizations/99808423033/roles/.artifact.registry.provisioner"
  role = "roles/artifactregistry.writer"
  member = "serviceAccount:sa-pipeline@-.iam.gserviceaccount.com"
}
# 
# 

# 
# 
resource "google_artifact_registry_repository_iam_member" "read_member_0" {
  project    = google_artifact_registry_repository.artifact-registry-repo.project
  location   = google_artifact_registry_repository.artifact-registry-repo.location
  repository = google_artifact_registry_repository.artifact-registry-repo.repository_id
  # role       = google_project_iam_custom_role.artifact_registry_iam_only_role.id
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:service-121471442119@serverless-robot-prod.iam.gserviceaccount.com"
}
# 
# 
resource "google_artifact_registry_repository_iam_member" "read_member_1" {
  project    = google_artifact_registry_repository.artifact-registry-repo.project
  location   = google_artifact_registry_repository.artifact-registry-repo.location
  repository = google_artifact_registry_repository.artifact-registry-repo.repository_id
  # role       = google_project_iam_custom_role.artifact_registry_iam_only_role.id
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:service-838836965106@serverless-robot-prod.iam.gserviceaccount.com"
}
# 
# 
resource "google_artifact_registry_repository_iam_member" "read_member_2" {
  project    = google_artifact_registry_repository.artifact-registry-repo.project
  location   = google_artifact_registry_repository.artifact-registry-repo.location
  repository = google_artifact_registry_repository.artifact-registry-repo.repository_id
  # role       = google_project_iam_custom_role.artifact_registry_iam_only_role.id
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:service-984675688163@serverless-robot-prod.iam.gserviceaccount.com"
}
# 
# 