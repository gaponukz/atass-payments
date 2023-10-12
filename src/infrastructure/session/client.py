import time
import uuid
import requests
import multiprocessing
import dataclass_factory
from src.infrastructure.session import dto


class HttpPaymentSession:
    def __init__(self, url: str, login: str, password: str):
        self._url = url
        self._login = login
        self._password = password
        self._session = requests.Session()
        self._factory = dataclass_factory.Factory()

        self._access_token: str | None = None
        self._refresh_token: str | None = None
        self._refresh_daemon_pid: int | None = None

        self._authorize_session()

    def _authorize_session(self):
        body = dto.AuthorizeDTO(
            params=dto.AuthorizeInnerDTO(
                login=self._login, password=self._password, client="transacter"
            )
        )

        response = self._session.post(
            f"{self._url}/auth/token", json=self._factory.dump(body)
        )

        output = self._factory.load(response.json(), dto.GetTokensDTO)
        self._access_token = output.data.access_token
        self._refresh_token = output.data.refresh_token

        proccess = multiprocessing.Process(
            target=self._start_refresh_daemon, args=(output.data.expires_in,)
        )
        proccess.start()

        self._refresh_daemon_pid = proccess.pid

    def _start_refresh_daemon(self, seconds_to_wait: int):
        time.sleep(seconds_to_wait - 2)
        # TODO: make request to refresh

    def _create_payment(self, amount_coins: int):
        transaction_id = str(uuid.uuid4())
        response = self._session.post(f"{self._url}/pga/transactions")
