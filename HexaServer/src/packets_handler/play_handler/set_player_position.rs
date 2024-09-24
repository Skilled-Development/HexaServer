

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::set_player_position_packet::SetPlayerPositionPacket;
use tokio::net::TcpStream;

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = socket;
    let _ = client;
    let _ = length;
    if buffer.clone().remaining() < 24 as usize {
        return Err("not_enough_data".to_string());
    }
    let packet = SetPlayerPositionPacket::read_packet(buffer,client.get_protocol_version());
    //println!("x: {}, y: {}, z: {}, on_ground: {}", packet.get_x(), packet.get_y(), packet.get_z(), packet.get_on_ground());
    Ok(())
}