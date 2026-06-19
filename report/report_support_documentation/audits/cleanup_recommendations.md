# Kết quả dọn dẹp tài liệu hỗ trợ

Ngày rà soát: 2026-06-19.

## Đã giữ

- Toàn bộ file LaTeX chính tại `report/`.
- `governance/project_source_of_truth.md`, `governance/OOAD.md` và `governance/report_structure.md`.
- Toàn bộ guideline đang được `AGENTS.md` tham chiếu.
- PlantUML và ảnh render đang được `chapter_2.tex`, `chapter_3.tex`, `chapter_4.tex` sử dụng.
- Đặc tả UC01--UC09.
- UML 2.5.1, bài giảng OOAD, báo cáo mẫu và ảnh minh họa có nguồn.
- Research brief có provenance.
- `report/main.pdf`.

## Đã sửa hoặc hợp nhất

- `audits/chapter_2_3_4_completion_audit.md` được cập nhật theo đúng trạng thái hiện tại, không xem phần chưa có biểu đồ là đã hoàn thành.
- Nội dung đề xuất nâng cấp Chương 1 đã được đối chiếu với `chapter_1.tex` và source of truth; proposal trùng lặp đã được xóa.
- Quyết định phân bổ ngân sách được rút gọn thành `governance/team_budget_hierarchy_decision.md`; proposal cũ đã được xóa.
- Research riêng về BCE/Robustness được hợp nhất vào `report/research/ooad/ooad-diagram-rules-multisource-research-2026-06-18.md`.
- Research thị trường đã có provenance riêng đặt cạnh brief.

## Đã xóa

- Output LaTeX sinh tự động trong `report/build/` sau khi build cuối thành công.
- LaTeX sinh cũ không được `main.tex` import.
- Toàn bộ file `.bak`, ảnh đánh số cũ và bản sao PlantUML/PNG không còn được tham chiếu.
- Các thư mục legacy không còn cần thiết; lịch sử được giữ bằng Git.

## Không xóa

- PDF tiêu chuẩn, bài giảng và báo cáo mẫu.
- Ảnh ví dụ Thinhnotes đang dùng làm tài liệu nghiên cứu.
- Research brief và provenance đang được plan hoặc guideline tham chiếu.
- PlantUML và ảnh render đang được báo cáo tham chiếu.

## Kiểm chứng

1. Không còn tham chiếu hoạt động tới file đã xóa.
2. `latexmk -pdf -g main.tex` thành công.
3. `report/main.pdf` vẫn được tạo tại vị trí chuẩn.
4. Các quyết định nghiệp vụ và nguồn trích dẫn còn bản canonical.

Khoảng trống mô hình hóa còn lại được theo dõi trong `chapter_2_3_4_completion_audit.md`; chúng không phải tài liệu rác và không thuộc phạm vi dọn dẹp.
