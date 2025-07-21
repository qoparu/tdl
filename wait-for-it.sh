#!/bin/sh

# Usage: ./wait-for-it.sh host:port [-- command args...]

HOST_PORT=$(echo "$1" | tr ':' ' ')
HOST=$(echo $HOST_PORT | awk '{print $1}')
PORT=$(echo $HOST_PORT | awk '{print $2}')
shift

while ! nc -z $HOST $PORT; do
  echo "Waiting for $HOST:$PORT..."
  sleep 1
done

echo "$HOST:$PORT is available!"

if [ "$1" = "--" ]; then
  shift
  exec "$@"
fi
