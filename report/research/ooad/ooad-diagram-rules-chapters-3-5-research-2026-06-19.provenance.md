# Provenance: Quy tắc biểu đồ Chương 3, 4 và 5

**Ngày truy cập:** 2026-06-19

## Nguồn chuẩn

1. Object Management Group. *OMG Unified Modeling Language Version 2.5.1*. December 2017.  
   https://www.omg.org/spec/UML/2.5.1  
   Bản PDF chuẩn: https://www.omg.org/spec/UML/2.5.1/PDF  
   Dùng để xác minh các loại biểu đồ UML, quan hệ lớp, interaction, state machine, component và deployment.

2. Object Management Group. *UML Testing Profile 2, Version 2.2*. August 2024, formal.  
   https://www.omg.org/spec/UTP2/2.2/About-UTP2  
   Dùng để xác minh UTP2 liên kết test artifact với requirement, risk, use case, business process và system specification.

3. Object Management Group. *UML Testing Profile 2, Version 2.3 beta*. March 2025.  
   https://www.omg.org/spec/UTP2/2.3/Beta1  
   Dùng để xác minh đây là bản beta, chỉ có tính thông tin; bản formal mới là cơ sở tuân thủ.

4. ISO/IEC/IEEE. *Systems and software engineering -- Architecture description*, ISO/IEC/IEEE 42010:2022.  
   https://www.iso.org/standard/74393.html  
   Dùng cho nguyên tắc phân biệt viewpoint, concern và mô tả kiến trúc.

5. ISO/IEC/IEEE. *Software and systems engineering -- Software testing -- Part 1: General concepts*, ISO/IEC/IEEE 29119-1:2022.  
   https://www.iso.org/standard/81291.html

6. ISO/IEC/IEEE. *Software and systems engineering -- Software testing -- Part 3: Test documentation*, ISO/IEC/IEEE 29119-3:2021.  
   https://www.iso.org/standard/79429.html

7. ISO/IEC. *Systems and software engineering -- Systems and software Quality Requirements and Evaluation (SQuaRE) -- Product quality model*, ISO/IEC 25010:2023.  
   https://www.iso.org/standard/78176.html

## Nguồn phương pháp và nghiên cứu

8. Ivar Jacobson, Magnus Christerson, Patrik Jonsson, Gunnar Övergaard. *Object-Oriented Software Engineering: A Use Case Driven Approach*. Addison-Wesley/ACM Press, 1992.  
   Dùng cho use-case-driven analysis và BCE.

9. Doug Rosenberg, Kendall Scott. “Robustness Analysis.” In *Use Case Driven Object Modeling with UML*. Apress.  
   https://doi.org/10.1007/978-1-4302-0369-8_5  
   Dùng cho vai trò cầu nối của Robustness Analysis và các phần tử boundary, control, entity.

10. Peter P. Chen. “The Entity-Relationship Model--Toward a Unified View of Data.” *ACM Transactions on Database Systems*, 1976.  
    https://doi.org/10.1145/320434.320440  
    Dùng để phân biệt ER modeling với UML Class Diagram.

## Tài liệu PlantUML

11. PlantUML. “Class Diagram.”  
    https://plantuml.com/class-diagram  
    Xác minh cú pháp generalization, realization, composition, aggregation, dependency, multiplicity và stereotype.

12. PlantUML. “Sequence Diagram.”  
    https://plantuml.com/sequence-diagram  
    Xác minh participant, message, return, activation và combined fragments.

13. PlantUML. “State Diagram.”  
    https://plantuml.com/state-diagram

14. PlantUML. “Component Diagram.”  
    https://plantuml.com/component-diagram

15. PlantUML. “Deployment Diagram.”  
    https://plantuml.com/deployment-diagram

16. PlantUML. “Information Engineering Diagram.”  
    https://plantuml.com/ie-diagram

## Nguồn nội bộ

17. `report/report_support_documentation/governance/project_source_of_truth.md`
18. `report/report_support_documentation/guidelines/diagrams/uml_2_5_1_drawing_rules.md`
19. `report/report_support_documentation/guidelines/diagrams/diagram_type_rubric.md`
20. `report/report_support_documentation/guidelines/diagrams/diagram_quality_guidelines.md`
21. `report/report_support_documentation/guidelines/diagrams/diagram_review_checklist.md`
22. `report/report_support_documentation/references/standards/formal-17-12-05.pdf`

## Phương pháp xác minh

- Dùng trang đặc tả OMG để xác minh phiên bản, trạng thái và tài liệu normative.
- Dùng bản PDF UML 2.5.1 đã lưu trong repository để đối chiếu ký pháp.
- Dùng trang PlantUML để xác minh cú pháp công cụ; không dùng PlantUML làm nguồn duy nhất cho ngữ nghĩa UML.
- Đánh dấu rõ quy tắc OOSE/ICONIX và quy ước dự án, tránh gán chúng cho OMG UML.
- Với ISO, chỉ sử dụng metadata và phạm vi công khai; không suy diễn nội dung chi tiết của bản tiêu chuẩn trả phí.
