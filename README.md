# 🚀 To-Do List (Distributed Edition)

[![Go](https://img.shields.io/badge/Go-1.21%2B-blue?logo=go)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18-blue?logo=react)](https://reactjs.org/)
[![MQTT](https://img.shields.io/badge/MQTT-Mosquitto-blue?logo=eclipse-mosquitto)](https://mosquitto.org/)
[![Docker](https://img.shields.io/badge/Docker-✓-blue?logo=docker)](https://www.docker.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/licenses/MIT)

A lightweight distributed To-Do List demonstrating microservices, REST API, and **publish-subscribe** messaging via MQTT (Mosquitto).  
Built for the *Distributed Programming for Web, IoT and Mobile Systems* exam.

---

## 🏛️ Architecture

This system is built on a microservice architecture. Components interact asynchronously via an MQTT message broker, ensuring loose coupling and high fault tolerance.

---

## 🎯 Basic Requirements  
- **Go 1.21+** (backend)  
- **Node.js 16+** (frontend)  
- **MQTT broker** (e.g. Mosquitto)  
- **Docker** (optional for broker)  
- **YAML** config file  
- **Git** & **GitHub Actions** for CI  

---

## 🛠️ Tech Stack  

| Layer               | Technology                               |
|---------------------|------------------------------------------|
| **Backend**         | Go, net/http                             |
| **Frontend**        | React (Vite)                             |
| **Messaging**       | MQTT (Mosquitto, Eclipse Paho)           |
| **Serialization**   | JSON (API) / YAML (config)               |
| **Storage**         | In-memory / PostgreSQL (optional)        |
| **Clocks & Sync**   | Lamport logical clocks                   |
| **Testing**         | Go testing & mocks                       |
| **CI/CD**           | GitHub Actions                           |

---

- **Server (API Gateway):** Accepts HTTP requests from the client, performs validation, and publishes task events to the MQTT broker.
- **Worker (optional):** Can subscribe to the MQTT topic and process events.
- **MQTT Broker:** Implements the **Publish-Subscribe** pattern for asynchronous messaging between components.

---

## 🚀 Quick Start (Docker)

This is the recommended way to run the project. Docker Compose will automatically bring up all components, including the MQTT broker.

1.  **Clone the repository** and navigate into it:
    ```bash
    git clone https://github.com/qoparu/tdl.git
    cd tdl
    ```
2.  **Run all containers** with a single command:
    ```bash
    docker-compose --profile client-server up --build
    ```
    This will start the Backend, Frontend, and Mosquitto MQTT broker.
    - Frontend: [http://localhost:3000](http://localhost:3000)
    - Backend: [http://localhost:8080](http://localhost:8080)

---

## 🛠️ Technologies and Concepts

This project demonstrates the following key concepts from the 'Distributed Programming' course:

* **Architecture:**
    * **Microservices:** The application is split into independent services (Server, optional Worker).
    * **Publish-Subscribe:** Implemented using MQTT for asynchronous communication.

* **Communication:**
    * **Message-Oriented Middleware:** Using MQTT (with the Eclipse Paho library in Go).
    * **REST API:** For client-server interaction.

* **Synchronization:**
    * **Lamport's Logical Clocks:** Partial ordering of events in a distributed system.

* **Data Serialization:**
    * **JSON:** For API requests and events.
    * **YAML:** For configuration files.

* **Containerization & CI/CD:**
    * **Docker & Docker Compose:** For environment isolation and orchestration.
    * **GitHub Actions:** For automatic testing on every commit.

---

## 🧪 Testing

To run backend tests:

```bash
go test ./...
