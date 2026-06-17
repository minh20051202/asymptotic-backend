# UC03 - Quản lý ngân sách đội ngũ

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC03` |
| Tên | Quản lý ngân sách đội ngũ |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Quản trị viên tổ chức muốn quản lý ngân sách đội ngũ để phân bổ hoặc thu hồi ngân sách sử dụng dịch vụ AI. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Quản trị viên tổ chức |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Quản trị viên tổ chức mở chức năng quản lý ngân sách đội ngũ. |
| Pre-Conditions | Tổ chức đang hoạt động; lập trình viên mục tiêu thuộc tổ chức. |
| Success Post-Conditions | Hạn mức của lập trình viên và ví tổ chức được cập nhật nhất quán; lịch sử điều chuyển được ghi nhận. |
| Failure Post-Conditions | Không thay đổi số dư hoặc hạn mức khi yêu cầu không hợp lệ; lỗi được ghi nhận. |

## Basic Flow

1. Quản trị viên tổ chức mở chức năng quản lý ngân sách đội ngũ.
2. Gateway hiển thị ngân sách khả dụng của tổ chức, danh sách lập trình viên, hạn mức hiện tại và lịch sử điều chuyển liên quan.
3. Quản trị viên chọn thao tác phân bổ ngân sách cho một lập trình viên.
4. Quản trị viên nhập số tiền cần phân bổ.
5. Gateway kiểm tra số tiền phân bổ không vượt ngân sách khả dụng của tổ chức.
6. Gateway điều chuyển ngân sách từ ví tổ chức sang hạn mức của lập trình viên.
7. Gateway ghi nhận lịch sử điều chuyển và thông báo kết quả cho quản trị viên.

## Alternative Flows

### 2a - Xem ngân sách tổ chức hoặc lập trình viên

1. Nếu chỉ cần tra cứu, quản trị viên xem ngân sách khả dụng của tổ chức hoặc hạn mức của từng lập trình viên.
2. Use Case kết thúc thành công.

### 2b - Xem lịch sử điều chuyển ngân sách

1. Nếu cần kiểm tra biến động ngân sách, quản trị viên chọn lịch sử điều chuyển.
2. Gateway hiển thị các lần phân bổ, thu hồi và trạng thái xử lý liên quan.
3. Use Case kết thúc thành công.

### 3a - Thu hồi ngân sách

1. Nếu cần thu hồi ngân sách, quản trị viên chọn thao tác thu hồi ngân sách của một lập trình viên.
2. Gateway hiển thị hạn mức hiện tại và phần ngân sách có thể thu hồi của lập trình viên.
3. Quản trị viên nhập số tiền cần thu hồi.
4. Gateway kiểm tra phần ngân sách có thể thu hồi.
5. Gateway điều chuyển ngân sách khả dụng về ví tổ chức và ghi nhận lịch sử điều chuyển.
6. Gateway thông báo kết quả cho quản trị viên.
7. Use Case kết thúc thành công.

### 5a - Ngân sách tổ chức không đủ

1. Gateway thông báo ngân sách khả dụng của tổ chức không đủ cho số tiền phân bổ.
2. Quản trị viên điều chỉnh số tiền phân bổ.
3. Use Case tiếp tục tại bước 5.

## Exception Flows

### 3b - Lập trình viên không thuộc tổ chức

1. Gateway từ chối thao tác và giữ nguyên trạng thái ngân sách.
2. Use Case kết thúc không thành công.

### 3c - Số tiền thu hồi vượt phần khả dụng

1. Gateway từ chối phần thu hồi vượt quá ngân sách chưa cam kết.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC03-01`: Không được phân bổ vượt ngân sách khả dụng của tổ chức.
- `BR-UC03-02`: Không được thu hồi phần ngân sách đang tạm giữ hoặc đã phân bổ cho tác nhân đang hoạt động.
- `BR-UC03-03`: Mọi điều chuyển nội bộ phải có lịch sử kiểm toán.
- `BR-UC03-04`: Ngân sách của lập trình viên là nguồn cấp trên cho hạn mức của các tác nhân AI thuộc lập trình viên đó.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.

## Traceability

- Functional Requirements: `FR01, FR05, FR09`.
- Diagram: `diagram.drawio`, `diagram.png`.
