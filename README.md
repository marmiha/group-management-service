# group-management-service
Group management REST API service written in Golang. Service provides a way to list, add, modify and remove users and groups where each of the users can belong to at most one group.

### Instructions

You have to install `docker-compose` and have set the `environment variables` either manually or by creating a `.env` file
in the root directory.

Set the following environment variables (each have their default values):

- `ExposedPort`=`3000` -> at which port the API will be accessible
- `DataService.User.Impl`=`postgresds` -> Which implementation of `UserDataServiceInterface` will be used. By this we sustain 
  the pluggable implementations in check with environment variables. Based on the set variable, we could for this implementation
  use something different from `postgresds` which denotes our `UserdataServiceInterface` implemented by underlying Postgres database.
  At this moment, only Postgres implemented although has the capability to include other, for instance an in memory implementation.
  or a distant microservice storage for our user data.
- `DataService.Group.Impl`=`postgresds` -> Which implementation of `GroupDataServiceInterface` will be used. Explanation follows the 
  same notion as the environment variable for UserDataInterface (mentioned under `DataService.User.Impl` environment variable).
- `DataService.Postgres.User`=`defaultuser`  -> Postgres user that will be used by group-management service for database operations.
- `DataService.Postgres.Password`=`defaultpassword` -> Postgres user password.
- `DataService.Postgres.ConnPoolSize`=`20` -> Postgres connection pool size.
- `DataService.Postgres.DropTableOnConn`=`true` -> Drop existing tables/entries upon postgres database connection (fresh start 💆‍♂).
- `- Adapter.Impl=rest`=`rest` -> Choose the Adapter implementation. Currently, only REST is supported, but the codebase supports ease
  of adding additional endpoints (GraphQL, RPC... 💆‍♂)
- `Adapter.Rest.JwtKey` -> This key is used to sign JWT authentication tokens. It's set because the default adapter (REST) implementation
needs the JwtKey for signing the Bearer tokens in part of authentication.

Run the following commands:
1) `docker-compose build`
2) `docker-compose up`