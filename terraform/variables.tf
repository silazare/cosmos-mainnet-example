# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# These variables are expected to be passed in by the operator.
# ---------------------------------------------------------------------------------------------------------------------

variable "project" {
  description = "The project ID where all resources will be launched."
  type        = string
}

variable "location" {
  description = "The location (region or zone) of the GKE cluster."
  type        = string
}

variable "region" {
  description = "The region for the network. If the cluster is regional, this must be the same region. Otherwise, it should be the region of the zone."
  type        = string
}

variable "machine_type" {
  description = "Machine type"
  type        = string
  default     = "e2-standard-2"
}

variable "disk_image" {
  description = "Disk image"
  type        = string
  default     = "cosmos-base-v1"
  #default     = "ubuntu-os-cloud/ubuntu-2004-lts"
}

variable "disk_size" {
  description = "Disk size"
  type        = number
  default     = 100
}

variable "disk_type" {
  description = "Disk type"
  type        = string
  default     = "pd-ssd"
}

variable "source_range" {
  description = "Source CIDR ranges"
  type        = string
  default     = "0.0.0.0/0"
}

variable "public_key_path" {
  description = "Path to the public key used for ssh access"
  type        = string
}

variable "private_key_path" {
  description = "Path to the private key used for ssh access"
  type        = string
}

variable "ssh_user" {
  description = "Username for SSH access"
  type        = string
}

variable "service_account_email" {
  description = "Service Account email"
  type        = string
}
