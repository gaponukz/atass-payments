import enum
import typing
import dataclasses

PaymentStatus = typing.Literal[
    "ACTIVE", "EXPIRED", "USED", "DELETED", "FAILED", "PENDING"
]


@dataclasses.dataclass
class GetTokensInnerDTO:
    access_token: str
    expires_in: int
    refresh_expires_in: int
    refresh_token: str


@dataclasses.dataclass
class GetTokensDTO:
    data: GetTokensInnerDTO


@dataclasses.dataclass
class AuthorizeInnerDTO:
    login: str
    password: str
    client: str


@dataclasses.dataclass
class AuthorizeDTO:
    params: AuthorizeInnerDTO


@dataclasses.dataclass
class CreatePaymenOptionsBackUrls:
    success: str
    error: str
    cancel: str


@dataclasses.dataclass
class CreatePaymenOptions:
    ttl: int
    create_short_url: bool
    backurl: CreatePaymenOptionsBackUrls


@dataclasses.dataclass
class CreatePaymentDTO:
    amount: int
    title: str
    description: str
    short_description: str
    external_id: str
    merchant_config_id: str
    config_id: str
    options: CreatePaymenOptions
    hold: bool = False
    lang: typing.Literal["UK", "RU", "EN"] = "UK"


@dataclasses.dataclass
class CreatePaymentResponse:
    id: str
    url: str
    signature: str
    short_url: str | None = None


@dataclasses.dataclass
class PaymentStatusDTO:
    id: str
    status: str
    external_id: str
