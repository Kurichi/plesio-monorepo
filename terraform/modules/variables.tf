variable "common" {
  type = object({
    project = string
    prefix  = string
    region  = string
    env     = string
    domain  = string
  })
}

variable "vpc" {
  type = object({
    # public_subnet_cidr  = string
    private_subnet_cidr = string
  })
}

variable "gke" {
  type = object({
    cluster_cidr = string
    service_cidr = string
    master_cidr  = string
  })
}

variable "bastion" {
  type = object({
    machine_type    = string
    ssh_sourcerange = string
  })

}
