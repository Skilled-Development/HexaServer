use std::sync::Arc;

use hexa_protocol::ServerVersion;

pub struct ServerConfig{
    pub server_port: u16,
    pub server_ip: String,
    pub server_name: String,
    pub max_players: i32,
    pub motd: String,
    pub versions: Vec<Arc<dyn ServerVersion + Send + Sync>>,
    pub versions_protocol: Vec<i32>,
    pub enable_monitoring: bool,
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
            versions_protocol: Vec::new(),
            enable_monitoring: false,
        }
    }

    pub fn get_server_name(&self) -> String {
        self.server_name.clone()
    }
    pub fn set_server_name(&mut self, server_name: String){
        self.server_name = server_name;
    }
    pub fn add_version(&mut self, version: Arc<dyn ServerVersion + Send + Sync>) {
        self.versions.push(version);
        self.update_versions_protocol();
    }
    pub fn update_versions_protocol(&mut self){
        self.versions_protocol = self.versions.iter().map(|version| version.protocol()).collect();
    }

    pub fn get_protocol_versions_array(&self) -> Vec<i32> {
        self.versions_protocol.clone()
    }
    
}