# Client Documentation

## Table of Contents

- [Overview](#overview)
- [Mutations](#mutations)
  - [Create Client](#create-client) 
- [Queries](#queries)
  - [Authenticate Client](#authenticate-client) 

## Overview

This documentation cover all client specific implementation done in the project.

### Create Client

### Mutations:

#### Description:

Creates a new client in the system with the provided information and returns the 
created client data along with a session token.

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

#### Possible Errors:

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

### Authenticate Client

#### Description:

Authenticates a client in the system using their credentials and returns a session 
token along with client data upon successful authentication.

#### Input Parameters:

```
email: String (required) - Client's registered email address
password: String (required) - Client's password
```
#### Process:

1. Validates input parameters for completeness and correctness
2. Establishes connection to MongoDB with timeout settings
3. Queries database for client record using email
4. Validates provided password against stored hash
5. Creates new session for authenticated client
6. Returns client data with session information

#### Sample:
```
mutation {
  authenticateClient(
    input: {
    email: "jone@mail.com"
    password: "jonPass1"
  }
  ) {
    company_name
    email
    id
    last_login
    status
    token
    type
    username
    updated_at
  }
}
```

#### Possible Errors:

- "invalid input arguments" - Missing or malformed input data
- "invalid email or password" - Empty or incorrect credential format
- "database connection error" - MongoDB connection issues
- "invalid credentials" - Incorrect email/password combination
- "internal server error" - Unexpected server-side issues
- "session creation failed" - Error creating user session

#### Returns:
```
{
  "data": {
    "authenticateClient": {
      "company_name": "doeCompany",
      "email": "jone@mail.com",
      "id": "7",
      "last_login": "12-04-2025 20:21:35",
      "status": "active",
      "token": "session-token-sample",
      "type": "client",
      "updated_at": "12-04-2025 20:21:35",
      "username": "jonDOe"
    }
  }
}
```

#### Security Considerations:

- Uses context timeout to prevent hanging operations
- Implements secure password validation
- Returns generic error messages to prevent information leakage
- Maintains secure session management
- Uses proper error wrapping for debugging
- Implements database connection timeout
- Validates input data before processing

#### Technical Details:

- Context Timeout: 30 seconds
- Database Collection: "clients"

#### required Dependencies:

- graphql
- mongoDB
- clientReusables
- helper (password validation)
- Standard Go packages (context, fmt, errors, strconv)

#### Performance Notes:
- Optimized database queries
- Efficient error handling flow
- Minimal memory usage
- Streamlined authentication process
- Single database query implementation

[Back to main](../README.md#features)
