use bytes::{Buf, BytesMut};
use hexa_protocol::packets::client::play::confirm_teleport_packet::ConfirmTeleportPacket;
use hexa_protocol_base::{chunk, PacketBuilder};
use tokio::net::TcpStream;
use tokio::task;

use crate::PlayerConnection;

pub async fn handle(
    length: i32,
    socket: &mut TcpStream,
    buffer: &mut BytesMut,
    client: &mut PlayerConnection,
) -> Result<(), String> {
    if buffer.remaining() < length as usize {
        return Err("not_enough_data".to_string());
    }
    ConfirmTeleportPacket::read_packet(buffer, client.get_protocol_version());

    let mut center_packet = PacketBuilder::new(0x54);
    center_packet.write_varint(0);
    center_packet.write_varint(0);
    center_packet.send(socket).await?;

    // Lanzar la tarea principal
    /*task::spawn(async move {
        for x in -20..=20 {
            for z in -20..=20 {
                generate_chunk_data_packet(socket, x, z).await;
            }
        }
    });*/

    Ok(())
}
async fn generate_chunk_data_packet(socket: &mut TcpStream, chunk_x: i32, chunk_y: i32) {
    let mut packet = PacketBuilder::new(0x27);
    packet.write_int(chunk_x);
    packet.write_int(chunk_y);
    packet.write_bytes([0x0a, 0x00]); //empty nbt heightmaps
                                      // A byte buffer where we'll store all the Chunk Section information
    let mut chunk_section: Vec<u8> = Vec::new();
    // Step 1: Store the number of blocks that are not air (4096 total blocks)
    // Since there are 4096 stone blocks, we store that value as a "short" (u16) of 2 bytes.
    chunk_section.extend_from_slice(&4096u16.to_be_bytes()); // Add 4096 as big-endian
                                                             // Step 2: Bits per entry = 0, because all blocks are stone (single value palette)
                                                             // Add 1 byte that represents 0 bits per entry
    chunk_section.push(0u8); // Represents the "Bits per entry"
                             // Step 3: Add the only palette value (stone) to the palette.
                             // We assume that the stone ID in the palette is 1.
    chunk_section.push(1u8); // Add the stone block ID to the palette (1 byte)
                             // Step 4: The data array is not added because it’s unnecessary with a single value palette.
                             // If we were using more than one block, we would need to store data here.
    chunk_section.push(0u8); // Data array size is 0 for a single value palette
                             // Step 5: Repeat the process for biomes.
                             // We assume there’s only one biome, and we use a single value palette for biomes as well.
    chunk_section.push(0u8); // 0 bits per entry (for biomes)
    chunk_section.push(1u8); // Add a value for the biome (ID 1 in the palette)
                             // The data array is also 0 for biomes, since there is only one biome.
    chunk_section.push(0u8); // Data array size is 0 for biomes.
                             //add the chunk data
    let chunk_sections = chunk_section.repeat(24);
    packet.write_varint(chunk_sections.len() as i32);
    packet.write_byte_array_no_length_prefixed(&chunk_sections);
    // Number of block entities
    packet.write_varint(0);
    // Sky light mask (empty)
    packet.write_varint(0);
    // Block light mask (empty)
    packet.write_varint(0);
    // Empty sky light mask (empty)
    packet.write_varint(0);
    // Empty block light mask (empty)
    packet.write_varint(0);
    // Sky light array count
    packet.write_varint(0);
    // Block light array count
    packet.write_varint(0);
    packet.send(socket).await.unwrap();
}
