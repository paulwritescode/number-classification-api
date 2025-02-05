# Number Classification API

## Overview

This is a simple REST API written in **Go** using the **Gorilla Mux** router. The API allows users to classify numbers based on various properties such as being an Armstrong number, even/odd, prime, or perfect. It also integrates CORS middleware to handle cross-origin requests.

---

## Features

- Classifies numbers based on mathematical properties.
- Handles user input and provides error messages for invalid inputs.
- Implements CORS support to allow cross-origin requests.
- Uses **Gorilla Mux** for routing.

---

## Technologies Used

- **Go** (Golang) - The main programming language.
- **Gorilla Mux** - Router for handling API routes.
- **JSON Encoding/Decoding** - To format API responses.
- **Net/HTTP** - Standard Go package for handling HTTP requests.
- **CORS Middleware** - Custom middleware to allow cross-origin requests.

---

## Installation

### **Prerequisites**

Ensure you have the following installed:

- [Go](https://go.dev/doc/install) (latest stable version)

### **Clone the Repository**

```sh
git clone https://github.com/yourusername/numbers-api.git
cd numbers-api
```

### **Install Dependencies**

```sh
go mod tidy
```

---

## Running the API

### **Start the Server**

```sh
go run main.go
```

The server will start on port `3000`.

---

## API Endpoints

### **Classify Number**

**Endpoint:** `GET /api/classify-number?number={value}`

#### **Request Example**

```sh
curl -X GET "http://localhost:3000/api/classify-number?number=371"
```

#### **Response Example**

```json
{
  "number": 371,
  "is_prime": false,
  "is_perfect": false,
  "properties": ["armstrong", "odd"]
}
```

#### **Error Handling**

If an invalid number is provided:

```json
{
  "number": "alphabet",
  "error": true
}
```

---

## CORS Handling

This API supports **CORS** to allow requests from different origins. The middleware is applied globally.

### **CORS Middleware Implementation**

```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

---

## Deployment

### **Run as a Binary**

```sh
go build -o classify-number
./classify-number
```

## Contributing

Feel free to contribute by forking the repo and submitting a pull request.

---

## License

This project is licensed under the MIT License.

