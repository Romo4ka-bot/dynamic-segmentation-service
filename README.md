# Dynamic Segmentation Service

Read this in other languages: [English](README.md), [Русский](README.ru.md)

This is a service designed for storing user profiles and the segments they belong to. 
It offers features such as creating, modifying, and deleting segments, as well as adding and removing users from segments, and creating new user profiles.

## Local launch of the service

1) The service utilizes Docker, so to launch it, simply execute the command below.

```bash
docker-compose up --build -d
```

Two containers will be raised (application and PostgreSQL database) as described in the docker-compose.yml file located at the project root. 
To confirm that the containers have been successfully brought up, you can use the command below and check the status (Up).

```bash
docker ps
```

You can stop the service using the following command:

```bash
docker-compose stop
```

## Main functionality:
1) Creating a segment
    * url: /api/v1/segments POST
    * incoming value:
   ```go
   package model
   type Segment struct {
	    Slug string
   }
   ```
    * return value: An identical model supplemented with some values
2) Getting generated pdf files
    * url: /api/v1/segments/{id} DELETE
    * incoming value: id of segment
    * return value: Upon a successful (200) response, null will be returned
3) Adding users to a segment and removing them from it
   * url: /internal/api/v1/template/pdf
   * incoming value:
   ```go
   package model
   type UserSegments struct {
    AddSlugs []string
    RemoveSlugs []string
    UserId int
   }
   ```
   * return value:
   ```go
   package model
   type User struct {
	   Id int
	   FirstName string
	   SecondName string
	   Username string
	   HashPassword string
   }
   ```
4) Getting the segments that the user is a member of
   * url: /api/v1/user-segments/users/{user_id} GET
   * incoming value: id of user(user_id)
   * return value: the User model, which is described in paragraph 3

More detailed description can be found in the Swagger documentation, where you can also test the requests. Swagger URL: http://localhost:8080/swagger/index.html.

Additionally, in the requests.http file located at the project root, examples of requests have been provided, which you can execute within a development environment of your choice or using any other method.

### List of libraries and technologies used for building the service
   * [gin](https://github.com/gin-gonic/gin) for describing API endpoints and handling requests
   * [net/http](https://pkg.go.dev/net/http) an embedded library for running the server
   * [gorm](https://github.com/go-gorm/gorm) used as an Object-Relational Mapping (ORM)
   * [swag](https://github.com/swaggo/swag) swagger for API documentation
   * [logrus](https://github.com/sirupsen/logrus) for logging errors in JSON format 
   * PostgreSQL database
   * Docker, docker-compose
