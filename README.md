# 🚀 DistDoDo

[![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)](https://golang.org/)  
[![React](https://img.shields.io/badge/React-18-blue?logo=react)](https://reactjs.org/)  
[![MQTT](https://img.shields.io/badge/MQTT-✓-orange?logo=mqtt)](https://mqtt.org/)  
[![ZeroMQ](https://img.shields.io/badge/ZeroMQ-✓-purple?logo=zeromq)](https://zeromq.org/)  
[![RabbitMQ](https://img.shields.io/badge/RabbitMQ-✓-red?logo=rabbitmq)](https://www.rabbitmq.com/)  
[![net/rpc](https://img.shields.io/badge/net--rpc-✓-green)](https://pkg.go.dev/net/rpc)  
[![YAML](https://img.shields.io/badge/YAML-✓-yellow?logo=yaml)](https://yaml.org/)  
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/licenses/MIT)  

Lightweight distributed to-do list with real-time push notifications and pluggable message brokers.

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
   ```bash
   git clone https://github.com/your-org/distdodo.git
   cd distdodo

<div align="center"> <h3>✨ Crafted with ❤️ by <a href="https://github.com/qoparu">Aru</a> ✨</h3> <p>For the <b>Automated Software Testing</b> exam</p> <img src="https://img.shields.io/badge/Java-Expert-important?logo=java" alt="Java Expert"> </div> 
