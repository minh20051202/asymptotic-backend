# Rà soát mức độ hoàn thiện Chương 2--4

Ngày rà soát: 2026-06-19.

Tài liệu này phản ánh trạng thái tạo tác hiện có. Tiêu chí được đối chiếu từ `AGENTS.md`, source of truth, guideline biểu đồ và các kế hoạch hoàn thiện Chương 2--4.

## Chương 2

Trạng thái: Cơ bản hoàn thiện.

- Đã có khảo sát hiện trạng, FR01--FR10, NFR01--NFR05 và phân tích actor.
- Đã có Use Case Diagram tổng quát, diagram chi tiết và đặc tả UC01--UC09.
- Đã có Activity Diagram cho UC01, đăng ký AI Agent và nạp tiền/phân bổ ngân sách.
- Các use case và yêu cầu dùng thuật ngữ canonical, phân biệt đăng ký Agent với quản lý API key.

Việc còn lại chủ yếu là rà soát hình thức, chất lượng ảnh và cảnh báo LaTeX sau mỗi lần sửa.

## Chương 3

Trạng thái: Hoàn thiện một phần.

Đã có:

- Cơ sở phân tích hướng đối tượng và phương pháp Boundary--Control--Entity.
- Analysis Class Diagram tổng quát.
- Ba góc nhìn hiện hành theo Boundary, Control và Entity.
- Robustness Diagram cho UC01 và UC04.
- Bảng đối chiếu UC01--UC09 với các lớp phân tích.

Chưa có theo Phase 2:

- Danh mục lớp và tài liệu dẫn xuất lớp dùng chung.
- Sáu Analysis Class Diagram tập trung: UC01; UC02--UC03; UC04--UC05; UC07--UC08; UC06; UC09.

Ba góc nhìn Boundary/Control/Entity hiện tại không thay thế hoàn toàn sáu góc nhìn theo nhóm nghiệp vụ. Khi bổ sung, phải giữ Analysis Class Diagram tổng quát thưa, đồng thời bảo đảm tên lớp, stereotype và ý nghĩa quan hệ nhất quán giữa các hình.

## Chương 4

Trạng thái: Hoàn thiện một phần.

Đã có:

- Kiến trúc modular monolith và thiết kế package.
- Design Class Diagram tổng quát, luồng Gateway và Finance/Ledger.
- Sequence Diagram cho UC01, UC04 và luồng nạp tiền/phân bổ ngân sách.
- State Machine cho AI Request và Financial Transaction.
- Mô tả dạng văn bản cho nhóm dữ liệu và nhóm API.

Chưa có theo Phase 3:

- Tài liệu dẫn xuất thực thể dữ liệu.
- Logical ERD theo các nhóm Organization/Access, Agent/API key, Finance/Budget, Provider/Catalog và Request/Trace.
- Góc nhìn ERD tổng quát chỉ cần tạo nếu vẫn đọc được ở kích thước trang báo cáo.

Phần dữ liệu hiện tại mới mô tả nhóm bảng, chưa đủ thay thế Logical ERD có định danh, quan hệ, cardinality, optionality và ràng buộc nghiệp vụ.

## Danh sách biểu đồ hiện có

- Chương 2: một Use Case Diagram tổng quát, chín Use Case Diagram chi tiết và ba Activity Diagram.
- Chương 3: một Analysis Class Diagram tổng quát, ba góc nhìn BCE và hai Robustness Diagram.
- Chương 4: một Package Diagram, ba Design Class Diagram, ba Sequence Diagram và hai State Machine Diagram.

## Ưu tiên tiếp theo

1. Hoàn thiện sáu Analysis Class Diagram tập trung và tài liệu dẫn xuất lớp.
2. Hoàn thiện Logical ERD theo nhóm nghiệp vụ và tài liệu dẫn xuất thực thể.
3. Build lại báo cáo, rà tham chiếu, caption, độ đọc của hình và cảnh báo LaTeX.

## Kiểm chứng

- Build ngày 2026-06-19 bằng `latexmk -pdf -g main.tex` thành công.
- `report/main.pdf` có 72 trang, khổ A4.
- Không có lỗi LaTeX, citation không xác định hoặc cross-reference không xác định.
- Còn cảnh báo `Overfull` và `Underfull` về dàn trang; cần tiếp tục rà khi bổ sung các biểu đồ còn thiếu.

Không xem một chương là hoàn thiện chỉ vì đã có văn bản thay cho tạo tác mô hình bắt buộc.
