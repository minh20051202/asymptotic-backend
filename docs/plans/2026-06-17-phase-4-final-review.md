# Phase 4 -- Final Review

Goal: final check Chương 2-4 before handoff.

Edit if needed:
- `report/chapter_2.tex`
- `report/chapter_3.tex`
- `report/chapter_4.tex`
- `report/report_support_documentation/audits/chapter_2_3_4_completion_audit.md`

Checks:
1. Source-of-truth positioning:
   - no tạo/train/deploy/host/orchestrate AI Agent;
   - no Agent direct provider API key;
   - Developer manages/registers external Agent;
   - UC04/UC05 separated;
   - budget chain Organization -> Team -> Developer -> Agent.
2. Placeholders:
   - no `Placeholder`, `TODO`, active `\textit{Placeholder}`.
3. Structure:
   - Chapter 2 has 6 main sections;
   - Chapter 3 has 6 main sections;
   - Chapter 4 has 6 main sections.
4. Figures:
   - every author-created diagram caption has `Nguồn: Tác giả xây dựng`;
   - no unreadable text/overcrowded overview.
5. UML:
   - no non-standard relation;
   - correct `extend` direction;
   - analysis class diagrams only in Chapter 3;
   - design diagrams only in Chapter 4.
6. Architecture:
   - modular monolith;
   - internal modules logical only;
   - internal calls in-process;
   - repositories not service-owned DB;
   - UC01/UC04/finance sequence use one shared DB;
   - financial atomic ops use one transaction boundary;
   - only AI Provider/Payment Provider external.
7. LaTeX:
   - no bad term-definition itemize unless intentional;
   - references/captions consistent;
   - no new severe overfull from Chapters 2-4.

Commands:

```bash
rg -n "tạo AI Agent|huấn luyện AI Agent|triển khai AI Agent|điều phối AI Agent|provider API key trực tiếp|AI Provider API key" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
rg -n "Placeholder|TODO|\\\\textit\\{Placeholder" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
rg -n "\\\\caption" report/chapter_2.tex report/chapter_3.tex report/chapter_4.tex
cd report
latexmk -pdf -g main.tex
rg -n "Overfull|Underfull|LaTeX Warning|Package .* Warning" main.log
```

Pass: build succeeds, audit updated, no blocker warning, `report/main.pdf` ready.
