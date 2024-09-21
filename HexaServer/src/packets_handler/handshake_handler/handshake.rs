
use std::{collections::HashMap, sync::Arc};

use bytes::BytesMut;
use hexa_protocol::{packets::server::handshake::status_response_packet, HandshakePacket, TextComponent};
use tokio::{ net::TcpStream, sync::Mutex};

use crate::{player_connection::ClientState, PlayerConnection};


pub async fn handle(
    length: i32,
    buffer: &mut BytesMut, 
    socket: &mut TcpStream,
    client: &mut PlayerConnection,
    clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>
) -> Result<(), String> {
    if length  > 3{
            let handshake_packet = HandshakePacket::read_packet(buffer);
            client.set_protocol_version(handshake_packet.get_player_protocol());
            let next_state = handshake_packet.next_state;
            if next_state == 2{
                client.set_client_state(ClientState::LOGIN);
            }       
       }else{
        let (server_name, server_versions, motd_text, server_icon_base64, mut player_count, max_player_count,sample_text) = {
            if let Some(server_config) = &client.server_config {
                let server_config_read_guard = server_config.read().unwrap();
                let server_name = server_config_read_guard.server_name.clone();
                let server_versions = server_config_read_guard.get_protocol_versions_array().clone();
                let motd_text = server_config_read_guard.motd.clone();
                let server_icon_base64 = server_config_read_guard.server_icon_base64
                    .lock()
                    .unwrap()
                    .clone()
                    .unwrap_or_else(|| "".to_string());
                let player_count = server_config_read_guard.player_count;
                let max_player_count = server_config_read_guard.max_player_count;
                let sample_text = server_config_read_guard.sample_text.clone();
                (server_name, server_versions, motd_text, server_icon_base64, player_count, max_player_count,sample_text)
            } else {
                return Err("Server config not present".to_string());
            }
        };
    
        if player_count == -1{
            let clients_guard = clients.lock().await;
            let player_counting = clients_guard
                .iter() 
                .filter_map(|(_, client)| client.try_lock().ok())
                .filter(|client_guard| client_guard.client_state == ClientState::PLAY)
                .count(); 

            player_count = player_counting as i32;
        }

        let mut motd = TextComponent::new();
        motd.set_text(&motd_text);
        let motd_json = motd.to_json();

        let _status_response_packet = status_response_packet::StatusResponsePacket::new(
            server_name,
            client.get_protocol_version(),
            server_versions,
            motd_json,
            server_icon_base64,
            player_count,
            max_player_count,
            sample_text

        ).build().send(socket).await?;
       }
       Ok(())
}
