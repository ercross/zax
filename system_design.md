### 1. **Objective**
Zax aims to build a system that:
- Provides real-time visibility and traceability for FMCG products from production to delivery across the global supply chain
- Ensure data availability and transparency for B2B stakeholders
- Provide B2C proof of authenticity and proof or origin

### 1. **System Overview**
The track and trace system enables:
- **B2B Track and Trace**: Provides real-time tracking of goods for stakeholders across the supply chain (manufacturers, distributors, warehouses, retailers).
- **B2C Authentication**: Allows end-users to verify product authenticity and origin, enhancing transparency and trust in the product’s quality.

### 2. **Functional Requirements**
- **Tracker-to-Product Mapping**: One-to-one mapping of Tracker chip data to product batch data
- **Hierarchical Arrangement of mapped data**: Mapped data are then arranged hierarchically as detailed in GS1 official specification for Track and Trace systems
- **Real-Time Tracking**: The last 2 points provides a continuous, real-time tracking and updates on product location and status.
    - **Receive Product-ownership status Data**: As products is transferred from the key actors of the supply chain, receive product exchange data
    - **Data Ingestion from iOT devices**: Collect data from various IoT devices, such as GPS trackers, and logistics systems regarding product location and movement to provide real-time tracking of product.

- **Product Authentication**: Enable consumers to scan product codes (e.g., QR codes) to verify authenticity and origin.
- **Notifications and Alerts**: Send notifications and alerts for anomalies (e.g., temperature deviation, unauthorized movement).
- **Reporting**: Generate reports on product movement, storage conditions, and delivery times.
- **Access Control**: Ensure secure access for different roles (admin, distributor, retailer, etc.).

### 3. **Non-Functional Requirements**
- **Scalability**: Handle high data throughput from potentially millions of FMCG products and multiple stakeholders.
- **Data Consistency**: Ensure eventual consistency for real-time data across services.
- **Latency**: Low latency for both B2B tracking and B2C authentication.
- **Fault Tolerance**: Redundancy to prevent data loss and maintain high availability.
- **Security**: Data protection and integrity, particularly in B2C components.

### 4. **System Architecture Components**
The architecture follows a **microservices** and **event-driven** approach. Below is a breakdown of key components:

#### 4.1 **Data Ingestion Layer**
- **IoT Gateways**: Aggregates data from GPS, RFID, and other IoT sensors. Each product package is tagged with an IoT device for monitoring location and environmental conditions (e.g., temperature).
- **API Gateway**: Manages incoming API requests and routes them to respective microservices. It provides authentication, rate limiting, and load balancing.

#### 4.2 **Core Microservices**
1. **Tracking Service**:
    - Responsible for managing the real-time location of products, updating positions as goods move through various supply chain stages.
    - Uses **Kafka** or **NATS** for asynchronous event streaming to push tracking updates to downstream services.

2. **Authentication Service**:
    - Provides a B2C interface that verifies product authenticity based on a unique identifier (e.g., QR code).

3. **Event Processing Service**:
    - Consumes raw data from the IoT gateway and applies business rules (e.g., temperature anomaly detection).
    - Aggregates data for storage and passes relevant tracking events to downstream services.

4. **Notification Service**:
    - Sends alerts and notifications to stakeholders based on pre-configured rules.
    - Integrates with SMS, email, or push notification services to provide real-time updates.

5. **Reporting and Analytics Service**:
    - Generates on-demand and scheduled reports on supply chain metrics (e.g., delivery times, storage conditions).
    - Connects to a data warehouse for analytical querying.

6. **Blockchain Integration Service**:
    - Records essential data on an immutable ledger to track product origin and authenticity.
    - Interacts with Hyperledger or similar blockchain for B2C transparency, especially when queried by end-users.

#### 4.3 **Storage Solutions**
- **Real-Time Data Store**: Uses a NoSQL database (e.g., Cassandra, DynamoDB) for low-latency real-time data handling.
- **Event Store**: Stores all incoming and processed events using a distributed event store (e.g., Kafka with compacted topics or Apache Pulsar).
- **Cold Storage and Archival**: For historical tracking data, uses a cold storage solution (e.g., Amazon S3) with data lifecycle management.

#### 4.4 **B2C Frontend**
- Provides consumers with a user-friendly interface to check product authenticity by scanning a QR code.
- Integrates with the Authentication Service, which retrieves data from the blockchain and verifies authenticity.
- Displays basic information on product origin, manufacturing details, and any additional authenticity certificates.

#### 4.5 **B2B Interface**
- **Web and API Access**: Web portal and RESTful APIs for supply chain partners to access tracking, reporting, and alert functionalities.
- **Access Control**: Role-based access control (RBAC) ensures stakeholders can access only their data.

### 5. **Detailed Flow: B2B and B2C Use Cases**
#### B2B Track and Trace
1. **IoT Data Ingestion**: Data from GPS trackers and RFID sensors are continuously sent to the IoT Gateway.
2. **Data Processing**: The Event Processing Service receives data, processes it, and applies any business rules (e.g., check for anomalies).
3. **Tracking Update**: Tracking Service updates the current status in the real-time data store and publishes events via Kafka for downstream processing.
4. **Notification**: Notification Service triggers alerts based on configured rules.
5. **Data Aggregation and Reporting**: Data is aggregated in the Reporting and Analytics Service for supply chain insights.

#### B2C Authentication
1. **Consumer Scans QR Code**: The consumer scans a code on the product, triggering an API call to the Authentication Service.
2. **Blockchain Query**: The service retrieves information from the blockchain (stored by Blockchain Integration Service) to verify product details.
3. **Display Authentication Results**: Authentication Service sends a response back to the consumer, showing the product’s origin and authenticity status.

### 6. **System Design Patterns and Key Decisions**
- **Event-Driven Architecture**: Kafka or NATS is essential for decoupling services and ensuring asynchronous data processing for real-time updates.
- **CQRS (Command Query Responsibility Segregation)**: The separation of tracking data (commands) from reporting data (queries) ensures scalability and optimizes each service independently.
- **Distributed Tracing**: Using tools like OpenTelemetry for tracing requests across microservices to track delays and bottlenecks.
- **Caching**: Redis or Memcached caches frequently accessed data, reducing load on the real-time database, particularly for high-frequency queries from the B2C authentication service.

### 7. **Security and Compliance Considerations**
- **Data Encryption**: Encrypt sensitive data both at rest and in transit to ensure security.
- **Auditing and Logging**: Keep detailed logs of all user actions and events in the system for auditing and regulatory compliance.
- **Access Controls**: Enforce strong access controls across B2B and B2C components to protect against unauthorized access.

### 8. **Scalability and Reliability**
- **Horizontal Scaling**: Each service can be scaled independently based on demand, ensuring the system can handle high volumes of IoT data and B2C requests.
- **Failover and Disaster Recovery**: Maintain multiple instances of key services and implement disaster recovery protocols for critical data (e.g., using cross-region replication in cloud storage).

### 9. **Conclusion**
This track-and-trace system is designed to handle the complexities of global FMCG supply chains by providing both B2B and B2C capabilities. With real-time data processing, strong security measures, and a scalable microservices architecture, it can support both internal stakeholders and end consumers in ensuring product authenticity and transparency across the supply chain. This design demonstrates a deep consideration of supply chain challenges and leverages modern architectural principles to deliver a robust solution.
