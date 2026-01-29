  [ $(uname -m) = aarch64 ] && curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.31.0/kind-linux-arm64
  chmod +x ./kind
    sudo mv ./kind /usr/local/bin/kind
        curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl"

        sudo mv kubectl /usr/local/bin/

        
