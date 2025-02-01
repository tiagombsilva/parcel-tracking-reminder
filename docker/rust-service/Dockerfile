FROM rust:latest AS build

WORKDIR /app

COPY Cargo.toml ./
COPY src ./src
COPY proto ./proto
COPY build.rs ./

RUN apt-get update && apt-get install -y protobuf-compiler && rm -rf /var/lib/apt/lists/*

RUN cargo build --release

# Final image using distroless
FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=build /app/target/release/discord-bot .

ENTRYPOINT ["/app/rust-discord-bot"]