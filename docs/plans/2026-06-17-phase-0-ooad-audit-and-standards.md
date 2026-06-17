# OOAD Audit And Standards Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Tạo audit nền cho Chương 2-4 để biết phần nào đã hoàn thiện, phần nào thiếu và diagram nào cần tạo/tách nhỏ.

**Architecture:** Phase này chỉ đọc, đối chiếu và ghi nhận trạng thái. Không chỉnh nội dung chương trừ khi cần sửa lỗi build nhỏ. Kết quả là tài liệu audit làm checklist cho các phase sau.

**Tech Stack:** Markdown, `rg`, LaTeX build check.

---

### Task 1: Read Mandatory Standards

**Files:**
- Read: `AGENTS.md`
- Read: `report/report_support_documentation/project_source_of_truth.md`
- Read: `report/report_support_documentation/OOAD.md`
- Read: `report/report_support_documentation/ooad_diagram_priority_list.md`
- Read: `report/report_support_documentation/formal-17-12-05.pdf`
- Read: `report/report_support_documentation/diagram_type_rubric.md`
- Read: `report/report_support_documentation/diagram_review_checklist.md`

**Step 1: Read all standards**

Run:

```bash
sed -n '1,260p' AGENTS.md
sed -n '1,360p' report/report_support_documentation/project_source_of_truth.md
sed -n '1,360p' report/report_support_documentation/OOAD.md
sed -n '1,260p' report/report_support_documentation/ooad_diagram_priority_list.md
sed -n '1,260p' report/report_support_documentation/diagram_type_rubric.md
sed -n '1,220p' report/report_support_documentation/diagram_review_checklist.md
```

Expected: clear understanding of scope, actors, use cases, FR/NFR, OOAD order and diagram rules.

**Step 2: Inspect `formal-17-12-05.pdf`**

Use available PDF/text inspection. Extract or visually inspect the UML rules relevant to:

```text
Use case diagram
Class diagram
Activity diagram
Sequence diagram
State machine diagram
Package diagram
```

Expected: notes are used to prevent non-standard UML relationships such as `<<precondition>>`.

### Task 2: Create Completion Audit

**Files:**
- Read: `report/chapter_2.tex`
- Read: `report/chapter_3.tex`
- Read: `report/chapter_4.tex`
- Create: `report/report_support_documentation/chapter_2_3_4_completion_audit.md`

**Step 1: Check placeholders**

Run:

```bash
rg -n "Placeholder|TODO|\\\\textit\\{Placeholder" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
```

Expected: every active placeholder is identified.

**Step 2: Check existing diagrams**

Run:

```bash
find report/report_support_documentation -maxdepth 4 -type f \( -name '*.puml' -o -name '*.drawio' -o -name '*.png' \) | sort
```

Expected: all existing diagram sources and exports are listed.

**Step 3: Write audit file**

Create `report/report_support_documentation/chapter_2_3_4_completion_audit.md` with:

```markdown
# Chapter 2-4 Completion Audit

## Sources Checked

- AGENTS.md
- project_source_of_truth.md
- OOAD.md
- ooad_diagram_priority_list.md
- formal-17-12-05.pdf
- diagram_type_rubric.md
- diagram_review_checklist.md

## Chapter 2 Status

| Section | Status | Gaps | Action |
|---|---|---|---|

## Chapter 3 Status

| Section | Status | Gaps | Action |
|---|---|---|---|

## Chapter 4 Status

| Section | Status | Gaps | Action |
|---|---|---|---|

## Diagram Status

| Diagram | Priority | Source | Export | In Report | Action |
|---|---|---|---|---|---|

## Diagram Breakdown Decisions

| Diagram | Keep Overview | Breakdown Needed | Breakdown Groups | Reason |
|---|---|---|---|---|
```

### Task 3: Build Baseline

**Files:**
- Read: `report/main.tex`
- Read: `report/main.log`

**Step 1: Build report**

Run:

```bash
cd report
latexmk -pdf -g main.tex
```

Expected: build exits `0`, `report/main.pdf` exists.

**Step 2: Commit**

```bash
git add report/report_support_documentation/chapter_2_3_4_completion_audit.md report/main.pdf
git commit -m "docs: audit ooad chapters 2 to 4"
```

