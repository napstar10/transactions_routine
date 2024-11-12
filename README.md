# transactions_routine
A test project

## Setup Project

- Clone the repository

- Move to the project's root

- run : `go mod tidy`


## Setup Postgres DB
- Use the database configs provided in the `confg.json` at the root
- Tables will be auto-migrated
- Used `gorm` with `postgres` driver


## Run Unit Tests
- Make sure you are in the root directory
- Run : `go test ./... -v`

## Easy Run
- Give permissions to run.sh : `chmod +x run.sh`
- execute : `./run.sh`
- The above command will run the tests, build and run the application.


## Manual Build & Run
- Move to the directory : transactions_routine/main
- `Run : go build -o transaction_routine ./main`
- execute : `./transaction_routine`

    
## Project Structure

- root 
    - main
      - all main files to start server with DI. Used `wire` as a DI tool
    - api
      - DTO
      - Handlers
      - Services
    - database
        - models
    - config

## Request flow
- A request will reach the server, and redirected to the specific route. Routes are written in `main/server.go`
- A handler mapped to the route will take up the request. Handler is used to marshall and unmarshall the request only
- Each handler will forward the request to the specific service, where the business logic is written.
- Each service will communicate with models and in turn to the database.