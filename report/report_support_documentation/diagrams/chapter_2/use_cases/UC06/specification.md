# UC06 - Theo dõi giao dịch, usage, cost và trace

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC06` |
| Tên | Theo dõi giao dịch, usage, cost và trace |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Người dùng có quyền muốn theo dõi giao dịch và mức độ sử dụng để kiểm tra chi phí, usage và trace trong phạm vi được phân quyền. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Quản trị viên tổ chức, Lập trình viên |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Người dùng mở chức năng báo cáo sử dụng. |
| Pre-Conditions | Người dùng đã đăng nhập và có quyền xem dữ liệu báo cáo. |
| Success Post-Conditions | Người dùng nhận được dữ liệu giao dịch và mức sử dụng phù hợp với phạm vi quyền truy cập. |
| Failure Post-Conditions | Dữ liệu nhạy cảm không bị lộ khi người dùng không có quyền hoặc yêu cầu lọc không hợp lệ. |

## Basic Flow

1. Người dùng mở chức năng theo dõi giao dịch và mức độ sử dụng.
2. Gateway xác định vai trò và phạm vi dữ liệu được phép xem.
3. Người dùng chọn tra cứu lịch sử giao dịch và nhập điều kiện lọc nếu cần.
4. Gateway truy xuất giao dịch, usage, cost và trace theo tổ chức, đội ngũ, lập trình viên và Agent phù hợp với phạm vi quyền.
5. Gateway hiển thị danh sách giao dịch và các chỉ số liên quan.

## Alternative Flows

### 2a - Xem tổng quan mức sử dụng

1. Nếu chỉ cần tổng quan, người dùng chọn xem mức sử dụng theo phạm vi được phân quyền.
2. Gateway tổng hợp chi phí, usage và trạng thái giao dịch.
3. Gateway hiển thị chỉ số tổng quan.
4. Use Case kết thúc thành công.

### 3a - Lọc dữ liệu theo thời gian

1. Nếu cần thu hẹp phạm vi tra cứu, người dùng chọn khoảng thời gian hoặc điều kiện lọc.
2. Use Case tiếp tục tại bước 4.

### 5a - Không có dữ liệu phù hợp

1. Gateway hiển thị trạng thái không có dữ liệu.
2. Use Case kết thúc thành công.

### 5b - Xem chi tiết giao dịch

1. Người dùng chọn một giao dịch trong danh sách.
2. Gateway hiển thị chi tiết vòng đời yêu cầu, mức sử dụng, chi phí và nhật ký thực thi liên quan.
3. Use Case kết thúc thành công.

### 5c - Xuất báo cáo

1. Người dùng yêu cầu xuất dữ liệu báo cáo.
2. Gateway tạo tệp báo cáo theo phạm vi được phép.
3. Use Case kết thúc thành công.

## Exception Flows

### 2b - Người dùng không có quyền

1. Gateway từ chối truy cập và không trả dữ liệu báo cáo.
2. Use Case kết thúc không thành công.

### 3b - Bộ lọc không hợp lệ

1. Gateway thông báo bộ lọc không hợp lệ và yêu cầu người dùng điều chỉnh điều kiện lọc.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC06-01`: Người dùng chỉ xem được dữ liệu trong phạm vi quyền của mình.
- `BR-UC06-02`: Mọi truy cập báo cáo nhạy cảm phải được ghi nhận.
- `BR-UC06-03`: Dữ liệu báo cáo phải khớp với giao dịch và bút toán đã ghi.
- `BR-UC06-04`: Quản trị viên tổ chức được tổng hợp theo tổ chức, đội ngũ, lập trình viên và Agent; lập trình viên chỉ xem dữ liệu của mình và các Agent do mình quản lý.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR02`: Bảo mật khóa truy cập.
- `NFR03`: Khả năng truy vết.

## Traceability

- Functional Requirements: `FR01, FR02, FR09`.
- Diagram: `diagram.puml`, `diagram.png`.
