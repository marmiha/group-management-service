# group-management-service

Group management REST API service written in Golang. Service provides a way to list, add, modify and remove users and
groups where each of the users can belong to at most one group.

### Instructions for testing and running

You have to install `docker-compose` and have set the `environment variables` either manually or by creating a `.env`
file in the root directory.

Set the following environment variables (each have their default values). The ones that are required are marked
with __*__:

- __*__ `ExposedPort`=`3000` -> At which port the API will be accessible
- __*__ `LogLevel`=`info` -> What is the microservice logging level (`trace`, `debug`, `info`, `error`, `fatal`
  , `panic`)
- __*__ `DataService.User.Impl`=`postgresds`-> Which implementation of `UserDataServiceInterface` will be used. By this
  we sustain the pluggable implementations in check with environment variables. Based on the set variable, we could for
  this implementation use something different from `postgresds` which denotes our `UserdataServiceInterface` implemented
  by underlying Postgres database. At this moment, only Postgres implemented although has the capability to include
  other, for instance an in memory implementation. or a distant microservice storage for our user data.
- __*__ `DataService.Group.Impl`=`postgresds` -> Which implementation of `GroupDataServiceInterface` will be used.
  Explanation follows the same notion as the environment variable for UserDataInterface (mentioned
  under `DataService.User.Impl`
  environment variable).
- `DataService.Postgres.User`=`defaultuser`  -> Postgres user that will be used by group-management service for database
  operations.
- `DataService.Postgres.Password`=`defaultpassword` -> Postgres user password.
- `DataService.Postgres.ConnPoolSize`=`20` -> Postgres connection pool size.
- `DataService.Postgres.DropTableOnConn`=`true` -> Drop existing tables/entries upon postgres database connection (fresh
  start 💆‍♂).
- __*__ `- Adapter.Impl=rest`=`rest` -> Choose the Adapter implementation. Currently, only REST is supported, but the
  codebase supports ease of adding additional endpoints (GraphQL, RPC... 💆‍♂)
- `Adapter.Rest.JwtKey`=`defaultjwtkey` -> This key is used to sign JWT authentication tokens. It's set because the
  default adapter (
  REST) implementation needs the JwtKey for signing the Bearer tokens in part of authentication.
- `Adapter.Rest.EnableLogging`=`true` -> If true then logging for HTTP requests is enabled.
- `POSTGRES_USER` -> Should equal `DataService.Postgres.User`
- `POSTGRES_˙PASSWORD` -> Should equal `DataService.Postgres.Password`

__Run__ the following commands for running the app:

- `docker-compose up --build --abort-on-container-exit`

__Tests__ are based on use cases and are executed on adapter level, therefore each adapter has its own tests. These are
discriminated between by the `Adapter.Impl`. At this moment only `rest` is supported as there is no other adapter
implementation, but the codebase supports ease of addition more adapters and therefore tests for each of these.

For each configuration of `docker-compose.yml` with the mentioned environment variables you can test it by setting them
the same in `docker-compose.test.yml` and running:

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

### Project Trivia

#### Inspiration

I wanted to build a Domain Driven Design microservice so everything is decoupled as much as possible with the use of
interfaces. Project structure is inspired by [this project](https://github.com/jfeng45/servicetmpl1) where Jin Feng
describes a nice project layout with maximum decoupling of three layers: business logic, data store logic and adapter
aka. application endpoints.

#### Structure

The most top level [adapter](./group-management-api/adapter) package. This is where the common interfaces and structs
for Adapter implementations are defined. The underlying folders represent each implementation of the interfaces (
currently only [restapi](./group-management-api/adapter/restapi) is implemented.

The underlying package is the [domain](./group-management-api/domain) which includes
the [usecase](./group-management-api/domain/usecase), [model](./group-management-api/domain/model)
and [payload](payload) packages. The `usecase` defines common interfaces which are used by overlying `adapter` package
to access the business logic. The sub folders\packages represent the implementations of these `usecase interfaces`. The
package `model` includes all the models that the business logic uses (`Group` and `User`), while the `payload`
encapsulates the general data for specific `usecase interfaces functions`.

Business logic relies on the data or should we say a data service. The
package [dataservice](./group-management-api/service/dataservice)
defines two interfaces that the `usecase interface implementations` uses: `GropDataInterface` and `UserdataInterface`.
This way the `domain` does not need to know the specifics of the implementations itself, but rather just get them
injected in runtime. This way group data might represent a remote microservice containing our data about groups, and a
local one for our user data. The combinations are endless, but these have to be implemented to suffice the data service
interface signatures. At this moment [postgresds](./group-management-api/service/dataservice/postgresds) complies with
both of the required signatures and is therefore used for both implementations of `GropDataInterface`
and `UserdataInterface`.

Application supports a dynamic startup from environment variables. A simple entry point is made
in [app](./group-management-api/app), with a call to initialize the application aka the group management microservice.
Firstly the environment variables are read into [config struct](./group-management-api/app/config/config.go) and then
saved inside a [container](./group-management-api/app/container/container.go). Secondly, the logger level is initialized
so that we can use it while the application gets built from top to bottom. Thirdly, the factory injection process
begins. Adapter being the top most layer so this is our entrypoint. The chosen Adapter
implementation ([ref](./group-management-api/app/factory/adapter.go#L21)) is read from the configuration. This
implementation factory method then calls the appropriate factory method for
the [usecases](/group-management-api/app/factory/usecase.go)) which the adapter implementation needs. The usecase
factory methods then call
[dataservice](/group-management-api/app/factory/dataservice.go) factory methods for their dependencies, and so forth. These factory methods,
follow a singleton pattern and thus each implementation of any interface is created and injected only once when it's first requested.
Whilst this process the factory methods can save functions inside the [container](./group-management-api/app/container/container.go)
called `ShutDownables`, which are called before the application exits. For example a Postgres DB needs its connections
to be closed before exiting, therefore the factory method for the postgres data service implementation registers a close 
connection shutdownable. Like the mentioned, the application also has to have an entry point which is saved inside the container.
It's a basic function returning an error. This is set while initializing the app. For example the Rest Adapter sets the function to
start listening and serving on some port. This is then called whenever after the container is successfully initialized.

The following picture also visualizes the concept of interfaces and their implementations.
![structure_graph](https://i.imgur.com/1hxYqg8.jpeg)