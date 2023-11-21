# Make a transaction in bank account

## Acceptance Criteria

Write an API to create a new transaction for an existing customer

- transaction can only be "withdrawl" or "deposit"
- amount cannot be negative
- withdrawal amount should be available in the account
- succesful transaction, should return the upadated balance with transaction id and response
- error handling should be done for bad request, validation and unexpected errors form the server side and should return the appropriate http status code with message
