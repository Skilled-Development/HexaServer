use bytes::{Buf, BytesMut};
use hexa_protocol_base::PacketReader;
use tokio::net::TcpStream;

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    if buffer.remaining() < length as usize {
        println!("Not enough data to read set item held packet");
        buffer.clear();
        return Ok(());
    }
    let mut reader = PacketReader::new(buffer);
    let hand = reader.read_varint();
    println!("hand: {}", hand);
    Ok(())
}