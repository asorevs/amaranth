# Real-Time Communication Microservice

This repository contains the source code for a real-time communication microservice built using Golang. The microservice provides a WebSocket API for users to send messages to each other and subscribe to specific channels for updates.

---

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Prerequsites](#prerequisites)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- Real-time communication via WebSocket connections
- User authentication using JWT tokens
- Message delivery to offline users upon reconnection
- Support for public and private channels
- Scalable architecture with concurrency and synchronization
- Containerized with Docker for easy deployment to Kubernetes

## Technologies Used

The microservice utilizes the following technologies:

- Golang for backend development
- WebSocket protocol for real-time communication
- JWT for user authentication and authorization
- MongoDB for data persistence
- Docker for containerization
- Kubernetes for deployment and scalability

## Prerequisites

Before running the Amaranth application, make sure you have the following prerequisites installed:

- Docker
- Go (at least version 1.19)

## Getting Started

To get started with Amaranth, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/asorevs/amaranth.git
   ```

2. Change to the project directory:

   ```bash
   cd amaranth
   ```

3. Build the Docker image:

   ```bash
   docker build -t amaranth .
   ```

4. Run the Docker container:

   ```bash
   docker run -d -p 27017:27017 --name amaranth-container amaranth
   ```

   You can also replace "your-username" and "your-password" with your desired MongoDB username and password.

   ```bash
   docker run -d -p 27017:27017 --name amaranth-container -e MONGODB_USERNAME="your-username" -e MONGODB_PASSWORD="your-password" amaranth
   ```

5. Verify that the container is running:

   ```bash
   docker ps
   ```

   You should see the "amaranth-container" listed.

6. Access the MongoDB shell inside the container:

   ```bash
   docker exec -it amaranth-container bash
   ```

7. You can also interact with the preconfigured data in the "amaranth" database.

   ```bash
   mongosh
   ```

## Contributing

Contributions and feedback are welcome. If you encounter any issues or have suggestions for improvements, please feel free to open issues or submit pull requests.

## License

This microservice is open source and released under the [MIT License](LICENSE).
