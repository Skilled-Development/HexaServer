


use bytes:: BytesMut;
use hexa_protocol_base::{protocol_util, PacketReader};
use tokio::{io::AsyncWriteExt, net::TcpStream};
extern crate rsa;
extern crate rand;
extern crate byteorder;

use crate::PlayerConnection;
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    let mut reader = PacketReader::new(buffer);
    let key = reader.read_identifier(); // Ignorar el identificador del paquete
    let has_payload = reader.read_boolean(); // Ignorar el booleano que indica si hay un payload
    if has_payload {
        let payload_length = reader.read_varint(); // Ignorar el payload
        let payload = reader.read_bytearray(5120);
        println!("Cokie request");
        println!("Key: {:?}", key);
        println!("Payload length: {:?}", payload_length);
        println!("Payload: {:?}", payload);
    }else{
        println!("Cokie request");
        println!("Key: {:?}", key);
        println!("Payload length: 0");
    }
    
    
    /*
        HERE I SEND THE FINISH CONFIGURATION PACKET
     */

    let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x03); 
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    packet.extend_from_slice(&response_packet);
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de ping: {:?}", e))?;
    Ok(())
}