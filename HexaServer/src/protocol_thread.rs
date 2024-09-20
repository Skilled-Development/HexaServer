use std::{collections::HashMap, sync::Arc};

use bytes::{Buf, BytesMut};
use tokio::{io::AsyncReadExt, net::{TcpListener, TcpStream}, sync::Mutex};

use crate::{packets_handler::{configuration_handler::{client_information, cookie_request, server_bound_configuration, server_bound_known_packs}, handshake_handler::{handshake, ping_request}, login_handler::{login_acknowledgement, login_start}}, player_connection::ClientState, PlayerConnection, ServerConfig};

pub struct ProtocolThread{
    pub port: u16,
    pub address: String,
    clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>,
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
            clients: Arc::new(Mutex::new(HashMap::new())),
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
            let address = format!("{}:{}", ip_address, port);
            let client = self.clients.lock().await.get(&address).cloned();
            let client = match client {
                Some(client) => client,
                None => {
                    let client = Arc::new(Mutex::new(PlayerConnection::new(
                        ip_address.clone(),
                        port,
                    )));
                    self.clients.lock().await.insert(client.lock().await.ip_address.clone(), client.clone());
                    client
                }
            };

            client.lock().await.set_server_config(self.server_config.clone());
            let clients = self.clients.clone(); 
            tokio::spawn(async move {
                let result = Self::handle_client(socket, client).await;
                if result.is_err() {
                    let mut clients_lock = clients.lock().await;
                    clients_lock.remove(&address);
                    println!("Cliente {} eliminado de la lista de clientes.", address);
                }
            });
        }
    }
    


  
    pub async fn handle_client(mut socket: TcpStream,  client: Arc<Mutex<PlayerConnection>>) -> Result<(), String> {
        let mut buffer = BytesMut::with_capacity(1024);
    
        loop {
            match socket.read_buf(&mut buffer).await {
                Ok(0) => {
                    return Err("error".to_string());
                }
                Ok(_) => {
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
                            Err(e) if e == "Datos incompletos" => break,
                            Err(e) => {
                                println!("Error al procesar el paquete: {:?}", e);
                                return Err("error".to_string());
                            }
                        }
                    }
                }
                Err(e) => {
                    println!("Error al leer del socket: {:?}", e);
                    return Err("error".to_string());
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
            0x00 => return client_information::handle(length,buffer,socket,client).await,
            0x01 => return cookie_request::handle(length,buffer,socket,client).await,
            0x02 => return server_bound_configuration::handle(length,buffer,socket,client).await,
            0x07 => return server_bound_known_packs::handle(length,buffer,socket,client).await,
            _ => println!("Unknown packet ID: {} in configuration handler", packet_id),
        }
        Ok(())
    }

    pub async fn login_handler( packet_id: i32, length: i32,buffer: &mut BytesMut,socket: &mut TcpStream, client: &mut PlayerConnection,  ) -> Result<(), String> {
        match packet_id{
            0x00 => return login_start::handle(length,buffer,socket,client).await,
            0x03 => return login_acknowledgement::handle(length,buffer,socket,client).await,
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
            0x00 => return handshake::handle(length,buffer,socket,client).await,
            0x01 => return ping_request::handle(length,buffer,socket,client).await,
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

