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

+----------+      +----------------+      +------------+      +----------+
|          |----->|                |----->|            |----->|          |
|  React   |      |  Server (Go)   |      |  RabbitMQ  |      | Worker   |
| (Client) |      |   (API Gateway)|      |  (Broker)  |      |  (Go)    |
|          |<-----|                |<-----|            |<-----|          |
+----------+      +----------------+      +------------+      +----------+
---

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

## ⚡ Quick Start  

1. **Clone**
1. **Run the server**
   ```bash
   go run ./cmd/server -config config.yaml
   ```

<div align="center"> <h3>✨ Crafted with ❤️ by <a href="https://github.com/qoparu">Aru</a> ✨</h3> <p>For the <b>DISTRIBUTED PROGRAMMING FOR WEB, IOT AND MOBILE SYSTEMS</b> exam</p>
