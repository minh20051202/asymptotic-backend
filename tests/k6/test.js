import http from "k6/http";
import { check } from "k6";
import { Counter } from "k6/metrics";

// Custom metrics to track specific errors
const soldOutErrors = new Counter("errors_sold_out");
const dbConnectionErrors = new Counter("errors_db_connection");
const otherErrors = new Counter("errors_other");

// ------------------------------------------------------------------
// CONFIGURATION
// ------------------------------------------------------------------
const EVENT_ID = "a97de01f-caea-40be-ba2b-4d39e75fbb1d";
const USER_IDS = [
  "f333ce48-6877-4967-9b12-7e62c3f9acc3",
  "e67cedc9-2c20-4768-88a1-954929c766a4",
  "60f88754-9b3b-4b58-a723-713cf6698fff",
];

export const options = {
  scenarios: {
    flash_sale: {
      executor: "per-vu-iterations",
      vus: 15000,
      iterations: 1,
    },
  },
};

export default function () {
  const url = "http://localhost:8080/transaction";
  const randomUser = USER_IDS[Math.floor(Math.random() * USER_IDS.length)];

  const payload = JSON.stringify({
    walletId: EVENT_ID,
    userId: randomUser,
    amount: 1,
  });

  const params = {
    headers: { "Content-Type": "application/json" },
  };

  // 1. Capture the response
  const res = http.post(url, payload, params);

  // 2. Check the Status Code
  const isSuccess = check(res, {
    "status is 200": (r) => r.status === 200,
  });

  // 3. Analyze Failures
  if (!isSuccess) {
    const body = res.body ? res.body.toString() : "";

    // Categorize the error based on the response text
    if (body.includes("sold out")) {
      soldOutErrors.add(1);
    } else if (
      body.includes("connection refused") ||
      body.includes("too many clients")
    ) {
      // This is likely your issue: Postgres ran out of connections
      dbConnectionErrors.add(1);
      console.error(`DB Error: ${body}`);
    } else {
      otherErrors.add(1);
      console.error(`Unexpected Error (${res.status}): ${body}`);
    }
  }
}
