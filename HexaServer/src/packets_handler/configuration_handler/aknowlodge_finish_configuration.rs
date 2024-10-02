use std::{collections::HashMap, sync::Arc, time::Instant};

use bytes::BytesMut;
use hexa_protocol_base::{packet_builder::PacketElement, PacketBuilder};
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};
extern crate byteorder;
extern crate rand;
extern crate rsa;

use crate::{entity::entity::Entity, player::player_connection::ClientState, Player};

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<Player>>,
    clients: Arc<Mutex<HashMap<String, Arc<Mutex<Player>>>>>,
) -> Result<(), String> {
    let _ = buffer;
    let _ = length;
    let _ = reader;
    let client_clone = client.clone();
    let mut client = client.lock().await;
    let connection = client.get_connection();
    let mut connection = connection.lock().await;
    let server_config_lock = connection.get_server_config();
    let server_config = server_config_lock.read().await;
    println!("Handled aknowlodge finish configuration");
    //Now we send the play packet
    let mut login_packet = PacketBuilder::new(0x2B);
    //ENTITY ID
    let entity_arc: Arc<Mutex<dyn Entity>> = client_clone as Arc<Mutex<dyn Entity>>;
    // Llamamos a next_entity_id con la entidad convertida
    let entity_id = server_config
        .get_entity_processor()
        .lock()
        .await
        .next_entity_id(entity_arc)
        .await;
    client.set_entity_id(entity_id);
    login_packet.write_int(entity_id);
    //HARDCORE
    login_packet.write_boolean(false);
    let dimension_names = vec![
        PacketElement::String("minecraft:overworld"),
        PacketElement::String("minecraft:the_nether"),
        PacketElement::String("minecraft:the_end"),
        PacketElement::String("minecraft:overworld_caves"),
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

    connection.send_packet_builder(login_packet).await;
    connection.set_last_keep_alive(Instant::now());
    client.set_position(0.0, 1000.0, 0.0);
    connection.set_client_state(ClientState::PLAY);

    let mut game_event_packet = PacketBuilder::new(0x22);
    game_event_packet.write_unsigned_byte(13);
    game_event_packet.write_float(0f32);
    connection.send_packet_builder(game_event_packet).await;
    let mut synchronize_position = PacketBuilder::new(0x40);
    synchronize_position.write_double(0.0);
    synchronize_position.write_double(1000.0);
    synchronize_position.write_double(0.0);
    synchronize_position.write_float(0.0);
    synchronize_position.write_float(0.0);
    synchronize_position.write_byte(0);
    synchronize_position.write_varint(0);
    connection.send_packet_builder(synchronize_position).await;

    let mut info_update = PacketBuilder::new(0x3E);
    info_update.write_byte(0x01);
    info_update.write_varint(1);
    info_update.write_uuid(client.get_uuid());
    info_update.write_string(&client.get_name());
    info_update.write_varint(0);

    //TODO: spawn player
    let mut spawn_player = PacketBuilder::new(0x01);
    //entity id
    spawn_player.write_varint(entity_id);
    //uuid
    spawn_player.write_uuid(client.get_uuid());
    //entity type
    spawn_player.write_varint(128);
    //position x
    spawn_player.write_double(0.0);
    //position y
    spawn_player.write_double(1000.0);
    //position z
    spawn_player.write_double(0.0);
    //rotation yaw
    spawn_player.write_angle(0.0);
    //rotation pitch
    spawn_player.write_angle(0.0);
    //head rotation
    spawn_player.write_angle(0.0);
    //data
    spawn_player.write_varint(0);
    //velocity x
    spawn_player.write_short(0);
    //velocity y
    spawn_player.write_short(0);
    //velocity z
    spawn_player.write_short(0);

    {
        let clients = clients.lock().await;
        println!("Clients size: {}", clients.len());
        for (client_id, other_client) in clients.iter() {
            if *client_id == connection.get_connection_id() {
                continue;
            }
            let other_client = other_client.lock().await;
            let other_connection = other_client.get_connection();
            let mut other_connection = other_connection.lock().await;
            other_connection
                .send_packet_builder(info_update.clone())
                .await;
            other_connection
                .send_packet_builder(spawn_player.clone())
                .await;
            let mut info_update = PacketBuilder::new(0x3E);
            info_update.write_byte(0x01);
            info_update.write_varint(1);
            info_update.write_uuid(other_client.get_uuid());
            info_update.write_string(&other_client.get_name());
            info_update.write_varint(0);
            connection.send_packet_builder(info_update).await;
            let mut spawn_player = PacketBuilder::new(0x01);
            spawn_player.write_varint(other_client.get_entity_id());
            spawn_player.write_uuid(other_client.get_uuid());
            spawn_player.write_varint(128);
            spawn_player.write_double(other_client.get_position().0);
            spawn_player.write_double(other_client.get_position().1);
            spawn_player.write_double(other_client.get_position().2);
            spawn_player.write_angle(other_client.get_rotation().0);
            spawn_player.write_angle(other_client.get_rotation().1);
            spawn_player.write_angle(other_client.get_rotation().1);
            spawn_player.write_varint(0);
            spawn_player.write_short(other_client.get_velocity().0);
            spawn_player.write_short(other_client.get_velocity().1);
            spawn_player.write_short(other_client.get_velocity().2);
            connection.send_packet_builder(spawn_player).await;
        }
    }
    Ok(())
}
