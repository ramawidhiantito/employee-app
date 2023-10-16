
# Employee APP

This is a simple app that manage Employee data. this app is coded in golang using Gin framework, postgres, and runs on docker


## Stacks Used

- Gin (Golang framework)
- Docker
- PostgreSQL




## Installation

This project runs on docker, if you doesnt have docker yet you can refer to this [Docker Installation](https://docs.docker.com/engine/install/).

I already created the `Dockerfile` and `docker-compose.yaml` in this project.
so after you clone my project. you should open terminal on project directory, and then you can run this command

```bash
  sudo docker-compose up --build
  
```
After succesful build you should have `PostgreSQL` and the `API` running. And you're good to go!

*Note : i already configured the `PostgreSQL` to run first and then comes the `API` in the `docker-compose.yaml` file


## Features
This employee APP have CRUD features
- Create Employee 
- Get Employee By ID
- Get All Employee
- Update Employee By ID
- Delete Employee

Dont worry! i already put `Auto Migration` to create the `Employee` table, so you can directly test these endpoints! 

For more detail about the endpoints, after you run the project, you can visit this link http://localhost:8080/swagger/index.html  or you can copy the endpoints URL and the payload.

I also provided postman collection if you want to test the employee APP. You can import `Employee APP.postman_collection.json` located in root project directory to you Postman!

I also provide unit tests on the `controller` folder