use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::swing_arm_packet::SwingArmPacket;
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
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    let packet = SwingArmPacket::read_packet(buffer, client.get_protocol_version());
    let hand = packet.get_hand();
    println!("Swing arm packet received with hand: {:?}", hand);
    Ok(())
}
