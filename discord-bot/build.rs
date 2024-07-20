fn main() {
    tonic_build::configure()
        .build_server(false)
        .compile(
            &[
                "../backend-api/src/main/proto/parcels.proto",
                "../backend-api/src/main/proto/accounts.proto",
            ],
            &["../backend-api/src/main/proto"],
        ).unwrap();
}