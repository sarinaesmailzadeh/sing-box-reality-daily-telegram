#!/bin/bash

install_dir=/root/xray-configuration


rm -rf /var/www/html/subscribe.*
cp  $install_dir/subscribe.* /var/www/html/


rm -rf /var/www/html/aggregate.*
cp  $install_dir/aggregate.* /var/www/html/

# Restart xray service
systemctl restart xray