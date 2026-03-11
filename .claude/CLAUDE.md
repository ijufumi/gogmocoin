# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go wrapper library for the [GMO Coin API](https://api.coin.z.com/docs) — a Japanese cryptocurrency exchange. Module path: `github.com/ijufumi/gogmocoin/v2`. Requires Go 1.25+.

## Build & Test Commands

```bash
go build -v ./...          # Build all packages
go test ./...              # Run all tests
go test ./api/public/...   # Run tests for a specific package subtree
```

Linting uses `golangci-lint` (run via CI, install locally with `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`).

## Commit Convention

Uses **conventional commits** enforced by commitlint + husky. Commit messages must follow the format: `type: description` (e.g., `feat: add new endpoint`, `fix: correct symbol parsing`). Semantic-release auto-publishes GitHub releases from `main`.

## Architecture

### Package Layout

- **`api/internal/api/`** — Base classes shared by all API clients:
  - `RestAPIBase` — HTTP client with GET/POST/PUT/DELETE, HMAC-SHA256 signing for private endpoints
  - `WSAPIBase` — WebSocket client with auto-reconnect, dual goroutines for send/receive, `atomic.Value` state machine
- **`api/public/`** — Unauthenticated API clients (REST + WebSocket)
- **`api/private/`** — Authenticated API clients requiring API key/secret (REST + WebSocket)
- **`api/common/`** — Shared types, constants, configuration, and utilities

### Key Design Patterns

**Interface segregation + composition**: Each API endpoint is a small interface (e.g., `Ticker`, `OrderBooks`). Client structs compose multiple interfaces into a single `Client` type.

**Strategy via function injection**: `RestAPIBase` accepts `HostFactoryFunc` and `HeaderCreationFunc` — public clients pass no-auth functions, private clients inject HMAC signing logic.

**Generics for WebSocket streams**: `RetrieveStream[T]` provides type-safe unmarshaling of WebSocket messages.

**Custom value types**: `Symbol` (validated string enum), `TimeInMillis` (millisecond timestamp), `decimal.Decimal` for all price/volume fields.

### Configuration

Private API clients read `API_KEY` and `API_SECRET` from a `.env` file (loaded via godotenv). Copy `.env.example` to `.env` for local development.

### Response Models

All REST responses embed `model.ResponseCommon` which contains `Status`, `Messages[]`, and `ResponseTime`. Each endpoint has its own typed `Data` field.
