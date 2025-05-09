# Admin Documentation

## Table of Contents

- [Overview](#overview)
- [Mutations](#mutations)
  - [Create Admin User](#create-admin-user)
  - [Create Admin User Request](#create-admin-user-requests)
- [Queries](#queries)
  - [Login Admin User](#login-admin-user) 
  - [Get Admin Users](#get-admin-users)
  - [Get Admin User Requests](#get-admin-user-requests)

## Overview

This documentation cover all adminUser specific implementation done in the project.

### Mutations:

### Create Admin User

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
            created_at
            email
            id
            last_login
            password
            role
            type
            updated_at
            username
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

### Create Admin User Requests
#### Description:

Creates a new administrative user request in the system. This function validates 
the requester's authorization, checks email uniqueness, and stores the request 
with appropriate timestamps.

#### Input Parameters:

```
  email: String (required) - Requested email address for admin account
```

#### Process:

1. Validates user authorization and permissions
2. Validates email format and uniqueness
3. Establishes connection to MongoDB with timeout settings
4. Generates new unique ID for the request
5. Creates request record with timestamp
6. Returns created request data

#### Sample:
``` 
mutation {
  createAdminUserRequest(
    email: "mailto:newadmin@example.com"
  ) {
    id
    email
    created_at
  }
}
```

#### Possible Errors:

- "not authorized" - User authentication failed
- "invalid user id" - Error parsing authorized user's ID
- "missing required argument(s)" - Email parameter not provided
- "invalid email format" - Email format validation failed
- "database connection error" - MongoDB connection issues
- "email already exists" - Duplicate email address
- "error generating user ID" - ID generation failed
- "failed to create user" - Request creation failed
- "internal server error" - Unexpected server-side issues

#### Returns:
``` 
{
  "data": {
    "createAdminUserRequest": {
      "id": "8",
      "email": "mailto:newadmin@example.com",
      "created_at": "2024-01-20T15:30:00Z",
    }
  }
}
```

#### Security Considerations:

- Implements user authorization validation
- Uses context timeout to prevent hanging operations
- Returns generic error messages to prevent information leakage
- Validates input data before processing
- Uses secure database connections
- Implements proper error wrapping
- Validates email format
- Checks for duplicate emails
- Uses atomic operations for ID generation

#### Technical Details:

- Context Timeout: 30 seconds
- Database Collection: "admin_user_request"
- Email Validation: RFC 5322 standards
- ID Generation: Auto-incrementing based on collection maximum
- Status Field: Default "pending"
- Timestamp Format: UTC RFC3339

#### Performance Notes:
- Optimized database queries
- Single database connection per request
- Efficient error handling flow
- Minimal memory usage
- Atomic ID generation
- Proper resource cleanup
- Index-based email uniqueness check
- Streamlined request creation process
- Efficient timestamp handling

### Queries:

### Login Admin User

#### Description: 
    Authenticates an administrative user in the system using their email 
    and password credentials. Performs validation, password verification, 
    and creates a session token upon successful authentication.

#### Input Parameters:

```
    email: String (required) - Admin user's email address
    password: String (required) - Admin user's password
```

#### Process:

- Validates input parameters for completeness and correctness
- Establishes connection to MongoDB admin_users collection
- Retrieves user record by email
- Verifies password against stored hash
- Creates new session token for authenticated user
- Returns authentication data including token

#### Sample:

```
mutation {
  loginAdminUser(
    input: {
      email: "mailto:admin@example.com"
      password: "SecurePass123!"
    }
    ) {
      token
      id
      email
  }
}
```

#### Possible Errors:

- "invalid input arguments" - Missing or malformed input data
- "database connection error" - MongoDB connection issues
- "invalid email or password combination" - Authentication failure
- "internal server error" - Unexpected server-side issues
- "session creation failed" - Error generating authentication token

#### Returns:

```
{
  "data": {
    "loginAdminUser": {
      "token": "eyJhbGciOiJIUzI1NiIs...",
      "id": "3",
      "email": "mailto:admin@example.com"
    }
  }
}

```

#### Security Considerations:
- Implements secure password verification
- Uses generic error messages to prevent user enumeration
- Implements context timeout for security
- Uses proper error wrapping for secure debugging
- Maintains consistent response timing to prevent timing attacks
- Returns minimal user information in response
- Implements secure session token generation

#### Technical Details:

- Context Timeout: 30 seconds
- Database Collection: "admin_users"
- Password Verification: Uses secure comparison
- Session Token: JWT-based authentication
- Error Handling: Implements proper error wrapping

#### Performance Notes:

- Optimized database queries using indexes
- Single database lookup per request
- Efficient error handling flow
- Quick response times for better user experience
- Minimal memory footprint
- Connection pooling for database operations
- Cached session management

### Get Admin Users

#### Description: 
    Retrieves a paginated list of administrative users from the system. 
    Requires administrator privileges and implements pagination for 
    efficient data retrieval and display.

#### Input Parameters:

``` 
  limit: Int (optional) - Number of items per page (default: 10)
  page: Int (optional) - Page number to retrieve (default: 1)

```

#### Process:
- Validates user authorization and administrator privileges
- Implements pagination with configurable page size
- Retrieves total count of users for pagination metadata
- Sorts users by ID in ascending order
- Returns paginated list with metadata
- Performs proper resource cleanup

#### Sample:

```
query {
  getAdminUsers(limit: 10, page: 1) {
    data {
      id
      email
      username
      role
      created_at
      updated_at
      last_login
    }
    page
    totalPages
    totalItems
  }
}
```

#### Possible Errors:

- "not authorized" - User authentication failed
- "invalid user id" - Error parsing user identifier
- "database connection error" - MongoDB connection issues
- "user not found" - Requesting user not found in system
- "access denied: administrator privileges required" - User lacks admin role
- "error counting documents" - Pagination counting failed
- "error fetching users" - Database query failed
- "no users found" - Empty result set
- "error decoding users" - Data parsing error
- "internal server error" - Unexpected server-side issues

#### Returns:

```
{
  "data": {
  "getAdminUsers": {
    "data":
    {
      "id": "1",
      "email": "mailto:admin@example.com",
      "username": "admin",
      "role": "administrator",
      "created_at": "2024-01-20T10:00:00Z",
      "updated_at": "2024-01-20T10:00:00Z",
      "last_login": "2024-01-20T15:30:00Z"
    }
    "page": 1,
    "totalPages": 5,
    "totalItems": 48
    }
  }
}
```

#### Security Considerations:

- Implements user authorization checks
- Validates administrator privileges
- Uses context timeout for security
- Returns generic error messages
- Implements proper resource cleanup
- Uses secure database connections
- Validates input parameters

#### Technical Details:

- Context Timeout: 10 seconds
- Database Collection: "admin_users"
- Default Page Size: 10 items
- Sorting: By ID ascending
- Pagination: Skip-based implementation
- Response Type: AdminUsersPaginatedResponse

#### Performance Notes:

- Uses efficient cursor-based pagination
- Implements proper cursor cleanup
- Uses index-based sorting
- Optimized document counting
- Efficient batch document retrieval
- Proper connection handling
- Memory-efficient result processing

### Get Admin User Requests

#### Description: 
    Retrieves a paginated list of administrative user requests from the system. 
    This endpoint requires administrator privileges and implements pagination 
    for efficient data retrieval and display. Requests are sorted by creation 
    date with newest first.

#### Input Parameters:

```
  limit: Int (optional) - Number of items per page (default: 10)
  page: Int (optional) - Page number to retrieve (default: 1)
```

#### Process:

- Validates user authorization and administrator privileges
- Implements pagination with configurable page size
- Retrieves total count of requests for pagination metadata
- Sorts requests by creation date in descending order
- Returns paginated list with metadata
- Performs proper resource cleanup

#### Sample:

```
query {
  getAdminUserRequests(limit: 10, page: 1) {
    data {
      id
      email
      created_at
    }
    page
    totalPages
    totalItems
  }
}
```

#### Possible Errors:

- "not authorized" - User authentication failed
- "invalid user id" - Error parsing user identifier
- "database connection error" - MongoDB connection issues
- "user not found" - Requesting user not found in system
- "access denied: administrator privileges required" - User lacks admin role
- "error counting documents" - Pagination counting failed
- "error fetching requests" - Database query failed
- "no requests found" - Empty result set
- "error decoding requests" - Data parsing error
- "internal server error" - Unexpected server-side issues

#### Returns:

```
{
"data": {
  "getAdminUserRequests": {
    "data": {
      "id": "1",
      "email": "mailto:requestor@example.com",
      "created_at": "2024-01-20T10:00:00Z",
    }
    "page": 1,
    "totalPages": 5,
    "totalItems": 48
    }
  }
}
```

#### Security Considerations:

- Implements user authorization checks
- Validates administrator privileges
- Uses context timeout for security
- Returns generic error messages
- Implements proper resource cleanup
- Uses secure database connections
- Validates input parameters

#### Technical Details:

- Context Timeout: 10 seconds
- Database Collection: "admin_user_request"
- Default Page Size: 10 items
- Sorting: By created_at descending
- Pagination: Skip-based implementation
- Response Type: AdminUserRequestsPaginatedResponse

#### Performance Notes:

- Uses efficient cursor-based pagination
- Implements proper cursor cleanup
- Uses index-based sorting
- Optimized document counting
- Efficient batch document retrieval
- Proper connection handling
- Memory-efficient result processing

[Back to main](../README.md#features)
