# Atlas DNS

Atlas Corp Drone Navigation System (DNS) [House anywhere assessment]

## Assumptions implemented

1. Every drone has its own unique ID. This is stored in the Mongo store.
2. The mongo store can hold other information as well related to the drone. In this case, the sector ID which it is deployed to is stored.
3. I have assumed that another service/endpoint will update the drone information in the mongo, including the Sector ID if the drone is deployed someplace else in future.

4. I have created the tests in a Suite manner. This makes the initialisation of values easier to be used during the test and makes everything reusable to reduce effort. 
5. Tests are using `gomock` and `testify` packages to build.
6. Unit tests covers the straightforward cases for the API built.

## Usage
1. Clone out the repository in your local.
2. To run the tests (located in the `tests` folder in the repo) \
   ` cd tests`\
   `go test ./... -v -coverpkg=./...`  -> this command prints out the code coverage as well.
3. The `Dockerfile` already has all the necessary configs required to run this application. 
4. The make file has commands to build and run the applications. `make run` -> this will build and run the application in Docker. PS: Make sure docker service is running in you local.
5. The mongo connected is an online datastore therefore it is not required to change the credentials and will run out of the box. If you would like to use your own Mongo store, You can change the credentials in `config/config.yml`
6. Application will serve on port `9000`. If you wish to change, this can be changed in the `config/config.yml` file and also in the `Dockerfile`.