# UC04 - Đăng ký và quản lý tác nhân AI

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC04` |
| Tên | Đăng ký và quản lý tác nhân AI |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Lập trình viên muốn đăng ký và quản lý tác nhân AI để hệ thống có thể kiểm soát trạng thái, hạn mức và chi phí phát sinh từ các yêu cầu AI. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Lập trình viên |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Lập trình viên mở chức năng đăng ký và quản lý tác nhân AI. |
| Pre-Conditions | Lập trình viên đã đăng nhập và thuộc một tổ chức đang hoạt động. |
| Success Post-Conditions | Tác nhân AI được đăng ký hoặc cấu hình quản lý được cập nhật; trạng thái và chính sách hạn mức sẵn sàng cho UC01 kiểm soát yêu cầu. |
| Failure Post-Conditions | Chính sách hiện tại không thay đổi nếu yêu cầu không hợp lệ. |

## Basic Flow

1. Lập trình viên mở chức năng đăng ký và quản lý tác nhân AI.
2. Lập trình viên yêu cầu đăng ký tác nhân AI mới.
3. Lập trình viên nhập tên, mô tả và phạm vi sử dụng dự kiến của tác nhân.
4. Gateway kiểm tra thông tin đăng ký và gắn tác nhân với tổ chức của lập trình viên.
5. Gateway tạo bản ghi tác nhân AI ở trạng thái hoạt động ban đầu.
6. Gateway thông báo đăng ký thành công và cho phép lập trình viên tiếp tục thiết lập hạn mức hoặc API key.

## Alternative Flows

### 2a - Xem thông tin tác nhân AI

1. Nếu chỉ cần tra cứu, lập trình viên xem thông tin đăng ký, hạn mức, phần đã sử dụng và trạng thái hiện tại của tác nhân AI.
2. Use Case kết thúc thành công.

### 3a - Thiết lập hạn mức tác nhân AI

1. Nếu cần kiểm soát chi phí, lập trình viên chọn một tác nhân AI đã đăng ký.
2. Gateway hiển thị hạn mức khả dụng của lập trình viên và chính sách hiện tại của tác nhân.
3. Lập trình viên nhập hạn mức và chu kỳ áp dụng.
4. Gateway kiểm tra hạn mức mới trong phạm vi ngân sách của lập trình viên.
5. Gateway lưu chính sách hạn mức mới và thông báo kết quả cho lập trình viên.
6. Use Case kết thúc thành công.

### 3b - Cập nhật chu kỳ ngân sách

1. Nếu cần thay đổi chu kỳ áp dụng, lập trình viên chọn thao tác cập nhật chu kỳ ngân sách.
2. Lập trình viên nhập chu kỳ mới.
3. Gateway kiểm tra chu kỳ hợp lệ và lưu chính sách mới.
4. Use Case kết thúc thành công.

### 3c - Tạm dừng tác nhân AI

1. Nếu cần dừng phát sinh chi phí mới, lập trình viên chọn tác nhân AI sau khi xem trạng thái hạn mức hiện tại.
2. Lập trình viên chọn thao tác tạm dừng tác nhân AI.
3. Gateway cập nhật trạng thái tác nhân và chặn các yêu cầu AI mới từ tác nhân này.
4. Use Case kết thúc thành công.

### 5a - Hạn mức mới thấp hơn mức đã sử dụng

1. Gateway ghi nhận chính sách mới nhưng chặn các yêu cầu mới vượt ngưỡng trong chu kỳ hiện tại.
2. Use Case kết thúc thành công.

## Exception Flows

### 5b - Hạn mức vượt phạm vi được cấp

1. Gateway từ chối yêu cầu và giữ nguyên chính sách hiện tại.
2. Use Case kết thúc không thành công.

### 2b - Tác nhân không thuộc quyền quản lý

1. Gateway từ chối truy cập.
2. Use Case kết thúc không thành công.

### 4b - Thông tin tác nhân không hợp lệ

1. Gateway từ chối đăng ký tác nhân và thông báo trường dữ liệu cần điều chỉnh.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC04-01`: Hạn mức của tác nhân không được vượt phần ngân sách lập trình viên được phép quản lý.
- `BR-UC04-02`: Chính sách hạn mức có hiệu lực với các yêu cầu AI tiếp theo.
- `BR-UC04-03`: Gateway phải chặn yêu cầu mới khi tác nhân vượt hạn mức.
- `BR-UC04-04`: Hạn mức của tác nhân chỉ được tiêu thụ thông qua UC01.
- `BR-UC04-05`: Mỗi tác nhân AI phải thuộc một tổ chức và có trạng thái quản lý rõ ràng trước khi được cấp API key.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.

## Traceability

- Functional Requirements: `FR02, FR05, FR09`.
- Diagram: `diagram.drawio`, `diagram.png`.
