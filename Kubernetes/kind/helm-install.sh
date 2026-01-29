  68  curl -fsSL https://get.helm.sh/helm-v3.12.0-linux-arm64.tar.gz -o helm-v3.12.0-linux-arm64.tar.gz
   69  ll
   70  tar -zxvf helm-v3.12.0-linux-arm64.tar.gz
   71  sudo mv linux-arm64/helm /usr/local/bin/helm
   72  sudo chmod +x /usr/local/bin/helm
   73  helm version
