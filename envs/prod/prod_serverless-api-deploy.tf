# All resources relating to the use of the 'serverless-api-deploy' module for the prod environment goes here.

# The inputs for this module are mapped to the variables file or the outputs of another module as required and shouldn't have to be changed unless you know what you are doing.
module "serverless-api-deploy" {
  source                          = "../../modules/serverless-api-deploy"
  app_environment                 = var.app_environment
  app_image_tag                   = var.app_image_tag
  app_name                        = var.app_name
  cloud_run_container_port        = var.cloud_run_container_port
  cloud_run_cpu_count             = var.cloud_run_cpu_count
  cloud_run_cpu_throttling        = var.cloud_run_cpu_throttling
  cloud_run_environment_variables = var.cloud_run_environment_variables
  cloud_run_ingress_traffic_type  = var.cloud_run_ingress_traffic_type
  cloud_run_max_instance_count    = var.cloud_run_max_instance_count
  cloud_run_memory_size           = var.cloud_run_memory_size
  cloud_run_min_instance_count    = var.cloud_run_min_instance_count
  cloud_run_service_invoker       = var.cloud_run_service_invoker
  cloud_run_service_vpc_connector = var.cloud_run_service_vpc_connector
  gcp_project_id                  = var.gcp_project_id
  gcp_region                      = var.gcp_region
}


# Outputs can be referenced by other modules and can be deleted if not required.
output "endpoint_url" {
  value       = module.serverless-api-deploy.endpoint_url
  description = "The URL of the Cloud Run service."
  sensitive   = false
}

output "service_account_email" {
  value       = module.serverless-api-deploy.service_account_email
  description = "The service account being used as the service's runtime account."
  sensitive   = false
}

