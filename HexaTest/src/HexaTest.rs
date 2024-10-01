use hexa_protocol_1_21::HexaProtocol1_21;
use hexa_server::HexaServer;
use std::sync::Arc;
#[tokio::main]
async fn main() {
    // Create an instance of HexaProtocol1_21
    let protocol_1_21 = HexaProtocol1_21::new();

    // Create an instance of HexaServer
    let mut server = HexaServer::new("HexaServer".to_string());

    // Add the version to the server
    server.add_version(Arc::new(protocol_1_21)).await;
    // Start the server
    server.start().await;
}
