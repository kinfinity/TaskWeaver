#!/bin/bash
#
# Setup Worker Nodes

set -euxo pipefail

# Join & setup hostname
/bin/bash /vagrant/configs/join.sh --v=5

sudo -i -u vagrant bash << EOF
whoami
mkdir -p /home/vagrant/.taskweaver
sudo cp -i /vagrant/configs/config /home/vagrant/.taskweaver/
sudo chown 1000:1000 /home/vagrant/.taskweaver/config
NODENAME=$(hostname -s)
EOF
