#!/usr/bin/env python3

import uuid
import sys

js = '''{
    "inbounds": [
        {
            "port": 6890,
            "protocol": "vmess",
            "settings": {
                "clients": [
                    {
                        "id": "%s",
                        "alterId": 6,
                        "security": "auto"
                    }
                ]
            },
            "streamSettings": {
                "network": "mkcp",
                "kcpSettings": {
                    "mtu": 1350,
                    "tti": 20,
                    "uplinkCapacity": 5,
                    "downlinkCapacity": 20,
                    "congestion": true,
                    "readBufferSize": 1,
                    "writeBufferSize": 1,
                    "header": {
                        "type": "utp"
                    }
                }
            }
        },
        {
            "port": 6891,
            "address": "127.0.0.1",
            "protocol": "vmess",
            "settings": {
                "clients": [
                    {
                        "id": "%s",
                        "alterId": 6,
                        "security": "auto"
                    }
                ]
            },
            "streamSettings": {
                "network": "ws",
                "wsSettings": {
                    "path": "/xx"
                }
            }
        },
        {
            "port": 6892,
            "protocol": "vmess",
            "settings": {
                "clients": [
                    {
                        "id": "%s",
                        "alterId": 6,
                        "security": "auto"
                    }
                ]
            },
            "streamSettings": {
                "network": "ws",
                "wsSettings": {
                    "path": "/bwh"
                }
            }
        }
    ],
    "outbounds": [
        {
            "protocol": "freedom",
            "settings": {}
        }
    ]
}
'''


jslite = '''{
    "inbounds": [
        {
            "port": 6891,
            "address": "127.0.0.1",
            "protocol": "vmess",
            "settings": {
                "clients": [
                    {
                        "id": "%s",
                        "alterId": 6,
                        "security": "auto"
                    }
                ]
            },
            "streamSettings": {
                "network": "ws",
                "wsSettings": {
                    "path": "/xx"
                }
            }
        }
    ],
    "outbounds": [
        {
            "protocol": "freedom",
            "settings": {}
        }
    ]
}
'''

defaultnginx = '''server {
	# listen 80 default_server;
	listen [::]:80 default_server;

	# SSL configuration
	#
	# listen 443 ssl default_server;
	listen [::]:443 ssl default_server;

	root /var/www/html;

	index index.html index.htm index.nginx-debian.html;

	server_name v6.xyzjdays.xyz;

	ssl_certificate       /root/ca/cert.pem;
	ssl_certificate_key   /root/ca/key.pem;
	ssl_protocols         TLSv1 TLSv1.1 TLSv1.2;
	ssl_ciphers           HIGH:!aNULL:!MD5;

	location /xx { # 与 V2Ray 配置中的 path 保持一致
	        proxy_redirect off;
	        proxy_pass http://127.0.0.1:6891;#假设WebSocket监听在环回地址的10000端口上
	        proxy_http_version 1.1;
	        proxy_set_header Upgrade $http_upgrade;
	        proxy_set_header Connection "upgrade";
	        proxy_set_header Host $http_host;
        }

	location /bwh { # 与 V2Ray 配置中的 path 保持一致
	        proxy_redirect off;
	        proxy_pass http://v4.xyzjdays.xyz:6892;#假设WebSocket监听在环回地址的10000端口上
	        proxy_http_version 1.1;
	        proxy_set_header Upgrade $http_upgrade;
	        proxy_set_header Connection "upgrade";
	        proxy_set_header Host $http_host;
        }

	location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
		try_files $uri $uri/ =404;
	}
}'''


runv2ray = '''#!/bin/bash 

/usr/bin/v2ray/v2ray -config=/root/conf/v2rays.json
'''

if __name__ == "__main__":
    u1 = uuid.uuid4().hex
    u2 = uuid.uuid4().hex
    u3 = uuid.uuid4().hex

    with open("/root/bin/runv2ray.sh", "w") as f:
        f.write(runv2ray)
        f.close()

    with open("/etc/nginx/sites-available/default", "w") as f:
        f.write(defaultnginx)
        f.close()

    with open("/root/conf/v2rays.json", "w") as f:
        if len(sys.argv) > 1:
            if sys.argv[1] == "full":
                f.write(js % (u1, u2, u3))
            else:
                f.write(jslite % (u1))
        else:
            f.write(jslite % (u1))
        f.close()
