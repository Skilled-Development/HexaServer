use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_1_21::read_data_file_to_bytesmut;
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};

use crate::PlayerConnection;
use hexa_protocol_base::{PacketBuilder, PacketReader};

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<PlayerConnection>>,
) -> Result<(), String> {
    let _ = reader;
    let _ = length;
    let mut client = client.lock().await;
    let mut reader = PacketReader::new(buffer);
    let kwnon_packs_count = reader.read_varint();
    println!("Known Packs Count: {}", kwnon_packs_count);
    let namespace = reader.read_string();
    println!("Namespace: {}", namespace);
    let id = reader.read_string();
    println!("ID: {}", id);
    let version = reader.read_string();
    println!("Version: {}", version);

    //send registrydata packet
    //TODO: ADD TO EACH VERSION
    match client.get_protocol_version() {
        767 => {
            let datas = vec![
                "trim_material.data",
                "trim_pattern.data",
                "banner_pattern.data",
                "biome.data",
                "chat_type.data",
                "damage_type.data",
                "dimension_type.data",
                "wolf_variant.data",
                "painting_variant.data",
            ];
            for data in datas {
                println!("Sending data: {}", data);
                let packet = read_data_file_to_bytesmut(data);
                client.send_packet_bytes(packet).await;
            }
        }
        _ => {
            let mut registry_data_packet = PacketBuilder::new(0x05);
            registry_data_packet.write_varint(0);
            client.send_packet_builder(registry_data_packet).await;
        }
    }
    println!("Sending finish configuration packet");
    //finish configuration
    let finish_configuration_packet = PacketBuilder::new(0x03);
    client
        .send_packet_builder(finish_configuration_packet)
        .await;
    Ok(())
}
