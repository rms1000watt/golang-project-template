#!/usr/bin/env bash

rm -rf certs &> /dev/null
mkdir certs &> /dev/null

set -e

echo "Generating CA key/crt"
openssl genrsa -out certs/ca.key 4096
openssl req -x509 -new -key certs/ca.key -out certs/ca.crt -days 730 -subj /CN="SuperNiceMegaCorpCA"

echo "Generating Server key/csr"
openssl genrsa -out certs/server.key 4096
openssl req -new -out certs/server.csr -key certs/server.key -config openssl.cnf

echo "Generating Server crt"
openssl x509 -req -in certs/server.csr -out certs/server.crt -days 730 -CAkey certs/ca.key -CA certs/ca.crt -CAcreateserial -CAserial certs/server.serial -extensions v3_ext -extfile openssl.cnf -sha256

echo "Building go binary"
go build

echo "Building docker"
docker build -t golang-project-template .
