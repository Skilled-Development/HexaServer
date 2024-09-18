use bytes::{Buf, BufMut, BytesMut};
use uuid::Uuid;

pub fn write_string(buffer: &mut BytesMut, value: &str) {
    let value_bytes = value.as_bytes();
    let length = value_bytes.len() as i32;
    write_varint(buffer, length); // Escribimos la longitud de la cadena como VarInt
    buffer.extend_from_slice(value_bytes); // Escribimos la cadena de bytes
}
pub fn write_angle(buffer: &mut BytesMut, angle: f32) {
    // Convierte el ángulo en un valor entre 0 y 255 (256 pasos por vuelta completa)
    let encoded_angle = ((angle / 360.0) * 256.0).round() as u8;
    
    // Escribe el valor en el buffer
    buffer.put_u8(encoded_angle);
}
pub fn write_identifier(buffer: &mut BytesMut, identifier: String) {
    // Asegurarse de que el identificador no sea demasiado largo
    if identifier.len() > 32767 {
        println!("Identificador demasiado largo, máximo permitido es 32767 caracteres");
    }
    
    // Convertir el string a bytes UTF-8
    let bytes = identifier.as_bytes();

    // Escribir la longitud de la cadena como VarInt
    write_varint(buffer, bytes.len() as i32);
    
    // Escribir los bytes del string en el buffer
    buffer.extend_from_slice(bytes);
}
pub fn write_short(buffer: &mut BytesMut, value: i16) {
    // Escribe el valor como un entero de 2 bytes en formato Big Endian
    buffer.put_i16(value);
}
pub fn write_double(buffer: &mut BytesMut, value: f64) {
    // Escribe el valor como un double (64 bits) en formato Big Endian
    buffer.put_f64(value);
}
pub fn write_uuid(buf: &mut BytesMut, uuid: Uuid) {
    // El UUID está compuesto por dos partes de 64 bits cada una
    let binding = uuid.as_u128().to_be().to_be_bytes();
    let (most_sig_bits, least_sig_bits) = binding.split_at(8);

    // Escribir los 64 bits más significativos (big-endian)
    buf.extend_from_slice(most_sig_bits);

    // Escribir los 64 bits menos significativos (big-endian)
    buf.extend_from_slice(least_sig_bits);
}
pub fn write_byte_array(buffer: &mut BytesMut, value: &[u8]) {
    let length = value.len() as i32;
    write_varint(buffer, length); // Escribimos la longitud del array de bytes como VarInt
    buffer.extend_from_slice(value); // Escribimos el array de bytes
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



pub fn read_int(buf: &mut &[u8]) -> Option<i32> {
    if buf.len() < 4 {
        return None; // No hay suficientes bytes para un int
    }
    
    let int_bytes = &buf[..4]; // Tomar los primeros 4 bytes
    *buf = &buf[4..]; // Avanzar en el buffer

    Some(i32::from_be_bytes([int_bytes[0], int_bytes[1], int_bytes[2], int_bytes[3]]))
}




pub fn read_string(buffer: &mut BytesMut) -> Result<String, String> {
    let length = read_varint(buffer)?;
    if buffer.remaining() < length as usize {
        return Err("Longitud de cadena inválida".to_string());
    }
    let string_bytes = buffer.split_to(length as usize);
    String::from_utf8(string_bytes.to_vec()).map_err(|e| format!("Error al convertir la cadena UTF-8: {:?}", e))
}
/* 

pub fn read_string(buf: &mut &[u8]) -> Option<String> {
    let length = read_varint(buf)? as usize;
    println!("Reading string with length: {}", length);
    
    if buf.len() < length {
        println!("Buffer size {} is smaller than the expected string length {}", buf.len(), length);
        return None;
    }

    let string_bytes = &buf[..length];
    *buf = &buf[length..];
    
    println!("Successfully read string: {:?}", String::from_utf8_lossy(string_bytes));
    Some(String::from_utf8_lossy(string_bytes).to_string())
}
*/

// Lee un Long (8 bytes) desde el buffer y elimina el dato del buffer
pub fn read_long(buf: &mut &[u8]) -> Option<i64> {
    if buf.len() < 8 {
        return None;
    }
    let long_bytes = &buf[..8];
    *buf = &buf[8..];
    Some(i64::from_be_bytes(long_bytes.try_into().unwrap()))
}

// Lee un unsigned short (2 bytes) desde el buffer y elimina el dato del buffer
pub fn read_unsigned_short(buf: &mut &[u8]) -> Option<u16> {
    if buf.len() < 2 {
        return None; // No hay suficiente data para leer un u16
    }
    // Lee los primeros 2 bytes y actualiza el buffer
    let (value_bytes, rest) = buf.split_at(2);
    *buf = rest; // Actualiza el buffer para eliminar los bytes leídos

    // Convierte los bytes a un u16
    let value = u16::from_be_bytes([value_bytes[0], value_bytes[1]]);
    Some(value)
}

pub fn write_int(buf: &mut BytesMut, mut value: i32) {
    loop {
        if (value & !0x7F) == 0 {
            buf.put_u8(value as u8);
            break;
        } else {
            buf.put_u8((value & 0x7F | 0x80) as u8);
            value >>= 7;
        }
    }
}

pub fn write_varint(buf: &mut BytesMut, mut value: i32) {
    while (value & 0xFFFFFF80u32 as i32) != 0 {
        buf.put_u8((value as u8 & 0x7F) | 0x80); // Escribe los primeros 7 bits con el bit 8 activado
        value >>= 7;                             // Desplaza 7 bits a la derecha
    }
    buf.put_u8(value as u8);                     // Escribe el último byte sin el bit 8 activado
}


pub fn write_long_be(buf: &mut BytesMut, value: i64) {
    buf.put_i64(value.to_be()); // Convierte a big-endian y escribe
}
//

// Escribe un byte (Byte) en el buffer
pub fn write_byte(buf: &mut BytesMut, value: u8) {
    buf.put_u8(value);
}

// Escribe un booleano (Boolean) en el buffer
pub fn write_boolean(buf: &mut BytesMut, value: bool) {
    buf.put_u8(if value { 1 } else { 0 });
}
