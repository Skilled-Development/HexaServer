use std::sync::Arc;

use bytes::BytesMut;
use hexa_protocol_base::PacketBuilder;
use tokio::{io::AsyncWriteExt, net::tcp::OwnedWriteHalf, sync::Mutex};

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
    pub name: Option<String>,
    pub ip_address: String,
    pub port: u16,
    pub client_state: ClientState,
    pub username: Option<String>,
    pub uuid: Option<uuid::Uuid>,
    pub server_config: Option<Arc<std::sync::RwLock<ServerConfig>>>,
    pub protocol_version: Option<i32>,
    pub last_keep_alive: Option<std::time::Instant>,
    pub keep_alive_id: Option<i64>,
    pub sended_blocks: bool,
    pub writer: Arc<Mutex<OwnedWriteHalf>>, // Ahora el writer está en un Arc<Mutex<>>
}

impl PlayerConnection {
    pub fn new(ip: String, port: u16, writer: OwnedWriteHalf) -> PlayerConnection {
        println!("Creating new connection with IP {}", ip);
        PlayerConnection {
            id: None,
            name: None,
            ip_address: ip,
            port,
            client_state: ClientState::HANDSHAKE,
            username: None,
            uuid: None,
            server_config: None,
            protocol_version: None,
            last_keep_alive: None,
            keep_alive_id: None,
            sended_blocks: false,
            writer: Arc::new(Mutex::new(writer)), // Envolver el writer en Arc<Mutex<>>
        }
    }

    /*
    pub fn set_socket_writer(&mut self, writer: WriteHalf<'_>) {
        self.socket_writer = Some(writer);
    }*/
    pub fn is_send_blocks(&self) -> bool {
        self.sended_blocks
    }
    pub fn set_send_blocks(&mut self, sended_blocks: bool) {
        self.sended_blocks = sended_blocks;
    }

    pub fn set_protocol_version(&mut self, protocol_version: i32) {
        self.protocol_version = Some(protocol_version);
    }

    pub fn set_last_keep_alive(&mut self, last_keep_alive: std::time::Instant) {
        self.last_keep_alive = Some(last_keep_alive);
    }
    pub fn set_keep_alive_id(&mut self, keep_alive_id: i64) {
        self.keep_alive_id = Some(keep_alive_id);
    }
    pub fn get_keep_alive_id(&self) -> i64 {
        self.keep_alive_id.clone().unwrap()
    }
    pub fn get_last_keep_alive(&self) -> std::time::Instant {
        self.last_keep_alive.clone().unwrap()
    }
    pub fn get_protocol_version(&self) -> i32 {
        self.protocol_version.clone().unwrap()
    }

    pub fn set_client_state(&mut self, client_state: ClientState) {
        self.client_state = client_state;
    }

    pub fn set_username(&mut self, username: String) {
        self.username = Some(username);
    }

    pub fn get_username(&self) -> String {
        self.username.clone().unwrap()
    }

    pub fn set_server_config(&mut self, server_config: Arc<std::sync::RwLock<ServerConfig>>) {
        self.server_config = Some(server_config);
    }

    pub fn set_uuid(&mut self, uuid: uuid::Uuid) {
        self.uuid = Some(uuid);
    }

    pub fn get_uuid(&self) -> uuid::Uuid {
        self.uuid.clone().unwrap()
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
