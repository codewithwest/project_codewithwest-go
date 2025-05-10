# Projects Documentation

## Table of Contents

- [Overview](#overview)
- [Mutations](#mutations)
    - [Create Project](#create-project)
- [Queries](#queries)
    - [Get Projects](#get-projects)

## Overview

This documentation cover all projects specific implementation done in the project.

### Mutations

### Create Project

### Description:

    Creates a new project in the system with validation of project category existence.
    Implements proper error handling, input validation, and database operations.
    Returns the created project data upon successful creation.

### Input Parameters:

```
input: {
    name: String! (required) - Project name
    project_category_id: Int! (required) - ID of associated project category
    description: String! (required) - Project description
    tech_stacks: [String!]! (required) - Array of technology stack names
    github_link: String (optional) - GitHub repository link
    live_link: String (optional) - Live project URL
    test_link: String (optional) - Test environment URL
}
```

### Process

- Validates input parameters
- Verifies project category existence 
- Generates new project ID
- Creates project record in database
- Returns created project data

### Samples

```
mutation {
    createProject(input: {
        name: "E-Commerce Platform",
        project_category_id: 1,
        description: "Full-stack e-commerce solution",
        tech_stacks: ["Go", "React", "MongoDB"],
        github_link: "https://github.com/user/project",
        live_link: "https://project.com",
        test_link: "https://test.project.com"
    }) {
        id
        name
        project_category_id
        description
        tech_stacks
        github_link
        live_link
        test_link
        created_at
        updated_at
    }
}
```

### Possible Errors:

- invalid input parameters - Input validation failed
- project category not found - Invalid project_category_id
- database connection error - MongoDB connection issues
- ID generation error - Failed to generate project ID
- insertion error - Project creation failed
- context timeout - Operation exceeded time limit

### Returns:

```
{
    "data": {
        "createProject": {
            "id": 1,
            "name": "E-Commerce Platform",
            "project_category_id": 1,
            "description": "Full-stack e-commerce solution",
            "tech_stacks": ["Go", "React", "MongoDB"],
            "github_link": "https://github.com/user/project",
            "live_link": "https://project.com",
            "test_link": "https://test.project.com",
            "created_at": "16-04-2025 17:34:47",
            "updated_at": "16-04-2025 17:34:47"
        }
    }
}
```

### Security Considerations:

- Implements context timeout for operation control
- Validates all input parameters
- Uses type-safe operations
- Returns sanitized error messages
- Implements proper error handling
- Manages database connections efficiently

### Technical Details:

- Context Timeout: 60 seconds
- Database Collections: "projects", "project_categories"
- Input Validation: Required fields checking
- Error Handling: Comprehensive error wrapping
- ID Generation: Auto-incrementing based on collection
- Database Operations: FindOne, InsertOne

### Performance Notes:

- Sequential database operations
- Context-based timeout management
- Efficient error handling
- Connection reuse through pooling
- Type-safe operations
- Proper resource cleanup

### Implementation Notes:

- Uses MongoDB for data storage
- Implements input validation
- Handles edge cases for project category existence
- Manages resource cleanup through defer
- Provides detailed error messages
- Supports optional fields

## Queries

### Get Projects

### Description:
  Retrieves a paginated list of projects from the system with efficient concurrent operations.
  Implements proper error handling, pagination, and optimized database queries.
  Returns projects data with pagination metadata.

### Input Parameters:

```
input: {
    limit: Int! (required) - Number of items per page (default: 10)
    page: Int (optional) - Page number to retrieve (default: 1)
}
```

### Process

- Validates pagination parameters
- Concurrently executes count and find operations
- Calculates pagination metadata
- Retrieves specified page of projects
- Returns paginated project data

### Sample Query:

```
query {
    getProjects(limit: 10, page: 1) {
        data {
            id
            name
            project_category_id
            description
            tech_stacks
            github_link
            live_link
            test_link
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

- invalid limit parameter - Input validation failed
- database connection error - MongoDB connection issues
- count operation error - Failed to count total documents
- find operation error - Failed to retrieve documents
- cursor error - Failed to process cursor
- context timeout - Operation exceeded time limit
- decoding error - Failed to decode document data

### Returns

```
{
    "data": {
        "getProjects": {
            "data": [
                {
                    "id": 1,
                    "name": "E-Commerce Platform",
                    "project_category_id": 1,
                    "description": "Full-stack e-commerce solution",
                    "tech_stacks": ["Go", "React", "MongoDB"],
                    "github_link": "https://github.com/user/project",
                    "live_link": "https://project.com",
                    "test_link": "https://test.project.com",
                    "created_at": "16-04-2025 17:34:47",
                    "updated_at": "16-04-2025 17:34:47"
                }
                // ... more projects
            ],
            "pagination": {
                "currentPage": 1,
                "perPage": 10,
                "count": 50,
                "offset": 0,
                "totalPages": 5,
                "totalItems": 50
            }
        }
    }
}

```

### Security Considerations:

- Implements context timeout for operation control
- Uses concurrent operations safely with WaitGroups
- Implements proper cursor cleanup
- Returns sanitized error messages
- Manages database connections efficiently
- Uses type-safe operations

### Technical Details:

- Context Timeout: 30 seconds
- Database Collection: "projects"
- Pagination: Skip-limit based
- Concurrent Operations: Count and Find
- Error Handling: Comprehensive error wrapping

### Resource Management: 

- Proper defer cleanup
- Performance Notes:
- Concurrent database operations
- Pre-allocated slices for results
- Efficient cursor.All() usage
- Connection reuse through pooling
- Context-based timeout management
- Proper resource cleanup
- Optimized pagination calculations

### Implementation Notes:
- Uses MongoDB for data storage
- Implements efficient pagination
- Uses goroutines for concurrent operations
- Manages resource cleanup through defer
- Provides detailed error messages
- Supports default pagination values
- Handles edge cases for page numbers
- Implements proper cursor management
- Returns structured pagination metadata

### Query Optimization:

- Concurrent execution of count and find operations
- Efficient use of MongoDB cursor
- Proper index usage for pagination
- Optimized batch document retrieval
- Memory-efficient slice allocation

