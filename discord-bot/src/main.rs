mod internal {
    tonic::include_proto!("grpc.parcels");
    tonic::include_proto!("grpc.accounts");
}

use internal::parcels_client::ParcelsClient;
use tonic::Request;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = ParcelsClient::connect("http://[::1]:9090").await?;

    let request = Request::new(());

    let message = client.get_parcels(request).await?;

    println!("RESPONSE={:?}", message);
    Ok(())
}