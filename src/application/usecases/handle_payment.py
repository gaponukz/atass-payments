import time
import typing
import threading
from src.application import dto


class PaymentRepository(typing.Protocol):
    def set_status(self, payment_id: str, status: dto.PaymentStatus):
        ...

    def submit_payment(self, payment_id: str):
        ...

    def get_unprocessed_payments(self) -> list[str]:
        ...


class PaymentExternalAPI(typing.Protocol):
    def get_payment_status(self, payment_id: str) -> dto.PaymentStatus:
        ...


class HandlePaymentUseCase:
    def __init__(self, storage: PaymentRepository, api: PaymentExternalAPI):
        self._storage = storage
        self._api = api

    def handle_unprocessed(self):
        for payment_id in self._storage.get_unprocessed_payments():
            threading.Thread(target=self.handle, args=(payment_id, 5, 8)).start()

    def handle(self, payment_id: str, retries: int, wait_before_check: int):
        try:
            self._handle(payment_id, retries, wait_before_check)

        finally:
            self._storage.submit_payment(payment_id)

    def _handle(self, payment_id: str, retries: int, wait_before_check: int):
        for _ in range(retries):
            status = self._api.get_payment_status(payment_id)

            if status in [dto.PaymentStatus.DONE, dto.PaymentStatus.NOT_EXIST]:
                self._storage.set_status(payment_id, status)
                return

            time.sleep(wait_before_check)

        self._storage.set_status(payment_id, dto.PaymentStatus.NOT_EXIST)
