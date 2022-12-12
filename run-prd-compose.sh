#!/bin/bash

echo "Stopping container and remove image"
docker-compose -f docker-compose-prd.yml down > /dev/null | true

echo "Starting Services"
docker-compose -f docker-compose-prd.yml up -d

echo "Attaching to API"
docker logs -f parcel-tracking-reminder-backend