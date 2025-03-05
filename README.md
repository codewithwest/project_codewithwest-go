# codewithwest-go
codewithwest-go

### This is the backend application that bridges the [Admin](https://github.com/codewithwest/project_codewithwest-go) and database interaction and feeds data also the [FrontEnd](https://codewithwest.vercel.app/)

## Functionality Provided


### Create Admin User
![image](https://github.com/user-attachments/assets/33b080e5-ca0c-406d-ad58-04dcf9e28491)
### Get Admin Users
![image](https://github.com/user-attachments/assets/568e062f-df54-4452-8e14-87ef28f76862)
### Request Admin User Access
![image](https://github.com/user-attachments/assets/83ea6242-6f4f-416a-9b6c-06fb9e1c0ab8)

### Get Admin user requests
![image](https://github.com/user-attachments/assets/a712da0d-903e-4733-8318-3fd2bf1f9dd9)

### Create project Category
![image](https://github.com/user-attachments/assets/47110f49-b0bb-49b9-8500-3391a70589c6)
### Get project categories
![image](https://github.com/user-attachments/assets/8b38a961-6364-4103-9165-41d249bb035b)

### Create Project
![image](https://github.com/user-attachments/assets/d54f9f05-536a-42e1-b00f-47d28914b551)
### Get projects
![image](https://github.com/user-attachments/assets/393b55ec-ad37-4214-8feb-cbcf6288fc01)
### LOgin Admin User
```
{
  loginAdminUser(
    input: { password: "string", username: "string", email: "string" }
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




