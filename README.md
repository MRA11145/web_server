# Simple Go Web Server

A basic web server implementation in Go that serves static files and handles form submissions.

## Project Structure 

## Features

- Static file serving
- Form handling
- Basic routing
- Hello endpoint

## Implementation Details

### 1. Static File Serving
The server serves static files from the `static` directory, including:
- A basic index page
- A form page for user input

### 2. Endpoints

#### Static Files (`/`)
- Serves files from the static directory
- Default route serves index.html

#### Form Handler (`/form`)
- Handles POST requests for form submissions
- Processes name and address fields
- Returns formatted response with submitted data

#### Hello Handler (`/hello`)
- Responds to GET requests only
- Returns a simple "hello!" message
- Includes proper error handling for invalid methods and paths

## How to Run

1. Ensure you have Go installed (version 1.23.3 or later)
2. Clone the repository
3. Run the server:

```bash
go run main.go
```
4. Access the server at `http://localhost:8000`

## API Endpoints

| Endpoint | Method | Description |
|----------|---------|------------|
| `/` | GET | Serves static files |
| `/form` | POST | Handles form submissions |
| `/hello` | GET | Returns hello message |

## Code References

### Main Server Setup
The main server configuration can be found in main.go:

```go:main.go
startLine: 36
endLine: 48
```

### Form Handler Implementation
Form processing logic:

```go:main.go
startLine: 9
endLine: 20
```

## Form Structure
The HTML form structure is defined in:

```html:static/form.html
startLine: 6
endLine: 13
```

## Error Handling

The application includes error handling for:
- Form parsing errors
- Invalid routes
- Unsupported HTTP methods
- Server startup issues

## Contributing

Feel free to submit issues and enhancement requests.

## License

This project is open-source and available under the MIT License.