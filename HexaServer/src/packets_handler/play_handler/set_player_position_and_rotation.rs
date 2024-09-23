use bytes::{Buf, BytesMut};
use hexa_protocol_base::PacketReader;
use tokio::net::TcpStream;

use crate::PlayerConnection;

pub async fn handle(length: i32, buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
    let _ = socket;
    let _ = client;
    let _ = length;

    if buffer.remaining() < 33 as usize {
        println!("Not enough data to read set player position and rotation packet");
        println!("Buffer remaining: {}, Length: {}", buffer.remaining(), length);
        //buffer.clear();
        return Err("not_enough_data".to_string());
    }
    let mut reader = PacketReader::new(buffer);
    let x = reader.read_double();
    let y = reader.read_double();
    let z = reader.read_double();
    let yaw = reader.read_float();
    let pitch = reader.read_float();
    let mut on_ground = false;
    if reader.buf.remaining() >= 1 {
        on_ground = reader.read_boolean();  
    }
    println!("x: {}, y: {}, z: {}, yaw: {}, pitch: {}, on_ground: {}", x, y, z, yaw, pitch, on_ground);
    Ok(())
}