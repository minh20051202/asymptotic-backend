# Complete OOAD Chapters 2-4 Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Hoàn thiện Chương 2, Chương 3 và Chương 4 của báo cáo OOAD cho Asymptotic -- AI Agent Financial Gateway bằng các phase nhỏ, dễ kiểm soát.

**Architecture:** Đây là plan index. Mỗi phase có file riêng để tránh một plan quá dài và giúp thực thi tuần tự: audit, Chương 2, Chương 3, Chương 4, final review. Tất cả phase đều bám `AGENTS.md`, `project_source_of_truth.md`, `OOAD.md`, `ooad_diagram_priority_list.md`, `formal-17-12-05.pdf`, `diagram_type_rubric.md` và `diagram_review_checklist.md`.

**Tech Stack:** LaTeX (`latexmk -pdf -g main.tex`), PlantUML qua `plantuml-skill`/Kroki, PNG exports nhúng vào report, Markdown support docs.

---

## Phase Files

1. [Phase 0 -- Audit And Standards](2026-06-17-phase-0-ooad-audit-and-standards.md)
2. [Phase 1 -- Complete Chapter 2](2026-06-17-phase-1-complete-chapter-2.md)
3. [Phase 2 -- Complete Chapter 3](2026-06-17-phase-2-complete-chapter-3.md)
4. [Phase 3 -- Complete Chapter 4](2026-06-17-phase-3-complete-chapter-4.md)
5. [Phase 4 -- Final Review](2026-06-17-phase-4-final-review.md)

## Global Rules

- Source of truth bắt buộc: `report/report_support_documentation/project_source_of_truth.md`.
- Quy trình OOAD bắt buộc: `report/report_support_documentation/OOAD.md`.
- Danh sách diagram ưu tiên: `report/report_support_documentation/ooad_diagram_priority_list.md`.
- Quy tắc UML bổ sung: `report/report_support_documentation/formal-17-12-05.pdf`.
- Checklist diagram: `report/report_support_documentation/diagram_type_rubric.md`, `report/report_support_documentation/diagram_review_checklist.md`.
- Quy tắc repo: `AGENTS.md`.
- Quy tắc hình thức đồ án từ `AGENTS.md` phải được áp dụng: tài liệu tham khảo đã liệt kê phải được trích dẫn; dữ liệu từ website phải trích dẫn trực tiếp tại vị trí sử dụng; hình không tự tạo phải có nguồn; hình tự tạo ghi `Nguồn: Tác giả xây dựng`; không viết lỗi chính tả, không giải thích thuật ngữ tiếng Anh theo kiểu máy móc; bố cục các chương/mục cần cân đối.
- Văn phong báo cáo phải ngắn gọn, đúng trọng tâm, súc tích và đủ ý. Không viết dài dòng, không lặp lại cùng một ý ở nhiều đoạn, không đưa chi tiết triển khai vào phần phân tích nếu chưa cần.
- Khi cần bổ sung cơ sở lý thuyết về system design, OOAD, UML, requirements, architecture hoặc software design, ưu tiên nguồn chính thống như tài liệu chuẩn, sách/giáo trình, tài liệu trường đại học, OMG UML, ISO/IEC/IEEE, SEI/CMU hoặc tài liệu nhà cung cấp chính thức. Nội dung lấy từ Internet phải được trích dẫn đầy đủ theo quy tắc đồ án.
- Dùng `plantuml-skill` cho các diagram mới sau này; chỉ dùng draw.io nếu PlantUML vẫn không đọc được sau khi đã tách nhỏ hoặc đơn giản hóa.
- Nếu dùng public Kroki, ghi rõ source `.puml` đã được upload lên `kroki.io`.
- Không nhồi quá nhiều thành phần vào diagram tổng quan. Nếu rối, giữ overview ở mức khái quát và tách thành diagram nhỏ theo nhóm nghiệp vụ, use case hoặc lớp trách nhiệm.
- Không mô tả hệ thống là nơi tạo, huấn luyện, triển khai hoặc điều phối AI Agent.
- Không mô tả AI Agent dùng trực tiếp provider API key.
- Hình tự tạo phải có caption kết thúc bằng `Nguồn: Tác giả xây dựng`.
- Sau mỗi phase, build từ `report/`:

```bash
latexmk -pdf -g main.tex
```

Expected: `Output written on main.pdf`, không có lỗi fatal.
