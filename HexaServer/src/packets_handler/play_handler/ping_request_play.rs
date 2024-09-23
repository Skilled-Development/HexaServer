

use bytes::{Buf, BytesMut};
use hexa_protocol_base::PacketReader;
use rsa::rand_core::le;
use tokio::net::TcpStream;

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    if buffer.remaining() < length as usize {
        println!("Not enough data to read ping request");
        buffer.clear();
        return Ok(());
    }
    let mut reader = PacketReader::new(buffer);
    let payload = reader.read_long_be();
    println!("Ping request {:?}", payload);
    Ok(())
}