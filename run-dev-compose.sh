#!/bin/bash

echo "Stopping container and remove image"
docker compose -f docker-compose-dev.yml down > /dev/null | true

echo "building maven project"
cd docker/java-service
./mvnw install -DskipTests

echo "Starting Services"
cd ../
docker compose -f docker-compose-dev.yml up --build -d

echo "Attaching to API"
docker logs -f backend-api