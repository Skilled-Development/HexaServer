use tokio::sync::Mutex;

use crate::Player;

use super::entity::Entity;
use std::{collections::HashMap, sync::Arc};

pub struct EntityProcessor {
    pub entities: Arc<Mutex<HashMap<i32, Arc<Mutex<dyn Entity>>>>>,
}

impl EntityProcessor {
    pub fn new() -> Self {
        EntityProcessor {
            entities: Arc::new(Mutex::new(HashMap::new())),
        }
    }
    pub async fn next_entity_id(&self, entity: Arc<Mutex<dyn Entity>>) -> i32 {
        let mut entities = self.entities.lock().await;
        let id = entities.len() as i32;
        entities.insert(id, entity);
        id
    }
    pub async fn remove_entity(&self, entity_id: i32) {
        let mut entities = self.entities.lock().await;
        entities.remove(&entity_id);
    }
    pub fn process(&self) {
        println!("Processing entity");
    }
    pub fn get_entities(&self) -> Arc<Mutex<HashMap<i32, Arc<Mutex<dyn Entity>>>>> {
        Arc::clone(&self.entities)
    }
    pub async fn get_clients(&self) -> Arc<Mutex<HashMap<i32, Arc<Mutex<dyn Entity>>>>> {
        let player_entities = self.entities.lock().await;
        let mut clients = HashMap::new();

        for (id, entity) in player_entities.iter() {
            // Intentar hacer downcast a Player
            let entity_guard = entity.lock().await; // Asegúrate de manejar posibles errores de bloqueo.
            if let Some(_player) = entity_guard.as_any().downcast_ref::<Player>() {
                // Si es un Player, lo insertamos en el HashMap
                clients.insert(*id, Arc::clone(entity));
            }
        }

        Arc::new(Mutex::new(clients))
    }
}
