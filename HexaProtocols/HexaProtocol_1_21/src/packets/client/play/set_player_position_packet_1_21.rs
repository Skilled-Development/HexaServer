use bytes::{Buf, BytesMut};

use hexa_protocol_base::{ PacketBuilder, PacketReader};

pub struct SetPlayerPositionPacket1_21{
    pub x: f64,
    pub y: f64,
    pub z: f64,
    pub on_ground: bool,
}



impl SetPlayerPositionPacket1_21{

    pub fn new(x:f64,y:f64,z:f64,on_ground:bool) -> SetPlayerPositionPacket1_21{
        SetPlayerPositionPacket1_21{
            x,y,z,on_ground
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->SetPlayerPositionPacket1_21 {
        let mut reader = PacketReader::new(reader);
        let x = reader.read_double();
        let y = reader.read_double();
        let z = reader.read_double();
        let mut on_ground = false;
        if reader.buf.remaining() >= 1 {
            on_ground = reader.read_boolean();  
        }
        SetPlayerPositionPacket1_21 {
            x,
            y,
            z,
            on_ground
        }
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer = PacketBuilder::new(0x1A);
        writer.write_double(self.x);
        writer.write_double(self.y);
        writer.write_double(self.z);
        writer.write_boolean(self.on_ground);
        writer
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