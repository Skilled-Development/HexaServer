use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::confirm_teleport_packet::ConfirmTeleportPacket;

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, client: &mut PlayerConnection) -> Result<(), String> {
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    ConfirmTeleportPacket::read_packet(buffer, client.get_protocol_version());
    Ok(())
}