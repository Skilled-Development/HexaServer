use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol::packets::client::play::ping_request_play_packet::PingRequestPlayPacket;
use tokio::sync::Mutex;

use crate::player::player::Player;

pub async fn handle(buffer: &mut BytesMut, client: Arc<Mutex<Player>>) -> Result<(), String> {
    let client = client.lock().await;
    let _packet = PingRequestPlayPacket::read_packet(buffer, client.get_protocol_version());
    Ok(())
}
