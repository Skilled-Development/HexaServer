use hexa_server::HexaServer;
use hexa_protocol_1_21::HexaProtocol1_21;
use std::{env, sync::Arc};

#[tokio::main]
async fn main() {
    //env::set_var("RUST_BACKTRACE", "1");
    // Crear una instancia de HexaProtocol1_21
    let protocol_1_21 = HexaProtocol1_21::new();

    // Crear una instancia de HexaServer
    let mut server = HexaServer::new("HexaServer".to_string());

    // Agregar la versión 1.21 al servidor
    server.add_version(Arc::new(protocol_1_21));

    // Iniciar el servidor
    server.start().await;

    
}
