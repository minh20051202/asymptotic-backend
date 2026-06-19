# Cách Viết FR Và NFR Trong OOAD

**Research cutoff:** 2026-06-16  
**Phạm vi:** cách viết Functional Requirement (FR) và Non-Functional Requirement (NFR) trong báo cáo OOAD/phân tích thiết kế hệ thống. Tập trung vào yêu cầu hệ thống, quan hệ với Use Case, Business Rule, Acceptance Criteria, Traceability, và ví dụ áp dụng cho MVP AI Agent Financial Gateway.

## Kết Luận Ngắn

Trong OOAD, **Use Case không thay thế FR/NFR**. Use Case mô tả tương tác giữa actor và hệ thống để đạt một mục tiêu; FR/NFR là các phát biểu yêu cầu có mã định danh, có thể kiểm tra, và có thể truy vết sang thiết kế/test.

Viết đúng theo tầng:

```text
Business Goal / Problem
  -> Business Rule / Constraint
  -> Use Case
  -> Functional Requirement + Non-Functional Requirement / Quality Attribute
  -> Acceptance Criteria / Test Case
  -> Design Class / Component
```

FR và NFR thường được viết song song cho cùng một phạm vi chức năng. FR trả lời: **hệ thống phải làm gì?**  
NFR trả lời: **hệ thống phải làm việc đó tốt đến mức nào, trong điều kiện nào, hoặc bị ràng buộc bởi gì?**

## Nguồn Chuẩn Chính

- ISO/IEC/IEEE 29148:2018 là chuẩn trung tâm về requirements engineering. ISO mô tả chuẩn này là quy định quy trình, nội dung thông tin, và format cho requirements trong vòng đời hệ thống/phần mềm [S01]. IEEE ghi rõ 29148 định nghĩa cấu trúc của yêu cầu tốt, thuộc tính/đặc tính của requirement, và thay thế IEEE 830/1233/1362 ở bản 2011 [S02].
- NASA Systems Engineering Handbook cung cấp checklist viết requirement tốt: dùng `shall` cho requirement, tránh mơ hồ, mỗi requirement một ý, có rationale/assumption, đúng tầng, không lẫn cách triển khai, và phải verifiable/testable [S03].
- OMG UML 2.5.1 xác định UML là ngôn ngữ đồ họa để visualize/specify/construct/document artifacts. Trong OOAD, UML/Use Case là mô hình hóa yêu cầu/hành vi, không phải danh sách requirement thay thế [S04].
- ISO/IEC 25010 cung cấp mô hình quality attributes: functional suitability, performance efficiency, compatibility, interaction capability, reliability, security, maintainability, portability, safety, flexibility [S05][S06].
  Lưu ý: `functional suitability` trong ISO 25010 là **thuộc tính chất lượng của sản phẩm**, không phải một FR riêng lẻ. FR là các hành vi/capability cụ thể mà hệ thống phải cung cấp.
- SEI Quality Attribute Workshop khuyến nghị dùng quality-attribute scenarios để phát hiện và làm rõ yêu cầu chất lượng trước khi kiến trúc được tạo [S07].
- Nghiên cứu về NFR chỉ ra NFR thường bị viết mơ hồ, thiếu định lượng, khó phân tích/test; nhiều NFR thực ra vẫn mô tả hành vi và cần được xử lý/test nghiêm túc như requirement khác [S12][S13].

## FR Là Gì?

Functional Requirement là yêu cầu về **hành vi, chức năng, kết quả hoặc năng lực mà hệ thống phải cung cấp**. Một FR tốt mô tả hành vi quan sát được từ bên ngoài hoặc thay đổi trạng thái nghiệp vụ rõ ràng.

Mẫu viết:

```text
FR-xx: Hệ thống phải [hành động] [đối tượng dữ liệu/nghiệp vụ]
       [khi điều kiện/kích hoạt] [để tạo kết quả].
```

Mẫu tiếng Anh nếu báo cáo dùng song ngữ:

```text
FR-xx: The system shall [verb] [object] [condition] [expected result].
```

Ví dụ tốt:

```text
FR-01: Hệ thống phải xác thực API key của AI Agent trước khi xử lý yêu cầu gọi provider.
FR-02: Hệ thống phải tạo bản ghi reservation khi ngân sách khả dụng của agent đủ chi phí dự kiến.
FR-03: Hệ thống phải từ chối yêu cầu và trả mã 402 khi ngân sách khả dụng nhỏ hơn chi phí dự kiến.
FR-04: Hệ thống phải trả lại kết quả đã lưu khi nhận lại yêu cầu có cùng idempotency key và cùng request hash.
```

Ví dụ chưa tốt:

```text
FR-x: Hệ thống phải nhanh.
```

Sai vì đây là quality/performance, không phải chức năng. Cần chuyển thành NFR có số đo.

```text
FR-y: Hệ thống dùng PostgreSQL và SELECT FOR UPDATE.
```

Sai ở tầng yêu cầu nếu chưa có lý do bắt buộc. Đây là giải pháp thiết kế. Requirement nên là “không cho số dư âm khi có yêu cầu đồng thời”; thiết kế có thể dùng PostgreSQL lock để thỏa mãn.

## NFR Là Gì?

NFR là yêu cầu về **thuộc tính chất lượng, ràng buộc vận hành, ràng buộc thiết kế, môi trường, bảo mật, hiệu năng, độ tin cậy, khả năng bảo trì, khả năng kiểm thử, tuân thủ**. Nên gọi rõ hơn là **Quality Requirement** hoặc **Quality Attribute Requirement** khi viết báo cáo.

Mẫu viết:

```text
NFR-xx: Trong [điều kiện], [chức năng/hệ thống] phải đạt [chỉ số đo được]
        bằng [phương pháp xác minh/test].
```

Mẫu quality-attribute scenario:

```text
Source of stimulus: ai_agent
Stimulus: gửi 50 request đồng thời
Environment: ví có số dư 10 đơn vị, mỗi request giá 1 đơn vị
Artifact: budget reservation service
Response: chỉ tối đa 10 request được reserve thành công
Response measure: số dư không âm; 40 request còn lại trả 402 trong <= X ms
```

Ví dụ tốt:

```text
NFR-01: Khi 50 request đồng thời cùng tiêu thụ một ví có số dư 10 đơn vị và mỗi request có chi phí 1 đơn vị, hệ thống phải đảm bảo không có trạng thái số dư âm; tối đa 10 request được chấp nhận.
NFR-02: Với 95% request proxy không tính thời gian chờ provider, gateway phải xử lý phần kiểm soát ngân sách trong không quá 20 ms ở tải 100 RPS.
NFR-03: API key thô của agent không được lưu trong cơ sở dữ liệu; hệ thống chỉ lưu giá trị hash hoặc dạng đã bảo vệ.
NFR-04: Nếu provider trả lỗi 5xx hoặc timeout, transaction phải ở trạng thái có thể đối soát và reservation phải được hoàn tiền hoặc đánh dấu pending reconciliation.
```

Lưu ý với invariant tài chính: “số dư không âm” có thể được chấm là **Business Rule** hoặc **Functional Invariant** nếu môn học yêu cầu tách FR/NFR nghiêm ngặt. Chỉ nên đặt nó dưới `NFR-CONS` khi bạn đang nhấn mạnh thuộc tính nhất quán dưới tải đồng thời; khi đó vẫn nên trace về BR/FR tương ứng.

Ví dụ chưa tốt:

```text
NFR-x: Hệ thống phải bảo mật.
NFR-y: Hệ thống phải ổn định.
NFR-z: Hệ thống phải dễ dùng.
```

Sai vì không đo được, không biết test bằng gì, và không rõ áp dụng cho chức năng nào.

## FR, NFR, Business Rule Khác Nhau Thế Nào?

| Loại | Câu hỏi | Ví dụ |
|---|---|---|
| Business Goal | Vì sao làm hệ thống? | Giảm rủi ro agent tiêu vượt ngân sách khi gọi API trả phí. |
| Business Rule | Quy tắc nghiệp vụ nào luôn đúng? | Agent không được tiêu vượt hạn mức được cấp. |
| FR | Hệ thống phải làm gì để thực thi rule? | Hệ thống phải kiểm tra ngân sách khả dụng trước khi gọi provider. |
| NFR | Hệ thống phải làm tốt đến mức nào/ràng buộc gì? | Khi có 50 request đồng thời, số dư không được âm và reservation phải nhất quán. |
| Acceptance Criteria | Kiểm chứng bằng gì? | Test 50 request song song, ví 10 đơn vị, chỉ 10 request thành công. |

Business Rule nên được viết riêng rồi FR tham chiếu đến nó. Không nên biến mọi Business Rule thành FR nếu rule đó là chính sách miền nghiệp vụ độc lập với phần mềm. Business Rules Group xem business rules là các quy tắc xác định/cấu trúc/kiểm soát hoạt động của enterprise; OMG SBVR là chuẩn để ghi nhận business vocabularies và business rules giữa tổ chức/công cụ [S11][S14].

## Quan Hệ Với Use Case Trong OOAD

Use Case mô tả **actor đạt mục tiêu gì qua tương tác với hệ thống**. FR được rút ra từ:

- Basic Flow: các hành vi hệ thống phải thực hiện.
- Alternative Flow: lựa chọn hợp lệ khác.
- Exception Flow: cách hệ thống xử lý lỗi.
- Pre/Post-condition: trạng thái bắt buộc trước/sau.
- Business Rule: constraint mà flow phải tuân thủ.

Ví dụ từ UC “Thực hiện yêu cầu AI”:

```text
UC01 step: Gateway kiểm tra khóa lũy đẳng và ghi nhận yêu cầu mới.
-> FR: Hệ thống phải phát hiện yêu cầu gửi lại dựa trên idempotency key và request hash.
-> NFR: Kiểm tra idempotency phải là atomic với việc tạo reservation để không phát sinh charge trùng.
-> AC: Gửi cùng idempotency key 2 lần; hệ thống chỉ tạo 1 transaction.
```

Không nên đưa chi tiết kỹ thuật nhỏ như “gọi hàm X”, “dùng mutex”, “query SQL Y” vào Use Case hoặc FR, trừ khi đó là ràng buộc bắt buộc của bài toán.

## Checklist Viết FR Tốt

Một FR nên có:

- `ID`: duy nhất, ví dụ `FR-01`.
- `Tên ngắn`: “Kiểm tra ngân sách trước thực thi”.
- `Mô tả shall`: hệ thống phải làm gì.
- `Actor/trigger`: khi nào yêu cầu xảy ra.
- `Input/output hoặc state change`: dữ liệu vào/ra hoặc trạng thái bị thay đổi.
- `Business rule liên quan`: ví dụ `BR-01`.
- `Use case liên quan`: ví dụ `UC01`.
- `Priority`: Must/Should/Could.
- `Acceptance criteria`: cách xác minh.

Mẫu bảng này theo tinh thần SRS/Volere: tách requirement, nguồn phát sinh, priority và cách kiểm chứng [S08].

| ID | Tên | Requirement | Source | Priority | Verification |
|---|---|---|---|---|---|
| FR-01 | Kiểm tra ngân sách | Hệ thống phải kiểm tra ngân sách khả dụng trước khi chuyển tiếp yêu cầu đến provider. | UC01, BR-01 | Must | Test request đủ/không đủ ngân sách |

## Checklist Viết NFR Tốt

Một NFR nên có:

- `ID`: ví dụ `NFR-PERF-01`, `NFR-SEC-01`, `NFR-REL-01`.
- `Quality category`: theo ISO 25010 hoặc nhóm riêng của đồ án.
- `Scope`: toàn hệ thống hay chức năng nào.
- `Condition`: tải, môi trường, dữ liệu, trạng thái lỗi.
- `Metric`: p95 latency, throughput, error rate, RTO, lock wait time, số request thành công, số dư âm bằng 0.
- `Threshold`: giá trị cụ thể.
- `Verification method`: test, inspection, analysis, demonstration.

Mẫu bảng NFR nên biến quality attribute thành fit criterion đo được thay vì tính từ chung chung [S07][S08].

| ID | Category | Requirement | Measurement | Verification |
|---|---|---|---|---|
| NFR-CONS-01 | Consistency | Hệ thống phải đảm bảo số dư ví không âm khi xử lý request đồng thời. | Min(balance) >= 0 sau test 50 request song song | Concurrency test |
| NFR-PERF-01 | Performance | Gateway overhead p95 không vượt 20 ms khi tải 100 RPS, không tính thời gian provider. | p95 <= 20 ms | Load test |
| NFR-SEC-01 | Security | API key thô không được lưu ở DB. | Không có plaintext API key trong bảng/log | Inspection + integration test |

## Lỗi Thường Gặp

1. **FR quá to, gom nhiều chức năng.**  
   “Hệ thống quản lý tài chính agent” nên tách thành kiểm tra ngân sách, reserve, settle, refund, report.

2. **NFR không đo được.**  
   “Nhanh”, “bảo mật”, “dễ dùng” phải đổi thành metric và điều kiện test.

3. **Lẫn yêu cầu với thiết kế.**  
   “Dùng PostgreSQL row lock” là thiết kế. Yêu cầu đúng hơn: “không phát sinh double-spend khi request đồng thời”.

4. **Use Case bị phân rã thành các bước kỹ thuật.**  
   Use Case là mục tiêu actor; các bước xác thực, idempotency, reserve, settle nên nằm trong flow hoặc activity/sequence diagram.

5. **Thiếu traceability.**  
   Mỗi FR/NFR cần biết xuất phát từ UC/BR nào và được test bằng gì.

6. **NFR không gắn vào chức năng.**  
   “System availability 99.9%” quá chung nếu đồ án chưa vận hành production. Với đồ án, nên ưu tiên consistency, security, testability, observability, performance có test được.

## Cấu Trúc Khuyến Nghị Cho Báo Cáo OOAD

```text
2.2 Đặc tả yêu cầu hệ thống
  2.2.1 Phạm vi yêu cầu
  2.2.2 Business goals và business rules
  2.2.3 Functional requirements
  2.2.4 Non-functional requirements / quality requirements
  2.2.5 Acceptance criteria
  2.2.6 Traceability matrix

2.3 Phân tích use case
  UC01...
  UC02...
```

Nếu báo cáo theo phong cách OOAD truyền thống, có thể đặt Use Case trước rồi FR/NFR sau. Nhưng phải có bảng traceability để tránh cảm giác Use Case và FR rời nhau.

Ví dụ traceability ngắn cho đồ án:

| Business Rule | Use Case | FR | NFR | Test/AC |
|---|---|---|---|---|
| BR-01: Agent không được tiêu vượt ngân sách | UC01 | FR-EXEC-03 | NFR-CONS-02 | TC-05: 50 request song song, ví 10 đơn vị |
| BR-02: Một idempotency key không tạo hai charge | UC01 | FR-IDEM-03 | NFR-IDEM-01 | TC-06: retry cùng key/hash |

## Áp Dụng Cho AI Agent Financial Gateway

Nên chia FR theo nhóm:

### Nhóm Core Execution

```text
FR-EXEC-01: Hệ thống phải xác thực API key của AI Agent trước khi xử lý yêu cầu.
FR-EXEC-02: Hệ thống phải tra cứu provider, endpoint và pricing hiện hành trước khi gọi provider.
FR-EXEC-03: Hệ thống phải kiểm tra ngân sách khả dụng trước khi gọi provider.
FR-EXEC-04: Hệ thống phải tạo reservation khi ngân sách đủ.
FR-EXEC-05: Hệ thống phải từ chối request bằng 402 khi ngân sách không đủ.
FR-EXEC-06: Hệ thống phải chuyển tiếp request hợp lệ đến provider đã cấu hình.
FR-EXEC-07: Hệ thống phải quyết toán transaction sau khi nhận kết quả provider.
FR-EXEC-08: Hệ thống phải hoàn tiền hoặc đánh dấu đối soát khi provider lỗi/time out.
```

### Nhóm Idempotency

```text
FR-IDEM-01: Hệ thống phải nhận idempotency key từ AI Agent cho mỗi request có khả năng retry.
FR-IDEM-02: Hệ thống phải lưu request hash cùng idempotency key.
FR-IDEM-03: Hệ thống phải trả kết quả đã lưu nếu idempotency key và request hash trùng với request đã xử lý.
FR-IDEM-04: Hệ thống phải từ chối request nếu idempotency key trùng nhưng request hash khác.
```

### Nhóm Administration

```text
FR-ADM-01: Hệ thống phải cho phép quản trị viên tạo/cập nhật/vô hiệu hóa provider.
FR-ADM-02: Hệ thống phải cho phép quản trị viên tạo/cập nhật/vô hiệu hóa endpoint.
FR-ADM-03: Hệ thống phải cho phép quản trị viên tạo pricing policy có thời điểm hiệu lực.
```

### Nhóm Agent Policy

```text
FR-POL-01: Hệ thống phải cho phép lập trình viên thiết lập hạn mức theo agent.
FR-POL-02: Hệ thống phải áp dụng hạn mức agent khi xử lý UC01.
FR-POL-03: Hệ thống phải chặn request mới của agent khi agent bị vô hiệu hóa hoặc vượt hạn mức.
```

### Nhóm Audit/Reporting

```text
FR-AUD-01: Hệ thống phải lưu execution trace cho mỗi request đi qua gateway.
FR-AUD-02: Hệ thống phải lưu ledger entry cho mỗi biến động số dư.
FR-AUD-03: Hệ thống phải cho phép người dùng có quyền xem lịch sử transaction theo phạm vi.
```

NFR phù hợp nhất với đồ án:

```text
NFR-CONS-01: Với mọi transaction tài chính, hệ thống phải đảm bảo số dư khả dụng không âm.
NFR-CONS-02: Khi 50 request đồng thời tiêu thụ cùng một ví có số dư 10 đơn vị, tối đa 10 request được reserve thành công nếu mỗi request có chi phí 1 đơn vị.
NFR-IDEM-01: Retry cùng idempotency key và request hash không được tạo thêm charge hoặc reservation mới.
NFR-PERF-01: Ở tải 100 RPS, p95 overhead của gateway không tính provider không vượt 20 ms.
NFR-SEC-01: API key của agent và provider secret không được lưu plaintext.
NFR-AUD-01: Mỗi transaction phải truy vết được từ request, agent, provider, pricing, reservation, settlement/refund đến ledger entry.
NFR-REC-01: Lỗi provider hoặc lỗi DB không được để transaction mất trạng thái; transaction phải ở một trong các trạng thái có thể đối soát.
```

## Quy Tắc Chấm Nhanh

Một requirement đạt chuẩn khi trả lời được:

1. Ai/cái gì là chủ thể của requirement?
2. Hệ thống phải làm gì hoặc đạt thuộc tính gì?
3. Điều kiện/kích hoạt là gì?
4. Kết quả quan sát được là gì?
5. Có đo/test/inspect được không?
6. Có đúng tầng không, hay đang lẫn thiết kế/code?
7. Có trace được tới nguồn yêu cầu hợp lệ không: stakeholder goal, use case, business rule, regulation, risk, operating constraint hoặc architecture constraint?
8. Có acceptance criteria không?

Nếu trả lời “không” ở câu 5 hoặc 7, requirement chưa đủ mạnh cho báo cáo OOAD.

## Nguồn

[S01] ISO, ISO/IEC/IEEE 29148:2018 Requirements Engineering, https://www.iso.org/standard/72089.html  
[S02] IEEE Standards Association, IEEE/ISO/IEC 29148-2011 page, https://standards.ieee.org/ieee/29148/5289/  
[S03] NASA, Systems Engineering Handbook, Appendix C, https://www.nasa.gov/wp-content/uploads/2018/09/nasa_systems_engineering_handbook_0.pdf  
[S04] OMG, UML 2.5.1 specification page, https://www.omg.org/spec/UML/2.5.1/About-UML  
[S05] ISO, ISO/IEC 25010:2023 Product Quality Model, https://www.iso.org/standard/78176.html  
[S06] ISO 25000 portal, ISO/IEC 25010 quality model summary, https://iso25000.com/index.php/en/iso-25000-standards/iso-25010  
[S07] SEI/CMU, Quality Attribute Workshops, Third Edition, https://insights.sei.cmu.edu/library/quality-attribute-workshops-qaws-third-edition/  
[S08] Volere Requirements Specification Template, https://www.volere.org/templates/volere-requirements-specification-template/  
[S09] Karl Wiegers, Software Requirements Essentials, https://softwarereqs.com/  
[S10] Karl Wiegers books page, https://www.karlwiegers.com/books.html  
[S11] Business Rules Group, "Defining Business Rules: What Are They Really?", July 2000, https://www.businessrulesgroup.org/first_paper/BRG-whatisBR_3ed.pdf  
[S12] Eckhardt, Vogelsang, Mendez Fernandez, "Are Non-functional Requirements really Non-functional?", arXiv, 2016, https://arxiv.org/abs/1611.08868  
[S13] Behutiye et al., "Non-functional Requirements Documentation in Agile Software Development", arXiv, 2017, https://arxiv.org/abs/1711.08894  
[S14] OMG, SBVR 1.5 specification page, https://www.omg.org/spec/SBVR/
