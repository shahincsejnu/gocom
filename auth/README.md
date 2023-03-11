# Auth Service

## Local URL (when run via `docker-compose`)
- `localhost:8084/`

## API Endpoints with examples
- `GET /` (Service's HealthCheck endpoint):
    - Request URL: `localhost:8084/`
    - Response:
        ```
        {
            "ok": true,
            "serviceName": "AuthService"
        }
        ```
- `POST /signup` (SignUp endpoint):
    - Request URL: `localhost:8084/signup`
    - Request Body:
        ```
        {
            "name": "demo",
            "email": "demo@gmail.com",
            "password": "demopass"
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Successfully Signed Up!"
        }
        ```
- `POST /login` (Login endpoint):
    - Request URL: `localhost:8084/login`
    - Request Body:
        ```
        {
            "name": "demo",
            "password": "demopass"
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Successfully Logged In!",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzY3NTUzNTksInVzZXJfaWQiOiI2NDBjZWJjYi1jZDI3LTRkZmYtODZjYi1jOTUxZTJkNjU4MjgifQ.hvw5wA0nn7ZjUWkMxwp4d8NTrKy3fQSLD7_kPvz_9Vk"
        }
        ```
- `POST /console` (JWT Token verification endpoint):
    - Request URL: `localhost:8084/console`
    - Request Body:
        ```
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzY3NTUzNTksInVzZXJfaWQiOiI2NDBjZWJjYi1jZDI3LTRkZmYtODZjYi1jOTUxZTJkNjU4MjgifQ.hvw5wA0nn7ZjUWkMxwp4d8NTrKy3fQSLD7_kPvz_9Vk"
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Successfully verified",
            "user": {
                "id": "640cebcb-cd27-4dff-86cb-c951e2d65828",
                "name": "demo",
                "email": "demo@gmail.com",
                "password": "$2a$10$lMRv0ZvYVCmhac4DkiLqQuItC1XEmqclPyEhLavVHaai9XJ0QcG6G",
                "isAdmin": false,
                "createdAt": "2023-02-18T20:21:40.404Z",
                "updatedAt": "2023-02-18T20:21:40.404Z"
            }
        }
        ```