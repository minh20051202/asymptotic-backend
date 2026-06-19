# Danh sách 20 diagram ưu tiên cho đồ án OOAD

Tài liệu này liệt kê các biểu đồ nên vẽ cho đồ án **AI Agent Financial Gateway** theo hướng OOAD. Danh sách được ưu tiên theo giá trị đối với báo cáo: làm rõ yêu cầu, chứng minh phân tích hướng đối tượng, hỗ trợ thiết kế và phục vụ kiểm thử/đối chiếu.

Phân loại notation:

- Analysis Class và Design Class là cách dùng UML Class Diagram.
- Robustness/BCE là kỹ thuật phương pháp OOAD.
- ERD và artifact kết quả kiểm thử không thuộc UML.

## Nhóm 1 -- Phân tích yêu cầu

### 1. Use Case Diagram tổng quát

- **Chương đề xuất:** Chương 2 -- Phân tích trường hợp sử dụng.
- **Mục đích:** Xác định phạm vi hệ thống và các actor tương tác với Gateway.
- **Actor chính:** Organization Admin, Developer, AI Agent, System Admin, AI Provider, Payment Provider.
- **Nội dung chính:** Đăng ký tổ chức, đăng ký AI Agent, cấp API key, nạp tiền, phân bổ ngân sách, gọi dịch vụ AI, cấu hình provider/model/pricing, xem usage/cost/trace.
- **Ưu tiên:** Rất cao.

### 2. Use Case Diagram nhóm Agent/Gateway

- **Chương đề xuất:** Chương 2.
- **Mục đích:** Làm rõ luồng AI Agent sử dụng Gateway thay vì gọi trực tiếp AI Provider.
- **Actor chính:** AI Agent, AI Provider.
- **Nội dung chính:** Xác thực API key, kiểm tra agent, kiểm tra policy/budget, route request, ghi nhận usage/cost/trace.
- **Ưu tiên:** Cao.

### 3. Use Case Diagram nhóm Finance/Admin

- **Chương đề xuất:** Chương 2.
- **Mục đích:** Làm rõ khía cạnh financial gateway.
- **Actor chính:** Organization Admin, Developer, Payment Provider, System Admin.
- **Nội dung chính:** Nạp tiền, xem số dư, phân bổ hạn mức theo chuỗi tổ chức--đội ngũ--Developer--Agent, thiết lập policy và xem giao dịch.
- **Ưu tiên:** Cao.

### 4. Activity Diagram UC01 -- Agent gọi AI qua Gateway

- **Chương đề xuất:** Chương 2 -- Mô hình hóa hành vi hệ thống.
- **Mục đích:** Mô tả luồng nghiệp vụ trung tâm của MVP.
- **Nội dung chính:** Nhận request, xác thực API key, kiểm tra agent, kiểm tra policy, kiểm tra ngân sách, route tới AI Provider, nhận phản hồi, ghi trace/cost, trả response.
- **Ưu tiên:** Rất cao.

### 5. Activity Diagram đăng ký AI Agent

- **Chương đề xuất:** Chương 2.
- **Mục đích:** Làm rõ hệ thống không tạo Agent mà cho phép Developer đăng ký Agent bên ngoài vào Gateway.
- **Nội dung chính:** Developer gửi thông tin Agent, hệ thống kiểm tra quan hệ thành viên, tạo bản ghi Agent chưa kích hoạt và gắn Developer quản lý. Cấp hạn mức và API key là các use case riêng.
- **Ưu tiên:** Cao.

### 6. Activity Diagram nạp tiền và phân bổ ngân sách

- **Chương đề xuất:** Chương 2.
- **Mục đích:** Mô tả nghiệp vụ tài chính trước khi agent phát sinh request.
- **Nội dung chính:** Organization Admin nạp tiền, hệ thống ghi nhận ví tổ chức, sau đó phân bổ hạn mức tổ chức--đội ngũ--Developer; Developer phân bổ hạn mức cho Agent.
- **Ưu tiên:** Cao.

## Nhóm 2 -- Phân tích hướng đối tượng

### 7. Analysis Class Diagram tổng quát

- **Chương đề xuất:** Chương 3 -- Phân tích hướng đối tượng.
- **Mục đích:** Trình bày mô hình lớp phân tích theo Boundary-Control-Entity.
- **Nhóm lớp chính:** Boundary, Control, Entity.
- **Ưu tiên:** Rất cao.

### 8. Boundary Class Diagram

- **Chương đề xuất:** Chương 3.
- **Mục đích:** Xác định các lớp giao tiếp giữa actor/hệ thống ngoài với hệ thống.
- **Lớp gợi ý:** AgentRequestBoundary, AdminBoundary, ProviderBoundary, PaymentBoundary.
- **Ưu tiên:** Cao.

### 9. Control Class Diagram

- **Chương đề xuất:** Chương 3.
- **Mục đích:** Xác định các lớp điều phối use case.
- **Lớp gợi ý:** AgentRegistrationControl, ApiKeyControl, BudgetControl, PolicyControl, GatewayRoutingControl, TraceControl.
- **Ưu tiên:** Cao.

### 10. Entity Class Diagram

- **Chương đề xuất:** Chương 3.
- **Mục đích:** Xác định các thực thể nghiệp vụ cốt lõi.
- **Lớp gợi ý:** Organization, AIAgent, ApiKey, Wallet, Budget, Policy, Provider, AIModel, RequestTrace, FinancialTransaction.
- **Ưu tiên:** Cao.

### 11. Robustness Diagram UC01 -- Agent gọi AI qua Gateway

- **Chương đề xuất:** Chương 3.
- **Mục đích:** Cầu nối giữa use case và class analysis.
- **Nội dung chính:** AI Agent -> AgentRequestBoundary -> GatewayRequestControl/BudgetControl/ProviderRoutingControl -> Entity.
- **Ưu tiên:** Trung bình cao.

### 12. Robustness Diagram đăng ký AI Agent

- **Chương đề xuất:** Chương 3.
- **Mục đích:** Làm rõ luồng đăng ký agent bên ngoài, tránh hiểu nhầm hệ thống tạo/huấn luyện agent.
- **Nội dung chính:** Developer -> OrganizationAdminBoundary -> AgentAccessControl -> Organization/Team/DeveloperProfile/Agent.
- **Ưu tiên:** Trung bình cao.

## Nhóm 3 -- Thiết kế hướng đối tượng

### 13. Package Diagram

- **Chương đề xuất:** Chương 4 -- Thiết kế hướng đối tượng.
- **Mục đích:** Trình bày cách chia module/gói trong thiết kế.
- **Gói gợi ý:** identity, organization, agent, apikey, ledger, policy, provider, gateway, trace, admin.
- **Ưu tiên:** Rất cao.

### 14. Design Class Diagram tổng quát

- **Chương đề xuất:** Chương 4.
- **Mục đích:** Chuyển mô hình phân tích sang mô hình thiết kế gần với hệ thống triển khai.
- **Nội dung chính:** Service, Repository, Entity, Adapter và quan hệ phụ thuộc chính.
- **Ưu tiên:** Rất cao.

### 15. Design Class Diagram cho Gateway Request Flow

- **Chương đề xuất:** Chương 4.
- **Mục đích:** Làm rõ các lớp tham gia xử lý request agent -> Gateway -> AI Provider.
- **Lớp gợi ý:** GatewayHandler, ApiKeyService, AgentService, PolicyService, BudgetService, ProviderRouter, ProviderAdapter, TraceService.
- **Ưu tiên:** Cao.

### 16. Design Class Diagram cho Finance/Ledger

- **Chương đề xuất:** Chương 4.
- **Mục đích:** Làm rõ thiết kế tài chính của Gateway.
- **Lớp gợi ý:** WalletService, BudgetService, BudgetReservationService, TransactionService, LedgerEntry, WalletRepository, TransactionRepository.
- **Ưu tiên:** Cao.

### 17. Sequence Diagram UC01 -- Agent gọi AI qua Gateway

- **Chương đề xuất:** Chương 4 -- Thiết kế tương tác.
- **Mục đích:** Biểu diễn tương tác runtime quan trọng nhất của hệ thống.
- **Nội dung chính:** AI Agent -> Gateway -> ApiKeyService -> PolicyService -> BudgetService -> ProviderAdapter -> AI Provider -> Trace/Ledger.
- **Ưu tiên:** Rất cao.

### 18. Sequence Diagram đăng ký AI Agent

- **Chương đề xuất:** Chương 4.
- **Mục đích:** Mô tả tương tác giữa Developer, các module nội bộ và persistence khi đăng ký Agent.
- **Nội dung chính:** Developer -> Admin API -> AgentService -> OrganizationStructureService -> Persistence dùng chung. Việc cấp API key thuộc sequence/use case UC05 riêng.
- **Ưu tiên:** Cao.

### 19. Sequence Diagram nạp tiền và phân bổ ngân sách

- **Chương đề xuất:** Chương 4.
- **Mục đích:** Mô tả tương tác nghiệp vụ tài chính trước khi agent sử dụng dịch vụ.
- **Nội dung chính:** Organization Admin -> WalletService -> Payment/Topup flow -> Ledger -> BudgetAllocationService; sau đó phân bổ theo ba quan hệ liền kề đến Agent.
- **Ưu tiên:** Cao.

### 20. State Machine Diagram cho AI Request và Financial Transaction

- **Chương đề xuất:** Chương 4 -- Thiết kế trạng thái.
- **Mục đích:** Mô tả vòng đời trạng thái của request và giao dịch tài chính.
- **AI Request states:** Received, Authenticated, PolicyChecked, BudgetChecked, Routed, Completed, Failed, Rejected.
- **Financial Transaction states:** Pending, Reserved, Settled, Released, Failed, Reversed.
- **Ưu tiên:** Rất cao.

## Ghi chú triển khai

- Nếu cần giảm số lượng biểu đồ, ưu tiên giữ các mục: 1, 4, 7, 10, 13, 14, 17, 20.
- Nếu muốn thể hiện OOAD rõ hơn, ưu tiên bổ sung Robustness Diagram cho các use case chính.
- Boundary, Control và Entity views hiện là các góc nhìn của cùng Analysis Class Model, không phải ba loại UML Diagram độc lập.
- Nếu cần mở rộng Chương 3, ưu tiên thêm focused Analysis Class Diagram theo nhóm use case thay vì tạo một sơ đồ lớn hơn.

## Biểu đồ bổ sung ngoài danh sách 20

### Logical ERD

- **Chương:** Chương 4.
- **Mục đích:** Mô tả entity dữ liệu, identifier, cardinality, optionality và associative entity.
- **Điều kiện:** Chỉ thêm khi mô hình dữ liệu logic đã chốt.
- **Phân loại:** Không thuộc UML.

### Component Diagram as-designed

- **Chương:** Chương 4.
- **Mục đích:** Mô tả component, interface và external system ở mức thiết kế.
- **Điều kiện:** Không biến package/module nội bộ thành microservice.
- **Phân loại:** UML.

### Deployment Diagram as-designed

- **Chương:** Chương 4.
- **Mục đích:** Mô tả artifact, node và communication path dự kiến.
- **Điều kiện:** Kiến trúc triển khai đã được chốt.
- **Phân loại:** UML.

### Component/Deployment Diagram as-built

- **Chương:** Chương 5.
- **Mục đích:** Chứng minh cấu trúc và môi trường thực tế đã triển khai.
- **Điều kiện:** Có artifact, node và bằng chứng triển khai.
- **Phân loại:** UML, tùy chọn.

### Minh chứng kiểm thử

- **Chương:** Chương 5.
- **Nội dung:** Ma trận Requirement--Use Case--Module--Test, bảng test case, kết quả và biểu đồ hiệu năng.
- **Điều kiện:** Có phương pháp đo và dữ liệu thực tế.
- **Phân loại:** Không thuộc UML.
