fn main() {
    tonic_build::configure()
        .build_server(false)
        //.out_dir("src/internal/common")  // you can change the generated code's location
        .compile(
            &[
                "../backend-api/src/main/proto/parcels.proto",
                "../backend-api/src/main/proto/accounts.proto",
            ],
            &["../backend-api/src/main/proto"], // specify the root location to search proto dependencies
        ).unwrap();
}