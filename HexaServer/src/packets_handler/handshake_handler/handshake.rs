
use bytes::BytesMut;
use hexa_protocol::{packets::server::handshake::status_response_packet, HandshakePacket};
use tokio:: net::TcpStream;

use crate::{player_connection::ClientState, PlayerConnection};


pub async fn handle(length: i32,buffer: &mut BytesMut, socket: &mut TcpStream,client: &mut PlayerConnection ) -> Result<(), String> {
    if length  > 3{
            let handshake_packet = HandshakePacket::read_packet(buffer);
            let next_state = handshake_packet.next_state;
            if next_state == 2{
                client.set_client_state(ClientState::LOGIN);
            }       
       }else{
        let server_name = {
            if let Some(server_config) = &client.server_config {
                let server_config_read_guard = server_config.read().unwrap();
                let server_name = server_config_read_guard.server_name.clone();
                server_name
            } else {
                return Err("Server config not present".to_string());
            }
        };

        let server_versions = {
            if let Some(server_config) = &client.server_config {
                let server_config_read_guard = server_config.read().unwrap();
                let server_versions = server_config_read_guard.get_protocol_versions_array().clone();
                server_versions
            } else {
                return Err("Server config not present".to_string());
            }
        };

        let _status_response_packet = status_response_packet::StatusResponsePacket::new(
            server_name,
            767,
            server_versions,
        ).build().send(socket).await?;
       }
       Ok(())
}
