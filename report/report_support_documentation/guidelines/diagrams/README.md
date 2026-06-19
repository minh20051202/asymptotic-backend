# Chỉ mục quy tắc và ví dụ diagram

Tài liệu này là điểm vào duy nhất khi tạo hoặc review biểu đồ cho báo cáo.

## 1. Thứ tự ưu tiên nguồn

1. [`project_source_of_truth.md`](../../governance/project_source_of_truth.md): nội dung, phạm vi, actor, use case và kiến trúc dự án.
2. [`uml_2_5_1_drawing_rules.md`](uml_2_5_1_drawing_rules.md): ngữ nghĩa UML và ký pháp.
3. [`diagram_type_rubric.md`](diagram_type_rubric.md): mục đích và mức chi tiết của từng loại biểu đồ.
4. [`diagram_quality_guidelines.md`](diagram_quality_guidelines.md): khả năng đọc, traceability và trình bày.
5. [`diagram_review_checklist.md`](diagram_review_checklist.md): cổng duyệt trước khi đưa vào báo cáo.
6. [`ooad_diagram_priority_list.md`](ooad_diagram_priority_list.md): thứ tự ưu tiên thực hiện.

## 2. Phân loại notation

| Nhóm | Biểu đồ/góc nhìn | Nguồn quy tắc |
|---|---|---|
| UML chuẩn | Class, Package, Sequence, State Machine, Component, Deployment | OMG UML 2.5.1 |
| Cách dùng UML | Analysis Class, Design Class | UML + phương pháp OOAD |
| Phương pháp OOAD | BCE, Robustness | OOSE/ICONIX + quy ước dự án |
| Ngoài UML | ERD | ER modeling/Information Engineering |
| Minh chứng | Test matrix, bảng kết quả, biểu đồ hiệu năng | Test documentation và quy tắc báo cáo |

PlantUML là công cụ biểu diễn. Cú pháp PlantUML không thay thế ngữ nghĩa của UML hoặc ERD.

## 3. Vị trí theo chương

### Chương 3 -- Phân tích hướng đối tượng

- Analysis Class Diagram tổng quát.
- Boundary, Control và Entity views.
- Robustness Diagram UC01.
- Robustness Diagram UC04.

Không đưa repository, adapter, framework, database table hoặc deployment node vào Chương 3.

### Chương 4 -- Thiết kế hướng đối tượng

- Package Diagram.
- Design Class Diagram tổng quát và theo flow.
- Sequence Diagram.
- State Machine Diagram.
- Logical ERD khi hoàn thiện thiết kế dữ liệu.
- Component/Deployment Diagram as-designed nếu cần làm rõ kiến trúc.

### Chương 5 -- Triển khai, kiểm thử và đánh giá

- As-built Component Diagram, tùy chọn.
- As-built Deployment Diagram, tùy chọn.
- Ma trận Requirement--Use Case--Module--Test.
- Bảng test case và kết quả.
- Biểu đồ hiệu năng hoặc chất lượng có phương pháp đo rõ.

Không bắt buộc tạo UML diagram cho Chương 5 khi chưa có dữ liệu triển khai.

## 4. Ví dụ chuẩn OMG

Chỉ mục ảnh trích từ UML 2.5.1:

- [`uml_2_5_1_examples/README.md`](../../references/examples/uml_2_5_1_examples/README.md)

Các ảnh này dùng để đối chiếu ký pháp, không chèn vào báo cáo như hình tự thiết kế.

## 5. Ví dụ dự án

| Chương | Loại/phạm vi | Trace chính | Source | Export |
|---|---|---|---|---|
| 3 | Analysis Class overview | UC01--UC09 | [`diagram.puml`](../../diagrams/chapter_3/analysis_classes/diagram.puml) | [`diagram.png`](../../diagrams/chapter_3/analysis_classes/diagram.png) |
| 3 | Boundary view | Actor/external interfaces | [`boundary.puml`](../../diagrams/chapter_3/analysis_classes/sections/boundary.puml) | [`boundary.png`](../../diagrams/chapter_3/analysis_classes/sections/boundary.png) |
| 3 | Control view | Use-case orchestration | [`control.puml`](../../diagrams/chapter_3/analysis_classes/sections/control.puml) | [`control.png`](../../diagrams/chapter_3/analysis_classes/sections/control.png) |
| 3 | Entity view | Domain concepts | [`entity.puml`](../../diagrams/chapter_3/analysis_classes/sections/entity.puml) | [`entity.png`](../../diagrams/chapter_3/analysis_classes/sections/entity.png) |
| 3 | Robustness | UC01 | [`diagram.puml`](../../diagrams/chapter_3/robustness/UC01/diagram.puml) | [`diagram.png`](../../diagrams/chapter_3/robustness/UC01/diagram.png) |
| 3 | Robustness | UC04 | [`diagram.puml`](../../diagrams/chapter_3/robustness/UC04/diagram.puml) | [`diagram.png`](../../diagrams/chapter_3/robustness/UC04/diagram.png) |
| 4 | Package | Modular monolith | [`diagram.puml`](../../diagrams/chapter_4/package/diagram.puml) | [`diagram.png`](../../diagrams/chapter_4/package/diagram.png) |
| 4 | Design Class overview | UC01--UC09 | [`overview.puml`](../../diagrams/chapter_4/design_classes/overview.puml) | [`overview.png`](../../diagrams/chapter_4/design_classes/overview.png) |
| 4 | Design Class | UC01 | [`gateway_request_flow.puml`](../../diagrams/chapter_4/design_classes/gateway_request_flow.puml) | [`gateway_request_flow.png`](../../diagrams/chapter_4/design_classes/gateway_request_flow.png) |
| 4 | Design Class | Finance/Ledger | [`finance_ledger.puml`](../../diagrams/chapter_4/design_classes/finance_ledger.puml) | [`finance_ledger.png`](../../diagrams/chapter_4/design_classes/finance_ledger.png) |
| 4 | Sequence | UC01 | [`sequence.puml`](../../diagrams/chapter_4/sequences/UC01/sequence.puml) | [`sequence.png`](../../diagrams/chapter_4/sequences/UC01/sequence.png) |
| 4 | Sequence | UC04 | [`sequence.puml`](../../diagrams/chapter_4/sequences/UC04_agent_registration/sequence.puml) | [`sequence.png`](../../diagrams/chapter_4/sequences/UC04_agent_registration/sequence.png) |
| 4 | Sequence | UC02/UC03 | [`sequence.puml`](../../diagrams/chapter_4/sequences/finance_budget/sequence.puml) | [`sequence.png`](../../diagrams/chapter_4/sequences/finance_budget/sequence.png) |
| 4 | State Machine | AI Request | [`ai_request_state.puml`](../../diagrams/chapter_4/states/ai_request_state.puml) | [`ai_request_state.png`](../../diagrams/chapter_4/states/ai_request_state.png) |
| 4 | State Machine | Financial Transaction | [`financial_transaction_state.puml`](../../diagrams/chapter_4/states/financial_transaction_state.puml) | [`financial_transaction_state.png`](../../diagrams/chapter_4/states/financial_transaction_state.png) |

Tên file như `diagram.puml` có thể lặp giữa các thư mục có phạm vi rõ. Cặp source/export phải nằm cùng thư mục.

## 6. Nghiên cứu nền

- [`ooad-diagram-rules-multisource-research-2026-06-18.md`](../../../research/ooad/ooad-diagram-rules-multisource-research-2026-06-18.md)
- [`ooad-diagram-rules-chapters-3-5-research-2026-06-19.md`](../../../research/ooad/ooad-diagram-rules-chapters-3-5-research-2026-06-19.md)

## 7. Trạng thái thiếu

- Chưa có logical ERD chính thức trong `diagrams/chapter_4/`.
- Chưa có Component Diagram chính thức.
- Chưa có as-built Deployment Diagram hoặc minh chứng diagram cho Chương 5.
- Các mục này chỉ được bổ sung khi mô hình dữ liệu, kiến trúc hoặc môi trường triển khai đã được chốt.
