use bytes::BytesMut;

use hexa_protocol_base::{ PacketBuilder, PacketReader};

pub struct ConfirmTeleportPacket1_21{
    pub tp_id: i32,
}



impl ConfirmTeleportPacket1_21{

    pub fn new(teleport_id:i32) -> ConfirmTeleportPacket1_21{
        ConfirmTeleportPacket1_21{
            tp_id:teleport_id
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->ConfirmTeleportPacket1_21 {
        let mut reader = PacketReader::new(reader);
        let tp_id = reader.read_varint();
        ConfirmTeleportPacket1_21::new(tp_id)
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer = PacketBuilder::new(0x01);
        writer.write_varint(self.tp_id);
        writer
    }

    pub fn get_tp_id(&self)-> i32{
        self.tp_id
    }
}