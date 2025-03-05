# Codewithwest-go server

### This is the backend application that bridges the [Admin](https://github.com/codewithwest/project_codewithwest-go) and database interaction and feeds data also the [FrontEnd](https://codewithwest.vercel.app/)

## Functionality Provided

### Create Admin User
```
    mutation createAdminUser($input: AdminUserInput!) {
        createAdminUser(input: $input) {
          id
          username
          email
          password
          role
          type
          created_at
          updated_at
          last_login
      }
    }
```
### Get Admin Users
```
    query getAdminUsers($limit: Int!) { # $filter is the parameter
       getAdminUsers(limit: $limit) {
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
### Request Admin User Access
``` 
mutation adminUserAccessRequest($email: String!){
      adminUserAccessRequest(email: $email) {
        created_at
        email
        id
      }
    }
```
### Get Admin user requests
```
      query getAdminUserAccessRequests($limit: Int!) {
        getAdminUserAccessRequests(limit: $limit) {
          created_at
          email
          id
        }
      }
```
### Create project Category
```
 mutation createProjectCategory($name: String!) {
      createProjectCategory(name: $name) {
        id
        name
        created_at
        updated_at
      }
    }
```
### Get project categories
```
    query getProjects($limit: Int!) {
        getProjects(limit: $limit) {
          created_at
          description
          github_link
          id
          live_link
          name
          project_category_id
          tech_stacks
          test_link
          updated_at
        }
      }
```
### Create Project
```
    mutation createProject($input: ProjectInput!) {
          createProject(input: $input) {
            id
            project_category_id
            name
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
### Get Projects
```
    query getProjects($limit: Int!) {
        getProjects(limit: $limit) {
          created_at
          description
          github_link
          id
          live_link
          name
          project_category_id
          tech_stacks
          test_link
          updated_at
        }
      }
```
### Login Admin User
```
    query loginAdminUser($input: AdminUserInput!) {
        loginAdminUser(input: $input) {
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




