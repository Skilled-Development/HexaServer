use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::set_player_position_and_rotation_packet::SetPlayerPositionAndRotationPacket;
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};

use crate::PlayerConnection;

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<PlayerConnection>>,
) -> Result<(), String> {
    let _ = reader;
    let _ = length;
    let client = client.lock().await;
    if buffer.remaining() < 33 as usize {
        return Err("not_enough_data".to_string());
    }
    let _packet =
        SetPlayerPositionAndRotationPacket::read_packet(buffer, client.get_protocol_version());
    /*println!(
        "x: {}, y: {}, z: {}, yaw: {}, pitch: {}, on_ground: {}",
        packet.get_x(),
        packet.get_y(),
        packet.get_z(),
        packet.get_yaw(),
        packet.get_pitch(),
        packet.get_on_ground()
    );*/
    Ok(())
}
