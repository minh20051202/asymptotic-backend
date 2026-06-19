# OOAD Chapters 2-4 Plan

Goal: hoàn thiện Chương 2-4 cho **Asymptotic -- AI Agent Financial Gateway**.

Source bắt buộc:
- `AGENTS.md`
- `report/report_support_documentation/governance/project_source_of_truth.md`
- `report/report_support_documentation/governance/OOAD.md`
- `report/report_support_documentation/guidelines/diagrams/ooad_diagram_priority_list.md`
- `report/report_support_documentation/guidelines/diagrams/diagram_type_rubric.md`
- `report/report_support_documentation/guidelines/diagrams/diagram_review_checklist.md`

Phase:
1. [Phase 0 -- Audit](2026-06-17-phase-0-ooad-audit-and-standards.md)
2. [Phase 1 -- Chapter 2](2026-06-17-phase-1-complete-chapter-2.md)
3. [Phase 2 -- Chapter 3](2026-06-17-phase-2-complete-chapter-3.md)
4. [Phase 3 -- Chapter 4](2026-06-17-phase-3-complete-chapter-4.md)
5. [Phase 4 -- Final Review](2026-06-17-phase-4-final-review.md)

Global rule:
- Bám source of truth.
- Không mô tả hệ thống tạo, train, deploy, host, orchestrate AI Agent.
- Developer tạo/đăng ký/quản lý AI Agent bên ngoài.
- API key management tách khỏi Agent registration.
- Budget chain: Organization -> Team -> Developer -> AI Agent.
- Organization wallet giữ tiền thật; Team/Developer/Agent là hạn mức.
- Provider credential nội bộ Gateway; không trả cho Developer/Agent.
- Chương 2 = yêu cầu + use case + activity.
- Chương 3 = analysis BCE, không repository/DB/framework.
- Chương 4 = design, modular monolith, shared physical DB.
- PlantUML-only. Không tạo/giữ draw.io.
- Hình tự tạo caption có `Nguồn: Tác giả xây dựng`.
- Văn ngắn, đúng OOAD, không lặp.

Build sau mỗi phase:

```bash
cd report
latexmk -pdf -g main.tex
```

Pass: `report/main.pdf` sinh thành công, không fatal error, không undefined refs/citations mới.
