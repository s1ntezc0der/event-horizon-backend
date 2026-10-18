# Event Horizon backend

This is a reference cut for showcase. The services are Go, gRPC, Postgres, Redis, NATS.

For more info ask: s1ntezc0der@gmail.com

Game platform backend: accounts, scores, shop, inventory, billing, realtime leaderboard. Each service is its own Go module under Clean Architecture (`handler` → `service` → `repository`). The public edge is HTTP. Domain calls are gRPC. Events go through NATS JetStream.

This tree is the backend only.
Synced from the main tree on 11.10.2026. Module path in this cut stays `github.com/s1ntezc0der/event-horizon-backend`.
 No frontend, no compose file, no deploy scripts.

## Shape

```text
Client
  │ HTTP
  ▼
Balancer
  │
  ▼
Gateway          JWT, HTTP → gRPC
  │
  ├── Auth, Game, Billing, Leaderboard, Shop
  ├── Inventory, Profile, Payment, Authors, History
  ├── Analytics (ClickHouse)
  └── NATS Hub
        │
        ▼
   JetStream ──► Fulfillment, Notification, Leaderboard, Profile
```

Leaderboard keeps a Redis sorted set and pushes updates over a websocket. Shop purchases are written with an outbox, then published (`purchase.paid`). Postgres is per service. Redis holds sessions, cache, and the leaderboard. Inventory also uses Mongo for author catalog fields.

## Services

| Service | Role |
|---------|------|
| gateway | HTTP API, auth check, routes into gRPC |
| balancer | Front door in front of gateway replicas |
| auth | Register, login, JWT, sessions in Redis |
| game | Score submit |
| billing | In-game balances |
| leaderboard | Top scores, Redis sorted set, websocket |
| shop | Catalog and purchase |
| inventory | Items, Postgres plus Mongo |
| profile | Player profile |
| payment | Subscription and merch gate |
| authors | Creator profiles |
| history | Event trail from NATS |
| analytics | Aggregates in ClickHouse |
| fulfillment | Consumes paid purchases |
| notification | Consumes events, sends notices |
| nats-hub | JetStream stream bootstrap |
| mcp | Local stdio tools for agents, not on the player path |

Shared code next to `services/`: `pkg/` (migrator, redis, sql builder), `platform/` (closer and common helpers), `contracts/events`.

## Stack

Go, gRPC, Gin on the gateway, PostgreSQL, Redis, MongoDB, NATS JetStream, ClickHouse, JWT, bcrypt, OpenTelemetry.

