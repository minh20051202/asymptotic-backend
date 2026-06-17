# Danh sách 20 diagram ưu tiên cho đồ án OOAD

Tài liệu này liệt kê các biểu đồ nên vẽ cho đồ án **AI Agent Financial Gateway** theo hướng OOAD. Danh sách được ưu tiên theo giá trị đối với báo cáo: làm rõ yêu cầu, chứng minh phân tích hướng đối tượng, hỗ trợ thiết kế và phục vụ kiểm thử/đối chiếu.

## Nhóm 1 -- Phân tích yêu cầu

### 1. Use Case Diagram tổng quát

- **Chương đề xuất:** Chương 2 -- Phân tích trường hợp sử dụng.
- **Mục đích:** Xác định phạm vi hệ thống và các actor tương tác với Gateway.
- **Actor chính:** Organization Admin, AI Agent, System Admin, AI Provider, Payment Provider.
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
- **Actor chính:** Organization Admin, Payment Provider, System Admin.
- **Nội dung chính:** Nạp tiền, xem số dư, phân bổ ngân sách, thiết lập quota/budget/policy, xem giao dịch.
- **Ưu tiên:** Cao.

### 4. Activity Diagram UC01 -- Agent gọi AI qua Gateway

- **Chương đề xuất:** Chương 2 -- Mô hình hóa hành vi hệ thống.
- **Mục đích:** Mô tả luồng nghiệp vụ trung tâm của MVP.
- **Nội dung chính:** Nhận request, xác thực API key, kiểm tra agent, kiểm tra policy, kiểm tra ngân sách, route tới AI Provider, nhận phản hồi, ghi trace/cost, trả response.
- **Ưu tiên:** Rất cao.

### 5. Activity Diagram đăng ký AI Agent và cấp API key

- **Chương đề xuất:** Chương 2.
- **Mục đích:** Làm rõ hệ thống không tạo agent, chỉ đăng ký agent bên ngoài vào Gateway.
- **Nội dung chính:** Organization Admin tạo agent record, hệ thống sinh API key, gắn agent với tổ chức, trả khóa truy cập.
- **Ưu tiên:** Cao.

### 6. Activity Diagram nạp tiền và phân bổ ngân sách

- **Chương đề xuất:** Chương 2.
- **Mục đích:** Mô tả nghiệp vụ tài chính trước khi agent phát sinh request.
- **Nội dung chính:** Organization Admin nạp tiền, Gateway ghi nhận ví/ngân sách, admin phân bổ budget cho agent hoặc nhóm agent.
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
- **Nội dung chính:** Organization Admin -> AdminBoundary -> AgentRegistrationControl -> Organization/AIAgent/ApiKey.
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

### 18. Sequence Diagram đăng ký AI Agent và cấp API key

- **Chương đề xuất:** Chương 4.
- **Mục đích:** Mô tả tương tác giữa admin, service và repository khi đăng ký agent.
- **Nội dung chính:** Organization Admin -> Admin API -> AgentService -> ApiKeyService -> Repository -> trả API key.
- **Ưu tiên:** Cao.

### 19. Sequence Diagram nạp tiền và phân bổ ngân sách

- **Chương đề xuất:** Chương 4.
- **Mục đích:** Mô tả tương tác nghiệp vụ tài chính trước khi agent sử dụng dịch vụ.
- **Nội dung chính:** Organization Admin -> WalletService -> Payment/Topup flow -> Ledger -> BudgetService.
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
- Nếu báo cáo cần gắn với backend triển khai, có thể thêm ERD, Component Diagram và Deployment Diagram ở Chương 4 hoặc Chương 5, nhưng không nên thay thế các biểu đồ OOAD cốt lõi.
