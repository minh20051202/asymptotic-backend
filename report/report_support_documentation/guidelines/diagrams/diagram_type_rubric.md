# Rubric đánh giá từng loại diagram OOAD

Tài liệu này dùng để review chi tiết từng loại biểu đồ trong đồ án OOAD. Mỗi loại diagram có mục đích khác nhau, vì vậy tiêu chí đánh giá cũng khác nhau.

Ký pháp và hướng quan hệ phải tuân theo `uml_2_5_1_drawing_rules.md`.

Phân loại:

- Analysis Class và Design Class là hai cách dùng UML Class Diagram.
- Robustness/BCE là kỹ thuật phương pháp OOAD, không phải loại diagram độc lập của OMG UML.
- ERD và biểu đồ kết quả kiểm thử không thuộc UML.

## 1. Use Case Diagram

### Mục đích đúng

Use Case Diagram trả lời câu hỏi: **Actor nào tương tác với hệ thống để đạt mục tiêu gì?**

### Tiêu chí đạt

- Có system boundary rõ ràng.
- Actor nằm ngoài system boundary.
- Use case nằm trong system boundary.
- Use case được đặt tên bằng động từ + bổ ngữ, thể hiện mục tiêu của actor.
- Actor chính và actor phụ được phân biệt hợp lý.
- Quan hệ `include`, `extend`, generalization nếu dùng phải có lý do rõ ràng.
- Mỗi use case tạo ra kết quả quan sát được cho actor hoặc stakeholder.
- Không biến từng API endpoint thành use case.
- Không biến bước trong flow, điều kiện rẽ nhánh hoặc thuật toán thành use case.
- Không đưa class, database table hoặc module kỹ thuật vào use case diagram.

### Lỗi thường gặp

- Actor đặt bên trong hệ thống.
- Use case quá nhỏ như "click button", "call API", "validate token".
- Use case quá kỹ thuật, không phản ánh mục tiêu nghiệp vụ.
- Sơ đồ chi tiết bị vẽ như Activity Diagram bằng các oval.
- Lạm dụng `include` và `extend`.
- Dùng relationship không chuẩn như `<<precondition>>`.
- Thiếu actor hệ thống ngoài như Payment Provider hoặc AI Provider.

### Checklist

- [ ] Boundary có tên hệ thống.
- [ ] Actor đúng danh sách chuẩn.
- [ ] Use case khớp danh sách use case chuẩn.
- [ ] Mỗi use case có mục tiêu và kết quả quan sát được.
- [ ] Không có chi tiết database/service.
- [ ] Không có bước kỹ thuật của UC01 bị tách thành use case riêng.
- [ ] Quan hệ include/extend có thể giải thích được.
- [ ] Không dùng relationship tự tạo ngoài UML chuẩn.

## 2. Use Case Specification

### Mục đích đúng

Use Case Specification mô tả chi tiết kịch bản thực hiện một use case.

### Tiêu chí đạt

- Có ID, tên, mô tả, actor, trigger.
- Có pre-condition và post-condition.
- Luồng chính rõ ràng, đánh số tuần tự.
- Mỗi bước có chủ thể rõ và chỉ chứa một hành động hoặc phản hồi chính.
- Luồng thay thế và ngoại lệ không trộn lẫn với luồng chính.
- Mỗi nhánh tham chiếu bước gốc và nêu rõ tiếp tục, thành công hoặc thất bại.
- Có business rules nếu use case liên quan chính sách nghiệp vụ.
- Có traceability về FR/NFR.

### Lỗi thường gặp

- Luồng chính viết như mô tả code.
- Basic flow chứa quá nhiều `if/else` hoặc chi tiết retry/fallback.
- Thiếu failure post-condition.
- Luồng ngoại lệ không nói rõ hệ thống rollback/giữ nguyên/trả lỗi ra sao.
- Nhánh không nói rõ Use Case tiếp tục ở bước nào hay kết thúc thế nào.
- Traceability sai với FR/NFR hiện hành.

### Checklist

- [ ] ID và tên use case đúng chuẩn.
- [ ] Actor đúng.
- [ ] Pre-condition và post-condition đủ rõ.
- [ ] Basic flow không quá kỹ thuật.
- [ ] Basic flow dùng câu chủ động và có chủ thể rõ.
- [ ] Alternative flow vẫn dẫn tới mục tiêu thành công.
- [ ] Exception flow có trạng thái kết thúc rõ.
- [ ] Use case tài chính nêu rõ trạng thái tiền/ngân sách sau lỗi.
- [ ] Traceability đúng source of truth.

## 3. Activity Diagram

### Mục đích đúng

Activity Diagram trả lời câu hỏi: **Quy trình hoặc use case diễn ra theo các bước và nhánh điều kiện nào?**

### Tiêu chí đạt

- Có điểm bắt đầu và kết thúc.
- Các action đặt tên bằng động từ.
- Decision node có điều kiện rõ.
- Nhánh lỗi/ngoại lệ quan trọng được thể hiện.
- Swimlane được dùng nếu cần phân biệt trách nhiệm giữa actor, system và external system.
- Không mô tả chi tiết message call giữa object như sequence diagram.

### Lỗi thường gặp

- Thiếu nhánh lỗi.
- Decision node không ghi điều kiện.
- Biểu đồ quá tuyến tính, không thể hiện điểm rẽ nhánh quan trọng.
- Vẽ như flowchart kỹ thuật quá chi tiết.
- Không phân biệt trách nhiệm giữa hệ thống và actor.

### Checklist

- [ ] Có start/end.
- [ ] Action rõ nghĩa.
- [ ] Nhánh decision có guard condition.
- [ ] Có luồng lỗi quan trọng.
- [ ] Nếu có nhiều bên tham gia, swimlane hợp lý.

## 4. Robustness Diagram

### Mục đích đúng

Robustness Diagram là cầu nối giữa use case và class analysis. Nó trả lời câu hỏi: **Use case cần boundary, control và entity nào?**

### Tiêu chí đạt

- Actor chỉ giao tiếp với Boundary.
- Boundary giao tiếp với Control.
- Control điều phối Entity.
- Entity không gọi ngược actor.
- Mỗi control nên gắn với một use case hoặc một nhóm trách nhiệm rõ.
- Các lớp trong robustness diagram có thể truy ra analysis class diagram.
- Quy tắc kết nối được ghi là convention OOSE/ICONIX của đồ án, không ghi là cú pháp bắt buộc của UML.

### Lỗi thường gặp

- Actor truy cập trực tiếp Entity.
- Boundary chứa quá nhiều logic nghiệp vụ.
- Không có Control class.
- Entity bị biến thành database table thuần túy.
- Lớp xuất hiện trong robustness nhưng không xuất hiện trong class analysis.

### Checklist

- [ ] Có đủ Boundary-Control-Entity.
- [ ] Actor không nối trực tiếp Entity.
- [ ] Control thể hiện trách nhiệm điều phối.
- [ ] Entity là khái niệm nghiệp vụ.
- [ ] Mapping được sang analysis class diagram.

## 5. Analysis Class Diagram

### Mục đích đúng

Analysis Class Diagram trả lời câu hỏi: **Trong bài toán có những lớp nghiệp vụ nào và chúng liên hệ ra sao?**

### Tiêu chí đạt

- Có phân loại Boundary-Control-Entity hoặc stereotype tương đương khi mô hình áp dụng BCE.
- Tên lớp phản ánh khái niệm nghiệp vụ.
- Quan hệ giữa lớp có ý nghĩa nghiệp vụ.
- Thuộc tính chỉ nên là thuộc tính quan trọng ở mức phân tích.
- Không đưa framework, repository, DTO, database migration vào analysis class.
- Entity có thể liên hệ trực tiếp với Entity khi đó là quan hệ cấu trúc nghiệp vụ.

### Lỗi thường gặp

- Biểu đồ là ERD trá hình.
- Có quá nhiều class kỹ thuật.
- Thiếu Control class cho các use case quan trọng.
- Quan hệ giữa lớp không ghi multiplicity khi cần.
- Class không trace được từ use case.

### Checklist

- [ ] Có Boundary-Control-Entity.
- [ ] Class xuất phát từ use case/khái niệm nghiệp vụ.
- [ ] Không lẫn implementation detail.
- [ ] Multiplicity quan trọng được thể hiện.
- [ ] Class quan trọng có trách nhiệm rõ.

## 6. Package Diagram

### Mục đích đúng

Package Diagram trả lời câu hỏi: **Hệ thống được chia thành các gói/module nào và phụ thuộc ra sao?**

### Tiêu chí đạt

- Package có tên theo domain hoặc responsibility.
- Quan hệ phụ thuộc có hướng rõ.
- Không để mọi package phụ thuộc lẫn nhau.
- Package tầng trên không bị phụ thuộc ngược bởi tầng thấp nếu kiến trúc không cho phép.
- Có thể giải thích được vì sao chia package như vậy.
- Việc đặt theo domain, tránh vòng phụ thuộc và giới hạn chiều phụ thuộc là quy ước kiến trúc, không phải ràng buộc cú pháp UML.

### Lỗi thường gặp

- Package đặt theo folder ngẫu nhiên.
- Quá nhiều package nhỏ.
- Thiếu package cho nghiệp vụ quan trọng.
- Dependency vòng tròn không được giải thích.

### Checklist

- [ ] Package phản ánh domain/responsibility.
- [ ] Dependency có hướng.
- [ ] Không có vòng phụ thuộc bất hợp lý.
- [ ] Package khớp với design class diagram.

## 7. Design Class Diagram

### Mục đích đúng

Design Class Diagram trả lời câu hỏi: **Các lớp thiết kế, service, repository, adapter, interface phối hợp ra sao để hiện thực hệ thống?**

### Tiêu chí đạt

- Có lớp thiết kế gần với implementation nhưng không cần liệt kê toàn bộ code.
- Service, repository, adapter, entity được phân biệt rõ.
- Interface được dùng ở điểm cần đảo chiều phụ thuộc hoặc thay provider.
- Quan hệ dependency, composition, association thể hiện hợp lý.
- Realization dùng từ implementation class tới interface; composition chỉ dùng cho sở hữu vòng đời độc quyền.
- Không quá chi tiết đến từng method phụ nếu làm rối biểu đồ.

### Lỗi thường gặp

- Copy toàn bộ codebase vào class diagram.
- Thiếu interface/adapter cho external provider.
- Repository phụ thuộc ngược vào service.
- Entity chứa logic không phù hợp hoặc service ôm quá nhiều trách nhiệm.

### Checklist

- [ ] Class thiết kế trace được từ analysis class.
- [ ] Service/repository/adapter rõ trách nhiệm.
- [ ] Quan hệ phụ thuộc hợp lý.
- [ ] Không quá chi tiết.
- [ ] Khớp package diagram.

## 8. Sequence Diagram

### Mục đích đúng

Sequence Diagram trả lời câu hỏi: **Các đối tượng tương tác theo thứ tự thời gian như thế nào để thực hiện một use case?**

### Tiêu chí đạt

- Lifeline đúng actor/object/service.
- Message theo đúng thứ tự thời gian.
- Có nhánh `alt`, `opt`, `loop` nếu có điều kiện quan trọng.
- Luồng thành công và lỗi quan trọng được thể hiện.
- External system được đặt rõ là bên ngoài.
- Không biến sequence diagram thành flowchart.
- Reply message chỉ cần xuất hiện khi kết quả trả về có ý nghĩa; không bắt buộc vẽ mọi reply.

### Lỗi thường gặp

- Thiếu return hoặc trạng thái kết quả quan trọng.
- Gọi trực tiếp database từ actor.
- Không thể hiện nhánh lỗi ngân sách/quyền truy cập.
- Message đặt tên mơ hồ như "process".
- Quá nhiều object nhỏ làm diagram khó đọc.

### Checklist

- [ ] Lifeline đúng vai trò.
- [ ] Message đúng thứ tự.
- [ ] Có nhánh lỗi quan trọng.
- [ ] External system được phân biệt.
- [ ] Khớp use case specification.

## 9. State Machine Diagram

### Mục đích đúng

State Machine Diagram trả lời câu hỏi: **Một đối tượng quan trọng đi qua những trạng thái nào trong vòng đời của nó?**

### Tiêu chí đạt

- Có initial state và final state nếu phù hợp.
- Trạng thái mô tả condition ổn định; tên danh từ/tính từ là heuristic, không phải cú pháp bắt buộc.
- Transition dùng dạng tổng quát `trigger [guard] / effect`; mỗi thành phần có thể được lược bỏ.
- Cho phép completion transition không có trigger.
- Có trạng thái lỗi/hủy/rollback nếu nghiệp vụ cần.
- Không vẽ state diagram cho đối tượng không có vòng đời trạng thái đáng kể.

### Lỗi thường gặp

- Trạng thái đặt tên như hành động: "Check Budget", "Call Provider".
- Thiếu trạng thái failed/rejected.
- Không có điều kiện chuyển trạng thái.
- Vẽ state quá chi tiết như activity diagram.

### Checklist

- [ ] State là trạng thái, không phải action.
- [ ] Transition có trigger/guard/effect khi cần; completion transition không bị gắn trigger giả.
- [ ] Có trạng thái lỗi quan trọng.
- [ ] Vòng đời khớp use case/sequence.

## 10. ERD / Database Schema Diagram

### Mục đích đúng

ERD trả lời câu hỏi: **Dữ liệu được lưu thành các thực thể/bảng nào và quan hệ ra sao?**

### Tiêu chí đạt

- Entity/table có tên rõ.
- Mức conceptual, logical hoặc physical được ghi rõ.
- Logical/physical model có identifier; physical schema có primary key và foreign key quan trọng.
- Quan hệ cardinality rõ.
- Các bảng tài chính có ràng buộc nhất quán.
- Không mâu thuẫn với Entity class ở analysis/design.

### Lỗi thường gặp

- Thiếu bảng trung gian cho quan hệ nhiều-nhiều.
- Thiếu trạng thái hoặc timestamp cần truy vết.
- Thiếu ràng buộc uniqueness cho API key/idempotency.
- Dùng kiểu tiền tệ không chính xác.

### Checklist

- [ ] Mức mô hình được ghi rõ.
- [ ] Identifier hoặc PK/FK phù hợp với mức mô hình.
- [ ] Cardinality rõ.
- [ ] Có trường audit/tracing quan trọng.
- [ ] Khớp entity/design class.
- [ ] Có ràng buộc cho dữ liệu tài chính và idempotency.

## 11. Component Diagram

### Mục đích đúng

Component Diagram trả lời câu hỏi: **Các component mô-đun, interface và dependency của hệ thống là gì?**

### Tiêu chí đạt

- Component đại diện cho đơn vị mô-đun tự chứa, có interface rõ và có thể thay thế trong môi trường của nó.
- Có external systems.
- Interface hoặc protocol giao tiếp rõ.
- Không vẽ quá nhiều class nhỏ.
- Không biến các package nội bộ của modular monolith thành microservice.
- Ghi rõ đây là góc nhìn as-designed hay as-built.
- Nếu kết hợp black-box và white-box, bố cục vẫn phải phân biệt rõ hai mức.

### Lỗi thường gặp

- Component diagram bị trộn với package diagram.
- Thiếu external provider.
- Không rõ hướng giao tiếp.

### Checklist

- [ ] Component có ranh giới và interface rõ.
- [ ] External system rõ.
- [ ] Giao tiếp/protocol rõ.
- [ ] Khớp kiến trúc triển khai.

## 12. Deployment Diagram

### Mục đích đúng

Deployment Diagram trả lời câu hỏi: **Hệ thống được triển khai trên node/môi trường nào?**

### Tiêu chí đạt

- Có node runtime cần thiết cho phạm vi đang mô tả; không bắt buộc mọi loại node xuất hiện.
- Có artifact hoặc service chạy trên node.
- Có kết nối giữa node.
- Phù hợp với môi trường triển khai thực tế hoặc nguyên mẫu.
- Ghi rõ đây là as-designed ở Chương 4 hay as-built ở Chương 5.

### Lỗi thường gặp

- Vẽ deployment như component diagram.
- Không thể hiện database/external provider.
- Vẽ cloud production phức tạp trong khi MVP chạy local.

### Checklist

- [ ] Node triển khai rõ.
- [ ] Artifact/service rõ.
- [ ] Kết nối giữa node rõ.
- [ ] Không phóng đại quá phạm vi MVP.

## 13. Minh chứng triển khai và kiểm thử

### Mục đích đúng

Nhóm artifact này trả lời câu hỏi: **Hệ thống đã triển khai được gì, được kiểm thử ra sao và kết quả đo có đủ bằng chứng không?**

Đây không phải một loại UML diagram.

### Tiêu chí đạt

- Có traceability từ requirement/use case tới module và test.
- Test case có expected result, actual result và trạng thái.
- Biểu đồ hiệu năng ghi workload, đơn vị, môi trường, cỡ mẫu và phương pháp đo.
- Kết quả UC01 phân biệt Gateway overhead với AI Provider latency khi có thể.
- Chỉ trình bày chức năng và môi trường thực tế đã triển khai.

### Lỗi thường gặp

- Gọi bảng test hoặc biểu đồ cột là UML diagram.
- Dùng kiến trúc dự kiến như bằng chứng as-built.
- Chỉ ghi pass/fail mà thiếu điều kiện kiểm thử.
- So sánh hiệu năng nhưng không ghi đơn vị, môi trường hoặc cỡ mẫu.

### Checklist

- [ ] Artifact được gọi đúng loại.
- [ ] Trace được về FR/NFR hoặc use case.
- [ ] Expected và actual result rõ.
- [ ] Phương pháp đo có thể tái hiện.
- [ ] Không phóng đại phần chưa triển khai.
