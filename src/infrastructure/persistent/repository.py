import json
import typing

import peewee
import dataclass_factory

from src.application import dto
from src.infrastructure.persistent.models import (
    PaymentModel,
    OutBoxModel,
    database_proxy,
)


class PaymentsRepository:
    def __init__(self, db: peewee.Database):
        self._db = db
        self._dataclass_factory = dataclass_factory.Factory()

        database_proxy.initialize(self._db)

    def prepare_on_first_startup(self):
        self._db.create_tables([PaymentModel, OutBoxModel])

    def save(self, data: dto.SavePaymentDTO):
        PaymentModel.create(
            external_id=data.external_id,
            payment_info=json.dumps(self._dataclass_factory.dump(data.payment)),
        ).save()

    def set_status(self, payment_id: str, status: dto.PaymentStatus):
        PaymentModel.update(status=status).where(
            PaymentModel.external_id == payment_id
        ).execute()

    def submit_payment(self, payment_id: str):
        payment = typing.cast(
            PaymentModel, PaymentModel.get(PaymentModel.external_id == payment_id)
        )

        with self._db.atomic():
            if payment.status == dto.PaymentStatus.DONE:
                OutBoxModel.create(payment_info=payment.payment_info).save()

            payment.is_submitted = True
            payment.save(only=[PaymentModel.is_submitted])

    def get_unprocessed_payments(self) -> list[str]:
        unprocessed_payments: list[str] = []

        for payment in (
            PaymentModel.select().where(PaymentModel.is_submitted == False).execute()
        ):
            payment = typing.cast(PaymentModel, payment)
            unprocessed_payments.append(str(payment.external_id))

        return unprocessed_payments
