"""Приложение FastAPI, предоставляющее API сервиса телеметрии."""

import os

from fastapi import FastAPI

from telemetry_service.models import TelemetryCreate, TelemetryRecord
from telemetry_service.store import TelemetryStore

app = FastAPI(title="Telemetry Service")
store = TelemetryStore()


@app.get("/health")
def health() -> dict[str, str]:
    """Сообщает о работоспособности сервиса."""
    return {"status": "ok"}


@app.post("/telemetry", response_model=TelemetryRecord, status_code=201)
def record_telemetry(payload: TelemetryCreate) -> TelemetryRecord:
    """Записывает новое измерение телеметрии для устройства."""
    return store.add(payload)


@app.get("/telemetry", response_model=list[TelemetryRecord])
def list_telemetry(device_id: str | None = None) -> list[TelemetryRecord]:
    """Возвращает записи телеметрии, опционально с фильтром по device_id."""
    return store.list(device_id)


def main() -> None:
    """Запускает сервис с помощью uvicorn."""
    import uvicorn

    port = int(os.environ.get("PORT", "8083"))
    uvicorn.run(app, host="0.0.0.0", port=port)


if __name__ == "__main__":
    main()
