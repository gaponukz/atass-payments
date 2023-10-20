import uuid
import typing
from src.domain import entities
from src.application import dto


class PaymentRepository(typing.Protocol):
    def save(self, data: dto.SavePaymentDTO):
        ...


class PaymentExternalAPI(typing.Protocol):
    def create_payment(
        self, data: dto.CreateExternalPaymentDTO
    ) -> dto.ExternalPaymentCreatedDTO:
        ...


class TextPresentationService(typing.Protocol):
    def get_text(
        self, payment: dto.CreatePaymentDTO
    ) -> dto.TextPresentaionOnPaymentCreationDTO:
        pass


class CreatePaymentUseCase:
    def __init__(
        self,
        storage: PaymentRepository,
        api: PaymentExternalAPI,
        presentation: TextPresentationService,
    ):
        self._storage = storage
        self._api = api
        self._presentation = presentation

    def create_payment(self, data: dto.CreatePaymentDTO) -> str:
        payment_id = str(uuid.uuid4())
        presentation = self._presentation.get_text(data)

        created = self._api.create_payment(
            dto.CreateExternalPaymentDTO(
                id=payment_id,
                amount=data.amount,
                ttl_seconds=300,
                after_payment_url=presentation.after_payment_url,
                title=presentation.title,
                description=presentation.description,
                short_description=presentation.short_description,
            )
        )

        self._storage.save(
            dto.SavePaymentDTO(
                external_id=created.id,
                payment=entities.Payment(
                    id=payment_id,
                    amount=data.amount,
                    route_id=data.route_id,
                    passenger=data.passenger,
                ),
            )
        )

        return created.url
