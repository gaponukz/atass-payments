import typing
import dataclasses


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
class CreatePaymentDTO:
    amount: int
    title: str
    description: str
    short_description: str
    external_id: str
    merchant_config_id: str
    config_id: str
    options: dict | None = None
    params: dict | None = None
    hold: bool = False
    lang: typing.Literal["UK", "RU", "EN"] = "UK"


@dataclasses.dataclass
class CreatePaymentResponse:
    id: str
    url: str
    short_url: str
    signature: str
