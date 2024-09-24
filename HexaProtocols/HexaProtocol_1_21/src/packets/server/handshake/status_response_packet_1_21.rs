
use serde_json::json;

use hexa_protocol_base::{packet_builder, Packet, PacketBuilder, PacketType};

pub struct StatusResponsePacket_1_21{
        pub server_name:String,
        pub player_protocol:i32,
        pub server_protocols:Vec<i32>,
        pub motd:serde_json::Value,
        pub server_icon:String,
        pub current_player_count:i32,
        pub max_player_count:i32,
        pub sample_text:Option<Vec<String>>
}

impl StatusResponsePacket_1_21{

    pub fn new(
        server_name:String,
        player_protocol:i32,
        server_protocols:Vec<i32>,
        motd:serde_json::Value,
        server_icon:String,
        current_player_count:i32,
        max_player_count:i32,
        sample_text:Option<Vec<String>>
    ) -> StatusResponsePacket_1_21{
        StatusResponsePacket_1_21{
            server_name,
            player_protocol,
            server_protocols,
            motd,
            server_icon,
            current_player_count,
            max_player_count,
            sample_text
        }
    }

    pub fn build(&self) -> PacketBuilder{
        let mut new_server_name = self.server_name.clone();
        let protocol = if self.server_protocols.contains(&self.player_protocol) {
            self.player_protocol
        } else {
            new_server_name = String::from("We don't support your MC version");
            0
        };
        //create an json array for the sample players
        let mut sample_players = Vec::new();
        if let Some(sample_text) = &self.sample_text {
            for text in sample_text {
                sample_players.push(json!({
                    "name": text,
                    "id": "4566e69f-c907-48ee-8d71-d7ba5aa00d20"
                }));
            }
        }
        

        let response = json!({
            "version": {
                "name": new_server_name,
                "protocol": protocol
            },
            "players": {
                "max": self.max_player_count,
                "online": self.current_player_count,
                "sample": sample_players
            },
            "description": self.motd,
            "favicon": "data:image/png;base64,".to_string() + &self.server_icon,
            "enforcesSecureChat": false
        });

        let response_str = serde_json::to_string(&response).unwrap();
        let mut packet = packet_builder::PacketBuilder::new(0);
        packet.write_string(response_str.as_str());
        packet
        }
    
}
