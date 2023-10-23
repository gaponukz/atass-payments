import time
import pytest
import dataclasses
from src.application import dto
from src.application.usecases.handle_payment import HandlePaymentUseCase


@dataclasses.dataclass
class Payment:
    status: dto.PaymentStatus
    is_submitted: bool


class PaymentRepositoryMock:
    def __init__(self):
        self.payments: dict[str, Payment] = {
            "1": Payment(status=dto.PaymentStatus.EXIST, is_submitted=False),
            "2": Payment(status=dto.PaymentStatus.EXIST, is_submitted=False),
            "3": Payment(status=dto.PaymentStatus.EXIST, is_submitted=False),
        }

    def set_status(self, payment_id: str, status: dto.PaymentStatus):
        self.payments[payment_id].status = status

    def submit_payment(self, payment_id: str):
        self.payments[payment_id].is_submitted = True

    def get_unprocessed_payments(self) -> list[str]:
        return [
            key for key in list(self.payments) if not self.payments[key].is_submitted
        ]


class PaymentExternalAPIMock:
    def __init__(self, history: dict[str, list[dto.PaymentStatus]], raise_error=False):
        self.history = history
        self._raise_error = raise_error

    def get_payment_status(self, payment_id: str) -> dto.PaymentStatus:
        if self._raise_error:
            raise ValueError("got unexpected exception")

        return self.history[payment_id].pop(0)


def test_hanlde_succeed():
    db = PaymentRepositoryMock()
    api = PaymentExternalAPIMock(
        {
            "1": [
                dto.PaymentStatus.EXIST,
                dto.PaymentStatus.EXIST,
                dto.PaymentStatus.DONE,
            ]
        }
    )
    service = HandlePaymentUseCase(db, api)

    service.handle("1", 5, 0)

    assert db.payments["1"].status == dto.PaymentStatus.DONE
    assert db.payments["1"].is_submitted == True


def test_payment_already_done():
    db = PaymentRepositoryMock()
    api = PaymentExternalAPIMock({"2": [dto.PaymentStatus.DONE]})
    service = HandlePaymentUseCase(db, api)

    service.handle("2", 1, 0)

    assert db.payments["2"].status == dto.PaymentStatus.DONE
    assert db.payments["2"].is_submitted == True


def test_payment_not_exist():
    db = PaymentRepositoryMock()
    api = PaymentExternalAPIMock({"1": [dto.PaymentStatus.NOT_EXIST]})
    service = HandlePaymentUseCase(db, api)

    service.handle("1", 1, 0)

    assert db.payments["1"].status == dto.PaymentStatus.NOT_EXIST
    assert db.payments["1"].is_submitted == True


def test_failure_handle():
    db = PaymentRepositoryMock()
    api = PaymentExternalAPIMock({"3": [dto.PaymentStatus.EXIST]}, True)
    service = HandlePaymentUseCase(db, api)

    with pytest.raises(ValueError):
        service.handle("3", 1, 0)

    assert db.payments["3"].is_submitted == True


def test_handle_unprocessed():
    db = PaymentRepositoryMock()
    api = PaymentExternalAPIMock(
        {
            "1": [dto.PaymentStatus.NOT_EXIST],
            "2": [dto.PaymentStatus.DONE],
            "3": [dto.PaymentStatus.EXIST, dto.PaymentStatus.DONE],
        }
    )

    service = HandlePaymentUseCase(db, api)
    service.handle_unprocessed()

    time.sleep(10)

    for _, value in db.payments.items():
        assert value.is_submitted
