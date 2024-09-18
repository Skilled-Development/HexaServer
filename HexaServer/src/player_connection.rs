
#[derive(Eq, Hash, PartialEq, Debug,Clone, Copy)]  // Add Debug trait to ClientState
pub enum ClientState {
    HANDSHAKE,
    LOGIN,
    CONFIGURATION,
    PLAY,
}

#[derive(Clone)]
pub struct PlayerConnection {


    
    pub id: Option<String>,
    pub name: Option<String>,
    pub ip_address: String,
    pub client_state: ClientState,
    pub server_name: String,
    pub server_versions: Vec<i32>,
    pub username: Option<String>,
    pub uuid: Option<uuid::Uuid>,

}

impl PlayerConnection {
    pub fn new(ip: String,server_name:String,server_versions:Vec<i32>) -> PlayerConnection {
        println!("Creating new connection with IP {}", ip);
        PlayerConnection {
            id: None,
            name: None,
            ip_address: ip,
            client_state: ClientState::HANDSHAKE,
            server_name,
            server_versions,
            username:None,
            uuid:None,
        }
    }

    pub fn set_client_state(&mut self, client_state: ClientState) {
        self.client_state = client_state;
    }

    pub fn set_username(&mut self, username: String) {
        self.username = Some(username);
    }

    pub fn get_username(&self) -> String {
        self.username.clone().unwrap()
    }


    pub fn set_uuid(&mut self, uuid: uuid::Uuid) {
        self.uuid = Some(uuid);
    }

    pub fn get_uuid(&self) -> uuid::Uuid {
        self.uuid.clone().unwrap()
    }
    

    pub fn get_client_state(&self) -> ClientState {
        self.client_state
    }


    
}
