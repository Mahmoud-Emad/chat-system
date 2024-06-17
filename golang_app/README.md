# Chat System API - Golang App

This project is part of a chat system API built with Golang. The Golang application handles the creation of chats and messages within the system. The system ensures unique numbering for chats within an application and messages within a chat.

## Features

- Create chats within applications, ensuring unique chat numbers per application.
- Create messages within chats, ensuring unique message numbers per chat.
- Handle race conditions using database transactions.
- Data is stored in MySQL.

## Prerequisites

- Docker and Docker Compose are installed on your system.

## Docker Compose

Build and run the Golang application using Docker Compose:
PS: Run it on the main repo root.

```bash
docker-compose up
```

This will start the following services:

- MySQL database
- Golang API

### API Endpoints

#### Chats

- **Create Chat**
  - Method: `POST`
  - Endpoint: `/applications/{application_token}/chats`
  - Response:

    ```json
    {
      "number": 1
    }
    ```

#### Messages

- **Create Message**
  - Method: `POST`
  - Endpoint: `/applications/{application_token}/chats/{chat_number}/messages`
  - Request Body:

    ```json
    {
      "body": "Hello there!"
    }
    ```

  - Response:

    ```json
    {
      "body": "Hello there!",
      "chat_number": "<chat_number>",
      "message_number": 1
    }
    ```

### Handling Race Conditions

The API ensures race conditions are handled by using database transactions with `FOR UPDATE` locks. This prevents multiple requests from creating chats or messages with the same number simultaneously, the race conditions are only managed in the `golang_app` just to do the point in the task, no need to do it over and over since it's not a real app.

### Running the Application

To run the Golang application:

1. You can execute the `./main` after doing `go build -o main .`.
2. You can run it via `Docker` by running `docker compose up --build golang_app`.
3. You can run the whole stack using `docker compose up --build -d` in the root path.
