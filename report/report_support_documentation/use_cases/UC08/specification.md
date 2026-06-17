# UC08 - Quản lý lập trình viên

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC08` |
| Tên | Quản lý lập trình viên |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Quản trị viên tổ chức muốn quản lý lập trình viên để kiểm soát thành viên, quyền truy cập và ngân sách trong tổ chức. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Quản trị viên tổ chức |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Quản trị viên tổ chức yêu cầu thay đổi danh sách lập trình viên. |
| Pre-Conditions | Quản trị viên tổ chức đã đăng nhập; tổ chức đang hoạt động. |
| Success Post-Conditions | Tài khoản lập trình viên và quyền truy cập liên quan được cập nhật theo yêu cầu. |
| Failure Post-Conditions | Không thay đổi tài khoản hoặc ngân sách nếu yêu cầu không hợp lệ. |

## Basic Flow

1. Quản trị viên tổ chức mở chức năng quản lý lập trình viên.
2. Gateway hiển thị danh sách lập trình viên và trạng thái hiện tại.
3. Quản trị viên chọn thao tác tạo lập trình viên mới.
4. Quản trị viên nhập thông tin lập trình viên.
5. Gateway kiểm tra thông tin tài khoản và phạm vi tổ chức.
6. Gateway tạo tài khoản lập trình viên, gắn với tổ chức và thiết lập hạn mức ban đầu.
7. Gateway thông báo kết quả và cập nhật danh sách thành viên.

## Alternative Flows

### 2a - Xem chi tiết lập trình viên

1. Nếu cần tra cứu, quản trị viên chọn một lập trình viên trong danh sách.
2. Gateway hiển thị thông tin tài khoản, hạn mức, Agent liên quan và trạng thái API key.
3. Use Case kết thúc thành công.

### 3a - Cập nhật thông tin lập trình viên

1. Nếu cần thay đổi thông tin, quản trị viên chọn lập trình viên cần cập nhật.
2. Quản trị viên nhập thông tin cần thay đổi.
3. Gateway kiểm tra dữ liệu và cập nhật thông tin lập trình viên.
4. Use Case kết thúc thành công.

### 3b - Vô hiệu hóa lập trình viên

1. Nếu cần ngừng quyền sử dụng, quản trị viên chọn lập trình viên cần vô hiệu hóa.
2. Gateway ngăn lập trình viên thực hiện thao tác mới.
3. Gateway thực hiện UC05.3 để thu hồi các API key liên quan.
4. Gateway thu hồi phần ngân sách khả dụng về ví tổ chức.
5. Use Case kết thúc thành công.

## Exception Flows

### 5b - Email đã tồn tại hoặc dữ liệu không hợp lệ

1. Gateway từ chối yêu cầu và thông báo nguyên nhân.
2. Use Case kết thúc không thành công.

### 3c - Lập trình viên còn giao dịch chưa quyết toán

1. Khi vô hiệu hóa lập trình viên còn giao dịch chưa quyết toán, Gateway chặn yêu cầu mới và chuyển phần ngân sách liên quan sang chờ xử lý.
2. Use Case kết thúc không thành công một phần.

## Business Rules

- `BR-UC08-01`: Lập trình viên thuộc đúng một tổ chức trong phạm vi hệ thống.
- `BR-UC08-02`: Vô hiệu hóa lập trình viên phải ngăn phát sinh yêu cầu AI mới.
- `BR-UC08-03`: Ngân sách đang tạm giữ chỉ được thu hồi sau khi giao dịch liên quan được quyết toán.

## Non-Functional Requirements

- `NFR02`: Bảo mật khóa truy cập.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.

## Traceability

- Functional Requirements: `FR01, FR03, FR05, FR09`.
- Diagram: `diagram.drawio`, `diagram.png`.
