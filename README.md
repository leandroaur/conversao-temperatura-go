# Temperature Converter in Go with MongoDB

This is a temperature converter in Go that uses MongoDB to store the converted temperatures. This README file will provide instructions on how to use the program.

## Prerequisites

To use this program, you will need:

- Go installed on your computer
- MongoDB instance running and accessible at the default port (27017) 
  - if you don't have MongoDB installed, you can download it from the official website: https://www.mongodb.com/try/download/community
- Docker installed on your computer (optional)

## Running the program

### Running the program with Go

1. Clone this repository to your local machine.
git clone https://github.com/your-username/temperature-converter.git

2. Navigate to the project directory.

cd temperature-converter

3. Make sure that MongoDB is running and accessible on the default port (27017).

4. Build the program using Go.

`go build convert.go`

5. Run the program.

`go run convert.go`

The program should now be running and ready to accept temperature inputs.

### Running the program with Docker

1. Clone this repository to your local machine.

git clone https://github.com/your-username/temperature-converter.git

2. Navigate to the project directory.

`cd temperature-converter`

3. Build the Docker image.

`docker build -t temperature-converter .`

4. Run the Docker container.

`docker run -p 8080:8080 -d temperature-converter`

The program should now be running inside the Docker container and accessible on port 8080.

## Usage

Once the program is running, you can use it to convert temperatures. The program accepts input temperatures in Fahrenheit and converts them to Celsius.

To convert a temperature, open your web browser and navigate to `http://localhost:8080/convert/<temperature>`, where `<temperature>` is the temperature you want to convert in Fahrenheit.

For example, to convert a temperature of 75 degrees Fahrenheit, navigate to `http://localhost:8080/convert/75` in your web browser. The program will display the converted temperature in Celsius.

## Conclusion

That's it! You should now be able to use the temperature converter in Go with MongoDB. If you have any questions or issues, feel free to contact me. Thank you for using my program!

