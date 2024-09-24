use bytes::BytesMut;

use hexa_protocol_1_21::packets::client::play::set_player_position_packet_1_21::SetPlayerPositionPacket1_21;
use hexa_protocol_base::{ Packet, PacketType};

pub struct SetPlayerPositionPacket{
    pub x: f64,
    pub y: f64,
    pub z: f64,
    pub on_ground: bool,
    pub protocol_version:i32
}


impl Packet for SetPlayerPositionPacket {
    fn get_packet_id(&self,protocol_version:i32) -> i32 {
        match protocol_version {
            767 => 0x1A,
            _ => 0x1A
        }
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::CLIENT
    }
    
}

impl SetPlayerPositionPacket{

    pub fn new(x:f64,y:f64,z:f64,on_ground:bool,protocol_version:i32) -> SetPlayerPositionPacket{
        SetPlayerPositionPacket{
            x,y,z,on_ground,protocol_version
        }
    }

    pub fn read_packet(reader: &mut BytesMut,protocol_version:i32) ->SetPlayerPositionPacket {
        match protocol_version {
            767 => {
                let packet_1_21 = SetPlayerPositionPacket1_21::read_packet(reader);
                SetPlayerPositionPacket::new(packet_1_21.get_x(),packet_1_21.get_y(),packet_1_21.get_z(),packet_1_21.get_on_ground(),protocol_version)
            },
            _ => SetPlayerPositionPacket::new(0f64,100f64,0f64,false,protocol_version)
            
        }
    }

    pub fn get_x(&self)-> f64{
        self.x
    }
    pub fn get_y(&self)-> f64{
        self.y
    }
    pub fn get_z(&self)-> f64{
        self.z
    }
    pub fn get_on_ground(&self)-> bool{
        self.on_ground
    }
    
}