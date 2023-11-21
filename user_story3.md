# Write an API to create a new account for an existing customer

## Acceptance criteria

- A new account can only be opened with a minimum deposit of 5000.00
- Account can only be of saving or checking type
- In case of an unexpected error, API should return status code 500 (Internal Server Error) along with the error message
- The API should return the new account id, when thenew account is opened with the status code as 201 (CREATED)
