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
        let _status_response_packet = status_response_packet::StatusResponsePacket::new(
            "HexaServer".to_string(),
            767,
            Vec::new(),
        ).build().send(socket).await?;
       }
       Ok(())
}
