from src.application import dto
from src.infrastructure.persistent.dto import PaymentDTO


class PaymentsRepository:
    def __init__(self):
        self.payments: dict[str, PaymentDTO] = {}

    def save(self, data: dto.SavePaymentDTO):
        self.payments[data.external_id] = PaymentDTO(
            status=dto.PaymentStatus.EXIST, payment=data.payment
        )

    def set_status(self, payment_id: str, status: dto.PaymentStatus):
        self.payments[payment_id].status = status

    def submit_payment(self, payment_id: str):
        print(f"Sending payment: {self.payments[payment_id]}")
        del self.payments[payment_id]
