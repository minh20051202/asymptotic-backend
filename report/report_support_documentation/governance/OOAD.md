# Quy trình phân tích và thiết kế hướng đối tượng OOAD

OOAD là viết tắt của **Object-Oriented Analysis and Design**, nghĩa là **phân tích và thiết kế hướng đối tượng**. Mục tiêu của OOAD là chuyển một bài toán thực tế thành mô hình phần mềm có cấu trúc, trong đó hệ thống được nhìn qua các đối tượng, trách nhiệm, quan hệ và tương tác giữa chúng.

OOAD thường được chia thành hai phần lớn:

- **OOA -- Object-Oriented Analysis:** phân tích hệ thống cần làm gì, ai sử dụng hệ thống, các khái niệm nghiệp vụ chính là gì.
- **OOD -- Object-Oriented Design:** thiết kế hệ thống sẽ được tổ chức như thế nào, các lớp/gói/thành phần tương tác ra sao để hiện thực các yêu cầu.

Nói ngắn gọn: **OOA tập trung vào bài toán**, còn **OOD tập trung vào giải pháp phần mềm**.

## 1. Xác định bài toán và phạm vi hệ thống

### Mục tiêu

Bước đầu tiên là hiểu hệ thống cần giải quyết vấn đề gì, phục vụ ai và giới hạn phạm vi đến đâu. Nếu phạm vi không rõ, các bước sau dễ bị lan man hoặc thiết kế sai đối tượng.

### Câu hỏi cần trả lời

- Hệ thống được xây dựng để giải quyết vấn đề gì?
- Ai là người hoặc hệ thống bên ngoài tương tác với hệ thống?
- Hệ thống chịu trách nhiệm làm gì?
- Những gì nằm ngoài phạm vi hệ thống?
- Kết quả mong muốn của hệ thống là gì?

### Tạo tác thường có

- Mô tả bài toán.
- Mục tiêu hệ thống.
- Phạm vi hệ thống.
- Danh sách actor sơ bộ.
- Danh sách chức năng sơ bộ.

### Vì sao cần bước này?

OOAD không bắt đầu bằng class diagram. Trước khi xác định lớp, cần biết hệ thống đang giải quyết bài toán nào. Nếu bỏ qua bước xác định phạm vi, class diagram thường biến thành danh sách bảng dữ liệu hoặc danh sách class kỹ thuật, không phản ánh đúng nghiệp vụ.

## 2. Khảo sát hiện trạng và bối cảnh nghiệp vụ

### Mục tiêu

Bước này giúp hiểu môi trường mà hệ thống sẽ hoạt động: quy trình hiện tại, hệ thống bên ngoài, dữ liệu đang tồn tại, ràng buộc nghiệp vụ và các vấn đề của cách làm cũ.

### Câu hỏi cần trả lời

- Hiện tại người dùng hoặc tổ chức đang xử lý bài toán này như thế nào?
- Có hệ thống bên ngoài nào liên quan?
- Quy trình hiện tại gặp vấn đề gì?
- Có ràng buộc pháp lý, bảo mật, tài chính, hiệu năng hoặc vận hành nào không?
- Dữ liệu nào cần được ghi nhận, xử lý hoặc trao đổi?

### Tạo tác thường có

- Mô tả hiện trạng.
- Mô tả hệ thống bên ngoài.
- Vấn đề/khoảng trống cần giải quyết.
- Ràng buộc nghiệp vụ và kỹ thuật.

### Vì sao cần bước này?

OOAD không chỉ mô hình hóa phần mềm, mà còn mô hình hóa nghiệp vụ. Khảo sát hiện trạng giúp tránh thiết kế hệ thống theo giả định chủ quan.

## 3. Xác định actor và use case

### Mục tiêu

Actor là người dùng hoặc hệ thống bên ngoài tương tác với hệ thống. Use case mô tả mục tiêu mà actor muốn đạt được thông qua hệ thống.

### Câu hỏi cần trả lời

- Ai sử dụng hệ thống?
- Hệ thống bên ngoài nào tương tác với hệ thống?
- Mỗi actor muốn thực hiện mục tiêu gì?
- Use case nào là chính, use case nào là phụ?
- Use case nào thuộc phạm vi hiện tại, use case nào nằm ngoài phạm vi?

### Tạo tác thường có

- Danh sách actor.
- Danh sách use case.
- Use Case Diagram tổng quát.
- Use Case Diagram phân rã theo nhóm chức năng nếu hệ thống lớn.

### Vì sao vẽ Use Case Diagram ở bước này?

Use Case Diagram giúp xác định **hệ thống làm gì từ góc nhìn bên ngoài**. Nó chưa quan tâm bên trong hệ thống có class nào, database nào hay công nghệ nào. Đây là bước cầu nối giữa mô tả bài toán và phân tích chi tiết.

## 4. Đặc tả use case chi tiết

### Mục tiêu

Sau khi xác định use case, cần mô tả chi tiết từng use case quan trọng. Đặc tả use case giúp thống nhất hành vi hệ thống trước khi đi vào class hoặc sequence diagram.

### Nội dung thường có

- Mã use case.
- Tên use case.
- Actor chính.
- Actor phụ.
- Mục tiêu.
- Tiền điều kiện.
- Hậu điều kiện.
- Luồng chính.
- Luồng thay thế.
- Luồng ngoại lệ.
- Dữ liệu vào/ra.
- Quy tắc nghiệp vụ liên quan.

### Vì sao cần đặc tả use case?

Use Case Diagram chỉ cho biết actor có thể làm gì, nhưng không mô tả chi tiết quá trình thực hiện. Đặc tả use case cung cấp nội dung để vẽ activity diagram, robustness diagram, sequence diagram và xác định lớp phân tích.

## 5. Phân tích luồng xử lý bằng Activity Diagram

### Mục tiêu

Activity Diagram mô tả trình tự hoạt động, điều kiện rẽ nhánh, vòng lặp và điểm kết thúc của một quy trình hoặc use case.

### Khi nào nên vẽ?

- Khi use case có nhiều bước xử lý.
- Khi có nhiều nhánh thành công/thất bại.
- Khi cần mô tả quy trình nghiệp vụ.
- Khi cần làm rõ điều kiện trước khi hệ thống thực hiện hành động quan trọng.

### Tạo tác thường có

- Activity Diagram cho use case chính.
- Activity Diagram cho các quy trình nghiệp vụ quan trọng.

### Vì sao Activity Diagram thường vẽ trước Class Diagram chi tiết?

Activity Diagram giúp hiểu **dòng chảy nghiệp vụ**. Từ dòng chảy đó có thể xác định hệ thống cần những trách nhiệm nào, từ đó mới phân bổ trách nhiệm cho các lớp phù hợp. Nếu vẽ class trước khi hiểu luồng nghiệp vụ, class dễ bị thiếu hành vi hoặc bị tách sai trách nhiệm.

## 6. Xác định yêu cầu chức năng và phi chức năng

### Mục tiêu

Yêu cầu chức năng mô tả hệ thống phải làm gì. Yêu cầu phi chức năng mô tả hệ thống phải đạt chất lượng gì.

### Yêu cầu chức năng

Ví dụ:

- Hệ thống cho phép người dùng đăng nhập.
- Hệ thống cho phép tạo đơn hàng.
- Hệ thống gửi thông báo khi giao dịch hoàn tất.

### Yêu cầu phi chức năng

Ví dụ:

- Bảo mật.
- Hiệu năng.
- Tính sẵn sàng.
- Khả năng mở rộng.
- Khả năng truy vết.
- Tính nhất quán dữ liệu.

### Tạo tác thường có

- Danh sách Functional Requirements.
- Danh sách Non-functional Requirements.
- Bảng liên kết yêu cầu với use case.

### Vì sao cần bước này?

Use case mô tả tương tác, nhưng yêu cầu hệ thống cần rõ ràng hơn để thiết kế và kiểm thử. FR/NFR là cơ sở để đánh giá hệ thống có đáp ứng bài toán hay không.

## 7. Nhận diện lớp phân tích

### Mục tiêu

Lớp phân tích là các lớp ở mức nghiệp vụ, dùng để hiểu hệ thống cần những đối tượng và trách nhiệm nào. Đây chưa phải class code cuối cùng.

### Cách nhận diện lớp

Một cách phổ biến là phân loại lớp theo mô hình **Boundary-Control-Entity**:

- **Boundary class:** lớp giao tiếp giữa actor/hệ thống ngoài và hệ thống.
- **Control class:** lớp điều phối một use case hoặc một luồng nghiệp vụ.
- **Entity class:** lớp biểu diễn khái niệm nghiệp vụ có dữ liệu và vòng đời riêng.

### Nguồn để tìm lớp

- Danh từ trong mô tả bài toán.
- Actor và use case.
- Dữ liệu cần lưu trữ.
- Quy tắc nghiệp vụ.
- Các bước trong activity diagram.

### Tạo tác thường có

- Danh sách lớp Boundary.
- Danh sách lớp Control.
- Danh sách lớp Entity.
- Analysis Class Diagram.

### Vì sao dùng Boundary-Control-Entity?

Boundary-Control-Entity giúp tách trách nhiệm rõ ràng:

- Boundary không chứa nghiệp vụ phức tạp.
- Control điều phối luồng xử lý.
- Entity giữ trạng thái và dữ liệu nghiệp vụ.

Cách phân tách này làm cho mô hình phân tích dễ hiểu và dễ chuyển sang thiết kế.

## 8. Vẽ Robustness Diagram

### Mục tiêu

Robustness Diagram là cầu nối giữa use case và class diagram. Nó cho thấy actor tương tác với boundary nào, boundary gọi control nào, control sử dụng entity nào.

### Thành phần

- Actor.
- Boundary object.
- Control object.
- Entity object.

### Khi nào nên vẽ?

- Khi muốn kiểm tra use case đã đủ lớp phân tích chưa.
- Khi muốn tránh nhảy trực tiếp từ use case sang sequence diagram.
- Khi cần chứng minh phân tích hướng đối tượng rõ hơn.

### Vì sao Robustness Diagram hữu ích?

Nó giúp phát hiện lỗi phân tích sớm:

- Use case có bước nhưng chưa có control xử lý.
- Entity bị actor truy cập trực tiếp.
- Boundary chứa quá nhiều logic.
- Thiếu entity để lưu trạng thái nghiệp vụ.

## 9. Xây dựng Analysis Class Diagram

### Mục tiêu

Analysis Class Diagram mô tả các lớp phân tích và quan hệ nghiệp vụ giữa chúng. Biểu đồ này tập trung vào khái niệm và trách nhiệm, chưa cần đi sâu vào framework, database hoặc implementation detail.

### Nội dung thường có

- Tên lớp.
- Thuộc tính nghiệp vụ chính.
- Quan hệ association, aggregation, composition, generalization nếu cần.
- Phân loại Boundary-Control-Entity.

### Vì sao chưa nên đưa chi tiết kỹ thuật vào đây?

Analysis Class Diagram dùng để hiểu bài toán. Nếu đưa sớm repository, framework, DTO, controller hoặc database table vào giai đoạn này, mô hình sẽ bị lệch sang thiết kế/triển khai và mất vai trò phân tích.

## 10. Chuyển từ phân tích sang thiết kế

### Mục tiêu

Sau khi có mô hình phân tích, bước tiếp theo là quyết định hệ thống sẽ được tổ chức như thế nào trong phần mềm thực tế.

### Câu hỏi cần trả lời

- Hệ thống chia thành những module/gói nào?
- Lớp phân tích nào trở thành lớp thiết kế?
- Cần thêm service, repository, adapter, controller không?
- Các lớp phụ thuộc vào nhau như thế nào?
- Có pattern kiến trúc nào phù hợp?
- Dữ liệu được lưu ở đâu?
- API hoặc giao diện hệ thống được thiết kế như thế nào?

### Tạo tác thường có

- Package Diagram.
- Design Class Diagram.
- Component Diagram.
- ERD hoặc database schema.
- API specification.

### Vì sao cần bước chuyển đổi?

Lớp phân tích không nhất thiết trùng một-một với class trong code. Thiết kế cần bổ sung các yếu tố kỹ thuật để hệ thống có thể triển khai, kiểm thử và vận hành.

## 11. Thiết kế gói bằng Package Diagram

### Mục tiêu

Package Diagram mô tả cách chia hệ thống thành các gói hoặc module lớn.

### Nội dung thường có

- Tên package/module.
- Quan hệ phụ thuộc giữa package.
- Ranh giới giữa các phần của hệ thống.

### Vì sao Package Diagram quan trọng?

Nó giúp kiểm soát độ phức tạp của hệ thống. Với hệ thống lớn, chỉ nhìn class diagram tổng quát có thể rất rối. Package Diagram cho thấy kiến trúc module ở mức cao trước khi đi vào chi tiết class.

## 12. Thiết kế lớp bằng Design Class Diagram

### Mục tiêu

Design Class Diagram mô tả các lớp ở mức thiết kế, gần với triển khai hơn Analysis Class Diagram.

### Nội dung thường có

- Class thiết kế.
- Interface.
- Service.
- Repository.
- Adapter.
- Thuộc tính và phương thức quan trọng.
- Quan hệ phụ thuộc.

### Khác gì với Analysis Class Diagram?

| Analysis Class Diagram | Design Class Diagram |
|---|---|
| Tập trung vào khái niệm nghiệp vụ | Tập trung vào cấu trúc phần mềm |
| Ít chi tiết kỹ thuật | Có service, repository, adapter, interface |
| Dùng để hiểu bài toán | Dùng để hướng dẫn triển khai |
| Không phụ thuộc framework | Có thể phản ánh kiến trúc triển khai |

## 13. Thiết kế tương tác bằng Sequence Diagram

### Mục tiêu

Sequence Diagram mô tả các đối tượng tương tác theo thời gian để thực hiện một use case.

### Khi nào nên vẽ?

- Cho các use case quan trọng.
- Cho luồng có nhiều service hoặc hệ thống ngoài.
- Cho luồng cần kiểm tra thứ tự xử lý.
- Cho luồng có transaction, điều kiện lỗi hoặc gọi API ngoài.

### Nội dung thường có

- Actor.
- Boundary/controller.
- Service/control.
- Entity/repository.
- External system.
- Message call/return.
- Nhánh điều kiện nếu cần.

### Vì sao Sequence Diagram thường nằm ở giai đoạn thiết kế?

Ở giai đoạn phân tích, ta cần biết hệ thống làm gì. Ở giai đoạn thiết kế, ta cần biết các đối tượng phối hợp ra sao để làm điều đó. Sequence Diagram giúp kiểm tra thiết kế class có đủ phương thức và trách nhiệm chưa.

## 14. Thiết kế trạng thái bằng State Machine Diagram

### Mục tiêu

State Machine Diagram mô tả vòng đời của một đối tượng có trạng thái quan trọng.

### Khi nào nên vẽ?

- Đối tượng có nhiều trạng thái.
- Trạng thái ảnh hưởng đến hành vi hệ thống.
- Có chuyển trạng thái do sự kiện.
- Có trạng thái lỗi, hủy, hoàn tất hoặc rollback.

### Ví dụ đối tượng nên có State Diagram

- Order.
- Payment.
- Request.
- Transaction.
- Account.
- Document approval.
- Ticket.

### Vì sao cần State Machine Diagram?

Class Diagram cho biết đối tượng có thuộc tính gì, nhưng không mô tả vòng đời trạng thái. State Machine Diagram giúp tránh lỗi như chuyển trạng thái không hợp lệ, xử lý thiếu trạng thái lỗi hoặc không rõ điều kiện hoàn tất.

## 15. Thiết kế dữ liệu

### Mục tiêu

Thiết kế dữ liệu xác định cách lưu trữ các thực thể và quan hệ trong hệ quản trị cơ sở dữ liệu hoặc kho dữ liệu khác.

### Tạo tác thường có

- ERD.
- Database schema.
- Bảng dữ liệu.
- Khóa chính, khóa ngoại.
- Ràng buộc toàn vẹn.
- Index quan trọng.

### Quan hệ với OOAD

ERD không phải UML OOAD thuần, nhưng thường cần trong đồ án phần mềm vì hệ thống thực tế phải lưu dữ liệu. ERD nên xuất phát từ Entity class và các yêu cầu lưu trữ.

## 16. Thiết kế thành phần và triển khai

### Mục tiêu

Bước này mô tả hệ thống ở mức thành phần triển khai: service, database, external system, message broker, client, server.

### Tạo tác thường có

- Component Diagram.
- Deployment Diagram.
- Mô tả môi trường triển khai.
- Mô tả công nghệ sử dụng.

### Vì sao bước này nằm sau thiết kế hướng đối tượng?

Component và deployment phản ánh cách hệ thống được đóng gói và chạy trong môi trường thật. Chúng phụ thuộc vào quyết định thiết kế trước đó, nên thường không nên đặt trước analysis class hoặc design class.

## 17. Kiểm thử và truy vết yêu cầu

### Mục tiêu

Kiểm thử xác nhận hệ thống đã đáp ứng yêu cầu. Truy vết yêu cầu giúp chứng minh mỗi yêu cầu đã được phân tích, thiết kế và kiểm tra.

### Tạo tác thường có

- Test case.
- Test scenario.
- Test result.
- Traceability Matrix.

### Traceability Matrix là gì?

Traceability Matrix là bảng liên kết:

- Requirement.
- Use case.
- Class/module thiết kế.
- Test case.
- Kết quả kiểm thử.

### Vì sao quan trọng?

Nó giúp chứng minh báo cáo không chỉ có biểu đồ rời rạc. Mỗi yêu cầu đều có đường đi từ phân tích đến thiết kế và kiểm thử.

## 18. Thứ tự diagram trong quy trình OOAD

Một thứ tự phổ biến là:

1. Use Case Diagram.
2. Use Case Specification.
3. Activity Diagram.
4. Robustness Diagram.
5. Analysis Class Diagram.
6. Package Diagram.
7. Design Class Diagram.
8. Sequence Diagram.
9. State Machine Diagram.
10. ERD hoặc Database Schema.
11. Component Diagram.
12. Deployment Diagram.
13. Test Case và Traceability Matrix.

Thứ tự này đi từ ngoài vào trong:

- Từ actor và mục tiêu.
- Đến luồng nghiệp vụ.
- Đến lớp phân tích.
- Đến lớp thiết kế.
- Đến tương tác runtime.
- Đến trạng thái, dữ liệu, triển khai và kiểm thử.

## 19. Lỗi thường gặp khi làm OOAD

### Nhảy thẳng vào database

Nếu bắt đầu bằng bảng dữ liệu, mô hình dễ trở thành thiết kế database, không phải phân tích hướng đối tượng.

### Nhảy thẳng vào class code

Nếu class diagram chứa quá nhiều controller, repository, DTO ngay từ đầu, đó là thiết kế kỹ thuật, không phải analysis model.

### Use case quá nhỏ hoặc quá kỹ thuật

Use case nên mô tả mục tiêu của actor, không nên là từng API endpoint nhỏ.

### Không phân biệt analysis class và design class

Analysis class dùng để hiểu nghiệp vụ. Design class dùng để triển khai phần mềm. Hai loại này liên quan nhưng không giống nhau.

### Thiếu traceability

Nếu yêu cầu, use case, class và test không liên kết với nhau, báo cáo sẽ giống tập hợp biểu đồ rời rạc.

## 20. Tóm tắt quy trình

Quy trình OOAD có thể tóm tắt như sau:

1. Hiểu bài toán và phạm vi.
2. Khảo sát hiện trạng và bối cảnh.
3. Xác định actor và use case.
4. Đặc tả use case.
5. Mô hình hóa luồng xử lý.
6. Xác định FR/NFR.
7. Nhận diện lớp phân tích.
8. Vẽ robustness diagram nếu cần.
9. Xây dựng analysis class diagram.
10. Chuyển sang thiết kế gói, lớp, tương tác, trạng thái.
11. Thiết kế dữ liệu và thành phần triển khai.
12. Kiểm thử và truy vết yêu cầu.

Điểm cốt lõi của OOAD là **không bắt đầu từ code**, mà bắt đầu từ **actor, mục tiêu, nghiệp vụ và trách nhiệm của đối tượng**. Các biểu đồ được vẽ theo thứ tự đó để bảo đảm thiết kế phần mềm có căn cứ từ yêu cầu, thay vì chỉ là mô tả lại implementation.
