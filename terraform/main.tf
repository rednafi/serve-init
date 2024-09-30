terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = ">= 1.0.0"  # You can set the version based on your needs
    }
  }
}

provider "hcloud" {
  token = var.hcloud_api_token
}

variable "hcloud_api_token" {
  description = "Hetzner Cloud API token"
  type        = string
  sensitive   = true
}

resource "hcloud_server" "hetzner_ubuntu" {
  name        = "hetzner-ubuntu"
  image       = "ubuntu-24.04"
  server_type = "cx22"
  location    = "nbg1"    # Germany Nuremberg
}

output "server_ip" {
  value = hcloud_server.hetzner_ubuntu.ipv4_address
}
