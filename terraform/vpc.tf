resource "google_compute_network" "cosmos_vpc_network" {
  name                    = "cosmos-network"
  description             = "Cosmos VPC"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "cosmos_vpc_subnet" {
  name          = "cosmos-subnetwork"
  ip_cidr_range = "192.168.100.0/24"
  region        = var.region
  network       = google_compute_network.cosmos_vpc_network.self_link
}

resource "google_compute_firewall" "cosmos_firewall_ssh" {
  name        = "cosmos-allow-ssh"
  description = "Allow SSH for defined CIDR range"
  network     = google_compute_network.cosmos_vpc_network.self_link

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["cosmos"]
}

resource "google_compute_firewall" "cosmos_firewall_http" {
  name        = "cosmos-allow-tcp"
  description = "Allow TCP for defined CIDR range"
  network     = google_compute_network.cosmos_vpc_network.self_link

  allow {
    protocol = "tcp"
    ports    = ["26656", "26657"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["cosmos"]
}


resource "google_compute_address" "cosmos_ephemeral_ip" {
  name        = "cosmos-ephemeral-ip"
  description = "Cosmos Ephemeral IP (Static External IP)"
}
