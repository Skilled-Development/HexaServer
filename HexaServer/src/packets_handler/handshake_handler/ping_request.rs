use bytes::BytesMut;
use hexa_protocol::packets::{
    client::handshake::ping_request_packet::PingRequestPacket,
    server::handshake::ping_response_packet::PingResponsePacket,
};
use tokio::net::TcpStream;

use crate::PlayerConnection;
pub async fn handle(
    length: i32,
    buffer: &mut BytesMut,
    socket: &mut TcpStream,
    client: &mut PlayerConnection,
) -> Result<(), String> {
    let _ = client;
    let _ = length;
    let request_packet = PingRequestPacket::read_packet(buffer);
    let _response_packet = PingResponsePacket::new(request_packet.get_payload());
    _response_packet.build().send(socket).await?;
    //client.send_packet(_response_packet.build().build()).await;
    Ok(())
}
