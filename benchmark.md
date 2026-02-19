---

# ⚡ Performance Benchmark: Optimistic vs. Pessimistic Locking

This section documents the performance characteristics of two concurrency control strategies under a **High Contention "Flash Sale" Scenario**.

### 🧪 The Test Scenario

We simulated a massive traffic spike where demand exceeds supply, forcing high contention on database rows.

- **Inventory:** 10,000 Tickets
- **Traffic:** 15,000 Concurrent Users (VUs) trying to buy simultaneously.
- **Goal:** Sell out inventory without overselling (Data Integrity) and maximize throughput.

---

### 📊 Summary of Results

| Metric          | 🛡️ Pessimistic Locking (Winner) | 🤞 Optimistic Locking     | Difference        |
| --------------- | ------------------------------- | ------------------------- | ----------------- |
| **Throughput**  | **476.85 reqs/s**               | 414.54 reqs/s             | **+15% Faster**   |
| **Avg Latency** | **17.79s**                      | 23.34s                    | **5.5s Faster**   |
| **P95 Latency** | **30.13s**                      | 35.09s                    | **More Stable**   |
| **Failures**    | 0 System Errors                 | **5 System Errors (503)** | **More Reliable** |
| **Mechanism**   | Queueing (DB Lock)              | Retry Loop (App Logic)    |                   |

> **Verdict:** Under extreme contention (1.5 users per 1 item), **Pessimistic Locking** outperformed Optimistic Locking in throughput, latency, and stability.

---

### 🧠 Analysis: Why did Optimistic Locking lose?

In low-contention scenarios (e.g., 200 items, 250 users), Optimistic Locking usually wins. However, at **15,000 concurrent users**, it suffered from a **"Retry Storm."**

#### 1. The Retry Storm (Optimistic Failure Mode)

When 15,000 users try to update the same rows:

1. **Conflict:** ~14,900 users fail to update the version on their first try.
2. **Retry:** The application logic catches the failure, waits (jitter), and sends the request _again_.
3. **Amplification:** The database is not just handling 15,000 requests; it is handling **15,000 × N Retries**.
4. **Result:** The CPU spends more time processing failures and retries than successful transactions. This is why latency spiked to **23s** and we saw `503 Service Unavailable` errors.

#### 2. The Efficient Queue (Pessimistic Success Mode)

Pessimistic locking used `SELECT ... FOR UPDATE` to create a "single-file line" at the database level.

1. **Queueing:** Users waited their turn inside the database connection pool.
2. **Efficiency:** Once a user got the lock, they bought the ticket 100% of the time. There was **zero wasted work** on failed attempts.
3. **Result:** Although queueing adds latency, it is mathematically more efficient than processing thousands of doomed retry attempts.

---

### 📄 Raw Benchmark Logs

<details>
<summary><strong>Click to view Pessimistic Locking Logs (Winner)</strong></summary>

```txt
  █ TOTAL RESULTS

    checks_total.......: 15000  476.856027/s
    checks_succeeded...: 66.66% 10000 out of 15000
    checks_failed......: 33.33% 5000 out of 15000

    CUSTOM
    errors_sold_out................: 5000   158.952009/s

    HTTP
    http_req_duration..............: avg=17.79s min=71.54ms  med=18.38s max=30.77s p(95)=30.13s
      { expected_response:true }...: avg=12.18s min=71.54ms  med=11.85s max=29.89s p(95)=24.69s
    http_req_failed................: 33.33% 5000 out of 15000
    http_reqs......................: 15000  476.856027/s

```

</details>

<details>
<summary><strong>Click to view Optimistic Locking Logs (Degraded)</strong></summary>

```txt
WARN[0002] Could not get a VU from the buffer for 400ms
ERRO[0032] Unexpected Error (503): {"error":"system busy, please try again"}

  █ TOTAL RESULTS

    checks_total.......: 15000  414.540345/s
    checks_succeeded...: 66.66% 10000 out of 15000
    checks_failed......: 33.33% 5000 out of 15000

    CUSTOM
    errors_other...................: 5      0.13818/s (System Overload)
    errors_sold_out................: 4995   138.041935/s

    HTTP
    http_req_duration..............: avg=23.34s min=67.29ms  med=27.51s max=35.38s p(95)=35.09s
      { expected_response:true }...: avg=17.96s min=67.29ms  med=17.52s max=34.88s p(95)=32.6s
    http_req_failed................: 33.33% 5000 out of 15000
    http_reqs......................: 15000  414.540345/s

```

</details>

---

### 🚀 Conclusion

For our Flash Sale system:

- We use **Pessimistic Locking** (`FOR UPDATE`) for inventory management because the cost of "Retry Storms" is too high when contention is extreme.
- We use **Optimistic Locking** (`version`) for user settings and shopping carts where contention is low.
