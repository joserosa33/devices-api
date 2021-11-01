# Devices API

This repository resutls from a challenge made to develop a API to manage devices written in go, a language that previously I didn't had any contact. 
So the result is a API that supports the following operations:

 - Add device
 - Get device by id
 - Update a device
 - Get all devices
 - Delete a device
 - Get all devices froma brand

This api is connected to a postgres database.

### Steps
#### Using docker-compose
Being on the project root just run

    docker-compose up

#### Manually
Before running the solution please execute

    docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres

After just run

    go run main.go

### Future improvements

 - Create swagger 