from src.application import dto


class TextPresentationService:
    def get_text(
        self, payment: dto.CreatePaymentDTO
    ) -> dto.TextPresentaionOnPaymentCreationDTO:
        return dto.TextPresentaionOnPaymentCreationDTO(
            title=f"Оплата за автобусний квиток",
            description=f"Шановний {payment.passenger.full_name}, дякуємо вам за вибір Atass для ваших подорожей.",
            short_description=f"Шановний {payment.passenger.full_name}, дякуємо вам за вибір Atass для ваших подорожей.",
            after_payment_url="http://localhost:8000",
        )
