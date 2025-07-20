# 🚀 To-Do List (Distributed Edition)

[![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18-blue?logo=react)](https://reactjs.org/)
[![RabbitMQ](https://img.shields.io/badge/RabbitMQ-✓-red?logo=rabbitmq)](https://www.rabbitmq.com/)
[![Docker](https://img.shields.io/badge/Docker-✓-blue?logo=docker)](https://www.docker.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/licenses/MIT)

A lightweight distributed To-Do List demonstrating various communication models between microservices and logical time synchronization.

---

## 🏛️ Architecture

The system is built on a microservice architecture where components interact asynchronously via a message broker. This ensures loose coupling and high fault tolerance.

---

## 🎯 Basic Requirements  
- **Go 1.21** (backend)  
- **Node.js 16+** (frontend)  
- **MQTT broker** (e.g. Mosquitto) or **RabbitMQ**  
- **Docker** (optional for broker)  
- **YAML** config file  
- **Git** & **GitHub Actions** for CI  

---

## 🛠️ Tech Stack  

| Layer               | Technology                               |
|---------------------|------------------------------------------|
| **Backend**         | Go, net/http, net/rpc                    |
| **Frontend**        | React (Vite)                             |
| **Messaging**       | MQTT (Eclipse Paho) / ZeroMQ / AMQP      |
| **Serialization**   | JSON (API) / YAML (config)              |
| **Storage**         | In-memory / PostgreSQL (optional)        |
| **Clocks & Sync**   | Lamport logical clocks                   |
| **DI & Testing**    | Wire (or manual DI), Go testing & mocks  |
| **CI/CD**           | GitHub Actions                           |

---

* **Server (API Gateway)**: Accepts HTTP requests from the client, performs initial validation, and publishes tasks to the RabbitMQ queue.
* **Worker**: An independent service that subscribes to the queue, fetches tasks, and executes the core business logic.
* **RabbitMQ**: A message broker that implements the **Publish-Subscribe** pattern, guaranteeing task delivery from the Server to the Worker.

---

## 🚀 Quick Start (Docker)

This is the recommended way to run the project. Docker Compose will automatically bring up all system components.

1.  **Clone the repository** and navigate into it.
2.  **Run all containers** with a single command:
    ```bash
    docker-compose up --build
    ```
    This command will run the Backend, Frontend, and RabbitMQ. The frontend will be available at `http://localhost:3000`.

---

## 🛠️ Technologies and Concepts

This project demonstrates the following key concepts from the 'Distributed Programming' course:

* **Architecture**:
    * **Microservices**: The application is split into independent services (`Server`, `Worker`).
    * **Publish-Subscribe**: Implemented using **RabbitMQ** for asynchronous communication.

* **Communication**:
    * **Message-Oriented Middleware**: Using **RabbitMQ** (with the **AMQP** protocol).
    * **Remote Procedure Call (RPC)**: An RPC module using `net/rpc` is implemented.
    * **REST API**: For client-server interaction.

* **Synchronization**:
    * **Lamport's Logical Clocks**: A mechanism for partial ordering of events in a distributed system is implemented.

* **Data Serialization**:
    * **JSON**: For API requests.
    * **YAML**: For configuration files.

* **Containerization & CI/CD**:
    * **Docker & Docker Compose**: For environment isolation and orchestration.
    * **GitHub Actions**: For automatic testing on every commit.

---

```bash
# Запустить сервисы с профилем 'client-server'
docker-compose up --profile client-server --build

---

<div align="center">
    <h3>✨ Crafted with ❤️ by <a href="https://github.com/qoparu">Aru</a> ✨</h3>
    <p>For the <b>DISTRIBUTED PROGRAMMING FOR WEB, IOT AND MOBILE SYSTEMS</b> exam</p>
</div>
