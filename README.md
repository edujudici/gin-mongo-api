# gin-mongo-api

Build a REST API with Golang and MongoDB - Gin-gonic Version
This repository shows the source code for building a user management application with Golang using the Gin-gonic framework and MongoDB.

[Article Link](https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m)


## Prerequisites
The following steps in this post require Golang experience. Experience with MongoDB isn’t a requirement, but it’s nice to have.

We will also be needing the following:
A [MongoDB](https://www.mongodb.com/) account to host database. [Signup](https://www.mongodb.com/cloud/atlas/register) is completely free.
[Postman](https://www.postman.com/downloads/) or any API testing application of your choice

## Dependencies
github.com/gin-gonic/gin is a framework for building web applications.

go.mongodb.org/mongo-driver/mongo is a driver for connecting to MongoDB.

github.com/joho/godotenv is a library for managing environment variables.

github.com/go-playground/validator/v10 is a library for validating structs and fields.

## Run application
```
$ go run main.go
```

## Endpoints to Access on Postman

### Create New User
```
POST localhost:6000/users
{
    "name": "user test",
    "location": "Brazil",
    "title": "development"
}
```

### Update User
```
PUT localhost:6000/users/<id>
{
    "name": "user test",
    "location": "Brazil",
    "title": "development"
}
```

### Find User
```
GET localhost:6000/users/<id>
```

### Remove User
```
DELETE localhost:6000/users/<id>
```

### List All Users
```
GET localhost:6000/users
```