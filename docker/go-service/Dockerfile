FROM golang:1.23 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o go-service .

# Final image using distroless
FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=build /app/go-service .

CMD ["/app/go-service"]