use std::{
    sync::{Arc, RwLock},
    time::Duration
};
use sysinfo::{System, Pid};
use tokio::time::sleep; 

pub struct Monitor{
    pub pid: i32,
    pub system: Arc<RwLock<System>>,
}

impl Monitor{
    pub fn new(pid: i32) -> Self{
        Monitor{
            pid,
            system: Arc::new(RwLock::new(System::new_all())),
        }
    }

    pub fn print_memory_usage(&self) {
        {
            let mut sys = self.system.write().unwrap();
            sys.refresh_specifics(sysinfo::RefreshKind::everything()); 
        }
        
        

        let sys = self.system.read().unwrap();

        if let Some(process) = sys.process(Pid::from(self.pid as usize)) {
            let memory = process.memory(); 
            let virtual_memory = process.virtual_memory(); 
            let cpu_usage = process.cpu_usage(); 
        
            println!("=== Process Monitoring (PID: {}) ===", self.pid);
            println!("Physical memory used: {:.2} MB", memory as f64 / 1024.0 / 1024.0);
            println!("Virtual memory used: {:.2} MB", virtual_memory as f64 / 1024.0 / 1024.0);
            println!("CPU usage: {:.2}%", cpu_usage);
            println!("=======================================");
        } else {
            println!("Could not find the process with PID: {}", self.pid);
        }        
    
    }

    pub async fn start_memory_monitor(&mut self) {
        loop {
            self.print_memory_usage();
            sleep(Duration::from_secs(5)).await; 
        }
    }
}