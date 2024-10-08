use std::{collections::HashMap, sync::Arc};

use bytes::{Buf, BufMut, BytesMut};
use hexa_protocol_base::PacketBuilder;
use tokio::{
    io::AsyncReadExt,
    net::{tcp::OwnedReadHalf, TcpListener},
    sync::{mpsc::UnboundedSender, Mutex, RwLock},
};

use crate::{
    packet::packet_buffer::PacketBuffer,
    packets_handler::{
        configuration_handler::{
            aknowlodge_finish_configuration, client_information, cookie_request,
            server_bound_configuration, server_bound_known_packs,
        },
        handshake_handler::{handshake, ping_request},
        login_handler::{login_acknowledgement, login_start},
    },
    player::{player::Player, player_connection::ClientState},
    PlayerConnection, ServerConfig,
};

pub struct ProtocolThread {
    pub port: u16,
    pub address: String,
    pub clients: Arc<Mutex<HashMap<String, Arc<Mutex<Player>>>>>,
    pub server_name: String,
    pub server_versions: Vec<i32>,
    pub server_config: Arc<RwLock<ServerConfig>>,
    pub packet_sender: UnboundedSender<PacketBuffer>,
}

impl ProtocolThread {
    pub fn new(
        port: u16,
        address: String,
        server_name: String,
        server_versions: Vec<i32>,
        server_config: Arc<RwLock<ServerConfig>>,
        packet_sender: UnboundedSender<PacketBuffer>,
    ) -> Self {
        let protocol_thread = ProtocolThread {
            port,
            address,
            clients: Arc::new(Mutex::new(HashMap::new())),
            server_name,
            server_versions,
            server_config,
            packet_sender,
        };
        protocol_thread
    }

    pub async fn start(&mut self) {
        let addr_str = format!("{}:{}", self.address, self.port);
        let listener = TcpListener::bind(addr_str).await.unwrap();
        println!(
            "Servidor de Minecraft escuchando en la ip {} , con en el puerto {}...",
            self.address, self.port
        );

        loop {
            let (socket, addr) = listener.accept().await.unwrap();
            let (reader, writer) = socket.into_split(); // Este produce `tokio::net::tcp::WriteHalf`
            let ip_address = addr.ip().to_string();
            let port = addr.port();
            println!("-------------------------------------");
            println!("New connection from {}:{}..", ip_address, port);
            let address = format!("{}:{}", ip_address, port);
            let client = self.clients.lock().await.get(&address).cloned();
            let (connection, client) = match client {
                Some(client) => {
                    let connection = client.lock().await.get_connection();
                    (connection, client.clone())
                }
                None => {
                    let connection = Arc::new(Mutex::new(PlayerConnection::new(
                        ip_address.clone(),
                        port,
                        writer,
                    )));
                    let client = Arc::new(Mutex::new(Player::new(connection.clone())));
                    let connection_clone = connection.clone();
                    let locked = connection_clone.lock().await;
                    self.clients.lock().await.insert(
                        locked.ip_address.clone() + ":" + &locked.port.to_string(),
                        client.clone(),
                    );
                    (connection, client)
                }
            };
            connection
                .lock()
                .await
                .set_server_config(self.server_config.clone());
            let clients = self.clients.clone();
            let client_clone = client.clone();
            let sender = self.packet_sender.clone();
            tokio::spawn({
                let clients_clone = clients.clone();
                async move {
                    let result = Self::handle_client(sender, reader, client, clients_clone).await;
                    if result.is_err() {
                        let mut clients_lock = clients.lock().await;
                        let client_clone = client_clone.lock().await;
                        let deleted_entity_id = client_clone.get_entity_id();
                        if deleted_entity_id == -1 {
                            clients_lock.remove(&address);
                            return;
                        }
                        let connection = client_clone.get_connection();
                        let connection = connection.lock().await;
                        let server_config_lock = connection.get_server_config();
                        let server_config = server_config_lock.read().await;
                        let entity_processor = server_config.get_entity_processor();
                        let entity_processor = entity_processor.lock().await;
                        entity_processor.remove_entity(deleted_entity_id).await;
                        clients_lock.remove(&address);
                        let mut remove_entity_packet = PacketBuilder::new(0x42);
                        remove_entity_packet.write_varint(1);
                        remove_entity_packet.write_varint(deleted_entity_id);

                        for (_client_id, other_client) in clients_lock.iter() {
                            let other_client = other_client.lock().await;
                            let other_connection = other_client.get_connection();
                            let mut other_connection = other_connection.lock().await;
                            other_connection
                                .send_packet_builder(remove_entity_packet.clone())
                                .await;
                        }
                        println!("Client {} deleted from list of clients.", address);
                    }
                }
            });
        }
    }

    pub async fn handle_client(
        packet_sender: UnboundedSender<PacketBuffer>,
        mut reader: OwnedReadHalf,
        client: Arc<Mutex<Player>>,
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<Player>>>>>,
    ) -> Result<(), String> {
        let mut buffer = BytesMut::with_capacity(8192);
        loop {
            match reader.read_buf(&mut buffer).await {
                Ok(0) => {
                    return Err("error".to_string());
                }
                Ok(_) => {
                    while buffer.len() > 0 {
                        match Self::process_packet(
                            packet_sender.clone(),
                            &mut buffer,
                            &mut reader,
                            client.clone(),
                            clients.clone(),
                        )
                        .await
                        {
                            Ok(_) => {
                                continue;
                            }
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
        packet_sender: UnboundedSender<PacketBuffer>,
        buffer: &mut BytesMut,
        reader: &mut OwnedReadHalf,
        client: Arc<Mutex<Player>>,
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<Player>>>>>,
    ) -> Result<(), String> {
        let client_state = {
            let real_client = client.lock().await;
            let connection = real_client.get_connection();
            let connection = connection.lock().await;
            connection.get_client_state()
        };
        if buffer.is_empty() {
            println!("Empty buffer");
            return Ok(());
        }
        if buffer.remaining() < 1 {
            println!("Uncomplete data 1");
            return Err("Datos incompletos".to_string());
        }
        if buffer.is_empty() {
            println!("Empty buffer");
            return Ok(());
        }
        let length = match read_varint(buffer) {
            Ok(len) => len,
            Err(e) => {
                println!("Error trying to read length: {:?}", e);
                return Ok(());
            }
        };
        let packet_id = match read_varint(buffer) {
            Ok(id) => id,
            Err(e) => {
                println!("Error trying to read packet id: {:?}", e);
                return Ok(());
            }
        };
        let result: Result<(), String> = match client_state {
            ClientState::HANDSHAKE => {
                Self::handshake_handler(packet_id, length, buffer, reader, client.clone(), clients)
                    .await
            }
            ClientState::LOGIN => {
                Self::login_handler(packet_id, length, buffer, reader, client.clone()).await
            }
            ClientState::CONFIGURATION => {
                Self::configuration_handler(
                    packet_id,
                    length,
                    buffer,
                    reader,
                    client.clone(),
                    clients,
                )
                .await
            }
            ClientState::PLAY => {
                Self::play_handler(packet_sender, packet_id, length, buffer, client.clone()).await
            }
        };
        if result.is_err() {
            let error = result.unwrap_err();
            if error == "not_enough_data" {
                let mut temp_buffer = BytesMut::with_capacity(1024);
                write_varint(&mut temp_buffer, length);
                write_varint(&mut temp_buffer, packet_id);
                let buffer_clone = buffer.clone();
                buffer.clear();
                buffer.extend_from_slice(&temp_buffer);
                buffer.extend_from_slice(&buffer_clone);
                let mut readed_buffer = BytesMut::with_capacity(1024);
                reader.read_buf(&mut readed_buffer).await.unwrap();
                buffer.extend_from_slice(&readed_buffer);
                return Ok(());
            }
        }
        Ok(())
    }

    pub async fn play_handler(
        packet_sender: UnboundedSender<PacketBuffer>,
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        client: Arc<Mutex<Player>>,
    ) -> Result<(), String> {
        let client_clone = client.clone();
        let packet_data_length = (length as usize) - varint_length(packet_id);
        let buffer_len = buffer.len();

        if packet_data_length > buffer_len {
            return Err("not_enough_data".to_string());
        }

        // Clonar los datos relevantes del buffer para enviarlo a través de packet_sender
        let packet_data = buffer.clone().split_to(packet_data_length as usize);
        let packet_buffer = PacketBuffer::new(
            packet_data,
            packet_id,
            packet_data_length.try_into().unwrap(),
            client_clone.clone(),
        );

        // Enviar el paquete a través de packet_sender
        if let Err(e) = packet_sender.send(packet_buffer) {
            println!("Error sending packet to server process: {:?}", e);
            return Err("Failed to send packet".to_string());
        }
        buffer.clear();
        Ok(())
    }

    pub async fn configuration_handler(
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        reader: &mut OwnedReadHalf,
        client: Arc<Mutex<Player>>,
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<Player>>>>>,
    ) -> Result<(), String> {
        match packet_id {
            0x00 => return client_information::handle(length, buffer, reader, client).await,
            0x01 => return cookie_request::handle(length, buffer, reader, client).await,
            0x02 => {
                return server_bound_configuration::handle(length, buffer, reader, client).await
            }
            0x03 => {
                return aknowlodge_finish_configuration::handle(buffer, reader, client, clients)
                    .await
            }
            0x07 => return server_bound_known_packs::handle(length, buffer, reader, client).await,
            _ => println!("Unknown packet ID: {} in configuration handler", packet_id),
        }
        Ok(())
    }

    pub async fn login_handler(
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        reader: &mut OwnedReadHalf,
        client: Arc<Mutex<Player>>,
    ) -> Result<(), String> {
        match packet_id {
            0x00 => return login_start::handle(length, buffer, reader, client).await,
            0x03 => return login_acknowledgement::handle(length, buffer, reader, client).await,
            _ => println!("Unknown packet ID: {} in login handler", packet_id),
        }
        Ok(())
    }
    pub async fn handshake_handler(
        packet_id: i32,
        length: i32,
        buffer: &mut BytesMut,
        reader: &mut OwnedReadHalf,
        client: Arc<Mutex<Player>>,
        clients: Arc<Mutex<HashMap<String, Arc<Mutex<Player>>>>>,
    ) -> Result<(), String> {
        match packet_id {
            0x00 => {
                return handshake::handle(length, buffer, reader, client, clients.clone()).await
            }
            0x01 => return ping_request::handle(length, buffer, reader, client).await,
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

pub fn write_varint(buffer: &mut BytesMut, mut value: i32) -> &mut BytesMut {
    while (value & 0xFFFFFF80u32 as i32) != 0 {
        buffer.put_u8((value as u8 & 0x7F) | 0x80);
        value >>= 7;
    }
    buffer.put_u8(value as u8);
    buffer
}
pub fn trim_zeros(buffer: &mut BytesMut) {
    // Encuentra el índice del primer byte no nulo
    let start = buffer.iter().position(|&x| x != 0).unwrap_or(0);
    // Encuentra el índice del último byte no nulo
    let end = buffer
        .iter()
        .rposition(|&x| x != 0)
        .unwrap_or(buffer.len() - 1);

    // Recorta el buffer a la parte no vacía
    let trimmed = buffer.split_off(start);
    buffer.clear();
    buffer.extend_from_slice(&trimmed[..=end - start]);
}
fn varint_length(mut value: i32) -> usize {
    let mut length = 0;

    // Mientras haya más bits que procesar, agregamos otro byte.
    while (value & !0x7F) != 0 {
        length += 1;
        value >>= 7;
    }

    // Contar el último byte
    length + 1
}
