# Nghiên cứu sâu tài liệu đối thủ phục vụ Chương 2

**Ngày nghiên cứu:** 19/06/2026  
**Phạm vi báo cáo:** Mục 2.1 “Khảo sát hiện trạng”  
**Đối tượng nghiên cứu:** AI/LLM Gateway và các hệ thống kiểm soát usage, budget, billing liên quan  
**Nguyên tắc nguồn:** Chỉ dùng tài liệu kỹ thuật, API reference, repository và trang sản phẩm chính thức

## 1. Kết luận chính

Nghiên cứu sâu các nhánh tài liệu Enterprise, architecture, budget, Agent, billing và database làm thay đổi đáng kể nhận định trước đó.

### 1.1. LiteLLM có hierarchical allocation

LiteLLM không chỉ có budget rời rạc. Enterprise documentation mô tả rõ:

```text
Organization
  → Team
    ├─ User
    │   → Virtual Key
    └─ Project
        → Virtual Key
```

Ngoài hierarchy trên, LiteLLM còn có Agent Registry, Agent permission, Agent spend attribution và Agent/session budget.

LiteLLM hỗ trợ:

- budget tại Organization, Team, Team Member/User, Project và Key;
- ràng buộc Team budget không vượt Organization budget;
- ràng buộc User budget không vượt Team budget;
- Team Admin điều chỉnh Team budget trong giới hạn Organization;
- project-level budget và spend isolation;
- chargeback/showback theo Organization, Team, User và Key;
- billing integration với Lago;
- Agent-linked trace và spend attribution.

Vì vậy, “hierarchical allocation” không thể tiếp tục dùng như khác biệt tổng quát của Asymptotic.

### 1.2. Portkey cũng có allocation nhiều cấp

Portkey hỗ trợ năm lớp kiểm soát độc lập:

1. API key;
2. Workspace;
3. Integration/provider;
4. Usage Limit Policy;
5. Rate Limit Policy.

Workspace được mô tả như sub-organization. Portkey cho phép cấp budget riêng cho từng Workspace dùng chung một Integration và đặt Integration-level ceiling phía trên.

### 1.3. OpenRouter đã trở thành đối thủ trực tiếp đáng kể

OpenRouter có:

- Organization;
- Workspace;
- shared organization credits;
- key-level USD limit;
- member/workspace guardrails;
- centralized key management;
- usage tracking;
- provider/model allowlist;
- BYOK credential.

Do đó, OpenRouter nên có mặt trong khảo sát đối thủ trực tiếp của Chương 2.

### 1.4. Cloudflare và Helicone có prepaid credits

Cloudflare Unified Billing và Helicone managed-provider routing đều sử dụng credit balance. Cloudflare hỗ trợ manual top-up và auto-top-up. Vì vậy, “wallet/top-up” ở mức prepaid inference credit cũng không còn là khác biệt độc lập.

### 1.5. Mục tiêu khảo sát không phải chứng minh tính mới tuyệt đối

Theo cách trình bày trong báo cáo mẫu `DOAN1-NGUYENHUYENTRANG.pdf`, khảo sát sản phẩm hiện có nhằm:

1. xác định cách các sản phẩm đang giải quyết bài toán;
2. ghi nhận các chức năng và cách tổ chức tốt để kế thừa;
3. nhận diện hạn chế hoặc điểm chưa phù hợp với phạm vi đề tài;
4. lựa chọn và điều chỉnh các chức năng cho hệ thống đề xuất.

Vì vậy, Asymptotic không cần chứng minh mọi chức năng đều mới. Việc sử dụng lại các mô hình đã được thị trường kiểm chứng như virtual key, hierarchical budget, provider credential isolation, spend tracking và pre-request enforcement là hợp lý.

Phần cần làm rõ là cách Asymptotic lựa chọn và kết hợp các chức năng đó theo phạm vi nghiệp vụ:

```text
organization-owned source of funds
  → allocation to Team
  → allocation to Developer
  → allocation to registered external Agent
  → estimated-cost reservation before provider call
  → actual-usage reconciliation
  → release of unused reservation
  → idempotent financial posting
  → transaction ledger linked to request trace
```

Tổ hợp trên là **mô hình hệ thống đề xuất**, không nhất thiết phải được trình bày như một phát minh hoặc khả năng duy nhất trên thị trường.

### 1.6. Ba loại kết quả cần rút ra từ khảo sát

| Loại kết quả | Ý nghĩa | Ví dụ cho Asymptotic |
|---|---|---|
| Kế thừa | Chức năng đã phổ biến và phù hợp, nên đưa vào hệ thống | Gateway API, virtual key, provider credential isolation, rate limit, hierarchical budget, usage/cost tracking |
| Điều chỉnh | Ý tưởng đã có nhưng cần thay đổi theo actor và nghiệp vụ của đề tài | Biến User/Project budget thành Team–Developer–Agent allocation; gắn Agent với Developer quản lý |
| Tự thiết kế | Cần thiết cho luồng nghiệp vụ nhưng chưa thấy mô tả phù hợp trong sản phẩm khảo sát | Reservation, release, idempotent financial posting, transaction state và recovery |

## 2. Phương pháp nghiên cứu

### 2.1. Quy trình đọc tài liệu

Không dừng ở product landing page. Với mỗi sản phẩm, nghiên cứu đi qua các nhóm trang:

1. Overview và architecture.
2. Enterprise hoặc organization management.
3. Authentication, virtual key và provider credential.
4. Budget, quota và rate limit.
5. Agent registry, Agent permission và Agent cost tracking.
6. Cost estimation, pricing và spend tracking.
7. Billing, credits hoặc wallet.
8. Database, event processing và concurrency.
9. Idempotency, deduplication và audit logs.
10. Feature maturity: OSS, Enterprise, beta, preview hoặc GA.

Các `llms.txt` chính thức được dùng để khám phá các trang con của Portkey, Cloudflare, Helicone và Kong.

### 2.2. Trạng thái capability

| Mã | Ý nghĩa |
|---|---|
| F | Tài liệu mô tả đầy đủ capability theo định nghĩa |
| P | Có chức năng gần tương đương nhưng khác ngữ nghĩa hoặc bảo đảm |
| U | Chưa xác định được từ tài liệu công khai đã đọc |
| X | Tài liệu nói rõ không hỗ trợ |
| NA | Không phù hợp với loại sản phẩm |

`U` không có nghĩa là sản phẩm chắc chắn không có chức năng.

### 2.3. Độ mạnh bằng chứng

| Mã | Loại nguồn |
|---|---|
| E4 | Technical documentation hoặc API reference chính thức |
| E3 | Tài liệu chính thức cho beta/preview |
| E2 | Repository hoặc sample chính thức |
| E1 | Product page, blog hoặc vendor announcement |
| E0 | Chưa tìm thấy bằng chứng công khai đủ dùng |

## 3. Tiêu chí kiểm chứng

| Mã | Tiêu chí | Điều kiện đạt |
|---|---|---|
| C01 | Multi-provider gateway | Một endpoint/gateway gọi được nhiều AI Provider |
| C02 | Gateway-issued credential | Client không cần dùng trực tiếp provider API key |
| C03 | Provider credential isolation | Provider credential được lưu và sử dụng tại gateway |
| C04 | Managed organization hierarchy | Có entity Organization/Workspace/Team/User/Project được quản lý |
| C05 | Registered Agent entity | Agent có ID/registry riêng, không chỉ là metadata |
| C06 | Hierarchical budget allocation | Budget cấp dưới nằm trong hoặc chịu ràng buộc từ budget cấp trên |
| C07 | Agent budget | Có budget/limit gắn trực tiếp với Agent hoặc Agent session |
| C08 | Pre-provider enforcement | Policy được kiểm tra trước khi gọi AI Provider |
| C09 | Estimated-cost reservation | Ước lượng và giữ chi phí trước request để chống concurrent overspend |
| C10 | Actual-usage reconciliation | Điều chỉnh số giữ/chi phí theo usage thực tế |
| C11 | Prepaid source of funds | Có balance/credit được nạp trước và trừ khi sử dụng |
| C12 | Request deduplication | Request gửi lại không tạo provider execution trùng |
| C13 | Idempotent financial posting | Một usage event không ghi chi phí/giao dịch nhiều lần |
| C14 | Financial ledger | Có transaction/entry và trạng thái tài chính, không chỉ request log |
| C15 | Unified financial trace | Liên kết identity, policy, provider call, usage, cost và posting |
| C16 | Chargeback/showback | Phân bổ hoặc báo cáo chi phí theo đơn vị nội bộ |
| C17 | Agent ownership | Agent được gắn rõ với người hoặc đơn vị chịu trách nhiệm |

## 4. LiteLLM: nghiên cứu sâu

### 4.1. Enterprise capability

Trang Enterprise xác nhận LiteLLM OSS đã có:

- OpenAI-compatible gateway;
- virtual key;
- spend tracking;
- budget;
- fallback;
- request/response logging.

Enterprise bổ sung:

- Organization và Organization Admin;
- delegated Team Admin;
- project management;
- fine-grained RBAC;
- audit log;
- key rotation;
- secret manager;
- project/tag/model-specific budget;
- temporary budget increase;
- spend report;
- multi-region control plane.

Nguồn: [LL01].

### 4.2. Hierarchical allocation

Multi-Tenant Architecture mô tả:

```text
Organization → Team → User → Key
```

Tài liệu ghi rõ:

- Organization có budget, model access, admin, Team và spend tracking.
- Team kế thừa constraint từ Organization.
- Team không được vượt Organization budget/model constraint.
- User có budget riêng.
- User budget không được vượt Team budget.
- Request bị chặn nếu bất kỳ cấp nào vượt budget.
- Team Admin có thể cập nhật Team budget trong Organization limit.
- Mỗi API call được gắn Organization, Team, User và Key context.

Ví dụ chính thức:

```text
Organization: $10,000/month
├─ Team 1: $6,000/month
│  ├─ User A: $3,000/month
│  └─ User B: $3,000/month
└─ Team 2: $4,000/month
```

Nguồn: [LL02].

Kết luận: LiteLLM đạt `C06 Hierarchical budget allocation`.

### 4.3. Project-level allocation

Project Management, hiện được đánh dấu beta và Enterprise, đặt Project giữa Team và Key:

```text
Organization → Team → Project → Key
```

Project có:

- Team owner;
- model allowlist;
- `max_budget`;
- TPM/RPM;
- budget duration;
- blocked status;
- metadata/cost center;
- project spend và model spend;
- key riêng.

Nguồn: [LL03].

Điều này cho phép tách nhiều ứng dụng/use case bên trong một Team. Project có thể gần với AI Agent hoặc workload về mặt cost center, nhưng Project không đồng nghĩa với registered external Agent.

### 4.4. Team Member budget

LiteLLM hỗ trợ `max_budget_in_team`. Khi User được thêm vào Team, Team Admin có thể đặt mức chi tối đa của User trong Team. Key gắn cả `user_id` và `team_id`; spend của key cập nhật spend User và Team.

Nếu key có `team_id`, Team budget được dùng thay cho personal budget. Để giới hạn User trong Team, phải dùng Team Member budget.

Nguồn: [LL04].

### 4.5. Agent Registry và Agent access

A2A Agent Gateway cho phép đăng ký:

- A2A Agent;
- Azure AI Foundry Agent;
- Vertex AI Agent Engine;
- Bedrock AgentCore Agent;
- LangGraph Agent;
- Pydantic AI Agent.

Agent Permission Management cho phép giới hạn Agent được gọi bởi:

- Virtual Key;
- Team;
- access group.

Request bị từ chối 403 nếu key/Team không có quyền gọi Agent.

Nguồn: [LL05], [LL06].

### 4.6. Agent spend attribution và trace

LiteLLM gửi:

- `X-LiteLLM-Trace-Id` để nhóm các call trong cùng Agent execution;
- `X-LiteLLM-Agent-Id` để gắn spend với Agent.

Agent có thể yêu cầu trace ID cho request vào/ra. Request thiếu trace ID bị từ chối nếu cấu hình enforcement được bật.

Nguồn: [LL05].

### 4.7. Agent/session budget

Tài liệu budget chung ghi rõ Agent có:

- TPM/RPM;
- iteration cap;
- session dollar budget.

Agent session budget kiểm tra accumulated spend trước mỗi call. Cost được ghi nhận sau call thành công.

Nguồn: [LL04], [LL07].

Đây là pre-call enforcement trên accumulated spend, nhưng chưa phải reservation của estimated cost cho request đang bắt đầu.

### 4.8. Spend tracking và database

LiteLLM tính cost từ model pricing và ghi `LiteLLM_SpendLogs` với:

- API key hash;
- internal User;
- Team;
- request tags;
- end user;
- model;
- provider endpoint;
- spend;
- token usage.

Database có các bảng riêng cho:

- Organization;
- Team;
- User;
- membership;
- Virtual Key;
- Budget;
- Spend Logs;
- Audit Logs.

Nguồn: [LL08], [LL09].

`SpendLogs` là request-level usage/cost log. Đây là bằng chứng mạnh cho accounting và chargeback, nhưng không tự chứng minh double-entry financial ledger.

### 4.9. Billing

LiteLLM tích hợp Lago để bill:

- Team;
- internal User;
- external customer.

Event gửi sang Lago có `transaction_id`, customer ID, input/output token, model và calculated response cost.

Nguồn: [LL10].

Đây là usage-based billing integration, không phải organization prepaid wallet được LiteLLM quản lý native.

### 4.10. Cost estimation

`/cost/estimate` dự báo cost từ:

- model;
- expected input token;
- expected output token;
- request volume.

Tính năng này phục vụ planning/model comparison. Tài liệu không cho thấy kết quả `/cost/estimate` được dùng tự động để reserve budget cho request runtime.

Nguồn: [LL11].

### 4.11. Request lifecycle và consistency

Life of a Request mô tả:

1. kiểm tra key và budget;
2. kiểm tra rate limit;
3. gửi request qua router tới provider;
4. sau response, chạy background task;
5. `_ProxyDBLogger` cập nhật spend/usage async.

Tài liệu nói rõ database transaction không gắn với lifecycle của request; ngoài key lookup, các DB transaction khác chạy bất đồng bộ.

Nguồn: [LL12].

High Availability documentation cho biết các instance tích lũy spend update vào Redis queue, sau đó một instance giữ lock và flush aggregate update vào database.

Nguồn: [LL13].

**Suy luận có căn cứ:** Cơ chế trên không thể hiện strict estimated-cost reservation trước provider call. Vì spend update diễn ra sau call và bất đồng bộ, không nên khẳng định LiteLLM có cùng consistency model với reservation/settlement của Asymptotic.

### 4.12. Đánh giá LiteLLM

| Tiêu chí | Đánh giá |
|---|---|
| C01–C04 | F/E4 |
| C05 Registered Agent | F/E4 |
| C06 Hierarchical allocation | F/E4 |
| C07 Agent budget | F/E4 |
| C08 Pre-provider enforcement | F/E4 |
| C09 Estimated reservation | U/E0 |
| C10 Actual reconciliation | P/E4 — post-call spend accounting, không thấy hold/release |
| C11 Prepaid funds | U/E0 |
| C12 Request deduplication | U/E0 |
| C13 Financial posting idempotency | U/E0 |
| C14 Financial ledger | P/E4 — spend logs và billing events |
| C15 Unified trace | P/E4 |
| C16 Chargeback/showback | F/E4 |
| C17 Agent ownership | P/E4 — Agent attribution có, developer ownership chưa nối rõ |

### 4.13. Kết luận riêng cho LiteLLM

Phát biểu đúng:

> LiteLLM hỗ trợ hierarchical budget allocation và enforcement tại Organization, Team, User/Team Member, Project và Key; đồng thời hỗ trợ Agent registry, Agent/session budget và Agent spend attribution.

Phát biểu còn dùng được cho Asymptotic:

> Tài liệu LiteLLM chưa thể hiện một organization-owned prepaid wallet được phân bổ tới Developer và Agent, strict estimated-cost reservation trước provider call, release phần dư theo actual usage, request deduplication và idempotent wallet-ledger posting.

## 5. Portkey: nghiên cứu sâu

### 5.1. Organization và Workspace

Portkey Organization là container cao nhất cho User, entity và Workspace. Workspace là sub-organization dùng để tách Team/project, data, visibility và resource.

Workspace có:

- Manager/Member role;
- service-account key;
- user key;
- scoped API;
- provider, prompt, config và guardrail riêng.

Nguồn: [PK01], [PK02].

### 5.2. Provider credential governance

Integration lưu credential một lần ở Organization rồi provision tới nhiều Workspace. API key của provider:

- được mã hóa;
- không lộ cho end user;
- có model allowlist;
- có budget/rate limit riêng theo Workspace;
- có thể revoke Workspace access.

Nguồn: [PK03].

### 5.3. Năm lớp limit

Portkey mô tả năm lớp:

| Lớp | Budget |
|---|---|
| API Key | Cost/token/rate |
| Workspace | Cost/token/rate |
| Integration/provider | Shared ceiling và per-workspace sub-limit |
| Usage Policy | Cost/token/request, group-by metadata |
| Rate Policy | Request/token rolling window |

Mỗi request đi qua mọi applicable check trước provider. Request bị budget/token cap trả 412 và không phát sinh provider spend.

Nguồn: [PK04].

Portkey do đó đạt hierarchical allocation theo mô hình:

```text
Integration hard ceiling
  → per-Workspace allocation
    → API Key limit
```

Ngoài ra policy metadata có thể tạo per-user hoặc per-customer counter.

### 5.4. Agent Registry

Portkey Agent Registry cho phép Admin provision Agent và kiểm soát access tới Agent/skill.

Tại thời điểm khảo sát, tài liệu đánh dấu “coming soon” cho:

- access theo User cụ thể;
- skill access trong Workspace cụ thể;
- skill access theo User.

Nguồn: [PK05].

Agent Registry tồn tại, nhưng Agent chưa được thể hiện như nút cuối của financial hierarchy từ Organization → Workspace → Developer → Agent.

### 5.5. Pricing caveat

Cost được tính theo input/output token và pricing data. Tuy nhiên:

- model không có pricing support có cost bằng 0;
- request đó không được tính vào provider budget;
- streaming request cần provider trả usage hoặc bật `include_usage`.

Nguồn: [PK06].

Đây là limitation quan trọng đối với budget enforcement dựa trên tiền.

### 5.6. Đánh giá Portkey

| Tiêu chí | Đánh giá |
|---|---|
| C01–C04 | F/E4 |
| C05 Registered Agent | F/E4 |
| C06 Hierarchical allocation | F/E4 |
| C07 Agent budget | U/E0 |
| C08 Pre-provider enforcement | F/E4 |
| C09 Estimated reservation | U/E0 |
| C10 Actual reconciliation | P/E4 — cost accounting, không thấy reserve/release |
| C11 Prepaid funds | U/E0 |
| C12–C13 Idempotency | U/E0 |
| C14 Financial ledger | U/E0 |
| C15 Unified trace | P/E4 |
| C16 Chargeback/showback | F/E4 |
| C17 Agent ownership | P/E4 |

## 6. Cloudflare AI Gateway: nghiên cứu sâu

### 6.1. Unified Billing và credits

Cloudflare Unified Billing:

- yêu cầu mua credit;
- hiển thị account credit balance;
- hỗ trợ manual top-up;
- hỗ trợ threshold-based auto-top-up;
- tự động trừ credit;
- gọi third-party provider không cần provider API key;
- thu phí 5% trên credit purchase.

Nguồn: [CF01].

Cloudflare đạt prepaid source of funds ở account level.

### 6.2. Spend limits

Spend limit:

- tính bằng USD;
- fixed hoặc rolling window;
- scope theo model/provider/custom metadata;
- split budget bucket theo `user_id`, Team hoặc application;
- kiểm tra tất cả rule trước provider dispatch;
- block 429 hoặc fallback sang model rẻ hơn.

Nguồn: [CF02].

Tuy nhiên, tài liệu nói rõ:

- current request cost chỉ ghi sau khi hoàn tất;
- rule eventually consistent;
- burst concurrent request có thể vượt budget tạm thời;
- cost là best-effort estimate, provider bill mới là authoritative.

Vì vậy, Cloudflare có pre-request blocking nhưng không có strict reservation.

### 6.3. Authentication limitation

Cloudflare AI Gateway token ở account scope. Quyền `AI Gateway Run` không thể giới hạn vào một gateway; token có thể gọi mọi gateway trong account và sử dụng stored provider credential.

Nguồn: [CF03].

Đây là khác biệt đáng kể so với per-Agent credential ownership.

### 6.4. Metadata không phải hierarchy

Custom metadata có thể tạo budget bucket theo User/Team/application nhưng các giá trị này không được tài liệu hóa như managed entity có:

- owner;
- membership;
- delegated admin;
- allocation transaction;
- transfer/revoke workflow.

Do đó Cloudflare đạt segmentation, không đạt managed hierarchical allocation theo nghĩa LiteLLM/Portkey.

### 6.5. Đánh giá Cloudflare

| Tiêu chí | Đánh giá |
|---|---|
| C01–C03 | F/E4 hoặc E3 với BYOK beta |
| C04 Managed hierarchy | P/E4 |
| C05 Registered Agent | U/E0 |
| C06 Hierarchical allocation | P/E3 — metadata buckets, không managed hierarchy |
| C07 Agent budget | U/E0 |
| C08 Pre-provider enforcement | F/E3 |
| C09 Estimated reservation | X/E4 — eventual consistency được nêu rõ |
| C10 Actual reconciliation | P/E4 |
| C11 Prepaid funds | F/E4 |
| C12–C14 | U/E0 |
| C15 Unified trace | P/E4 |
| C16 Chargeback/showback | P/E4 |
| C17 Agent ownership | U/E0 |

## 7. Helicone: nghiên cứu sâu

### 7.1. AI Gateway và managed credits

Helicone hiện cung cấp:

- OpenAI-compatible AI Gateway;
- hơn 100 model/provider;
- Helicone credential;
- BYOK;
- Helicone-managed provider key;
- credit balance;
- BYOK priority;
- managed-key fallback;
- cost-based routing;
- retry và cache.

Khi không đủ Helicone credit và không có BYOK, gateway trả 429. Tài liệu cũng mô tả trạng thái “wallet suspended”.

Nguồn: [HC01], [HC02].

### 7.2. Cost-based limiting

Custom Rate Limits hỗ trợ:

- request count;
- cost theo cent;
- global;
- per User;
- per custom property như Organization.

Request bị limit tại gateway và trả 429.

Tại thời điểm khảo sát, token-based limit và nhiều policy đồng thời được đánh dấu “coming soon”.

Nguồn: [HC03].

### 7.3. Cost attribution

Helicone sử dụng:

- Session để nhóm nhiều request trong một workflow;
- custom property để phân đoạn User tier, feature, environment;
- request/session cost;
- report, alert và analytics.

Nguồn: [HC04].

Đây là observability mạnh, nhưng User/Organization custom property không phải native allocation hierarchy.

### 7.4. Đánh giá Helicone

| Tiêu chí | Đánh giá |
|---|---|
| C01–C03 | F/E4 |
| C04 Managed hierarchy | P/E4 |
| C05 Registered Agent | U/E0 |
| C06 Hierarchical allocation | P/E4 — custom property, không managed allocation |
| C07 Agent budget | U/E0 |
| C08 Pre-provider enforcement | P/E4 |
| C09–C10 Reservation/reconciliation | U/E0 |
| C11 Prepaid funds | F/E4 |
| C12–C14 | U/E0 |
| C15 Unified trace | P/E4 |
| C16 Chargeback/showback | P/E4 |
| C17 Agent ownership | U/E0 |

## 8. OpenRouter: đối thủ cần bổ sung

### 8.1. Organization, Workspace và shared credits

OpenRouter Organization hỗ trợ:

- shared credits;
- centralized API key management;
- usage tracking giữa các member;
- Workspace;
- member role.

Mỗi API key thuộc một Workspace. Workspace có guardrail riêng và kế thừa account-level policy.

Nguồn: [OR01], [OR02].

### 8.2. USD limits

API key có:

- USD credit limit;
- daily/weekly/monthly reset;
- remaining credits;
- usage;
- option tính BYOK usage vào limit;
- automatic disable khi vượt limit.

Nguồn: [OR03], [OR04].

Enterprise quickstart mô tả guardrail cho member và key, gồm:

- spend limit;
- provider/model allowlist;
- Zero Data Retention requirement.

Nguồn: [OR05].

### 8.3. Provider credential

OpenRouter hỗ trợ:

- OpenRouter-managed credits;
- BYOK;
- encrypted provider keys;
- Workspace BYOK settings.

Nguồn: [OR06].

### 8.4. Đánh giá OpenRouter

| Tiêu chí | Đánh giá |
|---|---|
| C01–C04 | F/E4 |
| C05 Registered Agent | U/E0 |
| C06 Hierarchical allocation | P/E4 — Organization/Workspace/member/key controls |
| C07 Agent budget | U/E0 |
| C08 Pre-provider enforcement | F/E4 tại key/guardrail level |
| C09–C10 | U/E0 |
| C11 Prepaid funds | F/E4 |
| C12–C14 | U/E0 |
| C15–C16 | P/E4 |
| C17 Agent ownership | U/E0 |

## 9. Azure API Management AI Gateway

### 9.1. Capability

Azure API Management AI Gateway hỗ trợ:

- OpenAI-compatible và passthrough LLM API;
- unified multi-provider model API ở trạng thái preview;
- Amazon Bedrock và provider ngoài Microsoft;
- self-hosted model;
- MCP server;
- A2A Agent API import;
- managed identity/provider authentication;
- load balancing và circuit breaker;
- token metric, logging và audit.

Nguồn: [AZ01].

### 9.2. Hierarchical quota

`llm-token-limit` dùng arbitrary `counter-key`, ví dụ:

- subscription;
- client IP;
- User;
- application;
- Team;
- department.

Policy hỗ trợ:

- TPM;
- hourly/daily/weekly/monthly/yearly token quota;
- estimated prompt token trước backend;
- actual completion token sau response.

Nguồn: [AZ02].

Đây là hierarchical/segmented token allocation, không phải monetary wallet allocation.

### 9.3. Đánh giá Azure

Azure là comparator mạnh cho Agent/API governance và token quota, nhưng tài liệu được đọc chưa thể hiện:

- organization prepaid AI wallet;
- USD budget hierarchy;
- financial reservation;
- usage-cost transaction ledger.

## 10. Kong AI Gateway và Metering & Billing

### 10.1. AI Gateway

Kong hỗ trợ:

- nhiều provider;
- AI Proxy Advanced;
- load balancing/fallback/retry;
- Consumer/Consumer Group;
- token/cost rate limiting;
- A2A traffic gateway;
- logging và OpenTelemetry.

Nguồn: [KG01], [KG02].

### 10.2. Cost limit

AI Rate Limiting Advanced dùng token data từ provider để tính query cost và giới hạn cost của Consumer theo time window. Policy có thể scope theo:

- IP;
- credential;
- Consumer;
- service;
- header;
- path;
- Consumer Group.

Nguồn: [KG02], [KG03].

### 10.3. Metering & Billing

Konnect Metering & Billing cung cấp:

- real-time event metering;
- AI token metering;
- event deduplication;
- usage attribution;
- plan/subscription;
- limit enforcement;
- invoice;
- payment-provider/ERP integration.

Nguồn: [KG04], [KG05], [KG06].

Kong là trường hợp quan trọng vì có deduplication và billing event. Tuy nhiên, Metering & Billing chủ yếu phục vụ monetization/billing customer usage, không được tài liệu hóa như pre-provider reservation trên organization AI spend.

## 11. Ma trận cập nhật

| Capability | LiteLLM | Portkey | Cloudflare | Helicone | OpenRouter | Azure APIM | Kong |
|---|---|---|---|---|---|---|---|
| Multi-provider gateway | F | F | F | F | F | F | F |
| Managed hierarchy | F | F | P | P | F | F/P | F |
| Hierarchical allocation | F | F | P | P | P/F | P token | P |
| Registered external Agent | F | F | U | U | U | F A2A API | F A2A traffic |
| Agent budget | F | U | U | U | U | P token policy | U |
| Pre-provider budget check | F | F | F | P | F | F token | P |
| Strict estimated reservation | U | U | X | U | U | U | U |
| Actual-usage reconciliation | P | P | P | P | P | P token | P |
| Prepaid credits | U | U | F | F | F | U | NA/P |
| Request deduplication | U | U | U | U | U | U | U |
| Financial event deduplication | U | U | U | U | U | U | F in Metering |
| Financial ledger | P spend logs | U | U | U | U | U | P billing events |
| Chargeback/showback | F | F | P | P | F/P | P | F |
| Agent ownership chain | P | P | U | U | U | P | U |

## 12. Ma trận kế thừa chức năng cho Asymptotic

Ma trận này quan trọng hơn một bảng chỉ tập trung vào “có/không”. Nó cho biết kết quả khảo sát được sử dụng thế nào trong hệ thống đề xuất.

| Chức năng khảo sát | Sản phẩm tham khảo chính | Hướng áp dụng cho Asymptotic |
|---|---|---|
| Unified multi-provider gateway | LiteLLM, Portkey, Cloudflare, Helicone, OpenRouter, Azure, Kong | Kế thừa mô hình một gateway che giấu khác biệt provider |
| Gateway-issued/virtual key | LiteLLM, Portkey, OpenRouter | Kế thừa; cấp key riêng cho từng registered AI Agent |
| Provider credential isolation | Tất cả gateway chính | Kế thừa; credential chỉ do System Admin/Gateway quản lý |
| Organization hierarchy | LiteLLM, Portkey, OpenRouter | Kế thừa nguyên tắc multi-tenant và delegated administration |
| Hierarchical budget | LiteLLM, Portkey | Kế thừa; ánh xạ thành Organization → Team → Developer → Agent |
| Team member budget | LiteLLM | Điều chỉnh thành Developer allocation trong Team |
| Project budget | LiteLLM | Tham khảo cho workload/Agent allocation nhưng không đồng nhất Project với Agent |
| Agent Registry | LiteLLM, Portkey, Azure | Kế thừa khái niệm đăng ký external Agent; bổ sung Developer ownership và handover |
| Agent permission | LiteLLM, Portkey | Kế thừa; key chỉ được gọi model/provider hoặc endpoint cho phép |
| Agent spend attribution | LiteLLM | Kế thừa; mọi request phải có `agent_id` và cost owner path |
| Pre-request budget enforcement | LiteLLM, Portkey, Cloudflare, OpenRouter | Kế thừa; request bị từ chối trước provider nếu không đủ hạn mức |
| Cost estimation | LiteLLM và pricing catalog của các gateway | Kế thừa công thức ước lượng; dùng trực tiếp trong runtime authorization |
| Spend/cost log | LiteLLM, Portkey, Cloudflare, Helicone | Kế thừa; mở rộng thành trace liên kết request và transaction |
| Prepaid credits | Cloudflare, Helicone, OpenRouter | Tham khảo cho ví tổ chức và top-up prototype |
| Chargeback/showback | LiteLLM, Portkey, Kong | Kế thừa cách tổng hợp chi phí theo đơn vị nội bộ |
| Billing integration | LiteLLM–Lago, Kong Metering & Billing | Tham khảo cho reporting và future integration; không cần triển khai invoice production trong MVP |
| Reservation và release | Chưa thấy mô hình tương đương rõ trong các gateway chính | Tự thiết kế để bảo vệ hạn mức khi request đồng thời |
| Request idempotency | Chưa thấy contract phù hợp trong tài liệu gateway đã đọc | Tự thiết kế để tránh provider call trùng do retry |
| Ledger posting idempotency | Kong có event dedup ở billing layer nhưng khác phạm vi | Điều chỉnh/tự thiết kế cho transaction nội bộ |
| Financial recovery state | Chưa thấy đầy đủ trong tài liệu khảo sát | Tự thiết kế cho timeout, provider success/DB failure và streaming interruption |

## 13. Ý nghĩa đối với FR/NFR của Asymptotic

### FR01 và FR05

LiteLLM và Portkey cho thấy hierarchical budget là một mô hình thực tế nên được kế thừa. Asymptotic điều chỉnh hierarchy theo actor chuẩn:

- một nguồn tiền tại Organization;
- Team, Developer và Agent là allocation, không phải ví độc lập;
- allocation/revoke được ghi như transaction;
- Developer chỉ cấp Agent trong hạn mức nhận được.

### FR02

External Agent registration đã xuất hiện tại LiteLLM, Portkey và Azure. Đây là bằng chứng cho thấy chức năng phù hợp với xu hướng sản phẩm và nên được kế thừa.

Asymptotic điều chỉnh theo nghiệp vụ:

- Agent phải gắn với Developer quản lý;
- bàn giao Agent khi thay đổi Team/Developer;
- Agent tham gia trực tiếp vào financial ownership path.

### FR03 và FR04

Gateway key và provider credential isolation đã phổ biến. Đây là các thực hành tốt cần kế thừa, không cần được trình bày như tính năng mới.

### FR06

Pre-provider blocking cũng đã phổ biến. FR06 nên được làm rõ theo mức bảo đảm:

1. estimate request cost;
2. atomically reserve against all relevant limits;
3. reject before provider nếu reservation thất bại;
4. tránh concurrent overspend.

### FR08

Request idempotency vẫn là điểm khác biệt khả thi. Tài liệu đối thủ được đọc chưa mô tả rõ một idempotency key bảo đảm không tạo provider call và financial posting trùng.

### FR09 và NFR01

Request/cost log đã phổ biến. Asymptotic phải phân biệt log với ledger:

- immutable transaction identity;
- allocation/reservation/charge/release entry;
- trạng thái pending/committed/reversed;
- idempotent posting;
- balance invariant.

### NFR04

LiteLLM cho thấy spend update có thể chạy async. Asymptotic nên mô hình hóa rõ failure recovery khi:

- provider thành công nhưng database update lỗi;
- client timeout nhưng provider đã xử lý;
- streaming dừng giữa chừng;
- settlement/reconciliation retry;
- duplicate callback hoặc duplicate request.

## 14. Cách tổng hợp kết quả khảo sát cho Asymptotic

### 14.1. Không nên dùng

- “Sản phẩm đầu tiên có hierarchical budget.”
- “Các AI Gateway chưa có Agent registration.”
- “Đối thủ chỉ theo dõi chi phí sau thực thi.”
- “Chỉ Asymptotic có wallet/top-up.”
- “LiteLLM không có phân bổ ngân sách.”

### 14.2. Phát biểu đề xuất

> Kết quả khảo sát cho thấy các AI Gateway hiện có đã áp dụng hiệu quả nhiều chức năng như quản trị đa nhà cung cấp, virtual key, Agent registry, budget nhiều cấp, pre-request enforcement, prepaid credit và usage-based billing. Asymptotic kế thừa các nguyên tắc này và điều chỉnh theo cơ cấu Organization–Team–Developer–AI Agent. Hệ thống đồng thời bổ sung các cơ chế cần thiết cho luồng nghiệp vụ của đề tài, gồm reservation theo chi phí dự kiến, đối soát theo usage thực tế, xử lý lũy đẳng và truy vết tài chính theo từng request.

### 14.3. Kết luận theo đúng phong cách báo cáo mẫu

Sau mỗi sản phẩm nên trình bày:

- chức năng nổi bật;
- ưu điểm có thể học hỏi;
- hạn chế hoặc điểm không phù hợp;
- chức năng Asymptotic lựa chọn kế thừa.

Cuối phần khảo sát không cần kết luận “thị trường chưa có”. Có thể kết luận:

> Qua khảo sát các sản phẩm hiện có, nhóm thực hiện lựa chọn mô hình Gateway đa nhà cung cấp, khóa truy cập trung gian, quản lý credential tập trung, ngân sách phân cấp, kiểm soát trước request và theo dõi usage/cost làm nền tảng cho hệ thống đề xuất. Các chức năng được điều chỉnh theo actor và luồng nghiệp vụ của Asymptotic; đồng thời bổ sung reservation, lũy đẳng và ledger để bảo đảm tính nhất quán tài chính.

### 14.4. Mức độ chắc chắn

- “Các capability đã phổ biến”: bằng chứng mạnh.
- “Chưa thấy complete combination”: bằng chứng vừa; chỉ áp dụng cho tài liệu công khai đã khảo sát.
- “Asymptotic duy nhất”: không có đủ bằng chứng để tuyên bố.

## 15. Cấu trúc đề xuất cho mục 2.1

### 2.1.1. Bối cảnh

- Agent có request loop, retry và concurrency.
- Provider pricing phụ thuộc actual usage.
- Provider key trực tiếp làm phân tán credential và cost attribution.

### 2.1.2. Phương pháp khảo sát

- mốc thời gian;
- tập sản phẩm;
- taxonomy F/P/U/X;
- official-source-only;
- limitation.

### 2.1.3. Đối thủ trực tiếp

Nên chọn tối đa 5 sản phẩm để giữ cân bằng chương:

1. LiteLLM;
2. Portkey;
3. Cloudflare AI Gateway;
4. OpenRouter;
5. Kong hoặc Azure API Management.

Helicone có thể xuất hiện ngắn hơn hoặc trong bảng vì overlap với Cloudflare/OpenRouter.

### 2.1.4. Bảng so sánh

Sử dụng hai bảng:

**Bảng 1 — Chức năng của sản phẩm**

- managed hierarchy;
- Agent registration;
- hierarchical allocation;
- monetary pre-request enforcement;
- strict reservation;
- actual reconciliation;
- prepaid funds;
- request idempotency;
- ledger idempotency;
- financial ledger;
- Agent ownership chain.

**Bảng 2 — Hướng áp dụng cho Asymptotic**

- kế thừa;
- điều chỉnh;
- chưa đưa vào MVP;
- tự thiết kế.

Routing/cache/retry vẫn có thể nêu như chức năng được kế thừa, nhưng không cần phân tích dài vì đã là capability phổ biến.

### 2.1.5. Tổng hợp và đề xuất hệ thống

Tách:

- capability phổ biến;
- chức năng được kế thừa trực tiếp;
- chức năng cần điều chỉnh theo actor và scope;
- chức năng cần tự thiết kế;
- liên kết các lựa chọn với FR/NFR.

## 16. Nguồn chính thức

### LiteLLM

- [LL01] Enterprise: <https://docs.litellm.ai/docs/enterprise>
- [LL02] Multi-Tenant Architecture: <https://docs.litellm.ai/docs/proxy/multi_tenant_architecture>
- [LL03] Project Management: <https://docs.litellm.ai/docs/proxy/project_management>
- [LL04] Budgets and Rate Limits: <https://docs.litellm.ai/docs/proxy/users>
- [LL05] A2A Agent Gateway: <https://docs.litellm.ai/docs/a2a>
- [LL06] Agent Permission Management: <https://docs.litellm.ai/docs/a2a_agent_permissions>
- [LL07] Agent Iteration Budgets: <https://docs.litellm.ai/docs/a2a_iteration_budgets>
- [LL08] Spend Tracking: <https://docs.litellm.ai/docs/proxy/cost_tracking>
- [LL09] Database Information: <https://docs.litellm.ai/docs/proxy/db_info>
- [LL10] Billing with Lago: <https://docs.litellm.ai/docs/proxy/billing>
- [LL11] Pricing Calculator: <https://docs.litellm.ai/docs/proxy/pricing_calculator>
- [LL12] Life of a Request: <https://docs.litellm.ai/docs/proxy/architecture>
- [LL13] High Availability Spend Updates: <https://docs.litellm.ai/docs/proxy/db_deadlocks>

### Portkey

- [PK01] Organizations: <https://portkey.ai/docs/product/enterprise-offering/org-management/organizations>
- [PK02] Workspaces: <https://portkey.ai/docs/product/enterprise-offering/org-management/workspaces>
- [PK03] Integrations: <https://portkey.ai/docs/product/model-catalog/integrations>
- [PK04] Enforcing Limits and Budgets: <https://portkey.ai/docs/guides/use-cases/enforcing-limits-and-budgets>
- [PK05] Agent Registry: <https://portkey.ai/docs/product/agent-gateway/registry>
- [PK06] Cost Management: <https://portkey.ai/docs/product/observability/cost-management>

### Cloudflare

- [CF01] Unified Billing: <https://developers.cloudflare.com/ai-gateway/features/unified-billing/>
- [CF02] Spend Limits: <https://developers.cloudflare.com/ai-gateway/features/spend-limits/>
- [CF03] Authenticated Gateway: <https://developers.cloudflare.com/ai-gateway/configuration/authentication/>
- [CF04] BYOK: <https://developers.cloudflare.com/ai-gateway/configuration/bring-your-own-keys/>
- [CF05] Custom Metadata: <https://developers.cloudflare.com/ai-gateway/observability/custom-metadata/>
- [CF06] Logging: <https://developers.cloudflare.com/ai-gateway/observability/logging/>

### Helicone

- [HC01] AI Gateway: <https://docs.helicone.ai/gateway/overview>
- [HC02] Provider Routing: <https://docs.helicone.ai/gateway/provider-routing>
- [HC03] Custom Rate Limits: <https://docs.helicone.ai/features/advanced-usage/custom-rate-limits>
- [HC04] Cost Tracking: <https://docs.helicone.ai/guides/cookbooks/cost-tracking>
- [HC05] Error Handling and Credits: <https://docs.helicone.ai/gateway/concepts/error-handling>

### OpenRouter

- [OR01] Organization Management: <https://openrouter.ai/docs/cookbook/administration/organization-management>
- [OR02] Workspaces: <https://openrouter.ai/docs/guides/features/workspaces>
- [OR03] Management API Keys: <https://openrouter.ai/docs/guides/overview/auth/management-api-keys>
- [OR04] Key Limits: <https://openrouter.ai/docs/api/reference/limits>
- [OR05] Enterprise Quickstart: <https://openrouter.ai/docs/cookbook/get-started/enterprise-quickstart>
- [OR06] BYOK: <https://openrouter.ai/docs/guides/overview/auth/byok>

### Azure API Management

- [AZ01] AI Gateway Capabilities: <https://learn.microsoft.com/en-us/azure/api-management/genai-gateway-capabilities>
- [AZ02] LLM Token Limit Policy: <https://learn.microsoft.com/en-us/azure/api-management/llm-token-limit-policy>

### Kong

- [KG01] Kong AI Gateway: <https://developer.konghq.com/ai-gateway/>
- [KG02] AI Rate Limiting Advanced: <https://developer.konghq.com/plugins/ai-rate-limiting-advanced/>
- [KG03] AI Rate Limiting Reference: <https://developer.konghq.com/plugins/ai-rate-limiting-advanced/reference/>
- [KG04] Metering and Billing: <https://developer.konghq.com/metering-and-billing/>
- [KG05] Metering: <https://developer.konghq.com/metering-and-billing/metering/>
- [KG06] Meter AI Tokens: <https://developer.konghq.com/plugins/metering-and-billing/examples/meter-ai-tokens/>

## 17. Hạn chế nghiên cứu

- Chưa thử nghiệm trực tiếp Enterprise account của từng sản phẩm.
- Documentation có thể mô tả behavior dự kiến khác implementation tại một phiên bản cụ thể.
- Tập đối thủ chưa bao gồm toàn bộ thị trường.
- Một số sản phẩm tách gateway, billing và organization management thành module khác nhau.
- “Không tìm thấy” không chứng minh “không tồn tại”.
- Việc Asymptotic có thực sự đạt reservation, ledger và idempotency phải được kiểm chứng tại Chương 5; không được trình bày như kết quả đã triển khai nếu nguyên mẫu chưa đáp ứng.
