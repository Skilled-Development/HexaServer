use bytes::BytesMut;

use crate::{Packet, PacketBuilder, PacketReader, PacketType};

pub struct DataRegistryPacket{
}

impl Packet for DataRegistryPacket {
    fn get_packet_id(&self) -> i32 {
        0x01 // 1
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::SERVER
    }
    
}
impl DataRegistryPacket{

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