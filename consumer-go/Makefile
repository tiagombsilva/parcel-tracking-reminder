run-cron:
	@go run main.go
gen:
	@protoc \
		--proto_path=../backend-api/src/main/proto "parcels.proto"\
		--go_out=internal/common/parcels --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/parcels --go-grpc_opt=paths=source_relative