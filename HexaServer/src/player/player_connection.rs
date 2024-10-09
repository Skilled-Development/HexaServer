use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_base::PacketBuilder;
use tokio::{
    io::AsyncWriteExt,
    net::tcp::OwnedWriteHalf,
    sync::{Mutex, RwLock},
};

use crate::ServerConfig;

#[derive(Eq, Hash, PartialEq, Debug, Clone, Copy)] // Add Debug trait to ClientState
pub enum ClientState {
    HANDSHAKE,
    LOGIN,
    CONFIGURATION,
    PLAY,
}

pub struct PlayerConnection {
    pub id: Option<String>,
    pub ip_address: String,
    pub port: u16,
    pub client_state: ClientState,
    pub server_config: Option<Arc<RwLock<ServerConfig>>>,
    pub sended_blocks: bool,
    pub writer: Arc<Mutex<OwnedWriteHalf>>, // Ahora el writer está en un Arc<Mutex<>>
}

impl PlayerConnection {
    pub fn new(ip: String, port: u16, writer: OwnedWriteHalf) -> PlayerConnection {
        println!("Creating new connection with IP {}", ip);
        PlayerConnection {
            id: None,
            ip_address: ip,
            port,
            client_state: ClientState::HANDSHAKE,
            server_config: None,
            sended_blocks: false,
            writer: Arc::new(Mutex::new(writer)), // Envolver el writer en Arc<Mutex<>>
        }
    }

    pub fn get_connection_id(&self) -> String {
        self.ip_address.clone() + ":" + &self.port.to_string()
    }

    pub fn get_server_config(&self) -> Arc<RwLock<ServerConfig>> {
        self.server_config.clone().unwrap()
    }
    pub fn is_send_blocks(&self) -> bool {
        self.sended_blocks
    }
    pub fn set_send_blocks(&mut self, sended_blocks: bool) {
        self.sended_blocks = sended_blocks;
    }

    pub fn set_client_state(&mut self, client_state: ClientState) {
        self.client_state = client_state;
    }

    pub fn set_server_config(&mut self, server_config: Arc<RwLock<ServerConfig>>) {
        self.server_config = Some(server_config);
    }

    pub fn get_client_state(&self) -> ClientState {
        self.client_state
    }

    pub async fn send_packet_bytes(&mut self, packet: BytesMut) {
        self.writer.lock().await.write_all(&packet).await.unwrap();
    }

    pub async fn send_packet_builder(&mut self, mut packet: PacketBuilder) {
        self.writer
            .lock()
            .await
            .write_all(&packet.build())
            .await
            .unwrap();
    }
}
