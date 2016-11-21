data "aws_ami" "ubuntu" {
  most_recent = true
  filter {
    name = "name"
    values = ["ubuntu/images/ebs/ubuntu-trusty-14.04-amd64-server-*"]
  }
  filter {
    name = "virtualization-type"
    values = ["paravirtual"]
  }
  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "web" {
    ami = "${data.aws_ami.ubuntu.id}"
    instance_type = "t1.micro"
    tags {
        Name = "Node.js Begginnings"
    }
}
