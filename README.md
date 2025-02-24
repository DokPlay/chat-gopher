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





