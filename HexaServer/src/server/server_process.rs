use bytes::BytesMut;
use std::{sync::Arc, time::Duration};
use tokio::sync::{mpsc::UnboundedReceiver, Mutex};

pub struct ServerProcess {
    pub packet_receiver: Arc<Mutex<UnboundedReceiver<BytesMut>>>, // Cambiado a Arc<Mutex<...>>
    pub packets: Arc<Mutex<Vec<BytesMut>>>,                       // Lista de paquetes para procesar
}

impl ServerProcess {
    pub async fn run(self) {
        // Clonamos el Arc para el acceso a la lista de paquetes
        let packets = Arc::clone(&self.packets);
        let packet_receiver = Arc::clone(&self.packet_receiver);

        // Lanzamos el hilo para recibir paquetes
        self.start_packet_receiver(packet_receiver).await;

        loop {
            let start = std::time::Instant::now();

            // Procesamos los paquetes
            self.process_packets().await;

            // Calculamos el tiempo transcurrido
            let elapsed = start.elapsed();
            println!("Elapsed: {:?}", elapsed);

            // Esperamos el tiempo restante del tick, asegurando que sean 50ms
            let remaining_time = Duration::from_millis(50).saturating_sub(elapsed);
            if remaining_time > Duration::ZERO {
                tokio::time::sleep(remaining_time).await;
            }
        }
    }

    async fn start_packet_receiver(
        &self,
        packet_receiver: Arc<Mutex<UnboundedReceiver<BytesMut>>>,
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

    pub async fn process_packets(&self) {
        let mut packets = self.packets.lock().await;
        for packet in packets.drain(..) {
            println!("Procesando paquete: {:?}", packet);
        }
    }
}
