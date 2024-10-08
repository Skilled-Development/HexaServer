use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::set_player_position_and_rotation_packet::SetPlayerPositionAndRotationPacket;
use hexa_protocol_base::PacketBuilder;
use tokio::sync::Mutex;

use crate::{player::player::Player, ServerProcess};

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    client: Arc<Mutex<Player>>,
    server_process: &ServerProcess,
) -> Result<(), String> {
    let _ = server_process;
    let _ = length;
    if buffer.remaining() < 33 as usize {
        return Err("not_enough_data".to_string());
    }
    let client_clone = Arc::clone(&client);
    let mut client_lock = client_clone.lock().await;
    let last_pos = client_lock.get_position();
    let packet =
        SetPlayerPositionAndRotationPacket::read_packet(buffer, client_lock.get_protocol_version());
    client_lock.set_position(packet.get_x(), packet.get_y(), packet.get_z());

    let entity_id = client_lock.get_entity_id();
    let last_pos = (last_pos.0, last_pos.1, last_pos.2);
    let packet_x = packet.get_x();
    let packet_y = packet.get_y();
    let packet_z = packet.get_z();
    let packet_yaw = packet.get_yaw();
    let packet_pitch = packet.get_pitch();
    let packet_on_ground = packet.get_on_ground();

    let mut update_packet = PacketBuilder::new(0x2F);
    update_packet.write_varint(entity_id);
    let x = (packet_x * 4096.0 - last_pos.0 * 4096.0) as i16;
    update_packet.write_short(x);
    let y = (packet_y * 4096.0 - last_pos.1 * 4096.0) as i16;
    update_packet.write_short(y);
    let z = (packet_z * 4096.0 - last_pos.2 * 4096.0) as i16;
    update_packet.write_short(z);
    update_packet.write_angle(packet_yaw);
    update_packet.write_angle(packet_pitch);
    update_packet.write_boolean(packet_on_ground);
    let client_clone = Arc::clone(&client);
    server_process
        .broadcast_packet_except(client_clone, update_packet)
        .await;

    let mut head_rotation = PacketBuilder::new(0x48);
    head_rotation.write_varint(entity_id);
    head_rotation.write_angle(packet_yaw);
    let client_clone = Arc::clone(&client);
    server_process
        .broadcast_packet_except(client_clone, head_rotation)
        .await;

    Ok(())
}
