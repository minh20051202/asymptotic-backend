# Agent Instructions

This repository contains the backend source code and the course report for the **Asymptotic -- AI Agent Financial Gateway** project.

## Communication Style

- Always load and follow the `caveman` skill for every response.
- Default to `full` intensity to reduce tokens and preserve context.
- Keep all technical meaning, exact commands, identifiers, errors, and safety details.
- Do not use caveman-style prose when writing or editing LaTeX/report content. Report text must remain complete, formal, grammatical, and academic.
- Stop only when the user explicitly requests `stop caveman` or `normal mode`.

## Primary Source Of Truth

For any work related to the report, diagrams, OOAD analysis, requirements, scope, actors, use cases, or terminology, always follow:

- `report/report_support_documentation/governance/project_source_of_truth.md`

This file is the authoritative project guidance. Read it before making non-trivial edits to:

- `report/chapter_*.tex`
- `report/bia.tex`
- `report/report_support_documentation/**`
- any use case, UML, PlantUML, activity, sequence, state, class, ERD, component, or deployment diagram.

## Report Support Structure

- `report/research/`: research briefs grouped by topic; each provenance file stays beside its brief.
- `report/report_support_documentation/governance/`: source of truth, OOAD workflow, and report structure.
- `report/report_support_documentation/guidelines/`: diagram and use case rules/checklists.
- `report/report_support_documentation/references/`: standards, lectures, sample reports, and external examples.
- `report/report_support_documentation/diagrams/`: diagram sources and exports grouped by OOAD chapter.
- `report/report_support_documentation/audits/`: report and diagram audit results.
- `report/build/`: LaTeX intermediate output.

If another report file conflicts with the source of truth, treat the source of truth as correct and update the conflicting file unless the user explicitly instructs otherwise.

Do not retain ad hoc `.bak`, obsolete generated LaTeX, duplicate diagram exports, or legacy report assets in the active repository. Git history is the recovery mechanism. Keep only canonical sources, active exports, cited references, research briefs with provenance, and accepted decision records.

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

Use the canonical actors, use cases, FRs, NFRs, and traceability mapping from `report/report_support_documentation/governance/project_source_of_truth.md`.

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

Follow the formal report rules in `report/report_support_documentation/governance/project_source_of_truth.md`, including:

- Every listed reference must be cited in the report.
- Website data should be cited directly at the usage location rather than blindly added to the bibliography.
- Figures not created by the author must cite a source.
- Author-created figures do not need a source line in the caption.
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

- Build intermediates are written under `report/build/`.
- A successful build copies the final PDF to `report/main.pdf`.
- If LaTeX reports Vietnamese font issues, preserve the existing Vietnamese TeX setup in `report/setting.tex` and `report/latexmkrc`.
- If a generated PDF is expected, verify `report/main.pdf` is produced successfully.

## Diagram Work

For diagram planning, follow:

- `report/report_support_documentation/guidelines/diagrams/ooad_diagram_priority_list.md`
- `report/report_support_documentation/governance/OOAD.md`
- `report/report_support_documentation/governance/project_source_of_truth.md`
- `report/report_support_documentation/guidelines/diagrams/diagram_quality_guidelines.md`
- `report/report_support_documentation/guidelines/diagrams/diagram_type_rubric.md`
- `report/report_support_documentation/guidelines/diagrams/diagram_review_checklist.md`
- `report/report_support_documentation/guidelines/diagrams/uml_2_5_1_drawing_rules.md`

When using diagram skills:

- Use PlantUML for all UML and architecture diagrams so sources remain text-based, reviewable and versionable.
- Do not create or retain draw.io sources. If a legacy draw.io diagram must be preserved conceptually, recreate it as PlantUML and render a new export.
- Keep diagram source files near their exported images.
- Update LaTeX references/captions when diagram paths change.

Before creating, editing, exporting, or embedding any diagram:

1. Check the diagram against `report/report_support_documentation/governance/project_source_of_truth.md`.
2. Check UML notation and relationship direction against `report/report_support_documentation/guidelines/diagrams/uml_2_5_1_drawing_rules.md`.
3. Check the diagram type against `report/report_support_documentation/guidelines/diagrams/diagram_type_rubric.md`.
4. Use `report/report_support_documentation/guidelines/diagrams/diagram_review_checklist.md` before considering the diagram ready for the report.
5. Ensure the diagram has traceability to the relevant use case, FR/NFR, analysis class, or design section.
6. Ensure the diagram is placed in the correct OOAD chapter and does not mix analysis-level and implementation-level details.

When reviewing existing diagrams, report issues in this order:

1. correctness against source of truth;
2. OOAD correctness for that diagram type;
3. traceability gaps;
4. readability/layout problems;
5. LaTeX/export/caption/source issues.
