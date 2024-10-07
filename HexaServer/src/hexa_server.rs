use crate::{Monitor, ProtocolThread, ServerProcess};
use bytes::BytesMut;
use hexa_protocol_base::ServerVersion;
use std::{process, sync::Arc};
use tokio::{
    runtime::Handle,
    sync::{mpsc, Mutex, RwLock},
};

pub struct HexaServer {
    server_config: Arc<RwLock<crate::ServerConfig>>,
    pid: Option<usize>, // Agregamos el campo PID para almacenar el ID del proceso
}

impl HexaServer {
    pub fn new(server_name: String) -> Self {
        HexaServer {
            server_config: Arc::new(RwLock::new(crate::ServerConfig::new(
                server_name,
                25565,
                "0.0.0.0".to_string(),
            ))),
            pid: None,
        }
    }

    pub fn init_pid(&mut self) {
        self.pid = Some(process::id() as usize);
    }

    pub async fn add_version(&mut self, version: Arc<dyn ServerVersion + Send + Sync>) {
        let mut config_guard = self.server_config.write().await;
        config_guard.add_version(version);
    }

    pub fn set_server_name(&mut self, server_name: String) {
        let mut config_guard = Handle::current().block_on(self.server_config.write());
        config_guard.set_server_name(server_name);
    }

    pub fn get_server_name(&self) -> String {
        let config_guard = Handle::current().block_on(self.server_config.write());
        config_guard.get_server_name()
    }

    pub async fn start(&mut self) {
        self.init_pid();
        let server_config = self.server_config.read().await;
        if server_config.enable_monitoring {
            let mut monitor_thread = Monitor::new(self.pid.unwrap().try_into().unwrap());
            let _monitor_handle = tokio::spawn(async move {
                monitor_thread.start_memory_monitor().await;
            });
        }
        let versions = server_config.versions.clone();
        if versions.is_empty() {
            println!("No versions available. Shutting down HexaServer...");
            return;
        }

        println!("HexaServer is starting with {} versions...", versions.len());

        let versions_vector: Vec<i32> = versions.iter().map(|v| v.protocol()).collect();
        let (tx, rx): (
            mpsc::UnboundedSender<BytesMut>,
            mpsc::UnboundedReceiver<BytesMut>,
        ) = mpsc::unbounded_channel();
        let mut protocol_thread = ProtocolThread::new(
            server_config.server_port,
            server_config.server_ip.clone(),
            server_config.server_name.clone(),
            versions_vector,
            Arc::clone(&self.server_config),
            tx,
        );

        let mut server_process = ServerProcess {
            packet_receiver: Arc::new(Mutex::new(rx)),
            packets: Arc::new(Mutex::new(Vec::new())),
        };

        // Ejecutar ambos en paralelo
        tokio::join!(protocol_thread.start(), server_process.run(),);

        println!("HexaServer has stopped.");
    }
}
