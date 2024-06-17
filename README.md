# Chat System API

This project is an idea about implementing a chat system API `Rails` and `Golang`.

- The [Golang server](./golang_app/README.md) handles the creation of chats and messages within the system. The system ensures unique numbering for chats within an application and messages within a chat.

- The [Rails server](./rails_app/README.md) handles the creation of applications, gets chats, and gets messages within the system It also provides an endpoint for searching messages using Elasticsearch. It uses MySQL as the primary data store.

## Run the whole stack

This project uses `Docker` and `Docker Compose`, make sure that these packages are installed on your machine.

After creating the `.env.production` file on the main root and adding all of the required env vars see the [sample](./.env.sample):

```bash
  docker compose up --build
```

Want to see all of the [endpoints](./ENDPOINTS.md)?
**PS:** It was on my mind to use the 'Redis DB' as a caching layer, but I didn't have the time. :(
