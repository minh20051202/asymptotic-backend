# Hướng dẫn đánh giá chất lượng diagram trong đồ án OOAD

Tài liệu này dùng để kiểm tra chất lượng các biểu đồ trước khi đưa vào báo cáo. Mục tiêu không chỉ là biểu đồ "đẹp", mà phải **đúng quy trình OOAD, đúng nội dung nghiệp vụ, đúng mức trừu tượng và có khả năng truy vết** về yêu cầu hệ thống.

Quy tắc ký pháp UML chuẩn được tổng hợp tại `uml_2_5_1_drawing_rules.md`. Khi đánh giá, phải kiểm tra đúng ngữ nghĩa UML trước khi tối ưu bố cục.

## 1. Nguyên tắc chung

Một diagram đạt yêu cầu cần thỏa mãn các tiêu chí sau:

- **Đúng mục đích:** biểu đồ phải trả lời đúng câu hỏi của loại biểu đồ đó.
- **Đúng mức trừu tượng:** biểu đồ phân tích không lẫn chi tiết triển khai; biểu đồ thiết kế không quá mơ hồ.
- **Đúng phạm vi:** chỉ vẽ những thành phần nằm trong phạm vi đang mô tả.
- **Đúng thuật ngữ:** tên actor, use case, class, package, trạng thái phải thống nhất với source of truth.
- **Có traceability:** biểu đồ phải liên kết được với requirement, use case hoặc quyết định thiết kế tương ứng.
- **Dễ đọc:** bố cục rõ, ít đường cắt nhau, tên gọi ngắn gọn, không nhồi quá nhiều nội dung.
- **Đúng quy tắc nguồn:** hình từ nguồn bên ngoài phải trích dẫn; hình tự thiết kế không cần dòng nguồn.
- **Đúng loại quy tắc:** phân biệt UML chuẩn, quy tắc phương pháp và quy ước dự án.

## 2. Câu hỏi đánh giá nhanh

Trước khi chấp nhận một diagram, trả lời các câu hỏi sau:

1. Biểu đồ này phục vụ chương/mục nào?
2. Biểu đồ này mô tả yêu cầu, phân tích, thiết kế hay triển khai?
3. Các actor/use case/class/package có đúng tên chuẩn không?
4. Có thành phần nào nằm ngoài phạm vi đề tài không?
5. Có chi tiết kỹ thuật xuất hiện quá sớm không?
6. Người đọc có hiểu được thông điệp chính trong 30 giây không?
7. Biểu đồ có quá nhiều node hoặc đường nối không?
8. Biểu đồ có liên kết được với FR/NFR hoặc use case nào không?
9. Caption và quy tắc nguồn hình đã đúng chưa?
10. File source của biểu đồ có được lưu cùng output không?

## 3. Mức điểm đánh giá

Có thể chấm mỗi diagram theo thang 5 điểm:

| Điểm | Ý nghĩa |
|---|---|
| 5 | Rất tốt: đúng nội dung, đúng OOAD, dễ đọc, trace rõ |
| 4 | Tốt: đúng phần lớn, chỉ cần chỉnh nhỏ |
| 3 | Chấp nhận được: truyền tải được ý chính nhưng còn thiếu hoặc rối |
| 2 | Yếu: sai mức trừu tượng, thiếu thành phần quan trọng hoặc khó đọc |
| 1 | Không đạt: sai mục đích biểu đồ hoặc không khớp hệ thống |

## 4. Tiêu chí chung theo nhóm điểm

### 4.1 Tính đúng đắn nội dung

- Actor, hệ thống ngoài, class, trạng thái, gói phải đúng với source of truth.
- Không mô tả hệ thống như nơi tạo/huấn luyện/điều phối AI Agent.
- Không để AI Agent dùng trực tiếp provider API key.
- Gateway phải là nơi kiểm soát identity, policy, quota, budget và provider routing.
- Luồng tài chính phải thể hiện kiểm soát trước khi phát sinh chi phí.

### 4.2 Tính đúng đắn OOAD

- Use case diagram mô tả mục tiêu của actor, không mô tả API endpoint.
- Activity diagram mô tả luồng xử lý/nghiệp vụ, không thay thế sequence diagram.
- Analysis class diagram tập trung vào Boundary-Control-Entity, không nhồi repository/framework.
- Design class diagram có thể có service, repository, adapter, interface.
- Sequence diagram phải thể hiện thứ tự tương tác theo thời gian.
- State machine diagram phải có trạng thái, sự kiện chuyển trạng thái và trạng thái kết thúc.
- Robustness/BCE là quy ước phương pháp, không được trình bày như loại diagram chuẩn riêng của UML.
- ERD và biểu đồ kết quả kiểm thử không được gọi là UML diagram.

### 4.3 Tính nhất quán

- Tên actor trong diagram phải khớp với tài liệu actor.
- Tên use case phải khớp với danh sách use case chuẩn.
- Tên FR/NFR trong mô tả hoặc trace phải khớp với source of truth.
- Cùng một khái niệm không được đổi tên tùy ý giữa các diagram.

### 4.4 Tính dễ đọc

- Mỗi diagram nên có một thông điệp chính.
- Nếu diagram quá lớn, nên tách thành diagram tổng quát và diagram phân rã.
- Tránh đường nối chồng chéo hoặc cắt nhau quá nhiều.
- Nhóm thành phần liên quan gần nhau.
- Dùng màu hoặc stereotype nhất quán nếu có.
- Không dùng quá nhiều kiểu mũi tên trong cùng một biểu đồ.

### 4.5 Tính truy vết

Mỗi diagram nên trả lời được:

- Nó hỗ trợ use case nào?
- Nó liên quan FR/NFR nào?
- Nó được dùng ở chương/mục nào?
- Nó có ảnh hưởng đến class/sequence/state diagram nào khác không?

## 5. Checklist trước khi đưa diagram vào LaTeX

- [ ] File nguồn PlantUML (`.puml`) đã được lưu cạnh file export.
- [ ] File ảnh export rõ nét (`.png`, `.svg`, hoặc `.pdf`).
- [ ] Tên file không gây lỗi LaTeX hoặc đã được include đúng cách.
- [ ] Caption có tên biểu đồ rõ ràng.
- [ ] Hình từ nguồn bên ngoài có trích dẫn; hình tự thiết kế không có dòng nguồn thừa.
- [ ] Label LaTeX đặt đúng và không trùng.
- [ ] Diagram được nhắc tới trong nội dung trước hoặc sau hình.
- [ ] Diagram không bị tràn trang hoặc quá nhỏ khi render PDF.
- [ ] Nội dung diagram khớp source of truth.

## 6. Quy tắc đặt tên file

Khuyến nghị:

- Use case tổng quát: `use_case_overview.puml`, `use_case_overview.png`
- Activity UC01: `UC01_activity.puml`, `UC01_activity.png`
- Sequence UC01: `UC01_sequence.puml`, `UC01_sequence.png`
- Analysis class tổng quát: `analysis_class_overview.puml`
- State machine: `ai_request_state.puml`, `financial_transaction_state.puml`

Tránh:

- tên file quá dài;
- tên file có nhiều dấu cách;
- tên file trùng nhau trong cùng phạm vi;
- đổi tên file ảnh mà không cập nhật LaTeX.

Tên chung như `diagram.puml` hoặc `sequence.puml` được phép lặp nếu mỗi cặp source/export nằm trong thư mục use case hoặc loại biểu đồ riêng, và đường dẫn LaTeX không mơ hồ.

## 7. Quy tắc theo vị trí chương

- **Chương 3:** chỉ mô hình phân tích; không có repository, adapter, framework hoặc deployment node.
- **Chương 4:** mô hình thiết kế; Component/Deployment nếu có phải ghi là as-designed.
- **Chương 5:** minh chứng as-built và kiểm thử; không dùng kiến trúc dự kiến thay cho hệ thống đã triển khai.
- Chương 5 không bắt buộc có UML diagram khi chưa có dữ liệu triển khai.

## 8. Khi nào cần vẽ lại diagram?

Nên vẽ lại hoặc chỉnh diagram nếu:

- source of truth thay đổi;
- actor/use case/FR/NFR đổi tên;
- use case được tách hoặc gộp;
- diagram đang mô tả sai phạm vi hệ thống;
- diagram có chi tiết implementation trong phần analysis;
- diagram không còn khớp với nội dung chương;
- ảnh bị mờ hoặc không đọc được trong PDF.
