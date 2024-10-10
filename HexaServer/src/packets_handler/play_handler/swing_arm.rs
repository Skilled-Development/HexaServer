use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol::packets::client::play::swing_arm_packet::SwingArmPacket;
use hexa_protocol_base::{player::hand::Hand, PacketBuilder};
use tokio::sync::Mutex;

use crate::{player::player::Player, ServerProcess};

pub async fn handle(
    buffer: &mut BytesMut,
    client: Arc<Mutex<Player>>,
    server_process: &ServerProcess,
) -> Result<(), String> {
    let _ = server_process;
    let client_clone = Arc::clone(&client);
    let client = client.lock().await;
    let packet = SwingArmPacket::read_packet(buffer, client.get_protocol_version());
    let hand = packet.get_hand();
    //TODO
    let mut animation_packet = PacketBuilder::new(0x03);
    animation_packet.write_varint(client.get_entity_id());
    if hand == Hand::MainHand {
        animation_packet.write_varint(0);
    } else {
        animation_packet.write_varint(3);
    }
    server_process
        .broadcast_packet_except(client_clone, animation_packet)
        .await;
    Ok(())
}
