
use serde_json::json;

use crate::{packet_builder, Packet, PacketBuilder, PacketType};

pub struct StatusResponsePacket{
    pub server_name: String,
    pub player_protocol: i32,
    pub server_protocols: Vec<i32>

}
impl Packet for StatusResponsePacket {
    fn get_packet_id(&self) -> i32 {
        0x00 // 0 
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::SERVER
    }
    

}
impl StatusResponsePacket{

    pub fn new(server_name:String,player_protocol:i32,server_protocols:Vec<i32>) -> StatusResponsePacket{
        StatusResponsePacket{
            server_name,
            player_protocol,
            server_protocols
        }
    }

    pub fn build(&self) -> PacketBuilder{
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
        let mut packet = packet_builder::PacketBuilder::new(self.get_packet_id());
        packet.write_string(response_str.as_str());
        packet
        }
    
}