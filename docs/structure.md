# Project Structure

## Entry Point

The entry point of the project is the cmd/api/main.go file. This file is responsible for starting the server and setting up the routes.

## /internal

The /internal directory contains all the application-specific code. It contains the following subdirectories. This is how data flows throughout the app:

Request → Handler → Service → Repository → Database
Response ← Handler ← Service ← Repository ← Database

### /handlers

The /handlers directory contains the HTTP request handlers for the application. Each file in this directory is responsible for handling a specific set of routes.

#### Things to Consider:

Consider these things when writing handlers:

- Inputs Parameters
- Input Validation
- Authentication and Authorization (using Clerk API)
- Error Handling
- Security
  - Input sanitization

### /services

The /services directory contains the business logic for the application.

### /repository

The /repository directory contains the database repository for the application. This is the data access layer.

### /models

The /models directory contains the data models for the application. Each file in this directory is responsible for defining a specific data model.

## Libraries

- Gin is used to scaffold the API server and handle routing
