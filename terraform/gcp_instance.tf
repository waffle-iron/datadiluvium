resource "google_compute_instance" "sample-nodejs-1" {
  name = "sample-nodejs-1"
  machine_type = "f1-micro"
  tags = [
    "http-server",
    "https-server",
    "dev-ssh"]

  zone = "us-west1-a"

  disk {
    image = "ubuntu-1404-trusty-v20160809a"
    size = 16
  }

  network_interface {
    subnetwork = "developer-space"
    access_config {
    }
  }

  service_account {
    scopes = [
      "userinfo-email",
      "compute-ro",
      "storage-ro"]
  }
}