# UC08 - Quản lý đội ngũ và thành viên tổ chức

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC08` |
| Tên | Quản lý đội ngũ và thành viên tổ chức |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Quản trị viên tổ chức quản lý đội ngũ, lập trình viên, quan hệ thành viên và việc bàn giao Agent khi cơ cấu tổ chức thay đổi. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Quản trị viên tổ chức |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Quản trị viên tổ chức mở chức năng quản lý đội ngũ và thành viên. |
| Pre-Conditions | Quản trị viên đã đăng nhập; tổ chức đang hoạt động. |
| Success Post-Conditions | Đội ngũ, thành viên và quan hệ quản lý được cập nhật nhất quán; đường dẫn ngân sách của Agent vẫn hợp lệ. |
| Failure Post-Conditions | Không thay đổi cơ cấu nếu còn hạn mức hoặc giao dịch chưa được xử lý an toàn. |

## Basic Flow

1. Quản trị viên mở chức năng quản lý đội ngũ và thành viên.
2. Gateway hiển thị danh sách đội ngũ, lập trình viên, trạng thái và quan hệ hiện hành.
3. Quản trị viên chọn tạo đội ngũ mới.
4. Quản trị viên nhập tên, mô tả và thông tin quản lý đội ngũ.
5. Gateway kiểm tra dữ liệu và tính duy nhất trong tổ chức.
6. Gateway tạo đội ngũ ở trạng thái hoạt động.
7. Gateway ghi lịch sử thay đổi và thông báo kết quả.

## Alternative Flows

### 2a - Thêm lập trình viên vào tổ chức

1. Quản trị viên nhập hoặc chọn người dùng.
2. Gateway tạo vai trò lập trình viên và gán vào một đội ngũ.
3. Use Case kết thúc thành công.

### 2b - Cập nhật đội ngũ hoặc thành viên

1. Quản trị viên chọn đối tượng và nhập thông tin cần thay đổi.
2. Gateway kiểm tra quyền và dữ liệu.
3. Gateway cập nhật thông tin.
4. Use Case kết thúc thành công.

### 2c - Chuyển lập trình viên sang đội ngũ khác

1. Quản trị viên chọn lập trình viên và đội ngũ đích.
2. Gateway kiểm tra hạn mức và các Agent do lập trình viên quản lý.
3. Quản trị viên xử lý hoặc tái phân bổ hạn mức theo UC03.
4. Gateway cập nhật quan hệ đội ngũ sau khi đường dẫn ngân sách hợp lệ.
5. Use Case kết thúc thành công.

### 2d - Bàn giao Agent

1. Quản trị viên chọn Agent và lập trình viên quản lý mới.
2. Gateway xử lý hạn mức theo UC03 và quan hệ quản lý theo UC04.
3. Gateway cập nhật đường dẫn chịu chi phí.
4. Use Case kết thúc thành công.

### 2e - Vô hiệu hóa lập trình viên

1. Quản trị viên yêu cầu vô hiệu hóa lập trình viên.
2. Gateway chặn thao tác mới của lập trình viên và các Agent chưa được bàn giao.
3. Quản trị viên bàn giao Agent, thu hồi API key khi cần và xử lý hạn mức còn lại.
4. Gateway vô hiệu hóa lập trình viên.
5. Use Case kết thúc thành công.

### 2f - Vô hiệu hóa đội ngũ

1. Quản trị viên yêu cầu vô hiệu hóa đội ngũ.
2. Gateway kiểm tra đội ngũ không còn lập trình viên, Agent hoặc giao dịch chưa xử lý.
3. Gateway thu hồi phần hạn mức khả dụng theo UC03 và vô hiệu hóa đội ngũ.
4. Use Case kết thúc thành công.

## Exception Flows

### 5b - Dữ liệu bị trùng hoặc không hợp lệ

1. Gateway từ chối thay đổi và thông báo nguyên nhân.
2. Use Case kết thúc không thành công.

### 2g - Còn yêu cầu hoặc giao dịch chưa hoàn tất

1. Gateway chặn thao tác chuyển hoặc vô hiệu hóa.
2. Gateway yêu cầu hoàn tất, hủy hoặc đối soát giao dịch liên quan.
3. Use Case kết thúc không thành công.

### 2h - Đường dẫn ngân sách sau thay đổi không hợp lệ

1. Gateway từ chối thay đổi cơ cấu.
2. Dữ liệu đội ngũ, lập trình viên và Agent giữ nguyên.

## Business Rules

- `BR-UC08-01`: Mỗi lập trình viên thuộc tối đa một đội ngũ tại một thời điểm.
- `BR-UC08-02`: Mỗi Agent có một lập trình viên quản lý chính tại một thời điểm.
- `BR-UC08-03`: Chuyển lập trình viên hoặc Agent phải xử lý hạn mức và các yêu cầu đang xử lý trước khi cập nhật quan hệ.
- `BR-UC08-04`: Vô hiệu hóa lập trình viên hoặc đội ngũ phải ngăn phát sinh yêu cầu AI mới.
- `BR-UC08-05`: Mọi thay đổi cơ cấu phải có lịch sử kiểm toán.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR02`: Bảo mật và phân quyền.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.

## Traceability

- Functional Requirements: `FR01, FR02, FR05, FR09, FR10`.
- Diagram: `diagram.puml`, `diagram.png`.
