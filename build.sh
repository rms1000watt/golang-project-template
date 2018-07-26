#!/usr/bin/env bash

mkdir certs &> /dev/null

set -e

if [[ ! -f certs/ca.key ]] || [[ ! -f certs/ca.crt ]]; then
  echo "Generating CA key/crt"
  openssl genrsa -out certs/ca.key 4096
  openssl req -x509 -new -key certs/ca.key -out certs/ca.crt -days 730 -subj /CN="SuperNiceMegaCorpCA"
fi

if [[ ! -f certs/server.key ]] || [[ ! -f certs/server.csr ]]; then
  echo "Generating Server key/csr"
  openssl genrsa -out certs/server.key 4096
  openssl req -new -out certs/server.csr -key certs/server.key -config openssl.cnf
fi

if [[ ! -f certs/server.crt ]]; then
  echo "Generating Server crt"
  openssl x509 -req -in certs/server.csr -out certs/server.crt -days 730 -CAkey certs/ca.key -CA certs/ca.crt -CAcreateserial -CAserial certs/server.serial -extensions v3_ext -extfile openssl.cnf -sha256
fi

echo "Building go binary"
go build

echo "Building docker"
docker build -t golang-project-template .
