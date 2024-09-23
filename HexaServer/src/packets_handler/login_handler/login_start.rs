use std::io::Cursor;

use byteorder::{BigEndian, WriteBytesExt};
use bytes::BytesMut;
use hexa_protocol_base::{ PacketBuilder, PacketReader};
use tokio:: net::TcpStream;
extern crate rsa;
extern crate rand;
extern crate byteorder;

use rsa::{RsaPrivateKey, RsaPublicKey};
use rand::rngs::OsRng;
use rsa::pkcs1::EncodeRsaPublicKey;
use uuid::Uuid;

use crate::PlayerConnection;
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = length;
    let mut reader = PacketReader::new(buffer);
    let username = reader.read_string();
    println!("Username: {}", username);
    let uuid = reader.read_uuid();
    println!("UUID: {}", uuid);

    client.set_username(username.clone());
    client.set_uuid(uuid);


    //let (public_key, verify_token) = generate_keys();
    
    // Crear el paquete con public key y verify token
    let _ = send_login_success_packet(socket, &username, uuid.to_string()).await;

    Ok(())
}


// Función para generar las claves y el verify token
fn _generate_keys() -> (RsaPublicKey, Vec<u8>) {
    let mut rng = OsRng;
    let bits = 1024;
    let private_key = RsaPrivateKey::new(&mut rng, bits).expect("Error al generar clave privada");
    let public_key = RsaPublicKey::from(&private_key);
    let verify_token: [u8; 4] = rand::random();
    
    (public_key, verify_token.to_vec())
}

// Función para crear el paquete con public key y verify token
fn _create_packet(public_key: &RsaPublicKey, verify_token: &[u8]) -> Vec<u8> {
    
    let mut packet = Vec::new();
    
    // Serializar la clave pública en formato PEM (si necesitas DER, cambia este bloque)
    let public_key_pem = public_key.to_pkcs1_pem(rsa::pkcs1::LineEnding::LF).unwrap();
    let public_key_bytes = public_key_pem.as_bytes();
    
    // Escribir la longitud de la clave pública (varint en BigEndian)
    let public_key_length = public_key_bytes.len() as u32;
    let mut cursor = Cursor::new(Vec::new());
    WriteBytesExt::write_u32::<BigEndian>(&mut cursor, public_key_length).unwrap();
    packet.extend(cursor.into_inner());

    // Escribir los bytes de la clave pública
    packet.extend_from_slice(public_key_bytes);
    
    // Escribir la longitud del verify token
    let verify_token_length = verify_token.len() as u32;
    let mut cursor = Cursor::new(Vec::new());
    WriteBytesExt::write_u32::<BigEndian>(&mut cursor, verify_token_length).unwrap();
    packet.extend(cursor.into_inner());

    // Escribir los bytes del verify token
    packet.extend_from_slice(verify_token);

    
    
    packet
}

async fn send_login_success_packet(socket: &mut tokio::net::TcpStream, username: &str, uuid_string: String) -> Result<(), String> {
    let player_uuid = Uuid::parse_str(&uuid_string).unwrap();
    PacketBuilder::new(0x02)
        .write_uuid(player_uuid)
        .write_string(username)
        .write_varint(0)
        .write_boolean(false)
        .send(socket).await?;

    Ok(())
}