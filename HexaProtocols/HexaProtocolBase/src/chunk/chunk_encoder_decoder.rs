pub struct ChunkEncoderDecoder;

impl ChunkEncoderDecoder {
    pub fn encode_chunk_section_position(protocol_version:i32,section_x: u32, section_y: u32, section_z: u32) -> i64 {
        match protocol_version{
            767 => {
                let encoded = ((section_x & 0x3FFFFF) as u64) << 42 |
                      (section_y & 0xFFFFF) as u64 |
                      ((section_z & 0x3FFFFF) as u64) << 20;
                encoded.try_into().unwrap()
            },
            _ => 0
        }
    }
    pub fn decode_chunk_section_position(protocol_version:i32,encoded: i64) -> (u32, u32, u32) {
        match protocol_version{
            767 => {
                let section_x = ((encoded >> 42) & 0x3FFFFF) as u32;
                let section_y = (encoded & 0xFFFFF) as u32;
                let section_z = ((encoded >> 20) & 0x3FFFFF) as u32;
                (section_x, section_y, section_z)
            },
            _ => (0, 0, 0)
        }
    }

    pub fn encode_local_block(protocol_version:i32,block_state_id: u32, local_x: u32, local_y: u32, local_z: u32) -> i64 {
        match protocol_version{
            767 => {
                let encoded = (block_state_id as u64) << 12 |
                      ((local_x & 0xF) as u64) << 8 |
                      ((local_z & 0xF) as u64) << 4 |
                      (local_y & 0xF) as u64;
                encoded.try_into().unwrap()
            },
            _ => 0
            
        }
    }

    pub fn decode_local_block(protocol_version:i32,encoded: u64) -> (u32, u32, u32, u32) {
        match protocol_version {
            767 =>{
            let block_state_id = (encoded >> 12) as u32;
            let local_x = ((encoded >> 8) & 0xF) as u32;
            let local_y = (encoded & 0xF) as u32;
            let local_z = ((encoded >> 4) & 0xF) as u32;

            (block_state_id, local_x, local_y, local_z)
            },
            _ => (0, 0, 0, 0)
        }
    }
}
