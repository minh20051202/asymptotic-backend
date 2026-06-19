# UC01 - Thực hiện yêu cầu AI qua Gateway

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC01` |
| Tên | Thực hiện yêu cầu AI qua Gateway |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | AI Agent gửi yêu cầu qua Gateway để nhận kết quả AI trong phạm vi quyền, chính sách và ngân sách được cấp. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | AI Agent |
| Supporting Actor | Nhà cung cấp dịch vụ AI |
| Priority | Must Have |
| Trigger | AI Agent gửi yêu cầu AI tới Gateway. |
| Pre-Conditions | Agent đang hoạt động, có API key hợp lệ, thuộc đường dẫn tổ chức–đội ngũ–lập trình viên–Agent hợp lệ; nhà cung cấp, mô hình và biểu giá đã được cấu hình. |
| Success Post-Conditions | Agent nhận được kết quả AI; usage, cost, trace và trạng thái ngân sách liên quan được ghi nhận nhất quán. |
| Failure Post-Conditions | Agent nhận được thông báo lỗi; yêu cầu không bị tính phí trùng và trạng thái tài chính vẫn có thể kiểm tra, đối soát. |

## Basic Flow

1. AI Agent gửi API key, khóa lũy đẳng, nội dung yêu cầu và các tùy chọn sử dụng.
2. Gateway xác thực Agent và xác định tổ chức, đội ngũ, lập trình viên chịu chi phí.
3. Gateway kiểm tra yêu cầu có trùng với yêu cầu đã xử lý hay không.
4. Gateway kiểm tra trạng thái Agent, policy, quota và ngân sách khả dụng.
5. Gateway chọn nhà cung cấp, mô hình phù hợp và chuyển tiếp yêu cầu bằng credential nội bộ.
6. Nhà cung cấp dịch vụ AI xử lý và trả kết quả cùng thông tin sử dụng.
7. Gateway ghi nhận usage, cost, trace và cập nhật trạng thái ngân sách.
8. Gateway trả kết quả AI cho Agent.

## Alternative Flows

### 3a - Yêu cầu đã được xử lý

1. Gateway trả kết quả đã lưu và không gọi lại nhà cung cấp.
2. Use Case kết thúc thành công.

### 5a - Chọn phương án thay thế

1. Gateway chọn mô hình hoặc nhà cung cấp thay thế theo chính sách.
2. Use Case tiếp tục tại bước 6.

## Exception Flows

### 2b - Agent hoặc API key không hợp lệ

1. Gateway từ chối yêu cầu và thông báo nguyên nhân.
2. Use Case kết thúc không thành công.

### 4b - Vi phạm policy, vượt quota hoặc không đủ ngân sách

1. Gateway từ chối yêu cầu trước khi gọi nhà cung cấp.
2. Gateway trả lý do từ chối cho Agent.
3. Use Case kết thúc không thành công.

### 6b - Nhà cung cấp xử lý thất bại

1. Gateway ghi nhận trace lỗi và trạng thái chi phí nếu có.
2. Gateway trả thông báo lỗi cho Agent.
3. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC01-01`: Khóa lũy đẳng là duy nhất theo AI Agent.
- `BR-UC01-02`: Yêu cầu trùng khóa lũy đẳng không được tạo thêm chi phí.
- `BR-UC01-03`: Gateway chỉ gọi nhà cung cấp sau khi ngân sách liên quan đáp ứng chính sách.
- `BR-UC01-04`: AI Agent chỉ sử dụng API key do Gateway cấp; provider credential nội bộ không được trả về cho AI Agent.
- `BR-UC01-05`: Mỗi request phải xác định duy nhất đường dẫn chịu chi phí gồm tổ chức, đội ngũ, lập trình viên và Agent.
- `BR-UC01-06`: Việc cập nhật chi phí và hạn mức phải nhất quán tại tất cả các cấp.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR02`: Bảo mật khóa truy cập.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.
- `NFR05`: Hiệu năng.

## Traceability

- Functional Requirements: `FR03, FR04, FR05, FR06, FR07, FR08, FR09, FR10`.
- Diagram: `diagram.puml`, `diagram.png`.
