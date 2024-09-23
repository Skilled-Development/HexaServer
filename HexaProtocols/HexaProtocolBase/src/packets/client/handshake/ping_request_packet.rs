use bytes::BytesMut;

use crate::{Packet, PacketBuilder, PacketReader, PacketType};

pub struct PingRequestPacket{
    pub ping_payload: i64
}

impl Packet for PingRequestPacket {
    fn get_packet_id(&self) -> i32 {
        0x01 // 1
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    
}
impl PingRequestPacket{

    pub fn new(ping_payload:i64) -> PingRequestPacket{
        PingRequestPacket{
            ping_payload
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->PingRequestPacket {
        let mut reader = PacketReader::new(reader);
        let ping_payload = reader.read_long_be();
        PingRequestPacket::new(ping_payload)
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer = PacketBuilder::new(self.get_packet_id());
        writer.write_long_be(self.ping_payload);
        writer
    }

    pub fn get_payload(&self)-> i64{
        self.ping_payload
    }
}