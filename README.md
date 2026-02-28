Distributed Rate Limiter + Analytics System

A horizontally scalable rate-limiting and analytics system inspired by architectures used by companies like Cloudflare and Amazon Web Services.

This project demonstrates:

Distributed state management

Token bucket rate limiting

Event-driven architecture

Kafka-based async processing

Redis-backed analytics

Horizontal scalability

Production-grade service separation

🏗 Architecture Overview
Client
   ↓
Node.js Gateway (Routing + Optional Cache)
   ↓
Go Rate Limiter Service
   ↓
Redis (Token State Storage)
   ↓
Kafka (Event Stream)
   ↓
Go Analytics Worker
   ↓
Redis (Aggregated Metrics)
🧠 System Design Decisions
1. Stateless Services

All services are stateless and horizontally scalable.

2. Fast Write Path

Rate limiting must remain low latency:

Redis handles token state

Kafka decouples analytics

3. Event-Driven Analytics

Analytics processing is:

Asynchronous

Eventually consistent

Horizontally scalable via Kafka consumer groups

📦 Project Structure
distributed-rate-limiter/
│
├── docker-compose.yml
├── .env
├── README.md
│
├── gateway/              # Node.js API Gateway
├── rate-limiter/         # Go rate limiter service
└── analytics-worker/     # Go Kafka consumer
⚙️ Components
🔵 Gateway (Node.js)

Responsibilities:

Accept client requests

Forward rate-check calls

Expose analytics endpoint

Runs on:

http://localhost:3000
🟢 Rate Limiter (Go)

Implements:

Token bucket algorithm

Redis-backed token state

Kafka event publishing

Endpoint:

POST /check
{
  "user_id": "123"
}

Response:

{
  "allowed": true
}

Runs on:

http://localhost:8080
🟣 Analytics Worker (Go)

Consumes Kafka events

Updates Redis counters

Uses consumer groups for scaling

🔴 Redis

Used for:

Token state

Analytics counters

Optional caching layer

🟡 Kafka

Used for:

Decoupling write path

Asynchronous analytics processing

Horizontal scaling via consumer groups

🚀 Running the System
1️⃣ Start Everything

From project root:

docker-compose up --build
2️⃣ Test Rate Limiting
curl -X POST http://localhost:3000/api \
  -H "Content-Type: application/json" \
  -d '{"user_id":"user1"}'

After exceeding limit:

HTTP 429
3️⃣ Check Analytics
curl http://localhost:3000/analytics/user1

Response example:

{
  "total": 15,
  "blocked": 5
}
📈 Scaling the System

You can scale services independently:

docker-compose up --scale rate-limiter=3
docker-compose up --scale analytics-worker=3

Kafka manages:

Partitioning

Consumer group rebalancing

Redis provides:

Shared token state across instances

⚠️ Failure Handling Strategy
Failure Scenario	Handling Strategy
Redis down	Configurable fail-open or fail-closed
Kafka down	Producer retries
Worker crash	Kafka consumer group rebalancing
Gateway overload	HTTP 429 responses
🧪 Benchmarking

Using wrk:

wrk -t4 -c200 -d30s http://localhost:3000/api

Expected:

Low latency (<20ms locally)

Stable throughput

No write blocking due to async analytics

🔧 Configuration

All configuration is handled via .env:

Rate limit parameters

Redis connection

Kafka broker

Ports

🧱 Tradeoffs
Why Redis?

O(1) operations

High throughput

Ideal for ephemeral token state

Why Kafka?

Decouples write path

Durable event log

Scalable consumer model

Why Go?

Lightweight concurrency via goroutines

High throughput

Efficient memory usage
