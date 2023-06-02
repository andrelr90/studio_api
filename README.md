# Studio API Project
This repository contains the implementation of an API for managing classes and bookings in a studio.

### Installing


### Executing


### Available endpoints and usage


### General approach considerations
This project assumes that each date has only one class, therefore, overlaps between classes are validated. The API implementation includes the fundamental RESTful endpoints: GET, POST, PUT, and DELETE. These endpoints facilitate essential operations such as retrieving data, creating new resources, updating existing resources, and deleting resources as required by the business logic.

This project was developed using Go, a powerful programming language known for its simplicity and high efficiency. To develop the API, this project uses the Gin library, a widely used and trusted framework, to handle the implementation of the endpoints. By leveraging Gin, it is possible to efficiently define and manage the RESTful API endpoints, enabling seamless communication between the client and server. To maintain the integrity and validity of the basic data models, the project employs bindings and validators provided by Gin. Finally, the project has implemented a specialized date entity that focuses solely on representing dates without including the time component.

### Architecture
The main architecture of the project follows a layered-based architecture. The implemented layers are "model," "api," and "repository." The model layer represents the core business domain of the application, dealing with the representations of classes and bookings, and validating the intrinsic logic of the entities created. The API layer coordinates actions, handles business input/output logic, and is projected to eventually operate as an intermediary between the presentation layer and the model layer. Finally, the repository layer handles information retrieval and manipulation logic. In this version, it is implemented completely without persistence. The approach taken has several advantages regarding code maintance.

### Testing
Using a layered architecture also has some advantages regarding tests, as responsibilities are well distributed across entities in different packages. Tests for the API layer check the availability of endpoints, validate the input and output of data, and also verify the possible code and message errors returned by each endpoint. Regarding the model layer, date representations and logic are also unit-tested. Unfortunately, tests for the bindings and validators of Gin are not possible directly through the model layer, so they depend on the API tests. Still, using a well-defined validation tool has advantages regarding maintenance and even reliability. Finally, repository tests validate the manipulation logic and data structures of the set of classes and bookings.

Though coverage might be a controversial metric, the code developed has 100% test coverage in all files and modules, except for the main() function. Tests not only iterate through all the code but have the intention of effectively verifying corner cases, checking specific validations, and ensuring that the code can be properly maintained. Tests can be performed by running the following command in the project's root folder:

```
go test -cover ./...
```

Test coverage can also be checked through the file `cover.html` in the root folder.

### Performance considerations
As Go is a high-performance language, most of the features implemented in this project are designed to handle the problem with excellent performance. While many performance issues in similar projects are typically addressed by databases, this project utilizes non-persistent repositories, requiring performant data structures. Whenever possible, costs have been reduced to constant or logarithmic time complexity. A hash data structure is used for storing books, while a custom data structure is employed for storing classes, specifically focusing on the cost of finding available classes on a given date, which is necessary for some of the most frequently used booking operations. This approach is driven by the understanding that classes are typically less frequently manipulated or modified compared to bookings, making higher costs more acceptable. It is worth noting that the performance of these data structures could be further improved through incremental enhancements in future implementations.
