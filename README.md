
# Track-and-Trace System for FMCG Products

This repository contains the backend mock code for an **end-to-end track-and-trace system** prototype,
designed to monitor and authenticate FMCG (Fast-Moving Consumer Goods) products across the global supply chain. 
The system enables real-time tracking, verification of product authenticity, and regulatory compliance, 
serving both B2B and B2C requirements.


## Table of Contents

1. [Project Overview](#project-overview)
2. [System Architecture](#system-architecture)
3. [Key Features](#key-features)
4. [Tech Stack](#tech-stack)
5. [Getting Started](#getting-started)
6. [Configuration](#configuration)
7. [Usage](#usage)
8. [Contributing](#contributing)
9. [License](#license)

---

## Project Overview

The track-and-trace system provides end-to-end visibility and verification for FMCG products 
through each stage of the supply chain. It captures real-time data from IoT devices, sensors, 
and partner APIs to monitor product location, condition, and authenticity. 

The system serves two main components:
- **B2B Service**: Enables supply chain partners to view real-time product status and track logistics.
- **B2C Service**: Allows end-users to verify product authenticity, offering proof of origin 
and journey details for quality assurance.


## System Architecture

The system architecture is divided into several key components:

1. **Data Ingestion Layer**:
    - Collects data from IoT devices and external sources, validates, normalizes, 
      and enriches it before sending it downstream.
2. **Real-Time Data Store**:
    - A NoSQL store (e.g., Cassandra, DynamoDB) optimized for low-latency access, 
      holding recent data for operational tracking.
3. **Event Store**:
    - Distributed log (e.g., Kafka, Apache Pulsar) that stores all events in a replayable format, 
      supporting stream processing and data continuity.
4. **Cold Storage and Archival**:
    - Long-term storage for historical data, leveraging cloud storage solutions (e.g., Amazon S3) 
      with lifecycle management for cost efficiency.
5. **Microservices**:
    - Modular services for handling data ingestion, transformation, analytics, and B2B/B2C API endpoints.
6. **APIs**:
    - RESTful endpoints that provide access to real-time product tracking and authenticity verification.

### Data Flow Diagram

```
IoT Devices & APIs  ➔  Ingestion Layer  ➔  Event Store ➔ Real-Time Data Store ➔ Cold Storage
                                     ↘      ↘              ↙
                                        Microservices & Analytics
```

## Key Features

- **Real-Time Tracking**: Monitor products at each checkpoint in the supply chain.
- **Proof of Authenticity**: Allows end-users to verify product authenticity and origin.
- **Scalable Storage**: Data is tiered across Real-Time, Event, & Cold storage layers for optimized performance & cost.
- **Replayability**: Distributed event log enables data recovery and system reliability.
- **Modular Microservices**: Enables efficient scaling, fault tolerance, and isolated updates.

## Tech Stack

- **Language**: Go, Python (for auxiliary scripts)
- **Data Storage**: Cassandra/DynamoDB (Real-Time), Kafka/Pulsar (Event Store), Amazon S3 (Cold Storage)
- **Message Broker**: Kafka or Apache Pulsar
- **API Gateway**: Kong or NGINX
- **Monitoring & Analytics**: Prometheus, Grafana
- **Containerization & Orchestration**: Docker, Kubernetes
- **CI/CD**: GitHub Actions, Jenkins

## Getting Started

### Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/dl/) (v1.18 or later)
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Kafka](https://kafka.apache.org/quickstart) or [Apache Pulsar](https://pulsar.apache.org/docs/en/standalone/) 
  (for event streaming)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/username/track-and-trace-system.git
   cd track-and-trace-system
   ```

2. **Build the Project**:
   ```bash
   go build -o track-and-trace ./cmd
   ```

3. **Set Up Docker**:
   ```bash
   docker-compose up -d
   ```

4. **Run Migrations**:
   Set up your database schemas and seed data:
   ```bash
   go run scripts/migrate.go
   ```

### Configuration

Edit the `config.yml` file to set database connection strings, API keys, and other necessary configuration options. 
Example configuration parameters include:

```yaml
database:
  real_time_store: "cassandra://username:password@hostname:port/keyspace"
  event_store: "kafka://localhost:9092"
  cold_storage: "s3://bucket-name"

api:
  b2b_endpoint: "https://api.example.com/b2b"
  b2c_endpoint: "https://api.example.com/b2c"
```

## Usage

### Running the Application

To start the application:
```bash
./track-and-trace
```

### API Endpoints

#### B2B API
- **GET** `/api/v1/track/{product_id}` - Retrieve the current location and condition of a product.
- **POST** `/api/v1/report` - Submit an event related to product movement.

#### B2C API
- **GET** `/api/v1/authenticate/{product_id}` - Verify the authenticity of a product based on its identifier.

### Testing

Run tests with:
```bash
go test ./...
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any feature requests or bug reports.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---