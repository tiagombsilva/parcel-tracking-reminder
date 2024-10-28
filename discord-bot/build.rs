use std::env;
use std::path::PathBuf;


fn main() -> Result<(), Box<dyn std::error::Error>> {

    let proto_file = "./proto/notification.proto";
    let out_dir = PathBuf::from(env::var("OUT_DIR").unwrap());

    tonic_build::configure()
           .build_client(true)
           .build_server(true)
           .file_descriptor_set_path(out_dir.join("notification.bin"))
           .out_dir("./src")
           .compile_protos(&[proto_file], &["proto"])?;

    Ok(())
}