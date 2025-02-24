# Пример кода для Chat Gopher

Этот репозиторий содержит пример реализации для чат-приложения Chat Gopher. Ниже представлена таблица с кодом для различных файлов, используемых в проекте.

| Название файла    | Код                                        |
|-------------------|--------------------------------------------|
| `main.go`         | ```go                                      |
|                   | package main                               |
|                   | import "fmt"                               |
|                   | func main() {                              |
|                   |     fmt.Println("Hello, World!")           |
|                   | }                                          |
|                   | ```                                        |
| `server.go`       | ```go                                      |
|                   | package server                             |
|                   | import "net/http"                          |
|                   | func StartServer() {                       |
|                   |     http.ListenAndServe(":8080", nil)      |
|                   | }                                          |
|                   | ```                                        |
| `handlers.go`     | ```go                                      |
|                   | package handlers                           |
|                   | import (                                   |
|                   |     "net/http"                             |
|                   |     "chat_gopher/internal/models"          |
|                   |     "github.com/gorilla/websocket"         |
|                   |     "github.com/sirupsen/logrus"           |
|                   | )                                          |
|                   | var (                                       |
|                   |     upgrader = websocket.Upgrader{         |
|                   |         CheckOrigin: func(r *http.Request) bool { return true }, |
|                   |     }                                        |
|                   |     clientsMu sync.Mutex                   |
|                   |     clients = make(map[*websocket.Conn]bool) |
|                   | )                                          |
|                   | // WebSocketHandler - хендлер для работы с WebSocket |
|                   | func WebSocketHandler(w http.ResponseWriter, r *http.Request) { |
|                   |     logrus.Info("Обработка WebSocket-подключения") |
|                   |     conn, err := upgrader.Upgrade(w, r, nil) |
|                   |     if err != nil {                        |
|                   |         logrus.WithError(err).Error("Ошибка апгрейда соединения до WebSocket") |
|                   |         return                              |
|                   |     }                                        |
|                   |     logrus.Info("WebSocket-подключение успешно") |
|                   |     clientsMu.Lock()                        |
|                   |     clients[conn] = true                    |
|                   |     clientsMu.Unlock()                      |
|                   |     go readLoop(conn)                       |
|                   | }                                          |
|                   | ```                                        |

## Описание

1. **`main.go`** - Основной файл, содержащий точку входа для приложения, настройку и запуск сервера.
2. **`server.go`** - Файл с функцией для старта HTTP-сервера на указанном порту.
3. **`handlers.go`** - Хендлеры для работы с WebSocket, включая подключение клиентов и рассылку сообщений.

## Лицензия

Этот проект использует **MIT License**.

Copyright (c) 2025, DokPlay

Данное ПО предоставляется "как есть", без каких-либо выраженных или подразумеваемых гарантий, включая, но не ограничиваясь, гарантией товарной пригодности или пригодности для конкретной цели. В любом случае авторы не несут ответственности за любой ущерб, возникающий из использования этого ПО.
