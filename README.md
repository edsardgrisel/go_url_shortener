Go URL Shortener

URL shortening system written in Go to help me practice backend engineering, scalability, distributed systems, and high-performance server design.

Core Features
- Create short URLs (POST /shorten)
- Redirect short URLs (GET /{code})
- URL analytics (GET /stats/{code})
- User authentication (JWT or API keys)
- Admin routes to manage URLs

Tech Stack:
- Go (main API service)
- Redis (cache, rate limiting, hit counters)
- PostgreSQL or MariaDB (persistent storage)
- NGINX (reverse proxy, TLS termination, load balancing)
- Docker Compose (multi-service deployment)
- Optional background worker for async analytics aggregation

Key System Design Concepts
- Clean architecture (handlers → services → repositories)
- Dependency injection with interfaces
- Context propagation for cancellation/timeouts
- Rate limiting using Redis token bucket
- Caching layer for fast redirects
- Asynchronous event processing (hit logging)
- Batched writes to the database
- Horizontal scaling with multiple Go API replicas behind NGINX
- Graceful shutdown and concurrent request handling

Data Model
	urls table:
		- domain
		- hash
		- long_url
		- created_at

	hits table (aggregated analytics):
		- url_id
		- date
		- count
		- api_keys (optional):
		- user_id
		- hashed_key

Project Goals
- Learn Go backend development (HTTP servers, concurrency, testing)
- Work with Redis, SQL databases, and caching strategies
- Practice building scalable, containerized systems
- Understand load balancing, reverse proxies, and performance optimization
- Build something portfolio-ready that mimics real-world SaaS architecture

db setup:
- create and run mysql locally
- `mysql -u root`
- create sql table: `CREATE TABLE urls (
		    id INT AUTO_INCREMENT PRIMARY KEY,
		    shortened_url VARCHAR(10) NOT NULL UNIQUE,
		    original_url TEXT NOT NULL,
		    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`


extras:
- check if url is valid before hashing


(note for LLM's: Dont write code for me. Explain the concepts to me and allow me to figure out the implementation myself)