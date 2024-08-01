mod internal {
    tonic::include_proto!("grpc.parcels");
    tonic::include_proto!("grpc.accounts");
}

use futures::stream::StreamExt;
use internal::parcels_client::ParcelsClient;
use internal::ParcelMessage;
use tonic::transport::Channel;
use tonic::Request;

#[derive(Clone)]
pub struct ParcelsServiceClient {
    client: ParcelsClient<Channel>,
}

impl ParcelsServiceClient {
    pub async fn new(addr: &str) -> Result<Self, Box<dyn std::error::Error>> {
        let client = ParcelsClient::connect(addr.to_string()).await?;
        Ok(ParcelsServiceClient { client })
    }

    pub async fn get_parcels(&mut self) -> Result<Vec<ParcelMessage>, Box<dyn std::error::Error>> {
        let mut stream = self
            .client
            .get_parcels(Request::new(()))
            .await?
            .into_inner();
        let mut parcels = Vec::new();

        while let Some(message) = stream.next().await {
            match message {
                Ok(parcel) => parcels.push(parcel),
                Err(e) => eprintln!("Error receiving parcel: {:?}", e),
            }
        }

        Ok(parcels)
    }
}
