use std::time::Instant;

use bytes::{Buf, BytesMut};
use hexa_protocol::{packet_builder::PacketElement, PacketBuilder};
use tokio::{io::AsyncWriteExt, net::TcpStream};
extern crate rsa;
extern crate rand;
extern crate byteorder;

use crate::{player_connection::ClientState, PlayerConnection};
// Asumiendo que tienes estas funciones

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = buffer;
    let _ = client;
    let _ = length;
    println!("Handled aknowlodge finish configuration");
    //Now we send the play packet
    let mut login_packet = PacketBuilder::new(0x2B);
    //ENTITY ID
    login_packet.write_int(0);
    //HARDCORE
    login_packet.write_boolean(false);
    let dimension_names = vec![
        PacketElement::String("minecraft:overworld"),
         PacketElement::String("minecraft:the_nether"), 
         PacketElement::String("minecraft:the_end"),
         PacketElement::String("minecraft:overworld_caves")
         ];
    //DIMENSIONS COUNT
    login_packet.write_varint(dimension_names.len() as i32);
    //DIMENSIONS NAMES
    login_packet.write_array(&dimension_names);
    //MAX PLAYERS
    login_packet.write_varint(2024);
    //VIEW DISTANCE
    login_packet.write_varint(5);
    //SIMULATION DISTANCE
    login_packet.write_varint(5);
    //REDUCED DEBUG INFO
    login_packet.write_boolean(false);
    //ENABLE RESPAWN SCREEN
    login_packet.write_boolean(true);
    //DO LIMITED CRAFTING
    login_packet.write_boolean(false);
    //DIMENSION TYPE
    login_packet.write_varint(0);
    //DIMENSION NAME
    login_packet.write_identifier("minecraft:overworld".to_string());
    //HASHED SEED (IDK)
    let hex = "79aa9a41";
    let value = i64::from_str_radix(hex, 16).expect("Invalid hex string");
    login_packet.write_long_be(value);
    //GAMEMODE
    login_packet.write_unsigned_byte(2);
    //previous gamemode
    login_packet.write_byte(1);
    //IS DEBUG
    login_packet.write_boolean(false);
    //IS FLAT
    login_packet.write_boolean(true);
    //HAS DEATH LOCATION
    login_packet.write_boolean(false);
    //PORTAL COOLDOWN
    login_packet.write_varint(0);
    //ENFORCES SECURE CAHT
    login_packet.write_boolean(false);
    login_packet.send(socket).await?;
    client.set_last_keep_alive(Instant::now());
    client.set_client_state(ClientState::PLAY);


    let mut game_event_packet = PacketBuilder::new(0x22);
    game_event_packet.write_unsigned_byte(13);
    game_event_packet.write_float(0f32);
    game_event_packet.send(socket).await?;
    /*let mut synchronize_position = PacketBuilder::new(0x40);
    synchronize_position.write_double(0.0);
    synchronize_position.write_double(100.0);
    synchronize_position.write_double(0.0);
    synchronize_position.write_float(0.0);
    synchronize_position.write_float(0.0);
    synchronize_position.write_byte(0);
    synchronize_position.write_varint(0);
    synchronize_position.send(socket).await?;*/
    Ok(())
}