use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::swing_arm_packet::SwingArmPacket;
use tokio::net::TcpStream;

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = socket;
    let _ = client;
    let _ = length;
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    let packet = SwingArmPacket::read_packet(buffer, client.get_protocol_version());
    let hand = packet.get_hand();
    println!("Swing arm packet received with hand: {:?}", hand);
    Ok(())
}