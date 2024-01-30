# SMS Gateway

## Overview

SMS Gateway is a scalable messaging platform built to efficiently handle the sending and receiving of SMS messages. The
architecture is designed for horizontal scaling, leveraging MongoDB for data storage and Go (Golang) for improved
scalability.

## Architecture

### Components:

#### 1. MongoDB Database:

- ##### Purpose: Persistent storage for message data.
- ##### Configuration: MongoDB configuration can be set in config.go for scalability and horizontal scaling purposes.

#### 2. Golang Application:

- ##### Purpose: The core of the SMS Gateway system, responsible for handling message processing, routing, and communication with the MongoDB database.
- ##### Scalability: Golang is chosen for its concurrency support, making it well-suited for handling concurrent requests and improving scalability.

### Horizontal Scaling:

The system is designed for horizontal scaling to handle increasing loads by adding more instances of the application.
MongoDB plays a crucial role in this architecture, allowing data to be distributed across multiple nodes for improved
performance.
###  Build and Run Instructions:

#### 1. Prerequisites:
Install MongoDB and Golang on your machine.

#### 2. Configuration:
Set up MongoDB configurations in [config.go](internal/config/config.go).

#### 3. Run the Application:
Execute `go run cmd/main.go` in the root directory.

#### 4. Scaling:
To horizontally scale the application, deploy multiple instances of the Golang application and configure them to connect to the same MongoDB cluster.

Project Structure:

```bash

├── cmd
│   └── main.go            # Entry point for the application
├── contract               # Postman collection documentation 
├── internal
│   ├── config
│   │   └── config.go      # Configuration settings
│   ├── api/http
│   │   ├── handlers.go    # Process and handler
│   │   ├── middlewares.go # Middleware pipeline
│   │   ├── server.go      # Routing logic and serve  
│   ├── service            # Business logic for sms handling
├── pkg
│   ├── queue
│   │   └── queue.go       # Custom concurrent queue
└── README.md              # Documentation for the project
```