use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::swing_arm_packet::SwingArmPacket;
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
    let client = client.lock().await;
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    let packet = SwingArmPacket::read_packet(buffer, client.get_protocol_version());
    let _hand = packet.get_hand();
    //println!("Swing arm packet received with hand: {:?}", hand);
    Ok(())
}
