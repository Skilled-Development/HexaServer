use bytes::{Buf, BytesMut};
use hexa_protocol::{packet_builder, protocol_util};
use serde_json::json;
use tokio::{io::AsyncWriteExt, net::TcpStream};

use crate::{player_connection::ClientState, PlayerConnection};


pub async fn handle(length: i32,buffer: &mut BytesMut, socket: &mut TcpStream,client: &mut PlayerConnection ) -> Result<(), String> {
    // Leer los datos del handshake (Protocol Version, Server Address, Server Port, Next State)
    if length  > 3{
            let _protocol_version = protocol_util::read_varint(buffer)?;
            let _server_address = protocol_util::read_string(buffer)?;
            let _server_port = buffer.get_u16();
            let next_state = protocol_util::read_varint(buffer)?;
            if next_state == 2{
                client.set_client_state(ClientState::LOGIN);
            }        
       }else{
          // Responder con el Status Response
          let response = json!({
           "version": {
               "name": "1.21.1",
               "protocol": 767
           },
           "players": {
               "max": 2024,
               "online": 9,
               "sample": [
                   {
                       "name": "§aProbando",
                       "id": "4566e69f-c907-48ee-8d71-d7ba5aa00d20"
                   },
                   {
                       "name": "§aEste",
                       "id": "4566e69f-c907-48ee-8d71-d7ba5aa00d20"
                   },
                   {
                       "name": "§aServidor",
                       "id": "4566e69f-c907-48ee-8d71-d7ba5aa00d20"
                   }
               ]
           },
           "description": {
               "text": "",
               "extra": [
                   { "text": "Este es ", "color": "red" },
                   { "text": "un ejemplo ", "color": "green" },
                   { "text": "de texto ", "color": "blue", "bold": true },
                   { "text": "con varios colores y estilos." }
               ]
           },
           "favicon": "data:image/png;base64,...",
           "enforcesSecureChat": false
       });
       let response_str = serde_json::to_string(&response).unwrap();
       packet_builder::PacketBuilder::new(0x00)
           .write_string( response_str.as_str())
           .send(socket).await?;
       }
       Ok(())
}
