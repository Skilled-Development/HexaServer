use bytes::BytesMut;
use hexa_protocol_base:: PacketReader;
use tokio:: net::TcpStream;
extern crate rsa;
extern crate rand;
extern crate byteorder;


use crate:: PlayerConnection;
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = client;
    let _ = socket;
    let _ = length;
    let mut reader  = PacketReader::new(buffer);
    let channel = reader.read_identifier();
    let data = reader.read_bytearray(32767);
    println!("channel {:?}",channel);
    println!("data {:?}",data);

    Ok(())
}