# Complete Chapter 4 Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Hoàn thiện Chương 4 về thiết kế hướng đối tượng: kiến trúc, package, design class, sequence, state, dữ liệu và API.

**Architecture:** Chương 4 là design-level, được phép đưa service, repository, adapter, handler và transaction boundary. Nội dung phải trace được từ Chương 2-3 và không mâu thuẫn với analysis model.

**Tech Stack:** LaTeX, PlantUML via `plantuml-skill`/Kroki for package, design class, sequence and state diagrams.

---

### Task 1: Architecture And Package Design

**Files:**
- Modify: `report/chapter_4.tex`
- Create: `report/report_support_documentation/package_diagram/diagram.puml`
- Create: `report/report_support_documentation/package_diagram/diagram.png`
- Read: `report/report_support_documentation/OOAD.md`
- Read: `report/report_support_documentation/formal-17-12-05.pdf`

**Step 1: Complete architecture prose**

Explain:

```text
Presentation Layer
Application Layer
Domain Layer
Infrastructure Layer
Why modular monolith fits MVP financial consistency
Why Gateway is financial enforcement point
```

**Step 2: Create package diagram using PlantUML**

Create:

```text
report/report_support_documentation/package_diagram/diagram.puml
```

Packages:

```text
presentation, application, domain, infrastructure
identity, agent, apikey, policy, ledger, provider, payment, reporting
```

No circular dependencies. If too crowded, split by layers and domain packages.

**Step 3: Render and embed**

Render via Kroki, validate `HTTP 200`, inspect PNG, then embed in `report/chapter_4.tex` with caption ending `Nguồn: Tác giả xây dựng`.

### Task 2: Design Class Diagrams

**Files:**
- Modify: `report/chapter_4.tex`
- Create: `report/report_support_documentation/design_class_diagram/overview.puml`
- Create: `report/report_support_documentation/design_class_diagram/overview.png`
- Create: `report/report_support_documentation/design_class_diagram/gateway_request_flow.puml`
- Create: `report/report_support_documentation/design_class_diagram/gateway_request_flow.png`
- Create: `report/report_support_documentation/design_class_diagram/finance_ledger.puml`
- Create: `report/report_support_documentation/design_class_diagram/finance_ledger.png`

**Step 1: Create overview design class diagram**

Use service/repository/adapter level classes:

```text
GatewayHandler
AdminHandler
AIRequestService
AgentService
APIKeyService
BudgetService
PolicyService
ProviderRoutingService
SettlementService
WalletService
TransactionService
TraceService
ProviderAdapter
PaymentAdapter
Repository interfaces
Domain entities
```

If too crowded, keep it high-level and rely on the two focused diagrams.

**Step 2: Create Gateway Request Flow diagram**

Focus UC01:

```text
GatewayHandler
AIRequestService
APIKeyService
IdempotencyService
PolicyService
BudgetReservationService
ProviderRouter
ProviderAdapter
SettlementService
TraceService
WalletRepository
TransactionRepository
```

**Step 3: Create Finance/Ledger diagram**

Focus:

```text
WalletService
BudgetAllocationService
BudgetReservationService
TransactionService
LedgerService
PaymentService
WalletRepository
FinancialTransactionRepository
LedgerEntryRepository
Wallet
FinancialTransaction
LedgerEntry
PaymentTransaction
```

**Step 4: Render and embed**

Use `plantuml-skill` render loop. Embed all readable diagrams with captions.

### Task 3: Sequence Diagrams

**Files:**
- Modify: `report/chapter_4.tex`
- Create: `report/report_support_documentation/sequence_diagrams/UC01/sequence.puml`
- Create: `report/report_support_documentation/sequence_diagrams/UC01/sequence.png`
- Create: `report/report_support_documentation/sequence_diagrams/UC04_agent_registration/sequence.puml`
- Create: `report/report_support_documentation/sequence_diagrams/UC04_agent_registration/sequence.png`
- Create: `report/report_support_documentation/sequence_diagrams/finance_budget/sequence.puml`
- Create: `report/report_support_documentation/sequence_diagrams/finance_budget/sequence.png`

**Step 1: UC01 sequence**

Show:

```text
AI Agent -> GatewayHandler -> AIRequestService
API key authentication
Idempotency check
Policy/budget check
Budget reservation
Provider call using internal credential
Usage metadata
Settlement and trace
Response
```

Include `alt` for invalid key, insufficient budget, provider error and idempotent replay.

If too wide, split into:

```text
Phase 1: authenticate/idempotency/policy/budget
Phase 2: provider call/streaming/settlement/trace
```

**Step 2: UC04 sequence**

Show external Agent registration and API key issuance. Do not imply the system creates or hosts the Agent.

**Step 3: Finance/budget sequence**

Show top-up, payment callback validation, wallet credit and budget allocation.

### Task 4: State Machine Diagrams

**Files:**
- Modify: `report/chapter_4.tex`
- Create: `report/report_support_documentation/state_diagrams/ai_request_state.puml`
- Create: `report/report_support_documentation/state_diagrams/ai_request_state.png`
- Create: `report/report_support_documentation/state_diagrams/financial_transaction_state.puml`
- Create: `report/report_support_documentation/state_diagrams/financial_transaction_state.png`

**Step 1: AI request states**

Use:

```text
Received, Authenticated, PolicyChecked, BudgetReserved, Routed, Completed, Rejected, Failed, PendingReconciliation
```

**Step 2: Financial transaction states**

Use:

```text
Pending, Reserved, Settled, Released, Failed, Reversed, PendingReconciliation
```

**Step 3: Render and embed**

Render via Kroki and embed under `\section{Thiết kế trạng thái}`.

### Task 5: Data And API Design

**Files:**
- Modify: `report/chapter_4.tex`
- Optional create: `report/report_support_documentation/api_design_summary.md`
- Optional create: `report/report_support_documentation/data_design_summary.md`

**Step 1: Write data design**

Cover:

```text
Organization/User/Developer/Agent/APIKey
Wallet/BudgetPolicy/FinancialTransaction/LedgerEntry/PaymentTransaction
AIProvider/AIModel/PricingPolicy
IdempotencyRecord/ExecutionTrace
```

**Step 2: Write API design**

Describe endpoint groups:

```text
Agent request API
Organization/admin API
Agent/API key management API
Wallet/payment API
Provider/catalog admin API
Reporting API
Payment callback API
```

Clarify provider credentials are internal and never returned.

### Task 6: Build And Commit Chapter 4

**Files:**
- Modify: `report/chapter_4.tex`
- Create/modify: `report/report_support_documentation/package_diagram/**`
- Create/modify: `report/report_support_documentation/design_class_diagram/**`
- Create/modify: `report/report_support_documentation/sequence_diagrams/**`
- Create/modify: `report/report_support_documentation/state_diagrams/**`

**Step 1: Build**

```bash
cd report
latexmk -pdf -g main.tex
```

Expected: build succeeds.

**Step 2: Commit**

```bash
git add report/chapter_4.tex report/report_support_documentation/package_diagram report/report_support_documentation/design_class_diagram report/report_support_documentation/sequence_diagrams report/report_support_documentation/state_diagrams report/main.pdf
git commit -m "docs: complete chapter 4 object design"
```

