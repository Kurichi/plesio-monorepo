module "main" {
  source = "../modules"
  common = local.common
  vpc    = local.vpc
  gke    = local.gke
}
