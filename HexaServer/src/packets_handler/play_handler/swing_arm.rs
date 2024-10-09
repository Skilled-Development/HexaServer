use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::swing_arm_packet::SwingArmPacket;
use tokio::sync::Mutex;

use crate::{player::player::Player, ServerProcess};

pub async fn handle(
    buffer: &mut BytesMut,
    client: Arc<Mutex<Player>>,
    server_process: &ServerProcess,
) -> Result<(), String> {
    let _ = server_process;
    let client = client.lock().await;
    let packet = SwingArmPacket::read_packet(buffer, client.get_protocol_version());
    let _hand = packet.get_hand();
    //println!("Swing arm packet received with hand: {:?}", hand);
    Ok(())
}
