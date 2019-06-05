#!/usr/bin/env python3

import uuid

js = '''{
    "inbounds": [
        {
            "port": 6890,
            "protocol": "vmess",
            "settings": {
                "clients": [
                    {
                        "id": "{0}",
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
                        "id": "{1}",
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
                        "id": "{2}",
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

if __name__ == "__main__":
    u1 = uuid.uuid4().hex
    u2 = uuid.uuid4().hex
    u3 = uuid.uuid4().hex

    with open("/root/conf/v2rays.json", "w") as f:
        f.write(js.format(u1, u2, u3))
        f.close()
