

use bytes::BytesMut;

pub fn read_data_file_to_bytesmut(file_name: &str) -> BytesMut {
    // Incluye el archivo como bytes en tiempo de compilación
    //let data: &'static [u8] = include_bytes!("src/packets_files/registry/{}", file_name);
    match file_name{
        "trim_material.data" => {
            let data: &'static [u8] = include_bytes!("saved/trim_material.data" );
            BytesMut::from(data)
        }
        "trim_pattern.data" => {
            let data: &'static [u8] = include_bytes!("saved/trim_pattern.data" );
            BytesMut::from(data)
        },
        "banner_pattern.data" => {
            let data: &'static [u8] = include_bytes!("saved/banner_pattern.data" );
            BytesMut::from(data)
        }
        "biome.data" => {
            let data: &'static [u8] = include_bytes!("saved/biome.data" );
            BytesMut::from(data)
        }
        "chat_type.data" => {
            let data: &'static [u8] = include_bytes!("saved/chat_type.data" );
            BytesMut::from(data)
        }
        "damage_type.data" => {
            let data: &'static [u8] = include_bytes!("saved/damage_type.data" );
            BytesMut::from(data)
        }
        "dimension_type.data" => {
            let data: &'static [u8] = include_bytes!("saved/dimension_type.data" );
            BytesMut::from(data)
            /*let mut packet = PacketBuilder::new(0x07);
            packet.write_identifier("minecraft:dimension_type".to_string());
            packet.write_varint(0);
            packet.build()*/
        }
        "wolf_variant.data" => {
            let data: &'static [u8] = include_bytes!("saved/wolf_variant.data" );
            BytesMut::from(data)
        },
        "painting_variant.data" => {
            let data: &'static [u8] = include_bytes!("saved/painting_variant.data" );
            BytesMut::from(data)
        },
        _ => {
            let data: &'static [u8] = include_bytes!("saved/trim_material.data" );
            BytesMut::from(data)
        }
    }
}