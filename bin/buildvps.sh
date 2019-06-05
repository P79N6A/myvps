#!/bin/bash

echo "mkdir"
mkdir -p $HOME/ca
mkdir -p $HOME/confs
mkdir -p $HOME/bin

echo "apt install"
apt install -y nginx miredo curl socat

echo "install acme"
curl https://get.acme.sh | sh

echo "install v2ray"
curl https://install.direct/go.sh | sh

echo "copy config"
curl https://raw.githubusercontent.com/xyzj/docker-vps/master/buildfiles/confs/default.nginx > /etc/nginx/sites-avaiable/default
curl https://raw.githubusercontent.com/xyzj/docker-vps/master/buildfiles/bin/v2ray.sh > $HOME/bin/runv2ray.sh
chmod +x $HOME/bin/*
/root/bin/confv2ray.py

echo "make https"
$HOME/.acme.sh/acme.sh --issue -d v6.xyzjdays.xyz -w /var/www/html --ecc --keylength ec-256
$HOME/.acme.sh/acme.sh --install-cert -d v6.xyzjdays.xyz --key-file=/root/ca/key.pem --fullchain-file=/root/ca/cert.pem

echo "start servers"
systemctl restart nginx
start-stop-daemon --start --background --exec $HOME/runv2ray.sh

echo "done"