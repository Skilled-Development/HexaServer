use bytes::BytesMut;

use hexa_protocol_base::{ PacketBuilder, PacketReader};

pub struct SetItemHeldPacket1_21{
    pub slot: i16,
}



impl SetItemHeldPacket1_21{

    pub fn new(slot:i16) -> SetItemHeldPacket1_21{
        SetItemHeldPacket1_21{
            slot
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->SetItemHeldPacket1_21 {
        let mut reader = PacketReader::new(reader);
        let slot = reader.read_short();
        SetItemHeldPacket1_21::new(slot)
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer: PacketBuilder = PacketBuilder::new(0x2F);
        writer.write_short(self.slot);
        writer
    }

    pub fn get_slot(&self) -> i16 {
        self.slot
    }
}