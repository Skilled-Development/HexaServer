use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::set_player_position_packet::SetPlayerPositionPacket;
use hexa_protocol_base::PacketBuilder;
use tokio::sync::Mutex;

use crate::{player::player::Player, ServerProcess};

pub async fn handle(
    buffer: &mut BytesMut,
    client: Arc<Mutex<Player>>,
    server_process: &ServerProcess,
) -> Result<(), String> {
    let mut player: tokio::sync::MutexGuard<'_, Player> = client.lock().await;
    if buffer.clone().remaining() < 24 as usize {
        return Err("not_enough_data".to_string());
    }
    let last_pos = player.get_position();
    let packet = SetPlayerPositionPacket::read_packet(buffer, player.get_protocol_version());
    let entity_id = player.get_entity_id();
    let (packet_x, packet_y, packet_z) = (packet.get_x(), packet.get_y(), packet.get_z());
    let on_ground = packet.get_on_ground();

    let mut update_packet = PacketBuilder::new(0x2E);
    update_packet.write_varint(entity_id);
    update_packet.write_short(((packet_x - last_pos.0) * 4096.0) as i16);
    update_packet.write_short(((packet_y - last_pos.1) * 4096.0) as i16);
    update_packet.write_short(((packet_z - last_pos.2) * 4096.0) as i16);
    update_packet.write_boolean(on_ground);
    server_process
        .broadcast_packet_except(Arc::clone(&client), update_packet)
        .await;

    player.set_position(packet_x, packet_y, packet_z);
    player.set_on_ground(on_ground);
    Ok(())
}
