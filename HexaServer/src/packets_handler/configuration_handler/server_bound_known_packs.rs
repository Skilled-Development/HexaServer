use bytes::BytesMut;
use tokio::net::TcpStream;

use crate::PlayerConnection;
use hexa_protocol::{PacketBuilder, PacketReader};

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    let mut reader  = PacketReader::new(buffer);
    let kwnon_packs_count = reader.read_varint();
    println!("Known Packs Count: {}", kwnon_packs_count);


    send_registry_data_packet(socket, "minecraft:blocks", vec!["minecraft:stone", "minecraft:grass"]).await?;
    
    // Enviar Registry Data para "items"
    send_registry_data_packet(socket, "minecraft:items", vec!["minecraft:diamond", "minecraft:iron_ingot"]).await?;

    // Enviar Registry Data para "biomes"
    send_registry_data_packet(socket, "minecraft:biomes", vec!["minecraft:plains", "minecraft:desert"]).await?;

    let mut finish_configuration_packet = PacketBuilder::new(0x03);
    finish_configuration_packet.send(socket).await?;
    Ok(())
}

async fn send_registry_data_packet(socket: &mut TcpStream, registry_id: &str, entries: Vec<&str>) -> Result<(), String> {
    let packet_id = 0x07; // ID del paquete Registry Data
    let mut packet_builder = PacketBuilder::new(packet_id);

    // Escribir el Registry ID (ejemplo: "minecraft:blocks")
    packet_builder.write_identifier(registry_id.to_string());

    // Escribir el Entry Count
    packet_builder.write_varint(entries.len() as i32); // Entry Count

    // Escribir las Entries
    for entry in entries {
        packet_builder.write_identifier(entry.to_string());  // Escribir cada Entry ID
        packet_builder.write_boolean(false);  // Sin NBT data en este ejemplo
    }

    // Enviar el paquete
    packet_builder.send(socket).await
}