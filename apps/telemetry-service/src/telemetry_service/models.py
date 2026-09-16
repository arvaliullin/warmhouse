"""Схемы Pydantic для записей телеметрии."""

from datetime import datetime

from pydantic import BaseModel


class TelemetryCreate(BaseModel):
    """Данные, необходимые для записи нового измерения телеметрии."""

    device_id: str
    metric: str
    value: str


class TelemetryRecord(BaseModel):
    """Измерение телеметрии, сохраненное сервисом."""

    id: int
    device_id: str
    metric: str
    value: str
    recorded_at: datetime
