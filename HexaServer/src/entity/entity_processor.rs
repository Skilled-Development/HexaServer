use tokio::sync::Mutex;

use super::entity::Entity;
use std::{collections::HashMap, sync::Arc};

pub struct EntityProcessor {
    pub entities: Arc<Mutex<HashMap<i32, Box<dyn Entity>>>>,
}

impl EntityProcessor {
    pub fn new() -> Self {
        EntityProcessor {
            entities: Arc::new(Mutex::new(HashMap::new())),
        }
    }
    pub fn process(&self) {
        println!("Processing entity");
    }

    // pub fn next_entity_id(&self, entity: Entity) -> i32 {}
}
