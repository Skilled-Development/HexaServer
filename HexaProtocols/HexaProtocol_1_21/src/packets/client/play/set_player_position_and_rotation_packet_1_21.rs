use bytes::{Buf, BytesMut};

use hexa_protocol_base::{ PacketBuilder, PacketReader};

pub struct SetPlayerPositionAndRotationPacket1_21{
    pub x: f64,
    pub y: f64,
    pub z: f64,
    pub yaw: f32,
    pub pitch: f32,
    pub on_ground: bool,
}



impl SetPlayerPositionAndRotationPacket1_21{

    pub fn new(x:f64,y:f64,z:f64, yaw: f32, pitch: f32,on_ground:bool) -> SetPlayerPositionAndRotationPacket1_21{
        SetPlayerPositionAndRotationPacket1_21{
            x,y,z,yaw,pitch,on_ground
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->SetPlayerPositionAndRotationPacket1_21 {
        let mut reader = PacketReader::new(reader);
        let x = reader.read_double();
        let y = reader.read_double();
        let z = reader.read_double();
        let yaw = reader.read_float();
        let pitch = reader.read_float();
        let mut on_ground = false;
        if reader.buf.remaining() >= 1 {
            on_ground = reader.read_boolean();  
        }
        SetPlayerPositionAndRotationPacket1_21 {
            x,
            y,
            z,
            yaw,
            pitch,
            on_ground
        }
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer = PacketBuilder::new(0x1B);
        writer.write_double(self.x);
        writer.write_double(self.y);
        writer.write_double(self.z);
        writer.write_float(self.yaw);
        writer.write_float(self.pitch);
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

    pub fn get_yaw(&self)-> f32{
        self.yaw
    }
    pub fn get_pitch(&self)-> f32{
        self.pitch
    }
    
    
}