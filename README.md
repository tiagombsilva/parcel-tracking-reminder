# Parcel tracking and reminder

## Features

- Saves your tracking url
- Provides updated tracking 
- Reminds if the package is on deliver and/or custom one
- Discord bot integration

## Tech

[Go] Service:
- Fetches updates from the external API.
- Determines if there's a change in parcel status.
- Stores changes in the Java REST API.
- If a change is detected, triggers a notification via gRPC

[Java] REST API:
- Acts as the central database manager for storing and retrieving tracking updates.
- Exposes endpoints to fetch Accounts and Parcel data

[Rust] Discord Bot:
- Listens for gRPC notifications.
- Sends Discord messages to users based on the provided Discord IDs and update details.

## Starting

Run script run-prd-compose.sh

## External API used for consumer

https://parcelsapp.com/
