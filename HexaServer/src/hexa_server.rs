use std::{
    sync::{Arc, RwLock},
    process,
};
use crate::{Monitor, ProtocolThread};
use hexa_protocol::ServerVersion;

pub struct HexaServer {
    server_config: Arc<RwLock<crate::ServerConfig>>,
    pid: Option<usize>, // Agregamos el campo PID para almacenar el ID del proceso
}

impl HexaServer {
    pub fn new(server_name: String) -> Self {
        HexaServer {
            server_config: Arc::new(RwLock::new(crate::ServerConfig::new(
                25565,
                "localhost".to_string(),
                server_name,
                20,
                "A Minecraft Server".to_string(),
            ))),
            pid: None, 
        }
    }

    pub fn init_pid(&mut self) {
        self.pid = Some(process::id() as usize); 
    }

    

    pub fn add_version(&mut self, version: Arc<dyn ServerVersion + Send + Sync>) {
        let mut config_guard = self.server_config.write().unwrap();
        config_guard.add_version(version);
    }

    pub fn set_server_name(&mut self, server_name: String) {
        let mut config_guard = self.server_config.write().unwrap();
        config_guard.set_server_name(server_name);
    }

    pub fn get_server_name(&self) -> String {
        self.server_config.read().unwrap().get_server_name()
    }



    pub async fn start(&mut self) {
        self.init_pid();
        if self.server_config.read().unwrap().enable_monitoring{
            let mut monitor_thread = Monitor::new(self.pid.unwrap().try_into().unwrap());
            let _monitor_handle = tokio::spawn(async move {
                monitor_thread.start_memory_monitor().await;
            });
        }   
        let versions = self.server_config.read().unwrap().versions.clone();
        if versions.is_empty() {
            println!("No versions available. Shutting down HexaServer...");
            return;
        }

        println!(
            "HexaServer is starting with {} versions...",
            versions.len()
        );

        let versions_vector: Vec<i32> = versions.iter().map(|v| v.protocol()).collect();
        

        let mut protocol_thread = ProtocolThread::new(
            25565,
            "0.0.0.0".to_string(),
            self.server_config.read().unwrap().server_name.clone(),
            versions_vector,
            Arc::clone(&self.server_config),
        );

        

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
