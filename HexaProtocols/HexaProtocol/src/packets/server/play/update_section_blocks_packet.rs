use bytes::BytesMut;

use hexa_protocol_1_21::packets::server::play::update_section_blocks_packet_1_21::UpdateSectionBlocksPacket1_21;
use hexa_protocol_base::{ Packet, PacketBuilder, PacketType};

pub struct UpdateSectionBlocksPacket{
    pub chunk_section_x: u32,
    pub chunk_section_y: u32,
    pub chunk_section_z: u32,
    pub blocks_count: i32,
    pub data: Vec<i64>,
    pub protocol_version:i32
}


impl Packet for UpdateSectionBlocksPacket {
    fn get_packet_id(&self,protocol_version:i32) -> i32 {
        match protocol_version {
            767 => 0x49,
            _ => 0x49
        }
    }
    fn get_packet_type(&self) -> PacketType{
        PacketType::SERVER
    }
    
}

impl UpdateSectionBlocksPacket{

    pub fn new(
        chunk_section_x: u32,
        chunk_section_y: u32,
        chunk_section_z: u32,
        blocks_count: i32,
        data: Vec<i64>,
        protocol_version:i32
    ) -> UpdateSectionBlocksPacket{
        UpdateSectionBlocksPacket{
            chunk_section_x,
            chunk_section_y,
            chunk_section_z,
            blocks_count,
            data,
            protocol_version
        }
    }

    pub fn read_packet(reader: &mut BytesMut,protocol_version:i32) ->UpdateSectionBlocksPacket {
        match protocol_version {
            767 => {
                let packet_1_21 = UpdateSectionBlocksPacket1_21::read_packet(reader);
                UpdateSectionBlocksPacket::new(
                    packet_1_21.get_chunk_section_x(),
                    packet_1_21.get_chunk_section_y(),
                    packet_1_21.get_chunk_section_z(),
                    packet_1_21.get_blocks_count(),
                    packet_1_21.get_data(),
                    767
                )
            },
            _ => UpdateSectionBlocksPacket::new(0,0,0,0,Vec::new(),0)
            
        }
    }

    pub fn build(&self) -> PacketBuilder {
        match self.protocol_version {
            767 => {
                let packet_1_21 = UpdateSectionBlocksPacket1_21::new(
                    self.chunk_section_x,
                    self.chunk_section_y,
                    self.chunk_section_z,
                    self.blocks_count,
                    self.data.clone()
                );
                packet_1_21.build()
            },
            _ => PacketBuilder::new(0x49)
        }
    }

    pub async fn send(&self,socket: &mut tokio::net::TcpStream) {
        let mut packet = self.build();
        let _ = packet.send(socket).await;
    }

    pub fn get_chunk_section_x(&self)-> u32{
        self.chunk_section_x
    }
    pub fn get_chunk_section_y(&self)-> u32{
        self.chunk_section_y
    }
    pub fn get_chunk_section_z(&self)-> u32{
        self.chunk_section_z
    }
    pub fn get_blocks_count(&self)-> i32{
        self.blocks_count
    }
    pub fn get_data(&self)-> Vec<i64>{
        self.data.clone()
    }
}