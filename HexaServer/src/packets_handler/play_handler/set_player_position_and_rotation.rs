use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::set_player_position_and_rotation_packet::SetPlayerPositionAndRotationPacket;
use hexa_protocol_base::PacketBuilder;
use tokio::sync::Mutex;

use crate::{player::player::Player, ServerProcess};

pub async fn handle(
    buffer: &mut BytesMut,
    client: Arc<Mutex<Player>>,
    server_process: &ServerProcess,
) -> Result<(), String> {
    if buffer.remaining() < 33 as usize {
        return Err("not_enough_data".to_string());
    }
    let packet = {
        let client_lock = client.lock().await;
        SetPlayerPositionAndRotationPacket::read_packet(buffer, client_lock.get_protocol_version())
    };

    {
        let mut client_lock = client.lock().await;
        let last_pos = client_lock.get_position();
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
        update_packet.write_short((packet_x * 4096.0 - last_pos.0 * 4096.0) as i16);
        update_packet.write_short((packet_y * 4096.0 - last_pos.1 * 4096.0) as i16);
        update_packet.write_short((packet_z * 4096.0 - last_pos.2 * 4096.0) as i16);
        update_packet.write_angle(packet_yaw);
        update_packet.write_angle(packet_pitch);
        update_packet.write_boolean(packet_on_ground);

        server_process
            .broadcast_packet_except(Arc::clone(&client), update_packet)
            .await;

        let mut head_rotation = PacketBuilder::new(0x48);
        head_rotation.write_varint(entity_id);
        head_rotation.write_angle(packet_yaw);

        server_process
            .broadcast_packet_except(Arc::clone(&client), head_rotation)
            .await;
    }
    Ok(())
}
