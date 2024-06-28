# HTTP Proxy Server

This project implements a simple HTTP proxy server in Go. The server is designed to proxy HTTP requests to external services based on JSON data received from clients.

## Features

- **HTTP Request Proxying**: The server accepts HTTP requests from clients and forwards them to specified URLs provided in JSON format.
  
- **Request and Response Logging**: Uses `sync.Map` for temporary storage of request and response data within the server.

- **JSON Format**: Incoming requests are expected in JSON format with mandatory fields for request method (`method`), URL (`url`), and optional headers (`headers`).

- **JSON Response**: Clients receive structured JSON response containing request ID, HTTP status from the external service, headers, and response length.

## Usage

1. Initialize the new Go module:
   ```
   go mod init myproxyserver
   ```
2. Install the necessary dependencies. To do this, make sure that you have the package installed github.com/google/uuid . If not, install it:
   ```
   go get github.com/google/uuid
   ```
3. **Run the Server**: Start the server listening on port `8080` to begin accepting incoming requests.
   ```
   go run main.go
   ``` 
4. **Send Requests**: Use `curl` or any HTTP client to send POST requests to `/proxy`, providing data in JSON format.

   Open new terminal or PowerShell or CMD.
   
   Example for PowerShell:
   ```
   curl -X POST http://localhost:8080/proxy -d '{
     "method": "GET",
     "url": "http://www.google.com",
     "headers": {
       "User-Agent": "curl"
     }
   }' -H "Content-Type: application/json"
   ```
   For a more beautiful and easy-to-read output, you can use "jq": https://jqlang.github.io/jq/download/

   Install "jq" via CMD. For example:
   ```
   winget install jqlang.jq
   ```

   Example for CMD:
   ```
   curl -X POST http://localhost:8080/proxy -d "{\"method\":\"GET\", \"url\":\"http://www.google.com\", \"headers\":{\"User-Agent\":\"curl\"}}" -H "Content-Type: application/json" | jq .
   ```

## Output example
![image](https://github.com/RaimbekovA/http-proxy-server/assets/63358961/717d3c5e-f386-4d63-b9f0-a8dfbf920384)

