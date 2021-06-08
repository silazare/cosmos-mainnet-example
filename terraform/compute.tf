resource "random_string" "suffix" {
  length  = 4
  special = false
  upper   = false
}

resource "google_compute_instance" "cosmos_node" {
  name                      = "cosmos-node-${random_string.suffix.result}"
  description               = "Cosmos node"
  machine_type              = var.machine_type
  zone                      = var.location
  allow_stopping_for_update = true

  service_account {
    email = var.service_account_email
    // To allow full access to all Cloud APIs, need to set anyway even if we define SA email and not scope
    scopes = ["cloud-platform"]
  }

  boot_disk {
    initialize_params {
      image = var.disk_image
      size  = var.disk_size
      type  = var.disk_type
    }
  }

  network_interface {
    network    = google_compute_network.cosmos_vpc_network.self_link
    subnetwork = google_compute_subnetwork.cosmos_vpc_subnet.self_link

    access_config {
      nat_ip = google_compute_address.cosmos_ephemeral_ip.address
    }
  }

  metadata = {
    sshKeys = "${var.ssh_user}:${file(var.public_key_path)}"
  }

  tags = ["cosmos"]

  provisioner "remote-exec" {
    inline = ["echo 'SSH connected!'"]

    connection {
      host        = google_compute_instance.cosmos_node.network_interface[0].access_config[0].nat_ip
      type        = "ssh"
      user        = var.ssh_user
      private_key = file(var.private_key_path)
    }
  }
}
