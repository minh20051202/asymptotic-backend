# Report Support Documentation

Thư mục này chứa tài liệu hỗ trợ trực tiếp cho báo cáo. Các file LaTeX chính vẫn nằm tại `report/`.

## Cấu trúc

- `governance/`: nguồn sự thật, quy trình OOAD và cấu trúc báo cáo.
- `guidelines/diagrams/`: quy tắc, rubric và checklist vẽ biểu đồ.
- `guidelines/use_cases/`: hướng dẫn phân tích và đặc tả use case.
- `references/standards/`: tiêu chuẩn chính thức.
- `references/lectures/`: tài liệu bài giảng.
- `references/sample_reports/`: báo cáo mẫu dùng để tham khảo cách trình bày.
- `references/examples/`: hình và ví dụ minh họa từ nguồn ngoài.
- `diagrams/chapter_2/`: Use Case và Activity Diagram.
- `diagrams/chapter_3/`: Analysis Class và Robustness Diagram.
- `diagrams/chapter_4/`: Package, Design Class, Sequence, State và các thiết kế bổ sung đã được chốt.
- `diagrams/chapter_5/`: biểu đồ as-built và minh chứng trực quan khi có dữ liệu triển khai.
- `governance/team_budget_hierarchy_decision.md`: decision record của mô hình ngân sách đã chấp thuận.
- `audits/`: kết quả rà soát báo cáo và đề xuất dọn dẹp.

## Quy tắc

1. Đọc `governance/project_source_of_truth.md` trước khi sửa nội dung báo cáo hoặc biểu đồ.
2. Giữ `.puml` cạnh ảnh render.
3. Research brief nằm tại `report/research/`, không đặt lẫn vào guideline.
4. Quyết định đã chấp thuận phải được ghi vào `governance/` và đồng bộ với source of truth.
5. Không giữ file `.bak`, output LaTeX cũ hoặc bản sao biểu đồ không còn dùng; Git lưu lịch sử thay đổi.
6. Bắt đầu công việc diagram từ [`guidelines/diagrams/README.md`](guidelines/diagrams/README.md).
7. Không gọi Robustness, ERD hoặc biểu đồ kết quả kiểm thử là loại UML chuẩn khi chúng không thuộc phân loại đó.
