# ==============================================================================
# General variables
# ==============================================================================

variable "gcp_project_id" {
  type        = string
  description = "The ID of the project in which the resource belongs. If it is not provided, the default is used."
}

variable "gcp_region" {
  type        = string
  description = "The default region to manage resources in, ie 'us-central1'. If another region is specified on a regional resource, it will take precedence."
  default     = "us-central1"
  validation {
    condition     = var.gcp_region == "us-central1" || var.gcp_region == "us" || var.gcp_region == "us-east4" || var.gcp_region == "asia-south1" || var.gcp_region == "asia-south2" || var.gcp_region == "asia" || var.gcp_region == "europe" || var.gcp_region == "europe-west2" || var.gcp_region == "europe-west3"
    error_message = "The gcp_region value must be us-central1, us, us-east4, asia-south1, asia-south2, asia, europe, europe-west2 or europe-west3."
  }
}

variable "app_name" {
  description = "The name of the application, ie 'myapp'"
  type        = string
}

variable "app_environment" {
  description = "Application environment, ie preprod/prod"
  type        = string
}

# ==============================================================================
# CloudRun variables
# ==============================================================================

variable "app_image_tag" {
  type        = string
  description = "The url of image with which the service has to be created. The url should be fully qualified url. The image should be in the project's artifact registry. The format is <repo>/<ProjectID>/<ImageName>:<Tag>."
  default     = "us-docker.pkg.dev/-d10b9d29d5da968bfea18bec/test-images/willitconnect"
}

variable "cloud_run_service_vpc_connector" {
  type        = string
  description = "The VPC Network Connector that this cloud run can connect to. The Cloud Run Service Agent requires access to the host project's VPC connector, see requirements. See https://github.com/gcp/tfm-cloud-run#other-requirements for valid values."
  default     = "projects/prj-pp-gen-preprod-net-acc7/locations/us-central1/connectors/preprod-gen-central1"
}


variable "cloud_run_service_invoker" {
  type        = list(any)
  description = "Accounts that can access the cloud run service. You can pass these values as a list of user, serviceAccount and group. Refer [example](https://github.com/gcp/tfm-cloud-run/blob/61e6538afcd92cb75f039f92d84f169effe4d939/examples/basic/main.tf#L7)."
  default     = []
}

variable "cloud_run_min_instance_count" {
  type        = number
  description = "Minimum number of Cloud Run instances per service."
  default     = "0"
}

variable "cloud_run_max_instance_count" {
  type        = number
  description = "Maximum number of Cloud Run instances per service. A value over 1000 requires a quota increase: https://cloud.google.com/run/quotas#how_to_increase_quota_2."
  default     = "100"
}

variable "cloud_run_environment_variables" {
  description = "List of environment variables to set in the container."
  type        = list(any)
  default     = []
}

variable "cloud_run_memory_size" {
  type        = number
  description = "Memory in MB to allocate for the container. Enter value as integer between 135 and 34359."
  default     = "537"
}

variable "cloud_run_cpu_count" {
  type        = number
  description = "Number of CPU to allocate for the container. See the following link for CPU limit considerations and limitations: https://cloud.google.com/run/docs/configuring/cpu."
  default     = "1"
}

variable "cloud_run_container_port" {
  type        = number
  description = "Enter the container port number. The container must be listening on this port or it will fail to start up healthy."
  default     = "8080"
}

variable "cloud_run_ingress_traffic_type" {
  type        = string
  description = "Ingress traffic allowed to the Cloud Run service. Valid options are 'internal' and 'internal-and-cloud-load-balancing'. A value of 'all' requires an org policy exception."
  default     = "internal"
}

variable "cloud_run_cpu_throttling" {
  description = "Sets the CPU allocation type. Setting to 'true' means CPU is only allocated during request processing. See https://cloud.google.com/run/docs/configuring/cpu-allocation for more details."
  type        = bool
  default     = true
}
