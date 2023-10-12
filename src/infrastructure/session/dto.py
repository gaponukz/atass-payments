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
    external_id: str
    merchant_config_id: str
