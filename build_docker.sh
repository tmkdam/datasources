#!/bin/bash

docker build -t untoldone/bloom-datasources .
docker tag untoldone/bloom-datasources untoldone/bloom-datasources:latest
