# Devices API

For every reader please be aware that before this contact I didn't any knowledge about the language go.
This repository is a implementation of an API written in go with the objective of executing the following operations:

 - Add device
 - Get device by id
 - Update a device
 - Get all devices
 - Delete a device
 - Get all devices froma brand

This api is connected to a postgres database.

### Steps

Before running the solution please execute

    docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres

After just run

    go run main.go


### Future improvements

 - Create swagger 
 - Create a docker-compose
