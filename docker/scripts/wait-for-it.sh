#!/bin/bash

# Usage: ./wait-for-it.sh <host>/<port> <command_to_run>
HOST=$1
shift 2
CMD="$@"

# Function to check the availability of the host and port
check_port() {
    (echo > /dev/tcp/$HOST) &>/dev/null
    return $?
}

# Loop until the service is available
until check_port; do
  echo "Waiting for $HOST to become available..."
  sleep 1
done

echo "$HOST is now available!"

# Execute the command passed to the script
echo "Executing command: $CMD"
exec $CMD
