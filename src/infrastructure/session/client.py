import time
import requests
import dataclasses
import multiprocessing
import dataclass_factory
from src.infrastructure.session import dto
from src.application.dto import CreateExternalPaymentDTO
from src.application.dto import ExternalPaymentCreatedDTO


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
        self._refresh_token: str | None = None
        self._refresh_daemon_pid: int | None = None

        self._authorize_session()

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
        response = self._factory.load(response, dto.CreatePaymentResponse)

        return ExternalPaymentCreatedDTO(id=response.id, url=response.url)

    def get_payment_status(self, payment_id: str) -> dto.PaymentStatus:
        response = self._make_request("GET", f"/frames/links/pga/{payment_id}")
        data = self._factory.load(response, dto.PaymentStatusDTO)

        return data.status

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

    def _authorize_session(self):
        body = dto.AuthorizeDTO(
            params=dto.AuthorizeInnerDTO(
                login=self._creds.login,
                password=self._creds.password,
                client="transacter",
            )
        )

        response = self._session.post(
            f"{self._url}/auth/token", json=self._factory.dump(body)
        )

        output = self._factory.load(response.json(), dto.GetTokensDTO)
        self._access_token = output.data.access_token
        self._refresh_token = output.data.refresh_token

        # proccess = multiprocessing.Process(
        #     target=self._start_refresh_daemon, args=(output.data.expires_in,)
        # )
        # proccess.start()

        # self._refresh_daemon_pid = proccess.pid

    def _start_refresh_daemon(self, seconds_to_wait: int):
        time.sleep(seconds_to_wait - 2)
        # TODO: make request to refresh
