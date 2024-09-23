
pub trait Packet {
    fn get_packet_id(&self) -> i32;
    fn get_packet_type(&self) -> super::packet_type::PacketType;
}
