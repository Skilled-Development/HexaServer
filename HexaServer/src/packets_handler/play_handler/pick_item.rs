use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol_base::PacketReader;
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};

use crate::PlayerConnection;

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<PlayerConnection>>,
) -> Result<(), String> {
    let _ = reader;
    let _ = client;
    let _ = length;
    if buffer.remaining() < length as usize {
        println!("Not enough data to read set item held packet");
        println!(
            "Buffer remaining: {}, Length: {}",
            buffer.remaining(),
            length
        );
        //buffer.clear();
        return Err("not_enough_data".to_string());
    }
    let mut reader = PacketReader::new(buffer);
    let slot = reader.read_varint();
    println!("Slot: {}", slot);
    Ok(())
}
