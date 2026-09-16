"""Хранение записей телеметрии в памяти."""

from datetime import datetime, timezone
from threading import Lock

from telemetry_service.models import TelemetryCreate, TelemetryRecord


class TelemetryStore:
    """Хранит записи телеметрии в памяти на время работы процесса."""

    def __init__(self) -> None:
        self._records: list[TelemetryRecord] = []
        self._lock = Lock()
        self._next_id = 1

    def add(self, payload: TelemetryCreate) -> TelemetryRecord:
        """Сохраняет новую запись телеметрии и возвращает ее."""
        with self._lock:
            record = TelemetryRecord(
                id=self._next_id,
                device_id=payload.device_id,
                metric=payload.metric,
                value=payload.value,
                recorded_at=datetime.now(timezone.utc),
            )
            self._next_id += 1
            self._records.append(record)
            return record

    def list(self, device_id: str | None = None) -> list[TelemetryRecord]:
        """Возвращает сохраненные записи телеметрии, опционально с фильтром по device_id."""
        with self._lock:
            if device_id is None:
                return list(self._records)
            return [record for record in self._records if record.device_id == device_id]
