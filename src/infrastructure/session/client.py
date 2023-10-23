import time
import requests
import dataclasses
import threading
import dataclass_factory
from src.infrastructure.session import dto
from src.application.dto import CreateExternalPaymentDTO
from src.application.dto import ExternalPaymentCreatedDTO
from src.application.dto import PaymentStatus


@dataclasses.dataclass
class SessionCreds:
    url: str
    login: str
    password: str
    merchant_config_id: str
    config_id: str


class HttpPaymentSession:
    def __init__(self, creds: SessionCreds):
        self._url = creds.url
        self._creds = creds
        self._session = requests.Session()
        self._factory = dataclass_factory.Factory()

        self._access_token: str | None = None

        threading.Thread(target=self._authorize_session).start()

    def create_payment(
        self, data: CreateExternalPaymentDTO
    ) -> ExternalPaymentCreatedDTO:
        json_data = self._factory.dump(
            dto.CreatePaymentDTO(
                amount=data.amount,
                title=data.title,
                description=data.description,
                short_description=data.short_description,
                external_id=data.id,
                merchant_config_id=self._creds.merchant_config_id,
                config_id=self._creds.config_id,
                options=dto.CreatePaymenOptions(
                    ttl=data.ttl_seconds,
                    create_short_url=False,
                    backurl=dto.CreatePaymenOptionsBackUrls(
                        success=f"{data.after_payment_url}/success",
                        error=f"{data.after_payment_url}/error",
                        cancel=f"{data.after_payment_url}/cancel",
                    ),
                ),
            )
        )
        response = self._make_request("POST", "/frames/links/pga", json_data)
        created = self._factory.load(response, dto.CreatePaymentResponse)

        return ExternalPaymentCreatedDTO(id=created.id, url=created.url)

    def get_payment_status(self, payment_id: str) -> PaymentStatus:
        response = self._make_request("GET", f"/frames/links/pga/{payment_id}")
        data = self._factory.load(response, dto.PaymentStatusDTO)

        return self._map_payment_status(data.status)

    def _map_payment_status(
        self, payment_status: dto.ApiPaymentStatus
    ) -> PaymentStatus:
        if payment_status in ["ACTIVE", "PENDING"]:
            return PaymentStatus.EXIST

        if payment_status == "USED":
            return PaymentStatus.DONE

        return PaymentStatus.NOT_EXIST

    def _make_request(self, method: str, path: str, data: object = None) -> dict:
        return self._session.request(
            method,
            f"{self._url}{path}",
            json=data,
            headers={
                "Authorization": f"Bearer {self._access_token}",
                "Content-Type": "application/json",
            },
        ).json()

    def _authorize_session(self, time_to_wait: int = 0):
        body = dto.AuthorizeDTO(
            params=dto.AuthorizeInnerDTO(
                login=self._creds.login,
                password=self._creds.password,
                client="transacter",
            )
        )

        time.sleep(time_to_wait)

        response = self._session.post(
            f"{self._url}/auth/token", json=self._factory.dump(body)
        )

        print(f"body: {response.json()}")

        output = self._factory.load(response.json(), dto.GetTokensDTO)
        self._access_token = output.data.access_token
        self._refresh_token = output.data.refresh_token

        print(f"_authorize_session: new token after {output.data.expires_in} seconds")
        self._authorize_session(output.data.expires_in)
