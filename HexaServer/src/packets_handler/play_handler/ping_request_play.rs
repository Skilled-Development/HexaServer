use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::ping_request_play_packet::PingRequestPlayPacket;
use tokio::sync::Mutex;

use crate::player::player::Player;

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    client: Arc<Mutex<Player>>,
) -> Result<(), String> {
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    let client = client.lock().await;
    let connection = client.get_connection();
    let connection = connection.lock().await;
    let _packet = PingRequestPlayPacket::read_packet(buffer, client.get_protocol_version());
    Ok(())
}
