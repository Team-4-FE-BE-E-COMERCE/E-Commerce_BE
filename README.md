<h1><b><center> Alta E-Commerce </center></b> </h1>

#  📋 description
Alta E-commerce is a buying and selling application that has many features along with an explanation of the application.

1. `Users` can register a new account and login
2. `Users` can post profile photos
3. `Users` update profile and delete account
4. `Users` can create products, view their own products and other people's products
5. `Users` can add products to their own cart
6. `Users` can change their own cart contents
7. `Users` can place orders and view their order history

# 📍 Endpoints:
Read the API documentation here [API Endpoint Documentation](https://app.swaggerhub.com/apis-docs/HERIBUDIYANA/E-Commerce/1.0.0) (Swagger)

`Feature Users`
- POST /users (Register New User)
- POST /login (Login User)
- PUT /users (Update User)
- DELETE /users (Delete Account)
- GET /users (Get My Profile)
- GET /users/:id (Get Profile Other Users)

`Feature Products`
- POST /products (Add New Products)
- PUT /products/:id (Update Products)
- GET /products (Get All Products)
- DELETE /products/:id (Delete Products)

`Feature Carts`
- POST /carts (Add Product To Carts)
- PUT /carts/:id (Update Carts)
- GET /carts (Get All Carts)
- DELETE /carts (Delete Carts)

`Feature Orders`
- POST /orders (Checkout Product at Cart)
- PUT /orders/confirm (Confirm Order)
- PUT /orders/cancel (Cancel Order)
- GET /orders (Get History Order)

#  👨🏼‍💻 Contributors :
- [Heri By](https://github.com/BangHer99) (creators)
- [Mochammad Dany](https://github.com/mdanys) (creators)

#  📉 Database Schema:
![ERD](https://github.com/Team-4-FE-BE-E-COMERCE/E-Commerce_BE/blob/main/ERD/ERD%20E-Commerce.jpg?raw=true)


## Run Locally

```bash
git clone git@github.com:Team-4-FE-BE-E-COMERCE/E-Commerce_BE.git

source .env
go run main.go
