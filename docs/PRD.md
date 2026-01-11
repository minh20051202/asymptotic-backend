# Product Requirements Document

**Project:** Distributed Ticket Reservation System
**Version:** 1.0

## 1. Problem Statement

During high-demand "Flash Sales," the current system fails to guarantee inventory consistency. This leads to **Overselling** (selling 105 tickets for 100 seats), causing customer backlash and operational liability. Additionally, high traffic causes the database to crash, resulting in total service denial.

## 2. Actors

- **Admin:** Internal staff who configures the Event, sets the Total Quota (Inventory), and defines purchase limits.
- **User (Buyer):** A customer attempting to purchase tickets during the sale.

## 3. Functional Requirements

### Admin

- Must be able to create an Event with a fixed `total_quota` (e.g., 100).
- Must be able to set a `max_purchase_limit` (e.g., ).

### User (Buyer)

- Must be able to view the `available_quantity` of an Event.
- Must be able to purchase a specific `quantity` of tickets.
- **Constraint:** Requested quantity must not exceed `max_purchase_limit`.

### System (Core Logic)

- **Zero Overselling:** Must enforce `available_quantity >= 0` at all times.
- **Feedback:** Must respond with HTTP `409 Conflict` if the Event is Sold Out.
- **Performance:** Must handle concurrent requests without data corruption.

## 4. Non-Functional Requirements (NFRs)

### Consistency (Crucial)

- **Strict Consistency:** The system acts as a **CP System** (Consistency over Availability). It is acceptable to reject requests (Error 500/429) rather than allow Overselling.
- **Atomicity:** Purchase transactions must be atomic; either the Inventory decrements AND the Ticket is issued, or neither happens.

### Performance

- **Throughput:** The system must handle **1,000 write-requests per second** during peak traffic.
- **Latency:** 95% of requests (p95) must receive a response within **500ms**.

### Scalability

- The API layer must be stateless and horizontally scalable (capable of running on 3+ replicas behind a Load Balancer).

## 5. User Stories

### Admin

1.  "As an Admin, I want to set a fixed `total_quota` for an Event, so that I can strictly limit the number of attendees."
2.  "As an Admin, I want to define a `max_purchase_limit` per request, so that I can prevent scalpers from hoarding tickets."

### User

3.  "As a User, I want to view the real-time `available_quantity`, so that I know if tickets are still in stock before attempting to buy."
4.  "As a User, I want to purchase multiple tickets (up to the limit) in one transaction, so that I can book seats for my group/family."
