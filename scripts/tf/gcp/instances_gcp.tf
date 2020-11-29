provider "google" {
  project = "chattery-297109"
  region  = "asia-east2"
  zone    = "asia-east2-a"
}

resource "google_compute_network" "timemon_network" {
  name                    = "timemon-network"
  auto_create_subnetworks = "true"
}

resource "google_compute_firewall" "allow-inbound" {
  name    = "allow-inbound"
  network = google_compute_network.timemon_network.self_link

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["3530", "22"]
  }

  source_ranges = ["0.0.0.0/0"]
}


resource "google_compute_instance" "coordinator" {
  name         = "coordinator-vm"
  machine_type = "e2-micro"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  network_interface {
    # A default network is created for all GCP projects
    network = google_compute_network.timemon_network.self_link
    access_config {
    }
  }

  provisioner "remote-exec" {
    inline = [
      "wget https://github.com/arush15june/TimeMon/releases/download/v0.1/coordinator && chmod +x coordinator",
      "nohup ./coordinator &",
      "sleep 1"
    ]
    
    connection {
      type     = "ssh"
      user     = "charge"
      private_key = file("C:/Users/arush/.ssh/id_rsa")
      host     = self.network_interface.0.access_config.0.nat_ip
    }
  }

  metadata = {
    ssh-keys = "charge:${file("C:/Users/arush/.ssh/id_rsa.pub")}"
  }

  labels = { 
    name = "coordinator"
  }
}

resource "google_compute_instance" "follower" {
  count = 2
  name         = "follower-vm-${count.index+1}"
  machine_type = "e2-micro"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  network_interface {
    # A default network is created for all GCP projects
    network = google_compute_network.timemon_network.self_link
    access_config {
    }
  }

  provisioner "file" {
    source = "start-follower.sh"
    destination = "/home/charge/start-follower.sh"

    connection {
      type     = "ssh"
      user     = "charge"
      private_key = file("C:/Users/arush/.ssh/id_rsa")
      host     = self.network_interface.0.access_config.0.nat_ip
    }
  }

  provisioner "remote-exec" {
    inline = [
      "wget https://github.com/arush15june/TimeMon/releases/download/v0.1/follower && chmod +x follower",
      "echo ${self.labels.name} > follower.label",
      "echo ${google_compute_instance.coordinator.network_interface.0.network_ip}:3530 > coordinator.address",
      "chmod +x ./start-follower.sh && ./start-follower.sh",
      "sleep 1"
    ]
    
    connection {
      type     = "ssh"
      user     = "charge"
      private_key = file("C:/Users/arush/.ssh/id_rsa")
      host     = self.network_interface.0.access_config.0.nat_ip
    }
  }

  metadata = {
    ssh-keys = "charge:${file("C:/Users/arush/.ssh/id_rsa.pub")}"
  }

  labels = { 
    name = "follower-${count.index+1}"
  }
}

output "coordinator_public_ip" {
  value = google_compute_instance.coordinator.network_interface.0.access_config.0.nat_ip
}

output "coordinator_private_ip" {
  value = google_compute_instance.coordinator.network_interface.0.network_ip
}

output "follower_public_ips" {
  value = ["${google_compute_instance.follower.*.network_interface.0.access_config.0.nat_ip}"]
}

output "follower_private_ips" {
  value = ["${google_compute_instance.follower.*.network_interface.0.network_ip}"]
}
