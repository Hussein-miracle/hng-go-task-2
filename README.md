
# Person API Documentation

This API allows you to manage person records in a MongoDB database using the Go programming language and the Gin web framework. You can perform basic CRUD (Create, Read, Update, Delete) operations on person records.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the API](#running-the-api)
- [API Endpoints](#api-endpoints)
- [Error Handling](#error-handling)

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your system.
- MongoDB installed and running.

## Installation

To set up the API, follow these steps:

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/Hussein-miracle/hng-go-task-2
   ```
2. Change to the project directory:
   ```bash
   cd hng-go-task-2
   ```
3. Install the required Go packages using:
   ```bash
   go get .
   ```

## Configuration

Before running the API, you need to configure your MongoDB connection settings. Add a `.env` file with your MONGODB_URI gotten from your MongoDB database:

   ```
   MONGODB_URI='Replace with URI from Mongodb'
   ```

## Running the API

To run the API, follow these steps:

1. Start your MongoDB server if it's not already running.

2. Run the following command from the project directory:
   ```bash
   go run .
   ```
   This will start the API server on port 8080 by default.

3. The API is now running and ready to accept requests.

## API Endpoints
The API provides the following endpoints:

* `GET /person/:param`: Retrieve a person record by ID or name.
* `POST /person`: Create a new person record.
* `PUT /person/:param`: Update an existing person record by ID or name.
* `DELETE /person/:param`: Delete a person record by ID or name.

For detailed information on each endpoint, refer to the [API Documentation](https://github.com/Hussein-miracle/hng-go-task-2/blob/master/DOCUMENTATION.md).

## Error Handling
The API returns appropriate HTTP status codes and error messages for different scenarios. Refer to the Error Handling section in the API Documentation for more details.
