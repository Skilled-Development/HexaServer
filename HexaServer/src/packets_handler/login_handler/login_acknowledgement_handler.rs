use std::io::Cursor;

use byteorder::{BigEndian, WriteBytesExt};
use bytes::BytesMut;
use hexa_protocol::{protocol_util, PacketReader};
use tokio::{io::AsyncWriteExt, net::TcpStream};
extern crate rsa;
extern crate rand;
extern crate byteorder;

use rsa::{RsaPrivateKey, RsaPublicKey};
use rand::rngs::OsRng;
use rsa::pkcs1::EncodeRsaPublicKey;
use uuid::{Uuid};

use crate::{player_connection::ClientState, PlayerConnection};
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    client.set_client_state(ClientState::CONFIGURATION);
    println!("Login Acknowledgement Handler {:?}",buffer);
    Ok(())
}