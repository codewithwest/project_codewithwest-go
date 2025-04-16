# Client Documentation

## Table of Contents

- [Overview](#overview)
- [Usages](#usages)
    - [Mutations](#mutations)
    - [Queries](#queries)

## Overview

This documentation cover all adminUser specific implementation done in the project.

## Usages

### Mutations:
#### Description: 
    Creates a new administrative user in the system with the 
    provided credentials and user information. Performs validation, 
    password hashing, and ensures email uniqueness before creating 
    the user record.

#### Input Parameters:
```
    email: String (required) - Admin user's email address
    password: String (required) - Admin user's password
    username: String (required) - Admin user's username
    first_name: String (required) - Admin user's first name
    last_name: String (required) - Admin user's last name
    phone: String (required) - Admin user's phone number
```

#### Process:
- Validates all input parameters for completeness and correctness
- Establishes connection to MongoDB admin_users collection
- Hashes the provided password securely
- Checks for existing email to prevent duplicates
- Generates new user ID based on highest existing ID
- Creates new admin user record with provided information
- Returns created user data

#### Sample:
```
mutation {
    createAdminUser(
        input: {
            email: "admin@example.com"
            password: "SecurePass123!"
            username: "admin_user"
        }
        ) {
            id
            email
            username
            created_at
            updated_at
            status
    }
}
```
#### Possible Errors:
- "invalid input arguments" - Missing or malformed input data
- "password hashing failed" - Error in password security processing
- "database connection error" - MongoDB connection issues
- "email already exists" - Duplicate email address
- "user creation failed" - Error inserting new user record
- "user retrieval failed" - Error fetching created user data
- "internal server error" - Unexpected server-side issues

#### Returns:
```
{
  "data": {
    "createAdminUser": {
      "created_at": "16-04-2025 17:34:47",
      "email": "strisng@west",
      "id": "3",
      "last_login": null,
      "role": "administrator",
      "type": "user",
      "updated_at": "16-04-2025 17:34:47",
      "username": "string"
    }
  }
}
```

#### Security Considerations:
- Implements secure password hashing
- Validates input data before processing
- Prevents duplicate email addresses
- Uses context timeout to prevent hanging operations
- Returns generic error messages to prevent information leakage
- Uses proper error wrapping for debugging
- Implements database connection timeout

#### Technical Details:
- Context Timeout: 60 seconds
- Database Collection: "admin_users"
- Password Hashing: Uses secure hashing algorithm
- ID Generation: Auto-incrementing based on highest existing ID

#### Performance Notes:
- Optimized database queries
- Single database connection per request
- Efficient error handling flow
- Minimal memory usage
- Atomic ID generation process

### Queries:

[Back to main](../README.md#features)

