use std::sync::Arc;

use bytes::BytesMut;
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};
extern crate byteorder;
extern crate rand;
extern crate rsa;

use crate::{player::player_connection::ClientState, PlayerConnection};
// Asumiendo que tienes estas funciones

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<PlayerConnection>>,
) -> Result<(), String> {
    let _ = reader;
    let _ = length;
    let mut client = client.lock().await;
    client.set_client_state(ClientState::CONFIGURATION);
    println!("Login Acknowledgement Handler {:?}", buffer);
    Ok(())
}
