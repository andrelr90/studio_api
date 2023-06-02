# Studio API Project
This repository contains the implementation of an API for managing classes and bookings in a studio.

## Installing
This project was developed in `go version go1.20.4 windows/386`. If you don't have Go installed yet, you can get it from https://go.dev/doc/install. 

If you need to reset the dependencies of this project, you can delete the files `/src/go.mod` and `/src/go.sum` execute the following in the `/src` folder:
```
go mod init studio_api/src
go mod init tidy
```

## Executing
Assuming the project was installed successfully, navigate to the `/src` folder of the project and run:
```
go run main.go
```


## Available endpoints and usage
The API implementation includes the fundamental RESTful endpoints: GET, POST, PUT, and DELETE. These endpoints facilitate essential operations such as retrieving data, creating new resources, updating existing resources, and deleting resources as required by the business logic. Both classes and bookings are populated with two examples at the start of the program.

### Classes:
#### GET /classes/
Returns all classes currently available. The output should be similar to the following:
```
[
    {
        "id": 0,
        "name": "Pilates",
        "start_date": "2023-01-01",
        "end_date": "2023-02-01",
        "capacity": 30
    },
    {
        "id": 1,
        "name": "Yoga",
        "start_date": "2023-02-02",
        "end_date": "2023-03-01",
        "capacity": 25
    }
]
```

#### GET /classes/{id}
Returns a class given its id. The output should be similar to the following assuming the id 0:
```
{
    "id": 0,
    "name": "Pilates",
    "start_date": "2023-01-01",
    "end_date": "2023-02-01",
    "capacity": 30
}
```

#### POST /classes/
Adds a class. The input should be similar to the following:
```
{
    "name": "STRING",
    "start_date": "YYYY-MM-DD",
    "end_date": "YYYY-MM-DD",
    "capacity": INT
}
```
The output of this call should be similar to:
```
{
    "id": INT,
    "name": "STRING",
    "start_date": "YYYY-MM-DD",
    "end_date": "YYYY-MM-DD",
    "capacity": INT
}
```
Notice that the id is automatically given incrementally.

#### DELETE /classes/{id}
Deletes a class given its id. The output should be similar to the following assuming the id 0:
```
{
    "message": "Class deleted successfully"
}
```
Deleting a class deletes its bookings cascade.

#### PUT /classes/{id}
Updates a class given its id. The input should be similar to the following:
```
{
    "name": "STRING",
    "start_date": "YYYY-MM-DD",
    "end_date": "YYYY-MM-DD",
    "capacity": INT
}
```
The output of this call should be similar to:
```
{
    "id": INT,
    "name": "STRING",
    "start_date": "YYYY-MM-DD",
    "end_date": "YYYY-MM-DD",
    "capacity": INT
}
```
Notice that the ids cannot be updated. When updating a class's start or end date, any bookings scheduled for the dates that fall outside the new class duration will be automatically deleted.

### Bookings:
#### GET /bookings/
Returns all bookings currently available. The output should be similar to the following:
```
[
    {
        "id": 0,
        "name": "John Doe",
        "date": "2023-01-01"
    },
    {
        "id": 1,
        "name": "Jane Smith",
        "date": "2023-01-01"
    }
]
```

#### GET /bookings/{id}
Returns a booking given its id. The output should be similar to the following assuming the id 0:
```
{
    "id": 0,
    "name": "John Doe",
    "date": "2023-01-01"
}
```

#### POST /bookings/
Adds a booking. The input should be similar to the following:
```
{
    "name": "STRING",
    "date": "YYYY-MM-DD"
}
```
The output of this call should be similar to:
```
{
    "id": INT,
    "name": "STRING",
    "date": "YYYY-MM-DD"
}
```
Notice that the id is automatically given incrementally.

#### DELETE /bookings/{id}
Deletes a booking given its id. The output should be similar to the following assuming the id 0:
```
{
    "message": "Class booking deleted successfully"
}
```

#### PUT /bookings/{id}
Updates a booking given its id. The input should be similar to the following:
```
{
    "name": "STRING",
    "date": "YYYY-MM-DD"
}
```
The output of this call should be similar to:
```
{
    "id": INT,
    "name": "STRING",
    "date": "YYYY-MM-DD"
}
```
Notice that the ids cannot be updated.


## Keypoints
The following has been considered during development.

### General approach
This project assumes that each date has only one class, therefore, overlaps between classes are validated. The project was developed using Go, a powerful programming language known for its simplicity and high efficiency. To develop the API, this project uses the Gin library, a widely used and trusted framework, to handle the implementation of the endpoints. By leveraging Gin, it is possible to efficiently define and manage the RESTful API endpoints, enabling seamless communication between the client and server. To maintain the integrity and validity of the basic data models, the project employs bindings and validators provided by Gin. Finally, the project has implemented a specialized date entity that focuses solely on representing dates without including the time component.

### Architecture
The main architecture of the project follows a layered-based architecture. The implemented layers are "model," "api," and "repository." The model layer represents the core business domain of the application, dealing with the representations of classes and bookings, and validating the intrinsic logic of the entities created. The API layer coordinates actions, handles business input/output logic, and is projected to eventually operate as an intermediary between the presentation layer and the model layer. Finally, the repository layer handles information retrieval and manipulation logic. In this version, it is implemented completely without persistence. The approach taken has several advantages regarding code maintance.

### Testing
Using a layered architecture also has some advantages regarding tests, as responsibilities are well distributed across entities in different packages. Tests for the API layer check the availability of endpoints, validate the input and output of data, and also verify the possible code and message errors returned by each endpoint. Regarding the model layer, date representations and logic are also unit-tested. Unfortunately, tests for the bindings and validators of Gin are not possible directly through the model layer, so they depend on the API tests. Still, using a well-defined validation tool has advantages regarding maintenance and even reliability. Finally, repository tests validate the manipulation logic and data structures of the set of classes and bookings.

Though coverage might be a controversial metric, the code developed has 100% test coverage in all files and modules, except for the main() function. Tests not only iterate through all the code but have the intention of effectively verifying corner cases, checking specific validations, and ensuring that the code can be properly maintained. Tests can be performed by running the following command in the project's `/src` folder:

```
go test -cover ./...
```

Test coverage can also be checked through the file `cover.html` in the `/test_outputs` folder.

### Performance
As Go is a high-performance language, most of the features implemented in this project are designed to handle the problem with excellent performance. While many performance issues in similar projects are typically addressed by databases, this project utilizes non-persistent repositories, requiring performant data structures. Whenever possible, costs have been reduced to constant or logarithmic time complexity. A hash data structure is used for storing books, while a custom data structure is employed for storing classes, specifically focusing on the cost of finding available classes on a given date, which is necessary for some of the most frequently used booking operations. This approach is driven by the understanding that classes are typically less frequently manipulated or modified compared to bookings, making higher costs more acceptable. It is worth noting that the performance of these data structures could be further improved through incremental enhancements in future implementations.

### Possible next steps
The following have been disconsidered for simplicity and can be addressed in further implementations:
- Implement Overbooking check
- Use a database

