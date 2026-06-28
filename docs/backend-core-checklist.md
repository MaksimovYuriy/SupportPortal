# Backend Core Checklist

## Backend Core

- [x] Basic models: `User`, `Agent`, `Queue`, `Flow`, `FlowStep`, `Ticket`
- [x] Repositories for core entities
- [x] Services for core entities
- [x] REST handlers and routes
- [x] Ticket lifecycle: `new -> in_queue -> in_progress -> closed`
- [x] Engine v1: moves tickets into routes and assigns agents
- [x] Basic ticket filters
- [x] HTTP and Engine logging
- [x] Basic error handling
- [x] Unit tests for services and Engine

## Auth / Security

- [ ] New database migrations are allowed when auth/security needs additional fields or tables
- [ ] Add password hashing service, for example `internal/security/password`
- [ ] Use `bcrypt`
- [ ] Accept a raw password when creating a user through `POST /users`
- [ ] Store only `password_hash`
- [ ] Never return `password_hash` in DTOs
- [ ] Add `AuthService`
- [ ] Add login by email/password
- [ ] Add JWT service
- [ ] Add `POST /auth/login`
- [ ] Add auth middleware
- [ ] Protect agent/admin endpoints with auth middleware
- [ ] Keep only required endpoints public, for example `POST /auth/login` and maybe bot ticket creation
- [ ] Add tests for hashing, login, and middleware

## API / Handlers

- [ ] Add handler tests for `POST /tickets`
- [ ] Add handler tests for `POST /tickets/{id}/complete`
- [ ] Add handler tests for auth errors: `401`, `403`, invalid token
- [ ] Check that all handlers return errors in one format

## Engine

- [ ] Add per-tick summary log: started, assigned, skipped, failed
- [ ] Consider worker pool / goroutines for processing ticket batches
- [ ] Keep the current simple Engine v1 as a stable base
- [ ] Add tests for concurrent scenarios if concurrency is introduced

## Errors / Validation

- [ ] Add more specific errors when useful:
  - `ErrUnauthorized`
  - `ErrForbidden`
  - `ErrInvalidCredentials`
  - `ErrInvalidTicketState`
  - `ErrAgentUnavailable`
  - `ErrFlowHasNoSteps`
- [ ] Check which errors should map to `400/401/403/404/409`
- [ ] Check which errors should be logged as internal errors

## Database / Repositories

- [ ] Add new migrations when changing schema; avoid editing old migrations after schema stabilizes
- [ ] Consider indexes:
  - `tickets.status`
  - `tickets.current_flow_step_id`
  - `tickets.assigned_agent_id`
  - `users.email`
- [ ] Add integration tests at least for `TicketRepository`
- [ ] Later add integration tests for `UserRepository` and `FlowRepository`

## Docs

- [ ] Write `README.md`
- [ ] Document project startup
- [ ] Document env variables
- [ ] Document migrations
- [ ] Document ticket lifecycle
- [ ] Document Engine
- [ ] Add API examples / Postman flow

## Next Stage

- [ ] Design Python Telegram bot
- [ ] Design RabbitMQ events
- [ ] Add `internal/messaging`
- [ ] Document message contracts
- [ ] Connect Telegram ticket creation with backend core

## Suggested Order

1. Password hashing
2. User creation with `password_hash`
3. Auth login / JWT
4. Auth middleware
