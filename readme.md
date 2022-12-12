# Golang JWT

A simple CLI Gin api that uses jwt token authentication.

Basically my quick and dirty golang project

## Technologies

### Stack
Project is created with: 
* golang
* sqlite

### Packages
Project uses the following packages: 
* gorm
* go-jwt
* gin

## Launch
All commands should be in the project directory

### Setup
Follow these steps to build the project for your system: 
     
    - cd to the `httpd` folder
    - get dependencies with `go mod download`
    - run build command here to get exec file

### Usage
Use any client of your choice to access these endpoints.
- GET "user/"
  - simple get request: returns all registered users
- POST "user/sign-up"
  - requires payload `json:"name", "email", "password"`: adds a user to the backend
- POST "user/login"
  - requires payload `json:"email", "password"`: returns and sets a jwt token as session
- POST "user/log-out"
  - requires payload `json:"email", "password"`: removes session
- GET "secret/"
  - simple get request: returns a message for successful authentication