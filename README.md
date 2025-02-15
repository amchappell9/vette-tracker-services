# Vette Tracker Services

The API services for the Vette Tracker application. Written in Go, utilizes PostgreSQL for the database.

## To Run

`docker compose up --build`

This will start the API server and the PostgreSQL database.

## Project Structure

See the [project structure](docs/structure.md) for more information on the layout of the project.

## TODOs

- ~~Separate business logic from handlers~~
- ~~Use Interfaces~~
- Create custom errors and error handling
- Configuration - move database configuration to a dedicated config file

- Logging
- Testing
- Caching
- Input sanitization
