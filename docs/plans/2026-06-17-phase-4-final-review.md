# Final OOAD Chapter Review Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Rà soát cuối Chương 2-4 để bảo đảm đúng source of truth, đúng UML/OOAD, không còn placeholder, diagram đọc được và PDF build ổn.

**Architecture:** Phase này không thêm nội dung lớn. Nó kiểm tra traceability, hình thức LaTeX, caption/source, diagram readability và build warnings.

**Tech Stack:** LaTeX build, `rg`, image/PDF inspection.

---

### Task 1: Cross-Chapter Consistency Review

**Files:**
- Modify if needed: `report/chapter_2.tex`
- Modify if needed: `report/chapter_3.tex`
- Modify if needed: `report/chapter_4.tex`
- Modify: `report/report_support_documentation/chapter_2_3_4_completion_audit.md`

**Step 1: Check forbidden positioning**

Run:

```bash
rg -n "tạo AI Agent|huấn luyện AI Agent|triển khai AI Agent|điều phối AI Agent|provider API key trực tiếp|AI Provider API key" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
```

Expected: no incorrect description. If a match is legitimate, document why.

**Step 2: Check placeholders**

```bash
rg -n "Placeholder|TODO|\\\\textit\\{Placeholder" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
```

Expected: no active placeholders remain.

**Step 3: Check section balance**

```bash
rg -n "^\\\\section\\{" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
rg -n "^\\\\subsection\\{" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
```

Expected:

```text
Chapter 2: 6 main sections.
Chapter 3: 6 main sections.
Chapter 4: 6 main sections.
No numbered subsection unless intentionally needed.
```

### Task 2: Figure And Diagram Review

**Files:**
- Modify if needed: `report/chapter_2.tex`
- Modify if needed: `report/chapter_3.tex`
- Modify if needed: `report/chapter_4.tex`

**Step 1: Check captions**

```bash
rg -n "\\\\caption" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
```

Expected: every author-created diagram caption includes `Nguồn: Tác giả xây dựng`.

**Step 2: Check diagram complexity**

For each overview diagram:

```text
If labels overlap, text is unreadable, or relationship lines dominate the figure, split into smaller diagrams.
Do not shrink diagrams until text becomes unreadable.
```

**Step 3: Check UML rules**

Use `formal-17-12-05.pdf` and `diagram_type_rubric.md`:

```text
No non-standard UML relationships.
No reversed extend arrows.
No use case diagram modeled like activity diagram.
Class diagrams are in the correct chapter level: analysis in Chapter 3, design in Chapter 4.
```

### Task 3: LaTeX Quality Review

**Files:**
- Modify if needed: `report/chapter_2.tex`
- Modify if needed: `report/chapter_3.tex`
- Modify if needed: `report/chapter_4.tex`

**Step 1: Check semantic list anti-patterns**

Look for term-definition lists incorrectly written as bold-label itemize:

```bash
rg -n "\\\\item \\\\textbf\\{[^}]+:\\}" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
```

Expected: if a list is term-definition, convert to `description`; if it is a requirement bullet, keep as-is only when visually and semantically acceptable for the report style.

**Step 2: Check references**

```bash
rg -n "Figure~\\\\ref|Section~\\\\ref|Hình~\\\\ref|Bảng~\\\\ref" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
```

Expected: consistent with existing report style. If the report preamble does not support `cleveref`, do not introduce `\cref` without checking `setting.tex`.

### Task 4: Final Build

**Files:**
- Read: `report/main.log`
- Read: `report/main.pdf`

**Step 1: Build**

```bash
cd report
latexmk -pdf -g main.tex
```

Expected: build exits `0`, `report/main.pdf` exists.

**Step 2: Review warnings**

```bash
rg -n "Overfull|Underfull|LaTeX Warning|Package .* Warning" report/main.log
```

Expected: severe overfull caused by tables/figures in Chapters 2-4 is fixed. Minor warnings may remain.

**Step 3: Commit**

```bash
git add report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex report/report_support_documentation report/main.pdf
git commit -m "docs: complete ooad chapters 2 to 4"
```

