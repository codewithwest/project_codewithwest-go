# Project categories Documentation

## Table of Contents

- [Overview](#overview)
- [Mutations](#mutations)
  - [Project Categories](#create-project-category)
- [Queries](#queries)
  - [Get Project Categories](#get-project-categories) 
   

## Overview

This documentation cover all project categories specific implementation done in the project.

### Mutations:

### Create Project Category

#### Description: 
    Creates a new project category in the system with the provided name. 
    Performs validation, ensures name uniqueness, and manages MongoDB 
    operations for category creation. Implements auto-incrementing IDs 
    and proper error handling throughout the process.

#### Input Parameters:

```
name: String (required) - Project category name
```

#### Process:

- Validates input parameter for completeness and correctness
- Establishes connection to MongoDB project_categories collection
- Checks for existing category name to prevent duplicates
- Generates new category ID based on highest existing ID
- Creates new project category record with provided information
- Performs document insertion into MongoDB
- Returns created category data

#### Sample:

``` 
  mutation {
      createProjectCategory(
          input: {
              name: "Web Development"
          }
      ) {
          id
          name
          created_at
          updated_at
      }
  }
```

#### Possible Errors:

- missing or invalid name argument - Missing or malformed input data
- failed to connect to database - MongoDB connection issues
- error checking category existence - Category validation failed
- project category already exists - Duplicate category name
- error getting highest ID - ID generation failed
- failed to create project category - Category insertion failed
- invalid inserted ID type - MongoDB ID conversion error
- failed to retrieve created category - Category retrieval failed

#### Returns:

``` 
  {
      data: {
          createProjectCategory: {
              id: 1,
              name: Web Development,
              created_at: 16-04-2025 17:34:47,
              updated_at: 16-04-2025 17:34:47
          }
      }
  }
```

#### Security Considerations:

- Validates input data before processing
- Prevents duplicate category names
- Uses context timeout to prevent hanging operations
- Returns generic error messages to prevent information leakage
- Uses proper error wrapping for debugging
- Implements database connection timeout
- Sanitizes input data before database operations

#### Technical Details:

- Context Timeout: 30 seconds
- Database Collection: &quot;project_categories&quot;
- ID Generation: Auto-incrementing based on highest existing ID

#### Performance Notes:

- Optimized database queries with proper indexing
- Single database connection per request
- Efficient error handling flow
- Minimal memory usage
- Atomic ID generation process
- Context-based timeout management
- Structured query execution path

### Queries:

### Get Project Categories

#### Description:

Retrieves a paginated list of project categories from the system. 
Implements efficient pagination, concurrent database operations, and proper error handling. 
Returns both category data and pagination metadata for client-side navigation.

#### Input Parameters:
```
  limit: Int! (required) - Number of items per page
  page: Int (optional) - Page number (defaults to 1)
```

### Process:

- Validates and processes pagination parameters
- Establishes connection to MongoDB project_categories collection
- Concurrently executes count and find operations
- Processes retrieved documents into structured response
- Calculates pagination metadata
- Returns combined category and pagination data

### Sample:

```
  query {
      getProjectCategories(
          limit: 10,
          page: 1
      ) {
          data {
              id
              name
              created_at
              updated_at
          }
          pagination {
              currentPage
              perPage
              count
              offset
              totalPages
              totalItems
          }
      }
  }
```

### Possible Errors:
- invalid or missing limit argument - Pagination parameter validation failed
- database connection error - MongoDB connection issues
- error counting documents - Total count operation failed
- error finding documents - Document retrieval failed
- error decoding documents - Document parsing failed
- context timeout - Operation exceeded time limit

### Returns:

```
{
    "data": {
        "getProjectCategories": {
            "data": [
                {
                    "id": 1,
                    "name": "Web Development",
                    "created_at": "16-04-2025 17:34:47",
                    "updated_at": "16-04-2025 17:34:47"
                }
            ],
            "pagination": {
                "currentPage": 1,
                "perPage": 10,
                "count": 25,
                "offset": 0,
                "totalPages": 3,
                "totalItems": 25
            }
        }
    }
}
```

### Security Considerations:
- Implements context timeout for operation control
- Uses connection pooling for database efficiency
- Returns sanitized error messages
- Validates input parameters
- Implements proper cursor handling
- Uses type-safe operations
- Manages memory efficiently

### Technical Details:
- Context Timeout: 30 seconds
- Database Collection: "project_categories"
- Default Page Size: 10 items
- Concurrent Operations: Count and Find
- Connection Pooling: Enabled
- Cursor Management: Automatic cleanup

### Performance Notes:
- Concurrent database operations
- Pre-allocated memory for results
- Efficient cursor handling with cursor.All()
- Optimized pagination calculations
- Minimal memory footprint
- Efficient error handling
- Connection reuse through pooling
- Context-based timeout management
- Batch document processing

### Implementation Notes:
- Uses goroutines for concurrent operations
- Implements wait groups for synchronization
- Provides default pagination values
- Handles edge cases for page numbers
- Calculates correct offset values
- Manages resource cleanup
- Supports scalable data retrieval