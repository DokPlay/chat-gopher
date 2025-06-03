<<<<<<< HEAD
[→ Читать по-русски (README.RU.md)](README.RU.md)

Chat Gopher
Chat Gopher is a messaging application using WebSocket and Go. It supports sending and receiving messages, as well as storing messages in the database.

Features:
WebSocket-based real-time messaging.
Creating messages with storage in the database.
Historical message data can be fetched within a specified date range.
Installation and Setup
You can run the application using either Docker or a standard Go installation.

1. Clone the repository:
   ```
   git clone https://github.com/DokPlay/chat-gopher.git
   cd chat-gopher
   ```
   2. Run with Docker
For easy deployment, you can use Docker.

Create a .env file with database connection parameters:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=username
DB_PASSWORD=password
DB_NAME=chat_gopher
```
Run the container:
```
docker-compose up --build
```
3. Run manually with Go
If you prefer not to use Docker, you can run the server manually with Go.

Install dependencies:
```
go mod tidy
```
Create the .env file with database connection parameters as shown above and run the server:
```
go run main.go
```
The server will be available at http://localhost:8080.

API
Get a list of messages within a date range
GET /api/messages

This endpoint allows you to get a list of messages within the specified date range.

Parameters:

from (required) — start date in RFC3339 format.
to (required) — end date in RFC3339 format.
Example request:
```
GET /api/messages?from=2025-02-01T00:00:00Z&to=2025-02-10T23:59:59Z
```
Responses:

200 OK — success, returns a list of messages.
400 Bad Request — if parameters are incorrect.
500 Internal Server Error — if an error occurred on the server.
Send a single message
POST /api/messages

This endpoint allows you to send a single message. It will be saved to the database and broadcast to all connected clients via WebSocket.

Example request body:
```
{
  "sequence": 1,
  "text": "Hello, World!"
}
```
Responses:

200 OK — success, returns message information including created_at, id, sequence, and text.
400 Bad Request — if the request is incorrect.
500 Internal Server Error — if an error occurred on the server.
Data Structures
Message:
```
{
  "created_at": "2025-02-01T12:00:00Z",
  "id": 1,
  "sequence": 1,
  "text": "Hello, World!"
}
```
ErrorResponse:
```
{
  "error": "Request error"
}
```
Important Notes
WebSocket: The server supports WebSocket for real-time messaging.
Configuration: All database connection parameters and other settings can be found in the configuration file.

License
This project is licensed under the MIT License.

MIT License

Copyright (c) [25.02.2025] [DokPlay]

This software is provided "as is", without any express or implied warranties, including but not limited to warranties of merchantability or fitness for a particular purpose. In no event shall the authors be liable for any damages arising from the use of this software.






>>>>>>> 89c6bed18d1fea9599b7904ca731adae594d4d89
