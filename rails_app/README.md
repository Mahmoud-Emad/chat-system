# Chat System API - Rails App

## Overview

This API allows clients to create applications, get chats, and get messages. Each application has a unique token used by devices to identify and send chats to the application. It also provides an endpoint for searching messages using Elasticsearch. It uses MySQL as the primary data store.

## Features

1. **Applications**: Create new applications with a unique token and name.
2. **Chats**: Get chats within an application.
3. **Messages**: Get messages within a chat.
4. **Search**: Search through messages of a specific chat using partial matches.

## Endpoints

### Applications

- **Create Application**
  - `POST /applications`
  - Request Body: `{ "name": "Application Name" }`
  - Response: `{ "token": "generated_token", "name": "Application Name" }`

- **Get Applications**
  - `GET /applications`
  - Response: `[ { "token": "app_token", "name": "Application Name" } ]`

- **Update Application**
  - `PUT /applications/:token`
  - Request Body: `{ "name": "Updated Name" }`
  - Response: `{ "token": "app_token", "name": "Updated Name" }`

### Chats

- **Get Chats**
  - `GET /applications/:token/chats`
  - Response: `[ { "number": 1, "created_at": "timestamp", "updated_at": "timestamp" } ]`

### Messages

- **Get Messages**
  - `GET /applications/:token/chats/:number/messages`
  - Response: `[ { "number": 1, "created_at": "timestamp", "updated_at": "timestamp", "body": "Message body" } ]`

- **Search Messages**
  - `GET /applications/:token/chats/:number/messages/search`
  - Query Params: `?query=search_term`
  - Response: `[ { "number": 1, "created_at": "timestamp", "updated_at": "timestamp", "body": "Message body" } ]`

## Running the Application

### Prerequisites

- Docker
- Docker Compose

### Setup

1. **Clone the repository**
2. **Create a `.env.production` file `in the root dir`** with the following environment variables:

    ```INI
    MYSQL_ROOT_PASSWORD=password
    MYSQL_USER=root
    MYSQL_DATABASE=mysql
    MYSQL_PASSWORD=password
    ELASTICSEARCH_URL=<http://elasticsearch:9200>
    SECRET_KEY_BASE=xxx
    ```

3. **Run Docker Compose `in the root dir`**:

  ```bash
    docker-compose up
  ```

This command will build the Docker images and start the containers for the web application, MySQL, and Elasticsearch.

## Design Considerations

### Indexing and Optimization

- Appropriate indices are added to the MySQL tables to optimize the endpoints for creating, updating, and reading applications, chats, and messages.
- Elasticsearch is used to efficiently handle the search functionality, allowing for fast and scalable message searches.

## Future Enhancements

- **Bonus Features**:
  - Implementing the chat and message creation endpoints as a Golang app.
