# UC07 - Đăng ký tổ chức

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC07` |
| Tên | Đăng ký tổ chức |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Người dùng muốn đăng ký tổ chức mới để bắt đầu sử dụng hệ thống với vai trò quản trị viên tổ chức. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Người dùng |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Người dùng gửi yêu cầu tạo tổ chức. |
| Pre-Conditions | Người dùng đã có danh tính hợp lệ; thông tin tổ chức chưa tồn tại trong hệ thống. |
| Success Post-Conditions | Tổ chức mới được khởi tạo; người dùng hiện tại trở thành quản trị viên tổ chức; ví tổ chức được tạo. |
| Failure Post-Conditions | Không tạo dữ liệu tổ chức không hoàn chỉnh; lỗi được thông báo cho người đăng ký. |

## Basic Flow

1. Người dùng nhập thông tin tổ chức cần đăng ký.
2. Gateway kiểm tra tính hợp lệ và tính duy nhất của thông tin tổ chức.
3. Gateway khởi tạo tổ chức mới.
4. Gateway gán người dùng hiện tại làm quản trị viên đầu tiên của tổ chức.
5. Gateway tạo ví tổ chức với số dư ban đầu bằng không.
6. Gateway thông báo đăng ký thành công và cho phép người dùng bắt đầu phiên làm việc.

## Alternative Flows

Không có.

## Exception Flows

### 2b - Thông tin đăng ký không hợp lệ hoặc bị trùng

1. Gateway từ chối yêu cầu và thông báo trường dữ liệu cần điều chỉnh.
2. Use Case kết thúc không thành công.

### 3b - Khởi tạo tổ chức thất bại

1. Gateway hủy toàn bộ dữ liệu đã tạo trong yêu cầu đăng ký.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC07-01`: Người dùng tạo tổ chức thành công trở thành quản trị viên đầu tiên của tổ chức đó.
- `BR-UC07-02`: Tổ chức mới phải có ví tổ chức trước khi thực hiện UC02 hoặc UC03.
- `BR-UC07-03`: Dữ liệu đăng ký phải được tạo nhất quán.

## Non-Functional Requirements

- `NFR01`: Tính nhất quán tài chính.
- `NFR02`: Bảo mật khóa truy cập.

## Traceability

- Functional Requirements: `FR01`.
- Diagram: `diagram.drawio`, `diagram.png`.
