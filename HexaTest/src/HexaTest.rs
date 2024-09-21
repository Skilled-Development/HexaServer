use bytes::{Buf, BytesMut};
use hexa_protocol::PacketReader;
use hexa_server::HexaServer;
use hexa_protocol_1_21::HexaProtocol1_21;
use std::sync::Arc;
#[tokio::main]
async fn main() {
   /*  let byte_string: &[u8] = &[
        0x11, 0x0, 0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x4D, 0x69, 0x6E, 0x65, 0x63,
        0x72, 0x61, 0x66, 0x74, 0x21, 0x40, 0x24, 0x27, 0x52, 0x71, 0x3D, 0x28, 0x40, 0x97,
        0xA7, 0x7B, 0x1D, 0x40, 0x66, 0x79, 0xF1, 0x5B, 0x24,
    ];*/

    // Convertir a BytesMut
    let byte_data: &[u8] = b"0\x01\x11Hello, Minecraft!@\x0fminecraft:stone\x00d@\t!\xfbTD.\xea\xe8\x07";

    // Crea un BytesMut a partir del byte array
    let mut bytes_mut = BytesMut::from(byte_data);
    //bytes_mut.extend_from_slice(byte_string);

    // Imprimir el resultado
    println!("{:?}", bytes_mut);
    let length = read_varint(&mut bytes_mut);
    println!("Length: {}", length);
    let packet_id = read_varint(&mut bytes_mut);
    println!("Packet ID: {}", packet_id);
    let mut reader = PacketReader::new(&mut bytes_mut);
    let string = reader.read_string();
    println!("String: {}", string);
    let angle = reader.read_angle();
    println!("Angle: {}", angle);
    let identifier = reader.read_identifier();
    println!("Identifier: {}", identifier);
    let short = reader.read_unsigned_short();
    println!("Short: {}", short);
    let double = reader.read_double();
    println!("Double: {}", double);
    let int = reader.read_int();
    println!("Int: {}", int);

    
   /* // Create an instance of HexaProtocol1_21
    let protocol_1_21 = HexaProtocol1_21::new();

    // Create an instance of HexaServer
    let mut server = HexaServer::new("HexaServer".to_string());

    // Add the version to the server
    server.add_version(Arc::new(protocol_1_21));
    // Start the server
    server.start().await;*/

    
}
fn read_varint(buffer: &mut BytesMut) -> i32 {
    if buffer.is_empty() {
        return -1;
    }

    let mut result = 0;
    let mut shift = 0;
    loop {
        if buffer.is_empty() {
            return -1;
        }

        let byte = buffer.get_u8();
        result |= ((byte & 0x7F) as i32) << shift;
        if byte & 0x80 == 0 {
            break;
        }
        shift += 7;
        if shift > 35 {
            return -1;
        }
    }
    result
}