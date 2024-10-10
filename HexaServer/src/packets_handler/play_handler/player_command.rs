use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_base::PacketReader;
use tokio::sync::Mutex;

use crate::Player;

pub async fn handle(buffer: &mut BytesMut, client: Arc<Mutex<Player>>) -> Result<(), String> {
    let _ = client;
    let mut packet = PacketReader::new(buffer);
    let _entity_id = packet.read_varint();
    let action_id = packet.read_varint();
    let _jump_boost = packet.read_varint();
    //TODO: Implement player command handling
    let mut player = client.lock().await;
    if action_id == 0 {
        player.set_sneaking(true);
    } else if action_id == 1 {
        player.set_sneaking(false);
    } else if action_id == 3 {
        player.set_sprinting(true);
    } else if action_id == 4 {
        player.set_sprinting(false);
    }
    Ok(())
}
