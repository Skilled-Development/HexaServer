use bytes::BytesMut;

use hexa_protocol_base::{Packet, PacketBuilder, PacketReader, PacketType};

pub struct PingResponsePacket{
    pub ping_payload: i64
}

impl Packet for PingResponsePacket {
    fn get_packet_id(&self) -> i32 {
        0x01 // 1
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::SERVER
    }
    
}
impl PingResponsePacket{

    pub fn new(ping_payload:i64) -> PingResponsePacket{
        PingResponsePacket{
            ping_payload
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->PingResponsePacket {
        let mut reader = PacketReader::new(reader);
        let ping_payload = reader.read_long_be();
        PingResponsePacket::new(ping_payload)
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer = PacketBuilder::new(self.get_packet_id());
        writer.write_long_be(self.ping_payload);
        writer
    }
}