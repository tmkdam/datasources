#!/bin/bash

sleep 30 # wait for db to come up if started at the same time

cd /go/bin/linux_amd64/datasources/
./runner bootstrap || true
./runner update