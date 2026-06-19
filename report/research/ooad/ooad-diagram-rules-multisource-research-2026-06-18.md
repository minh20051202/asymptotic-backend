# Research Brief: Quy tắc và ký pháp cho bộ biểu đồ OOAD

**Ngày nghiên cứu:** 2026-06-18  
**Phạm vi:** Các biểu đồ trong `ooad_diagram_priority_list.md`, gồm Use Case, Activity, Analysis Class/BCE, Robustness, Package, Design Class, Sequence, State Machine; bổ sung ERD, Component và Deployment.  
**Mục tiêu:** Xây dựng hướng dẫn vẽ dựa trên nhiều nguồn, không chỉ dựa vào OMG UML 2.5.1.

Tài liệu này đã hợp nhất và thay thế research brief riêng về Analysis Class, BCE và Robustness ngày 18/06/2026.

**Cập nhật:** Các hiệu chỉnh về phân loại UML/phương pháp, transition không có trigger, reply tùy chọn, ERD theo mức mô hình và ranh giới Chương 4--5 nằm tại `ooad-diagram-rules-chapters-3-5-research-2026-06-19.md`.

## 1. Kết luận điều hành

1. OMG UML cung cấp ngữ nghĩa và ký pháp chuẩn, nhưng không đủ để hướng dẫn lựa chọn mức chi tiết, bố cục và cách phân rã biểu đồ trong báo cáo. Cần kết hợp sách OOAD/ICONIX, Agile Modeling, tài liệu công cụ và nguồn thiết kế dữ liệu.
2. **Analysis Class Diagram** và **Robustness Diagram** không phải cùng một loại biểu đồ:
   - Analysis Class Diagram nên dùng hộp lớp UML với stereotype `<<boundary>>`, `<<control>>`, `<<entity>>`, thuộc tính nghiệp vụ chính và quan hệ cấu trúc.
   - Robustness Diagram có thể dùng biểu tượng tròn BCE của Jacobson/ICONIX; nó mô tả cộng tác trong một use case và áp dụng quy tắc kết nối BCE chặt hơn.
3. Entity phân tích là khái niệm nghiệp vụ có thông tin và hành vi lâu dài; không được định nghĩa đơn giản là bảng cơ sở dữ liệu.
4. Có thể và nên vẽ cả:
   - một **Analysis Class Diagram tổng quát** để cung cấp bản đồ toàn hệ thống;
   - nhiều **Analysis Class Diagram tập trung** để giải thích từng nhóm nghiệp vụ/use case.
5. Các sơ đồ không phải bản sao độc lập. Chúng là các góc nhìn của cùng mô hình. Lớp trùng tên phải giữ nguyên stereotype, trách nhiệm và ý nghĩa quan hệ. Sơ đồ tổng quát chỉ giữ lớp/quan hệ cốt lõi; sơ đồ con bổ sung thuộc tính, trách nhiệm, multiplicity và lớp điều khiển/biên liên quan.
6. Không nên ép mọi biểu đồ thành một trang lớn. Nhiều góc nhìn ở các mức chi tiết khác nhau giúp kể đúng câu chuyện cho từng đối tượng đọc; mỗi góc nhìn phải dùng chung một model và giữ nhất quán tên, stereotype và quan hệ.

## 2. Phân loại nguồn và cách sử dụng

- **Chuẩn:** OMG UML 2.5.1 là cơ sở khi xác định ý nghĩa phần tử và quan hệ.
- **Phương pháp OOAD:** Jacobson/Unified Process và ICONIX dùng để dẫn xuất BCE, robustness và use-case realization.
- **Hướng dẫn thực hành:** Agile Modeling dùng cho mức chi tiết, khả năng đọc và lỗi mô hình phổ biến.
- **Tài liệu công cụ:** PlantUML dùng để xác nhận khả năng biểu diễn bằng mã nguồn; không dùng cú pháp công cụ để thay thế ngữ nghĩa UML.
- **ERD:** Dùng tài liệu mô hình dữ liệu/Information Engineering riêng; ERD không phải UML Class Diagram dù có phần ký pháp gần nhau.

## 3. Quy tắc theo từng loại biểu đồ

### 3.1 Use Case Diagram

**Mục đích**

Thể hiện phạm vi hệ thống, actor bên ngoài và mục tiêu có giá trị mà hệ thống cung cấp. Sơ đồ không mô tả thứ tự xử lý hay dữ liệu truyền qua lại.

**Ký pháp**

- Actor nằm ngoài system boundary.
- Use case là hình ellipse, đặt tên bằng động từ và bổ ngữ.
- Association actor-use case là đường liền; không cần đầu mũi tên nếu dễ bị hiểu nhầm thành data flow.
- `<<include>>`: mũi tên nét đứt từ use case cơ sở tới use case được bao gồm; hành vi được dùng là bắt buộc trong ngữ cảnh đó.
- `<<extend>>`: mũi tên nét đứt từ use case mở rộng về use case cơ sở; hành vi bổ sung xảy ra có điều kiện.
- Generalization: tam giác rỗng hướng về actor/use case tổng quát.

**Quy tắc trình bày**

- Sơ đồ tổng quan giữ mục tiêu lớn và actor chính.
- Dùng sơ đồ nhóm/chi tiết khi tổng quan quá dày.
- Không dùng package hay đường nối để biến use case diagram thành sơ đồ phân rã chức năng.

**Lỗi thường gặp**

- Tách bước xác thực, ghi log, cập nhật DB thành use case dù actor không có mục tiêu độc lập.
- Dùng `include` để nối các lựa chọn loại trừ nhau.
- Đảo chiều `extend`.
- Đặt precondition thành quan hệ UML.
- Nối actor vào use case được include/extend dù actor chỉ khởi tạo use case cơ sở.

**Nguồn**

- [Agile Modeling: UML Use Case Diagrams](https://agilemodeling.com/artifacts/useCaseDiagram.htm)
- [Visual Paradigm: What is Use Case Diagram?](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-use-case-diagram/)
- [PlantUML: Use Case Diagram](https://plantuml.com/use-case-diagram)

### 3.2 Activity Diagram

**Mục đích**

Mô tả workflow, luồng điều khiển, điều kiện, vòng lặp và xử lý song song của một use case hoặc quy trình nghiệp vụ.

**Ký pháp**

- Initial node, action, control flow, activity final/flow final.
- Decision có một luồng vào và nhiều luồng ra; mỗi nhánh cần guard rõ.
- Merge hợp nhất các nhánh lựa chọn; không chờ mọi nhánh.
- Fork tách luồng song song; join đồng bộ và chờ các luồng cần thiết.
- Partition/swimlane thể hiện bên chịu trách nhiệm, không phải module kỹ thuật tùy tiện.

**Quy tắc trình bày**

- Dòng chính nên đi từ trên xuống hoặc trái sang phải.
- Tên action bắt đầu bằng động từ.
- Dùng swimlane khi có actor, Gateway và external system.
- Tách sub-activity nếu một nhánh quá dài.
- Chỉ thể hiện lỗi/nhánh nghiệp vụ quan trọng; chi tiết message thuộc Sequence Diagram.

**Lỗi thường gặp**

- Dùng decision thay fork hoặc merge thay join.
- Guard thiếu, trùng hoặc không bao phủ trường hợp còn lại.
- Một action chứa nhiều hành động.
- Mô tả lời gọi method/object như sequence diagram.

**Nguồn**

- [Agile Modeling: UML Activity Diagrams](https://agilemodeling.com/artifacts/activityDiagram.htm)
- [Visual Paradigm: What is Activity Diagram?](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-activity-diagram/)
- [PlantUML: Activity Diagram](https://plantuml.com/activity-diagram-beta)
- Conrad Bock, *UML 2 Activity and Action Models*, Journal of Object Technology, 2003–2004.

### 3.3 Analysis Class Diagram theo BCE

**Mục đích**

Mô tả cấu trúc khái niệm ở mức phân tích: lớp nào cần để thực hiện yêu cầu, trách nhiệm của chúng và quan hệ nghiệp vụ giữa chúng.

**Ký pháp khuyến nghị**

- Dùng hộp lớp UML với `<<boundary>>`, `<<control>>`, `<<entity>>`.
- Có thể dùng icon BCE, nhưng không nên dùng icon tròn khi cần hiển thị thuộc tính, trách nhiệm, association và multiplicity.
- Boundary đại diện điểm tương tác với actor hoặc external system.
- Control điều phối kịch bản/use case hoặc quy tắc nghiệp vụ.
- Entity đại diện thông tin, khái niệm và hành vi nghiệp vụ có vòng đời đáng kể.

**Mức chi tiết**

- Có tên lớp, stereotype, 2–5 trách nhiệm hoặc thuộc tính nghiệp vụ chính nếu cần.
- Không cần visibility, kiểu ngôn ngữ lập trình, getter/setter, repository, DTO, controller framework hay bảng DB.
- Có thể dùng association, aggregation/composition khi ý nghĩa vòng đời thật sự rõ, generalization và multiplicity quan trọng.
- Entity–Entity association hợp lệ trong Analysis Class Diagram. Quy tắc cấm Entity–Entity chỉ là cách diễn giải sai khi chuyển quy tắc robustness sang class diagram.

**Dẫn xuất lớp**

- Boundary: từ từng loại actor-system interaction trong use case.
- Control: từ trách nhiệm điều phối của use case, business rule hoặc hoạt động liên quan nhiều lớp.
- Entity: từ danh từ nghiệp vụ, dữ liệu lâu dài, pre/postcondition, business rule và thông tin phải truy vết.

**Lỗi thường gặp**

- Đồng nhất Entity với table.
- Vẽ architecture layer hoặc code class thay vì lớp phân tích.
- Một control chung xử lý toàn hệ thống.
- Lặp control cho từng bước nhỏ trong basic flow.
- Không trace được lớp về use case/FR.

**Nguồn**

- Ivar Jacobson et al., *Object-Oriented Software Engineering: A Use Case Driven Approach*, Addison-Wesley/ACM Press, 1992.
- Ivar Jacobson, Grady Booch, James Rumbaugh, *The Unified Software Development Process*, Addison-Wesley, 1999.
- [Agile Modeling: UML Class Diagrams](https://agilemodeling.com/artifacts/classDiagram.htm)
- [Visual Paradigm: UML Class Diagram Tutorial](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/uml-class-diagram-tutorial/)
- [PlantUML: Class Diagram](https://plantuml.com/class-diagram)

### 3.4 Robustness Diagram

**Mục đích**

Kiểm tra logic use case và tạo cầu nối từ use case text sang lớp phân tích, sequence và thiết kế.

**Ký pháp**

- Actor, Boundary, Control, Entity.
- Icon tròn BCE của Jacobson/ICONIX phù hợp loại biểu đồ này.
- Không cần thuộc tính, method, kiểu dữ liệu hay multiplicity chi tiết.
- Mỗi robustness diagram thường bám một use case hoặc một scenario đủ nhỏ.

**Quy tắc kết nối**

- Actor giao tiếp với Boundary.
- Boundary giao tiếp với Actor và Control.
- Control giao tiếp với Boundary, Entity và Control khác khi cần.
- Trong convention nghiêm dùng cho đồ án, không vẽ Entity–Entity trong Robustness Diagram; quan hệ cấu trúc Entity–Entity được chuyển sang Analysis Class Diagram.
- Không nối Actor trực tiếp Control/Entity hoặc Boundary trực tiếp Entity.

**Lưu ý chuẩn**

Robustness Diagram không phải một loại diagram độc lập được chuẩn hóa đầy đủ như Use Case hay Class Diagram trong UML 2.5.1. Nó là kỹ thuật ICONIX/OOSE, thường biểu diễn bằng collaboration/communication-like view và stereotypes BCE.

Một số mô tả ECB rộng cho phép Entity biết Entity khác. Tuy nhiên, Visual Paradigm minh họa convention Robustness nghiêm không cho Entity–Entity. Đồ án chọn convention nghiêm để giữ robustness tập trung vào luồng cộng tác; điều này không cấm association Entity–Entity trong Analysis Class Diagram.

**Lỗi thường gặp**

- Biến robustness thành class diagram với thuộc tính/method.
- Dùng service/repository cụ thể.
- Một diagram bao phủ nhiều use case không liên quan.
- Vẽ mọi bước use case thành một Control.
- Dùng quy tắc robustness để cấm association giữa các Entity trong Analysis Class Diagram.

**Nguồn**

- Doug Rosenberg, Kendall Scott, *Use Case Driven Object Modeling with UML*, Addison-Wesley.
- [Springer/Apress: Robustness Analysis](https://link.springer.com/chapter/10.1007/978-1-4302-0369-8_5)
- [Agile Modeling: Robustness Diagrams](https://agilemodeling.com/artifacts/robustnessDiagram.htm)
- [Visual Paradigm: Robustness Diagram Connection Rules](https://online.visual-paradigm.com/diagrams/templates/robustness-diagram/robustness-diagram-connection-rules/)

### 3.5 Package Diagram

**Mục đích**

Tổ chức model elements theo namespace/module và thể hiện dependency giữa các package.

**Ký pháp**

- Package là hình thư mục có tab.
- Dependency là đường nét đứt, đầu mũi tên hướng về supplier.
- `package import` và `package merge` chỉ dùng khi thật sự cần đúng semantics; dependency module thông thường không nên bị gán nhãn tùy tiện.

**Quy tắc trình bày**

- Package đặt theo domain/responsibility, không sao chép cây thư mục.
- Hiển thị dependency cấp cao; chi tiết class nằm ở Design Class Diagram.
- Tránh vòng phụ thuộc; nếu tồn tại phải có giải thích thiết kế.
- Với modular monolith, package là module nội bộ cùng tiến trình, không vẽ như microservice.

**Lỗi thường gặp**

- Mọi package nối với mọi package.
- Trộn component, deployment node và database table.
- Dùng package chỉ để tạo khung màu.

**Nguồn**

- [Agile Modeling: UML Package Diagrams](https://agilemodeling.com/artifacts/packageDiagram.htm)
- [Visual Paradigm: What is Package Diagram?](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-package-diagram/)
- OMG UML 2.5.1, Packages.

### 3.6 Design Class Diagram

**Mục đích**

Mô tả solution model: lớp/interface hiện thực các trách nhiệm phân tích trong kiến trúc đã chọn.

**Ký pháp**

- Hộp lớp/interface; thuộc tính và operation có thể gồm visibility, parameter, return type khi có giá trị thiết kế.
- Association cho liên kết cấu trúc lâu dài.
- Dependency cho quan hệ sử dụng tạm thời.
- Generalization cho quan hệ “is-a”.
- Realization từ class/component tới interface.
- Composition chỉ dùng khi part thuộc độc quyền và vòng đời phụ thuộc whole; shared aggregation nên dùng thận trọng.

**Quy tắc trình bày**

- Tách diagram tổng quát và diagram theo flow/module.
- Tổng quát chỉ hiển thị module facade/service/interface và dependency chính.
- Sơ đồ con hiển thị service, repository contract, adapter, entity/value object cần cho use case.
- Không copy toàn bộ codebase hoặc mọi method.

**Lỗi thường gặp**

- Association và dependency dùng thay nhau.
- Repository phụ thuộc application service.
- Concrete provider class lan vào domain/application.
- Composition được dùng chỉ vì foreign key hoặc “has-a”.

**Nguồn**

- [Agile Modeling: UML Class Diagrams](https://agilemodeling.com/artifacts/classDiagram.htm)
- [Visual Paradigm: UML Class Diagram Tutorial](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/uml-class-diagram-tutorial/)
- Martin Fowler, *UML Distilled*, Addison-Wesley.
- [PlantUML: Class Diagram](https://plantuml.com/class-diagram)

### 3.7 Sequence Diagram

**Mục đích**

Mô tả tương tác có thứ tự thời gian của một scenario/use case realization.

**Ký pháp**

- Actor/participant, lifeline, activation, synchronous/asynchronous message và return.
- Thời gian đi từ trên xuống.
- Combined fragments: `alt`, `opt`, `loop`, `par`, `break`, `critical` khi có đúng semantics.
- Guard đặt trong ngoặc vuông.
- Creation/destruction chỉ thể hiện khi có giá trị thiết kế.

**Quy tắc trình bày**

- Một diagram cho một scenario chính; scenario lỗi lớn có thể tách riêng.
- Participant sắp từ actor, boundary/API, application service/control, domain/persistence adapter, external system.
- Message đặt theo ý định nghiệp vụ, tránh tên mơ hồ như `process()`.
- Với modular monolith, module nội bộ trao đổi bằng call; không dùng network arrow hay message broker nếu thiết kế không có.

**Lỗi thường gặp**

- Actor gọi DB trực tiếp.
- Dùng sequence để mô tả mọi dòng code.
- Return message dày đặc nhưng không thêm nghĩa.
- Dùng `alt` không có guard loại trừ rõ.
- Không khớp basic/alternative flow của use case.

**Nguồn**

- [Agile Modeling: UML Sequence Diagrams](https://agilemodeling.com/artifacts/sequenceDiagram.htm)
- [Visual Paradigm: What is Sequence Diagram?](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-sequence-diagram/)
- [PlantUML: Sequence Diagram](https://plantuml.com/sequence-diagram)

### 3.8 State Machine Diagram

**Mục đích**

Mô tả vòng đời của một classifier/object có trạng thái đáng kể và phản ứng với event.

**Ký pháp**

- Initial/final pseudostate, state, transition.
- Nhãn transition theo dạng `trigger [guard] / effect`.
- State có thể có `entry`, `do`, `exit`.
- Composite state và region chỉ dùng khi giúp giảm độ phức tạp.

**Quy tắc trình bày**

- State là tình trạng ổn định, thường đặt bằng danh từ/tính từ; action không phải state.
- Transition dùng dạng tổng quát `trigger [guard] / effect`; mỗi thành phần có thể được lược bỏ. UML cho phép completion transition không có trigger.
- Guard từ cùng một state cho cùng event nên không chồng lấn.
- Tách state machine cho `AIRequest` và `FinancialTransaction`; không gộp hai vòng đời độc lập.

**Lỗi thường gặp**

- Vẽ activity bằng các state.
- Thiếu failure/rejection/reversal.
- Transition không có nguyên nhân.
- Trạng thái phản ánh từng method nội bộ.

**Nguồn**

- David Harel, “Statecharts: A Visual Formalism for Complex Systems”, *Science of Computer Programming*, 1987.
- [Visual Paradigm: What is State Machine Diagram?](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-state-machine-diagram/)
- [PlantUML: State Diagram](https://plantuml.com/state-diagram)

### 3.9 ERD

**Mục đích**

Mô tả cấu trúc dữ liệu khái niệm hoặc logic: entity, identifier, relationship, cardinality, optionality và constraint quan trọng.

**Ký pháp**

- Chọn một notation nhất quán: Crow’s Foot/Information Engineering hoặc Chen.
- Nếu dùng Crow’s Foot, mỗi đầu quan hệ thể hiện minimum và maximum cardinality.
- Logical ERD có identifier, thuộc tính chính và associative entity; không cần SQL type/index.
- Physical schema diagram mới cần kiểu DBMS, PK/FK cụ thể, index và constraint triển khai.

**Quy tắc trình bày**

- Ghi rõ mức conceptual, logical hay physical.
- Chương 4 nên dùng logical ERD, dẫn xuất từ Entity analysis và business rules.
- Dùng ERD tổng quan thưa + ERD con theo cụm khi số entity lớn.
- Không dùng Entity phân tích và database table như hai khái niệm đồng nhất.

**Lỗi thường gặp**

- Không ghi optionality.
- Dùng nhiều-many mà thiếu associative entity ở logical model.
- Trộn service/control vào ERD.
- Nhồi index, partition, migration vào phần thiết kế logic.

**Nguồn**

- Peter P. Chen, “The Entity-Relationship Model—Toward a Unified View of Data”, *ACM TODS*, 1976.
- [IBM: What is Data Modeling?](https://www.ibm.com/think/topics/data-modeling)
- [IBM: What is an Entity Relationship Diagram?](https://www.ibm.com/think/topics/entity-relationship-diagram)
- [PlantUML: Information Engineering Diagram](https://plantuml.com/ie-diagram)

### 3.10 Component Diagram

**Mục đích**

Mô tả các đơn vị phần mềm có interface cung cấp/yêu cầu và cách chúng được ghép thành hệ thống.

**Ký pháp**

- Component dùng `<<component>>` hoặc biểu tượng component.
- Provided interface dùng lollipop; required interface dùng socket.
- Assembly connector nối interface tương thích.
- Dependency dùng khi không cần biểu diễn interface đầy đủ.

**Quy tắc trình bày**

- Component phải có ranh giới và hợp đồng rõ; package nội bộ không tự động là component.
- Với Asymptotic modular monolith, có thể vẽ một backend component chứa các module hoặc vài component triển khai lớn; không tách mỗi module thành microservice.
- External AI Provider, Payment Provider và DB được phân biệt rõ.

**Lỗi thường gặp**

- Component Diagram giống Package Diagram đổi hình.
- Nối component bằng đường không giải thích interface/protocol.
- Phóng đại hệ thống thành kiến trúc phân tán production.

**Nguồn**

- [Agile Modeling: UML Component Diagrams](https://agilemodeling.com/artifacts/componentDiagram.htm)
- [Visual Paradigm: What is Component Diagram?](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-component-diagram/)
- [PlantUML: Component Diagram](https://plantuml.com/component-diagram)

### 3.11 Deployment Diagram

**Mục đích**

Mô tả execution environment, device/node, artifact được triển khai và communication path.

**Ký pháp**

- Device node và execution environment là node 3D.
- Artifact đặt trong/deploy lên node.
- Communication path nối node; có thể ghi protocol.
- External service có thể dùng node/component stereotype phù hợp nhưng phải phân biệt khỏi artifact nội bộ.

**Quy tắc trình bày**

- Bám môi trường MVP hoặc prototype thực tế.
- Phân biệt logical module với physical deployment.
- Nếu backend modular monolith chạy một process, phải thể hiện một deployable backend, không biến module thành node.

**Lỗi thường gặp**

- Dùng component thay artifact/node.
- Vẽ cloud cluster, HA, autoscaling chưa tồn tại.
- Không thể hiện DB hoặc external systems.

**Nguồn**

- [Agile Modeling: UML Deployment Diagrams](https://agilemodeling.com/artifacts/deploymentDiagram.htm)
- [Visual Paradigm: What is Deployment Diagram?](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-deployment-diagram/)
- [PlantUML: Deployment Diagram](https://plantuml.com/deployment-diagram)

## 4. Cách kết hợp Analysis Class Diagram tổng quát và sơ đồ con

### 4.1 Nguyên tắc

UML phân biệt **model** và **diagram**: diagram là góc nhìn một phần của model. Vì vậy cùng một lớp có thể xuất hiện trên nhiều diagram mà không tạo ra nhiều định nghĩa lớp khác nhau.

### 4.2 Analysis Class Diagram tổng quát

Mục tiêu: cung cấp bản đồ 30–60 giây về toàn bộ hệ thống.

Nên hiển thị:

- Các nhóm BCE chính hoặc package nghiệp vụ.
- Boundary chính cho AI Agent, quản trị tổ chức, quản trị hệ thống và external providers.
- Control cấp use case/domain chính.
- Entity cốt lõi.
- Chỉ quan hệ trục chính; không hiển thị toàn bộ thuộc tính.

Không nên hiển thị:

- Mọi boundary/control của chín use case.
- Mọi association và multiplicity.
- Chi tiết workflow.
- Repository/service/framework.

Khuyến nghị giới hạn mềm: khoảng 20–30 lớp nhìn thấy. Nếu vượt và nhiều đường cắt nhau, giảm quan hệ hoặc dùng package/grouping.

### 4.3 Các Analysis Class Diagram tập trung

Đề xuất sáu cụm:

1. **Gateway request:** UC01; AIRequest, IdempotencyRecord, UsagePolicy, BudgetReservation, FinancialTransaction và ExecutionTrace.
2. **Tài chính và ngân sách:** UC02, UC03; Wallet, BudgetLimit, BudgetAllocation, BudgetReservation, FinancialTransaction, LedgerEntry và PaymentTransaction.
3. **Agent và API key:** UC04, UC05; AIAgent, ApiKey, AgentAssignment và AgentStatus.
4. **Tổ chức và thành viên:** UC07, UC08; Organization, Team, User, DeveloperProfile, OrganizationMembership và TeamMembership.
5. **Usage, cost và trace:** UC06; AIRequest, UsageRecord, CostRecord, FinancialTransaction, LedgerEntry và RequestTrace.
6. **Provider, model và chính sách:** UC09 và phần định tuyến của UC01; AIProvider, ProviderCredential, AIModel, ModelPricing và RoutingPolicy.

Mỗi sơ đồ con nên có:

- Boundary gắn actor/external system của nhóm use case.
- Control chịu trách nhiệm điều phối use case.
- Entity và quan hệ cần cho nhóm.
- Thuộc tính/trách nhiệm mức phân tích.
- Trace tới UC/FR/NFR.

### 4.4 Quy tắc chống lặp và mâu thuẫn

- Cùng lớp: cùng tên, stereotype và trách nhiệm cốt lõi trên mọi diagram.
- Overview là nguồn định hướng; subdiagram là nơi thêm chi tiết.
- Quan hệ xuất hiện ở overview phải giữ cùng semantics trong subdiagram.
- Lớp dùng chung có thể lặp như “anchor”, nhưng chỉ hiển thị quan hệ liên quan phạm vi hiện tại.
- Không sao chép toàn bộ overview vào từng subdiagram.
- Lưu một glossary/class catalog riêng: class, stereotype, trách nhiệm, nguồn UC/FR.
- Review hai chiều:
  - mỗi class phải trace được về requirement/use case;
  - mỗi use case quan trọng phải có Boundary, Control và Entity phù hợp.

### 4.5 Bố cục

Với Analysis Class Diagram:

- Có thể xếp Boundary → Control → Entity theo trái-phải hoặc trên-dưới để tăng khả năng đọc.
- Đây là quy ước bố cục, không phải quy tắc UML bắt buộc.
- Entity có thể nối Entity; không cần ép mọi đường qua Control.
- Tối thiểu hóa đường cắt; đặt lớp dùng chung ở trung tâm nhóm.
- Dùng cách tiếp cận “overview + zoomed views”: overview kể câu chuyện toàn hệ thống, từng sơ đồ con kể một phạm vi hẹp hơn. C4 dùng nguyên tắc mức zoom khác nhau cho các audience khác nhau; ở đây chỉ mượn nguyên tắc phân tầng góc nhìn, không thay ký pháp UML bằng C4.

Với Robustness Diagram:

- Actor → Boundary → Control → Entity là dòng đọc ưu tiên.
- Áp dụng quy tắc BCE nghiêm hơn.
- Chỉ tập trung một scenario/use case.

## 5. Khuyến nghị cập nhật bộ tài liệu dự án

1. Giữ mục **Analysis Class Diagram tổng quát** trong `ooad_diagram_priority_list.md`.
2. Thay các mục Boundary/Control/Entity riêng rời bằng các **Analysis Class Diagram tập trung theo nhóm nghiệp vụ**, vì một sơ đồ chỉ chứa mọi Boundary hoặc mọi Control làm mất ngữ cảnh use case.
3. Giữ hai Robustness Diagram ưu tiên cho UC01 và UC04.
4. Bổ sung vào `diagram_type_rubric.md`:
   - Analysis Class khác Robustness;
   - icon BCE là tùy chọn;
   - Entity không đồng nghĩa table;
   - Entity–Entity hợp lệ trong class analysis;
   - quy tắc BCE chặt áp dụng cho robustness.
5. Thêm trường vào checklist:
   - Mức mô hình: requirements/analysis/design/deployment.
   - Diagram là overview hay focused view.
   - Model element trùng có nhất quán với diagram khác không.
6. Tạo nguồn PlantUML riêng cho từng view; không dùng `newpage` để nhét nhiều view khác mục đích vào một file ảnh.
7. Dùng một class catalog chung và, nếu phù hợp, PlantUML include chung để giảm nguy cơ lệch tên/stereotype giữa overview và sáu sơ đồ con.

## 6. Quy tắc chốt cho Asymptotic

- Chương 2: Use Case và Activity; không có class/service/repository.
- Chương 3:
  - một Analysis Class Diagram tổng quát;
  - sáu Analysis Class Diagram tập trung;
  - Robustness Diagram UC01 và UC04.
- Chương 4:
  - Package Diagram;
  - Design Class overview và các design class theo flow;
  - Sequence, State Machine;
  - Logical ERD;
  - Component/Deployment khi cần chứng minh kiến trúc và môi trường chạy.
- Tất cả diagram dùng PlantUML, có source cạnh export, trace tới UC/FR/NFR và được review bằng rubric riêng của đúng loại diagram.

## 7. Giới hạn nghiên cứu

- Một số tài liệu vendor mô tả notation theo khả năng công cụ, không phải chuẩn semantics; vì vậy chúng chỉ được dùng cho cú pháp và thực hành trình bày.
- Robustness Diagram là kỹ thuật phương pháp luận, không phải diagram chuẩn hóa độc lập ngang hàng với các UML diagram chính.
- Các ngưỡng số lớp/node là heuristic cho khả năng đọc, không phải quy định UML.
