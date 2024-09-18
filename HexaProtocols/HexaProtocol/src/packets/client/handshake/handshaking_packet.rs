use bytes:: BytesMut;

use crate::{packets::{packet::Packet, packet_type::PacketType}, PacketReader};

// Import the PlayerConnection type if it is defined in another module

#[derive(Debug)]
pub struct HandshakingPacket {
    pub protocol_version: i32,
    pub server_address: String,
    pub server_port: u16,
    pub next_state: i32,
}

impl Packet for HandshakingPacket {
    fn get_packet_id(&self) -> i32 {
        0 
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    

}
impl HandshakingPacket {


        pub fn read_packet(packet_data: &mut BytesMut) -> (&mut BytesMut,HandshakingPacket) {
            let mut reader = PacketReader::new(packet_data);
            let protocol_version = reader.read_varint();
            let server_address = reader.read_string();
            let server_port = reader.read_unsigned_short();
            let next_state = reader.read_varint();
            let handshaking_packet = HandshakingPacket::new(protocol_version, server_address, server_port, next_state);


            (reader.buf,handshaking_packet)
        }

    

    pub fn new(protocol_version: i32, server_address: String, server_port: u16, next_state: i32) -> HandshakingPacket {
        HandshakingPacket {
            protocol_version,
            server_address,
            server_port,
            next_state
        }
    }
    pub fn get_next_state(&self)-> i32{
        self.next_state
    }

    pub fn get_player_protocol(&self )-> i32{
        self.protocol_version
    }


}



