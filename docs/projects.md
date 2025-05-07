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