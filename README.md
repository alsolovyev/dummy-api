# Dummy API

This Go application adheres to the clean architecture pattern and provides a straightforward implementation of a REST API that facilitates the management of user data, offering endpoints for adding and modifying information without requiring application restart.


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

## This project uses the following directory structure to organize the codebase:

* The `cmd` directory contains the main entry point of the application. The <project-name> directory holds the `api/main.go` file that sets up and starts the application.
* The `internal` directory holds the core logic of the application and is organized into four main subdirectories, each of which represents a layer in accordance with the clean architecture pattern:
  * `entity`: the entity layer represents the core business entities of the application. It represents the objects that are used to model the domain and encapsulate the data and behavior of the application. These entities are independent of any framework, library, or external dependency and are therefore highly reusable and testable. They define the application's data model and the relationships between objects, and are the source of truth for the application's business logic. The Entity layer is isolated from other layers, such as the Presentation or Data Access layer, making it easy to change or replace these layers without affecting the Entity layer.

  * `repository`: the repository layer represents the persistence logic of the application. It defines the interactions between the entities and the data storage, and it is responsible for loading and saving the entities. The repository layer is typically composed of several repositories, each one representing a specific type of entity.

  * `usecase`: the use case layer represents the business logic of the application. It defines the interactions between the entities and the external systems, and it is responsible for performing the operations that the application needs to accomplish. The use case layer is typically composed of several use cases, each one representing a specific operation or set of operations.

  * `controller`: the controller layer represents the interface of the application. It defines the interactions between the use cases and the external systems, such as the user interface or an API. The controller layer is typically composed of several controllers, each one representing a specific set of use cases. The controllers are responsible for receiving the requests from the external systems, validating the input data, and calling the appropriate use cases to handle the request.


* The `pkg` directory holds code that is intended to be reused across multiple applications. It contains helper packages, such as an HTTP client or a custom error handling package.
* The `README.md` file provides documentation and instructions for the project.


<br />

## Build With

* [Golang](https://go.dev)
* [Chi](https://github.com/go-chi/chi)


<br />

## License
This project is licensed under the MIT License
