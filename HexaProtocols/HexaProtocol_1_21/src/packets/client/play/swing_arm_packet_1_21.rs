use bytes::BytesMut;

use hexa_protocol_base::{ PacketBuilder, PacketReader};

pub struct SwingArmPacket1_21{
    pub hand: i32,
}

impl SwingArmPacket1_21{

    pub fn new(hand:i32) -> SwingArmPacket1_21{
        SwingArmPacket1_21{
            hand
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->SwingArmPacket1_21 {
        let mut reader = PacketReader::new(reader);
        let hand = reader.read_varint   ();
        SwingArmPacket1_21::new(hand)
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer: PacketBuilder = PacketBuilder::new(0x36);
        writer.write_varint(self.hand);
        writer
    }

    pub fn get_hand(&self) -> i32 {
        self.hand
    }
}