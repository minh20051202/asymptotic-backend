# UC04 - Đăng ký và quản lý AI Agent

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC04` |
| Tên | Đăng ký và quản lý AI Agent |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Lập trình viên đăng ký AI Agent bên ngoài do mình phát triển hoặc quản lý, đồng thời quản lý thông tin và trạng thái của Agent. Quản trị viên tổ chức giám sát và bàn giao Agent khi cần. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Lập trình viên |
| Supporting Actor | Quản trị viên tổ chức |
| Priority | Must Have |
| Trigger | Lập trình viên hoặc quản trị viên tổ chức mở chức năng đăng ký và quản lý AI Agent. |
| Pre-Conditions | Người dùng đã đăng nhập và có quyền quản lý Agent; tổ chức đang hoạt động; lập trình viên quản lý thuộc một đội ngũ hợp lệ. |
| Success Post-Conditions | Agent được đăng ký hoặc thông tin quản lý được cập nhật; Agent có đường dẫn tổ chức–đội ngũ–lập trình viên rõ ràng. |
| Failure Post-Conditions | Không tạo hoặc thay đổi Agent khi dữ liệu, quyền hoặc quan hệ quản lý không hợp lệ. |

## Basic Flow

1. Lập trình viên mở chức năng đăng ký và quản lý AI Agent.
2. Lập trình viên yêu cầu đăng ký Agent mới.
3. Lập trình viên nhập tên, mô tả và phạm vi sử dụng của Agent.
4. Gateway kiểm tra quyền, trạng thái và quan hệ đội ngũ của lập trình viên.
5. Gateway tạo Agent, gắn với tổ chức và chính lập trình viên đăng ký ở trạng thái chưa kích hoạt.
6. Gateway thông báo đăng ký thành công.

## Alternative Flows

### 2a - Xem thông tin Agent

1. Người dùng có quyền chọn Agent trong phạm vi quản lý.
2. Gateway hiển thị thông tin, lập trình viên quản lý, đội ngũ, trạng thái, hạn mức tổng quan và API key ở dạng che giấu.
3. Use Case kết thúc thành công.

### 2b - Cập nhật thông tin Agent

1. Người dùng có quyền chọn Agent và nhập thông tin cần thay đổi.
2. Gateway kiểm tra dữ liệu và quyền quản lý.
3. Gateway cập nhật thông tin không thuộc ngân sách hoặc API key.
4. Use Case kết thúc thành công.

### 2c - Kích hoạt Agent

1. Người dùng có quyền yêu cầu kích hoạt Agent.
2. Gateway kiểm tra Agent có lập trình viên quản lý, đường dẫn ngân sách hợp lệ, hạn mức phù hợp và ít nhất một API key hoạt động.
3. Gateway chuyển Agent sang trạng thái hoạt động.
4. Use Case kết thúc thành công.

### 2d - Tạm dừng hoặc vô hiệu hóa Agent

1. Người dùng có quyền chọn Agent cần dừng.
2. Gateway chặn yêu cầu mới và cập nhật trạng thái Agent.
3. Use Case kết thúc thành công.

### 2e - Bàn giao Agent

1. Quản trị viên tổ chức chọn lập trình viên quản lý mới.
2. Gateway kiểm tra lập trình viên mới và xử lý hạn mức hiện tại theo UC03.
3. Gateway cập nhật quan hệ quản lý khi các điều kiện bàn giao đã hoàn tất.
4. Use Case kết thúc thành công.

## Exception Flows

### 4b - Thông tin hoặc quan hệ quản lý không hợp lệ

1. Gateway từ chối đăng ký hoặc cập nhật và thông báo dữ liệu cần điều chỉnh.
2. Use Case kết thúc không thành công.

### 2f - Người dùng không có quyền quản lý Agent

1. Gateway từ chối truy cập.
2. Use Case kết thúc không thành công.

### 2g - Không thể kích hoạt Agent

1. Gateway thông báo thành phần còn thiếu: đường dẫn ngân sách, hạn mức hoặc API key.
2. Agent giữ nguyên trạng thái chưa kích hoạt.

## Business Rules

- `BR-UC04-01`: Mỗi Agent thuộc một tổ chức và có một lập trình viên quản lý chính tại một thời điểm.
- `BR-UC04-02`: Đội ngũ của Agent được xác định từ lập trình viên quản lý hiện hành.
- `BR-UC04-03`: Phân bổ hạn mức cho Agent thuộc UC03; quản lý API key thuộc UC05.
- `BR-UC04-04`: Agent không tự quản lý thông tin, hạn mức hoặc API key.
- `BR-UC04-05`: Chỉ Agent hoạt động, có đường dẫn ngân sách hợp lệ và API key hoạt động mới được thực hiện UC01.
- `BR-UC04-06`: Bàn giao Agent phải xử lý hạn mức và các yêu cầu đang xử lý trước khi thay đổi người quản lý.

## Non-Functional Requirements

- `NFR02`: Bảo mật và phân quyền.
- `NFR03`: Khả năng truy vết.

## Traceability

- Functional Requirements: `FR02, FR05, FR09, FR10`.
- Diagram: `diagram.puml`, `diagram.png`.
