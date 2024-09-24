pub mod hexa_protocol_1_21;
pub use hexa_protocol_1_21::HexaProtocol1_21; 
pub mod packets {
    pub mod client{
        pub mod handshake{
            pub mod handshake_packet;
            pub mod ping_request_packet;
        }
        pub mod play{
            pub mod confirm_teleport_packet_1_21;
            pub mod set_player_position_packet_1_21;
            pub mod set_player_position_and_rotation_packet_1_21;
            pub mod ping_request_play_packet_1_21;
            pub mod set_item_held_packet_1_21;
            pub mod swing_arm_packet_1_21;
        }
    }
    pub mod server{
        pub mod handshake{
            pub mod status_response_packet_1_21;
            pub mod ping_response_packet;
        }
        pub mod configuration{
            pub mod data_registry{
                pub mod data_registry_packet_1_21;
            }
        }
        pub mod play{
            pub mod update_section_blocks_packet_1_21;
        }
    }
}

pub use packets::server::configuration::data_registry::data_registry_packet_1_21::read_data_file_to_bytesmut;