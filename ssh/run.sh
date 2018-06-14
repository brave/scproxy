#!/usr/bin/env bash

if [ -z $SCPROXY_BACKEND ]; then
  echo "[ERROR] Must have SCPROXY_BACKEND env var set."
  exit 1
fi

USER=$(echo $SCPROXY_BACKEND|python3 -c 'from urllib import parse; user=parse.urlparse(input()).username; user and print(user)')
PORT=$(echo $SCPROXY_BACKEND|python3 -c 'from urllib import parse; print(parse.urlparse(input()).port)')

set -euo pipefail

HOST=$(echo $SCPROXY_BACKEND|python3 -c 'from urllib import parse; print(parse.urlparse(input()).hostname)')
USER=${USER:-user}
PORT=${PORT:-2222}

set -x

eval $(ssh-agent)
ssh-add "/root/.ssh/${SSH_KEY:-id_rsa}"

ssh -N -L 0.0.0.0:6379:scproxy:6379 -o ExitOnForwardFailure=yes -p "$PORT" "${USER}@${HOST}"
