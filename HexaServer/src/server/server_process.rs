use bytes::Buf;
use crab_nbt::{Nbt, NbtCompound};
use hexa_protocol_base::PacketBuilder;
use rand::Rng;
use std::{collections::HashMap, sync::Arc, time::Duration};
use tokio::{
    sync::{mpsc::UnboundedReceiver, Mutex, RwLock},
    time::Instant,
};

use crate::{
    packet::packet_buffer::PacketBuffer,
    packets_handler::play_handler::{
        confirm_teletransportation, keep_alive, ping_request_play, set_player_position,
        set_player_position_and_rotation, set_player_rotation, swing_arm,
    },
    Player, ServerConfig,
};

pub struct ServerProcess {
    pub packet_receiver: Arc<Mutex<UnboundedReceiver<PacketBuffer>>>,
    pub packets: Arc<Mutex<Vec<PacketBuffer>>>,
    pub server_config: Arc<RwLock<ServerConfig>>,
    pub broadcast_packet_map: HashMap<Arc<Mutex<Player>>, Vec<PacketBuilder>>,
    pub tick_number: i32,
    pub mspt_list: Vec<f64>,
}

impl ServerProcess {
    pub fn new(
        packet_receiver: Arc<Mutex<UnboundedReceiver<PacketBuffer>>>,
        server_config: Arc<RwLock<ServerConfig>>,
    ) -> ServerProcess {
        ServerProcess {
            packet_receiver,
            packets: Arc::new(Mutex::new(Vec::new())),
            server_config,
            broadcast_packet_map: HashMap::new(),
            tick_number: 0,
            mspt_list: Vec::new(),
        }
    }

    pub async fn run(mut self) {
        self.start_packet_receiver().await;
        loop {
            let start = Instant::now();

            // Procesamos los paquetes
            if let Err(e) = self.process_packets().await {
                eprintln!("Error procesando paquetes: {:?}", e);
            }
            // Calculamos el tiempo transcurrido
            let elapsed = start.elapsed();
            self.mspt_list.push(elapsed.as_secs_f64() * 1000.0);

            // Cada 20 ticks, calculamos el TPS y lo transmitimos
            if self.tick_number == 19 {
                self.tick_number = 0;
                self.broadcast_mspt_and_tps().await;
            } else {
                self.tick_number += 1;
            }
            // Dormimos lo que queda de los 50ms de cada tick
            let remaining_time = Duration::from_millis(50).saturating_sub(elapsed);
            if !remaining_time.is_zero() {
                tokio::time::sleep(remaining_time).await;
            }
        }
    }

    async fn start_packet_receiver(&mut self) {
        let packet_receiver = Arc::clone(&self.packet_receiver);
        let packets = Arc::clone(&self.packets);
        tokio::spawn(async move {
            let mut receiver = packet_receiver.lock().await;
            while let Some(packet) = receiver.recv().await {
                packets.lock().await.push(packet);
            }
        });
    }

    pub async fn process_packets(&self) -> Result<(), String> {
        let packets = Arc::clone(&self.packets);
        let mut packets = packets.lock().await;

        if packets.is_empty() {
            return Ok(());
        }

        let packets_to_process = std::mem::take(&mut *packets);
        drop(packets);

        let mut futures = vec![];
        for packet in packets_to_process {
            futures.push(self.handle_packet(packet));
        }

        futures::future::join_all(futures).await;

        Ok(())
    }

    async fn handle_packet(&self, mut packet: PacketBuffer) -> Result<(), String> {
        let packet_id = packet.get_packet_id();
        let length = packet.get_packet_length();
        let client = Arc::clone(&packet.get_client());
        let buffer = packet.get_mut_buffer();
        if buffer.remaining() < length as usize {
            return Err("not_enough_data".to_string());
        }
        let _ = length;
        let result_on_read = match packet_id {
            0x00 => confirm_teletransportation::handle(buffer, client.clone()).await,
            0x21 => ping_request_play::handle(buffer, client.clone()).await,
            0x1A => set_player_position::handle(buffer, client.clone(), self).await,
            0x1B => set_player_position_and_rotation::handle(buffer, client.clone(), self).await,
            0x36 => swing_arm::handle(buffer, client.clone(), self).await,
            0x18 => keep_alive::handle(buffer, client.clone()).await,
            0x1C => set_player_rotation::handle(buffer, client.clone(), self).await,
            _ => {
                println!("Unknown packet ID: 0x{:x}", packet_id);
                buffer.clear();
                return Ok(());
            }
        };

        if result_on_read.is_err() {
            return result_on_read;
        }

        self.handle_keep_alive(client).await;
        Ok(())
    }

    async fn handle_keep_alive(&self, client: Arc<Mutex<Player>>) {
        let mut client_guard = client.lock().await;
        if client_guard.get_last_keep_alive().elapsed().as_millis() > 17000 {
            let connection = client_guard.get_connection();
            let mut connection_guard = connection.lock().await;

            let mut keep_alive_packet = PacketBuilder::new(0x26);
            let random_id: i64 = rand::thread_rng().gen();
            keep_alive_packet.write_long_be(random_id);

            connection_guard
                .send_packet_builder(keep_alive_packet)
                .await;
            client_guard.set_keep_alive_id(random_id);
            client_guard.set_last_keep_alive(Instant::now().into());
        }
    }

    pub async fn broadcast_mspt_and_tps(&mut self) {
        let median_mspt = self.mspt_list.iter().sum::<f64>() / self.mspt_list.len() as f64;
        let tps = 1000.0 / median_mspt;
        let message = format!("MSPT: {:.4}ms TPS: {:.2}", median_mspt, tps);

        self.broadcast_message(message).await;
        self.mspt_list.clear();
    }

    pub async fn broadcast_message(&self, message: String) {
        let server_config = Arc::clone(&self.server_config);

        let entity_processor = server_config.read().await.entity_processor.clone(); // Clonamos el Arc
        let entity_processor = entity_processor.lock().await; // Bloqueamos el Mutex

        // Obtenemos los clientes y los bloqueamos
        let clients_arc = entity_processor.get_clients().await; // Obtenemos el Arc
        let clients = clients_arc.lock().await; // Bloqueamos el Mutex
                                                // Aquí `clients` tiene un tiempo de vida más largo

        let mut chat_packet = PacketBuilder::new(0x6D);
        let nbt = Nbt::new(
            "root".to_owned(),
            NbtCompound::from_iter([("text".to_owned(), message.to_owned().into())]),
        );
        // Escribimos el NBT como un string en el paquete
        chat_packet.write_nbt(nbt.clone());
        chat_packet.write_nbt(nbt);

        for (_id, client) in clients.iter() {
            // Bloqueamos cada cliente
            let client = client.lock().await;

            // Intentamos hacer downcast a Player
            if let Some(player) = client.as_any().downcast_ref::<Player>() {
                let connection_arc = player.get_connection();
                let mut connection = connection_arc.lock().await; // Bloqueamos la conexión

                connection.send_packet_builder(chat_packet.clone()).await;
            }
        }
    }

    pub async fn broadcast_packet_except(&self, except: Arc<Mutex<Player>>, packet: PacketBuilder) {
        let server_config = Arc::clone(&self.server_config);
        let packet_clone = packet.clone();

        tokio::spawn(async move {
            let entity_processor = server_config.read().await.entity_processor.clone();
            let entity_processor = entity_processor.lock().await;

            let clients_arc = entity_processor.get_clients().await;
            let clients = clients_arc.lock().await;
            let locked_player = except.lock().await;
            let entity_id = locked_player.get_entity_id();
            drop(locked_player);
            for (_id, client) in clients.iter() {
                let client = client.lock().await;
                if let Some(player) = client.as_any().downcast_ref::<Player>() {
                    if player.get_entity_id() == entity_id {
                        continue;
                    }
                    let connection_arc = player.get_connection();
                    let mut connection = connection_arc.lock().await; // Bloqueamos la conexión
                    connection.send_packet_builder(packet_clone.clone()).await;
                    drop(connection);
                }
            }
        });
    }
}
