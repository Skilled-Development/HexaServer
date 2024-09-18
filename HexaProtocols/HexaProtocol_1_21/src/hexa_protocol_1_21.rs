
use hexa_protocol::ServerVersion;

pub struct HexaProtocol1_21 {
    version: String,
    protocol: i32,
}

impl HexaProtocol1_21 {
    // Método para crear una nueva instancia de HexaProtocol1_21 con valores "hardcodeados"
    pub fn new() -> Self {
        HexaProtocol1_21 {
            version: "1.21".to_string(),
            protocol: 767,
        }
    }
}

impl ServerVersion for HexaProtocol1_21 {
    fn version(&self) -> String {
        self.version.clone()  // Retorna el valor del campo version
    }

    fn protocol(&self) -> i32 {
        self.protocol.to_string().parse().unwrap()  // Retorna el valor del campo protocol convertido a i32
    }

    fn get_version_info(&self) -> String {
        format!("Version: {} - Protocol: {}", self.version(), self.protocol())
    }
}
