# UC05 - Quản lý API key của Agent

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC05` |
| Tên | Quản lý API key của Agent |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Lập trình viên muốn quản lý API key của Agent để kiểm soát quyền gọi Gateway của AI Agent. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Lập trình viên |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Lập trình viên yêu cầu quản lý API key của một AI Agent. |
| Pre-Conditions | Lập trình viên đã đăng nhập; AI Agent thuộc quyền quản lý của lập trình viên. |
| Success Post-Conditions | API key của Agent được tạo hoặc thu hồi theo yêu cầu; trạng thái khóa được cập nhật trên Gateway. |
| Failure Post-Conditions | Không tạo hoặc thay đổi API key nếu yêu cầu không hợp lệ; sự kiện lỗi được ghi nhận. |

## Basic Flow

1. Lập trình viên chọn AI Agent cần quản lý API key.
2. Gateway hiển thị danh sách API key hiện có của Agent.
3. Lập trình viên yêu cầu tạo API key mới và nhập thông tin mô tả.
4. Gateway tạo API key mới cho Agent.
5. Gateway hiển thị giá trị API key một lần cho lập trình viên.
6. Lập trình viên xác nhận đã lưu API key.
7. Gateway kích hoạt API key và ghi nhận sự kiện bảo mật.

## Alternative Flows

### 2a - Xem danh sách API key

1. Nếu chỉ cần tra cứu, lập trình viên xem danh sách API key và trạng thái của từng khóa.
2. Use Case kết thúc thành công.

### 3a - Thu hồi API key

1. Nếu cần thu hồi quyền truy cập, lập trình viên chọn một API key đang hoạt động từ danh sách API key.
2. Gateway thu hồi API key và ngăn khóa này xác thực các yêu cầu AI mới.
3. Use Case kết thúc thành công.

## Exception Flows

### 1b - Agent không thuộc quyền quản lý

1. Gateway từ chối truy cập.
2. Use Case kết thúc không thành công.

### 4b - Không thể tạo API key

1. Gateway thông báo lỗi và không kích hoạt API key mới.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC05-01`: Giá trị API key thô chỉ được hiển thị một lần.
- `BR-UC05-02`: Gateway chỉ lưu giá trị bảo vệ của API key, không lưu khóa thô.
- `BR-UC05-03`: API key đã thu hồi không được dùng để thực hiện UC01.
- `BR-UC05-04`: API key của Agent là khóa do Asymptotic cấp, không phải provider API key của AI Provider.
- `BR-UC05-05`: Provider credential nội bộ của Asymptotic không được hiển thị cho tổ chức, lập trình viên hoặc Agent.

## Non-Functional Requirements

- `NFR02`: Bảo mật khóa truy cập.
- `NFR03`: Khả năng truy vết.

## Traceability

- Functional Requirements: `FR03`.
- Diagram: `diagram.drawio`, `diagram.png`.
