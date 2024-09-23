

pub mod packets {
    pub mod client{
        pub mod handshake{
            pub mod handshake_packet;
            pub mod ping_request_packet;
        }
    }
    pub mod server{
        pub mod handshake{
            pub mod status_response_packet;
            pub mod ping_response_packet;
        }
    }
}
pub use packets::client::handshake::handshake_packet::HandshakePacket;
pub use packets::server::handshake::status_response_packet::StatusResponsePacket;
