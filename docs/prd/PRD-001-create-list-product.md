# PRD-001 — Create & List a Product

## Problem
FreshCart has no system of record for its products. We will give merchants one.

## Target user
FreshCart merchant (Dana, Marco).

## In-scope user story (exactly one)
As a FreshCart merchant, I can create a product so customers can later buy it.

## Success metric
A merchant creates a valid product in under 60 seconds and immediately sees it
in the product list.

## Acceptance criteria (handed to the BA to sharpen, Step 6)
- A product can be created with: name, price, category (stock flag optional,
  defaults to "out of stock").
- Invalid products are rejected with a clear, specific error.
- Every product is permanently associated with a tenant (`tenant_id`).
- Created products can be listed back, scoped to the tenant.

## Explicitly out of scope for this PRD
Cart, checkout, payments, images, search, customer accounts.

## Dependencies
None — this is the first feature. It depends only on the Sprint-1 skeleton.