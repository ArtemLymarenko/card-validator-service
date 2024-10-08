# Card Validator Project
This service implements an API for validating user credit cards using Luhn algorithm.

## API 
### `/api/v1` 

## API Endpoints
### `POST` /card/validate

#### Parameters 
`None`

#### Request body
- ``number`` **string**: The credit card number to validate.
- ``expYear`` **int**: Expiration year.
- ``expMonth`` **int**: Expiration month.

#### Response Body
```json
  {
    "valid": true,
    "error": {
        "code": "001",
        "message": "Some message"
    }
  }
```
#### Response Body Status Codes
```
001: Invalid Card Number
002: Invalid Card Year
003: Invalid Card Month
004: Bad Request
```

## How to run
Clone the repository to your local machine:
```sh
git clone https://github.com/ArtemLymarenko/card-validator-service.git
cd card-validator-service
```

***!!!*** Fill `.env` file in the root directory with the necessary environment variables, using `.env.example` (just copy all data)
### Make
To build binary file:
```sh
make build
```
Start docker container:
```sh
make up
```
Stop docker containers:
```sh
make down
```

### Docker Compose
To build binary file:
```sh
docker-compose up --build -d
```
Start docker container:
```sh
docker-compose up -d
```
Stop the containers:
```sh
docker-compose down
```

