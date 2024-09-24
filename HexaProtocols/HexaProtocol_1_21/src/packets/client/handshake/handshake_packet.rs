use bytes:: BytesMut;

use hexa_protocol_base::{packets::{packet::Packet, packet_type::PacketType}, PacketBuilder, PacketReader};

// Import the PlayerConnection type if it is defined in another module

#[derive(Debug)]
pub struct HandshakePacket {
    pub protocol_version: i32,
    pub server_address: String,
    pub server_port: u16,
    pub next_state: i32,
}

impl Packet for HandshakePacket {
    fn get_packet_id(&self,protocol_version:i32) -> i32 {
        match protocol_version {
            767=> 0x00,
            _=> 0x00
        }
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    

}
impl HandshakePacket {


    pub fn read_packet(packet_data: &mut BytesMut) ->HandshakePacket {
        let mut reader = PacketReader::new(packet_data);
        let protocol_version = reader.read_varint();
        let server_address = reader.read_string();
        let server_port = reader.read_unsigned_short();
        let next_state = reader.read_varint();
        let handshaking_packet = HandshakePacket::new(protocol_version, server_address, server_port, next_state);


        handshaking_packet
    }


    pub fn new(protocol_version: i32, server_address: String, server_port: u16, next_state: i32) -> HandshakePacket {
        HandshakePacket {
            protocol_version,
            server_address,
            server_port,
            next_state
        }
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer = PacketBuilder::new(0);
        writer.write_varint(self.protocol_version);
        writer.write_string(self.server_address.as_str());
        writer.write_unsigned_short(self.server_port);
        writer.write_varint(self.next_state);
        writer
    }

    pub fn get_next_state(&self)-> i32{
        self.next_state
    }

    pub fn get_player_protocol(&self )-> i32{
        self.protocol_version
    }


}



