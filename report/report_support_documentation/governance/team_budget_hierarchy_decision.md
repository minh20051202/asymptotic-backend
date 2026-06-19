# Quyết định mô hình phân bổ ngân sách theo đội ngũ

Trạng thái: Đã chấp thuận và tích hợp.

## Quyết định

Asymptotic sử dụng một ví tiền thật ở cấp tổ chức và chuỗi hạn mức:

```text
Organization -> Team -> Developer -> AI Agent
```

- `Wallet` thuộc Organization và lưu tiền thật.
- Team, Developer và Agent chỉ giữ hạn mức kiểm soát, không có ví độc lập.
- Quản trị viên tổ chức phân bổ hoặc thu hồi hạn mức giữa Organization, Team và Developer.
- Developer phân bổ hoặc thu hồi hạn mức của các Agent do mình quản lý.
- Chi phí thực tế luôn được khấu trừ từ ví tổ chức và đồng thời ghi nhận mức sử dụng tại Team, Developer và Agent.

## Quy tắc chính

- Phân bổ cấp dưới không được vượt phần khả dụng của cấp cha.
- `available = allocated - reserved - spent`.
- Không thu hồi phần đã sử dụng, đang tạm giữ hoặc thuộc giao dịch chưa hoàn tất.
- Request chỉ được gọi AI Provider khi ví tổ chức và mọi cấp hạn mức đều cho phép.
- Reservation, settlement, release và reversal phải được cập nhật nguyên tử và có dấu vết kiểm toán.
- Khi chuyển Team, Developer hoặc Agent, hệ thống phải xử lý hạn mức và giao dịch đang xử lý trước khi hoàn tất bàn giao.

## Phạm vi MVP

- Mỗi Developer thuộc tối đa một Team tại một thời điểm.
- Mỗi Agent có một Developer quản lý chính tại một thời điểm.
- Quan hệ nhiều Team hoặc nhiều người quản lý chính là hướng mở rộng.

## Tạo tác đã tích hợp

- `project_source_of_truth.md`: luồng nghiệp vụ, UC03, UC08, FR01, FR05, FR06, FR09 và FR10.
- `chapter_1.tex`: bài toán, mục tiêu và phạm vi.
- `chapter_2.tex`: yêu cầu, actor, UC03, UC08 và các đặc tả liên quan.
- `chapter_3.tex`: Team, TeamMembership, DeveloperProfile, BudgetPolicy và BudgetAllocation.
- `chapter_4.tex`: modular monolith, Finance/Ledger, sequence và thiết kế dữ liệu.

Tài liệu canonical:

- [`project_source_of_truth.md`](project_source_of_truth.md)
- [Đặc tả UC03](../diagrams/chapter_2/use_cases/UC03/specification.md)
- [Đặc tả UC08](../diagrams/chapter_2/use_cases/UC08/specification.md)

