

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::ping_request_play_packet::PingRequestPlayPacket;

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut,client: &mut PlayerConnection) -> Result<(), String> {
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    let _packet = PingRequestPlayPacket::read_packet(buffer,client.get_protocol_version());
    Ok(())
}