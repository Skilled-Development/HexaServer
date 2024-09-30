use bytes::{Buf, BytesMut};
use uuid::Uuid; // Asegúrate de tener la dependencia `uuid` en tu proyecto

pub struct PacketReader<'a> {
    pub buf: &'a mut BytesMut,
}

impl<'a> PacketReader<'a> {
    pub fn new(buffer: &'a mut BytesMut) -> Self {
        PacketReader { buf: buffer }
    }

    // Leer un Byte (8 bits)
    pub fn read_byte(&mut self) -> i8 {
        self.ensure_remaining(1);
        self.buf.get_i8()
    }

    // Leer un Unsigned Byte (8 bits)
    pub fn read_unsigned_byte(&mut self) -> u8 {
        self.ensure_remaining(1);
        self.buf.get_u8()
    }
    pub fn read_short(&mut self) -> i16 {
        self.ensure_remaining(2);
        self.buf.get_i16()
    }
    pub fn read_float(&mut self) -> f32 {
        self.ensure_remaining(4);
        self.buf.get_f32() 
    }

    pub fn read_boolean(&mut self) -> bool {
        self.ensure_remaining(1); // Asegura que haya al menos 1 byte en el buffer
        let byte = self.read_unsigned_byte(); // Leer un byte sin signo
        match byte {
            0x00 => false, // 0x00 representa `false`
            0x01 => true,  // 0x01 representa `true`
            _ => false // Cualquier otro valor no es válido
        }
    }
    pub fn read_position(&mut self) -> (i32, i32, i32) {
        self.ensure_remaining(8); // Asegura que haya al menos 8 bytes en el buffer
        let val = self.buf.get_u64();

        let x = ((val >> 38) & 0x3FFFFFF) as i32;
        let y = ((val >> 26) & 0xFFF) as i32;
        let z = (val & 0x3FFFFFF) as i32;

        // Ajustar los valores para el rango correcto
        let x = if x >= 0x2000000 { x - 0x4000000 } else { x };
        let y = if y >= 0x800 { y - 0x1000 } else { y };
        let z = if z >= 0x2000000 { z - 0x4000000 } else { z };

        (x, y, z)
    }
    pub fn read_varlong(&mut self) -> i64 {
        let mut value = 0;
        let mut position = 0;

        loop {
            self.ensure_remaining(1); // Asegúrate de que haya al menos un byte para leer
            let current_byte = self.read_unsigned_byte();
            value |= ((current_byte & 0x7F) as i64) << position;

            if (current_byte & 0x80) == 0 {
                break;
            }

            position += 7;
            if position >= 64 {
                panic!("VarLong is too big");
            }
        }

        value
    }
    // Leer un Unsigned Short (16 bits sin signo)
    pub fn read_unsigned_short(&mut self) -> u16 {
        self.ensure_remaining(2);
        self.buf.get_u16() // Lee un entero sin signo de 16 bits (big-endian)
    }

    // Leer un VarInt (tamaño variable)
    pub fn read_varint(&mut self) -> i32 {
        let mut value = 0;
        let mut position = 0;

        loop {
            self.ensure_remaining(1); // Asegúrate de que haya al menos un byte para leer
            let current_byte = self.read_unsigned_byte();
            value |= ((current_byte & 0x7F) as i32) << position;

            if (current_byte & 0x80) == 0 {
                break;
            }

            position += 7;
            if position >= 32 {
                panic!("VarInt is too big");
            }
        }

        value
    }

    // Leer una cadena (String)
    pub fn read_string(&mut self) -> String {
        // Primero leer la longitud de la cadena como VarInt
        let length = self.read_varint();
        // Verificar que el buffer tenga suficiente longitud para la cadena
        self.ensure_remaining(length as usize);
        // Luego leer el contenido de la cadena (codificada en UTF-8)
        let mut buf = vec![0; length as usize];
        self.buf.copy_to_slice(&mut buf);
        String::from_utf8(buf).expect("Invalid UTF-8 string")
    }
    pub fn read_angle(&mut self) -> f32 {
        // Asegúrate de que haya al menos 1 byte para leer
        self.ensure_remaining(1);
        // Lee el byte y conviértelo de nuevo a un ángulo de 0 a 360 grados
        let encoded_angle = self.read_unsigned_byte();
        (encoded_angle as f32 / 256.0) * 360.0
    }
    // Leer un Long (64 bits)
    pub fn read_long_be(&mut self) -> i64 {
        self.ensure_remaining(8); // Asegura que haya al menos 8 bytes en el buffer
        let value = self.buf.get_i64(); // Lee el valor en formato big-endian
        value.to_be() // Convierte el valor a big-endian
    }
    pub fn read_int(&mut self) -> i32 {
        self.ensure_remaining(4); // Asegura que haya al menos 4 bytes para leer un entero
        self.buf.get_i32() // Lee un entero de 32 bits en formato big-endian
    }

    // Leer un Double (64 bits)
    pub fn read_double(&mut self) -> f64 {
        self.ensure_remaining(8); // Asegura que haya al menos 8 bytes para leer un double
        self.buf.get_f64() // Lee un double de 64 bits en formato big-endian
    }
    // Leer un UUID (128 bits)
    pub fn read_uuid(&mut self) -> Uuid {
        self.ensure_remaining(16); // Un UUID tiene 16 bytes (128 bits)

        // Leer los dos enteros de 64 bits
        let most_sig_bits = self.buf.get_u64();
        let least_sig_bits = self.buf.get_u64();

        // Crear un UUID a partir de los bits más y menos significativos
        Uuid::from_u64_pair(most_sig_bits, least_sig_bits)
    }

    // Verificar si hay suficientes bytes para leer un determinado tamaño
    pub fn has_remaining(&self, size: usize) -> bool {
        self.buf.len() >= size
    }

    // Asegura que el buffer tenga al menos `size` bytes restantes
    fn ensure_remaining(&self, size: usize) {
        if self.buf.len() < size {
            panic!("Not enough bytes remaining in the buffer");
        }
    }

    pub fn read_identifier(&mut self) -> String {
        // Leer la longitud del identificador como VarInt
        let length = self.read_varint();
        
        // Verificar que el tamaño del identificador sea válido
        if length < 1 || length > (32767 * 3) + 3 {
            panic!("Identifier length is invalid");
        }
    
        // Asegurarse de que el buffer tenga suficientes bytes para leer el identificador
       // self.ensure_remaining(length as usize);
    
        // Leer el contenido del identificador (codificado en UTF-8)
        let mut buf = vec![0; length as usize];
        self.buf.copy_to_slice(&mut buf);
        
        // Convertir los bytes a una cadena y retornar
        String::from_utf8(buf).expect("Invalid UTF-8 identifier")
    }
    pub fn read_bytearray(&mut self, max_length: i32) -> Vec<u8> {
        // Leer la longitud del arreglo de bytes como VarInt
        let length = self.read_varint();
    
        // Si se pasa -1 como max_length, no hay límite en la longitud
        if max_length != -1 {
            // Verificar que la longitud esté dentro de los límites aceptables
            if length < 0 || length > max_length {
                panic!("Byte array length exceeds the allowed maximum");
            }
        }
    
        // Asegurar que hay suficientes bytes para leer el arreglo
        //self.ensure_remaining(length as usize);
    
        // Leer el contenido del arreglo de bytes
        let mut buf = vec![0; length as usize];
        self.buf.copy_to_slice(&mut buf);
    
        buf
    }
    
}
