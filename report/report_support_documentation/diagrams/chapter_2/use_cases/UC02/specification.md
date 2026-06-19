# UC02 - Nạp tiền vào ví tổ chức

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC02` |
| Tên | Nạp tiền vào ví tổ chức |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Quản trị viên tổ chức muốn nạp tiền vào ví tổ chức để tạo ngân sách sử dụng dịch vụ AI. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Quản trị viên tổ chức |
| Supporting Actor | Nhà cung cấp dịch vụ thanh toán |
| Priority | Must Have |
| Trigger | Quản trị viên tổ chức yêu cầu nạp tiền. |
| Pre-Conditions | Quản trị viên tổ chức đã đăng nhập; tổ chức và ví tổ chức đang hoạt động. |
| Success Post-Conditions | Ví tổ chức được cộng số tiền hợp lệ; giao dịch nạp tiền được ghi nhận. |
| Failure Post-Conditions | Số dư ví không thay đổi nếu thanh toán không hợp lệ; giao dịch lỗi được ghi nhận để đối soát. |

## Basic Flow

1. Quản trị viên tổ chức nhập số tiền cần nạp và chọn phương thức thanh toán.
2. Gateway kiểm tra thông tin nạp tiền.
3. Gateway gửi yêu cầu thanh toán tới nhà cung cấp dịch vụ thanh toán.
4. Quản trị viên hoàn tất thanh toán tại nhà cung cấp.
5. Nhà cung cấp thanh toán gửi kết quả thanh toán về Gateway.
6. Gateway xác thực kết quả và kiểm tra giao dịch chưa được ghi nhận trước đó.
7. Gateway cập nhật số dư ví tổ chức và ghi nhận giao dịch nạp tiền.
8. Gateway thông báo kết quả nạp tiền cho quản trị viên.

## Alternative Flows

### 5a - Thanh toán đang chờ xác nhận

1. Gateway thông báo giao dịch đang chờ xử lý.
2. Khi nhận được kết quả hợp lệ, Use Case tiếp tục tại bước 6.

## Exception Flows

### 2b - Số tiền nạp không hợp lệ

1. Gateway từ chối yêu cầu và thông báo nguyên nhân.
2. Use Case kết thúc không thành công.

### 6b - Kết quả thanh toán không hợp lệ

1. Gateway từ chối cập nhật ví và ghi nhận giao dịch cần kiểm tra.
2. Use Case kết thúc không thành công.

### 7b - Cập nhật ví thất bại

1. Gateway không thay đổi số dư ví và ghi nhận giao dịch cần xử lý.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC02-01`: Mỗi mã giao dịch thanh toán chỉ được ghi nhận một lần.
- `BR-UC02-02`: Chỉ kết quả thanh toán đã được nhà cung cấp thanh toán xác thực mới được cập nhật vào ví.
- `BR-UC02-03`: Mọi biến động ví phải có bút toán tương ứng.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR02`: Bảo mật khóa truy cập.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.

## Traceability

- Functional Requirements: `FR01, FR09`.
- Diagram: `diagram.puml`, `diagram.png`.
