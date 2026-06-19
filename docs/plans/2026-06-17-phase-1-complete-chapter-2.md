# Phase 1 -- Complete Chapter 2

Goal: hoàn thiện khảo sát, FR/NFR, actor, use case, use case spec, activity diagrams.

Edit:
- `report/chapter_2.tex`
- `report/report_support_documentation/diagrams/chapter_2/use_cases/**`
- `report/report_support_documentation/diagrams/chapter_2/activity_diagrams/**`

Rules:
- Chapter 2 only requirement/business behavior.
- No class/repository/DB/design sequence detail.
- Use canonical actor/UC/FR/NFR from source of truth.
- PlantUML-only.

Canonical actors:
- Người dùng
- Quản trị viên tổ chức
- Lập trình viên
- AI Agent
- Quản trị viên hệ thống
- AI Provider
- Payment Provider

Canonical UCs:
- UC01 -- Thực hiện yêu cầu AI qua Gateway
- UC02 -- Nạp tiền vào ví tổ chức
- UC03 -- Quản lý và phân bổ ngân sách nội bộ
- UC04 -- Đăng ký và quản lý AI Agent
- UC05 -- Quản lý API key của Agent
- UC06 -- Theo dõi giao dịch, usage, cost và trace
- UC07 -- Đăng ký tổ chức
- UC08 -- Quản lý đội ngũ và thành viên tổ chức
- UC09 -- Quản lý provider, model và chính sách giá

UC rule:
- UC04: Developer đăng ký/quản lý Agent bên ngoài.
- UC05: cấp/thu hồi API key.
- UC03: budget chain Organization -> Team -> Developer -> Agent.
- Agent không nhận provider credential.

Each UC spec needs:
- tên;
- mô tả;
- tác nhân;
- ưu tiên;
- trigger;
- precondition;
- main flow;
- alternative flow;
- exception flow;
- success/failure postcondition;
- FR/NFR links.

Traceability:
- UC01: FR03, FR04, FR05, FR06, FR07, FR08, FR09, FR10; NFR01-NFR05
- UC02: FR01, FR09; NFR01-NFR04
- UC03: FR01, FR05, FR09, FR10; NFR01-NFR03
- UC04: FR02, FR05, FR09, FR10; NFR02, NFR03
- UC05: FR02, FR03, FR09, FR10; NFR02, NFR03
- UC06: FR01, FR02, FR09; NFR01-NFR03
- UC07: FR01, FR09; NFR01-NFR03
- UC08: FR01, FR02, FR05, FR09, FR10; NFR01-NFR03
- UC09: FR04, FR05, FR07, FR09; NFR02-NFR04

Diagrams:
- overview use case;
- detailed use case diagrams if overview crowded;
- activity UC01;
- activity UC04 Agent registration;
- activity finance/budget.

UML checks:
- no `<<precondition>>`;
- correct `<<extend>>` direction;
- no activity step as use case;
- no actor inside system boundary;
- caption has `Nguồn: Tác giả xây dựng`.

Build:

```bash
cd report
latexmk -pdf -g main.tex
```

Pass: Chapter 2 complete, specs consistent, diagrams readable, build succeeds.
