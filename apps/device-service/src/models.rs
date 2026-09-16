//! Типы данных, описывающие устройства, которыми управляет сервис.

use serde::{Deserialize, Serialize};
use uuid::Uuid;

/// Устройство, подключенное к экосистеме умного дома.
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Device {
    pub id: Uuid,
    pub type_id: String,
    pub house_id: String,
    pub serial_number: String,
    pub status: String,
}

/// Данные, необходимые для регистрации нового устройства.
#[derive(Debug, Deserialize)]
pub struct DeviceCreate {
    pub type_id: String,
    pub house_id: String,
    pub serial_number: String,
}
