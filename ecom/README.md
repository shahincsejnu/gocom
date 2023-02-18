# Ecom service

## Local URL (when run via `docker-compose`)
- `localhost:8085/`

## API endpoints with examples
- `GET /` (Service's HealthCheck endpoint):
    - Request URL: `localhost:8085/`
    - Response:
        ```
        {
            "ok": true,
            "serviceName": "EcomService"
        }
        ```
- `GET /users/:userID` (Get a user by userID):
    - Request URL: `localhost:8085/users/640cebcb-cd27-4dff-86cb-c951e2d65828`
    - Response:
        ```
        {
            "ok": true,
            "message": "User is fetched",
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
- `GET /users/:userID/orders` (Get all the orders of a user):
    - Request URL: `localhost:8085/users/640cebcb-cd27-4dff-86cb-c951e2d65828/orders`
    - Response:
        ```
        {
            "ok": true,
            "message": "User Orders fetching is succeeded",
            "orders": null
        }
        ```
- `GET /users/:userID/addresses` (Get all the addresses of a user):
    - Request URL: `localhost:8085/users/640cebcb-cd27-4dff-86cb-c951e2d65828/addresses`
    - Response:
        ```
        {
            "ok": true,
            "message": "User Addresses fetching is succeeded",
            "addresses": null
        }
        ```

- `GET /products` (Get all the products):
    - Request URL: `localhost:8085/products`
    - Response:
        ```
        {
            "ok": true,
            "message": "Products fetching is succeeded",
            "products": [
                {
                    "id": "dcab11a7-7c33-4cc8-b177-7e3d00f7af0a",
                    "name": "demo2",
                    "price": 11,
                    "description": "demo product2",
                    "stock": 3,
                    "createdAt": "2023-02-14T07:52:45.809Z",
                    "updatedAt": "2023-02-14T07:52:45.809Z"
                }
            ]
        }
        ```
- `POST /products` (Create a product):
    - Request URL: `localhost:8085/products`
    - Request Body:
        ```
        {
            "name": "sample",
            "price": 100,
            "description": "sample product",
            "stock": 1
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Product is created successfully",
            "productID": "9df4ae40-9cf7-4acd-8d2b-df8be230eaf7"
        }
        ```
- `PATCH /products/:productID` (Update a product):
    - Request URL: `localhost:8085/products/9df4ae40-9cf7-4acd-8d2b-df8be230eaf7`
    - Request Body:
        ```
        {
            "price": 10,
            "description": "sample product",
            "stock": 2
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Product is updated successfully"
        }
        ```
- `GET /products/:productID` (Get details of a product):
    - Request URL: `localhost:8085/products/9df4ae40-9cf7-4acd-8d2b-df8be230eaf7`
    - Response:
        ```
        {
            "ok": true,
            "message": "Product fetching is succeeded",
            "product": {
                "id": "9df4ae40-9cf7-4acd-8d2b-df8be230eaf7",
                "name": "sample",
                "price": 10,
                "description": "sample product",
                "stock": 2,
                "createdAt": "2023-02-18T20:57:24.064Z",
                "updatedAt": "2023-02-18T20:57:24.064Z"
            }
        }
        ```
- `DELETE /products/:productID` (Delete a product):
    - Request URL: `localhost:8085/products/9df4ae40-9cf7-4acd-8d2b-df8be230eaf7`
    - Response:
        ```
        {
            "ok": true,
            "message": "Product deletion is succeeded"
        }
        ```

- `POST /orders` (Create an order):
    - Request URL: `localhost:8085/orders`
    - Request Body:
        ```
        {
            "productID": "dcab11a7-7c33-4cc8-b177-7e3d00f7af0a",
            "userID": "640cebcb-cd27-4dff-86cb-c951e2d65828",
            "addressID": "f599d514-c341-432d-95e4-2587b81f733d",
            "quantity": 1
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Order is created successfully",
            "orderID": "4e9a012b-0a80-4ada-8dbd-c87f2045c9aa"
        }
        ```
- `GET /orders/:orderID` (Get an order):
    - Request URL: `localhost:8085/orders/4e9a012b-0a80-4ada-8dbd-c87f2045c9aa`
    - Response:
        ```
        {
            "ok": true,
            "message": "Order fetching is succeeded",
            "order": {
                "id": "4e9a012b-0a80-4ada-8dbd-c87f2045c9aa",
                "productID": "dcab11a7-7c33-4cc8-b177-7e3d00f7af0a",
                "userID": "640cebcb-cd27-4dff-86cb-c951e2d65828",
                "addressID": "f599d514-c341-432d-95e4-2587b81f733d",
                "quantity": 1,
                "createdAt": "2023-02-18T21:17:00.752Z",
                "updatedAt": "2023-02-18T21:17:00.752Z"
            }
        }
        ```
- `PATCH /orders/:orderID` (Update an order):
    - Request URL: `localhost:8085/orders/4e9a012b-0a80-4ada-8dbd-c87f2045c9aa`
    - Request Body:
        ```
        {
            "quantity": 2
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Order is updated successfully"
        }
        ```
- `DELETE /orders/:orderID` (Delete an order):
    - Request URL: `localhost:8085/orders/4e9a012b-0a80-4ada-8dbd-c87f2045c9aa`
    - Response:
        ```
        {
            "ok": true,
            "message": "Order deletion is succeeded"
        }
        ```

- `POST /addresses` (Create an address):
    - Request URL: `localhost:8085/addresses`
    - Request Body:
        ```
        {
            "userID": "640cebcb-cd27-4dff-86cb-c951e2d65828",
            "country": "demo",
            "city": "demo",
            "streetAddress": "demo"
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Address is created successfully",
            "addressID": "f599d514-c341-432d-95e4-2587b81f733d"
        }
        ```
- `GET /addresses/:addressID` (Get an address):
    - Request URL: `localhost:8085/addresses/f599d514-c341-432d-95e4-2587b81f733d`
    - Response:
        ```
        {
            "ok": true,
            "message": "Address fetching is succeeded",
            "address": {
                "id": "f599d514-c341-432d-95e4-2587b81f733d",
                "userID": "640cebcb-cd27-4dff-86cb-c951e2d65828",
                "country": "demo",
                "city": "demo",
                "streetAddress": "demo",
                "createdAt": "2023-02-18T21:10:27.298Z",
                "updatedAt": "2023-02-18T21:10:27.298Z"
            }
        }
        ```
- `PATCH /addresses/:addressID` (Update an address):
    - Request URL: `localhost:8085/addresses/f599d514-c341-432d-95e4-2587b81f733d`
    - Request Body:
        ```
        {
            "country": "demo country",
            "city": "demo city",
            "streetAddress": "demo street"
        }
        ```
    - Response:
        ```
        {
            "ok": true,
            "message": "Address is updated successfully"
        }
        ```
- `DELETE /addresses/:addressID` (Delete an address):
    - Request URL: `localhost:8085/addresses/f599d514-c341-432d-95e4-2587b81f733d`
    - Response:
        ```
        {
            "ok": true,
            "message": "Address deletion is succeeded"
        }
        ```
