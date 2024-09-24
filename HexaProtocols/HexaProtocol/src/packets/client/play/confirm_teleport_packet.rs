use bytes::BytesMut;

use hexa_protocol_1_21::packets::client::play::confirm_teleport_packet_1_21::ConfirmTeleportPacket1_21;
use hexa_protocol_base::{ Packet,  PacketType};

pub struct ConfirmTeleportPacket{
    pub tp_id: i32,
    pub protocol_version:i32
}


impl Packet for ConfirmTeleportPacket {
    fn get_packet_id(&self,protocol_version:i32) -> i32 {
        match protocol_version {
            767 => 0x01,
            _ => 0x00
        }
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    
}

impl ConfirmTeleportPacket{

    pub fn new(teleport_id:i32,protocol_version:i32) -> ConfirmTeleportPacket{
        ConfirmTeleportPacket{
            protocol_version,
            tp_id:teleport_id
        }
    }

    pub fn read_packet(reader: &mut BytesMut,protocol_version:i32) ->ConfirmTeleportPacket {
        match protocol_version {
            767 => {
                let packet_1_21 = ConfirmTeleportPacket1_21::read_packet(reader);
                ConfirmTeleportPacket::new(packet_1_21.get_tp_id(),protocol_version)
            },
            _ => ConfirmTeleportPacket::new(0,protocol_version)
            
        }
    }

    pub fn get_tp_id(&self)-> i32{
        self.tp_id
    }
}