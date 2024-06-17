# Chat System API `Endpoints`

1. Rails server

- `/applications`

- Create application:
- - URL : <http://localhost:8000/applications>
- - Payload: ```json { "name": "<app_name>" }```
- - Method: POST

- Get all applications:
- - URL : <http://localhost:8000/applications>
- - Method: GET

- Show application:
- - URL : <http://localhost:8000/applications/{application_token}>
- - Method: GET

- `/chats`

- Get all chats:
- - URL : <http://localhost:8000/applications/{application_token}/chats>
- - Method: GET

- Show chat:
- - URL : <http://localhost:8000/applications/{application_token}/chats/{chat_number}>
- - Method: GET

- `/messages`

- Get all messages:
- - URL : <http://localhost:8000/applications/{application_token}/chats/{chat_number}>
- - Method: GET

- Show message:
- - URL : <http://localhost:8000/applications/{application_token}/chats/{chat_number}/messages/{message_number}>
- - Method: GET

- `/search`

- List searched messages:
- - URL : <http://localhost:8000/applications/{application_token}/chats/12/messages/search/?query={message}>
- - Method: GET

2. GO server

- `/chats`

- Create chat:
- - URL : <http://localhost:8001/applications/{application_token}/chats/>
- - Method: POST

- `/messages`

- Create message:
- - URL : <http://localhost:8001/applications/{application_token}/chats/messages/>
- - Payload: ```json { "body": "<message_body>" }```
- - Method: POST
