#!/bin/bash
#
# Setup for Control Plane (Master) servers

set -euxo pipefail

MASTER_IP="10.0.0.10"
NODENAME=$(hostname -s)

mkdir -p "$HOME"/.taskweaver
sudo cp -i /etc/taskweaver/admin.conf "$HOME"/.taskweaver/config
sudo chown "$(id -u)":"$(id -g)" "$HOME"/.taskweaver/config

# Save Configs to shared /Vagrant location

# For Vagrant re-runs, check if there is existing configs in the location and delete it for saving new configuration.

config_path="/vagrant/configs"

if [ -d $config_path ]; then
  rm -f $config_path/*
else
  mkdir -p $config_path
fi

cp -i /etc/taskweaver/admin.conf /vagrant/configs/config
touch /vagrant/configs/join.sh
chmod +x /vagrant/configs/join.sh
