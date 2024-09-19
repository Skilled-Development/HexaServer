use std::sync::Arc;

use hexa_protocol::ServerVersion;

pub struct ServerConfig{
    pub server_port: u16,
    pub server_ip: String,
    pub server_name: String,
    pub max_players: i32,
    pub motd: String,
    pub versions: Vec<Arc<dyn ServerVersion + Send + Sync>>,
}

impl ServerConfig{
    pub fn new(server_port: u16, server_ip: String, server_name: String, max_players: i32, motd: String) -> Self{
        ServerConfig{
            server_port,
            server_ip,
            server_name,
            max_players,
            motd,
            versions: Vec::new(),
        }
    }
    
}