use base64::{engine::general_purpose, Engine as _};
use hexa_protocol_base::ServerVersion;
use reqwest::blocking::get;
use std::error::Error;
use std::sync::{Arc, Mutex};
use std::thread;

pub struct ServerConfig {
    pub server_port: u16,
    pub server_ip: String,
    pub server_name: String,
    pub motd: String,
    pub versions: Vec<Arc<dyn ServerVersion + Send + Sync>>,
    pub versions_protocol: Vec<i32>,
    pub enable_monitoring: bool,
    pub server_icon_url: Option<String>,
    pub server_icon_base64: Arc<Mutex<Option<String>>>,
    pub player_count: i32,
    pub max_player_count: i32,
    pub sample_text: Option<Vec<String>>,
}

impl ServerConfig {
    pub fn new(server_name: String, server_port: u16, server_ip: String) -> Self {
        let config = ServerConfig {
            server_port,
            server_ip,
            server_name,
            motd: "&a&lHexaServer: §fThe §e§lUltimate §4§lRust-Powered §fMinecraft Experience"
                .to_string(),
            versions: Vec::new(),
            versions_protocol: Vec::new(),
            enable_monitoring: false,
            server_icon_url: Some("https://i.imgur.com/jQpVKY7.png".to_string()),
            server_icon_base64: Arc::new(Mutex::new(None)),
            player_count: -1,
            max_player_count: 2024,
            sample_text: Some(vec![
                "§6HexaServer - §eEngineered for §aEfficiency ⚙️".to_string(),
                "§bBuilt with Rust for ultimate performance.".to_string(),
                "§3Fast, scalable, and §csecure server solutions.".to_string(),
                "§9Empower your server with cutting-edge technology.".to_string(),
                "§eHexaServer: Where §dspeed meets reliability.".to_string(),
            ]),
        };
        config.update_server_icon_base64();
        config
    }

    pub fn update_server_icon_base64(&self) {
        let server_icon_url = self.server_icon_url.clone();

        // Usar Arc y Mutex para manejar el estado compartido
        let server_icon_base64 = Arc::clone(&self.server_icon_base64);

        thread::spawn(move || match server_icon_url.as_deref() {
            Some(url) => match Self::download_image_as_base64(url) {
                Ok(base64_string) => {
                    let mut icon_base64 = server_icon_base64.lock().unwrap();
                    *icon_base64 = Some(base64_string);
                }
                Err(e) => {
                    println!("Failed to download server icon: {}", e);
                }
            },
            None => {
                println!("Server icon URL is not set.");
            }
        });
    }

    pub fn get_server_name(&self) -> String {
        self.server_name.clone()
    }

    pub fn set_server_name(&mut self, server_name: String) {
        self.server_name = server_name;
    }

    pub fn add_version(&mut self, version: Arc<dyn ServerVersion + Send + Sync>) {
        self.versions.push(version);
        self.update_versions_protocol();
    }

    pub fn update_versions_protocol(&mut self) {
        self.versions_protocol = self
            .versions
            .iter()
            .map(|version| version.protocol())
            .collect();
    }

    pub fn get_protocol_versions_array(&self) -> Vec<i32> {
        self.versions_protocol.clone()
    }

    pub fn set_server_icon_url(&mut self, server_icon_url: String) {
        self.server_icon_url = Some(server_icon_url);
        self.update_server_icon_base64();
    }

    fn download_image_as_base64(url: &str) -> Result<String, Box<dyn Error>> {
        let response = get(url)?;

        if !response.status().is_success() {
            return Err(format!("Failed to download image: HTTP {}", response.status()).into());
        }

        let image_bytes = response.bytes()?;
        let base64_string = general_purpose::STANDARD.encode(&image_bytes);

        Ok(base64_string)
    }
}
