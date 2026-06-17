# Complete Chapter 3 Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Hoàn thiện Chương 3 về phân tích hướng đối tượng theo Boundary-Control-Entity và mapping use case sang lớp phân tích.

**Architecture:** Chương 3 là analysis-level. Nó mô tả lớp biên, lớp điều khiển và lớp thực thể nghiệp vụ, không mô tả repository, DTO, framework, SQL hoặc implementation detail.

**Tech Stack:** LaTeX, PlantUML via `plantuml-skill`/Kroki for analysis class diagrams.

---

### Task 1: Rebuild Analysis Class Diagram

**Files:**
- Modify: `report/report_support_documentation/class_diagram/diagram.puml`
- Create/Modify: `report/report_support_documentation/class_diagram/diagram.png`
- Modify: `report/chapter_3.tex`
- Read: `report/report_support_documentation/formal-17-12-05.pdf`
- Read: `report/report_support_documentation/diagram_type_rubric.md`

**Step 1: Use PlantUML**

Use `plantuml-skill`. The source must be in:

```text
report/report_support_documentation/class_diagram/diagram.puml
```

**Step 2: Keep diagram analysis-level**

Include:

```text
Boundary: GatewayRequestBoundary, OrganizationAdminBoundary, SystemAdminBoundary, PaymentCallbackBoundary, AIProviderBoundary, ReportingBoundary
Control: AIRequestControl, AuthenticationControl, IdempotencyControl, BudgetControl, RoutingControl, SettlementControl, AgentAccessControl, PaymentControl, BudgetAllocationControl, ReportingControl, ProviderCatalogControl
Entity: Organization, User, DeveloperProfile, Agent, APIKey, Wallet, BudgetPolicy, AIProvider, AIModel, PricingPolicy, FinancialTransaction, LedgerEntry, PaymentTransaction, IdempotencyRecord, ExecutionTrace
```

Do not include:

```text
Repository
DTO
Migration
Controller implementation
Framework class
Database table-only names
```

**Step 3: Apply diagram complexity rule**

If the overview is too crowded:

```text
Keep BCE groups in overview.
Keep core entity multiplicities.
Move detailed control/entity dependencies to smaller section diagrams or robustness diagrams.
```

**Step 4: Render**

```bash
http=$(curl -s -w "%{http_code}" -o report/report_support_documentation/class_diagram/diagram.png \
  -X POST https://kroki.io/plantuml/png \
  -H "Content-Type: text/plain" \
  --data-binary "@report/report_support_documentation/class_diagram/diagram.puml")
echo "HTTP $http"
file report/report_support_documentation/class_diagram/diagram.png
```

Expected: `HTTP 200`, `PNG image data`.

**Step 5: Embed**

Use in `report/chapter_3.tex`:

```tex
\includegraphics[width=0.95\textwidth]{report_support_documentation/class_diagram/diagram.png}
\caption[Biểu đồ lớp phân tích tổng quát]{Biểu đồ lớp phân tích tổng quát của hệ thống. Nguồn: Tác giả xây dựng}
```

### Task 2: Write Chapter 3 Prose

**Files:**
- Modify: `report/chapter_3.tex`
- Read: `report/report_support_documentation/OOAD.md`
- Read: `report/report_support_documentation/project_source_of_truth.md`

**Step 1: Complete Section 3.1**

Write the basis of object-oriented analysis:

```text
Analysis focuses on problem-domain responsibilities.
Boundary-Control-Entity separates interaction, orchestration and business state.
Analysis classes are derived from use cases, activity diagrams, FR/NFR and business concepts.
```

**Step 2: Complete Section 3.2**

Introduce the analysis class diagram and explain what each group means.

**Step 3: Complete Sections 3.3-3.5**

Describe:

```text
Boundary classes and actor/system touchpoints.
Control classes and use case orchestration responsibilities.
Entity classes and core business state/lifecycle.
```

**Step 4: Complete Section 3.6**

Create a real mapping table:

```tex
\begin{longtable}{|L{0.14\textwidth}|L{0.26\textwidth}|L{0.28\textwidth}|L{0.22\textwidth}|}
\caption{Đối chiếu trường hợp sử dụng với lớp phân tích}\\
\hline
\textbf{Use Case} & \textbf{Lớp biên} & \textbf{Lớp điều khiển} & \textbf{Lớp thực thể chính} \\
\hline
...
\end{longtable}
```

Fill UC01-UC09 with real classes from the diagram.

### Task 3: Build And Commit Chapter 3

**Files:**
- Modify: `report/chapter_3.tex`
- Modify: `report/report_support_documentation/class_diagram/diagram.puml`
- Modify: `report/report_support_documentation/class_diagram/diagram.png`

**Step 1: Build**

```bash
cd report
latexmk -pdf -g main.tex
```

Expected: build succeeds.

**Step 2: Commit**

```bash
git add report/chapter_3.tex report/report_support_documentation/class_diagram/diagram.puml report/report_support_documentation/class_diagram/diagram.png report/main.pdf
git commit -m "docs: complete chapter 3 object analysis"
```

