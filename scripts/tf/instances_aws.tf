provider "aws" {
  region = "ap-south-1"
}

resource "aws_security_group" "allow_all" {
  name        = "allow_all"
  description = "Allow all inbound traffic"

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "coordinator" {
  ami = "ami-0a4a70bd98c6d6441"
  instance_type = "t2.micro"
  key_name = "EC2_ROOT"
  
  vpc_security_group_ids = [aws_security_group.allow_all.id]
  
  provisioner "remote-exec" {
    inline = [
      "git clone https://github.com/arush15june/TimeMon.git",
    ]
    
    connection {
      type     = "ssh"
      user     = "ubuntu"
      private_key = file("E:/Keystore/SSHKeys/EC2_ROOT.pem")
      host     = self.public_ip
    }
  }

  tags = { 
    name = "coordinator"
  }
}

resource "aws_instance" "follower" {
  count = 2
  
  ami = "ami-0a4a70bd98c6d6441"
  instance_type = "t2.micro"
  key_name = "EC2_ROOT"
  
  vpc_security_group_ids = [aws_security_group.allow_all.id]
  
  provisioner "remote-exec" {
    inline = [
      "git clone https://github.com/arush15june/TimeMon.git",
    ]
    
    connection {
      type     = "ssh"
      user     = "ubuntu"
      private_key = file("E:/Keystore/SSHKeys/EC2_ROOT.pem")
      host     = self.public_ip
    }
  }
  
  tags = { 
    name = "follower-${count.index}"
  }
}

output "coordinator_ip" {
  value = aws_instance.coordinator.public_ip
}

output "follower_ips" {
  value = ["${aws_instance.follower.*.public_ip}"]
}

