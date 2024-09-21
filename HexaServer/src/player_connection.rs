use std::sync::Arc;

use crate::ServerConfig;


#[derive(Eq, Hash, PartialEq, Debug,Clone, Copy)]  // Add Debug trait to ClientState
pub enum ClientState {
    HANDSHAKE,
    LOGIN,
    CONFIGURATION,
    PLAY,
}

#[derive(Clone)]
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

}

impl PlayerConnection {
    pub fn new(ip: String, port: u16) -> PlayerConnection {
        println!("Creating new connection with IP {}", ip);
        PlayerConnection {
            id: None,
            name: None,
            ip_address: ip,
             port,
            client_state: ClientState::HANDSHAKE,
            username:None,
            uuid:None,
            server_config: None,
            protocol_version: None,
        }
    }

    pub fn set_protocol_version(&mut self, protocol_version: i32) {
        self.protocol_version = Some(protocol_version);
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


    
}
