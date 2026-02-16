---

# Product Requirements Document

**Project:** Asymptotic (Universal API Gateway)
**Version:** 1.0
**Tagline:** The Financial Firewall for AI Agents.

## 1. Problem Statement

AI Agents operate autonomously and at high speed, needing access to dozens of paid APIs (Search, LLMs, Flight Data).

1. **Management Friction:** Developers must manage multiple API keys and credit cards, creating a maintenance nightmare.
2. **Financial Risk:** High-concurrency agent loops can accidentally trigger "Double-Spend" events, where an agent with $10.00 balance successfully fires $50.00 worth of API calls before the balance updates, leaving the platform with bad debt (Negative Balance).

## 2. Actors

- **AI Agent (Client):** The autonomous software sending HTTP requests to the Gateway.
- **External Provider:** The third-party service being consumed (e.g., Google Search, OpenAI, Bloomberg).
- **System (Asymptotic):** The reverse proxy responsible for routing, metering, and enforcing financial invariants.

## 3. Functional Requirements

### 3.1. The Gateway (Reverse Proxy)

- **REQ-1:** The System MUST act as a passthrough proxy. It receives requests at `POST /v1/proxy/{provider}/{service}` and forwards them to the corresponding External Provider.
- **REQ-2:** The System MUST hide the upstream API Keys. The Agent only provides one `Asymptotic-Key`.
- **REQ-3:** The System MUST support **Idempotency**. If an Agent sends the header `X-Idempotency-Key`, the System must return a cached response if the key was already processed.

### 3.2. Financial Logic (The Core)

- **REQ-4 (The Invariant):** The System MUST enforce a **Zero-Balance Invariant**. An Agent’s wallet balance MUST NEVER fall below $0.00.
- **REQ-5:** The System MUST deduct the estimated cost of the API call **before** forwarding the request to the External Provider.
- **REQ-6 (Refunds):** If the External Provider returns a server error (`5xx`) or times out, the System MUST automatically refund the deducted amount to the Agent's wallet.

### 3.3. Audit & Observability

- **REQ-7:** The System MUST log every transaction (Request ID, Cost, Status, Timestamp) to a persistent database for auditing.
- **REQ-8:** Logging MUST be performed asynchronously (via Worker Pool) to prevent adding latency to the Agent's request.

## 4. Non-Functional Requirements (NFRs)

### 4.1. Consistency (Crucial)

- **Strict Consistency:** The financial ledger acts as a **CP System** (Consistency over Availability).
- **Concurrency Control:** The System MUST use **Pessimistic Locking** (`SELECT FOR UPDATE`) on the Wallet row to prevent Race Conditions (Double-Spending) during simultaneous requests.
- **Isolation:** Financial transactions MUST run with an isolation level sufficient to prevent "Lost Updates" (Read Committed with Locking).

### 4.2. Performance

- **Throughput:** The System must handle **1,000 requests per second (RPS)** per wallet without locking the database for more than 10ms per request.
- **Latency Overhead:** The Gateway MUST NOT add more than **20ms** of overhead to the upstream API call time.
- **Optimization:** The "Split Transaction" pattern MUST be used: The database lock must be released _before_ the slow network call to the External Provider occurs.

### 4.3. Reliability

- **Fault Tolerance:** Network failures between the Gateway and the Agent (Lost ACKs) MUST NOT result in double-charging (handled via Idempotency).
- **Graceful Degradation:** If the Logging System (Worker Pool) is overloaded, the System SHOULD drop logs rather than block the API traffic (Availability over Observability for logs).

## 5. User Stories

### AI Agent (The Client)

1. "As an Agent, I want to use a single API Key to access Google, OpenAI, and Weather services, so that I don't crash due to missing credentials."
2. "As an Agent, I want to send an `Idempotency-Key` when I retry a request, so that I don't get charged twice if my WiFi disconnects."

### Platform Admin

3. "As an Admin, I want to ensure that even if an Agent sends 50 parallel requests for $1.00 each, and they only have $10.00, only 10 requests succeed and 40 fail with `402 Payment Required`."
4. "As an Admin, I want the system to automatically refund the Agent if the upstream provider (e.g., Google) is down (Error 500), so that users only pay for successful calls."
