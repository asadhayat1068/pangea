# Pangea — Marketplace Service

Multi-tenant commerce platform. Release 1: Marketplace MVP for tenant FreshCart.

## Quick start (one command)
    make up      # starts PostgreSQL via Docker
    make run     # runs the marketplace service
    make test    # runs all tests
    make down    # stops PostgreSQL

## Architecture
Modular monolith, hexagonal (ports & adapters). See docs/adr for decisions.
Every entity is scoped by tenant_id (ADR-002).

## Migrations Setup

### macOS:
    brew install golang-migrate

### Linux (prebuilt binary):
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/migrate

### Any OS with Go (alternative):
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest