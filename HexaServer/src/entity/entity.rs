pub trait Entity: Send {
    fn get_id(&self) -> i32;
}
