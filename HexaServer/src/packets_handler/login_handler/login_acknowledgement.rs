use std::sync::Arc;

use bytes::BytesMut;
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};
extern crate byteorder;
extern crate rand;
extern crate rsa;

use crate::{player::player_connection::ClientState, Player};
// Asumiendo que tienes estas funciones

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<Player>>,
) -> Result<(), String> {
    let _ = reader;
    let _ = length;
    let client = client.lock().await;
    let connection = client.get_connection();
    let mut connection = connection.lock().await;
    connection.set_client_state(ClientState::CONFIGURATION);
    println!("Login Acknowledgement Handler {:?}", buffer);
    Ok(())
}
