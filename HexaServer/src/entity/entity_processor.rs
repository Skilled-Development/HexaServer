use tokio::sync::Mutex;

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
}
