#!/bin/bash

apt update 

echo "Uninstalling..."


install_dir=/root/xray-configuration

# Remove files
rm $install_dir/config.json
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



cp $install_dir/config.json /usr/local/etc/xray/config.json

