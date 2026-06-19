# Quy tắc vẽ UML 2.5.1 cho đồ án

Tài liệu này tóm lược các quy tắc ký pháp cần áp dụng khi vẽ biểu đồ cho đồ án, dựa trên:

- **Object Management Group, Unified Modeling Language, Version 2.5.1, formal/2017-12-05**;
- [`project_source_of_truth.md`](../../governance/project_source_of_truth.md);
- quy trình OOAD và quy ước trình bày của báo cáo.

Tệp chuẩn gốc: [`formal-17-12-05.pdf`](../../references/standards/formal-17-12-05.pdf).

Ảnh ví dụ đã trích: [`uml_2_5_1_examples/README.md`](../../references/examples/uml_2_5_1_examples/README.md).

## 1. Phạm vi áp dụng

UML 2.5.1 quy định cú pháp và ngữ nghĩa của ký pháp. Chuẩn không quy định chi tiết cách căn lề, màu sắc, kích thước chữ hoặc cách tối ưu đường nối cho một báo cáo cụ thể.

Tài liệu phân biệt ba lớp quy tắc:

- **Quy tắc UML:** lấy từ OMG UML 2.5.1.
- **Quy tắc phương pháp:** lấy từ OOAD, OOSE/ICONIX hoặc kỹ thuật use-case-driven analysis.
- **Quy ước dự án:** mức chi tiết, bố cục, traceability và cách biểu diễn kiến trúc Asymptotic.

Analysis Class Diagram và Design Class Diagram là hai cách dùng UML Class Diagram ở các mức khác nhau. BCE và Robustness là kỹ thuật phương pháp; ERD và biểu đồ kết quả kiểm thử không thuộc UML.

Khi có xung đột, ưu tiên theo thứ tự:

1. [`project_source_of_truth.md`](../../governance/project_source_of_truth.md) về nội dung hệ thống;
2. UML 2.5.1 về ý nghĩa và hướng của ký pháp;
3. `diagram_type_rubric.md` về mục đích OOAD;
4. `diagram_quality_guidelines.md` về khả năng đọc;
5. lựa chọn bố cục của PlantUML.

Không được sửa sai ngữ nghĩa UML chỉ để biểu đồ đẹp hơn.

## 2. Quy tắc chung

- Mỗi biểu đồ phải có một mục đích và một mức trừu tượng rõ ràng.
- Chỉ dùng ký pháp có ý nghĩa xác định; không dùng một kiểu mũi tên cho nhiều quan hệ khác nhau.
- Tên phần tử phải nhất quán với source of truth và phần văn bản tương ứng.
- Quan hệ phải có đúng hướng, đúng đầu mũi tên và đúng stereotype.
- Không dùng màu để thay thế ý nghĩa của ký pháp.
- Hạn chế đường cắt nhau; ưu tiên nhóm phần tử liên quan và sắp xếp theo hướng đọc.
- Chỉ đưa chi tiết đủ để trả lời câu hỏi của biểu đồ.
- File `.puml` phải được lưu cạnh ảnh xuất.
- Ảnh đưa vào báo cáo phải đọc được ở kích thước in thực tế.

## 3. Use Case Diagram

Nguồn chuẩn: Điều 18, đặc biệt 18.1 và 18.2.

### Phần tử và vị trí

- Actor biểu diễn **vai trò** tương tác với hệ thống, không nhất thiết là một người cụ thể.
- Actor có thể là người, phần cứng hoặc hệ thống bên ngoài.
- Actor đặt ngoài system boundary.
- Use case đặt trong system boundary và biểu diễn bằng hình ellipse.
- Tên use case mô tả mục tiêu tạo ra kết quả quan sát được, thường dùng động từ và bổ ngữ.
- System boundary là hình chữ nhật có tên chủ thể/hệ thống.

### Quan hệ

- Association giữa actor và use case là đường liền.
- `<<include>>` là đường đứt, mũi tên mở từ use case cơ sở đến use case được bao gồm.
- Dùng `include` khi hành vi được bao gồm là bắt buộc và được tái sử dụng.
- `<<extend>>` là đường đứt, mũi tên mở từ use case mở rộng đến use case cơ sở.
- Dùng `extend` khi hành vi bổ sung là tùy chọn hoặc có điều kiện tại extension point.
- Generalization hướng từ phần tử chuyên biệt đến phần tử tổng quát bằng tam giác rỗng.
- Không nối hai use case bằng association thông thường.
- Không dùng `include` để biểu diễn thứ tự trước/sau hoặc điều kiện tiên quyết.

### Áp dụng cho đồ án

- Sơ đồ tổng quát dùng đúng chín use case chuẩn.
- AI Agent là actor ngoài hệ thống.
- AI Provider và Payment Provider là supporting actors bên ngoài.
- Các bước xác thực khóa, kiểm tra budget và ghi trace thuộc luồng UC01, không tự động trở thành use case độc lập.
- Không cần `include` hoặc `extend` nếu chín use case đều là mục tiêu độc lập.

## 4. Activity Diagram

Nguồn chuẩn: Điều 15.

- Initial node là chấm tròn đặc và không có cạnh đi vào.
- Activity final là vòng tròn kép, kết thúc toàn bộ activity.
- Flow final là vòng tròn có dấu X, chỉ kết thúc luồng đi vào nó.
- Action đặt tên bằng động từ hoặc cụm động từ.
- Control flow là đường có mũi tên mở.
- Decision node là hình thoi, thường có một cạnh vào và nhiều cạnh ra.
- Mỗi nhánh ra từ decision phải có guard rõ; có thể dùng `[else]`.
- Merge node gộp các nhánh thay thế, không đồng bộ hóa.
- Fork có một cạnh vào và nhiều cạnh ra, dùng cho luồng song song.
- Join có nhiều cạnh vào và một cạnh ra, dùng để đồng bộ.
- Không dùng merge thay join hoặc decision thay fork.
- Activity partition/swimlane biểu diễn bên chịu trách nhiệm; nó không thay đổi ngữ nghĩa token flow.
- Object flow chỉ dùng khi cần thể hiện dữ liệu hoặc đối tượng được truyền giữa action.

### Áp dụng cho đồ án

- Chương 2 ưu tiên luồng nghiệp vụ, không đưa tên controller, repository hoặc bảng dữ liệu.
- UC01 cần thể hiện nhánh từ chối do khóa, policy, quota hoặc budget.
- Chỉ dùng fork/join khi các công việc thật sự có thể diễn ra song song.
- Swimlane nên phản ánh actor, Gateway và external provider, không phản ánh từng module thiết kế.

## 5. Class Diagram

Nguồn chuẩn: Điều 9 và Điều 11.

- Class dùng hình chữ nhật, có thể chia các ngăn tên, thuộc tính và operation.
- Không bắt buộc hiển thị mọi thuộc tính và operation.
- Visibility dùng `+`, `-`, `#`, `~` khi cần.
- Association là đường liền; role name và multiplicity đặt gần đầu liên kết tương ứng.
- Multiplicity phải phản ánh đúng số lượng đối tượng tham gia quan hệ.
- Navigability dùng mũi tên mở tại đầu có thể điều hướng; không tự suy diễn khi không vẽ.
- Shared aggregation dùng hình thoi rỗng tại phía whole.
- Composite aggregation dùng hình thoi đặc tại phía whole và chỉ dùng khi có quan hệ sở hữu vòng đời mạnh.
- Generalization là đường liền với tam giác rỗng hướng về classifier tổng quát.
- Realization là đường đứt với tam giác rỗng hướng về interface/specification.
- Dependency là đường đứt với mũi tên mở từ client đến supplier.

### Áp dụng cho đồ án

- Analysis class diagram chỉ giữ khái niệm nghiệp vụ và lớp Boundary-Control-Entity.
- Design class diagram mới đưa service, repository, adapter và interface.
- Không dùng composition chỉ vì khóa ngoại tồn tại.
- Các multiplicity quan trọng như Organization-Team-Developer-Agent phải thống nhất với use case và mô hình dữ liệu.

### Quan hệ trong Design Class Diagram

| Quan hệ | Dùng khi | PlantUML |
|---|---|---|
| Association | Có liên kết cấu trúc hoặc tham chiếu lâu dài | `--`, `-->` |
| Dependency | Một lớp sử dụng lớp khác tạm thời | `..>` |
| Generalization | Lớp chuyên biệt kế thừa lớp tổng quát | `--|>` |
| Realization | Lớp hiện thực interface/contract | `..|>` |
| Shared aggregation | Whole-part nhưng part tồn tại độc lập | `o--` |
| Composition | Whole sở hữu độc quyền và chi phối vòng đời part | `*--` |

Shared aggregation không bắt buộc cho mọi quan hệ whole-part; association thường rõ hơn. Composition chỉ dùng khi có ngữ nghĩa sở hữu vòng đời, không suy ra từ foreign key hoặc quan hệ `has-a`.

## 6. Sequence Diagram

Nguồn chuẩn: Điều 17, đặc biệt 17.2, 17.3, 17.4, 17.6 và 17.8.

- Lifeline có phần đầu hình chữ nhật và đường dọc; thời gian tăng từ trên xuống.
- Khoảng cách dọc không biểu diễn thời lượng tuyệt đối nếu không có time constraint.
- Message phải đi ngang hoặc đi xuống từ send event đến receive event, không đi ngược lên theo thời gian.
- Synchronous call dùng mũi tên đặc.
- Asynchronous message dùng mũi tên mở.
- Reply, nếu được biểu diễn, thường dùng đường đứt và mũi tên mở. UML không bắt buộc vẽ mọi reply.
- Create message kết thúc tại đầu lifeline được tạo.
- Destruction occurrence dùng dấu X ở cuối lifeline.
- Activation/execution specification là hình chữ nhật mảnh trên lifeline.
- Message name nên phản ánh operation, signal hoặc kết quả có nghĩa.

### Combined fragment

- `alt`: chọn tối đa một nhánh theo guard; có thể có `[else]`.
- `opt`: thực hiện một nhánh khi guard đúng hoặc bỏ qua.
- `loop`: lặp theo guard hoặc giới hạn.
- `par`: các operand có thể xen kẽ/song song.
- `break`: khi guard đúng, thay thế phần còn lại của interaction bao quanh.
- `critical`: vùng không được xen kẽ bởi event khác.

### Áp dụng cho đồ án

- Actor không gọi trực tiếp database.
- Sequence UC01 phải thể hiện Gateway gọi AI Provider bằng credential nội bộ.
- Nhánh từ chối trước provider call phải nằm trước message gọi provider.
- Modular monolith dùng lifeline theo module/trách nhiệm trong cùng ứng dụng; không vẽ mỗi module như một microservice độc lập.
- Chỉ thể hiện persistence chung hoặc repository phù hợp, không tạo một database cho mỗi module.

## 7. State Machine Diagram

Nguồn chuẩn: Điều 14.

- State là hình chữ nhật bo góc và mô tả một điều kiện ổn định của đối tượng.
- Tên state nên mô tả một điều kiện ổn định có ý nghĩa. Dùng danh từ, tính từ hoặc trạng thái kết quả là heuristic trình bày, không phải cú pháp UML bắt buộc.
- Initial pseudostate là chấm tròn đặc, có tối đa một transition đi ra.
- Transition từ initial pseudostate không có trigger hoặc guard.
- Final state là vòng tròn kép và không có transition đi ra.
- Transition label có dạng tổng quát `trigger [guard] / effect`; mỗi thành phần có thể được lược bỏ.
- Transition không có trigger có thể là completion transition.
- Choice dùng hình thoi để chọn nhánh theo guard.
- Fork/join dùng thanh đậm cho vùng trạng thái song song.
- Composite state có thể chứa region; các orthogonal region được phân cách bằng đường đứt.
- Entry, exit và do behavior đặt trong compartment của state khi cần.

### Áp dụng cho đồ án

- Không biến các bước `Kiểm tra ngân sách`, `Gọi provider` thành state nếu chúng chỉ là action tức thời.
- State của AI Request phải phản ánh trạng thái lưu được hoặc trạng thái nghiệp vụ có ý nghĩa.
- Transition phải khớp với sequence và quy tắc lỗi của use case.

## 8. Package Diagram

Nguồn chuẩn: Điều 12.

- Package dùng ký hiệu thư mục có tab.
- Package là namespace/đơn vị tổ chức mô hình, không tự động là đơn vị triển khai.
- Dependency giữa package phải có hướng từ client đến supplier.
- Package import, access hoặc merge phải dùng đúng stereotype nếu cần biểu diễn các ngữ nghĩa này.
- Package merge là đường đứt, mũi tên mở từ receiving package đến merged package, nhãn `<<merge>>`.

### Áp dụng cho đồ án

- Package diagram Chương 4 biểu diễn module của modular monolith.
- Không vẽ một database riêng cho từng package.
- Đặt package theo domain/responsibility và tránh dependency vòng là quy ước kiến trúc của dự án, không phải cú pháp UML.
- Dependency phải theo trách nhiệm nghiệp vụ và kiến trúc đã chốt.

## 9. Component Diagram

Nguồn chuẩn: Điều 11.6.

- Component là đơn vị mô-đun, tự chứa và có thể thay thế trong môi trường của nó.
- Component dùng hình chữ nhật có `<<component>>` hoặc biểu tượng component.
- Provided interface dùng ký hiệu lollipop.
- Required interface dùng ký hiệu socket.
- Dependency có thể dùng ở mức tổng quan khi chưa cần biểu diễn interface chi tiết.
- Black-box interface view và white-box internal structure view đều hợp lệ. Nên tách thành hai biểu đồ nếu việc trộn hai mức làm giảm khả năng đọc.

### Áp dụng cho đồ án

- Modular monolith có thể là một component ứng dụng lớn, bên trong được phân rã bằng package/module.
- Không gọi mỗi module nội bộ là một deployable service nếu chúng không triển khai độc lập.
- AI Provider, Payment Provider và database có thể xuất hiện như component/external system liên quan, nhưng phải phân biệt với module nội bộ.

## 10. Deployment Diagram

Nguồn chuẩn: Điều 19.

- Node là computational resource có thể chứa artifact được triển khai.
- Device là node vật lý, dùng `<<device>>`.
- Execution environment là môi trường chạy phần mềm, dùng `<<executionEnvironment>>`.
- Artifact là phần tử vật lý có thể triển khai, dùng `<<artifact>>`.
- Deployment có thể biểu diễn bằng artifact nằm trong node hoặc dependency `<<deploy>>`.
- Communication path giữa các node dùng association.
- Device có thể chứa execution environment.

### Áp dụng cho đồ án

- Phân biệt rõ component logic, artifact triển khai và node runtime.
- Sơ đồ phải phản ánh MVP/modular monolith, không phóng đại thành cụm microservice.
- Database là node/runtime resource hoặc hệ quản trị dữ liệu, không phải package nghiệp vụ.
- External providers nằm ngoài hạ tầng Asymptotic.
- UML không bắt buộc mọi Deployment Diagram phải có client, server, database và external provider. Các phần tử được chọn theo phạm vi câu hỏi của biểu đồ.
- Nếu đặt ở Chương 4, ghi rõ đây là deployment **as-designed**. Nếu đặt ở Chương 5, biểu đồ phải phản ánh deployment **as-built**.

## 11. Robustness Diagram và BCE

Robustness Diagram không phải một loại biểu đồ độc lập được UML 2.5.1 chuẩn hóa. Đây là quy ước phân tích sử dụng UML object/class notation cùng stereotype:

- `<<boundary>>`;
- `<<control>>`;
- `<<entity>>`.

Quy tắc actor-boundary-control-entity là quy tắc của phương pháp robustness/BCE, không được ghi là quy tắc bắt buộc của OMG UML:

- actor giao tiếp với boundary;
- boundary chuyển yêu cầu cho control;
- control điều phối entity;
- actor không truy cập entity trực tiếp.

Trong Analysis Class Diagram, Entity có thể liên hệ trực tiếp với Entity. Không áp dụng quy tắc kết nối Robustness Diagram để cấm association cấu trúc giữa các Entity.

## 12. ERD

ERD không phải loại biểu đồ thuộc UML 2.5.1.

Trong đồ án:

- có thể dùng PlantUML entity notation cho database schema;
- hoặc dùng UML class notation với stereotype phù hợp;
- phải ghi rõ mức conceptual, logical hoặc physical;
- conceptual ERD không bắt buộc hiển thị PK/FK;
- logical ERD cần identifier, cardinality, optionality và associative entity quan trọng;
- physical schema diagram mới cần PK/FK, kiểu dữ liệu, index và constraint triển khai;
- phải ghi rõ đây là ERD/database schema diagram;
- không trộn cardinality ERD với multiplicity UML trong cùng một biểu đồ mà không giải thích;
- ERD không thay thế analysis class diagram hoặc design class diagram.

## 13. Biểu đồ và minh chứng Chương 5

Chương 5 có thể dùng:

- As-built Component Diagram, nếu cần mô tả module/component thực tế đã triển khai.
- As-built Deployment Diagram, nếu có node, artifact và môi trường chạy thực tế.
- Ma trận Requirement--Use Case--Module--Test.
- Bảng test case và kết quả.
- Biểu đồ hiệu năng, tỷ lệ pass/fail hoặc chi phí.

Hai mục đầu là UML. Ma trận, bảng và biểu đồ kết quả là minh chứng báo cáo, không được gọi là UML diagram.

Không bắt buộc tạo UML diagram cho Chương 5 khi chưa có dữ liệu triển khai. Biểu đồ hiệu năng phải ghi workload, đơn vị, môi trường, cỡ mẫu và cách đo. Với UC01, nên tách Gateway overhead khỏi thời gian phản hồi của AI Provider.

## 14. Quy tắc PlantUML

- Dùng PlantUML cho toàn bộ nguồn biểu đồ.
- Chọn cú pháp đúng loại biểu đồ, không dùng generic rectangle thay mọi phần tử UML.
- Dùng alias ngắn, ổn định; label hiển thị dùng thuật ngữ chuẩn.
- Dùng layout hint và hidden link chỉ để cải thiện bố cục, không tạo quan hệ giả.
- Không dựa vào màu để phân biệt primary/supporting actor hoặc loại quan hệ.
- Sau khi render, kiểm tra ảnh bằng mắt: chữ, mũi tên, đường cắt, tỷ lệ và khả năng đọc trong PDF.
- Nếu biểu đồ quá dày, giảm nội dung hoặc tách theo mục đích; không thu nhỏ đến mức không đọc được.

## 15. Kiểm tra trước khi duyệt

- [ ] Nội dung khớp [`project_source_of_truth.md`](../../governance/project_source_of_truth.md).
- [ ] Phân biệt rõ quy tắc UML, quy tắc phương pháp và quy ước dự án.
- [ ] Loại biểu đồ đúng mục đích và đúng chương OOAD.
- [ ] Ký pháp, hướng mũi tên và stereotype đúng UML.
- [ ] Không dùng `include`, `extend`, aggregation hoặc composition sai nghĩa.
- [ ] Không mô hình hóa modular monolith thành microservice.
- [ ] Actor và external system nằm đúng ranh giới.
- [ ] Tên phần tử nhất quán với văn bản.
- [ ] Biểu đồ đọc được ở kích thước báo cáo.
- [ ] `.puml` và ảnh xuất được lưu cùng nhau.
- [ ] Caption và nguồn hình đúng quy tắc báo cáo.
