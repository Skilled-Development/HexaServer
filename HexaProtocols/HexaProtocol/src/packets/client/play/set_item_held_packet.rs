use bytes::BytesMut;

use hexa_protocol_1_21::packets::client::play::set_item_held_packet_1_21::SetItemHeldPacket1_21;
use hexa_protocol_base::{ Packet, PacketType};

pub struct SetItemHeldPacket{
    pub slot: i16,
    pub protocol_version:i32
}


impl Packet for SetItemHeldPacket {
    fn get_packet_id(&self,protocol_version:i32) -> i32 {
        match protocol_version {
            767 => 0x2F,
            _ => 0x2F
        }
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    
}

impl SetItemHeldPacket{

    pub fn new(slot:i16,protocol_version:i32) -> SetItemHeldPacket{
        SetItemHeldPacket{
            slot,
            protocol_version
        }
    }

    pub fn read_packet(reader: &mut BytesMut,protocol_version:i32) ->SetItemHeldPacket {
        match protocol_version {
            767 => {
                let packet_1_21 = SetItemHeldPacket1_21::read_packet(reader);
                SetItemHeldPacket::new(packet_1_21.get_slot(),protocol_version)
            },
            _ => SetItemHeldPacket::new(0,protocol_version)
            
        }
    }

    pub fn get_slot(&self)-> i16{
        self.slot
    }
}