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

    pub fn read_boolean(&mut self) -> bool {
        self.ensure_remaining(1); // Asegura que haya al menos 1 byte en el buffer
        let byte = self.read_unsigned_byte(); // Leer un byte sin signo
        match byte {
            0x00 => false, // 0x00 representa `false`
            0x01 => true,  // 0x01 representa `true`
            _ => panic!("Invalid boolean value in packet: {}", byte), // Cualquier otro valor no es válido
        }
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

    // Leer un Long (64 bits)
    pub fn read_long_be(&mut self) -> i64 {
        self.ensure_remaining(8); // Asegura que haya al menos 8 bytes en el buffer
        let value = self.buf.get_i64(); // Lee el valor en formato big-endian
        value.to_be() // Convierte el valor a big-endian
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
