-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: CreateProduct :exec
INSERT INTO products (
    id, name, price, description, stock
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
    set price = $2,
    description = $3,
    stock = $4
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: CreateOrder :exec
INSERT INTO orders (
    id, product_id, user_id, address_id, quantity
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: GetOrdersByUser :many
SELECT * FROM orders
WHERE user_id = $1;

-- name: GetOrderById :one
SELECT * FROM orders
WHERE id = $1;

-- name: UpdateOrder :exec
UPDATE orders
    set quantity = $2
WHERE id = $1;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;

-- name: CreateAddress :exec
INSERT INTO addresses (
    id, user_id, country, city, street_address
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: GetAddressesByUser :many
SELECT * FROM addresses
WHERE user_id = $1;

-- name: GetAddressById :one
SELECT * FROM addresses
WHERE id = $1;

-- name: UpdateAddress :exec
UPDATE addresses
    set country = $2,
    city = $3,
    street_address = $4
WHERE id = $1;

-- name: DeleteAddress :exec
DELETE FROM addresses
WHERE id = $1;
