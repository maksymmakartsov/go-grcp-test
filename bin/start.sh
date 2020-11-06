#!/bin/bash

dir=./tmp/pid
if [ -e "$dir" ]; then
  echo "process is running already"
  exit 1
fi

mkdir -p tmp/pid

export PG_USER=postgres
export PG_PASS=postgres
export PG_DB_NAME=countries_grpc
export PG_FIXTURES=third_party/fixtures

./bin/init-db.sh

./build/package/server &
echo $! >tmp/pid/server.pid
