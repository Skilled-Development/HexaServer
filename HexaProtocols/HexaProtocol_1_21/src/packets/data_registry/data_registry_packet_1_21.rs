use std::{fs::{self, File}, io::{self, Read}, path::{Path, PathBuf}};

use bytes::BytesMut;

pub fn read_data_file_to_bytesmut(file_name: &str) -> BytesMut {
    // Incluye el archivo como bytes en tiempo de compilación
    //let data: &'static [u8] = include_bytes!("src/packets_files/registry/{}", file_name);
    match file_name{
        "trim_material.data" => {
            let data: &'static [u8] = include_bytes!( "trim_material.data" );
            BytesMut::from(data)
        }
        "trim_pattern.data" => {
            let data: &'static [u8] = include_bytes!( "trim_pattern.data" );
            BytesMut::from(data)
        },
        "banner_pattern.data" => {
            let data: &'static [u8] = include_bytes!( "banner_pattern.data" );
            BytesMut::from(data)
        }
        "biome.data" => {
            let data: &'static [u8] = include_bytes!( "biome.data" );
            BytesMut::from(data)
        }
        "chat_type.data" => {
            let data: &'static [u8] = include_bytes!( "chat_type.data" );
            BytesMut::from(data)
        }
        "damage_type.data" => {
            let data: &'static [u8] = include_bytes!( "damage_type.data" );
            BytesMut::from(data)
        }
        "dimension_type.data" => {
            let data: &'static [u8] = include_bytes!( "dimension_type.data" );
            BytesMut::from(data)
        }
        "wolf_variant.data" => {
            let data: &'static [u8] = include_bytes!( "wolf_variant.data" );
            BytesMut::from(data)
        }
        _ => {
            let data: &'static [u8] = include_bytes!( "trim_material.data" );
            BytesMut::from(data)
        }
    }
}

fn list_files_in_directory(path: &Path) {
    match fs::read_dir(path) {
        Ok(entries) => {
            for entry in entries {
                match entry {
                    Ok(entry) => {
                        let entry_path = entry.path();

                        if entry_path.is_dir() {
                            // Si es un directorio, imprimimos el nombre y hacemos una llamada recursiva
                            println!("Directory: {:?}", entry_path);
                            list_files_in_directory(&entry_path);
                        } else {
                            // Si es un archivo, imprimimos su nombre
                            println!("File: {:?}", entry_path);
                        }
                    }
                    Err(e) => println!("Error reading entry: {:?}", e),
                }
            }
        }
        Err(e) => println!("Could not read directory {:?}: {:?}", path, e),
    }
}