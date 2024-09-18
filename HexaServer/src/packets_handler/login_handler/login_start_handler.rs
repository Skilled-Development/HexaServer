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

use crate::{player_connection, PlayerConnection};
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
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
fn generate_keys() -> (RsaPublicKey, Vec<u8>) {
    let mut rng = OsRng;
    let bits = 1024;
    let private_key = RsaPrivateKey::new(&mut rng, bits).expect("Error al generar clave privada");
    let public_key = RsaPublicKey::from(&private_key);
    let verify_token: [u8; 4] = rand::random();
    
    (public_key, verify_token.to_vec())
}

// Función para crear el paquete con public key y verify token
fn create_packet(public_key: &RsaPublicKey, verify_token: &[u8]) -> Vec<u8> {
    
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

    let mut response_packet = BytesMut::new();
    protocol_util::write_varint(&mut response_packet, 0x02); // Packet ID 0x02 (Login Success)

    // Escribir el UUID del jugador
    protocol_util::write_uuid(&mut response_packet, player_uuid);

    // Escribir el nombre de usuario (con una longitud máxima de 16 caracteres)
    protocol_util::write_string(&mut response_packet, username);

    // Escribir la cantidad de propiedades (0 en modo offline)
    protocol_util::write_varint(&mut response_packet, 0); // No properties in offline mode

    // Strict Error Handling - Lo ajustamos a `false` (no desconectar si hay errores)
    protocol_util::write_boolean(&mut response_packet, false);

    // Crear el paquete final que contiene la longitud del paquete
    let mut packet = BytesMut::new();
    protocol_util::write_varint(&mut packet, response_packet.len() as i32); // Longitud del paquete
    packet.extend_from_slice(&response_packet);

    // Enviar el paquete al cliente
    socket.write_all(&packet).await.map_err(|e| format!("Error al enviar el paquete de Login Success: {:?}", e))?;

    Ok(())
}