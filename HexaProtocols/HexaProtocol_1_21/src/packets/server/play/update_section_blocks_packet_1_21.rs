use bytes::BytesMut;

use hexa_protocol_base::{ chunk::chunk_encoder_decoder::ChunkEncoderDecoder, PacketBuilder, PacketReader};

pub struct UpdateSectionBlocksPacket1_21{
    pub chunk_section_x: u32,
    pub chunk_section_y: u32,
    pub chunk_section_z: u32,
    pub blocks_count: i32,
    pub data: Vec<i64>
}



impl UpdateSectionBlocksPacket1_21{

    pub fn new( 
        chunk_section_x: u32,
        chunk_section_y: u32,
        chunk_section_z: u32,
        blocks_count: i32,
        data: Vec<i64>,
        ) -> UpdateSectionBlocksPacket1_21{
        UpdateSectionBlocksPacket1_21{
            chunk_section_x,
            chunk_section_y,
            chunk_section_z,
            blocks_count,
            data
        }
    }

    pub fn read_packet(reader: &mut BytesMut) ->UpdateSectionBlocksPacket1_21 {
        let mut reader = PacketReader::new(reader);
        let chunk_section = reader.read_long_be();
        let blocks_size = reader.read_varint();
        let mut data = Vec::new();
        for _ in 0..blocks_size {
            data.push(reader.read_varlong());
        }
        let (chunk_section_x, chunk_section_y, chunk_section_z) = ChunkEncoderDecoder::decode_chunk_section_position(767, chunk_section);
        UpdateSectionBlocksPacket1_21::new(chunk_section_x,chunk_section_y,chunk_section_z,blocks_size,data)
    }

    pub fn build(&self) -> PacketBuilder {
        let mut writer = PacketBuilder::new(0x49);
        let encode = ChunkEncoderDecoder::encode_chunk_section_position(767, self.chunk_section_x, self.chunk_section_y, self.chunk_section_z);
        writer.write_long_be(encode);
        writer.write_varint(self.blocks_count);
        for i in 0..self.blocks_count {
            writer.write_varlong(self.data[i as usize]);
        }
        writer
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