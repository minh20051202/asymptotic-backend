# Agent Instructions

This repository contains the backend source code and the course report for the **Asymptotic -- AI Agent Financial Gateway** project.

## Primary Source Of Truth

For any work related to the report, diagrams, OOAD analysis, requirements, scope, actors, use cases, or terminology, always follow:

- `report/report_support_documentation/project_source_of_truth.md`

This file is the authoritative project guidance. Read it before making non-trivial edits to:

- `report/chapter_*.tex`
- `report/bia.tex`
- `report/generated/*.tex`
- `report/report_support_documentation/**`
- any use case, UML, PlantUML, draw.io, activity, sequence, state, class, ERD, component, or deployment diagram.

If another report file conflicts with the source of truth, treat the source of truth as correct and update the conflicting file unless the user explicitly instructs otherwise.

## Report Positioning

Keep the project positioned as:

- Vietnamese title: **THIẾT KẾ VÀ XÂY DỰNG CỔNG TÀI CHÍNH THỜI GIAN THỰC CHO TÁC NHÂN TRÍ TUỆ NHÂN TẠO**
- System name: **Asymptotic -- AI Agent Financial Gateway**

The system is a real-time financial gateway for external AI Agents. It does **not** create, train, deploy, host, or orchestrate AI Agents.

Core flow:

1. Organization funds its wallet/budget.
2. Organization or authorized developer registers an external AI Agent.
3. Gateway issues an Asymptotic API key for the Agent.
4. Agent calls Gateway using the Asymptotic API key.
5. Gateway checks identity, policy, quota, and budget.
6. Gateway calls the AI Provider using internal provider credentials.
7. Gateway records usage, cost, trace, and financial state.

Never describe the Agent as using the AI Provider API key directly. Provider credentials are internal to the Gateway.

## OOAD Structure

When editing the report, preserve the OOAD progression:

1. Chapter 1: introduction, problem statement, target problem, goals/scope, method, report structure.
2. Chapter 2: current-state survey, system requirements, actor analysis, use case analysis, use case specification, business-flow/activity modeling.
3. Chapter 3: object-oriented analysis, OOAD basis, analysis class identification, Boundary-Control-Entity analysis, use-case-to-class mapping.
4. Chapter 4: object-oriented design, architecture, packages, classes, interactions, states, data/API design.
5. Chapter 5: implementation environment, implemented functions, testing, result evaluation, implementation limitations.
6. Chapter 6: achieved results, project limitations, future work.

Do not move implementation-heavy discussion into analysis sections. Keep technology, code, database implementation details, and deployment details in design/implementation chapters unless the source of truth says otherwise.

## Use Case And Requirement Consistency

Use the canonical actors, use cases, FRs, NFRs, and traceability mapping from `project_source_of_truth.md`.

Important rule:

- The use case **Đăng ký và quản lý AI Agent** must exist explicitly.
- API key management is a separate concern from AI Agent registration.
- Use case traceability must map back to the canonical FR/NFR IDs.

Before changing use cases or diagrams, check whether the change affects:

- actor names;
- use case names;
- FR/NFR mapping;
- diagram captions;
- generated LaTeX;
- Chapter 3 analysis classes;
- Chapter 4 sequence/state/class diagrams.

## Formal Report Rules

Follow the formal report rules in `project_source_of_truth.md`, including:

- Every listed reference must be cited in the report.
- Website data should be cited directly at the usage location rather than blindly added to the bibliography.
- Figures not created by the author must cite a source.
- Author-created figures should say: `Nguồn: Tác giả xây dựng`.
- Avoid spelling, punctuation, and incomplete sentence errors.
- Do not explain English terms in the style: `Học máy (machine learning) là...`.
- Keep chapter sizes and subsection structure reasonably balanced.
- Keep the cover consistent with the official school/institute template and the finalized project title.

## Editing And Build Rules

- Use `apply_patch` for manual edits.
- Do not delete user-created files or generated backups unless explicitly asked.
- After LaTeX/report edits, build from `report/` with:

```bash
latexmk -pdf -g main.tex
```

- If LaTeX reports Vietnamese font issues, preserve the existing Vietnamese TeX setup in `report/setting.tex` and `report/latexmkrc`.
- If a generated PDF is expected, verify `report/main.pdf` is produced successfully.

## Diagram Work

For diagram planning, follow:

- `report/report_support_documentation/ooad_diagram_priority_list.md`
- `report/report_support_documentation/OOAD.md`
- `report/report_support_documentation/project_source_of_truth.md`
- `report/report_support_documentation/diagram_quality_guidelines.md`
- `report/report_support_documentation/diagram_type_rubric.md`
- `report/report_support_documentation/diagram_review_checklist.md`

When using diagram skills:

- Prefer PlantUML for UML diagrams that should stay text-based and versionable.
- Prefer draw.io for polished diagrams that need richer layout or styling.
- Keep diagram source files near their exported images.
- Update LaTeX references/captions when diagram paths change.

Before creating, editing, exporting, or embedding any diagram:

1. Check the diagram against `project_source_of_truth.md`.
2. Check the diagram type against `diagram_type_rubric.md`.
3. Use `diagram_review_checklist.md` as the review checklist before considering the diagram ready for the report.
4. Ensure the diagram has traceability to the relevant use case, FR/NFR, analysis class, or design section.
5. Ensure the diagram is placed in the correct OOAD chapter and does not mix analysis-level and implementation-level details.

When reviewing existing diagrams, report issues in this order:

1. correctness against source of truth;
2. OOAD correctness for that diagram type;
3. traceability gaps;
4. readability/layout problems;
5. LaTeX/export/caption/source issues.
