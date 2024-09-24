use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::confirm_teleport_packet::ConfirmTeleportPacket;
use hexa_protocol_base::PacketReader;
use tokio::{io::AsyncWriteExt, net::TcpStream};

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    ConfirmTeleportPacket::read_packet(buffer, client.get_protocol_version());
    Ok(())
}