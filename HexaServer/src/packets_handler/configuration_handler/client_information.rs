
use bytes::BytesMut;
use hexa_protocol_base::{ PacketBuilder, PacketReader};
use tokio:: net::TcpStream;
extern crate rsa;
extern crate rand;
extern crate byteorder;

use crate::PlayerConnection;
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;

    println!("client infor {:?}",buffer);
    let mut reader  = PacketReader::new(buffer);
    let locale = reader.read_string();
    let view_distance = reader.read_byte();
    let chat_mode = reader.read_varint();
    let chat_colors = reader.read_boolean();
    let displayed_skin_parts = reader.read_unsigned_byte();
    let main_hand = reader.read_varint();
    let disable_text_filtering = reader.read_boolean();
    let server_listings = reader.read_boolean();
    println!("locale {:?}",locale);
    println!("view_distance {:?}",view_distance);
    println!("chat_mode {:?}",chat_mode);
    println!("chat_colors {:?}",chat_colors);
    println!("displayed_skin_parts {:?}",displayed_skin_parts);
    println!("main_hand {:?}",main_hand);
    println!("disable_text_filtering {:?}",disable_text_filtering);
    println!("server_listings {:?}",server_listings);
    /*let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x01);
    protocol_util::write_long_be(&mut response_packet, 20000);
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    packet.extend_from_slice(&response_packet);
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de ping: {:?}", e))?;*/


    //send_finish_configuration(socket).await?;
    //send_spawn_entity_packet(socket, 0, client.get_uuid(), 123, 0.0,0.0, 0.0, 
    //90.0, 90.0, 90.0, 0, 0, 0, 0).await?;



   /* let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x00); 
    protocol_util::write_identifier(&mut response_packet, client.get_uuid().to_string());
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    packet.extend_from_slice(&response_packet);
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de ping: {:?}", e))?;*/

    //Send clientbound knonw packs packet
    let packet_id = 0x0E;  // Según el protocolo
    let mut packet_builder = PacketBuilder::new(packet_id);
    packet_builder.write_varint(1); // Known Pack Count (0 packs)
    packet_builder.write_string("minecraft");
    packet_builder.write_string("core");
    packet_builder.write_string("1.21");

    packet_builder.send(socket).await?;
    Ok(())
}
