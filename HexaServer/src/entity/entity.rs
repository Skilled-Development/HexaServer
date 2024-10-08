use std::any::Any;

pub trait Entity: Send + Any {
    fn get_id(&self) -> i32;

    fn as_any(&self) -> &dyn Any;
}
