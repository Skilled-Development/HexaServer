use std::sync::Arc;

use bytes::{Buf, BytesMut};
use hexa_protocol_base::PacketReader;
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};

use crate::player::player::Player;

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<Player>>,
) -> Result<(), String> {
    let client = client.lock().await;
    let connection = client.get_connection();
    let connection = connection.lock().await;
    let _ = reader;
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
    let alive_id = reader.read_long_be();
    if alive_id != connection.get_keep_alive_id() {
        println!("Keep alive id is not the same as the last one");
        println!(
            "Received: {}, Player keep alive: {}",
            alive_id,
            connection.get_keep_alive_id()
        );
        return Err("Keep alive id is not the same as the last one".to_string());
    }
    Ok(())
}
