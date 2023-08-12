#!/bin/bash

install_dir=/root/sing-box


rm -rf /var/www/html/subscribe.*
cp  $install_dir/subscribe.* /var/www/html/


rm -rf /var/www/html/aggregate.*
cp  $install_dir/aggregate.* /var/www/html/

# Restart sing-box service
systemctl restart sing-box