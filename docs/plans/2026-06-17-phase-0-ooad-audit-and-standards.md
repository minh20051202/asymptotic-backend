# Phase 0 -- OOAD Audit

Goal: tạo baseline audit cho Chương 2-4.

Read:
- `AGENTS.md`
- `report/report_support_documentation/governance/project_source_of_truth.md`
- `report/report_support_documentation/governance/OOAD.md`
- `report/report_support_documentation/guidelines/diagrams/ooad_diagram_priority_list.md`
- `report/report_support_documentation/references/standards/formal-17-12-05.pdf`
- `report/report_support_documentation/guidelines/diagrams/diagram_type_rubric.md`
- `report/report_support_documentation/guidelines/diagrams/diagram_review_checklist.md`
- `report/chapter_2.tex`
- `report/chapter_3.tex`
- `report/chapter_4.tex`

Do:
1. Check placeholder/TODO in Chapters 2-4.
2. List PlantUML/PNG diagrams under `report/report_support_documentation`.
3. Check diagram gap against priority list.
4. Check architecture mismatch:
   - no microservice claim;
   - no database-per-service;
   - internal modules in-process;
   - one shared Asymptotic DB;
   - only AI Provider/Payment Provider external.
5. Write `report/report_support_documentation/audits/chapter_2_3_4_completion_audit.md`.

Audit file must contain:
- source checked;
- Chapter 2/3/4 status;
- diagram status;
- diagram split decision;
- architecture consistency table;
- remaining action list.

Build:

```bash
cd report
latexmk -pdf -g main.tex
```

Pass: audit exists, build succeeds, gaps clear.
