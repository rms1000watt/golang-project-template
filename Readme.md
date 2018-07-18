# Golang Project Template

## Introduction

This is the starting point for a Golang project with Gometalinter, Govendor, TDD, Precommit git hooks, security best practices, etc.

## Contents

- [Install](#install)
- [Build](#build)
- [Run](#run)
- [cURL](#curl)
- [TODO](#todo)

## Install

```bash
# Govendor
go get -u github.com/kardianos/govendor
govendor sync
```

## Build

```bash
./build.sh
```

## Run

```bash
# Binary
./golang-project-template serve

# Docker
docker-compose up -d
```

## cURL

```bash
curl -H "Origin:https://localhost" --cacert certs/ca.crt -d '{"id":"70640AC2-E6FA-415E-B70B-DE64F74FBF24","name":"ryan"}' -X POST https://localhost:8080/person
curl -H "Origin:https://localhost" --cacert certs/ca.crt -d '{"id":"70640AC2-E6FA-415E-B70B-DE64F74FBF24"}' https://localhost:8080/person
```

## TODO

- [ ] Travis ci
