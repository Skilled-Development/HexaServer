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
    pub broadcast_packets_flag: bool,
    pub broadcast_packet_map: HashMap<Arc<Mutex<Player>>, Vec<PacketBuilder>>,
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
            broadcast_packets_flag: false,
            broadcast_packet_map: HashMap::new(),
        }
    }

    pub async fn run(mut self) {
        // Clonamos el Arc para el acceso a la lista de paquetes
        let _packets = Arc::clone(&self.packets);
        let packet_receiver = Arc::clone(&self.packet_receiver);

        self.start_packet_receiver(packet_receiver).await;
        loop {
            let start = std::time::Instant::now();

            // Procesamos los paquetes
            let _ = self.process_packets().await;
            self.broadcast_packets_flag = true;
            // Calculamos el tiempo transcurrido
            let elapsed = start.elapsed();
            //println!("Elapsed: {:?}", elapsed);
            self.broadcast_message(format!("Time elapsed: {}", elapsed.as_micros().to_string()))
                .await;

            // Esperamos el tiempo restante del tick, asegurando que sean 50ms
            let remaining_time = Duration::from_millis(50).saturating_sub(elapsed);
            if remaining_time > Duration::ZERO {
                tokio::time::sleep(remaining_time).await;
            }
        }
    }

    async fn start_packet_receiver(
        &self,
        packet_receiver: Arc<Mutex<UnboundedReceiver<PacketBuffer>>>,
    ) {
        // Clonamos el Arc para moverlo al futuro
        let packets = Arc::clone(&self.packets);

        tokio::spawn(async move {
            // Debemos bloquear el Mutex para obtener el receiver
            let mut receiver = packet_receiver.lock().await;

            while let Some(packet) = receiver.recv().await {
                let mut packets = packets.lock().await;
                packets.push(packet);
            }
        });
    }

    pub async fn process_packets(&self) -> Result<(), String> {
        let mut packets = self.packets.lock().await;

        if packets.is_empty() {
            return Ok(());
        }

        // Liberar el lock antes del procesamiento de paquetes
        let packets_to_process = std::mem::take(&mut *packets);

        for mut packet in packets_to_process {
            let packet_id = packet.get_packet_id();
            let length = packet.get_packet_length();
            let client = Arc::clone(&packet.get_client());
            let buffer = packet.get_mut_buffer();

            // Desbloquear el buffer antes de procesar
            let result_on_read = match packet_id {
                0x00 => confirm_teletransportation::handle(length, buffer, client.clone()).await,
                0x21 => ping_request_play::handle(length, buffer, client.clone()).await,
                0x1A => set_player_position::handle(length, buffer, client.clone(), self).await,
                0x1B => {
                    set_player_position_and_rotation::handle(length, buffer, client.clone(), self)
                        .await
                }
                0x36 => swing_arm::handle(length, buffer, client.clone(), self).await,
                0x18 => keep_alive::handle(length, buffer, client.clone()).await,
                0x1c => set_player_rotation::handle(length, buffer, client.clone(), self).await,
                /*0x2F => set_item_held::handle(length, buffer, reader, client.clone()).await,


                0x20 => pick_item::handle(length, buffer, reader, client.clone()).await,*/
                _ => {
                    println!("Unknown packet ID: 0x{:x} in play handler", packet_id);
                    buffer.clear();
                    continue;
                }
            };

            // Solo bloquear cuando sea necesario
            if let Ok(_) = result_on_read {
                let client_guard = client.lock().await;
                let connection = client_guard.get_connection();
                let mut connection_guard = connection.lock().await;

                let client_last_keep_alive = connection_guard.get_last_keep_alive();
                if client_last_keep_alive.elapsed().as_millis() > 17000 {
                    println!("Client {} timed out", connection_guard.ip_address);

                    let mut keep_alive_packet = PacketBuilder::new(0x26);
                    let random_id: i64 = rand::thread_rng().gen();
                    keep_alive_packet.write_long_be(random_id);

                    connection_guard
                        .send_packet_builder(keep_alive_packet)
                        .await;
                    connection_guard.set_keep_alive_id(random_id);
                    connection_guard.set_last_keep_alive(Instant::now().into());

                    println!(
                        "Sent keep alive packet to client {} with alive id {}",
                        connection_guard.ip_address, random_id
                    );
                }
            } else {
                return result_on_read; // Devolver el error si alguna lectura falla
            }
        }

        Ok(())
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
