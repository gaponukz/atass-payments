# Atass service for managing payments &amp; adding passengers
## On successful payment callback

- URL: `/processPayment`
- Method: `POST`
- Description: Send notification about passanger's payment 
- Request body example:
```json
{
    "amount": 50.99,
    "routeId": "7c47bcb9-8179-49b5-93fc-089fafa793d3",
    "passenger": {
        "id": "e707e899-7aed-4ad8-82ed-4963401ed12b",
        "gmail": "john.doe@gmail.com",
        "fullName": "John Doe",
        "phoneNumber": "+1234567890",
        "movingFromId": "93b7f7ea-60f2-4524-82b9-a532721f2596",
        "movingTowardsId": "82ebef20-cde6-4b87-be3d-3030b8fe481d"
    }
}
```
- Response example:
```
f3c77884-3185-11ee-be56-0242ac120002
```
