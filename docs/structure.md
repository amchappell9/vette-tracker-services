# Project Structure

## Entry Point

The entry point of the project is the cmd/api/main.go file. This file is responsible for starting the server and setting up the routes.

## /internal

The /internal directory contains all the application-specific code. It contains the following subdirectories:

### /handlers

The /handlers directory contains the HTTP request handlers for the application. Each file in this directory is responsible for handling a specific set of routes.

### /models

The /models directory contains the data models for the application. Each file in this directory is responsible for defining a specific data model.

### /repository

The /repository directory contains the database repository for the application. This is the data access layer.

### /services

The /services directory contains the business logic for the application.
