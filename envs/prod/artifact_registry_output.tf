output "artifact_registry_repo_id" {
  value       = google_artifact_registry_repository.artifact-registry-repo.id
  description = "Artifact registry repository id"
}