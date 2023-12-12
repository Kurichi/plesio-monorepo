module "main" {
  source  = "../modules"
  common  = local.common
  vpc     = local.vpc
  gke     = local.gke
  bastion = local.bastion
}

output "command_to_connect_cluster" {
  value = "\n$ gcloud container clusters get-credentials ${module.main.cluster_name} --region ${module.main.cluster_location} --project ${module.main.cluster_project}\n"
}
