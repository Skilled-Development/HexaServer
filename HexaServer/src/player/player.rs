use crate::entity::entity::Entity;

pub struct Player {
    pub name: String,
    pub uuid: uuid::Uuid,
    pub entity_id: i32,
    pub health: i32,
    pub hunger: i32,
    pub experience: i32,
    pub level: i32,
    pub gamemode: i32,
    pub x: f64,
    pub y: f64,
    pub z: f64,
    pub yaw: f32,
    pub pitch: f32,
}

impl Entity for Player {
    fn get_id(&self) -> i32 {
        0
    }
}

impl Player {
    pub fn new(
        name: String,
        uuid: uuid::Uuid,
        entity_id: i32,
        health: i32,
        hunger: i32,
        experience: i32,
        level: i32,
        gamemode: i32,
        x: f64,
        y: f64,
        z: f64,
        yaw: f32,
        pitch: f32,
    ) -> Player {
        Player {
            name,
            uuid,
            entity_id,
            health,
            hunger,
            experience,
            level,
            gamemode,
            x,
            y,
            z,
            yaw,
            pitch,
        }
    }
}
