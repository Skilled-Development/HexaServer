use bytes::BytesMut;

use hexa_protocol_base::{ PacketBuilder, PacketReader};

pub struct PingRequestPlayPacket1_21{
    pub payload: i64
}


impl PingRequestPlayPacket1_21{

    pub fn new(payload:i64) -> PingRequestPlayPacket1_21{
        PingRequestPlayPacket1_21{
            payload
        }
    }

    pub fn read_packet(buffer: &mut BytesMut) ->PingRequestPlayPacket1_21 {
        let mut reader = PacketReader::new(buffer);
        let payload = reader.read_long_be();
        PingRequestPlayPacket1_21::new(payload)
    }

    pub fn build(&self)-> PacketBuilder{
        let mut builder = PacketBuilder::new(0x21);
        builder.write_long_be(self.payload);
        builder
    }

    pub fn get_payload(&self)-> i64{
        self.payload
    }
}