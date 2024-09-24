use std:: time::Instant;

use bytes::BytesMut;
use hexa_protocol_base::{packet_builder::PacketElement, PacketBuilder};
use tokio::net::TcpStream;
extern crate rsa;
extern crate rand;
extern crate byteorder;

use crate::{player::player_connection::ClientState, PlayerConnection};

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
    login_packet.write_unsigned_byte(1u8);
    //previous gamemode
    login_packet.write_byte(0);
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

    /*let mut clone_packet = login_packet.clone();
    let mut clone_packet_buffer = clone_packet.build();
    let length = read_varint(&mut clone_packet_buffer).unwrap();
    println!("Length: {}", length);
    let packetId = read_varint(&mut clone_packet_buffer).unwrap();
    println!("Packet ID: {}", packetId);
    let mut clone_reader = PacketReader::new(&mut clone_packet_buffer);
    let entity_id = clone_reader.read_int();
    println!("Entity ID: {}", entity_id);   
    let hardcore = clone_reader.read_boolean();
    println!("Hardcore: {}", hardcore);
    let dimensions_count = clone_reader.read_varint();
    println!("Dimensions count: {}", dimensions_count);
    let mut dimensions_names = Vec::new();
    for _ in 0..dimensions_count {
        dimensions_names.push(clone_reader.read_string());
    }
    println!("Dimensions names: {:?}", dimensions_names);
    let max_players = clone_reader.read_varint();
    println!("Max players: {}", max_players);
    let view_distance = clone_reader.read_varint();
    println!("View distance: {}", view_distance);
    let simulation_distance = clone_reader.read_varint();
    println!("Simulation distance: {}", simulation_distance);
    let reduced_debug_info = clone_reader.read_boolean();
    println!("Reduced debug info: {}", reduced_debug_info);
    let enable_respawn_screen = clone_reader.read_boolean();
    println!("Enable respawn screen: {}", enable_respawn_screen);
    let do_limited_crafting = clone_reader.read_boolean();
    println!("Do limited crafting: {}", do_limited_crafting);
    let dimension_type = clone_reader.read_varint();
    println!("Dimension type: {}", dimension_type);
    let dimension_name = clone_reader.read_string();
    println!("Dimension name: {}", dimension_name);
    let hashed_seed = clone_reader.read_long_be();
    println!("Hashed seed: {}", hashed_seed);
    let gamemode = clone_reader.read_unsigned_byte();
    println!("Gamemode: {}", gamemode);
    let previous_gamemode = clone_reader.read_byte();
    println!("Previous gamemode: {}", previous_gamemode);
    let is_debug = clone_reader.read_boolean();
    println!("Is debug: {}", is_debug);
    let is_flat = clone_reader.read_boolean();
    println!("Is flat: {}", is_flat);
    let has_death_location = clone_reader.read_boolean();
    println!("Has death location: {}", has_death_location);
    let portal_cooldown = clone_reader.read_varint();
    println!("Portal cooldown: {}", portal_cooldown);
    let enforces_secure_chat = clone_reader.read_boolean();
    println!("Enforces secure chat: {}", enforces_secure_chat);*/




    login_packet.send(socket).await?;
    client.set_last_keep_alive(Instant::now());
    client.set_client_state(ClientState::PLAY);


    let mut game_event_packet = PacketBuilder::new(0x22);
    game_event_packet.write_unsigned_byte(13);
    game_event_packet.write_float(0f32);
    game_event_packet.send(socket).await?;
    let mut synchronize_position = PacketBuilder::new(0x40);
    synchronize_position.write_double(0.0);
    synchronize_position.write_double(10000.0);
    synchronize_position.write_double(0.0);
    synchronize_position.write_float(0.0);
    synchronize_position.write_float(0.0);
    synchronize_position.write_byte(0);
    synchronize_position.write_varint(0);
    synchronize_position.send(socket).await?;
    Ok(())
}
