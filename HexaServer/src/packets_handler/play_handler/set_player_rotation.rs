use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_base::{PacketBuilder, PacketReader};
use tokio::sync::Mutex;

use crate::{Player, ServerProcess};

pub async fn handle(
    buffer: &mut BytesMut,
    client: Arc<Mutex<Player>>,
    server_process: &ServerProcess,
) -> Result<(), String> {
    let mut player = client.lock().await;
    let mut packet = PacketReader::new(buffer);
    let yaw = packet.read_float();
    let pitch = packet.read_float();
    let on_ground = packet.read_boolean();
    player.set_rotation(yaw, pitch);
    player.set_on_ground(on_ground);
    let entity_id = player.get_entity_id();

    let last_pos = player.get_position();
    let packet_x = last_pos.0;
    let packet_y = last_pos.1;
    let packet_z = last_pos.2;
    let packet_yaw = yaw;
    let packet_pitch = pitch;
    let packet_on_ground = on_ground;
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
    head_rotation.write_angle(yaw);
    let client_clone = Arc::clone(&client);
    server_process
        .broadcast_packet_except(client_clone, head_rotation)
        .await;

    Ok(())
}
