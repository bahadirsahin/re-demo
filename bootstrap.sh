#/bin/bash
# update os
sudo apt update -y
sudo apt upgrade -y

# install make
sudo apt install -y make

# install docker & docker compose
sudo mkdir -p /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update -y
sudo apt install -y containerd.io docker-ce docker-ce-cli docker-compose-plugin
sudo usermod -aG docker $USER

# check versions
make -v
docker -v
docker compose version
