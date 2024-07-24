# B3 Trades Application

## Overview

The B3 Trades application was developed to process financial trading data files from B3 (Brazilian Stock Exchange) 
and insert them into a PostgreSQL database. It also provides this data via an API, allowing queries for the following data:
```json
{
   "ticker": "PETR4",
   "max_range_value": 0,
   "max_daily_volume": 0,
   "TradedQuantity": 1000
}
```

## Features

The application performs the following tasks:

- Reads financial trading data files from a specified directory.
- Parses and processes the data into structured `Trade` objects.
- Inserts batches of `Trade` data into the PostgreSQL database.
- Supports concurrency to optimize file processing using goroutines.
- Clears the database table before inserting new data.

```go
type Trade struct {
	ID              string    // Unique trade ID
	Ticker          string    // Instrument code
	TradePrice      float64   // Trade price
	TradedQuantity  int       // Traded quantity
	ClosingTime     string    // Trade closing time (string format)
	TradeDate       time.Time // Trade date
}
```

## How It Works

1. **Initial Setup**:
    - Clone the repository and navigate to the project directory.
    - Ensure that Go is installed (`go version` should return a valid version).
    - Install PostgreSQL and Docker if they are not already installed.

2. **Preparing Data Files**:
    - Download financial trading data files from the official [B3](https://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/cotacoes/) website.
    - Save the downloaded files to a directory of your choice.
    - Note the directory path to set in `DIRECTORY_PATH`.

3. **Environment Configuration**:
    - Create a `.env` file at the root of the project to store database information:
      ```dotenv
      # PostgreSQL configuration
      POSTGRES_USER=your_postgres_user
      POSTGRES_PASSWORD=your_postgres_password
      POSTGRES_DB=your_postgres_db
      POSTGRES_HOST=db
      POSTGRES_PORT=5432
      ```
    - Create a `.env` file in `b3-insert` and fill it as follows:
    - Replace `/path/to/your/files/b3` with the actual path where you saved the B3 files.
    ```dotenv
      # Directory path for files
      DIRECTORY_PATH=/path/to/your/files/b3
      # Maximum number of workers (concurrent processing)
      MAX_WORKERS=20
      # PostgreSQL Database URL
      DATABASE_URL=postgres://user:pass@host:5432/postgres?sslmode=disable
    ```
    - Create a `.env` file in `b3-api` and fill it as follows:
    ```dotenv
      # API Configuration
      APP_PORT=8080
      # PostgreSQL Database URL
      DATABASE_URL=postgres://user:pass@host:5432/postgres?sslmode=disable
    ```

4. **Docker Configuration**:
    - Ensure that Docker is running on your system.
    - Use the provided `docker-compose.yml` file to start the PostgreSQL database:
      ```bash
      docker compose up
      ```

5. **Running B3-Insert**:
    - Build and run the `b3-insert` application:
      ```bash
      go build -o b3-insert ./cmd
      ./b3-insert
      ```
    - To run without building, use:
      ```bash
      go run ./cmd/main.go
      ```
    - The application will start processing files from the specified directory, inserting data into the PostgreSQL database.

6. **Running B3-Api**:
    - Build and run the `b3-api` application:
      ```bash
      go build -o b3-api ./cmd
      ./b3-api
      ```
    - To run without building, use:
      ```bash
      go run ./cmd/main.go
      ```
    - Once `b3-api` is running (via Docker), you can access it at `http://localhost:8080` (assuming the default configuration).

    - Example curl request:
    ```bash
    curl -X GET \
      'http://localhost:8080/api/aggregated-data/PETR4?date=2024-06-28' \
      -H 'Content-Type: application/json'
    ```

7. **Stopping the Application**:
    - To stop the database, use the command `docker compose down`.
    - To stop the application, use `Ctrl + C` in the terminal where it is running.

## Additional Notes

- Ensure you have sufficient permissions and disk space for database operations and data processing.
- Adjust the `MAX_WORKERS` value in the `.env` file according to your system's capabilities and performance requirements.
