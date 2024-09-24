use bytes::BytesMut;

use hexa_protocol_1_21::packets::client::play::swing_arm_packet_1_21::SwingArmPacket1_21;
use hexa_protocol_base::{ player::hand::Hand, Packet, PacketType};

pub struct SwingArmPacket{
    pub hand: i32,
    pub protocol_version:i32
}


impl Packet for SwingArmPacket {
    fn get_packet_id(&self,protocol_version:i32) -> i32 {
        match protocol_version {
            767 => 0x36,
            _ => 0x36
        }
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    
}

impl SwingArmPacket{

    pub fn new(hand:i32,protocol_version:i32) -> SwingArmPacket{
        SwingArmPacket{
            hand,
            protocol_version
        }
    }

    pub fn read_packet(reader: &mut BytesMut,protocol_version:i32) ->SwingArmPacket {
        match protocol_version {
            767 => {
                let packet_1_21 = SwingArmPacket1_21::read_packet(reader);
                SwingArmPacket::new(packet_1_21.get_hand(),protocol_version)
            },
            _ => SwingArmPacket::new(0,protocol_version)
            
        }
    }

    pub fn get_hand_number(&self)-> i32{
        self.hand
    }
    pub fn get_hand(&self)-> Hand{
        match self.protocol_version{
            767 => {
                match self.hand{
                    0 => Hand::MainHand,
                    1 => Hand::OffHand,
                    _ => Hand::MainHand
                }
            },
            _ => Hand::MainHand
        }
    }
}