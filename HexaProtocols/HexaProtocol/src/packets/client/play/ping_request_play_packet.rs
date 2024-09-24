use bytes::BytesMut;

use hexa_protocol_1_21::packets::client::play::ping_request_play_packet_1_21::PingRequestPlayPacket1_21;
use hexa_protocol_base::{ Packet, PacketType};

pub struct PingRequestPlayPacket{
    pub payload: i64,
    pub protocol_version:i32
}


impl Packet for PingRequestPlayPacket {
    fn get_packet_id(&self,protocol_version:i32) -> i32 {
        match protocol_version {
            767 => 0x21,
            _ => 0x21
        }
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    
}

impl PingRequestPlayPacket{

    pub fn new(payload:i64,protocol_version:i32) -> PingRequestPlayPacket{
        PingRequestPlayPacket{
            protocol_version,
            payload
        }
    }

    pub fn read_packet(reader: &mut BytesMut,protocol_version:i32) ->PingRequestPlayPacket {
        match protocol_version {
            767 => {
                let packet_1_21 = PingRequestPlayPacket1_21::read_packet(reader);
                PingRequestPlayPacket::new(packet_1_21.get_payload(),protocol_version)
            },
            _ => PingRequestPlayPacket::new(0,protocol_version)
            
        }
    }

    pub fn get_payload(&self)-> i64{
        self.payload
    }
}