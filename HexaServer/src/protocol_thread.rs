use std::{collections::HashMap, sync::Arc};
use chrono::Utc;

use bytes::{Buf, BytesMut};
use hexa_protocol::{packet_builder, PacketBuilder};
use tokio::{io::{AsyncReadExt, AsyncWriteExt}, net::{TcpListener, TcpStream}, sync::Mutex};

use crate::{packets_handler::{configuration_handler::{aknowlodge_finish_configuration, client_information, cookie_request, server_bound_configuration, server_bound_known_packs}, handshake_handler::{handshake, ping_request}, login_handler::{login_acknowledgement, login_start}, play_handler::{confirm_teletransportation, keep_alive, pick_item, ping_request_play, set_item_held, set_player_position, set_player_position_and_rotation, swing_arm}}, player_connection::ClientState, PlayerConnection, ServerConfig};
use rand::Rng;

pub struct ProtocolThread{
    pub port: u16,
    pub address: String,
    pub clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>,
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
            println!("New connection from {}:{}..", ip_address, port);
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
            tokio::spawn({
                let clients_clone = clients.clone();  /*Este clients_clone en cuestion */
                async move {
                    let result = Self::handle_client(socket, client, clients_clone).await;
                    if result.is_err() {
                        let mut clients_lock = clients.lock().await;
                        clients_lock.remove(&address);
                        println!("Client {} deleted from list of clients.", address);
                    }
                }
            });
        }
    }
    


  
    pub async fn handle_client(mut socket: TcpStream,  client: Arc<Mutex<PlayerConnection>>,clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>) -> Result<(), String> {
        let mut buffer = BytesMut::with_capacity(1024);
    
        loop {
            match socket.read_buf(&mut buffer).await {
                Ok(0) => {
                    return Err("error".to_string());
                }
                Ok(_) => {
                    while buffer.len() > 0 {
                        println!("===============================================");
                        let mut client_guard = client.lock().await;
                        match Self::process_packet(&mut buffer, &mut socket, &mut client_guard,clients.clone()).await {
                            Ok(_) => {
                                println!("Buffer after processing: {:?}", buffer);
                                continue;
                            },
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

    
    async fn process_packet(
        buffer: &mut BytesMut, 
        socket: &mut TcpStream, 
        client: &mut PlayerConnection,
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>
    ) -> Result<(), String> {
        println!("Processing packet...");
        println!("Buffer: {:?}", buffer);
        let client_state = client.client_state.clone();
        println!("Client state: {:?}", client_state);
        if buffer.is_empty() {
        
            println!("Empty buffer");
            buffer.clear();
            return Ok(());
        }
        if buffer.remaining() < 1 {
            println!("Uncomplete data 1");
            buffer.clear();
            return Err("Datos incompletos".to_string());
        }
        if buffer.is_empty() {
            println!("Empty buffer");
            buffer.clear();
            return Ok(());
        }
        let length = match read_varint(buffer) {
            Ok(len) => len,
            Err(e) => {
                println!("Error trying to read length: {:?}", e);
                buffer.clear();
                return Ok(())
            },
        };
        let packet_id = match read_varint(buffer) {
            Ok(id) => id,
            Err(e) => {
                println!("Error trying to read packet id: {:?}", e);
                buffer.clear();
                return Ok(())
            },
        
        };
        println!("Packet ID: {}", packet_id);
        println!("Packet ID: 0x{:X}", packet_id);
        /*if buffer.remaining() < length as usize {
            buffer.clear();
            return Ok(());
        }*/
        match client_state {
            ClientState::HANDSHAKE => Self::handshake_handler(packet_id,length,buffer,socket,  client,clients).await?,
            ClientState::LOGIN => Self::login_handler(packet_id,length,buffer,socket,  client,clients).await?,
            ClientState::CONFIGURATION => Self::configuration_handler(packet_id,length,buffer,socket,  client,clients).await?,
            ClientState::PLAY => Self::play_handler(packet_id,length,buffer,socket,  client,clients).await?,
        }

    
        Ok(())
    }

    pub async fn play_handler(
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        socket: &mut TcpStream,
        client: &mut PlayerConnection, 
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>
        ) -> Result<(), String> {
            let _ = clients;

            if packet_id > 100 {
                println!("Unknown packet ID: {} in play handler", packet_id);
                buffer.clear();
                return Ok(());
            }
            let result_on_read = match packet_id{
                0x00 =>  confirm_teletransportation::handle(length,buffer,socket,client).await,
                0x21 =>  ping_request_play::handle(length,buffer,socket,client).await,
                0x1A =>  set_player_position::handle(length,buffer,socket,client).await,
                0x1B =>  set_player_position_and_rotation::handle(length,buffer,socket,client).await,
                0x2F =>  set_item_held::handle(length,buffer,socket,client).await,
                0x36 =>  swing_arm::handle(length,buffer,socket,client).await,
                0x18 =>  keep_alive::handle(length,buffer,socket,client).await,
                0x20 => pick_item::handle(length,buffer,socket,client).await,
                0x3D => {
                    println!("Idk what to do with this packet");
                    buffer.clear();
                    socket.flush().await.unwrap();
                    Ok(())
                }
                0x57=>{
                    println!("Idk what to do with this packet");
                    buffer.clear();
                    socket.flush().await.unwrap();
                    Ok(())
                }
                0x53 =>{
                    println!("Idk what to do with this packet");
                    buffer.clear();
                    socket.flush().await.unwrap();
                    Ok(())
                }
                0x54 =>{
                    println!("Idk what to do with this packet");
                    buffer.clear();
                    socket.flush().await.unwrap();
                    Ok(())
                }
                0x40 => {
                    println!("Idk what to do with this packet");
                    buffer.clear();
                    socket.flush().await.unwrap();
                    Ok(())
                }
                0x10 =>{
                    println!("Idk what to do with this packet");
                    buffer.clear();
                    socket.flush().await.unwrap();
                    Ok(())
                }
                _ => todo!("Unknown packet ID: {} in play handler", packet_id),
            };

            if result_on_read.is_ok(){
                let client_last_keep_alive = client.get_last_keep_alive();
                if client_last_keep_alive.elapsed().as_millis() > 17000 {
                    println!("Client {} timed out", client.ip_address);
                    let mut keep_alive_packet = PacketBuilder::new(0x26);
                    let random_id: i64 = 346092730i64;//rand::thread_rng().gen();
                    keep_alive_packet.write_long_be(random_id);
                    keep_alive_packet.send(socket).await?;
                    client.set_keep_alive_id(random_id);
                    println!("Sent keep alive packet to client {} with alive id {}", client.ip_address, random_id);
                }
            }else{
                return Err("Error al procesar el paquete".to_string());
            }
            Ok(())
        }

    pub async fn configuration_handler(
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        socket: &mut TcpStream,
        client: &mut PlayerConnection, 
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>
        ) -> Result<(), String> {
        let _ = clients;
    
        match packet_id{
            0x00 => return client_information::handle(length,buffer,socket,client).await,
            0x01 => return cookie_request::handle(length,buffer,socket,client).await,
            0x02 => return server_bound_configuration::handle(length,buffer,socket,client).await,
            0x03 => return aknowlodge_finish_configuration::handle(length,buffer,socket,client).await,
            0x07 => return server_bound_known_packs::handle(length,buffer,socket,client).await,
    
            _ => println!("Unknown packet ID: {} in configuration handler", packet_id),
        }
        Ok(())
    }

    pub async fn login_handler( 
        packet_id: i32, 
        length: i32,
        buffer: &mut BytesMut,
        socket: &mut TcpStream, 
        client: &mut PlayerConnection, 
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>
    ) -> Result<(), String> {
        let _ = clients;
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
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<PlayerConnection>>>>>
        ) -> Result<(), String> {
        match packet_id{
            0x00 => return handshake::handle(length,buffer,socket,client,clients.clone()).await,
            0x01 => return ping_request::handle(length,buffer,socket,client).await,
            _ => println!("Unknown packet ID: {} in handshake_handler", packet_id),
        }
        Ok(())
    }
    

}

pub fn read_varint(buffer: &mut BytesMut) -> Result<i32, String> {
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

