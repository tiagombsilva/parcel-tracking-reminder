# Parcel tracking and reminder

## Features

- Saves your tracking url
- Provides updated tracking 
- Reminds if the package is on deliver and/or custom one
- Discord bot integration

## Tech

- [Spring-Boot] - Backend REST API (Browser/Discord Bot)
- [Spring-Boot] - Backend gRPC API (Between Consumer Go and SpringBoot)
- [Go] - Consumer of Parcels API
- [Rust] - Discord Bot
- [Postgres] - Database

## Installation

docker-compose up

## External API used for consumer

https://parcelsapp.com/
