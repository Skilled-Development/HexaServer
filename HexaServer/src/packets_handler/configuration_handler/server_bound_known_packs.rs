use bytes::BytesMut;
use hexa_protocol_1_21::read_data_file_to_bytesmut;
use tokio::{io::AsyncWriteExt, net::TcpStream};

use crate::{protocol_thread::read_varint, PlayerConnection};
use hexa_protocol::{PacketBuilder, PacketReader};

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    let mut reader  = PacketReader::new(buffer);
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
        767=>{
            let datas = vec !["trim_material.data",
            "trim_pattern.data",
            "banner_pattern.data",
            "biome.data",
            "chat_type.data",
            "damage_type.data",
            "dimension_type.data",
            "wolf_variant.data",
            "painting_variant.data"
            ];
            for data in datas {
                println!("Sending data: {}", data);
                let packet = read_data_file_to_bytesmut(data);
                socket.write_all(&packet).await.map_err(|e| e.to_string())?;
            }
        },
        _=>{
            let mut registry_data_packet = PacketBuilder::new(0x05);
            registry_data_packet.write_varint(0);
            registry_data_packet.send(socket).await?;
        }
        
    }
    println!("Sending finish configuration packet");
    //finish configuration
    let mut finish_configuration_packet = PacketBuilder::new(0x03);
    finish_configuration_packet.send(socket).await?;
    Ok(())
}

