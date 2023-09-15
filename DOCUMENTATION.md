# API Documentation for Person App

This documentation provides an overview of the API for managing person records using the Go programming language and the Gin web framework. The API allows you to perform basic CRUD (Create, Read, Update, Delete) operations on person records stored in a MongoDB database.

A documentation on postman is also included here [LINK](https://documenter.getpostman.com/view/20087483/2s9YC7RWQg) 

## Endpoints


### Create Person
**Description:** Update an existing person record by their ID or name.

**Endpoint:** `/person`

**Method:** `POST`

**Parameters:**
- `param` Create a new person record.

**Request Body:**
  ```json
  {
    "name": "Person's Name"
  }
  ```
**Success Response:**
  - **Status Code: 200 (OK)**
    ```json
    {
      "status": "success",
      "data": {
        "ID": "ObjectID",
        "name": "Person's Name",
        "createdAt": "Timestamp",
        "updatedAt": "Timestamp"
      },
      "message": "Person created successfully"
    }
    ```
**Error Responses:**
  - **Status Code: 500 (Internal Server Error)**
    ```json
    {
        "error": "Internal serval error",
        "status": "error"
    }
    ```
  - **Status Code: 400 (Bad Request)**
    ```json
    {
        "error": "an error occured",
        "status": "error"
    }
    ```

    

### Update Person
**Description:** Update an existing person record by their ID or name.

**Endpoint:** `/person/:param`

**Method:** `DELETE`

**Parameters:**
- `param` (string): The ID (ObjectID) or name of the person to update.

**Request Body:**
  ```json
  {
    "name": "New Person's Name"
  }
  ```
**Success Response:**
  - **Status Code: 200 (OK)**
    ```json
    {
      "success": "success",
      "message": "Person updated successfully",
      "data": {
        "ID": "ObjectId",
        "name": "Updated Person's Name",
        "createdAt": "Timestamp",
        "updatedAt": "Timestamp"
      }
    }
    ```
**Error Responses:**
  - **Status Code: 500 (Internal Server Error)**
    ```json
    {
        "error": "an error occured",
        "success": "error"
    }
    ```
  - **Status Code: 404 (Not Found)**
    ```json
    {
        "error": "person not found",
        "success": "error"
    }
    ```
  - **Status Code: 400 (Bad Request)**
    ```json
    {
        "error": string,
        "success": "error"
    }
    ```


### Get Person
**Description:** Retrieve a person record by either their ID or name.

**Endpoint:** `/person/:param`

**Method:** `GET`

**Parameters:**
- `param` (string): The ID (ObjectID) or name of the person to retrieve.

**Success Response:**
  - **Status Code: 200 (OK)**
    ```json
    {
      "status": "success",
      "data": {
        "ID": "ObjectId",
        "name": "Person's Name",
        "createdAt": "Timestamp",
        "updatedAt": "Timestamp"
      }
    }
    ```
**Error Responses:**
  - **Status Code: 500 (Internal Server Error)**
    ```json
    {
        "error": "an error occured",
        "status": "error"
    }
    ```
  - **Status Code: 404 (Not Found)**
    ```json
    {
        "error": "person not found",
        "status": "error"
    }
    ```



### Delete Person
**Description:** Delete a person record by their ID or name.

**Endpoint:** `/person/:param`

**Method:** `DELETE`

**Parameters:**
- `param` (string): The ID (ObjectID) or name of the person to delete.

**Success Response:**
  - **Status Code: 200 (OK)**
    ```json
    {
      "status": "success",
      "message": "User <Person's Name> deleted"
    }
    ```
**Error Responses:**
  - **Status Code: 500 (Internal Server Error)**
    ```json
    {
        "error": "an error occured",
        "status": "error"
    }
    ```
  - **Status Code: 404 (Not Found)**
    ```json
    {
        "error": "person not found",
        "status": "error"
    }
    ```

### Known Limitations
This API doesn't have authentication or authorization

### Setup and Deployment
Refer to the setup instructions in the [README.md](https://github.com/Hussein-miracle/hng-go-task-2/edit/master/README.md) file in this repo
