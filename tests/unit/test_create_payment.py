import pytest
from src.domain import entities
from src.application import dto
from src.application.usecases.create_payment import CreatePaymentUseCase


class PaymentRepositoryMock:
    def __init__(self):
        self.last_payment: dto.SavePaymentDTO | None = None

    def save(self, data: dto.SavePaymentDTO):
        self.last_payment = data


class PaymentExternalAPIMock:
    def __init__(self, raise_error=False):
        self._raise_error = raise_error
        self.last_payment: dto.CreateExternalPaymentDTO | None = None

    def create_payment(
        self, data: dto.CreateExternalPaymentDTO
    ) -> dto.ExternalPaymentCreatedDTO:
        if self._raise_error:
            raise ValueError("got unexpected exception")

        self.last_payment = data

        return dto.ExternalPaymentCreatedDTO(
            id="external_id_1", url="https://api.pay/test"
        )


class TextPresentationServiceStub:
    def get_text(
        self, payment: dto.CreatePaymentDTO
    ) -> dto.TextPresentaionOnPaymentCreationDTO:
        return dto.TextPresentaionOnPaymentCreationDTO(
            title="test title",
            description="test description",
            short_description="short description",
            after_payment_url="http://localhost:8000",
        )


def test_successful_payment():
    db = PaymentRepositoryMock()
    api = PaymentExternalAPIMock()
    text = TextPresentationServiceStub()
    service = CreatePaymentUseCase(db, api, text)

    creadted = service.create_payment(
        dto.CreatePaymentDTO(
            amount=1.5005,
            route_id="123",
            passenger=entities.Passenger(
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

    assert creadted.id == "external_id_1"
    assert creadted.url == "https://api.pay/test"

    assert db.last_payment.external_id == creadted.id
    assert db.last_payment.payment.amount == 150
    assert db.last_payment.payment.route_id == "123"
    assert db.last_payment.payment.passenger.full_name == "Alex Afg"

    assert api.last_payment.id == db.last_payment.payment.id
    assert api.last_payment.amount == db.last_payment.payment.amount

    assert api.last_payment.title == "test title"
    assert api.last_payment.description == "test description"
    assert api.last_payment.short_description == "short description"
    assert api.last_payment.after_payment_url == "http://localhost:8000"


def test_failure_payment():
    db = PaymentRepositoryMock()
    api = PaymentExternalAPIMock(raise_error=True)
    text = TextPresentationServiceStub()
    service = CreatePaymentUseCase(db, api, text)

    with pytest.raises(ValueError):
        service.create_payment(
            dto.CreatePaymentDTO(
                amount=1.5005,
                route_id="123",
                passenger=entities.Passenger(
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

    assert db.last_payment is None
