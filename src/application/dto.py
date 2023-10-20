import dataclasses
from src.domain import entities


@dataclasses.dataclass
class CreatePaymentDTO:
    amount: float
    route_id: str
    passenger: entities.Passenger


@dataclasses.dataclass
class SavePaymentDTO:
    external_id: str
    payment: entities.Payment


@dataclasses.dataclass
class TextPresentaionOnPaymentCreationDTO:
    title: str
    description: str
    short_description: str
    after_payment_url: str


@dataclasses.dataclass
class CreateExternalPaymentDTO:
    id: str
    ttl_seconds: int
    after_payment_url: str
    amount: int
    title: str
    description: str
    short_description: str


@dataclasses.dataclass
class ExternalPaymentCreatedDTO:
    id: str
    url: str
