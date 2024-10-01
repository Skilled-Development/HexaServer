use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol::packets::{
    client::handshake::ping_request_packet::PingRequestPacket,
    server::handshake::ping_response_packet::PingResponsePacket,
};
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};

use crate::PlayerConnection;
pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<PlayerConnection>>,
) -> Result<(), String> {
    let _ = reader;
    let mut client = client.lock().await;
    let _ = length;
    let request_packet = PingRequestPacket::read_packet(buffer);
    let _response_packet = PingResponsePacket::new(request_packet.get_payload());
    client
        .send_packet_bytes(_response_packet.build().build())
        .await;
    //_response_packet.build().send(socket).await?;
    //client.send_packet(_response_packet.build().build()).await;
    Ok(())
}
