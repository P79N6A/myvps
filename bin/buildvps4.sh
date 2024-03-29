#!/bin/bash

echo "mkdir"
mkdir -p $HOME/ca
mkdir -p $HOME/conf
mkdir -p $HOME/bin

echo "apt install"
apt install -y nginx curl socat nano
# apt install -y miredo miredo-server

echo "install acme"
curl https://get.acme.sh | sh

echo "install v2ray"
bash <(curl -L -s https://install.direct/go.sh)

echo "copy config"
chmod +x $HOME/bin/*
/root/myvps/bin/confv4.py $1
cp /root/myvps/ca/*.pem /root/ca/

echo "make https"
$HOME/.acme.sh/acme.sh --issue -d v4.xyzjdays.xyz -w /var/www/html --ecc --keylength ec-256
$HOME/.acme.sh/acme.sh --install-cert -d v4.xyzjdays.xyz --key-file /root/ca/key.pem --fullchain-file /root/ca/cert.pem --ecc

echo "start servers"
systemctl restart nginx
start-stop-daemon --start --background --exec $HOME/runv2ray.sh

echo "done"