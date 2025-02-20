# This file is used to define variable values for Terraform configurations, allowing for cleaner and more modular infrastructure code.
# Variables set here must also be defined in the corresponding Terraform configuration files located in the same directory.
# project_name is 'prj-iac-team-p'
# pipeline_account is 'sa-pipeline@-.iam.gserviceaccount.com'


app_environment                 = "prod" # Application environment, ie preprod/prod
app_image_tag                   = "us-central1-docker.pkg.dev/-123/-container-images/sln-golang-app@sha256:3c3b8eabc0c5f6c7f16694fd8224f7acee44cdd03b8d95f3dc1734169361bee2" # The url of image with which the service has to be created. The url should be fully qualified url. The image should be in the project's artifact registry. The format is <repo>/<ProjectID>/<ImageName>:<Tag>.
app_name                        = "go-kcloutie" # The name of the application, ie 'myapp'
cloud_run_service_invoker       = ["user:abaker9@.com", "user:kcloutie@.com", "user:jvandal3@.com"] # Accounts that can access the cloud run service. You can pass these values as a list of user, serviceAccount and group. Refer [example](https://github.com/gcp/tfm-cloud-run/blob/61e6538afcd92cb75f039f92d84f169effe4d939/examples/basic/main.tf#L7).
cloud_run_service_vpc_connector = "projects/prj-p-gen-priv-net-19b4/locations/us-central1/connectors/prod-priv-gen-central1-02" # The VPC Network Connector that this cloud run can connect to. The Cloud Run Service Agent requires access to the host project's VPC connector, see requirements. See https://github.com/gcp/tfm-cloud-run#other-requirements for valid values.
gcp_project_id                  = "-" # The ID of the project in which the resource belongs. If it is not provided, the default is used.
gcp_region                      = "us-central1" # The default region to manage resources in, ie 'us-central1'. If another region is specified on a regional resource, it will take precedence.
