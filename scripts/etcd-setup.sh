#!/bin/bash

# Check the distribution to determine the package manager
if command -v apt &> /dev/null; then
  # Install etcd on Debian/Ubuntu
  apt update
  apt install -y etcd
  systemctl enable etcd
  systemctl start etcd
elif command -v yum &> /dev/null; then
  # Install etcd on CentOS/RHEL
  yum install -y epel-release
  yum install -y etcd
  systemctl enable etcd
  systemctl start etcd
else
  echo "Unsupported Linux distribution. Please install etcd manually."
  exit 1
fi

# Verify the installation and service status
etcd --version
systemctl status etcd
