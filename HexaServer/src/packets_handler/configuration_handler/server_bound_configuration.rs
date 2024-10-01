use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_base::PacketReader;
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
    let _ = client;
    let _ = reader;
    let _ = length;
    let mut reader = PacketReader::new(buffer);
    let channel = reader.read_identifier();
    let data = reader.read_bytearray(32767);
    println!("channel {:?}", channel);
    println!("data {:?}", data);

    Ok(())
}
