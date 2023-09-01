#!/bin/bash

apt update 

echo "Uninstalling..."
# Stop and disable xray service
systemctl stop xray
systemctl disable xray

install_dir=/root/xray-configuration

# Remove files
rm /etc/systemd/system/xray.service
rm $install_dir/reality.json
rm $install_dir/subscribe.*
rm /var/www/html/subscribe.*


vnstat > $install_dir/log.txt


ls -laht /var/log
echo "" > /var/log/syslog
echo "" > /var/log/syslog.1

journalctl --vacuum-time=1d



echo "Unistall DONE!"



bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install --beta -u root



# Generate key pair
echo "Generating key pair..."
key_pair=$(xray x25519)
echo "Key pair generation complete."
echo $key_pair

#store public key in a file
touch $install_dir/key_pair.txt
echo $key_pair > $install_dir/key_pair.txt




systemctl daemon-reload
systemctl enable xray
systemctl start xray
systemctl restart xray