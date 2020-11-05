#!/bin/bash

kill -TERM $(cat tmp/pid/server.pid)

rm -rf tmp/pid
