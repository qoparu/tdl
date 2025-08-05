# 🚀 To-Do List (Distributed Edition)

[![Go](https://img.shields.io/badge/Go-1.21%2B-blue?logo=go)](https://golang.org/)
[![JavaScript](https://img.shields.io/badge/JavaScript-ES6%2B-yellow?logo=javascript)](https://www.javascript.com/)
[![MQTT](https://img.shields.io/badge/MQTT-Mosquitto-blue?logo=eclipse-mosquitto)](https://mosquitto.org/)
[![Docker](https://img.shields.io/badge/Docker-✓-blue?logo=docker)](https://www.docker.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/licenses/MIT)

A lightweight distributed To-Do List demonstrating microservices, REST API, and **publish-subscribe** messaging via MQTT (Mosquitto).  
Built for the *Distributed Programming for Web, IoT and Mobile Systems* exam.


<p align="center">
  <img src="TDL.png" alt="img" />
</p>

---

## 🏛️ Architecture

This system is built on a microservice architecture. Components interact asynchronously via an MQTT message broker, ensuring loose coupling and high fault tolerance.


<p align="center">
  <img src="docker.png" alt="img" />
</p>

---

## 🎯 Basic Requirements  
- **Go 1.21+** (backend)  
- **JS** (frontend)  
- **MQTT broker** (e.g. Mosquitto)  
- **Docker** (optional for broker)  
- **YAML** config file  
- **Git** & **GitHub Actions** for CI  

---

## 🛠️ Tech Stack  

| Layer               | Technology                               |
|---------------------|------------------------------------------|
| **Backend**         | Go, net/http                             |
| **Frontend**        | Vanilla JS, HTML5, CSS3                  |
| **Messaging**       | MQTT (Mosquitto, Eclipse Paho)           |
| **Serialization**   | JSON (API) / YAML (config)               |
| **Storage**         | In-memory / PostgreSQL (optional)        |
| **Testing**         | Go testing & mocks                       |
| **CI/CD**           | GitHub Actions                           |

---

- **Server (API Gateway):** Accepts HTTP requests from the client, performs validation, and publishes task events to the MQTT broker.
- **Worker (optional):** Can subscribe to the MQTT topic and process events.
- **MQTT Broker:** Implements the **Publish-Subscribe** pattern for asynchronous messaging between components.

---

## 🚀 Quick Start (Docker Compose)

This is the recommended way to run the project. Docker Compose will automatically build and run all services: the Go backend, the database, the MQTT broker, and the frontend.

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/qoparu/tdl.git](https://github.com/qoparu/tdl.git)
    cd tdl
    ```

2.  **Run the entire stack:**
    ```bash
    docker-compose --profile client-server up --build
    ```

3.  **Access the application:**
    * **Frontend:** [http://localhost:5173](http://localhost:5173)
    * **Backend API:** [http://localhost:8080](http://localhost:8080) / [http://localhost:8080/tasks](http://localhost:8080/tasks)
      
---

## 🛠️ Technologies and Concepts

This project demonstrates the following key concepts from the 'Distributed Programming' course:

* **Architecture:**
    * **Microservices:** The application is split into independent services (Server, optional Worker).
    * **Publish-Subscribe:** Implemented using MQTT for asynchronous communication.

* **Communication:**
    * **Message-Oriented Middleware:** Using MQTT (with the Eclipse Paho library in Go).
    * **REST API:** For client-server interaction.

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
```


<p align="right">
  <img src="db.jpg" alt="img" />
</p>

<p align="left">
  <img src="p1.jpg" alt="img" />
</p>

<div align="center"> <h3>✨ Crafted with ❤️ by <a href="https://github.com/qoparu">Aru</a> ✨</h3> <p>For the <b>DISTRIBUTED PROGRAMMING FOR WEB, IOT AND MOBILE SYSTEMS</b> exam</p> </div>
