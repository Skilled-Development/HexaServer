
use bytes::BytesMut;
use tokio:: net::TcpStream;
extern crate rsa;
extern crate rand;
extern crate byteorder;

use crate::{player_connection::ClientState, PlayerConnection};
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = socket;
    let _ = length;
    client.set_client_state(ClientState::CONFIGURATION);
    println!("Login Acknowledgement Handler {:?}",buffer);
    Ok(())
}