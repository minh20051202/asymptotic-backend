# Phase 2 -- Complete Chapter 3

Goal: hoàn thiện phân tích hướng đối tượng bằng BCE + analysis class diagrams.

Edit:
- `report/chapter_3.tex`
- `report/report_support_documentation/analysis_class_diagrams/**`

Read:
- source of truth;
- UC01-UC09 specs;
- activity diagrams;
- FR/NFR traceability;
- OOAD/diagram rubric/checklist.

Rules:
- Chapter 3 = analysis model only.
- No repository, DTO, endpoint, framework, DB, modular-monolith detail.
- Use one sparse system-wide Analysis Class Diagram plus focused analysis views.
- Overview and focused diagrams are views of one analysis model, not independent models.
- Class must trace to use case/spec/rule/FR/NFR.
- Code/schema only for cross-check, not source of truth.

Derivation:
1. Read UC spec.
2. Extract actor, trigger, pre/post, flows, rules, FR/NFR.
3. Find BCE candidates:
   - actor/system touchpoint -> Boundary;
   - orchestration/rule/state change -> Control;
   - business noun with identity/state/lifecycle -> Entity.
4. Remove technical/duplicate classes.
5. Add relation only when rule/flow proves it.
6. Record in `analysis_class_diagrams/class_derivation.md`.

Diagrams to create/check:
1. `overview/diagram.puml`: system-wide analysis overview.
   - Show core Boundary, Control, Entity classes and only major cross-cluster relationships.
   - No attributes, operations, repositories, DTOs, frameworks, or detailed multiplicities.
   - Target 20-30 visible classes; split or reduce relationships if page readability fails.
2. `ai_request/diagram.puml`: UC01.
   - Boundary: AgentRequestBoundary, AIProviderBoundary.
   - Control: AIRequestControl, AuthenticationControl, IdempotencyControl, PolicyEvaluationControl, BudgetControl, ProviderRoutingControl, SettlementControl, TraceControl.
   - Entity: Organization, Team, DeveloperProfile, Agent, APIKey, UsagePolicy, Wallet, BudgetLimit, BudgetReservation, AIProvider, AIModel, FinancialTransaction, IdempotencyRecord, ExecutionTrace.
3. `finance_budget/diagram.puml`: UC02, UC03.
   - Budget chain Organization -> Team -> Developer -> Agent.
   - Team/Developer/Agent are limits, not wallets.
4. `agent_access/diagram.puml`: UC04, UC05.
   - Developer registers/manages external Agent.
   - APIKey separate.
   - No Agent creation/training/hosting/orchestration.
5. `organization_membership/diagram.puml`: UC07, UC08.
   - User, Organization, OrganizationMembership, Team, TeamMembership, DeveloperProfile, Agent.
6. `reporting_trace/diagram.puml`: UC06.
   - usage, cost, transaction, trace.
7. `provider_catalog/diagram.puml`: UC09.
   - AIProvider, AIModel, ProviderEndpoint, ProviderCredential, PricingPolicy, FallbackPolicy.
   - ProviderCredential internal. Developer/Agent never receive provider key.

Diagram checks:
- 8-18 classes preferred.
- stereotypes: `<<boundary>>`, `<<control>>`, `<<entity>>`.
- Analysis Class diagrams use class boxes with stereotypes; BCE circle icons are reserved for Robustness diagrams.
- no Actor -> Control/Entity.
- no Boundary business rule.
- no Entity -> Boundary.
- Entity -> Entity association is allowed when it represents a proven business relationship.
- multiplicity only when source supports it.
- readable layout.
- same class name, stereotype, responsibility, and relationship semantics across overview and focused diagrams.
- maintain `analysis_class_diagrams/class_catalog.md` as the shared class definition and traceability source.

Chapter 3 prose:
- explain BCE method shortly;
- explain class derivation;
- embed six diagrams;
- discuss key Boundary/Control/Entity, not full class list;
- add UC-to-analysis-class table for UC01-UC09.

Transition check:
- Boundary -> Handler/API boundary in Chapter 4.
- Control -> service/module responsibility in Chapter 4.
- Entity -> domain entity/aggregate in Chapter 4.
- no design class without UC/FR/analysis source.

Build:

```bash
cd report
latexmk -pdf -g main.tex
```

Verify:

```bash
rg -n "analysis_class_diagrams" report/chapter_3.tex
```

Pass: overview and six focused diagrams readable, consistent with class catalog, build succeeds.
