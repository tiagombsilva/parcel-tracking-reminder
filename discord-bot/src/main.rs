mod internal {
    tonic::include_proto!("grpc.parcels");
    tonic::include_proto!("grpc.accounts");
}

use futures::stream::StreamExt;
use internal::parcels_client::ParcelsClient;
use tonic::Request;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = ParcelsClient::connect("http://[::1]:9090").await?;

    let mut stream = client.get_parcels(Request::new(())).await?.into_inner();

    while let Some(message) = stream.next().await {
        match message {
            Ok(parcel) => println!("Received parcel: {:?}", parcel),
            Err(e) => eprintln!("Error receiving parcel: {:?}", e),
        }
    }
    Ok(())
}
