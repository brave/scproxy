#!/usr/bin/env bash
set -euo pipefail

function check_exit() {
	return=$?;
	if [[ $return -ne 0 ]]; then
		echo "[ERROR] $0 failed"
	fi

  kill %1 #kill autossh process
  exit $return
}

if [[ $# -lt 1 ]]; then
	echo "help: $0 <redis host> [ssh key name]"
	exit 1
fi

if ! type autossh >/dev/null 2>&1; then
	echo "[ERROR] autossh not found"
	exit 2
fi

trap check_exit EXIT

REDIS_HOST=$1

sccache --stop-server >/dev/null 2>&1||:

set -x
autossh -M 0 -o "ServerAliveInterval 30" -o "ServerAliveCountMax 3" -o "ExitOnForwardFailure=yes" -N -L 6379:redis:6379 -p 2223 user@$REDIS_HOST &
set +x

until echo 'ping'|nc 127.0.0.1 6379|grep PONG; do
	if jobs|awk '{print $2}'|grep Exit; then
		echo "Autossh exited"
		exit
	fi
	sleep 2
done

SCCACHE_REDIS=redis://127.0.0.1:6379 sccache --start-server

while sccache -s; do sleep 5; done
