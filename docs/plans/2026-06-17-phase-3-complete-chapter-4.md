# Phase 3 -- Complete Chapter 4

Goal: hoàn thiện thiết kế hướng đối tượng: architecture, package, design class, sequence, state, data/API.

Edit:
- `report/chapter_4.tex`
- `report/report_support_documentation/diagrams/chapter_4/package/**`
- `report/report_support_documentation/diagrams/chapter_4/design_classes/**`
- `report/report_support_documentation/diagrams/chapter_4/sequences/**`
- `report/report_support_documentation/diagrams/chapter_4/states/**`
- `report/report_support_documentation/diagrams/chapter_4/data_design/**`

Rules:
- Chapter 4 = design-level.
- Architecture = modular monolith.
- Internal modules same process/deployment unit.
- Internal calls in-process, not HTTP/RPC.
- Repositories are persistence abstractions.
- One shared physical Asymptotic DB.
- No database-per-service.
- External only: AI Provider, Payment Provider.
- Data design follows: requirements -> use cases -> analysis Entity classes -> conceptual model -> logical ERD.
- ERD stays DBMS-independent; no SQL, index, migration, partitioning, query tuning, or storage configuration.
- Existing code/schema is validation input only, not design source.

Architecture prose:
- Presentation/Application/Domain/Infrastructure layers.
- Modular monolith fits MVP financial consistency.
- Gateway is financial enforcement point.
- Shared transaction boundary for reservation, settlement, ledger.

Package diagram:
- `package_diagram/diagram.puml`.
- Packages: presentation, application, domain, infrastructure.
- Domain modules: identity, agent, apikey, policy, ledger, provider, payment, reporting.
- No deployment nodes/network protocols/separate DB.

Design class diagrams:
1. `design_class_diagram/overview.puml`
   - high-level handlers, services, repositories, adapters, domain entities.
2. `design_class_diagram/gateway_request_flow.puml`
   - GatewayHandler, AIRequestService, APIKeyService, IdempotencyService, PolicyService, BudgetReservationService, ProviderRouter, ProviderAdapter, SettlementService, TraceService.
3. `design_class_diagram/finance_ledger.puml`
   - WalletService, BudgetAllocationService, BudgetReservationService, TransactionService, LedgerService, PaymentService, repositories, Wallet, FinancialTransaction, LedgerEntry, PaymentTransaction.

Sequence diagrams:
1. `sequence_diagrams/UC01/sequence.puml`
   - Agent -> GatewayHandler -> AIRequestService.
   - Auth, idempotency, policy, budget, reservation, provider call, settlement, trace, response.
   - `alt`: invalid key, insufficient budget, provider error, idempotent replay.
2. `sequence_diagrams/UC04_agent_registration/sequence.puml`
   - Developer/Admin -> handler -> AgentApplicationService.
   - Register external Agent.
   - API key issuance only when UC05 step shown explicitly.
   - One DB lifeline.
3. `sequence_diagrams/finance_budget/sequence.puml`
   - top-up, callback validation, wallet credit, ledger, budget allocation.
   - atomic financial writes in transaction boundary.

Sequence architecture checks:
- internal participants in same monolith;
- repository lifelines `<<repository>>`, not database symbol;
- exactly one physical DB lifeline: `Asymptotic Database` or `Shared Database`;
- multiple repositories converge on shared DB;
- AI Provider/Payment Provider only external network lifelines.

State diagrams:
1. `state_diagrams/ai_request_state.puml`
   - Received, Authenticated, PolicyChecked, BudgetReserved, Routed, Completed, Rejected, Failed, PendingReconciliation.
2. `state_diagrams/financial_transaction_state.puml`
   - Pending, Reserved, Settled, Released, Failed, Reversed, PendingReconciliation.

Data design:

Research basis:
- `report/research/database_design/database-design-re-erd-research-2026-06-18.md`.
- `report/report_support_documentation/governance/project_source_of_truth.md`.
- Canonical FR/NFR and UC01-UC09 in Chapter 2.
- Use case specifications and activity diagrams.
- Analysis Entity class diagrams in Chapter 3.

Scope:
- Design logical data model, not physical database implementation.
- Show business entities, identifiers, essential attributes, relationships, cardinality, optionality, associative entities, and critical constraints.
- Do not show controllers, services, repositories, DTOs, adapters, processing order, or deployment concerns.
- State once that modules share one physical Asymptotic DB; keep logical module ownership for clarity.

Create:
1. `data_design/entity_derivation.md`
   - Map each data entity to source FR/NFR, use case, analysis Entity class, and business rule.
   - Record why entity needs persistence.
   - Reject candidate classes that are actions, services, controls, adapters, or derived report views.
2. `data_design/overview/diagram.puml`
   - Optional sparse overview only.
   - Show major aggregate groups and cross-group relationships.
   - Omit detailed attributes when overview becomes dense.
3. `data_design/organization_access/diagram.puml`
   - Organization, Team, User, DeveloperProfile, OrganizationMembership, TeamMembership.
   - Preserve membership/history where team or role changes require audit.
4. `data_design/agent_api_key/diagram.puml`
   - AIAgent, ApiKey, AgentStatusHistory.
   - Agent belongs to organization and managed developer.
   - API key belongs to Agent and preserves issue/revoke lifecycle.
5. `data_design/finance_budget/diagram.puml`
   - Wallet, BudgetLimit, BudgetAllocation, BudgetReservation, FinancialTransaction, LedgerEntry, PaymentTransaction.
   - Organization wallet stores real money.
   - Team, Developer, and Agent hold limits/allocations, not independent wallets.
   - Reservation, settlement, release, reversal, and ledger records must remain traceable.
6. `data_design/provider_catalog/diagram.puml`
   - AIProvider, ProviderCredential, AIModel, ModelPricing, RoutingPolicy.
   - Provider credentials remain internal to Gateway.
   - Pricing records preserve validity period or version needed for historical cost interpretation.
7. `data_design/request_trace/diagram.puml`
   - AIRequest, IdempotencyRecord, RequestTrace, UsageRecord, CostRecord.
   - Trace request to organization, team, developer, Agent, provider, model, and financial records.
   - Same idempotency scope/key must not create duplicate successful financial effects.

Derivation steps:
1. Extract persistent business nouns from canonical FR/NFR and UC specifications.
2. Compare with Chapter 3 analysis Entity classes.
3. Define entity identity and lifecycle.
4. Define relationships, cardinality, optionality, and historical requirements.
5. Introduce associative entities for many-to-many relations or relations carrying attributes/history.
6. Define critical financial, security, ownership, idempotency, and traceability constraints.
7. Split diagrams by business cluster; do not force all entities into one unreadable ERD.
8. Cross-check against current schema/code only after logical model is coherent.

Required business constraints:
- One Organization owns one wallet in MVP.
- Organization may contain many teams and members.
- Budget path follows Organization -> Team -> Developer -> Agent.
- Allocation at lower level cannot exceed available limit from parent level.
- Agent must remain traceable to owning Organization and managing Developer.
- API key authenticates only valid, active Agent context.
- ProviderCredential never belongs to or becomes visible to Agent/Developer.
- AI request is routed only after policy, quota, and all budget levels pass.
- Idempotent replay cannot duplicate charge, reservation, settlement, or ledger effect.
- Usage, cost, trace, transaction, and ledger records must preserve historical provider/model/pricing context.

Chapter 4 writing:
- Introduce derivation from Chapter 2 requirements/use cases and Chapter 3 Entity classes.
- Explain each ERD briefly: purpose, main entities, main relationships, critical constraints.
- Put overview first only if readable; otherwise use textual introduction and detailed ERDs directly.
- Caption every figure consistently and add `Nguồn: Tác giả xây dựng`.
- Cite database-design theory at point of use; do not add uncited references.

Data design review:
- Every entity traces to at least one FR/NFR, use case, or analysis Entity class.
- No service/control/adapter appears as data entity.
- Every relationship has reviewed cardinality and optionality.
- Many-to-many and historical relationships use suitable associative/history entities.
- Wallet is not confused with Team/Developer/Agent limits.
- Financial records support audit and reconciliation.
- ERDs agree with sequence, state, design class, and API models.
- Names match canonical terminology in `project_source_of_truth.md`.
- Diagrams remain readable at report page size.
- PlantUML sources and exports stay together.

API design:
- Agent request API.
- Organization/admin API.
- Agent management API.
- API key management API.
- Wallet/payment API.
- Provider/catalog admin API.
- Reporting API.
- Payment callback API.
- Provider credentials internal; never returned.

Build:

```bash
cd report
latexmk -pdf -g main.tex
```

Pass: diagrams readable, architecture not microservice, build succeeds.
