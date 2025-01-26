FROM rust:1.73 AS build

WORKDIR /app

COPY Cargo.toml ./
COPY src ./src
COPY proto ./proto

RUN apt-get update && apt-get install -y protobuf-compiler && rm -rf /var/lib/apt/lists/*

ENV OUT_DIR=/app/proto

RUN cargo build --release

RUN cargo test

# Final image using distroless
FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=build /app/target/release/rust-service .

ENTRYPOINT ["/app/rust-service"]