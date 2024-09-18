use std::sync::Arc;
use crate::ProtocolThread;
use hexa_protocol::ServerVersion;
use sysinfo::System;

#[derive(Clone)]
pub struct HexaServer {
    pub versions: Vec<Arc<dyn ServerVersion + Send + Sync>>,
    pub server_name: String,
}

impl HexaServer {
    pub fn new(server_name: String) ->Self {
        HexaServer {
            server_name,
            versions: Vec::new(),
        }
    }

    pub fn add_version(&mut self, version: Arc<dyn ServerVersion + Send + Sync>) {
        self.versions.push(version);
    }

    pub fn set_server_name(&mut self, server_name: String) {
        self.server_name = server_name;
    }

    pub fn get_server_name(&self) -> String {
        self.server_name.clone()
    }

    pub async fn start(&mut self) {
        if self.versions.is_empty() {
            println!("No versions available. Shutting down HexaServer...");
            return;
        }
    
        println!(
            "HexaServer is starting with {} versions...",
            self.versions.len()
        );
    
        let _system = System::new_all();
        let _pid = std::process::id(); // PID del proceso actual
        let mut versions_vector: Vec<i32> = Vec::new();
    
        for version in &self.versions {
            versions_vector.push(version.protocol());
        }
    
        let mut protocol_thread = ProtocolThread::new(
            25565,
            "0.0.0.0".to_string(),
            self.server_name.clone(),
            versions_vector,
        );
        
        // Spawning the protocol thread
        let protocol_handle = tokio::spawn(async move {
            protocol_thread.start().await;
        });
    
        tokio::select! {
            _ = protocol_handle => {
                println!("Protocol thread has finished.");
            }
        }
    
        println!("HexaServer has stopped.");
    }
    
}
