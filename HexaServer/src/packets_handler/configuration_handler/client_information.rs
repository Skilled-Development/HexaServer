
use bytes::{BufMut, BytesMut};
use hexa_protocol::{protocol_util, PacketBuilder, PacketReader};
use tokio::{io::AsyncWriteExt, net::TcpStream};
extern crate rsa;
extern crate rand;
extern crate byteorder;
use uuid::Uuid;

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
    packet_builder.write_varint(0); // Known Pack Count (0 packs)
    packet_builder.send(socket).await?;
    Ok(())
}
async fn _send_finish_configuration(socket: &mut tokio::net::TcpStream) -> Result<(), String> {
    let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x03); 
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    packet.extend_from_slice(&response_packet);
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de ping: {:?}", e))?;
    let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x03); 
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    packet.extend_from_slice(&response_packet);
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de ping: {:?}", e))?;
    let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x01); 
    let mut packet = BytesMut::new();
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);

    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_long_be(&mut response_packet, 20000);
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    packet.extend_from_slice(&response_packet);
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de ping: {:?}", e))?;

    /*let mut response_packet = BytesMut::new();
    
    // El paquete Finish Configuration tiene el ID 0x03 y no tiene campos adicionales
    response_packet.put_u8(0x03); // Escribe el ID del paquete

    let mut packet = BytesMut::new();
    
    // Codifica el tamaño del paquete (VarInt)
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    
    // Agrega los datos del paquete
    packet.extend_from_slice(&response_packet);
    
    // Envía el paquete completo
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de finish configuration: {:?}", e))?;
    */
    Ok(())
}

async fn _send_spawn_entity_packet(
    socket: &mut tokio::net::TcpStream,
    entity_id: i32,
    entity_uuid: Uuid,
    entity_type: i32,
    x: f64,
    y: f64,
    z: f64,
    pitch: f32,
    yaw: f32,
    head_yaw: f32,
    data: i32,
    velocity_x: i16,
    velocity_y: i16,
    velocity_z: i16,
) -> Result<(), String> {
    let mut response_packet = BytesMut::new();

    // Packet ID (0x01 for Spawn Entity)
    protocol_util::write_varint(&mut response_packet, 0x01);

    // Entity ID (VarInt)
    protocol_util::write_varint(&mut response_packet, entity_id);

    // Entity UUID (UUID)
    protocol_util::write_uuid(&mut response_packet, entity_uuid);

    // Entity Type (VarInt)
    protocol_util::write_varint(&mut response_packet, entity_type);

    // Coordinates (X, Y, Z as Double)
    protocol_util::write_double(&mut response_packet, x);
    protocol_util::write_double(&mut response_packet, y);
    protocol_util::write_double(&mut response_packet, z);

    // Pitch (Angle)
    let encoded_pitch = ((pitch / 360.0) * 256.0) as u8;
    response_packet.put_u8(encoded_pitch);

    // Yaw (Angle)
    let encoded_yaw = ((yaw / 360.0) * 256.0) as u8;
    response_packet.put_u8(encoded_yaw);

    // Head Yaw (Angle)
    let encoded_head_yaw = ((head_yaw / 360.0) * 256.0) as u8;
    response_packet.put_u8(encoded_head_yaw);

    // Data (VarInt)
    protocol_util::write_varint(&mut response_packet, data);

    // Velocity (Short)
    protocol_util::write_short(&mut response_packet, velocity_x);
    protocol_util::write_short(&mut response_packet, velocity_y);
    protocol_util::write_short(&mut response_packet, velocity_z);

    // Create the packet with its size
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32); // Packet size (VarInt)
    packet.extend_from_slice(&response_packet); // Add the content of the packet

    // Send the packet through the socket
    socket
        .write_all(&packet)
        .await
        .map_err(|e| format!("Error al enviar el paquete Spawn Entity: {:?}", e))?;

    Ok(())
}