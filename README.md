# Temperature Converter

This is a simple web application that converts temperatures between Celsius and Fahrenheit and saves the results to a MongoDB database.

## How to use

### Using Go

To run the application from Go, you need to have Go and MongoDB installed on your computer. Once you have them installed, follow these steps:

1. Clone this repository to your computer.
2. Open a terminal and navigate to the root directory of the cloned repository.
3. Start MongoDB by running the command `mongod`.
4. Run the command `go run .` to start the application.
5. Open your web browser and go to `http://localhost:8080` to access the application.

### Using Docker

To run the application from Docker, you need to have Docker installed on your computer. Once you have it installed, follow these steps:

1. Clone this repository to your computer.
2. Open a terminal and navigate to the root directory of the cloned repository.
3. Run the command `docker-compose up` to start the application and MongoDB.
4. Open your web browser and go to `http://localhost:8080` to access the application.

Note that when running the application from Docker, the application will be able to connect to the MongoDB container by using the hostname `mongodb` instead of `localhost`. So, in the application code, you need to change the MongoDB URI to `mongodb://mongodb:27017` instead of `mongodb://localhost:27017`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

