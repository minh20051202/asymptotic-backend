# Source of Truth -- AI Agent Financial Gateway

Tài liệu này là nguồn tham chiếu chính để thống nhất phạm vi, thuật ngữ, actor, use case, yêu cầu và định hướng mô hình hóa cho đồ án. Khi viết lại báo cáo, vẽ biểu đồ hoặc chỉnh tài liệu phân tích, ưu tiên bám theo tài liệu này để tránh lệch định vị đề tài.

## 1. Tên và định vị đề tài

### Tên đề tài

**THIẾT KẾ VÀ XÂY DỰNG CỔNG TÀI CHÍNH THỜI GIAN THỰC CHO TÁC NHÂN TRÍ TUỆ NHÂN TẠO**

### Tên hệ thống

**Asymptotic -- AI Agent Financial Gateway**

### Định vị hệ thống

Asymptotic là một gateway trung gian giữa **AI Agent bên ngoài** và **các nhà cung cấp dịch vụ AI trả phí**.

Hệ thống cho phép tổ chức:

- nạp tiền vào hệ thống;
- đăng ký AI Agent bên ngoài vào Gateway;
- cấp API key do Asymptotic quản lý cho từng Agent;
- thiết lập hạn mức, quota, budget và policy;
- cho phép Agent gọi dịch vụ AI thông qua Gateway;
- kiểm soát ngân sách và chính sách trước khi request được gửi tới AI Provider;
- ghi nhận usage, cost, trace và giao dịch để phục vụ kiểm tra, báo cáo và đối soát.

## 2. Bản chất hệ thống

### Hệ thống là gì?

Hệ thống là **cổng tài chính thời gian thực cho AI Agent**. Trọng tâm của hệ thống là kiểm soát tài chính, quyền truy cập và khả năng truy vết đối với các request do AI Agent tự động phát sinh.

### Hệ thống không phải là gì?

Hệ thống không phải:

- hệ thống tạo AI Agent;
- nền tảng huấn luyện AI Agent;
- môi trường runtime để chạy AI Agent;
- hệ thống điều phối workflow hoặc multi-agent;
- AI Financial Gateway tổng quát cho mọi nghiệp vụ tài chính;
- Payment Gateway production đầy đủ;
- hệ thống billing chỉ tổng hợp chi phí sau khi đã sử dụng dịch vụ.

### Điểm khác biệt của yếu tố "AI Agent"

Điểm khác biệt không nằm ở việc Asymptotic tự có AI Agent. Điểm khác biệt nằm ở **đối tượng được kiểm soát là AI Agent bên ngoài**: một thực thể phần mềm có khả năng tự động gọi API nhiều lần, thử lại, chạy song song và phát sinh chi phí khó dự đoán.

Vì vậy Gateway phải kiểm soát request theo:

- organization;
- agent_id;
- API key;
- policy;
- quota;
- budget;
- reservation hoặc kiểm tra hạn mức;
- idempotency;
- usage/cost/trace.

## 3. Luồng nghiệp vụ cốt lõi

Luồng cốt lõi của hệ thống:

1. Tổ chức đăng ký hoặc được tạo trong hệ thống.
2. Tổ chức nạp tiền vào ví/ngân sách.
3. Quản trị viên tổ chức tạo đội ngũ, gán lập trình viên và phân bổ ngân sách theo chuỗi tổ chức → đội ngũ → lập trình viên.
4. Tổ chức hoặc người được phân quyền đăng ký AI Agent bên ngoài vào Gateway và xác định lập trình viên quản lý.
5. Lập trình viên cấp hạn mức cho Agent trong phạm vi hạn mức được đội ngũ cấp.
6. Gateway cấp API key cho Agent.
7. AI Agent gửi request tới Gateway bằng API key do Asymptotic cấp.
8. Gateway xác thực API key và xác định tổ chức, đội ngũ, lập trình viên và Agent chịu chi phí.
9. Gateway kiểm tra policy, quota và ngân sách tại tất cả các cấp.
10. Nếu hợp lệ, Gateway gọi AI Provider bằng provider credential nội bộ của Asymptotic.
11. Gateway nhận kết quả, ghi nhận usage/cost/trace và cập nhật trạng thái tài chính.
12. Gateway trả kết quả hoặc lỗi rõ ràng cho Agent.

Nguyên tắc quan trọng:

- AI Agent không dùng trực tiếp provider API key.
- Provider credential nội bộ không được trả về cho tổ chức, developer hoặc Agent.
- Request chỉ nên được gửi tới AI Provider sau khi Gateway đã kiểm soát định danh, chính sách và ngân sách.

## 4. Phạm vi MVP

MVP bao gồm:

- Backend API.
- Quản lý tổ chức.
- Quản lý ví/ngân sách tổ chức.
- Nạp tiền hoặc ghi nhận top-up ở mức nguyên mẫu.
- Phân bổ ngân sách theo chuỗi tổ chức → đội ngũ → lập trình viên → AI Agent.
- Đăng ký và quản lý AI Agent bên ngoài.
- Cấp, xác thực, thu hồi API key của Agent.
- Quản lý provider, model, endpoint, biểu giá và provider credential nội bộ.
- Thiết lập policy, quota, budget và hạn mức.
- Gateway/proxy request tới AI Provider.
- Ước lượng hoặc xác định chi phí dự kiến trước khi gọi provider.
- Kiểm tra budget/quota/policy trước khi route request.
- Hỗ trợ idempotency để hạn chế tính phí trùng khi request được gửi lại.
- Ghi nhận usage, cost, trace, transaction hoặc ledger entry.
- Trả lỗi rõ ràng khi request bị từ chối do sai khóa, agent không hợp lệ, vượt quota, vượt budget hoặc vi phạm policy.

## 5. Ngoài phạm vi MVP

Các nội dung sau không thuộc MVP hoặc chỉ mô hình hóa ở mức thiết kế:

- Tạo, huấn luyện, triển khai hoặc điều phối AI Agent.
- Agent orchestration hoặc multi-agent workflow.
- Cung cấp provider API key trực tiếp cho tổ chức, developer hoặc Agent.
- Payment Gateway production đầy đủ.
- Invoice production đầy đủ.
- Settlement production đầy đủ với AI Provider.
- Dashboard quản trị hoàn chỉnh.
- Multi-model routing tối ưu chi phí theo thời gian thực ở mức production.
- Triển khai phân tán, high availability hoặc autoscaling production.

## 6. Actor chuẩn

### A01 -- Người dùng

Người dùng là cá nhân có danh tính trong hệ thống. Người dùng có thể tạo tổ chức hoặc được gán vào tổ chức với một vai trò cụ thể.

### A02 -- Quản trị viên tổ chức

Quản trị viên tổ chức quản lý tổ chức, ví, đội ngũ, lập trình viên, Agent, ngân sách và báo cáo trong phạm vi tổ chức.

### A03 -- Lập trình viên

Lập trình viên thuộc một đội ngũ trong tổ chức, có thể đăng ký hoặc quản lý AI Agent theo quyền được cấp. Lập trình viên nhận hạn mức từ đội ngũ, cấp hạn mức cho các Agent do mình quản lý và nhận API key để tích hợp Agent bên ngoài với Gateway.

### A04 -- AI Agent

AI Agent là phần mềm bên ngoài hệ thống. Agent dùng API key do Asymptotic cấp để gọi Gateway. Agent không trực tiếp quản lý ngân sách và không biết provider credential nội bộ.

### A05 -- Quản trị viên hệ thống

Quản trị viên hệ thống quản lý cấu hình toàn cục như provider, model, endpoint, biểu giá, credential nội bộ và chính sách hệ thống.

### A06 -- AI Provider

AI Provider là hệ thống bên ngoài cung cấp dịch vụ AI trả phí. Gateway gọi AI Provider bằng credential nội bộ.

### A07 -- Payment Provider

Payment Provider là hệ thống bên ngoài xử lý thanh toán hoặc xác nhận top-up. Trong MVP, có thể chỉ mô hình hóa hoặc triển khai đơn giản.

## 7. Use case chuẩn

### UC01 -- Thực hiện yêu cầu AI qua Gateway

- **Primary Actor:** AI Agent.
- **Supporting Actor:** AI Provider.
- **Mục tiêu:** Agent gửi request qua Gateway để nhận kết quả AI trong phạm vi quyền và ngân sách được cấp.
- **Must Have:** Có.

### UC02 -- Nạp tiền vào ví tổ chức

- **Primary Actor:** Quản trị viên tổ chức.
- **Supporting Actor:** Payment Provider.
- **Mục tiêu:** Tổ chức nạp tiền để có ngân sách sử dụng dịch vụ AI.
- **Must Have:** Có trong thiết kế; triển khai production có thể nằm ngoài MVP.

### UC03 -- Quản lý và phân bổ ngân sách nội bộ

- **Primary Actor:** Quản trị viên tổ chức, Lập trình viên.
- **Mục tiêu:** Quản trị viên tổ chức phân bổ hoặc thu hồi ngân sách giữa tổ chức, đội ngũ và lập trình viên; lập trình viên phân bổ hoặc thu hồi hạn mức của các AI Agent do mình quản lý.
- **Must Have:** Có.
- **Ghi chú:** Ví tổ chức là nơi lưu tiền thật. Team, Developer và Agent là các cấp hạn mức; việc phân bổ không tạo ví độc lập.

### UC04 -- Đăng ký và quản lý AI Agent

- **Primary Actor:** Lập trình viên.
- **Supporting Actor:** Quản trị viên tổ chức.
- **Mục tiêu:** Lập trình viên đăng ký AI Agent bên ngoài do mình phát triển hoặc quản lý; cập nhật thông tin và trạng thái Agent. Quản trị viên tổ chức giám sát và thực hiện bàn giao khi cần.
- **Must Have:** Có.
- **Ghi chú:** Use case này cần tồn tại rõ ràng. API key thuộc UC05; phân bổ hạn mức cho Agent thuộc UC03.

### UC05 -- Quản lý API key của Agent

- **Primary Actor:** Lập trình viên hoặc Quản trị viên tổ chức.
- **Mục tiêu:** Tạo, xem trạng thái, thu hồi API key do Asymptotic cấp cho Agent.
- **Must Have:** Có.

### UC06 -- Theo dõi giao dịch, usage, cost và trace

- **Primary Actor:** Quản trị viên tổ chức, Lập trình viên.
- **Mục tiêu:** Kiểm tra lịch sử request, chi phí, usage, transaction và trace trong phạm vi quyền.
- **Must Have:** Có.

### UC07 -- Đăng ký tổ chức

- **Primary Actor:** Người dùng.
- **Mục tiêu:** Tạo tổ chức mới và khởi tạo vai trò quản trị viên tổ chức.
- **Must Have:** Có.

### UC08 -- Quản lý đội ngũ và thành viên tổ chức

- **Primary Actor:** Quản trị viên tổ chức.
- **Mục tiêu:** Quản lý đội ngũ, thành viên, quan hệ lập trình viên–đội ngũ và việc bàn giao Agent khi thay đổi cơ cấu.
- **Must Have:** Có.

### UC09 -- Quản lý provider, model và chính sách giá

- **Primary Actor:** Quản trị viên hệ thống.
- **Mục tiêu:** Cấu hình provider, model, endpoint, credential nội bộ, biểu giá và fallback policy.
- **Must Have:** Có.

## 8. Functional Requirements chuẩn

### FR01 -- Quản lý tổ chức và ví/ngân sách

Hệ thống cho phép ghi nhận tổ chức sử dụng dịch vụ, ví tổ chức và ngân sách phân cấp theo đội ngũ, lập trình viên và AI Agent; đồng thời ghi nhận các giao dịch nạp, phân bổ, thu hồi và chi phí tương ứng.

### FR02 -- Đăng ký và quản lý AI Agent

Hệ thống cho phép tổ chức hoặc người được phân quyền đăng ký AI Agent bên ngoài, gắn Agent với tổ chức, lập trình viên quản lý, trạng thái hoạt động và lịch sử sử dụng.

### FR03 -- Cấp và xác thực API key cho Agent

Hệ thống cấp API key cho Agent, lưu khóa ở dạng bảo vệ và dùng khóa này để xác thực request gửi vào Gateway.

### FR04 -- Quản lý provider, model và biểu giá

Hệ thống lưu cấu hình AI Provider, model, endpoint, credential nội bộ và đơn giá hoặc mức phí dự kiến.

### FR05 -- Thiết lập chính sách sử dụng

Hệ thống hỗ trợ hạn mức, quota, budget hoặc policy theo tổ chức, đội ngũ, lập trình viên và Agent. Lập trình viên chỉ được cấp hạn mức cho các Agent do mình quản lý trong phạm vi hạn mức còn khả dụng.

### FR06 -- Kiểm soát chi phí trước thực thi

Hệ thống ước lượng hoặc xác định chi phí dự kiến và chỉ chuyển tiếp request khi ví tổ chức cùng hạn mức của đội ngũ, lập trình viên và Agent đều cho phép.

### FR07 -- Định tuyến request tới AI Provider

Gateway chuyển tiếp request tới AI Provider bằng provider credential nội bộ, không yêu cầu Agent dùng trực tiếp khóa của provider.

### FR08 -- Kiểm soát lũy đẳng

Request gửi lại với cùng khóa lũy đẳng không được tạo thêm giao dịch hoặc chi phí trùng.

### FR09 -- Ghi nhận usage, cost và trace

Hệ thống lưu lịch sử request, trạng thái xử lý, tổ chức, đội ngũ, lập trình viên, Agent, provider, model, usage, cost, transaction và trace để phục vụ kiểm tra, báo cáo và đối soát.

### FR10 -- Trả lỗi khi vượt giới hạn

Hệ thống trả lỗi rõ ràng khi request bị từ chối do sai khóa, đường dẫn tổ chức–đội ngũ–lập trình viên–Agent không hợp lệ, một cấp không đủ hạn mức, vượt quota hoặc vi phạm policy.

## 9. Non-functional Requirements chuẩn

### NFR01 -- Tính nhất quán tài chính

Số dư hoặc ngân sách khả dụng không được bị sử dụng vượt mức. Cùng một khóa lũy đẳng không được tạo nhiều giao dịch thành công.

### NFR02 -- Bảo mật khóa truy cập

API key của Agent, mật khẩu và provider credential không được lưu ở dạng rõ. Các điểm cuối quản trị phải được phân quyền.

### NFR03 -- Khả năng truy vết

Mỗi request cần có đủ thông tin để xác định tổ chức, đội ngũ, lập trình viên, Agent, API key, provider, model, chi phí, trạng thái và thời điểm xử lý.

### NFR04 -- Khả năng phục hồi

Lỗi cơ sở dữ liệu, lỗi provider hoặc lỗi mạng không được để lại trạng thái tài chính không thể đối soát.

### NFR05 -- Hiệu năng

Thời gian xử lý tại Gateway cần được tách biệt với thời gian chờ provider để đánh giá đúng chi phí độ trễ của lớp kiểm soát.

## 10. Traceability chuẩn giữa Use Case và FR/NFR

| Use Case | FR liên quan | NFR liên quan |
|---|---|---|
| UC01 -- Thực hiện yêu cầu AI qua Gateway | FR03, FR04, FR05, FR06, FR07, FR08, FR09, FR10 | NFR01, NFR02, NFR03, NFR04, NFR05 |
| UC02 -- Nạp tiền vào ví tổ chức | FR01, FR09 | NFR01, NFR02, NFR03, NFR04 |
| UC03 -- Quản lý và phân bổ ngân sách nội bộ | FR01, FR05, FR09, FR10 | NFR01, NFR02, NFR03, NFR04 |
| UC04 -- Đăng ký và quản lý AI Agent | FR02, FR05, FR09, FR10 | NFR02, NFR03 |
| UC05 -- Quản lý API key của Agent | FR02, FR03, FR09, FR10 | NFR02, NFR03 |
| UC06 -- Theo dõi giao dịch, usage, cost và trace | FR01, FR02, FR09 | NFR01, NFR02, NFR03 |
| UC07 -- Đăng ký tổ chức | FR01, FR09 | NFR01, NFR02, NFR03 |
| UC08 -- Quản lý đội ngũ và thành viên tổ chức | FR01, FR02, FR05, FR09, FR10 | NFR01, NFR02, NFR03, NFR04 |
| UC09 -- Quản lý provider, model và chính sách giá | FR04, FR05, FR07, FR09 | NFR02, NFR03, NFR04 |

## 11. Diagram chuẩn cần bám theo

### Chương 2 -- Phân tích yêu cầu

- Use Case Diagram tổng quát.
- Use Case Diagram nhóm Agent/Gateway.
- Use Case Diagram nhóm Finance/Admin.
- Activity Diagram UC01 -- Agent gọi AI qua Gateway.
- Activity Diagram đăng ký và quản lý AI Agent.
- Activity Diagram nạp tiền và phân bổ ngân sách.

### Chương 3 -- Phân tích hướng đối tượng

- Analysis Class Diagram tổng quát.
- Boundary Class Diagram.
- Control Class Diagram.
- Entity Class Diagram.
- Robustness Diagram UC01.
- Robustness Diagram đăng ký và quản lý AI Agent.

### Chương 4 -- Thiết kế hướng đối tượng

- Package Diagram.
- Design Class Diagram tổng quát.
- Design Class Diagram cho Gateway Request Flow.
- Design Class Diagram cho Finance/Ledger.
- Sequence Diagram UC01.
- Sequence Diagram đăng ký và quản lý AI Agent.
- Sequence Diagram nạp tiền và phân bổ ngân sách.
- State Machine Diagram cho AI Request.
- State Machine Diagram cho Financial Transaction hoặc Budget Reservation.

### Chương 5 -- Triển khai, kiểm thử và đánh giá

Các artifact sau là tùy chọn và chỉ được thêm khi có dữ liệu triển khai thực tế:

- As-built Component Diagram.
- As-built Deployment Diagram.
- Ma trận Requirement--Use Case--Module--Test.
- Bảng test case và kết quả.
- Biểu đồ hiệu năng, tỷ lệ pass/fail hoặc chi phí.

Component/Deployment Diagram ở Chương 5 phải phản ánh hệ thống đã triển khai, không lặp lại kiến trúc dự kiến của Chương 4. Ma trận, bảng và biểu đồ kết quả kiểm thử không phải UML diagram. Không dựng thêm microservice, cloud cluster hoặc node chưa tồn tại để hoàn thiện hình thức báo cáo.

## 12. Cấu trúc báo cáo chuẩn

### Chương 1 -- Giới thiệu đề tài

- Đặt vấn đề.
- Bài toán cần giải quyết.
- Mục tiêu và phạm vi.
- Phương pháp thực hiện.
- Bố cục đồ án.

Ghi chú: Chương 1 là chương mở đầu nên có thể ít mục hơn các chương phân tích và thiết kế.

### Chương 2 -- Khảo sát và phân tích yêu cầu

- Khảo sát hiện trạng.
- Đặc tả yêu cầu hệ thống.
- Phân tích actor.
- Phân tích trường hợp sử dụng.
- Đặc tả trường hợp sử dụng.
- Mô hình hóa luồng nghiệp vụ.

Ghi chú: Chương 2 cần tách actor, use case diagram và use case specification thành các mục riêng để đúng quy trình phân tích yêu cầu và giúp bố cục cân đối hơn.

### Chương 3 -- Phân tích hướng đối tượng

- Cơ sở phân tích hướng đối tượng.
- Nhận diện lớp phân tích.
- Phân tích lớp biên.
- Phân tích lớp điều khiển.
- Phân tích lớp thực thể.
- Đối chiếu lớp phân tích với use case.

Ghi chú: Chương 3 là trọng tâm OOAD ở mức phân tích, cần làm rõ Boundary-Control-Entity và trace từ use case sang lớp phân tích.

### Chương 4 -- Thiết kế hướng đối tượng

- Kiến trúc tổng thể.
- Thiết kế gói.
- Thiết kế lớp.
- Thiết kế tương tác.
- Thiết kế trạng thái.
- Thiết kế dữ liệu và giao diện API.

### Chương 5 -- Triển khai, kiểm thử và đánh giá

- Môi trường và công nghệ triển khai.
- Các chức năng đã triển khai.
- Kiểm thử hệ thống.
- Đánh giá kết quả.
- Hạn chế triển khai.

### Chương 6 -- Kết luận và hướng phát triển

- Kết quả đạt được.
- Hạn chế của đồ án.
- Hướng phát triển.

## 13. Quy tắc khi chỉnh báo cáo

Khi chỉnh bất kỳ chương nào, cần kiểm tra:

1. Nội dung có còn đúng định vị **AI Agent Financial Gateway** không?
2. Có vô tình viết như hệ thống tạo/huấn luyện/điều phối AI Agent không?
3. Có vô tình viết như Payment Gateway production không?
4. Actor có dùng đúng danh sách chuẩn không?
5. Use case có trace được về FR/NFR không?
6. Class analysis có xuất phát từ use case không?
7. Sequence/state/design class có xuất phát từ analysis model không?
8. Nội dung triển khai có bị đưa quá sớm vào phần phân tích không?
9. Chương 4 có tập trung vào thiết kế, còn Chương 5 mới tập trung vào công nghệ/triển khai không?

## 14. Các quyết định đã chốt

- Dùng tên đề tài tiếng Việt: **Cổng tài chính thời gian thực cho tác nhân trí tuệ nhân tạo**.
- Dùng tên hệ thống: **Asymptotic -- AI Agent Financial Gateway**.
- Hệ thống **không tạo AI Agent**, chỉ đăng ký AI Agent bên ngoài vào Gateway.
- Agent dùng **API key do Asymptotic cấp**.
- Gateway dùng **provider credential nội bộ** để gọi AI Provider.
- Trọng tâm nghiệp vụ là **kiểm soát chi phí trước khi request phát sinh chi phí ở AI Provider**.
- MVP có thể mô hình hóa payment/top-up, nhưng không cần triển khai Payment Gateway production đầy đủ.
- Use case **Đăng ký và quản lý AI Agent** phải tồn tại rõ ràng trong mô hình.
- Chương 3 và Chương 4 phải thể hiện rõ quy trình OOAD, không chỉ mô tả backend implementation.

## 15. Quy tắc hình thức đồ án

Các quy tắc dưới đây dùng để kiểm tra hình thức báo cáo trước khi nộp.

### 15.1 Tài liệu tham khảo và trích dẫn

- Danh mục tài liệu tham khảo không cần phân loại tiếng Việt và tiếng Anh.
- Mỗi tài liệu tham khảo cần có tối thiểu các thông tin:
  - tên tác giả;
  - tên tài liệu;
  - nhà xuất bản hoặc nơi công bố;
  - năm xuất bản.
- Tài liệu tham khảo đã liệt kê trong danh mục phải được trích dẫn trong nội dung đồ án.
- Không để tài liệu trong danh mục tham khảo nếu không được trích dẫn trong báo cáo.
- Nếu sử dụng dữ liệu hoặc thông tin từ một website, trích dẫn trực tiếp website đó tại vị trí sử dụng; không đưa website đó vào danh mục tài liệu tham khảo nếu không được yêu cầu bởi mẫu báo cáo.

### 15.2 Hình ảnh, biểu đồ và nguồn

- Mọi hình ảnh, biểu đồ hoặc bảng nếu không phải do tác giả tự tạo ra phải có trích dẫn nguồn.
- Hình ảnh, biểu đồ hoặc bảng do tác giả tự xây dựng không cần ghi dòng nguồn trong caption.
- Hình ảnh phải đủ rõ, không vỡ chữ, không méo tỉ lệ.
- Tên hình và tên bảng phải thống nhất cách viết trong toàn bộ báo cáo.

### 15.3 Chính tả và câu văn

- Không có lỗi chính tả.
- Dấu chấm, dấu phẩy, dấu hai chấm và dấu chấm phẩy phải đặt đúng chỗ.
- Câu văn cần có đủ chủ ngữ và vị ngữ.
- Tránh câu quá dài, nhiều vế liệt kê khó đọc.
- Thuật ngữ phải dùng nhất quán trong toàn bộ báo cáo.

### 15.4 Cách dùng thuật ngữ tiếng Anh

- Không giải thích thuật ngữ tiếng Anh theo kiểu: **Học máy (machine learning) là...**
- Nếu cần dùng thuật ngữ tiếng Anh, ưu tiên một trong hai cách:
  - dùng tiếng Việt nếu thuật ngữ đã phổ biến và không gây mơ hồ;
  - dùng tiếng Anh trực tiếp nếu đó là tên kỹ thuật, pattern, API hoặc thuật ngữ chuyên ngành khó dịch.
- Khi đã chọn một cách gọi, dùng nhất quán trong toàn bộ báo cáo.
- Không lạm dụng tiếng Anh trong câu tiếng Việt nếu có thể viết rõ bằng tiếng Việt.

### 15.5 Bố cục và độ cân đối

- Bố cục các chương cần cân đối, tránh một chương quá dài trong khi chương khác quá ngắn.
- Các chương nên có số trang tương đối gần nhau, tùy vai trò của từng chương.
- Số lượng tiểu mục trong các mục cùng cấp nên tương đối cân bằng.
- Không tạo quá nhiều tiểu mục ngắn chỉ có một hoặc hai câu.
- Mỗi chương nên có phần dẫn nhập hoặc kết luận chương nếu mẫu báo cáo yêu cầu hoặc nếu cần giúp mạch đọc rõ hơn.

### 15.6 Mẫu bìa và định dạng

- Sử dụng đúng mẫu bìa do viện và trường quy định.
- Thông tin trên bìa phải thống nhất với tên đề tài đã chốt.
- Tên giảng viên hướng dẫn, sinh viên, mã số sinh viên, đơn vị đào tạo và thời gian phải chính xác.
- Định dạng font, lề, đánh số chương/mục, danh sách hình, danh sách bảng và tài liệu tham khảo phải thống nhất theo mẫu báo cáo.
