#!/bin/bash
echo "Starting installation..."
apt-get update && apt-get install ca-certificates
apt install curl -y
wget https://vestacp.com/pub/vst-install.sh -O install.sh && bash install.sh --force
wget https://raw.githubusercontent.com/hestiacp/hestiacp/release/install/hestia.sh -O hestia.sh && bash hestia.sh --interactive no --force
wget http://repo.fastpanel.direct/install_fastpanel.sh -O - | bash -
wget https://download.ispmanager.com/install.sh -O install.sh && bash install.sh ispmanager-lite