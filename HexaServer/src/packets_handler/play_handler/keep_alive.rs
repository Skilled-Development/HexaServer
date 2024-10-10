use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_base::PacketReader;
use tokio::sync::Mutex;

use crate::player::player::Player;

pub async fn handle(buffer: &mut BytesMut, client: Arc<Mutex<Player>>) -> Result<(), String> {
    let client = client.lock().await;
    let mut reader = PacketReader::new(buffer);
    let alive_id = reader.read_long_be();
    if alive_id != client.get_keep_alive_id() {
        return Err("Keep alive id is not the same as the last one".to_string());
    }
    Ok(())
}
