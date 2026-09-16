//! Хранение устройств в памяти.

use std::sync::Mutex;

use uuid::Uuid;

use crate::models::{Device, DeviceCreate};

/// Хранит устройства в памяти на время работы процесса.
#[derive(Default)]
pub struct DeviceStore {
    devices: Mutex<Vec<Device>>,
}

impl DeviceStore {
    /// Регистрирует новое устройство со сгенерированным id и статусом connected.
    pub fn add(&self, payload: DeviceCreate) -> Device {
        let device = Device {
            id: Uuid::new_v4(),
            type_id: payload.type_id,
            house_id: payload.house_id,
            serial_number: payload.serial_number,
            status: "connected".to_string(),
        };
        self.devices.lock().unwrap().push(device.clone());
        device
    }

    /// Возвращает все зарегистрированные устройства.
    pub fn list(&self) -> Vec<Device> {
        self.devices.lock().unwrap().clone()
    }
}
