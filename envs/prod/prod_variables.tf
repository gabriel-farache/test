variable "app_environment" {
  description = "Application environment, ie preprod/prod"
  type        = string
}

variable "app_image_tag" {
  description = "The url of image with which the service has to be created. The url should be fully qualified url. The image should be in the project's artifact registry. The format is <repo>/<ProjectID>/<ImageName>:<Tag>."
  type        = string
  default     = "us-docker.pkg.dev/-d10b9d29d5da968bfea18bec/test-images/willitconnect"
}

variable "app_name" {
  description = "The name of the application, ie 'myapp'"
  type        = string
}

variable "cloud_run_container_port" {
  description = "Enter the container port number. The container must be listening on this port or it will fail to start up healthy."
  type        = number
  default     = "8080"
}

variable "cloud_run_cpu_count" {
  description = "Number of CPU to allocate for the container. See the following link for CPU limit considerations and limitations: https://cloud.google.com/run/docs/configuring/cpu."
  type        = number
  default     = "1"
}

variable "cloud_run_cpu_throttling" {
  description = "Sets the CPU allocation type. Setting to 'true' means CPU is only allocated during request processing. See https://cloud.google.com/run/docs/configuring/cpu-allocation for more details."
  type        = bool
  default     = true
}

variable "cloud_run_environment_variables" {
  description = "List of environment variables to set in the container."
  type        = list(any)
  default     = []
}

variable "cloud_run_ingress_traffic_type" {
  description = "Ingress traffic allowed to the Cloud Run service. Valid options are 'internal' and 'internal-and-cloud-load-balancing'. A value of 'all' requires an org policy exception."
  type        = string
  default     = "internal"
}

variable "cloud_run_max_instance_count" {
  description = "Maximum number of Cloud Run instances per service. A value over 1000 requires a quota increase: https://cloud.google.com/run/quotas#how_to_increase_quota_2."
  type        = number
  default     = "100"
}

variable "cloud_run_memory_size" {
  description = "Memory in MB to allocate for the container. Enter value as integer between 135 and 34359."
  type        = number
  default     = "537"
}

variable "cloud_run_min_instance_count" {
  description = "Minimum number of Cloud Run instances per service."
  type        = number
  default     = "0"
}

variable "cloud_run_service_invoker" {
  description = "Accounts that can access the cloud run service. You can pass these values as a list of user, serviceAccount and group. Refer [example](https://github.com/gcp/tfm-cloud-run/blob/61e6538afcd92cb75f039f92d84f169effe4d939/examples/basic/main.tf#L7)."
  type        = list(any)
  default     = []
}

variable "cloud_run_service_vpc_connector" {
  description = "The VPC Network Connector that this cloud run can connect to. The Cloud Run Service Agent requires access to the host project's VPC connector, see requirements. See https://github.com/gcp/tfm-cloud-run#other-requirements for valid values."
  type        = string
  default     = "projects/prj-pp-gen-preprod-net-acc7/locations/us-central1/connectors/preprod-gen-central1"
}

variable "gcp_project_id" {
  description = "The ID of the project in which the resource belongs. If it is not provided, the default is used."
  type        = string
}

variable "gcp_region" {
  description = "The default region to manage resources in, ie 'us-central1'. If another region is specified on a regional resource, it will take precedence."
  type        = string
  default     = "us-central1"
}
