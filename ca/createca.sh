#!/bin/bash
rm ca* server* client*

echo
echo --- 生成ca，根证书 ---
openssl genrsa -out ca.key
openssl req -x509 -new -key ca.key -days 3650 -out ca.crt
echo
openssl genrsa -out server.key
# openssl req -new -x509 -key server.key -out server.crt -days 3650
openssl req -new -key server.key -out server.csr
echo
echo --- 签发服务器证书 ---
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650 -extfile wlst.cnf

cat ca.crt server.crt > server.pem
