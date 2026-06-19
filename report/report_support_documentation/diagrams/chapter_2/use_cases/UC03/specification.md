# UC03 - Quản lý và phân bổ ngân sách nội bộ

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC03` |
| Tên | Quản lý và phân bổ ngân sách nội bộ |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Quản trị viên tổ chức và lập trình viên phân bổ hoặc thu hồi hạn mức theo chuỗi tổ chức → đội ngũ → lập trình viên → AI Agent. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Quản trị viên tổ chức, Lập trình viên |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Quản trị viên tổ chức hoặc lập trình viên mở chức năng quản lý ngân sách trong phạm vi được phân quyền. |
| Pre-Conditions | Tổ chức và ví tổ chức đang hoạt động; đội ngũ, lập trình viên hoặc Agent mục tiêu thuộc đúng đường dẫn ngân sách. |
| Success Post-Conditions | Hạn mức nguồn và đích được cập nhật nhất quán; lịch sử phân bổ hoặc thu hồi được ghi nhận. |
| Failure Post-Conditions | Không thay đổi hạn mức khi yêu cầu không hợp lệ hoặc vượt quá phần khả dụng. |

## Basic Flow

1. Người dùng mở chức năng quản lý ngân sách.
2. Gateway xác định vai trò và phạm vi ngân sách người dùng được phép quản lý.
3. Gateway hiển thị hạn mức đã cấp, đã sử dụng và còn khả dụng tại các cấp liên quan.
4. Người dùng chọn đối tượng đích và thao tác phân bổ hạn mức.
5. Người dùng nhập số tiền cần phân bổ.
6. Gateway kiểm tra quyền, quan hệ nguồn–đích và hạn mức còn khả dụng.
7. Gateway cập nhật hạn mức của cấp đích.
8. Gateway ghi lịch sử điều chuyển và thông báo kết quả.

## Alternative Flows

### 3a - Xem ngân sách và hạn mức

1. Người dùng xem số liệu trong phạm vi được phân quyền.
2. Use Case kết thúc thành công.

### 4a - Quản trị viên cấp ngân sách cho đội ngũ

1. Quản trị viên tổ chức chọn đội ngũ đang hoạt động.
2. Gateway sử dụng ví/ngân sách khả dụng của tổ chức làm nguồn.
3. Use Case tiếp tục tại bước 5.

### 4b - Quản trị viên cấp hạn mức cho lập trình viên

1. Quản trị viên tổ chức chọn lập trình viên đang thuộc đội ngũ.
2. Gateway sử dụng hạn mức khả dụng của đội ngũ làm nguồn.
3. Use Case tiếp tục tại bước 5.

### 4c - Lập trình viên cấp hạn mức cho Agent

1. Lập trình viên chọn Agent do mình quản lý.
2. Gateway sử dụng hạn mức khả dụng của lập trình viên làm nguồn.
3. Use Case tiếp tục tại bước 5.

### 4d - Thu hồi hạn mức

1. Người dùng chọn đối tượng cần thu hồi hạn mức.
2. Gateway xác định phần hạn mức còn khả dụng có thể thu hồi.
3. Người dùng nhập số tiền cần thu hồi.
4. Gateway hoàn trả hạn mức về cấp nguồn trực tiếp và ghi lịch sử.
5. Use Case kết thúc thành công.

### 4e - Xem lịch sử điều chuyển

1. Người dùng chọn phạm vi và thời gian cần tra cứu.
2. Gateway hiển thị lịch sử phân bổ, thu hồi và người thực hiện.
3. Use Case kết thúc thành công.

## Exception Flows

### 2b - Người dùng không có quyền

1. Gateway từ chối truy cập và không thay đổi hạn mức.
2. Use Case kết thúc không thành công.

### 6b - Đường dẫn ngân sách không hợp lệ

1. Gateway phát hiện đội ngũ, lập trình viên hoặc Agent không thuộc quan hệ nguồn–đích hợp lệ.
2. Gateway từ chối thao tác.
3. Use Case kết thúc không thành công.

### 6c - Hạn mức nguồn không đủ

1. Gateway thông báo phần khả dụng không đủ cho số tiền yêu cầu.
2. Người dùng điều chỉnh số tiền hoặc hủy thao tác.

### 6d - Thu hồi vượt phần khả dụng

1. Gateway từ chối phần đã sử dụng hoặc đã cấp xuống cấp dưới.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC03-01`: Chuỗi phân bổ hợp lệ là tổ chức → đội ngũ → lập trình viên → AI Agent.
- `BR-UC03-02`: Ví tổ chức lưu tiền thật; đội ngũ, lập trình viên và Agent chỉ có hạn mức kiểm soát.
- `BR-UC03-03`: Tổng hạn mức cấp xuống không được vượt phần khả dụng của cấp nguồn.
- `BR-UC03-04`: Lập trình viên chỉ được cấp hoặc thu hồi hạn mức của Agent do mình quản lý.
- `BR-UC03-05`: Không được thu hồi phần đã sử dụng hoặc đã cấp xuống cấp dưới.
- `BR-UC03-06`: Mọi điều chuyển phải nguyên tử và có lịch sử kiểm toán.
- `BR-UC03-07`: Hạn mức khả dụng được xác định từ hạn mức đã cấp trừ phần đã sử dụng và đã phân bổ xuống cấp dưới.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR02`: Bảo mật và phân quyền.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.

## Traceability

- Functional Requirements: `FR01, FR05, FR09, FR10`.
- Diagram: `diagram.puml`, `diagram.png`.
