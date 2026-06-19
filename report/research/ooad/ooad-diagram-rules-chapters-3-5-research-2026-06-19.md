# Research Brief: Quy tắc biểu đồ cho Chương 3, 4 và 5

**Ngày nghiên cứu:** 2026-06-19  
**Phạm vi:** Biểu đồ phân tích hướng đối tượng, thiết kế hướng đối tượng, triển khai và minh chứng kiểm thử.  
**Quan hệ với nghiên cứu trước:** Tài liệu này bổ sung cho `ooad-diagram-rules-multisource-research-2026-06-18.md`, tập trung vào phân loại nguồn, ranh giới giữa các chương và các điểm dễ áp dụng sai.

## 1. Kết luận chính

### 1.1 Phải tách ba lớp quy tắc

1. **UML chuẩn:** xác định phần tử, quan hệ và ngữ nghĩa theo OMG UML 2.5.1.
2. **Phương pháp OOAD:** xác định cách dùng Analysis Class, BCE, Robustness và use-case realization.
3. **Quy ước dự án/công cụ:** xác định bố cục, tên file, mức chi tiết, PlantUML và traceability.

Không được mô tả quy ước BCE, ICONIX, kiến trúc phân lớp hoặc cách bố trí PlantUML như yêu cầu bắt buộc của UML.

### 1.2 Phân loại các biểu đồ

| Biểu đồ hoặc góc nhìn | Phân loại | Chương chính |
|---|---|---|
| Analysis Class Diagram | Cách dùng UML Class Diagram ở mức phân tích | 3 |
| Boundary, Control, Entity views | Góc nhìn BCE của mô hình lớp phân tích | 3 |
| Robustness Diagram | Kỹ thuật OOSE/ICONIX, dùng ký pháp BCE | 3 |
| Package Diagram | UML chuẩn | 4 |
| Design Class Diagram | Cách dùng UML Class Diagram ở mức thiết kế | 4 |
| Sequence Diagram | UML chuẩn | 4 |
| State Machine Diagram | UML chuẩn | 4 |
| Component Diagram | UML chuẩn | 4 hoặc 5 tùy câu hỏi |
| Deployment Diagram | UML chuẩn | 4 nếu là thiết kế; 5 nếu là triển khai thực tế |
| ERD | Ký pháp mô hình dữ liệu, không thuộc UML | 4 |
| Ma trận traceability, bảng test, biểu đồ hiệu năng | Minh chứng kỹ thuật, không thuộc UML | 5 |

### 1.3 Ranh giới Chương 4 và Chương 5

- Chương 4 trả lời **hệ thống được thiết kế như thế nào**.
- Chương 5 trả lời **hệ thống đã được triển khai và kiểm chứng như thế nào**.
- Component hoặc Deployment Diagram ở Chương 4 phải được ghi là **as-designed**.
- Component hoặc Deployment Diagram ở Chương 5 phải phản ánh **as-built** và chỉ chứa thành phần thực tế đã triển khai.
- Không cần tạo biểu đồ Chương 5 khi chưa có bằng chứng triển khai. Không được dựng kiến trúc cloud, cluster hoặc microservice chỉ để làm báo cáo dày hơn.

## 2. Quy tắc Chương 3

### 2.1 Analysis Class Diagram

Analysis Class Diagram không phải loại UML riêng. Đây là UML Class Diagram được giới hạn ở mức phân tích.

Nên thể hiện:

- lớp nghiệp vụ;
- stereotype `<<boundary>>`, `<<control>>`, `<<entity>>`;
- trách nhiệm hoặc thuộc tính nghiệp vụ quan trọng;
- association, generalization và multiplicity có ý nghĩa;
- nguồn truy vết từ use case hoặc yêu cầu.

Không nên thể hiện:

- controller framework, DTO, repository, migration;
- kiểu dữ liệu phụ thuộc ngôn ngữ lập trình;
- operation kỹ thuật hoặc toàn bộ codebase;
- bảng cơ sở dữ liệu như thể chúng đồng nhất với Entity phân tích.

### 2.2 BCE

`Boundary`, `Control` và `Entity` là các stereotype/phân loại phương pháp của mô hình phân tích, không phải ba loại UML Class Diagram.

- **Boundary:** điểm giao tiếp với actor hoặc hệ thống ngoài.
- **Control:** điều phối scenario, use case hoặc quy tắc nghiệp vụ.
- **Entity:** khái niệm nghiệp vụ có danh tính, thông tin hoặc vòng đời đáng kể.

Entity có thể liên hệ trực tiếp với Entity trong Analysis Class Diagram. Không áp dụng máy móc quy tắc kết nối Robustness Diagram cho quan hệ cấu trúc giữa các lớp.

### 2.3 Robustness Diagram

Robustness Diagram không được OMG UML 2.5.1 chuẩn hóa thành một loại biểu đồ độc lập. Đồ án dùng nó như kỹ thuật OOSE/ICONIX để kiểm tra use case và dẫn xuất lớp phân tích.

Quy ước dự án:

- Actor giao tiếp với Boundary.
- Boundary giao tiếp với Control.
- Control giao tiếp với Boundary, Entity hoặc Control khác.
- Không nối Actor trực tiếp tới Control hoặc Entity.
- Không đưa thuộc tính, method, repository hoặc chi tiết triển khai.
- Mỗi biểu đồ bám một use case hoặc một scenario hẹp.

Các quy tắc trên phải được ghi là **quy ước phương pháp của đồ án**, không phải quy tắc cú pháp UML.

## 3. Quy tắc Chương 4

### 3.1 Package Diagram

UML chỉ quy định package, namespace và dependency. Các quy tắc đặt package theo domain, tránh vòng phụ thuộc hoặc chỉ cho phép phụ thuộc theo một chiều là quyết định kiến trúc của dự án.

Đối với modular monolith:

- package biểu diễn module logic trong cùng ứng dụng;
- không vẽ mỗi package như một service triển khai độc lập;
- không gán một cơ sở dữ liệu riêng cho mỗi package;
- dependency hướng từ client package tới supplier package.

### 3.2 Design Class Diagram

Phải phân biệt đúng quan hệ:

| Quan hệ | Ý nghĩa | PlantUML |
|---|---|---|
| Association | Liên kết cấu trúc hoặc tham chiếu lâu dài | `--`, `-->` |
| Dependency | Sử dụng tạm thời, gọi qua tham số hoặc operation | `..>` |
| Generalization | Lớp chuyên biệt kế thừa lớp tổng quát | `--|>` |
| Realization | Lớp hiện thực interface/contract | `..|>` |
| Aggregation | Whole-part nhưng part tồn tại độc lập | `o--` |
| Composition | Whole sở hữu độc quyền và chi phối vòng đời part | `*--` |

Không dùng composition chỉ vì có khóa ngoại hoặc quan hệ “has-a”. Shared aggregation chỉ dùng khi thật sự cần diễn đạt whole-part; association thường rõ hơn.

### 3.3 Sequence Diagram

- Thời gian tăng từ trên xuống.
- Message phải phản ánh trách nhiệm hoặc operation có ý nghĩa.
- Dùng `alt`, `opt`, `loop`, `par`, `break`, `critical` đúng ngữ nghĩa.
- Reply message là tùy chọn; chỉ vẽ khi kết quả trả về quan trọng với scenario.
- Luồng lỗi quan trọng có thể đặt trong `alt` hoặc tách thành biểu đồ riêng nếu làm luồng chính khó đọc.
- Với modular monolith, module nội bộ là participant/lifeline trong cùng ứng dụng, không phải các microservice giao tiếp qua mạng.
- AI Agent không gọi cơ sở dữ liệu hoặc AI Provider bằng provider credential.

### 3.4 State Machine Diagram

Nhãn transition có dạng tổng quát:

```text
trigger [guard] / effect
```

Mỗi thành phần đều có thể được lược bỏ. Transition không có trigger có thể là completion transition. Vì vậy không được bắt buộc mọi transition phải có cả event, guard và effect.

Tên state bằng danh từ hoặc tính từ là heuristic để tăng khả năng đọc, không phải cú pháp UML. Tiêu chí quan trọng là state biểu diễn một điều kiện ổn định có ý nghĩa với vòng đời đối tượng.

### 3.5 Component Diagram

- Component biểu diễn đơn vị mô-đun có hợp đồng rõ.
- Package nội bộ không tự động là component.
- Provided interface dùng lollipop; required interface dùng socket khi cần làm rõ hợp đồng.
- Dependency có thể dùng trong góc nhìn tổng quan.
- Black-box và white-box đều hợp lệ; nên tách thành hai view nếu việc trộn làm biểu đồ khó đọc.
- Với modular monolith, backend có thể là một component triển khai lớn, bên trong được phân rã thành module/package.

### 3.6 Deployment Diagram

UML Deployment Diagram dùng node, device, execution environment, artifact, deployment và communication path. UML không bắt buộc mọi biểu đồ phải có client, server, database và external provider; đó là tiêu chí đầy đủ theo phạm vi dự án.

Đối với Asymptotic:

- Chương 4 chỉ vẽ deployment dự kiến nếu thiết kế triển khai đã được chốt.
- Chương 5 vẽ deployment thực tế nếu có môi trường chạy và bằng chứng.
- Phải phân biệt component logic, artifact triển khai và runtime node.
- Không mô hình hóa module nội bộ thành node hoặc service độc lập.

### 3.7 ERD

ERD không thuộc UML. Phải ghi rõ mức:

- **Conceptual:** entity và relationship chính;
- **Logical:** identifier, attribute, cardinality, optionality và associative entity;
- **Physical:** table, PK/FK, type, index và constraint phụ thuộc DBMS.

Yêu cầu PK/FK không áp dụng bắt buộc cho conceptual ERD. Chương 4 nên ưu tiên logical ERD; physical schema chỉ đưa vào khi cần chứng minh thiết kế dữ liệu triển khai.

## 4. Quy tắc Chương 5

### 4.1 Biểu đồ UML có thể dùng

- **As-built Component Diagram:** module/component thực tế đã triển khai.
- **As-built Deployment Diagram:** artifact, node, môi trường chạy và kết nối thực tế.

Hai biểu đồ này là tùy chọn. Chỉ dùng khi chúng bổ sung bằng chứng mà Chương 4 chưa thể hiện.

### 4.2 Minh chứng không phải UML

- ma trận Requirement--Use Case--Module--Test;
- bảng test case và kết quả;
- biểu đồ tỷ lệ pass/fail;
- biểu đồ độ trễ, throughput hoặc chi phí;
- bảng đối chiếu thiết kế với triển khai;
- ảnh giao diện hoặc log minh chứng.

Phải ghi đúng loại, không gọi các artifact này là UML diagram.

### 4.3 Quy tắc kiểm thử và đánh giá

- Mỗi test phải truy vết về FR, NFR hoặc use case.
- Kết quả phải phân biệt expected result và actual result.
- Biểu đồ hiệu năng phải ghi workload, đơn vị, môi trường, cỡ mẫu và cách đo.
- Khi đo UC01, nên phân biệt Gateway overhead với thời gian AI Provider để tránh kết luận sai.
- ISO/IEC/IEEE 29119 dùng làm hướng dẫn tài liệu kiểm thử; ISO/IEC 25010 dùng để nhóm mục tiêu chất lượng. Không cần áp dụng toàn bộ profile UTP2 cho MVP nếu không tạo mô hình kiểm thử UML chuyên biệt.

## 5. Ví dụ và mức tin cậy

- Ví dụ OMG đã trích trong repository dùng để kiểm tra ký pháp UML chuẩn.
- Ví dụ PlantUML dùng để kiểm tra cú pháp render, không thay thế ngữ nghĩa UML.
- Robustness/BCE phải dùng ví dụ phương pháp hoặc ví dụ dự án vì OMG không cung cấp Robustness Diagram chuẩn.
- Hình dự án tự thiết kế không cần dòng nguồn; hình trích từ tiêu chuẩn hoặc website phải có trích dẫn.

## 6. Quyết định áp dụng cho repository

1. Dùng `uml_2_5_1_drawing_rules.md` cho ngữ nghĩa UML và ghi rõ phần nào là quy ước dự án.
2. Dùng `diagram_type_rubric.md` cho mục đích, mức chi tiết và lỗi thường gặp.
3. Dùng `diagram_review_checklist.md` để duyệt từng biểu đồ.
4. Dùng `guidelines/diagrams/README.md` làm chỉ mục nguồn, ví dụ và vị trí chương.
5. Chương 5 chưa bắt buộc có UML diagram. Chỉ thêm as-built Component/Deployment Diagram khi dữ liệu triển khai đã tồn tại.

## 7. Giới hạn

- Nội dung tiêu chuẩn ISO công khai chủ yếu là metadata và abstract; không tuyên bố các chi tiết không thể xác minh từ bản tiêu chuẩn đầy đủ.
- UTP2 2.3 beta được OMG công bố tháng 3/2025 nhưng là bản beta. Khi cần tuân thủ chính thức, ưu tiên UTP2 2.2 formal tháng 8/2024.
- PlantUML có nhiều cú pháp trình bày ngoài UML; chỉ các phần tử có ngữ nghĩa phù hợp mới được dùng trong biểu đồ UML của đồ án.
