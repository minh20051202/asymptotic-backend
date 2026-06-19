# Research Brief: Database Design, Requirements Engineering và ERD ở mức thiết kế

Ngày: 2026-06-18  
Phạm vi: Thiết kế dữ liệu cho báo cáo OOAD, tập trung vào Requirements Engineering, mô hình dữ liệu khái niệm, mô hình dữ liệu logic và ERD. Không đi vào chi tiết triển khai như SQL, index, migration, partitioning, tuning hoặc cấu hình DBMS.

## 1. Kết luận ngắn

Thiết kế cơ sở dữ liệu trong đồ án OOAD không nên bắt đầu từ bảng dữ liệu. Hướng hợp lý hơn là đi theo chuỗi:

Yêu cầu nghiệp vụ → Use case → Lớp thực thể phân tích → Mô hình dữ liệu khái niệm → ERD logic → Thiết kế dữ liệu ở Chương 4.

Với đồ án Asymptotic, ERD nên được đặt ở Chương 4 như một thiết kế dữ liệu dẫn xuất từ Chương 2 và Chương 3. ERD không thay thế Analysis Class Diagram, Design Class Diagram hoặc Sequence Diagram. Nó chỉ trả lời câu hỏi: hệ thống cần lưu những thực thể dữ liệu nào, chúng liên hệ với nhau ra sao, và ràng buộc nghiệp vụ nào cần được bảo toàn ở mức dữ liệu.

## 2. Vai trò của Requirements Engineering trong thiết kế dữ liệu

IEEE/ISO/IEC 29148-2018 mô tả Requirements Engineering là tập hợp các quy trình và sản phẩm liên quan đến yêu cầu của hệ thống/phần mềm trong suốt vòng đời; tiêu chuẩn này nhấn mạnh cấu trúc của một yêu cầu tốt, thuộc tính của yêu cầu và việc áp dụng lặp lại các quy trình yêu cầu trong vòng đời hệ thống [IEEE 29148-2018](https://standards.ieee.org/ieee/29148/6937/).

Áp dụng vào thiết kế dữ liệu:

- Yêu cầu chức năng xác định hệ thống cần ghi nhận dữ liệu gì.
- Use case xác định dữ liệu được tạo, đọc, cập nhật hoặc dùng để ra quyết định ở từng luồng.
- Tiền điều kiện, hậu điều kiện và business rule xác định ràng buộc dữ liệu.
- Yêu cầu phi chức năng như nhất quán tài chính, bảo mật khóa truy cập và truy vết xác định các nhóm dữ liệu bắt buộc phải tồn tại.
- Traceability giúp chứng minh mỗi bảng/nhóm thực thể dữ liệu đều có nguồn gốc từ yêu cầu hoặc use case, tránh thiết kế bảng theo cảm tính.

Vì vậy, trong báo cáo, phần thiết kế dữ liệu nên ghi rõ ERD được dẫn xuất từ FR/NFR, use case specification, activity diagram và các lớp Entity ở Chương 3.

## 3. Phân biệt conceptual, logical và physical data model

IBM mô tả data modeling là quá trình tạo biểu diễn trực quan cho hệ thống thông tin hoặc một phần hệ thống, nhằm truyền đạt các kiểu dữ liệu được sử dụng/lưu trữ và quan hệ giữa chúng. IBM cũng nêu quy trình data modeling thường bắt đầu từ yêu cầu nghiệp vụ, sau đó chuyển business rule thành cấu trúc dữ liệu [IBM, Data Modeling](https://www.ibm.com/think/topics/data-modeling).

Ba mức mô hình nên tách rõ:

- Conceptual data model: Mức khái niệm, dùng để xác định các khái niệm nghiệp vụ chính và quan hệ lớn giữa chúng. Mô hình này phù hợp để trao đổi với stakeholder và liên kết với phân tích nghiệp vụ.
- Logical data model: Mức logic, bổ sung thuộc tính, định danh, quan hệ, bội số và ràng buộc nghiệp vụ. Mô hình này vẫn nên độc lập với DBMS cụ thể.
- Physical data model: Mức vật lý, mô tả cách dữ liệu được cài đặt trong DBMS cụ thể, gồm kiểu dữ liệu cụ thể, index, tối ưu hiệu năng và các chi tiết lưu trữ.

Đối với đồ án này, phần ERD trong Chương 4 nên dừng ở mức logical design. Có thể nêu khóa chính/khóa ngoại và ràng buộc chính để mô hình đủ rõ, nhưng không nên đi sâu vào index, câu lệnh SQL, migration hoặc tuning.

## 4. ERD nên thể hiện điều gì

IBM định nghĩa ERD là biểu diễn trực quan cách các đối tượng trong cơ sở dữ liệu liên hệ với nhau. ERD dùng các thành phần như entity, attribute, relationship và cardinality; nó thường đóng vai trò mô hình dữ liệu khái niệm cấp cao, chuẩn bị cho thiết kế dữ liệu chi tiết hơn [IBM, Entity Relationship Diagram](https://www.ibm.com/think/topics/entity-relationship-diagram).

Trong báo cáo OOAD, một ERD tốt nên thể hiện:

- Entity: Các đối tượng nghiệp vụ cần lưu dữ liệu, ví dụ Organization, Team, Developer, AIAgent, ApiKey, Wallet, BudgetLimit, Provider, AIModel, RequestTrace, FinancialTransaction.
- Attribute: Thuộc tính quan trọng để hiểu nghiệp vụ, không cần liệt kê mọi cột kỹ thuật.
- Identifier: Định danh của thực thể, đủ để phân biệt bản ghi.
- Relationship: Quan hệ nghiệp vụ giữa các entity.
- Cardinality: Bội số quan hệ, ví dụ một Organization có nhiều Team, một Team có nhiều Developer, một Developer quản lý nhiều AIAgent.
- Constraint: Ràng buộc quan trọng, ví dụ Agent thuộc một Organization, ApiKey thuộc một Agent, request trace phải gắn với Agent và model được dùng.
- Associative entity: Thực thể trung gian khi quan hệ nhiều-nhiều hoặc quan hệ có thuộc tính riêng, ví dụ membership/history, allocation, ledger entry.

ERD không nên thể hiện:

- Controller, service, repository, DTO hoặc adapter.
- Luồng xử lý theo thời gian; phần này thuộc activity/sequence diagram.
- Trạng thái vòng đời phức tạp; phần này thuộc state machine diagram.
- Tối ưu vật lý như index, partition, connection pool, query plan.

## 5. Nền tảng lý thuyết của ER model

Peter P. Chen là tác giả bài báo kinh điển “The Entity-Relationship Model - Toward a Unified View of Data” đăng trên ACM Transactions on Database Systems năm 1976. DBLP ghi nhận bài báo này là journal article, DOI `10.1145/320434.320440`, trang 9-36 [DBLP, Chen 1976](https://dblp.org/rec/journals/tods/Chen76.html).

Giá trị chính của ER model là mô tả thế giới nghiệp vụ bằng entity và relationship trước khi chuyển sang cấu trúc lưu trữ. Đây là lý do ERD phù hợp với phần thiết kế dữ liệu của đồ án: nó buộc người thiết kế mô tả quan hệ nghiệp vụ trước, thay vì nhảy thẳng sang bảng.

## 6. Quy trình đề xuất để vẽ ERD cho đồ án

### Bước 1: Gom nguồn yêu cầu

Nguồn đầu vào nên lấy theo thứ tự:

1. `project_source_of_truth.md`.
2. FR/NFR trong Chương 2.
3. Use case specification UC01-UC09.
4. Activity diagram của các nghiệp vụ chính.
5. Analysis Entity classes ở Chương 3.
6. Code/schema hiện có chỉ dùng để đối chiếu, không dùng làm nguồn quyết định.

### Bước 2: Trích danh từ nghiệp vụ thành entity ứng viên

Từ use case và FR/NFR, trích các khái niệm có định danh, trạng thái hoặc vòng đời riêng. Với Asymptotic, nhóm entity ứng viên gồm:

- Organization, Team, User, DeveloperProfile, OrganizationMembership.
- AIAgent, ApiKey, AgentStatus.
- Wallet, BudgetLimit, BudgetAllocation, BudgetReservation.
- FinancialTransaction, LedgerEntry.
- Provider, ProviderCredential, AIModel, ModelPricing, RoutingPolicy.
- AIRequest, IdempotencyRecord, RequestTrace, UsageRecord, CostRecord.

### Bước 3: Loại bỏ lớp xử lý khỏi ERD

Các khái niệm như GatewayRoutingControl, BudgetControl, PolicyService, ProviderAdapter, ApiKeyService không phải entity dữ liệu. Chúng thuộc class/sequence design, không đưa vào ERD.

### Bước 4: Xác định quan hệ và bội số

Mỗi quan hệ cần trả lời:

- Một bản ghi phía A liên hệ với bao nhiêu bản ghi phía B?
- Quan hệ này bắt buộc hay tùy chọn?
- Quan hệ có thuộc tính riêng không?
- Quan hệ có cần lịch sử không?
- Nếu thay đổi quan hệ, dữ liệu lịch sử có cần giữ nguyên không?

Ví dụ:

- Một Organization có nhiều Team.
- Một Team có nhiều DeveloperProfile theo thời điểm; nếu cần lịch sử chuyển team, dùng membership/history thay vì gắn trực tiếp cứng.
- Một DeveloperProfile quản lý nhiều AIAgent.
- Một AIAgent có nhiều ApiKey theo vòng đời tạo/thu hồi.
- Một AIRequest tạo một RequestTrace và có thể tạo BudgetReservation, CostRecord, LedgerEntry.

### Bước 5: Xác định ràng buộc nghiệp vụ

Các ràng buộc phải bám yêu cầu:

- Agent không được gọi provider trực tiếp bằng provider credential.
- ApiKey chỉ xác thực được Agent trong trạng thái hợp lệ.
- Request chỉ được route khi Organization, Team, Developer và Agent còn hạn mức hợp lệ.
- Idempotency key không được tạo chi phí trùng cho cùng request.
- Usage/cost/trace phải đủ để truy vết theo organization, team, developer, agent, provider và model.
- Ví tổ chức là nơi giữ tiền thật; Team, Developer và Agent là cấp hạn mức, không phải ví độc lập.

### Bước 6: Chia ERD thành nhiều sơ đồ nhỏ nếu quá nhiều entity

Vì hệ thống có nhiều nhóm nghiệp vụ, không nên ép toàn bộ entity vào một ERD lớn gây rối. Nên chia thành các ERD con:

- ERD tổ chức và phân quyền: Organization, Team, User, DeveloperProfile, Membership.
- ERD Agent và API key: AIAgent, ApiKey, AgentStatus.
- ERD tài chính và hạn mức: Wallet, BudgetLimit, Allocation, Reservation, Transaction, LedgerEntry.
- ERD provider/model/pricing: Provider, ProviderCredential, AIModel, ModelPricing, RoutingPolicy.
- ERD request/usage/trace: AIRequest, IdempotencyRecord, RequestTrace, UsageRecord, CostRecord.

Nếu báo cáo cần một hình tổng quan, chỉ nên vẽ ERD overview có entity nhóm chính và quan hệ lớn; các thuộc tính chi tiết để ở ERD con.

## 7. Áp dụng vào Asymptotic

### Vị trí trong báo cáo

ERD nên nằm ở Chương 4, trong phần thiết kế dữ liệu. Chương 3 chỉ nên nhận diện Entity class ở mức phân tích nghiệp vụ. Chương 4 mới chuyển các Entity class đó thành mô hình dữ liệu logic.

### Cách diễn đạt trong báo cáo

Có thể viết ngắn như sau:

> Mô hình dữ liệu của hệ thống được dẫn xuất từ các yêu cầu chức năng, đặc tả use case và các lớp thực thể ở giai đoạn phân tích. ERD trong mục này tập trung vào cấu trúc dữ liệu logic, quan hệ nghiệp vụ và các ràng buộc toàn vẹn chính. Các chi tiết triển khai vật lý như chỉ mục, tối ưu truy vấn hoặc migration không thuộc phạm vi trình bày của đồ án.

### Các câu hỏi kiểm tra trước khi chốt ERD

- Mỗi entity trong ERD có truy vết được về FR/NFR hoặc use case không?
- Có entity nào thực chất là hành động xử lý, service hoặc adapter không?
- Có quan hệ nhiều-nhiều nào cần associative entity không?
- Có quan hệ nào cần lưu lịch sử thay vì chỉ lưu trạng thái hiện tại không?
- Có ràng buộc tài chính nào cần thể hiện bằng entity riêng như LedgerEntry hoặc BudgetReservation không?
- Có đang nhầm ví tổ chức với hạn mức team/developer/agent không?
- Có đang đưa chi tiết triển khai vào ERD quá sớm không?

## 8. Khuyến nghị cho kế hoạch tiếp theo

Đối với Chương 4, nên vẽ ERD theo hướng “logical design” và chia nhỏ theo cụm nghiệp vụ. Trình tự tốt nhất:

1. Hoàn thiện Analysis Entity classes ở Chương 3.
2. Lập bảng ánh xạ Entity class → Data entity.
3. Vẽ ERD con theo từng cụm nghiệp vụ.
4. Viết đoạn giải thích ngắn cho từng ERD: mục đích, entity chính, quan hệ chính, ràng buộc chính.
5. Kiểm tra traceability từ ERD về FR/NFR và UC.

Không nên bắt đầu từ schema code hiện có, vì như vậy ERD dễ trở thành tài liệu reverse-engineering thay vì thiết kế hướng đối tượng có truy vết yêu cầu.

## 9. Nguồn tham khảo

- IEEE Standards Association. “IEEE/ISO/IEC 29148-2018: Systems and software engineering -- Life cycle processes -- Requirements engineering.” 2018. https://standards.ieee.org/ieee/29148/6937/
- IBM. “What is data modeling?” IBM Think. https://www.ibm.com/think/topics/data-modeling
- IBM. “What is an entity relationship diagram?” IBM Think. https://www.ibm.com/think/topics/entity-relationship-diagram
- Chen, Peter P. “The Entity-Relationship Model - Toward a Unified View of Data.” ACM Transactions on Database Systems, 1(1), 9-36, 1976. DOI: 10.1145/320434.320440. Metadata: https://dblp.org/rec/journals/tods/Chen76.html
