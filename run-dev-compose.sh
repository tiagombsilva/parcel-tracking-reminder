#!/bin/bash

echo "Stopping container and remove image"
docker compose -f docker-compose-dev.yml down > /dev/null | true

echo "Starting Services"
cd ../
docker compose -f docker-compose-dev.yml up --build -d