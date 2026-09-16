//! Точка входа сервиса устройств.

mod handlers;
mod models;
mod store;

use std::env;
use std::sync::Arc;

use axum::routing::get;
use axum::Router;

use store::DeviceStore;

#[tokio::main]
async fn main() {
    let store = Arc::new(DeviceStore::default());

    let app = Router::new()
        .route("/health", get(handlers::health))
        .route("/devices", get(handlers::list_devices).post(handlers::create_device))
        .with_state(store);

    let address = env::var("ADDRESS").unwrap_or_else(|_| "0.0.0.0:8082".to_string());
    let listener = tokio::net::TcpListener::bind(&address).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
