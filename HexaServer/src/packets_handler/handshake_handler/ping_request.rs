use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol::packets::{
    client::handshake::ping_request_packet::PingRequestPacket,
    server::handshake::ping_response_packet::PingResponsePacket,
};
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};

use crate::{player::player::Player, PlayerConnection};
pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<Player>>,
) -> Result<(), String> {
    let _ = reader;
    let client = client.lock().await;
    let connection: Arc<Mutex<PlayerConnection>> = client.get_connection();
    let mut connection = connection.lock().await;
    let _ = length;
    let request_packet = PingRequestPacket::read_packet(buffer);
    let _response_packet = PingResponsePacket::new(request_packet.get_payload());
    connection
        .send_packet_bytes(_response_packet.build().build())
        .await;
    Ok(())
}
