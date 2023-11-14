# MARVEL API-RESTFUL

This is a simple Go project that demonstrates how to create a RESTful API for performing CRUD (Create, Read, Update, Delete) operations on a PostgreSQL database regarding the MCU (Marvel Cinematic Universe).


## Prerequisites

Before you begin, make sure you have the following software installed on your system:
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install)
- [Task](https://taskfile.dev/install)

---

## Usage of `task` command

This project contains a file `Taskfile.yml` which contains a group os tasks you can execute with `task` command:
- `task setup` -> Installs all the required tools to deploy and run the project
- `task all:restart` -> Restarts the PostgreSQL container and recreates the database schema and tables
- `task service:up` -> Starts the PostgreSQL database using docker
- `task service:down` -> Stops the PostgreSQL database and removes its container
- `task db:create` -> Creates the database schema
- `task db:drop` -> Removes the database schema
- `task db:init` -> Creates the database tables
- `task db:clean` -> Removes the database tables
