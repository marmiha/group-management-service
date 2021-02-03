# group-management-service
Group management REST API service written in Golang. Service provides a way to list, add, modify and remove users and groups where each of the users can belong to at most one group.

### Instructions

You have to install `docker-compose` and have set the `environment variables` either manually or by creating a `.env` file
in the root directory.

Set the following environment variables (each have their defaut values):


- `API_PORT`=`3000` -> at which port the API will be accessible
- `POSTGRES_USER`=`defaultuser`  -> Postgres user that will access the from the group-management service.
- `POSTGRES_PASSWORD`=`defaultpassword` -> Postgres user password.
- `JWT_KEY`=`defaultjwtkey` -> This key is used to sign JWT authentication tokens.

Run the following commands:
1) `docker-compose build`
2) `docker-compose up`