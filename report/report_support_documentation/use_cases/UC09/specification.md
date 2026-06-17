# UC09 - Quản lý model, provider và chính sách giá

| Thuộc tính | Nội dung |
|---|---|
| ID | `UC09` |
| Tên | Quản lý model, provider và chính sách giá |
| Trạng thái | Thiết kế mục tiêu |
| Mô tả | Quản trị viên hệ thống muốn quản lý provider, model và chính sách giá để Gateway định tuyến, ước lượng chi phí và tính phí request AI. |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Quản trị viên hệ thống |
| Supporting Actor | Không có |
| Priority | Must Have |
| Trigger | Quản trị viên hệ thống yêu cầu cập nhật provider, model hoặc chính sách giá. |
| Pre-Conditions | Quản trị viên hệ thống đã đăng nhập. |
| Success Post-Conditions | Cấu hình provider, model và chính sách giá được cập nhật; Gateway có thể sử dụng cấu hình mới cho các yêu cầu tiếp theo. |
| Failure Post-Conditions | Cấu hình hiện tại không thay đổi nếu dữ liệu mới không hợp lệ hoặc gây xung đột. |

## Basic Flow

1. Quản trị viên hệ thống mở chức năng quản lý model, provider và chính sách giá.
2. Gateway hiển thị trạng thái provider, model, provider credential và chính sách giá hiện tại.
3. Quản trị viên chọn thao tác cập nhật thông tin model hoặc provider.
4. Quản trị viên nhập thông tin cần thay đổi.
5. Gateway kiểm tra cấu hình mới theo quy tắc hợp lệ.
6. Gateway lưu cấu hình model hoặc provider.
7. Gateway thông báo cấu hình đã sẵn sàng cho các yêu cầu tiếp theo.

## Alternative Flows

### 2a - Xem danh mục model/provider

1. Nếu chỉ cần tra cứu, quản trị viên hệ thống xem danh mục model, provider, trạng thái hỗ trợ và biểu giá hiện tại.
2. Use Case kết thúc thành công.

### 3a - Thêm model/provider

1. Nếu cần bổ sung cấu hình mới, quản trị viên hệ thống chọn thao tác thêm model hoặc provider.
2. Quản trị viên nhập thông tin cấu hình, trạng thái hỗ trợ và biểu giá ban đầu nếu có.
3. Gateway kiểm tra dữ liệu và lưu cấu hình mới.
4. Use Case kết thúc thành công.

### 3b - Vô hiệu hóa model

1. Nếu cần ngừng sử dụng một model, quản trị viên hệ thống chọn model từ danh mục model/provider.
2. Quản trị viên hệ thống chọn thao tác vô hiệu hóa model.
3. Gateway ngừng chọn model này cho yêu cầu mới.
4. Use Case kết thúc thành công.

### 3c - Cập nhật biểu giá

1. Nếu cần thay đổi biểu giá, quản trị viên hệ thống chọn model hoặc provider từ danh mục.
2. Quản trị viên hệ thống nhập giá mới và thời điểm hiệu lực.
3. Gateway lưu biểu giá mới, không làm thay đổi các giao dịch đã quyết toán.
4. Use Case kết thúc thành công.

### 3d - Thiết lập chính sách fallback

1. Nếu cần phương án dự phòng, quản trị viên hệ thống chọn model hoặc provider gốc từ danh mục.
2. Quản trị viên hệ thống cấu hình model hoặc provider thay thế.
3. Gateway lưu chính sách fallback để UC01 sử dụng khi model ban đầu không khả dụng hoặc không phù hợp với ngân sách.
4. Use Case kết thúc thành công.

### 3e - Model đang được sử dụng trong request đang xử lý

1. Nếu model đang được request sử dụng, Gateway áp dụng thay đổi cho yêu cầu mới và giữ nguyên cấu hình cũ cho request đang xử lý.
2. Use Case kết thúc thành công có điều kiện.

## Exception Flows

### 5b - Cấu hình không hợp lệ

1. Gateway từ chối thay đổi và giữ nguyên cấu hình hiện tại.
2. Use Case kết thúc không thành công.

## Business Rules

- `BR-UC09-01`: Mọi yêu cầu AI phải sử dụng provider, model và biểu giá đang có hiệu lực.
- `BR-UC09-02`: Thay đổi biểu giá không được làm sai lệch giao dịch đã quyết toán.
- `BR-UC09-03`: Model bị vô hiệu hóa không được chọn cho yêu cầu mới.
- `BR-UC09-04`: Provider credential nội bộ chỉ được Gateway sử dụng khi route request tới AI Provider.
- `BR-UC09-05`: Tổ chức, lập trình viên và Agent không được truy cập provider credential nội bộ.
- `BR-UC09-06`: Provider credential nội bộ phải được lưu bảo mật và không được xuất hiện trong log, báo cáo hoặc phản hồi lỗi.

## Non-Functional Requirements

- `NFR02`: Bảo mật khóa truy cập.
- `NFR03`: Khả năng truy vết.
- `NFR04`: Khả năng phục hồi.

## Traceability

- Functional Requirements: `FR04, FR05, FR07`.
- Diagram: `diagram.drawio`, `diagram.png`.
