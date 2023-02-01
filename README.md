# Dummy API

This Go application provides a straightforward implementation of a REST API that facilitates the management of user data, offering endpoints for adding and modifying information without requiring application restart.


<br />

## Getting Started

### Prerequisites
To run this application, you will need to have [Go installed](https://go.dev/dl/) on your machine.

### Installing
Clone the repository to your local machine using the following command:

```shell
git clone https://github.com/alsolovyev/dummy-api.git
```

Change into the directory:

```shell
cd dummy-api
```

Run the application:
```shell
make run
```


<br />

## Usage
Create a file that includes the necessary data within the "data" directory:
```shell
echo '[{"name": "Jane"}, {"name": "Willa"}]' > ./data/users.json
echo 'Lorem ipsum dolor sit amet' > ./data/text
```

Send a request to retrieve data from the file that you have created:
```shell
curl 'http://localhost:8181/api/v1/file/users.json' | json_pp
curl 'http://localhost:8181/api/v1/file/text'
```

It is now possible to modify, add, or delete files without the requirement of restarting the application.


<br />

## Endpoints
The following endpoints are available:

### GET /api/v1/file/{name}
Retrieves a specific file based on its name and parses its data based on its extension. Currently, it supports parsing of JSON formatted files with the ".json" extension, as well as text files without an extension.


<br />

## Build With

* [Golang](https://go.dev)
* [Chi](https://github.com/go-chi/chi)


<br />

## License
This project is licensed under the MIT License
