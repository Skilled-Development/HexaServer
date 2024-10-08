use crate::Player;
use bytes::BytesMut;
use std::sync::Arc;
use tokio::sync::Mutex;

pub struct PacketBuffer {
    pub buffer: BytesMut,
    pub packet_id: i32,
    pub packet_length: i32,
    pub client: Arc<Mutex<Player>>,
}

impl PacketBuffer {
    pub fn new(
        buffer: BytesMut,
        packet_id: i32,
        packet_length: i32,
        client: Arc<Mutex<Player>>,
    ) -> PacketBuffer {
        PacketBuffer {
            buffer,
            packet_id,
            packet_length,
            client,
        }
    }

    // Obtener una referencia al buffer
    pub fn get_buffer(&self) -> &BytesMut {
        &self.buffer
    }

    // Obtener una referencia mutable al buffer
    pub fn get_mut_buffer(&mut self) -> &mut BytesMut {
        &mut self.buffer
    }

    pub fn get_packet_id(&self) -> i32 {
        self.packet_id
    }

    pub fn get_packet_length(&self) -> i32 {
        self.packet_length
    }

    pub fn get_client(&self) -> Arc<Mutex<Player>> {
        self.client.clone()
    }
}
