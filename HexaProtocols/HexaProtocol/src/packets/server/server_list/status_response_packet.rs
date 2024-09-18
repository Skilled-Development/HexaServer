use bytes::BytesMut;
use serde_json::json;

use crate::protocol_util;

pub struct StatusResponsePacket{
    pub server_name: String,
    pub player_protocol: i32,
    pub server_protocols: Vec<i32>

}

impl StatusResponsePacket{

    pub fn new(server_name:String,player_protocol:i32,server_protocols:Vec<i32>) -> StatusResponsePacket{
        StatusResponsePacket{
            server_name,
            player_protocol,
            server_protocols
        }
    }

    pub fn generate_packet(&self) -> BytesMut{
        let mut _new_server_name = self.server_name.clone();
        let _protocol = if self.server_protocols.contains(&self.player_protocol) {
            self.player_protocol
        } else {
            _new_server_name = String::from("We don't support your MC version");
            0 // or any default value you want to assign to protocol
        };


        let response = json!({
            "version": {
                "name": "1.21",
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
        let mut response_packet = BytesMut::new();
        protocol_util::write_int(&mut response_packet, 0x00);
        protocol_util::write_int(&mut response_packet, response_str.len() as i32);
        response_packet.extend_from_slice(response_str.as_bytes());

        let mut packet = BytesMut::new();
        protocol_util::write_int(&mut packet, response_packet.len() as i32);
        packet.extend_from_slice(&response_packet);

        packet
    }
}