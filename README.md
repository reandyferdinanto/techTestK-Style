# techTestK-Style
This repository contains code for a Golang-based REST API with CRUD functionalities for storing and retrieving member data in a database.

## Installation
1. Install Golang on your local machine. For more information, visit the official Golang website.
2. Clone the repository using the following command:

```
git clone https://github.com/<your_username>/<your_repository>.git
```
3. Navigate to the repository's directory using the command line.
4. Install the required dependencies using the following command:
```
go mod tidy
```
5. Configure your database connection details in database.go file.
6. Start the server using the following command:
``` 
go run main.go
```

## Usage

The API exposes the following endpoints:

* POST /member - Create a new member
* GET /member - Retrieve all members
* PUT /member/:id - Update a member
* DELETE /member/:id - Delete a member

The request and response payloads are in JSON format.

## POSTMAN Documentation
https://documenter.getpostman.com/view/23401470/2s93Xu3mQ3

## Dependecies
Dependencies
Fiber - Web framework
GORM - ORM library for Golang
MySQL driver - MySQL driver for Golang

## License
This project is licensed under the MIT License. Feel free to use and modify this code as per your requirement.
