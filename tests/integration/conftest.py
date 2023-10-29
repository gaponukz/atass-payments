import peewee
import pytest
import requests
import requests_mock
from src.application.usecases.create_payment import CreatePaymentUseCase
from src.application.usecases.handle_payment import HandlePaymentUseCase
from src.infrastructure.persistent.repository import PaymentsRepository
from src.infrastructure.persistent.models import PaymentModel
from src.infrastructure.session.client import HttpPaymentSession, SessionCreds
from src.infrastructure.text_presentation.service import TextPresentationService


session = requests.Session()
adapter = requests_mock.Adapter()
session.mount("mock://", adapter)

adapter.register_uri(
    "POST",
    "mock://innsmouth.payhub.com.ua/auth/token",
    json={
        "data": {
            "access_token": "",
            "expires_in": 999999,
            "refresh_expires_in": 99999,
            "refresh_token": "",
        }
    },
)

api = HttpPaymentSession(
    SessionCreds(
        url="mock://innsmouth.payhub.com.ua",
        login="test_login",
        password="test_password",
        merchant_config_id="bla",
        config_id="bla",
    ),
    session,
)

db = peewee.SqliteDatabase(":memory:")
db.connect()
text_presentation = TextPresentationService()


@pytest.fixture
def create_payment_usecase():
    storage = PaymentsRepository(db)
    storage.prepare_on_first_startup()

    return CreatePaymentUseCase(storage, api, text_presentation)


@pytest.fixture
def handle_payment_usecase():
    storage = PaymentsRepository(db)
    storage.prepare_on_first_startup()

    return HandlePaymentUseCase(storage, api)


@pytest.fixture
def db_model():
    return PaymentModel
