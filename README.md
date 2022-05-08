# Database query analyzer app


Application returns information about queries executed in the database.

To run application locally you can use several approaches:

## Docker compose

To start application you can start it using next command:

    docker compose up -d   


Configuration will be taken from the `.env.docker` file in this case.

In this case database will be created automatically in the docker container.

## Using Makefile

Before you can run application you need to run PostgreSQL locally.

Then in the `.env.local` specify configuration using environment variables:

| Environment variable    | Default   | Description                                     |
| ----------------------- |-----------|-------------------------------------------------|
| HTTPSERVER_HOSTNAME         | localhost | application hostname                            |
| HTTPSERVER_PORT     | 8080      | application port                                |
| DATABASE_HOSTNAME      |           | database hostname                               |
| DATABASE_PORT            |           | database port                                   |
| DATABASE_NAME                    |       | database name which will be used by application |
| DATABASE_USERNAME           | .         | credential to connect to database, username     |
| DATABASE_PASSWORD                 |       | credentials to connect to database, password    |

After that you finally can run application by running 

    make run

command.

At first it will run migration scripts in the `./migrations` directory.
After that it will start application itself.

## Testing application

To test application you can use http commands which are defined in the `./http` folder.

You can run them using VS Code REST plugin or using Intellij IDEA.

OR you can run them using curl for this, for example:

    curl -X GET --location "http://localhost:8080/health"

should return `OK` string in the response. It means that app is running.

    curl -X GET --location "http://localhost:8080/v1/query?page=1&size=10&order_by=asc"

returns the list of queries

More information about the endpoint you can find in the `/api/swagger.yml` file.
