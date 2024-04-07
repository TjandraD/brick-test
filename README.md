# Money Transfer Service - Brick

This is a Money Transfer service that is able to validate financial account number, transfer/disburse money, and receive a callback when the transaction is completed.

# General Overview

![General Overview](overview.png)

The whole process consists of 3 components:
- This service
- A mock bank/financial service
- PostgreSQL database

The service is responsible for:
- Validating the account number via mockAPI
- Transfer money to the account via mockAPI & insert the transaction data into the database
- Receive a callback from the mockAPI when the transaction is completed and update the data into the database

# Requirements

Required things before running this service:
- Go version 1.22.1 or above
- Docker and Docker Compose installed
- Echo library for HTTP webserver (see [this reference](https://echo.labstack.com/docs/quick-start))
- env package ([installation](https://github.com/caarlos0/env?tab=readme-ov-file#example))
- Gorm & PostgreSQL driver ([installation](https://gorm.io/docs/index.html))

# How to Run

Run this command to copy the env file required to run the service:
```bash
make setup
```

Then, run docker that will start the service & the database:
```bash
make build
make run-docker
```
