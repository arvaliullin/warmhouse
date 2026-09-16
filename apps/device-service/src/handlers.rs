//! HTTP-обработчики сервиса устройств.

use std::sync::Arc;

use axum::extract::State;
use axum::http::StatusCode;
use axum::Json;

use crate::models::{Device, DeviceCreate};
use crate::store::DeviceStore;

/// Сообщает о работоспособности сервиса.
pub async fn health() -> Json<serde_json::Value> {
    Json(serde_json::json!({ "status": "ok" }))
}

/// Регистрирует новое устройство.
pub async fn create_device(
    State(store): State<Arc<DeviceStore>>,
    Json(payload): Json<DeviceCreate>,
) -> (StatusCode, Json<Device>) {
    let device = store.add(payload);
    (StatusCode::CREATED, Json(device))
}

/// Возвращает список всех зарегистрированных устройств.
pub async fn list_devices(State(store): State<Arc<DeviceStore>>) -> Json<Vec<Device>> {
    Json(store.list())
}
