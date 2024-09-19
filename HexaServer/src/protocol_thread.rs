use std::{collections::HashMap, sync::Arc};

use bytes::{Buf, BytesMut};
use tokio::{io::AsyncReadExt, net::{TcpListener, TcpStream}, sync::Mutex};

use crate::{packets_handler::{configuration_handler::{client_information_handler, cookie_request_handler, server_bound_configuration_handler}, handshake_handler::{handshake_handler, ping_request_handler}, login_handler::{login_acknowledgement_handler, login_start_handler}}, player_connection::ClientState, PlayerConnection, ServerConfig};

pub struct ProtocolThread{
    pub port: u16,
    pub address: String,
    clients: std::collections::HashMap<String, Arc<Mutex<PlayerConnection>>>,
    pub server_name: String,
    pub server_versions: Vec<i32>,
    pub server_config: Arc<std::sync::RwLock<ServerConfig>>,
}

impl ProtocolThread{
    pub fn new(
        port: u16,
        address: String,
        server_name: String,
        server_versions: Vec<i32>,
        server_config: Arc<std::sync::RwLock<ServerConfig>>,
    ) -> ProtocolThread {
        let protocol_thread = ProtocolThread {
            port,
            address,
            clients: HashMap::new(),
            server_name,
            server_versions,
            server_config
        };
        protocol_thread
    }


    pub async fn start(&mut self) {
        let addr_str = format!("{}:{}", self.address, self.port);
        let listener = TcpListener::bind(addr_str).await.unwrap();
        println!("Servidor de Minecraft escuchando en la ip {} , con en el puerto {}...", self.address, self.port);
    
        loop {
            let (socket, addr) = listener.accept().await.unwrap();
            let ip_address = addr.ip().to_string();
            let port = addr.port();
            println!("-------------------------------------");
            println!("Nueva conexión de {}:{}...", ip_address, port);
            
            // Aquí, addr ya contiene la IP y el puerto del cliente
            let client = self.clients.get(&ip_address).cloned();
            let client = match client {
                Some(client) => client,
                None => {
                    let client = Arc::new(Mutex::new(PlayerConnection::new(
                        ip_address.clone(),
                        self.server_name.clone(),
                        self.server_versions.clone(),
                    )));
                    self.clients.insert(client.lock().await.ip_address.clone(), client.clone());
                    client
                }
            };
    
            tokio::spawn(async move {
                Self::handle_client(socket, client).await;
            });
        }
    }
    
    pub async fn handle_client(mut socket: TcpStream,  client: Arc<Mutex<PlayerConnection>>) {
        let mut buffer = BytesMut::with_capacity(1024);
    
        // Loop principal para manejar la conexión con el cliente
        loop {
            match socket.read_buf(&mut buffer).await {
                Ok(0) => {
                    // Conexión cerrada por el cliente
                    return;
                }
                Ok(_) => {
                    // Procesar los paquetes si hay suficientes datos
                    while buffer.len() > 0 {
                        println!("===============================================");
                        println!("Procesando paquete...");
                        println!("Buffer: {:?}", buffer);
                        let mut client_guard = client.lock().await;
                        println!("Client state: {:?}", client_guard.client_state);
                        match Self::process_packet(&mut buffer, &mut socket, &mut client_guard).await {
                            Ok(_) => {
                                println!("Buffer después de procesar: {:?}", buffer);
                                continue;
                            },
                            Err(e) if e == "Datos incompletos" => break, // Esperar más datos si no están completos
                            Err(e) => {
                                println!("Error al procesar el paquete: {:?}", e);
                                return;
                            }
                        }
                    }
                }
                Err(e) => {
                    println!("Error al leer del socket: {:?}", e);
                    return;
                }
            }
        }
    }
    
    async fn process_packet(buffer: &mut BytesMut, socket: &mut TcpStream, client: &mut PlayerConnection) -> Result<(), String> {
        if buffer.remaining() < 1 {
            return Err("Datos incompletos".to_string());
        }
        let length = match read_varint(buffer) {
            Ok(len) => len,
            Err(e) => return Err(e),
        };
        if buffer.remaining() < length as usize {
            return Err("Datos incompletos".to_string());
        }
        let packet_id = read_varint(buffer)?;
        println!("Packet ID: {}", packet_id);
        match client.client_state {
            ClientState::HANDSHAKE => Self::handshake_handler(packet_id,length,buffer,socket,  client).await?,
            ClientState::LOGIN => Self::login_handler(packet_id,length,buffer,socket,  client).await?,
            ClientState::CONFIGURATION => Self::configuration_handler(packet_id,length,buffer,socket,  client).await?,
            ClientState::PLAY => todo!(),
        }

    
        Ok(())
    }

    pub async fn configuration_handler(
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        socket: &mut TcpStream,
        client: &mut PlayerConnection, 
        ) -> Result<(), String> {
        match packet_id{
            0x00 => return client_information_handler::handle(length,buffer,socket,client).await,
            0x01 => return cookie_request_handler::handle(length,buffer,socket,client).await,
            0x02 => return server_bound_configuration_handler::handle(length,buffer,socket,client).await,
            _ => println!("Unknown packet ID: {} in configuration handler", packet_id),
        }
        Ok(())
    }

    pub async fn login_handler( packet_id: i32, length: i32,buffer: &mut BytesMut,socket: &mut TcpStream, client: &mut PlayerConnection,  ) -> Result<(), String> {
        match packet_id{
            0x00 => return login_start_handler::handle(length,buffer,socket,client).await,
            0x03 => return login_acknowledgement_handler::handle(length,buffer,socket,client).await,
            _ => println!("Unknown packet ID: {} in login handler", packet_id),
        }
        Ok(())
    }
    pub async fn handshake_handler(
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        socket: &mut TcpStream,
        client: &mut PlayerConnection, 
        ) -> Result<(), String> {
        match packet_id{
            0x00 => return handshake_handler::handle(length,buffer,socket,client).await,
            0x01 => return ping_request_handler::handle(length,buffer,socket,client).await,
            _ => println!("Unknown packet ID: {} in handshake_handler", packet_id),
        }
        Ok(())
    }
    

}

fn read_varint(buffer: &mut BytesMut) -> Result<i32, String> {
    if buffer.is_empty() {
        return Err("Datos incompletos: Buffer vacío".to_string());
    }

    let mut result = 0;
    let mut shift = 0;
    loop {
        if buffer.is_empty() {
            return Err("Buffer vacío durante la lectura de VarInt".to_string());
        }

        let byte = buffer.get_u8();
        result |= ((byte & 0x7F) as i32) << shift;
        if byte & 0x80 == 0 {
            break;
        }
        shift += 7;
        if shift > 35 {
            return Err("VarInt demasiado grande".to_string());
        }
    }
    Ok(result)
}

