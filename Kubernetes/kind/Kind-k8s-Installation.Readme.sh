  [ $(uname -m) = aarch64 ] && curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.31.0/kind-linux-arm64
  chmod +x ./kind
    sudo mv ./kind /usr/local/bin/kind
        curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl"

        sudo mv kubectl /usr/local/bin/
37  # Install dependencies
   38  sudo dnf install -y dnf-plugins-core
   39  # Add Docker repo
   40  sudo dnf config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
   41  # Install Docker
   42  sudo dnf install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
   43  # Start Docker
   44  sudo systemctl enable --now docker
   45  # Add your user to docker group (so you can run without sudo)
   46  sudo usermod -aG docker $USER

docker --version

kind create cluster --name mycluster --config kind-3nodes-500mb.yaml 

kubectl config current-context
kubectl config get-contexts
kubectl config use-context kind-mycluster
