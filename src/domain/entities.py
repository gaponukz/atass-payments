import dataclasses


@dataclasses.dataclass
class Passenger:
    id: str
    gmail: str
    full_name: str
    phone_number: str
    move_from_id: str
    move_to_id: str
    is_anonymous: bool


@dataclasses.dataclass
class Payment:
    id: str
    amount: int
    route_id: str
    passenger: Passenger


@dataclasses.dataclass
class OutboxData:
    payment_id: str
    route_id: str
    passenger: Passenger
