# UC01 - Thực hiện yêu cầu AI

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC01` |
| Tên | Thực hiện yêu cầu AI |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | AI Agent muốn gửi yêu cầu qua Gateway để nhận kết quả AI theo luồng streaming trong phạm vi ngân sách được cấp. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | AI Agent |
| Supporting Actor | Nhà cung cấp dịch vụ AI |
| Priority | Must Have |
| Trigger | AI Agent gửi yêu cầu AI tới Gateway. |
| Pre-Conditions | AI Agent đã được cấp API key của Gateway; chính sách ngân sách, danh mục mô hình, nhà cung cấp và biểu giá đã được thiết lập. |
| Success Post-Conditions | Kết quả đã được truyền cho AI Agent; chi phí thực tế đã được quyết toán; phần tạm giữ dư đã được giải phóng; nhật ký thực thi và bản ghi lũy đẳng đã hoàn tất. |
| Failure Post-Conditions | Không thực thi hoặc tính phí trùng; khoản tạm giữ được giải phóng hoặc chuyển sang trạng thái chờ xử lý; lỗi được lưu để thử lại hoặc đối soát. |

## Basic Flow

1. AI Agent gửi API key, khóa lũy đẳng, nội dung tác vụ và các tùy chọn thực thi.
2. Gateway xác thực yêu cầu và kiểm tra quyền thực thi.
3. Gateway kiểm tra khóa lũy đẳng và ghi nhận yêu cầu mới.
4. Gateway chọn phương án thực thi, kiểm tra chính sách và ngân sách liên quan.
5. Gateway tạm giữ ngân sách cần thiết cho yêu cầu.
6. Gateway dùng provider credential nội bộ để gọi nhà cung cấp dịch vụ AI và truyền từng phần kết quả cho AI Agent.
7. Gateway xác định mức sử dụng thực tế sau khi luồng kết quả kết thúc.
8. Gateway quyết toán chi phí, giải phóng phần còn dư và hoàn tất bản ghi lũy đẳng.

## Alternative Flows

### 3a - Yêu cầu đã được xử lý

1. Gateway trả kết quả đã lưu và không gọi lại nhà cung cấp.
2. Use Case kết thúc thành công.

### 4a - Chọn phương án thay thế

1. Gateway chọn mô hình hoặc nhà cung cấp thay thế theo chính sách.
2. Use Case tiếp tục tại bước 5.

## Exception Flows

### 2b - Yêu cầu hoặc quyền truy cập không hợp lệ

1. Gateway từ chối yêu cầu; không phát sinh tạm giữ ngân sách.
2. Use Case kết thúc không thành công.

### 4b - Không đủ ngân sách và không có phương án thay thế

1. Gateway xác định ngân sách hiện tại không đủ cho yêu cầu và không có mô hình hoặc nhà cung cấp thay thế phù hợp.
2. Gateway từ chối yêu cầu; không gọi nhà cung cấp dịch vụ AI và không phát sinh chi phí.
3. Use Case kết thúc không thành công.

### 6b - Luồng xử lý không hoàn tất

1. Gateway ghi nhận trạng thái lỗi, quyết toán phần chi phí đã phát sinh nếu có và chuyển giao dịch sang trạng thái có thể phục hồi hoặc đối soát.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC01-01`: Khóa lũy đẳng là duy nhất theo AI Agent.
- `BR-UC01-02`: Không tự động thử lại sau khi streaming đã bắt đầu.
- `BR-UC01-03`: Gateway chỉ gọi nhà cung cấp sau khi ngân sách liên quan đáp ứng chính sách.
- `BR-UC01-04`: AI Agent chỉ sử dụng API key do Gateway cấp; provider credential nội bộ không được trả về cho AI Agent.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR02`: Bảo mật khóa truy cập.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.
- `NFR05`: Hiệu năng.

## Traceability

- Functional Requirements: `FR03, FR04, FR05, FR06, FR07, FR08, FR09, FR10`.
- Diagram: `diagram.drawio`, `diagram.png`.
