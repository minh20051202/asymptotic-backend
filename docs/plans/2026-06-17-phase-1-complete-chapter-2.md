# Complete Chapter 2 Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Hoàn thiện Chương 2 về khảo sát, yêu cầu, actor, use case, use case specification và activity diagrams.

**Architecture:** Chương 2 chỉ làm phân tích yêu cầu và hành vi nghiệp vụ. Không đưa class, repository, database implementation hoặc sequence design vào chương này.

**Tech Stack:** LaTeX, PlantUML via `plantuml-skill`/Kroki for any new diagrams, existing draw.io/PNG diagrams where already accepted.

---

### Task 1: Verify Requirements And Actors

**Files:**
- Modify: `report/chapter_2.tex`
- Read: `report/report_support_documentation/project_source_of_truth.md`
- Read: `report/report_support_documentation/use_case_analysis.md`

**Step 1: Verify canonical actors**

Check Chương 2 contains:

```text
Người dùng
Quản trị viên tổ chức
Lập trình viên
AI Agent
Nhà cung cấp dịch vụ AI
Nhà cung cấp dịch vụ thanh toán
Quản trị viên hệ thống
```

**Step 2: Verify FR/NFR**

Keep bullet format. Ensure each FR/NFR:

```text
Uses canonical ID FR01-FR10 or NFR01-NFR05.
Starts content with capital letter.
States observable system behavior or quality constraint.
Does not contain implementation details like SQL locks or repository names.
```

**Step 3: Build**

```bash
cd report
latexmk -pdf -g main.tex
```

Expected: build succeeds.

### Task 2: Verify Use Case Specifications

**Files:**
- Modify: `report/chapter_2.tex`
- Read: `report/report_support_documentation/use_cases/UC01/specification.md`
- Read: `report/report_support_documentation/use_cases/UC02/specification.md`
- Read: `report/report_support_documentation/use_cases/UC03/specification.md`
- Read: `report/report_support_documentation/use_cases/UC04/specification.md`
- Read: `report/report_support_documentation/use_cases/UC05/specification.md`
- Read: `report/report_support_documentation/use_cases/UC06/specification.md`
- Read: `report/report_support_documentation/use_cases/UC07/specification.md`
- Read: `report/report_support_documentation/use_cases/UC08/specification.md`
- Read: `report/report_support_documentation/use_cases/UC09/specification.md`

**Step 1: Confirm canonical use cases**

Chương 2 must contain:

```text
UC01 -- Thực hiện yêu cầu AI qua Gateway
UC02 -- Nạp tiền vào ví tổ chức
UC03 -- Quản lý ngân sách tổ chức/team/developer
UC04 -- Đăng ký và quản lý AI Agent
UC05 -- Quản lý API key của Agent
UC06 -- Theo dõi giao dịch, usage, cost và trace
UC07 -- Đăng ký tổ chức
UC08 -- Quản lý lập trình viên/thành viên tổ chức
UC09 -- Quản lý provider, model và chính sách giá
```

**Step 2: Confirm each specification structure**

Every UC table must include:

```text
Tên
Mô tả
Tác nhân
Mức ưu tiên
Sự kiện kích hoạt
Điều kiện tiên quyết
Luồng chính
Luồng thay thế
Luồng ngoại lệ
Hậu điều kiện thành công
Hậu điều kiện thất bại
Liên kết yêu cầu
```

**Step 3: Verify traceability**

Match `Liên kết yêu cầu` exactly to source of truth:

```text
UC01: FR03, FR04, FR05, FR06, FR07, FR08, FR09, FR10; NFR01, NFR02, NFR03, NFR04, NFR05
UC02: FR01, FR09; NFR01, NFR02, NFR03, NFR04
UC03: FR01, FR05, FR09, FR10; NFR01, NFR02, NFR03
UC04: FR02, FR05, FR09, FR10; NFR02, NFR03
UC05: FR02, FR03, FR09, FR10; NFR02, NFR03
UC06: FR01, FR02, FR09; NFR01, NFR02, NFR03
UC07: FR01, FR09; NFR01, NFR02, NFR03
UC08: FR01, FR02, FR05, FR09, FR10; NFR01, NFR02, NFR03
UC09: FR04, FR05, FR07, FR09; NFR02, NFR03, NFR04
```

### Task 3: Verify Chapter 2 Diagrams

**Files:**
- Modify: `report/chapter_2.tex`
- Read: `report/report_support_documentation/ooad_diagram_priority_list.md`
- Read: `report/report_support_documentation/formal-17-12-05.pdf`
- Read: `report/report_support_documentation/diagram_review_checklist.md`
- Existing: `report/report_support_documentation/activity_diagrams/UC01/activity.puml`
- Existing: `report/report_support_documentation/activity_diagrams/UC04_agent_registration/activity.puml`
- Existing: `report/report_support_documentation/activity_diagrams/finance_budget/activity.puml`
- Optional create: `report/report_support_documentation/use_cases/agent_gateway_group/diagram.puml`
- Optional create: `report/report_support_documentation/use_cases/finance_admin_group/diagram.puml`

**Step 1: Check required diagrams**

Chương 2 should cover:

```text
Use Case Diagram tổng quát
Use Case Diagram nhóm Agent/Gateway, or equivalent detailed UC diagrams
Use Case Diagram nhóm Finance/Admin, or equivalent detailed UC diagrams
Activity Diagram UC01
Activity Diagram đăng ký AI Agent và cấp API key
Activity Diagram nạp tiền và phân bổ ngân sách
```

**Step 2: Apply diagram complexity rule**

Do not force every UC and relationship into one overview. If too crowded:

```text
Overview: only actors + 9 main UCs.
Agent/Gateway group: UC01, UC04, UC05, UC06, UC09.
Finance/Admin group: UC02, UC03, UC07, UC08, UC06.
```

**Step 3: Use PlantUML for new diagrams**

If creating new diagrams, use `plantuml-skill`, render via Kroki and validate:

```bash
http=$(curl -s -w "%{http_code}" -o output.png \
  -X POST https://kroki.io/plantuml/png \
  -H "Content-Type: text/plain" \
  --data-binary "@diagram.puml")
echo "HTTP $http"
file output.png
```

Expected: `HTTP 200`, `PNG image data`.

**Step 4: UML checks**

Before embedding:

```text
No <<precondition>> relationship.
<<extend>> arrow points from extension use case to base use case.
No activity steps modeled as use cases.
No actor inside system boundary.
No actor connected to included/extended child UC unless actor directly initiates it.
```

**Step 5: Captions and labels**

Every figure must include:

```tex
\caption[Short caption]{Full caption. Nguồn: Tác giả xây dựng}
```

### Task 4: Build And Commit Chapter 2

**Files:**
- Modify: `report/chapter_2.tex`
- Modify as needed: `report/report_support_documentation/use_cases/**`
- Modify as needed: `report/report_support_documentation/activity_diagrams/**`

**Step 1: Build**

```bash
cd report
latexmk -pdf -g main.tex
```

Expected: build succeeds.

**Step 2: Commit**

```bash
git add report/chapter_2.tex report/report_support_documentation/use_cases report/report_support_documentation/activity_diagrams report/main.pdf
git commit -m "docs: complete chapter 2 requirement analysis"
```

