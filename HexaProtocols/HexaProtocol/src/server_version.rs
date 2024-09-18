
pub trait ServerVersion {
    fn version(&self) -> String;
    fn protocol(&self) -> i32;
    fn get_version_info(&self) -> String {
        format!("Version: {} - Protocol: {}", self.version(), self.protocol())
    }

}
