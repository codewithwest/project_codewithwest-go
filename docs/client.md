# Client Documentation

## Table of Contents

- [Overview](#overview)
- [Usages](#usages)
  - [Mutations](#mutations)
  - [Queries](#queries)

## Overview

This documentation cover all client specific implementation done in the project.

## Usages

### Mutations:

#### **_CreateClient_**

#### Description:

Creates a new client in the system with the provided information and returns the created client data along with a session token.

#### Input Parameters:

- email: String (required) - Client's email address
- userName: String (required) - Client's username
- password: String (required) - Client's password
- company_name: String (required) - Client's company name

#### Process:

1. Validates all input parameters
2. Checks for existing email and username to prevent duplicates
3. Hashes the password for secure storage
4. Generates a unique user ID
5. Creates and stores the client document in MongoDB
6. Creates a session for the new client
7. Generates a new token for client

#### Sample:

```
mutation {
  createClient(
    input: {
      password: "jonPass1"
      company_name: "doeComapny"
      username: "jonDOe"
      email: "jone@mail.com"
    }
  ) {
    company_name
    created_at
    email
    id
    last_login
    status
    token
    type
    updated_at
    username
  }
}
```

Possible Errors:

- "input validation failed" - Invalid or missing required fields
- "email already exists" - Email is already registered
- "username already exists" - Username is already taken
- "database connection failed" - MongoDB connection issues
- "password hashing failed" - Error in password encryption
- "client creation failed" - Error while saving to database
- "session creation failed" - Error creating user session

#### Returns

```
{
  "data": {
    "createClient": {
      "company_name": "doeComapny",
      "created_at": "12-04-2025 20:21:35",
      "email": "jone@mail.com",
      "id": "7",
      "last_login": null,
      "status": "active",
      "token": "token-sample",
      "type": "client",
      "updated_at": "12-04-2025 20:21:35",
      "username": "jonDOe"
    }
  }
}
```

### Queries

#### To-do

[Back to main](../README.md#features)
