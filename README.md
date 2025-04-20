# Reuros API

Reuros API is a simple service that provides currency conversion rates between two currencies. It fetches exchange rates using the [ExchangeRate-API](https://www.exchangerate-api.com/).

## Features

- Fetch exchange rates for a given base and target currency.
- Swagger documentation for API endpoints.

## Prerequisites

- Docker and Docker Compose installed on your system.
- An API key from [ExchangeRate-API](https://www.exchangerate-api.com/).

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/reuros-api.git
   cd reuros-api
   ```

2. Create a `.env` file in the root directory and add your ExchangeRate-API key:
   ```properties
   EXCHANGE_RATE_API_KEY=your_api_key_here
   ```

3. Build and run the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

4. The application will be available at `http://localhost:8080`.

## API Endpoints

### Get Currency Conversion Rate

- **Endpoint**: `/currency/{base}/{target}`
- **Method**: `GET`
- **Description**: Fetch the conversion rate between two currencies.
- **Parameters**:
  - `base` (path): Base currency code (e.g., USD).
  - `target` (path): Target currency code (e.g., EUR).
- **Responses**:
  - `200 OK`: Returns the conversion rate.
  - `400 Bad Request`: Invalid input or URL format.
  - `401 Unauthorized`: Missing or invalid API key.
  - `500 Internal Server Error`: Server error or failed to fetch rates.

Example:
```bash
curl http://localhost:8080/currency/USD/EUR
```

## Swagger Documentation

Swagger documentation is available at:
```
http://localhost:8080/swagger/
```

To generate a new version of the Swagger documentation, run the following command:
```bash
swag init -g cmd/server/main.go --output docs
```

## Deployment

The project is deployed using the infrastructure provided by [Railway](https://railway.app). You can access the services at the following URLs:

- **Backend Swagger**: [https://currency-api-production-01.up.railway.app/swagger/index.html](https://currency-api-production-01.up.railway.app/swagger/index.html)
- **UI**: [https://currency-app-production-01.up.railway.app/](https://currency-app-production-01.up.railway.app/)

## Project Structure

- `src/handlers`: Contains HTTP handlers for API endpoints.
- `src/services`: Contains business logic for fetching currency rates.
- `src/models`: Defines data models used in the application.
- `src/docs`: Auto-generated Swagger documentation files.

## Environment Variables

- `EXCHANGE_RATE_API_KEY`: API key for accessing the ExchangeRate-API.
- `PORT`: Port on which the application runs (default: 8080).
