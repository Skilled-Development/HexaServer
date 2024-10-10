use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol::packets::client::play::set_item_held_packet::SetItemHeldPacket;
use tokio::sync::Mutex;

use crate::player::player::Player;

pub async fn handle(buffer: &mut BytesMut, client: Arc<Mutex<Player>>) -> Result<(), String> {
    let mut client = client.lock().await;
    let packet = SetItemHeldPacket::read_packet(buffer, client.get_protocol_version());
    client.set_held_slot(packet.get_slot());
    Ok(())
}
