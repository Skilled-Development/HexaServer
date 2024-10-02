use std::sync::Arc;

use tokio::sync::Mutex;

use crate::{entity::entity::Entity, PlayerConnection};

use super::game_mode::GameMode;

pub struct Player {
    pub connection: Arc<Mutex<PlayerConnection>>,
    pub name: String,
    pub uuid: uuid::Uuid,
    pub entity_id: i32,
    pub health: i32,
    pub hunger: i32,
    pub experience: i32,
    pub level: i32,
    pub gamemode: GameMode,
    pub x: f64,
    pub y: f64,
    pub z: f64,
    pub yaw: f32,
    pub pitch: f32,
    pub on_ground: bool,
    pub protocol_version: i32,
    pub velocity: (i16, i16, i16),
}

impl Entity for Player {
    fn get_id(&self) -> i32 {
        0
    }
}

impl Player {
    pub fn new(connection: Arc<Mutex<PlayerConnection>>) -> Player {
        Player {
            connection,
            name: "EmptyName".to_string(),
            uuid: uuid::Uuid::new_v4(),
            entity_id: -1,
            health: 20,
            hunger: 20,
            experience: 0,
            level: 0,
            gamemode: GameMode::Survival,
            x: 0.0,
            y: 0.0,
            z: 0.0,
            yaw: 0.0,
            pitch: 0.0,
            on_ground: false,
            protocol_version: 0,
            velocity: (0, 0, 0),
        }
    }

    pub fn get_connection(&self) -> Arc<Mutex<PlayerConnection>> {
        self.connection.clone()
    }

    pub fn set_name(&mut self, name: String) {
        self.name = name;
    }

    pub fn set_uuid(&mut self, uuid: uuid::Uuid) {
        self.uuid = uuid;
    }

    pub fn set_entity_id(&mut self, entity_id: i32) {
        self.entity_id = entity_id;
    }

    pub fn get_name(&self) -> String {
        self.name.clone()
    }

    pub fn get_uuid(&self) -> uuid::Uuid {
        self.uuid
    }

    pub fn get_entity_id(&self) -> i32 {
        self.entity_id
    }

    pub fn get_health(&self) -> i32 {
        self.health
    }

    pub fn set_health(&mut self, health: i32) {
        self.health = health;
    }

    pub fn get_hunger(&self) -> i32 {
        self.hunger
    }

    pub fn set_hunger(&mut self, hunger: i32) {
        self.hunger = hunger;
    }

    pub fn get_experience(&self) -> i32 {
        self.experience
    }

    pub fn set_experience(&mut self, experience: i32) {
        self.experience = experience;
    }

    pub fn get_level(&self) -> i32 {
        self.level
    }

    pub fn set_level(&mut self, level: i32) {
        self.level = level;
    }

    pub fn get_gamemode(&self) -> &GameMode {
        &self.gamemode
    }

    pub fn set_gamemode(&mut self, gamemode: GameMode) {
        self.gamemode = gamemode;
    }

    pub fn get_position(&self) -> (f64, f64, f64) {
        (self.x, self.y, self.z)
    }

    pub fn set_position(&mut self, x: f64, y: f64, z: f64) {
        self.x = x;
        self.y = y;
        self.z = z;
    }

    pub fn get_rotation(&self) -> (f32, f32) {
        (self.yaw, self.pitch)
    }

    pub fn set_rotation(&mut self, yaw: f32, pitch: f32) {
        self.yaw = yaw;
        self.pitch = pitch;
    }
    pub fn is_on_ground(&self) -> bool {
        self.on_ground
    }

    pub fn set_on_ground(&mut self, on_ground: bool) {
        self.on_ground = on_ground;
    }

    pub fn get_protocol_version(&self) -> i32 {
        self.protocol_version
    }

    pub fn set_protocol_version(&mut self, protocol_version: i32) {
        self.protocol_version = protocol_version;
    }

    pub fn get_velocity(&self) -> (i16, i16, i16) {
        self.velocity
    }

    pub fn set_velocity(&mut self, velocity: (i16, i16, i16)) {
        self.velocity = velocity;
    }

    pub async fn get_connection_id(&self) -> String {
        let connection = self.connection.lock().await;
        connection.get_connection_id()
    }
}
