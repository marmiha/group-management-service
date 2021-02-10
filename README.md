# group-management-service

Group management REST API service written in Golang. Service provides a way to list, add, modify and remove users and
groups where each of the users can belong to at most one group.

### Instructions for testing and running

You have to install `docker-compose` and have set the `environment variables` either manually or by creating a `.env`
file in the root directory.

Set the following environment variables (each have their default values). The ones that are required are marked with __*__:

- __*__ `ExposedPort`=`3000` -> At which port the API will be accessible
- __*__ `LogLevel`=`info` -> What is the microservice logging level (`trace`, `debug`, `info`, `error`, `fatal`, `panic`)
- __*__ `DataService.User.Impl`=`postgresds`-> Which implementation of `UserDataServiceInterface` will be used. By this we
  sustain the pluggable implementations in check with environment variables. Based on the set variable, we could for
  this implementation use something different from `postgresds` which denotes our `UserdataServiceInterface` implemented
  by underlying Postgres database. At this moment, only Postgres implemented although has the capability to include
  other, for instance an in memory implementation. or a distant microservice storage for our user data.
- __*__ `DataService.Group.Impl`=`postgresds` -> Which implementation of `GroupDataServiceInterface` will be used. Explanation
  follows the same notion as the environment variable for UserDataInterface (mentioned under `DataService.User.Impl`
  environment variable).
- `DataService.Postgres.User`=`defaultuser`  -> Postgres user that will be used by group-management service for database
  operations.
- `DataService.Postgres.Password`=`defaultpassword` -> Postgres user password.
- `DataService.Postgres.ConnPoolSize`=`20` -> Postgres connection pool size.
- `DataService.Postgres.DropTableOnConn`=`true` -> Drop existing tables/entries upon postgres database connection (fresh
  start 💆‍♂).
- __*__ `- Adapter.Impl=rest`=`rest` -> Choose the Adapter implementation. Currently, only REST is supported, but the codebase
  supports ease of adding additional endpoints (GraphQL, RPC... 💆‍♂)
- `Adapter.Rest.JwtKey`=`defaultjwtkey` -> This key is used to sign JWT authentication tokens. It's set because the default adapter (
  REST) implementation needs the JwtKey for signing the Bearer tokens in part of authentication.
- `Adapter.Rest.EnableLogging`=`true` -> If true then logging for HTTP requests is enabled.
- `POSTGRES_USER` -> Should equal `DataService.Postgres.User`
- `POSTGRES_˙PASSWORD` -> Should equal `DataService.Postgres.Password`

__Run__ the following commands for running the app:

1) `docker-compose build`
2) `docker-compose up`

__Tests__ are based on use cases and are executed on adapter level, therefore
each adapter has its own tests. These are discriminated between by the `Adapter.Impl`. At this moment only `rest` is supported
as there is no other adapter implementation, but the codebase supports ease of addition more adapters and therefore tests for each of these.

For each configuration of `docker-compose.yml` with the mentioned environment variables you can test it by setting them the
same in `docker-compose.test.yml` and running:

 * `docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit`

By default `docker-compose.test.yml` tests the default configuration of `docker-compose.yml`.

### REST OpenAPI specification

After start up the application REST API endpoints will be accessible on `host:<ExposedPort>/api/v1/...` while the
OpenAPI specification of these are described in `swagger.yml` accessible at `host:<ExposedPort>/specs/swagger.yml`. The
hosted file can be pasted inside [swagger.io](https://editor.swagger.io/). Default host is `localhost:3000` so it has to
be changed manually if the microservice is accessible elsewhere.

The Swagger files are created with `go-swagger` framework from annotations(comments), which are mostly
in [restapi package](https://github.com/marmiha/group-management-service/tree/master/group-management-api/adapter/restapi)
. A nice example are the annotated `swagger-routes`
in [endpoints file](https://github.com/marmiha/group-management-service/blob/master/group-management-api/adapter/restapi/endpoints.go#L16)
, where for each route a definition, responses and security is defined.

If `go-swagger` is installed recompilation of `swagger.yml` file can be done via this command:

- `swagger generate spec -m -w group-management-api\adapter\restapi -o group-management-api\public\specs\swagger.yml`

This command can be used in case any endpoints get changed and the OpenAPI specification has to be recompiled.
