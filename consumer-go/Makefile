#run-orders:
#	@go run services/orders/*.go
gen:
	@protoc \
		--proto_path=../backend-api/src/main/proto "parcels.proto"\
		--go_out=internal/common/parcels --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/parcels --go-grpc_opt=paths=source_relative
	@protoc \
		--proto_path=../backend-api/src/main/proto "accounts.proto"\
		--go_out=internal/common/accounts --go_opt=paths=source_relative \
  		--go-grpc_out=internal/common/accounts --go-grpc_opt=paths=source_relative