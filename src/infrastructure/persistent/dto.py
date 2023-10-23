import dataclasses
from src.domain import entities
from src.application import dto


@dataclasses.dataclass
class PaymentDTO:
    status: dto.PaymentStatus
    payment: entities.Payment
