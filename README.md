# golang-gin-app
Golang App with GIN framework

This app contains the relationship between User and its Posts(User HAS MANY Posts)

There are CRUD operations with User and Post Model and the Authorization details help to associate a Post with the User.

# Features

- Serializers are used in different endpoint to send different response as per the requirement.

- Postgres DB is introduced with required migrations.

- Validators are used to validate Struct data

- Channels are used to handle some special cases.

# To run the app
- go run main.go\
OR
- To auto reload server when changes are made in your code(use nodemon)\
nodemon --exec go run main.go --signal SIGTERM

# Few User Endpoints
1. POST - localhost:9000/signup
2. GET  - localhost:9000/users
3. PUT  - localhost:9000/users/:id
4. GET  - localhost:9000/users/:id

# Few Post Endpoints
1. POST - localhost:9000/post
2. GET  - localhost:9000/posts
