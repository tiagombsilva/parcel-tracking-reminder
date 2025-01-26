use notification::notification_service_server::{NotificationService, NotificationServiceServer};
use notification::{NotificationReq, NotificationRes};
use tonic::{transport::Server, Request, Response, Status};

pub mod notification {
    tonic::include_proto!("notification");
}

#[derive(Default)]
pub struct MyNotificationService;

#[tonic::async_trait]
impl NotificationService for MyNotificationService {
    async fn send_notification(
        &self,
        request: Request<NotificationReq>,
    ) -> Result<Response<NotificationRes>, Status> {
        println!(
            "Received notification from Discord ID: {}",
            request.get_ref().discord_id
        );

        let reply = NotificationRes {
            status: String::from("Notification processed successfully!"),
        };

        Ok(Response::new(reply))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let notification_service = MyNotificationService::default();

    Server::builder()
        .add_service(NotificationServiceServer::new(notification_service))
        .serve(addr)
        .await?;

    Ok(())
}
