# Vette Tracker Services

The API services for the Vette Tracker application. Written in Go, utilizes PostgreSQL for the database.

## To Run

`docker compose up --build`

This will start the API server and the PostgreSQL database.

## Project Structure

See the [project structure](docs/structure.md) for more information on the layout of the project.

## API TODOs

- ~~Separate business logic from handlers~~
- ~~Use Interfaces~~
- ~~Create custom errors and error handling~~
- Logging
- Configuration - move database configuration to a dedicated config file
- Testing
- ~~Caching~~
- ~~Input sanitization~~
- OpenAPI Spec generation

## Larger TODOs

- Modify UI to call these services
- Integrate with Clerk API for authentication and authorization
- Figure out how to deploy to Digital Ocean droplet
- Implement a CI/CD pipeline to autodeploy changes
- Figure out how to deploy a dev and prod instance on the same droplet
