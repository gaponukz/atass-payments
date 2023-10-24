import peewee
import datetime
from src.application import dto

database_proxy = peewee.DatabaseProxy()


def _utc_now() -> datetime.datetime:
    return datetime.datetime.now(datetime.timezone.utc)


class PaymentModel(peewee.Model):
    external_id = peewee.UUIDField(primary_key=True)
    status = peewee.IntegerField(default=dto.PaymentStatus.EXIST)
    is_submitted = peewee.BooleanField(default=False)
    payment_info = peewee.TextField()

    class Meta:
        table_name = "payments"
        database = database_proxy


class OutBoxModel(peewee.Model):
    payment_info = peewee.TextField()
    append_datetime = peewee.DateTimeField(default=_utc_now)

    class Meta:
        table_name = "outbox"
        database = database_proxy
