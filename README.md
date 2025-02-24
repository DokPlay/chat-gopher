# Chat Gopher

Chat Gopher — это приложение для обмена сообщениями через WebSocket с использованием Go. Оно поддерживает отправку и получение сообщений, а также сохранение сообщений в базе данных.

## Особенности:
- Подключение через WebSocket для обмена сообщениями в реальном времени.
- Создание сообщений с сохранением в базе данных.
- Исторические данные сообщений могут быть получены за определённый период.
-Отправить сообщение (Первый клиент)
-Просмотр сообщений в реальном времени (Второй клиент)
-История сообщений (Третий клиент)
-Swagger-документация
## Установка и запуск

Для запуска приложения можно использовать как Docker, так и обычную установку через Go.

### 1. Клонируйте репозиторий:

```bash
git clone https://github.com/DokPlay/chat-gopher.git
cd chat-gopher
```

2. Запуск через Docker
Для упрощения развертывания, вы можете использовать Docker.

Создайте файл .env с параметрами подключения к базе данных:
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=username
DB_PASSWORD=password
DB_NAME=chat_gopher
```
Запустите контейнеры с помощью Docker Compose:
```bash
docker-compose up --build
```
Этот процесс развернет приложение и базу данных в Docker контейнерах. Приложение будет доступно на localhost:8080.

3. Запуск без Docker
Если вы предпочитаете запускать приложение без Docker, выполните следующие шаги:

 1.Установите зависимости:
```bash
go mod tidy
```
 2.Запустите сервер:
```bash
go run main.go
```
Сервер будет запущен на http://localhost:8080.

4. 
   ## API

Chat Gopher предоставляет несколько API-эндпоинтов для работы с сообщениями.

### Получить список сообщений за диапазон дат

**Метод:** `GET /api/messages`

Этот эндпоинт позволяет получить список сообщений за указанный диапазон дат.

**Параметры запроса:**

- `from` (обязательный) — начальная дата в формате RFC3339.
- `to` (обязательный) — конечная дата в формате RFC3339.

**Пример запроса:**

```http
GET /api/messages?from=2025-02-01T00:00:00Z&to=2025-02-10T23:59:59Z
```
```
Ответ:

200 OK — успешный ответ, возвращает список сообщений.
400 Bad Request — если параметры некорректны.
500 Internal Server Error — если произошла ошибка на сервере.
Пример ответа:
[
  {
    "created_at": "2025-02-01T12:00:00Z",
    "id": 1,
    "sequence": 1,
    "text": "Hello, World!"
  },
  {
    "created_at": "2025-02-02T12:00:00Z",
    "id": 2,
    "sequence": 2,
    "text": "How are you?"
  }
]
```
Отправить одно сообщение
Метод: POST /api/messages

Этот эндпоинт позволяет отправить одно сообщение. Оно будет сохранено в базе данных и рассылается всем подключенным клиентам через WebSocket.

Пример тела запроса:
```
{
  "sequence": 1,
  "text": "Hello, World!"
}
```

```
Ответ:

200 OK — успешный ответ, возвращает информацию о сообщении, включая created_at, id, sequence и text.
400 Bad Request — если запрос некорректен.
500 Internal Server Error — если произошла ошибка на сервере.

```
Пример ответа:
```
{
  "created_at": "2025-02-01T12:00:00Z",
  "id": 1,
  "sequence": 1,
  "text": "Hello, World!"
}
```
Структуры данных
Message
```
{
  "created_at": "2025-02-01T12:00:00Z",
  "id": 1,
  "sequence": 1,
  "text": "Hello, World!"
}
```

ErrorResponse
```
{
  "error": "Ошибка запроса"
}
```
Важные замечания:
WebSocket: Сервер поддерживает подключение через WebSocket для обмена сообщениями в реальном времени.
Конфигурация: Все параметры подключения к базе данных и другие настройки можно найти в файле конфигурации.

Лицензия
Этот проект использует лицензию MIT.

MIT License

Copyright (c) [25.02.2025] [DokPlay]

Данное ПО предоставляется "как есть", без каких-либо выраженных или подразумеваемых гарантий, включая, но не ограничиваясь, гарантией товарной пригодности или пригодности для конкретной цели. В любом случае авторы не несут ответственности за любой ущерб, возникающий из использования этого ПО.

 
    
    English Version:
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








