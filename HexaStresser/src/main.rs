use tokio::{io::AsyncWriteExt, net::TcpStream, time::sleep};
use bytes::{BytesMut, BufMut};
use std::time::{Duration, Instant};

async fn send_handshake_packet(mut socket: TcpStream) {
    let mut packet = BytesMut::new();
    
    // ID del paquete (Handshake)
    write_int(&mut packet, 0x00); // Packet ID
    write_int(&mut packet, 0); // Length (placeholder, can be updated based on actual data)
    write_int(&mut packet, 1); // Protocol version
    write_string(&mut packet, "localhost"); // Server address
    write_unsignedshort(&mut packet, 25565); // Server port
    write_int(&mut packet, 1); // Next state

    // Añadir longitud total del paquete
    let length = packet.len() as i32;
    let mut final_packet = BytesMut::new();
    write_int(&mut final_packet, length);
    final_packet.extend_from_slice(&packet);

    // Enviar paquete
    if let Err(err) = socket.write_all(&final_packet).await {
        println!("Error al enviar el paquete: {}", err);
    }
}

async fn simulate_connections(address: &str, port: u16, num_connections: usize) {
    let mut tasks = Vec::with_capacity(num_connections);
    
    for _ in 0..num_connections {
        let address = address.to_string();
        let port = port;
        tasks.push(tokio::spawn(async move {
            match TcpStream::connect(format!("{}:{}", address, port)).await {
                Ok(socket) => {
                    println!("Conectado a {}:{}", address, port);
                    send_handshake_packet(socket).await;
                },
                Err(err) => {
                    println!("Error al conectar: {}", err);
                }
            }
        }));
    }
    
    // Esperar a que todas las tareas se completen
    for task in tasks {
        if let Err(err) = task.await {
            println!("Error en la tarea: {:?}", err);
        }
    }
}

fn write_int(buf: &mut BytesMut, value: i32) {
    buf.put_i32_le(value);
}

fn write_long(buf: &mut BytesMut, value: i64) {
    buf.put_i64_le(value);
}

fn write_string(buf: &mut BytesMut, value: &str) {
    let length = value.len() as i16;
    buf.put_i16_le(length);
    buf.put_slice(value.as_bytes());
}

fn write_unsignedshort(buf: &mut BytesMut, value: u16) {
    buf.put_u16_le(value);
}

#[tokio::main]
async fn main() {
    let address = "127.0.0.1";
    let port = 25565;
    let num_connections = 1000;

    let start_time = Instant::now();
    let duration = Duration::new(10, 0); // 10 segundos

    while start_time.elapsed() < duration {
        simulate_connections(address, port, num_connections).await;

        // Espera un breve período antes de repetir la simulación para no saturar el servidor
        sleep(Duration::from_millis(100)).await;
    }
}