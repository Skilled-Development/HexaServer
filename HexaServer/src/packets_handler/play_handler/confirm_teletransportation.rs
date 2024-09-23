use bytes::{Buf, BytesMut};
use hexa_protocol_base::PacketReader;
use tokio::{io::AsyncWriteExt, net::TcpStream};

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = length;
    if length == 0{
        println!("Length confirm teletransportation is 0");
        buffer.clear();
        return Ok(());
    }
    if buffer.remaining() < length as usize {
        println!("Not enough data to read confirm teletransportation");
        buffer.clear();
        return Ok(());
    }
    println!("Length confirm teletransportation {:?}",length);
    println!("Buffer confirm teletransportation {:?}",buffer);
    let mut reader = PacketReader::new(buffer);
    let id = reader.read_varint();
    println!("Tp confirm {:?}",id);
    socket.flush().await.unwrap();
    Ok(())
}