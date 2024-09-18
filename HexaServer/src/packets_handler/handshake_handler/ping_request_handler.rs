use bytes::{Buf, BytesMut};
use hexa_protocol::protocol_util;
use tokio::{io::AsyncWriteExt, net::TcpStream};

use crate::PlayerConnection;
pub async fn handle(length: i32,buffer: &mut BytesMut, socket: &mut TcpStream,client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    let payload = buffer.get_i64();
    let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x01);
    protocol_util::write_long_be(&mut response_packet, payload);
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32);
    packet.extend_from_slice(&response_packet);
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de ping: {:?}", e))?;
    Ok(())
}




