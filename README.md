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