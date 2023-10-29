import uuid
import json
import typing
from requests_mock import Mocker
from src.domain.entities import Passenger
from src.application.dto import CreatePaymentDTO, PaymentStatus
from src.infrastructure.persistent.models import PaymentModel
from src.application.usecases.create_payment import CreatePaymentUseCase


def test(
    requests_mock: Mocker,
    create_payment_usecase: CreatePaymentUseCase,
    db_model: PaymentModel,
):
    external_id = str(uuid.uuid4())
    requests_mock.post(
        "mock://innsmouth.payhub.com.ua/frames/links/pga",
        json={
            "id": external_id,
            "url": "http://innsmouth.payhub.com/payhere",
            "signature": "123",
        },
    )

    created = create_payment_usecase.create_payment(
        CreatePaymentDTO(
            amount=1.50046,
            route_id="123",
            passenger=Passenger(
                id="1",
                gmail="gmail",
                full_name="Alex Afg",
                phone_number="111",
                move_from_id="1",
                move_to_id="2",
                is_anonymous=False,
            ),
        )
    )

    assert requests_mock.last_request is not None

    requests_data = requests_mock.last_request.json()

    assert requests_data["amount"] == 150
    assert requests_data["title"] == "Оплата за автобусний квиток"

    assert created.id == external_id
    assert created.url == "http://innsmouth.payhub.com/payhere"

    payment = typing.cast(PaymentModel, db_model.get_by_id(external_id))
    assert payment.status == PaymentStatus.EXIST
    assert payment.is_submitted == False

    paymeny_info = json.loads(payment.payment_info)
    assert paymeny_info["amount"] == 150
    assert paymeny_info["route_id"] == "123"
    assert paymeny_info["passenger"]["id"] == "1"
