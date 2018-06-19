#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 ]]; then
	echo "help: $0 <redis host> [ssh key name]"
	exit 1
fi


function check_exit() {
    return=$?;
    if [[ $return -eq 0 ]]; then
	echo "[INFO] $0 exiting"
    else
	echo "[ERROR] $0 failed"
    fi

	  docker-compose stop
    exit $return
}

trap check_exit EXIT

REDIS_HOST=$1
SSH_KEY=${2:-id_rsa}

set -x

sccache --stop-server||:

SCPROXY_BACKEND=scproxy://${REDIS_HOST}:2222 SSH_KEY=${SSH_KEY} docker-compose up -d --build

until echo 'ping'|nc 127.0.0.1 6379|grep PONG; do sleep 2; done

SCCACHE_REDIS=redis://127.0.0.1:6379 sccache --start-server

docker-compose logs -f &

while sccache -s; do sleep 2; done
