# Flight Route App

This is a simple Go application built using the Echo framework to manage flight routes.

## Installation

1. **Install Go:**

  Download and install Go from the [official website](https://golang.org/doc/install) according to your system's specifications.

2. **Clone the Repository:**

```sh
   git clone https://github.com/ayzahamid/go-coding-challenge.git
```
3. **Navigate to the Project Directory:**

```sh
   cd go-assessment
```

4. **Initialize Go Module:**
```sh
   go mod init flightrouteapp
```

5. **Install the Echo Framework:**
```sh
   go get -u github.com/labstack/echo/v4
```
6. **Run the Application:**
```sh
   go run main.go
```

This will start the application on port 8080 by default. You can access it in your browser at http://localhost:8080/.

## Development Choices

### Libraries and Tools

1. **Echo Framework:**

   The application is built using the Echo framework, a high-performance, minimalist Go web framework. Echo was chosen for its simplicity, ease of use, and features for building RESTful APIs.

2. **HTTP Package:**

   The standard `net/http` package from Go is used for handling HTTP requests and responses. This package provides a solid foundation for building web applications in Go.

### Code Design

1. **TicketPair Struct:**

   A `TicketPair` struct is defined to represent flight tickets with source and destination fields. This struct is used for parsing JSON request payloads and handling ticket information.

2. **Reconstruct Itinerary Function:**

   The `reconstructItinerary` function takes a slice of ticket pairs and reconstructs the itinerary based on the provided tickets. It performs validation checks on the input data and ensures a valid itinerary is generated.

### API Endpoints

1. **POST /flights:**

   The application exposes a POST endpoint at `/flights` for receiving flight ticket information in JSON format. It then reconstructs the itinerary based on the provided tickets and returns the itinerary as a JSON response.

### Running the Application

Once the application is running, you can test the `/flights` endpoint using HTTP POST requests with valid ticket data.
