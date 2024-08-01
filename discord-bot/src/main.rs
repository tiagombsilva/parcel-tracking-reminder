mod parcels;

use parcels::ParcelsServiceClient;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = ParcelsServiceClient::new("http://[::1]:9090").await?;

    let parcels = client.get_parcels().await?;

    println!("Received parcels: {:?}", parcels);
    Ok(())
}
