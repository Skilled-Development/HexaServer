use std::{collections::HashMap, sync::Arc};

use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::set_player_position_and_rotation_packet::SetPlayerPositionAndRotationPacket;
use hexa_protocol_base::PacketBuilder;
use tokio::{net::tcp::OwnedReadHalf, sync::Mutex};

use crate::player::player::Player;

pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    reader: &mut OwnedReadHalf,
    client: Arc<Mutex<Player>>,
    clients: Arc<Mutex<HashMap<String, Arc<Mutex<Player>>>>>,
) -> Result<(), String> {
    let _ = reader;
    let _ = length;
    if buffer.remaining() < 33 as usize {
        return Err("not_enough_data".to_string());
    }
    let mut client_lock = client.lock().await;
    let last_pos = client_lock.get_position();
    let packet =
        SetPlayerPositionAndRotationPacket::read_packet(buffer, client_lock.get_protocol_version());
    client_lock.set_position(packet.get_x(), packet.get_y(), packet.get_z());

    let client = Arc::clone(&client);
    tokio::spawn(async move {
        let client = client.lock().await;
        let clients = Arc::clone(&clients);
        let connection = client.get_connection();
        let connection = Arc::clone(&connection);
        let connection = connection.lock().await;
        let connection_id = connection.get_connection_id().to_string();
        let entity_id = client.get_entity_id();
        let last_pos = (last_pos.0, last_pos.1, last_pos.2);
        let packet_x = packet.get_x();
        let packet_y = packet.get_y();
        let packet_z = packet.get_z();
        let packet_yaw = packet.get_yaw();
        let packet_pitch = packet.get_pitch();
        let packet_on_ground = packet.get_on_ground();

        //TODO: Send packet to other clients
        let mut update_packet = PacketBuilder::new(0x2F);
        update_packet.write_varint(entity_id);
        //delta x
        let x = (packet_x * 4096.0 - last_pos.0 * 4096.0) as i16;
        update_packet.write_short(x);
        //delta y
        let y = (packet_y * 4096.0 - last_pos.1 * 4096.0) as i16;
        update_packet.write_short(y);
        //delta z
        let z = (packet_z * 4096.0 - last_pos.2 * 4096.0) as i16;
        update_packet.write_short(z);
        //yaw
        update_packet.write_angle(packet_yaw);
        //pitch
        update_packet.write_angle(packet_pitch);
        //on ground
        update_packet.write_boolean(packet_on_ground);
        let update_packet: BytesMut = update_packet.build();

        let mut head_rotation = PacketBuilder::new(0x48);
        head_rotation.write_varint(entity_id);
        head_rotation.write_angle(packet_yaw);
        let head_rotation: BytesMut = head_rotation.build();

        let clients = clients.lock().await; // Hacemos el lock una sola vez
        println!("Clients size: {}", clients.len());

        let mut tasks = vec![]; // Creamos una lista de tareas para ser ejecutadas en paralelo

        for (client_id, other_client) in clients.iter() {
            if *client_id == connection_id {
                continue;
            }
            let update_packet_clone = update_packet.clone();
            let head_rotation_clone = head_rotation.clone();
            let other_client = Arc::clone(other_client);

            let task = tokio::spawn(async move {
                let other_client = other_client.lock().await;
                let connection = other_client.get_connection();
                let mut connection = connection.lock().await;
                connection.send_packet_bytes(update_packet_clone).await;
                connection.send_packet_bytes(head_rotation_clone).await;
            });

            tasks.push(task);
        }

        // Esperamos que todas las tareas terminen
        for task in tasks {
            if let Err(e) = task.await {
                eprintln!("Error sending packet: {:?}", e);
            }
        }
    });
    Ok(())
}
