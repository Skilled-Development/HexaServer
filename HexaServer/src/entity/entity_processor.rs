use tokio::sync::Mutex;

use super::entity::{self, Entity};
use std::{collections::HashMap, sync::Arc};

pub struct EntityProcessor {
    pub entities: Arc<Mutex<HashMap<i32, Box<dyn Entity + Send>>>>,
}

impl EntityProcessor {
    pub fn new() -> Self {
        EntityProcessor {
            entities: Arc::new(Mutex::new(HashMap::new())),
        }
    }
    pub async fn next_entity_id(&self, entity: Box<dyn Entity + Send>) -> i32 {
        let mut entities = self.entities.lock().await;
        let id = entities.len() as i32;
        entities.insert(id, entity);
        id
    }
    pub fn process(&self) {
        println!("Processing entity");
    }
}
