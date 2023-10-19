import typing
import dataclasses

PaymentStatus = typing.Literal[
    "ACTIVE", "EXPIRED", "USED", "DELETED", "FAILED", "PENDING"
]


@dataclasses.dataclass
class CreatePaymentDTO:
    id: str
    ttl_seconds: int
    after_payment_url: str
    amount: int
    title: str
    description: str
    short_description: str


@dataclasses.dataclass
class PaymentCreatedDTO:
    id: str
    url: str
