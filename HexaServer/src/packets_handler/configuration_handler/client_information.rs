use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_base::{PacketBuilder, PacketReader};
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};
extern crate byteorder;
extern crate rand;
extern crate rsa;

use crate::PlayerConnection;
// Asumiendo que tienes estas funciones

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<PlayerConnection>>,
) -> Result<(), String> {
    let _ = length;
    let _ = reader;
    let mut client = client.lock().await;
    let mut reader = PacketReader::new(buffer);
    let locale = reader.read_string();
    let view_distance = reader.read_byte();
    let chat_mode = reader.read_varint();
    let chat_colors = reader.read_boolean();
    let displayed_skin_parts = reader.read_unsigned_byte();
    let main_hand = reader.read_varint();
    let disable_text_filtering = reader.read_boolean();
    let server_listings = reader.read_boolean();
    println!("locale {:?}", locale);
    println!("view_distance {:?}", view_distance);
    println!("chat_mode {:?}", chat_mode);
    println!("chat_colors {:?}", chat_colors);
    println!("displayed_skin_parts {:?}", displayed_skin_parts);
    println!("main_hand {:?}", main_hand);
    println!("disable_text_filtering {:?}", disable_text_filtering);
    println!("server_listings {:?}", server_listings);

    //Send clientbound knonw packs packet
    let packet_id = 0x0E; // Según el protocolo
    let mut packet_builder = PacketBuilder::new(packet_id);
    packet_builder.write_varint(1); // Known Pack Count (0 packs)
    packet_builder.write_string("minecraft");
    packet_builder.write_string("core");
    packet_builder.write_string("1.21");

    client.send_packet_builder(packet_builder).await;
    Ok(())
}
