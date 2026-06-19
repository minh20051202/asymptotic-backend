# Reorganize Report Artifacts Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Reorganize `report/` and move `outputs/` into topic-based report research folders without breaking LaTeX, documentation references, or diagram traceability.

**Architecture:** Keep primary LaTeX sources at `report/` root. Group research by topic, support documents by role, diagrams by OOAD chapter, references by source type, and build intermediates under `report/build/`. Preserve all backups under a dedicated legacy tree.

**Tech Stack:** Bash filesystem operations, Markdown, LaTeX, latexmk, PlantUML source files.

---

### Task 1: Create Destination Structure

Create:

```text
report/build/
report/research/{requirements,use_cases,ooad,database_design,market}/
report/images/{report_assets,legacy}/
report/report_support_documentation/{governance,guidelines/diagrams,guidelines/use_cases,references/standards,references/lectures,references/sample_reports,references/examples,diagrams/chapter_2,diagrams/chapter_3,diagrams/chapter_4,proposals,audits,legacy}/
```

### Task 2: Move Research and Supporting Documents

- Move every `outputs/*.md` and matching provenance file into the relevant `report/research/<topic>/`.
- Move governance, guideline, reference, proposal, and audit files into their destination folders.
- Move `report/report_structure.md` into governance.
- Preserve backup files under `report/report_support_documentation/legacy/`.

### Task 3: Move Diagrams by OOAD Chapter

- Chapter 2: use cases and activity diagrams.
- Chapter 3: analysis class diagrams and future robustness diagrams.
- Chapter 4: package, design class, sequence, state, and data design diagrams.
- Keep each `.puml` beside its rendered image.

### Task 4: Organize Report Images

- Move `logo_fami.jpg` into `report/images/report_assets/`.
- Move legacy numbered images and backup images into `report/images/legacy/`.
- Update active LaTeX image paths.

### Task 5: Update References

Modify:

- `AGENTS.md`
- `report/chapter_*.tex`
- `report/bia.tex`
- `report/generated/*.tex`
- all Markdown plans, research briefs, provenance files, audits, and guidelines containing old paths.

Rules:

- `project_source_of_truth.md` new canonical path:
  `report/report_support_documentation/governance/project_source_of_truth.md`.
- Diagram guidance new path:
  `report/report_support_documentation/guidelines/diagrams/`.
- Use case guidance new path:
  `report/report_support_documentation/guidelines/use_cases/use_case_analysis.md`.

### Task 6: Configure Build Directory

Modify `report/latexmkrc`:

- send intermediate and generated build output to `report/build/`;
- copy successful `build/main.pdf` to `report/main.pdf`;
- preserve shell escape and Vietnamese TeX configuration.

Move existing intermediate files and `_minted-main/` into `report/build/`.

### Task 7: Verify

Run:

```bash
rg -n "outputs/|report_support_documentation/(use_cases|activity_diagrams|class_diagram|package_diagram|design_class_diagram|sequence_diagrams|state_diagrams)|report/images/logo_fami.jpg" .
find report -maxdepth 4 -type d | sort
cd report
latexmk -pdf -g main.tex
test -f main.pdf
```

Pass criteria:

- no active references to old paths;
- no files lost;
- `.puml` and exports remain adjacent;
- `report/main.pdf` exists;
- LaTeX build exits successfully;
- build intermediates are under `report/build/`.
