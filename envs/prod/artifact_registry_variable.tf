
variable "shared_project_id" {
  type        = string
  description = "The ID of the project in which all shared resources like artifact registry and secrets will be created"
}
variable "artifact_registry_id" {
  type        = string
  description = "The ID of the Repository"
  default     = "tfm-artifact-registry"
}

variable "artifact_registry_description" {
  type        = string
  description = "repo description"
  default     = "TFM Artifact Registry"
}

variable "artifact_registry_gcp_region" {
  type        = string
  description = "The region in which the shared artifact registry exists"
  default = "us-central1"
  validation {
    condition     = var.artifact_registry_gcp_region == "us-central1" || var.artifact_registry_gcp_region == "us" || var.artifact_registry_gcp_region == "us-east4" || var.artifact_registry_gcp_region == "asia-south1" || var.artifact_registry_gcp_region == "asia-south2" || var.artifact_registry_gcp_region == "asia" || var.artifact_registry_gcp_region == "europe" || var.artifact_registry_gcp_region == "europe-west2" || var.artifact_registry_gcp_region == "europe-west3"
    error_message = "The artifact_registry_gcp_region value must be us-central1, us, us-east4, asia-south1, asia-south2, asia, europe, europe-west2 or europe-west3."
  }
}

variable "format" {
  type        = string
  description = "Container formats for the artifact registry"
  default     = "DOCKER"
}

variable "labels" {
  description = "a set of key/value label pairs for registries."
  type        = map(string)
  default     = null
}