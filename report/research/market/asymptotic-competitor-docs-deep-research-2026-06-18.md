# Nghiên cứu chuyên sâu tài liệu đối thủ của Asymptotic

**Ngày nghiên cứu:** 18/06/2026  
**Phạm vi nguồn:** Tài liệu kỹ thuật, API reference, repository, pricing page và thông báo chính thức của nhà cung cấp.  
**Đầu ra phục vụ:** Nâng cấp mục 2.1 “Khảo sát hiện trạng” của báo cáo Asymptotic -- AI Agent Financial Gateway.

## 1. Kết luận điều hành

Nghiên cứu sâu tài liệu sản phẩm cho thấy thị trường AI Gateway đã tiến xa hơn bản tổng hợp ban đầu:

- Cloudflare AI Gateway đã có spend limit kiểm tra trước khi gửi request tới provider, prepaid credit, top-up và auto-top-up.
- Portkey đã có Organization, Workspace, budget nhiều cấp, pre-request enforcement và Agent Registry.
- LiteLLM đã có hierarchy Organization–Team–User–Key, đăng ký Agent bên ngoài, Agent-linked key, Agent/session budget, spend attribution và trace.
- Helicone không còn chỉ là sản phẩm observability; sản phẩm hiện có AI Gateway, managed provider credential, prepaid credit, cost-based blocking, routing, retry, fallback và cache.

Do đó, Asymptotic không nên tự phân biệt bằng các capability riêng lẻ như:

- multi-provider gateway;
- virtual key;
- Agent registration;
- budget hoặc spend limit;
- cost tracking;
- prepaid credit;
- pre-provider blocking.

Trọng tâm thiết kế còn có thể bảo vệ được là tổ hợp:

```text
organization wallet
  → hierarchical internal allocation
  → developer-owned external AI Agent
  → estimated-cost reservation
  → provider execution
  → actual-usage reconciliation and release
  → idempotent ledger posting
  → unified policy–request–usage–cost–transaction trace
```

Tài liệu công khai đã khảo sát chưa cho thấy Cloudflare, Portkey, LiteLLM hoặc Helicone mô tả đầy đủ tổ hợp trên. Đây là kết luận về **trọng tâm tích hợp và tính nhất quán tài chính**, không phải tuyên bố rằng Asymptotic sở hữu từng chức năng riêng biệt đầu tiên trên thị trường.

## 2. Phương pháp nghiên cứu

### 2.1. Câu hỏi nghiên cứu

1. Các đối thủ trực tiếp đã hỗ trợ những capability nào liên quan đến FR01–FR10 của Asymptotic?
2. Budget có được kiểm tra trước provider call hay chỉ được báo cáo sau thực thi?
3. Sản phẩm có Agent như một thực thể được đăng ký, cấp key, giới hạn và truy vết riêng hay không?
4. “Wallet”, “credit” và “budget” của từng sản phẩm có cùng ý nghĩa với ví tổ chức và hạn mức nội bộ của Asymptotic hay không?
5. Sản phẩm có reservation, reconciliation theo usage thực tế, idempotent financial posting và ledger hay không?
6. Các sản phẩm agent-payment lân cận cung cấp nguyên tắc nào có thể tham khảo mà không làm lệch phạm vi đề tài?

### 2.2. Nhóm sản phẩm

**Đối thủ AI Gateway trực tiếp**

- Cloudflare AI Gateway.
- Portkey AI Gateway.
- LiteLLM Proxy.
- Helicone AI Gateway.

**Hạ tầng lân cận**

- Skyfire.
- Nevermined.
- x402.
- Stripe Machine Payments và Machine Payment Protocol.
- Google Agent Payments Protocol.
- Mastercard Agent Pay.
- Visa Trusted Agent Protocol.

Nhóm lân cận không được chấm chung với AI Gateway vì “budget”, “wallet”, “credential” và “settlement” trong thanh toán thương mại có ngữ nghĩa khác với kiểm soát chi phí AI Provider.

### 2.3. Quy tắc đánh giá capability

#### Trạng thái capability

- **F — Fully matched:** tài liệu mô tả rõ capability đúng với định nghĩa khảo sát.
- **P — Partial/analogous:** có capability gần tương đương nhưng khác ngữ nghĩa, cấp áp dụng hoặc bảo đảm.
- **X — Explicitly unsupported:** tài liệu chính thức nói rõ không hỗ trợ.
- **U — Undetermined:** chưa xác định được sau khi đọc tài liệu công khai; không đồng nghĩa với không tồn tại.
- **NA — Not applicable:** không phù hợp với loại sản phẩm.

#### Độ mạnh bằng chứng

- **E4:** tài liệu kỹ thuật hoặc API reference chính thức cho capability đang được cung cấp.
- **E3:** tài liệu chính thức cho beta, preview hoặc capability có giới hạn availability.
- **E2:** repository, sample hoặc SDK chính thức chứng minh cơ chế.
- **E1:** landing page, blog hoặc press release của nhà cung cấp.
- **E0:** chưa có bằng chứng công khai đủ dùng.

### 2.4. Hạn chế

- Nghiên cứu không thử nghiệm tài khoản enterprise của từng sản phẩm.
- Capability enterprise-only có thể không được mô tả đầy đủ công khai.
- Vendor documentation chứng minh sản phẩm tuyên bố và thiết kế capability, không tự chứng minh quy mô sử dụng hoặc hiệu quả production.
- Trạng thái sản phẩm thay đổi nhanh; kết quả chỉ phản ánh tài liệu truy cập ngày 18/06/2026.

## 3. Định nghĩa tiêu chí kiểm chứng

| Mã | Tiêu chí | Điều kiện được tính là hỗ trợ |
|---|---|---|
| D01 | Multi-provider gateway | Một gateway/proxy gọi được từ hai AI Provider trở lên |
| D02 | Gateway-issued credential | Client dùng credential của gateway thay vì provider key trực tiếp |
| D03 | Provider credential isolation | Provider credential được lưu và sử dụng ở phía gateway |
| D04 | Managed hierarchy | Có các entity quản trị như organization, workspace/team, user hoặc Agent |
| D05 | Registered Agent entity | Agent có bản ghi, ID hoặc registry riêng, không chỉ là metadata tự do |
| D06 | Scoped access control | Key/entity bị giới hạn theo model, provider, resource hoặc role |
| D07 | Request/token limits | Gateway kiểm tra RPM, TPM, concurrency hoặc quota |
| D08 | Monetary spend limit | Có giới hạn theo chi phí, không chỉ request/token |
| D09 | Pre-provider enforcement | Giới hạn được kiểm tra trước khi request được gửi tới provider |
| D10 | Strict reservation | Chi phí dự kiến được giữ trước call để chống overspend do concurrent request |
| D11 | Usage-cost attribution | Usage và cost được gắn với key, user, team hoặc Agent |
| D12 | Internal reconciliation | Chi phí dự kiến được điều chỉnh theo usage thực tế sau call |
| D13 | Request deduplication | Cùng idempotency key không tạo provider call trùng |
| D14 | Ledger idempotency | Cùng sự kiện không tạo financial posting trùng |
| D15 | Financial ledger | Có transaction/entry và trạng thái tài chính, không chỉ request log |
| D16 | Prepaid balance/top-up | Có monetary credit balance và luồng nạp tiền |
| D17 | Unified trace | Có thể liên kết identity, policy decision, provider call, usage, cost và posting |

“Budget” chỉ được ghi nhận tại D08 khi giới hạn dùng đơn vị tiền. Token quota không tự động được xem là monetary budget. “Pre-provider enforcement” không tự động chứng minh strict reservation.

## 4. Ma trận đối thủ AI Gateway trực tiếp

| Capability | Cloudflare | Portkey | LiteLLM | Helicone |
|---|---|---|---|---|
| D01 Multi-provider gateway | F/E4 | F/E4 | F/E4 | F/E4 |
| D02 Gateway-issued credential | P/E4 | F/E4 | F/E4 | F/E4 |
| D03 Provider credential isolation | F/E3 | F/E4 | F/E4 | F/E4 |
| D04 Managed hierarchy | P/E4 | F/E4 | F/E4 | P/E4 |
| D05 Registered Agent entity | U/E0 | F/E4 | F/E4 | U/E0 |
| D06 Scoped access control | P/E4 | F/E4 | F/E4 | P/E4 |
| D07 Request/token limits | F/E4 | F/E4 | F/E4 | F/E4 |
| D08 Monetary spend limit | F/E3 | F/E4 | F/E4 | P/E4 |
| D09 Pre-provider enforcement | F/E3 | F/E4 | P/E4 | P/E4 |
| D10 Strict reservation | U/E0 | U/E0 | U/E0 | U/E0 |
| D11 Usage-cost attribution | F/E4 | F/E4 | F/E4 | F/E4 |
| D12 Internal reconciliation | U/E0 | U/E0 | P/E4 | U/E0 |
| D13 Request deduplication | U/E0 | U/E0 | U/E0 | U/E0 |
| D14 Ledger idempotency | U/E0 | U/E0 | U/E0 | U/E0 |
| D15 Financial ledger | U/E0 | U/E0 | P/E4 | U/E0 |
| D16 Prepaid balance/top-up | F/E4 | U/E0 | U/E0 | F/E4 |
| D17 Unified financial trace | P/E4 | P/E4 | P/E4 | P/E4 |

Các ô `U/E0` chỉ có nghĩa là chưa tìm thấy mô tả đủ rõ trong tài liệu công khai được đọc.

## 5. Cloudflare AI Gateway

### 5.1. Năng lực được xác nhận

Cloudflare AI Gateway hỗ trợ provider-native endpoint và API tương thích OpenAI cho nhiều provider. Gateway cung cấp cache, retry, fallback, rate limiting, dynamic routing, logging, analytics, cost tracking, custom metadata và OpenTelemetry/Logpush export.[CF01][CF06][CF07]

Bring Your Own Keys lưu provider credential trong Cloudflare Secrets Store và hỗ trợ nhiều key alias. Tài liệu đánh dấu chức năng này là beta.[CF04]

Spend limit hỗ trợ:

- ngân sách theo USD;
- fixed hoặc rolling window;
- lọc theo model, provider và custom metadata;
- metadata có thể biểu diễn user hoặc team;
- kiểm tra rule trước khi dispatch tới provider;
- trả HTTP 429 khi vượt giới hạn.[CF02]

Đây là pre-provider enforcement rõ ràng. Tuy nhiên, tài liệu cũng nói cost của request hiện tại chỉ được ghi nhận sau khi request hoàn tất và enforcement có tính eventually consistent. Các request đồng thời có thể làm tổng chi phí vượt giới hạn trong một khoảng ngắn.[CF02]

Cloudflare Unified Billing sử dụng credit balance. Khách hàng có thể:

- mua credit;
- top-up thủ công;
- cấu hình auto-top-up theo threshold;
- dùng credit để thanh toán inference qua provider được hỗ trợ;
- để hệ thống tự động trừ credit.[CF03]

### 5.2. Giới hạn so với Asymptotic

- Authenticated Gateway dùng token ở cấp Cloudflare account; tài liệu lưu ý token có quyền `AI Gateway Run` có thể gọi mọi gateway trong account. Đây không phải Agent-specific virtual key hierarchy.[CF05]
- User/team trong spend rule là metadata bucket, chưa được thể hiện là entity tài chính có quan hệ sở hữu và phân bổ.
- Spend limit không phải strict reservation vì có eventual consistency.
- Chưa tìm thấy hold–capture–release lifecycle theo estimated và actual cost.
- Chưa tìm thấy request idempotency bảo đảm không gọi provider trùng.
- Chưa tìm thấy idempotent ledger posting hoặc customer-visible transaction ledger.

### 5.3. Kết luận

Cloudflare phủ phần lớn gateway, prepaid credit và pre-request spend control. Khoảng cách còn lại với Asymptotic nằm ở hierarchical financial ownership, strict reservation, reconciliation và ledger semantics.

## 6. Portkey AI Gateway

### 6.1. Năng lực được xác nhận

Portkey cung cấp Universal API, multi-provider routing, load balancing, circuit breaking, retry, fallback, cache và semantic cache.[PK01]

Provider credential được lưu trong Integrations ở cấp Organization và có thể cấp cho Workspace. Upstream credential không cần phân phối cho end user.[PK02]

Mô hình quản trị được tài liệu hóa:

```text
Organization
  → Workspace
    → Members / API keys / scoped resources
```

Organization và Workspace có role riêng. Các quyền chi tiết đối với analytics, logs, integrations và key phụ thuộc gói sản phẩm, trong đó nhiều khả năng governance nằm ở Enterprise.[PK07][PK08][PK09][PK12]

Portkey hỗ trợ rate limit hoặc budget tại nhiều lớp:

- API key;
- Workspace;
- integration/provider;
- usage policy;
- metadata group như user hoặc customer.[PK03][PK04][PK05][PK06]

Official guide mô tả mọi applicable check được thực hiện trước provider dispatch. Request bị budget policy chặn trả lỗi 412 và không tạo provider spend.[PK03]

Cost được tính từ input/output token và model pricing. Sản phẩm hỗ trợ custom pricing, discount và markup.[PK11]

Portkey Agent Gateway hiện có Agent Registry và virtual Agent Server. Vì vậy, không còn chính xác nếu mô tả Portkey không có Agent registration.[PK10]

### 6.2. Giới hạn so với Asymptotic

- Agent Registry chưa được tài liệu hóa như một entity con của developer chịu trách nhiệm tài chính.
- Hierarchy Organization → Workspace → Member không tương đương hoàn toàn Organization → Team → Developer → Agent.
- Nếu model không có pricing được hỗ trợ và log có cost bằng 0, request có thể không được tính vào provider budget.[PK11]
- `credit_limit` trong usage policy là giới hạn sử dụng, không phải bằng chứng về stored-value wallet hoặc top-up.
- Chưa tìm thấy strict monetary reservation.
- Chưa tìm thấy estimated-to-actual reconciliation state machine.
- Chưa tìm thấy inference idempotency hoặc financial posting idempotency.
- Chưa tìm thấy customer-visible financial ledger.

### 6.3. Kết luận

Portkey là đối thủ trực tiếp mạnh về organization/workspace governance, Agent registry và pre-request budget policy. Asymptotic chỉ có thể phân biệt ở mô hình trách nhiệm developer–Agent kết hợp nguồn tiền, allocation, reservation, reconciliation và ledger.

## 7. LiteLLM Proxy

### 7.1. Năng lực được xác nhận

LiteLLM cung cấp OpenAI-compatible gateway tới hơn 100 LLM/provider, routing, load balancing, retry, fallback, health routing, budget routing và cache.[LL01][LL08]

Hierarchy được tài liệu hóa:

```text
Organization
  → Team
    → Internal User
      → Virtual Key
```

Virtual key có thể gắn với user, team, cả hai hoặc Agent. Key có thể bị giới hạn model, budget và rate.[LL02][LL03]

LiteLLM A2A Agent Gateway hỗ trợ đăng ký các Agent bên ngoài, gồm A2A, LangGraph, Vertex AI Agent, Azure Foundry Agent và Bedrock AgentCore. Agent có thể có:

- Agent ID;
- virtual key;
- permission;
- request log;
- rate limit;
- session budget;
- trace grouping.[LL04][LL05]

Budget được hỗ trợ ở nhiều cấp:

- proxy;
- organization;
- team;
- team member;
- user;
- virtual key;
- model;
- end-user;
- tag;
- Agent;
- Agent session.[LL03][LL05]

Agent session budget kiểm tra accumulated spend trước mỗi call. Sau call thành công, cost mới được cộng vào accumulated spend.[LL05]

Cost tracking lưu spend theo key, user, team, tag và Agent. `spend_logs` cung cấp chi tiết theo từng call và có thể dùng cho audit hoặc ETL.[LL06]

`X-LiteLLM-Agent-Id` hỗ trợ Agent attribution; `X-LiteLLM-Trace-Id` liên kết các call trong cùng execution.[LL04][LL05]

### 7.2. Giới hạn so với Asymptotic

- Budget dựa trên accumulated spend sau các call thành công; chưa phải estimated-cost reservation chống overspend do concurrency.
- Agent attribution chưa thiết lập đầy đủ quan hệ một developer sở hữu/quản lý Agent và cấp hạn mức từ allocation được kế thừa.
- `spend_logs` là usage/cost record, chưa phải double-entry hoặc wallet transaction ledger.
- Chưa tìm thấy organization-owned prepaid balance và top-up.
- Chưa tìm thấy hold, release và capture theo estimated/actual cost.
- Trace ID và JSON-RPC request ID không được tài liệu hóa như duplicate-call hoặc duplicate-charge protection.
- Nhiều chức năng RBAC, secret manager, SSO/SCIM và governance yêu cầu Enterprise.[LL03][LL07][LL09]

### 7.3. Kết luận

LiteLLM là đối thủ gần Asymptotic nhất trong nhóm đã nghiên cứu. External Agent registration, Agent-linked key, Agent/session budgets và spend trace không còn là differentiation đủ mạnh. Phần còn lại cần nhấn mạnh strict financial consistency và ledger workflow.

## 8. Helicone AI Gateway

### 8.1. Năng lực được xác nhận

Helicone hiện cung cấp OpenAI-compatible AI Gateway tới hơn 100 model/provider, không chỉ observability.[HC01]

Gateway hỗ trợ:

- Helicone API key;
- BYOK provider credential;
- managed provider credential;
- provider routing;
- BYOK priority;
- managed-key fallback;
- cheapest-provider routing;
- retry;
- cache.[HC01][HC02][HC05][HC06]

Helicone Custom Rate Limits hỗ trợ giới hạn theo:

- số request;
- cost tính bằng cent;
- global scope;
- user;
- custom property, ví dụ organization.[HC03]

Cost-based sliding-window rule có thể chặn request ở gateway và trả HTTP 429. Đây là cơ chế gần với spend limit, nhưng tài liệu trình bày nó như rate limiting thay vì hierarchical budget allocation.[HC03]

Helicone theo dõi request, session, user, custom property, trace, cost, alert và report.[HC04][HC07]

Khi dùng managed provider key, khách hàng nạp credit vào Helicone account và inference tiêu thụ credit đó. Tài liệu có lỗi insufficient credit hoặc wallet suspended.[HC02]

### 8.2. Giới hạn so với Asymptotic

- Chưa tìm thấy hierarchy Team → Developer → Agent.
- Custom property có thể mô phỏng organization hoặc customer, nhưng không phải managed financial entity.
- Chưa tìm thấy Agent registry.
- Chưa xác định chính xác cost counter update order hoặc mức overshoot khi có concurrent request.
- Token-based custom limits và nhiều simultaneous policy được ghi là “coming soon” tại thời điểm khảo sát.[HC03]
- “Replay session” là công cụ evaluation gửi lại request, không phải replay protection.
- Chưa tìm thấy reservation, capture/release, financial posting idempotency hoặc transaction ledger.

### 8.3. Kết luận

Helicone phải được xếp vào nhóm AI Gateway trực tiếp. Asymptotic không nên so sánh Helicone như một sản phẩm chỉ quan sát chi phí sau thực thi.

## 9. So sánh ngữ nghĩa các khái niệm tài chính

| Khái niệm | Cloudflare | Portkey | LiteLLM | Helicone | Asymptotic |
|---|---|---|---|---|---|
| Budget | USD spend rule | Cost/token usage policy | Accumulated spend ceiling | Cost-based rate limit | Hạn mức trên nguồn tiền và allocation |
| Wallet/credit | Prepaid Cloudflare credits | Chưa xác định stored value | Chưa xác định stored value | Prepaid Helicone credits | Ví tổ chức và transaction history |
| Pre-request check | Có, eventually consistent | Có | Có theo accumulated spend | Có ở rate-limit layer | Kiểm tra đồng thời ví và hạn mức nhiều cấp |
| Reservation | Chưa tìm thấy | Chưa tìm thấy | Chưa tìm thấy | Chưa tìm thấy | Tạm giữ/kiểm tra estimated cost |
| Reconciliation | Cost ghi sau response | Cost tính từ usage | Spend cộng sau call | Cost tracking | Điều chỉnh theo actual usage |
| Ledger | Chưa tìm thấy | Chưa tìm thấy | Spend logs, tương tự một phần | Chưa tìm thấy | Transaction/ledger entry có lũy đẳng |

Các khái niệm này không nên được quy về dấu “Có/Không” đơn giản trong báo cáo.

## 10. Hạ tầng lân cận: Agent identity và machine payments

### 10.1. Skyfire

Skyfire tài liệu hóa Agent Account gắn với một user, Agent API key, buyer/seller wallet và KYA/PAY token. KYA token có thể mang identity của human/organization, Agent và Agent Platform. PAY hoặc KYA-PAY token chứa giá trị, tiền tệ, settlement type hoặc pricing scheme.[SF01][SF02][SF03][SF04]

Giá trị tham khảo:

- Agent là entity riêng nhưng gắn với principal;
- credential cho Agent phải có phạm vi;
- authorization được kiểm tra trước giao dịch;
- identity và payment commitment có bằng chứng riêng.

Không nên khẳng định Skyfire có immutable audit ledger nếu chỉ dựa trên transaction history. Độ trưởng thành và quy mô triển khai chủ yếu do vendor công bố.

### 10.2. Nevermined

Nevermined cung cấp đăng ký Agent/service, pricing, rate limit, access control, per-request metering, entitlement, x402 settlement, delegated/virtual card và spending rule theo transaction, ngày hoặc merchant category.[NV01][NV02]

Giá trị tham khảo:

- metering phải gắn trực tiếp với entitlement và payment;
- spending policy có thể được kiểm tra trước khi Agent tiêu thụ dịch vụ;
- revoke và analytics là thành phần của Agent financial control.

Agent registration của Nevermined không tự chứng minh verified Agent identity. Các tuyên bố về adoption và “immutable record” cần được xem là vendor assertion nếu thiếu mô tả kỹ thuật độc lập.

### 10.3. x402

x402 dùng HTTP 402 để server trả payment requirement, client ký payment payload và gửi lại request để facilitator verify/settle. Protocol hỗ trợ exact, `upto`, batch settlement, receipt và idempotency extension.[X401][X402][X403][X404]

Giá trị tham khảo:

- authorize-before-execute;
- request-level financial proof;
- replay/idempotency phải là một phần của giao thức;
- receipt phải liên kết được với request.

x402 không cung cấp đầy đủ principal/Agent ownership, hierarchical budget, plan management hoặc AI usage metering. FAQ nêu các phần này có thể cần hệ thống bên ngoài.[X402]

### 10.4. Stripe Machine Payments và MPP

Stripe hỗ trợ HTTP 402 machine-payment flow với crypto hoặc Shared Payment Token, PaymentIntent, authorization, retry, receipt, refund, reporting và settlement.[ST01][ST02]

Tài liệu có hướng dẫn live mode, nhưng availability phụ thuộc rail, khu vực và approval. Một số Crypto PaymentIntent API dùng preview version. Vì vậy, không nên gắn nhãn chung “beta” hoặc “GA” cho toàn bộ sản phẩm.

Stripe cung cấp payment lifecycle thực, nhưng không cung cấp hierarchical AI-consumption budget theo mô hình Asymptotic.

### 10.5. Google AP2

AP2 là protocol/specification cùng sample và reference implementation, chưa phải production payment network. Protocol mô hình hóa Intent, Cart, Checkout và Payment Mandate bằng verifiable credential và cryptographic signature.[AP01][AP02]

Nguyên tắc có thể áp dụng:

- intent phải được xác minh, không chỉ suy luận từ hành vi Agent;
- authorization artifact cần version, phạm vi và chữ ký;
- transaction evidence phải chống sửa đổi và hỗ trợ accountability.

### 10.6. Mastercard Agent Pay

Mastercard công bố Agent Pay ngày 29/04/2025. Tài liệu mô tả registered Agent, Agentic Token, tokenized payment credential, consumer consent, order intent và network visibility.[MC01][MC02]

Nhiều capability vẫn được diễn đạt theo hướng “will deliver”, “will collaborate” hoặc “we are building”. Vì vậy, nên mô tả đây là chương trình đã công bố và đang được phát triển/triển khai với đối tác, không khẳng định GA toàn cầu.

### 10.7. Visa Trusted Agent Protocol

Visa ghi rõ protocol đang trong quá trình phát triển và triển khai, có thể chưa khả dụng ở mọi thị trường. Protocol tài liệu hóa Agent intent, payment identifier, merchant/purpose-bound signature, timestamp, session/key metadata, registry/public key và anti-replay/relay.[VS01][VS02]

Giá trị tham khảo:

- credential gắn với merchant, purpose và thời gian;
- request cần anti-replay;
- Agent authentication tách khỏi payment settlement.

Protocol không tự cung cấp wallet budget, AI usage metering hoặc financial ledger.

## 11. Các phát hiện làm thay đổi bản tổng hợp trước

### 11.1. Cloudflare bị đánh giá thấp

Bản trước cần sửa vì Cloudflare đã có:

- prepaid credit;
- manual và automatic top-up;
- automatic deduction;
- USD spend limit;
- metadata-based user/team buckets;
- pre-provider enforcement.

Điểm yếu còn lại không phải “không có budget” mà là eventual consistency, thiếu strict reservation và thiếu hierarchical ledger workflow.

### 11.2. Portkey có Agent Registry

Không nên tiếp tục ghi “chưa xác định Agent registration”. Khoảng cách thực tế là Agent Registry chưa được mô tả như một phần của chuỗi financial ownership từ organization đến developer và Agent.

### 11.3. LiteLLM gần Asymptotic hơn dự kiến

LiteLLM đã có:

- external Agent registration;
- Agent-linked virtual key;
- per-Agent và session budget;
- iteration cap;
- Agent spend attribution;
- trace grouping;
- team/key access control cho Agent.

External Agent registration, Agent API key và Agent budget không còn là differentiation riêng lẻ.

### 11.4. Helicone là đối thủ trực tiếp

Helicone hiện có đầy đủ gateway, managed key, prepaid credit, routing và cost-based blocking. Không nên đặt Helicone chỉ trong nhóm observability.

### 11.5. Wallet/top-up không còn độc đáo

Cloudflare và Helicone đều có prepaid credit/top-up. Asymptotic cần phân biệt:

- tiền thuộc ví tổ chức;
- allocation nội bộ không tạo ví độc lập;
- trách nhiệm chi phí được xác định qua Team, Developer và Agent;
- ledger ghi allocation, revoke, reservation và actual charge.

## 12. Khoảng trống thị trường sau khi hiệu chỉnh

### 12.1. Capability đã phổ biến

- Multi-provider gateway.
- Gateway-issued credential hoặc virtual key.
- Provider credential isolation.
- Routing, retry, fallback và cache.
- Usage/cost tracking.
- Rate limit và quota.
- Monetary spend limit.
- Pre-provider blocking.
- Organization/team/user governance ở một số sản phẩm.

### 12.2. Capability đang xuất hiện mạnh

- Registered Agent entity.
- Agent-linked key.
- Agent/session budget.
- Agent spend attribution.
- Prepaid inference credit.
- Request-level payment proof.
- Mandate và Agent identity trong agentic commerce.

### 12.3. Capability chưa được xác lập đầy đủ trong tài liệu đối thủ đã đọc

- Organization wallet kết hợp allocation phân cấp tới Team, Developer và Agent.
- Developer là chủ thể quản lý và chịu trách nhiệm trực tiếp cho external Agent.
- Estimated-cost reservation trước provider call với bảo đảm chống overspend khi concurrent.
- Capture/release hoặc reconciliation dựa trên actual usage.
- Request deduplication bảo đảm không gọi provider trùng.
- Idempotent financial posting bảo đảm không ghi chi phí trùng.
- Customer-visible transaction ledger liên kết allocation, reservation, charge và release.
- Trace thống nhất từ policy decision đến provider execution và financial posting.

### 12.4. Phát biểu định vị có thể dùng

> Asymptotic tập trung vào kiểm soát tài chính nhất quán theo từng request của AI Agent bên ngoài. Hệ thống xác định chủ thể chịu chi phí theo cơ cấu tổ chức, kiểm tra đồng thời ví và hạn mức phân cấp trước khi gọi AI Provider, đối soát chi phí theo usage thực tế, ghi nhận ledger có lũy đẳng và duy trì liên kết truy vết giữa policy decision, provider execution và financial posting.

Nên gọi đây là **trọng tâm thiết kế** hoặc **khoảng trống chưa được thể hiện đầy đủ trong tài liệu công khai được khảo sát**, không gọi là khả năng duy nhất trên thị trường.

## 13. Khuyến nghị cho mục 2.1 của báo cáo

### 13.1. Cấu trúc

1. Bối cảnh và vấn đề hiện tại.
2. Phương pháp khảo sát và taxonomy bằng chứng.
3. Khảo sát Cloudflare, Portkey, LiteLLM và Helicone.
4. Bảng so sánh chỉ dành cho AI Gateway trực tiếp.
5. Phần ngắn về Agent identity và machine-payment trends.
6. Nhận xét capability phổ biến, capability mới xuất hiện và capability chưa xác định.
7. Bài toán đặt ra cho Asymptotic.

### 13.2. Quy tắc viết

- Mỗi capability quan trọng phải có citation tới đúng trang tài liệu.
- Phân biệt GA, beta, preview, pilot và announcement.
- Không dùng “không có” khi kết quả chỉ là chưa tìm thấy.
- Không gộp nhiều vendor trong một cột.
- Không so “Mục tiêu Asymptotic” với capability production như thể chúng có cùng trạng thái.
- Không dùng screenshot chỉ để trang trí; hình phải chứng minh capability.
- Dùng “đối soát nội bộ theo usage thực tế” thay cho “settlement với AI Provider” để giữ đúng phạm vi MVP.

## 14. Nguồn chính

### Cloudflare

- [CF01] Cloudflare, “AI Gateway — Overview”: <https://developers.cloudflare.com/ai-gateway/>
- [CF02] Cloudflare, “Spend limits”: <https://developers.cloudflare.com/ai-gateway/features/spend-limits/>
- [CF03] Cloudflare, “Unified Billing”: <https://developers.cloudflare.com/ai-gateway/features/unified-billing/>
- [CF04] Cloudflare, “Bring Your Own Keys”: <https://developers.cloudflare.com/ai-gateway/configuration/bring-your-own-keys/>
- [CF05] Cloudflare, “Authenticated Gateway”: <https://developers.cloudflare.com/ai-gateway/configuration/authentication/>
- [CF06] Cloudflare, “Dynamic routing”: <https://developers.cloudflare.com/ai-gateway/features/dynamic-routing/>
- [CF07] Cloudflare, “Costs”: <https://developers.cloudflare.com/ai-gateway/observability/costs/>
- [CF08] Cloudflare, “Pricing”: <https://developers.cloudflare.com/ai-gateway/reference/pricing/>
- [CF09] Cloudflare, “Limits”: <https://developers.cloudflare.com/ai-gateway/reference/limits/>

### Portkey

- [PK01] Portkey, “AI Gateway”: <https://portkey.ai/docs/product/ai-gateway>
- [PK02] Portkey, “Integrations”: <https://docs.portkey.ai/docs/product/model-catalog/integrations>
- [PK03] Portkey, “Enforcing Limits and Budgets”: <https://docs.portkey.ai/docs/guides/use-cases/enforcing-limits-and-budgets>
- [PK04] Portkey, “API Key Budget and Rate Limits”: <https://docs.portkey.ai/docs/product/administration/enforce-budget-and-rate-limit>
- [PK05] Portkey, “Workspace Budget and Rate Limits”: <https://docs.portkey.ai/docs/product/administration/enforce-workspace-budget-limts-and-rate-limits>
- [PK06] Portkey, “Usage and Rate Limit Policies”: <https://docs.portkey.ai/docs/product/enterprise-offering/budget-policies>
- [PK07] Portkey, “Organizations”: <https://docs.portkey.ai/docs/product/enterprise-offering/org-management/organizations>
- [PK08] Portkey, “Workspaces”: <https://docs.portkey.ai/docs/product/enterprise-offering/org-management/workspaces>
- [PK09] Portkey, “User Roles and Permissions”: <https://docs.portkey.ai/docs/product/enterprise-offering/org-management/user-roles-and-permissions>
- [PK10] Portkey, “Agent Registry”: <https://docs.portkey.ai/docs/product/agent-gateway/registry>
- [PK11] Portkey, “Cost Management”: <https://docs.portkey.ai/docs/product/observability/cost-management>
- [PK12] Portkey, “Product Feature Comparison”: <https://portkey.ai/docs/product/product-feature-comparison>
- [PK13] Portkey, “Pricing”: <https://portkey.ai/pricing>

### LiteLLM

- [LL01] LiteLLM, “Proxy Server”: <https://docs.litellm.ai/docs/simple_proxy>
- [LL02] LiteLLM, “Virtual Keys”: <https://docs.litellm.ai/docs/proxy/virtual_keys>
- [LL03] LiteLLM, “Budgets and Rate Limits”: <https://docs.litellm.ai/docs/proxy/users>
- [LL04] LiteLLM, “A2A Agent Gateway”: <https://docs.litellm.ai/docs/a2a>
- [LL05] LiteLLM, “Agent Iteration Budgets”: <https://docs.litellm.ai/docs/a2a_iteration_budgets>
- [LL06] LiteLLM, “Spend Tracking”: <https://docs.litellm.ai/docs/proxy/cost_tracking>
- [LL07] LiteLLM, “Access Control”: <https://docs.litellm.ai/docs/proxy/access_control>
- [LL08] LiteLLM, “Routing and Load Balancing”: <https://docs.litellm.ai/docs/routing-load-balancing>
- [LL09] LiteLLM, “Secret Managers”: <https://docs.litellm.ai/docs/secret_managers/overview>
- [LL10] LiteLLM, “Enterprise”: <https://www.litellm.ai/enterprise>

### Helicone

- [HC01] Helicone, “AI Gateway Overview”: <https://docs.helicone.ai/gateway/overview>
- [HC02] Helicone, “Provider Routing”: <https://docs.helicone.ai/gateway/provider-routing>
- [HC03] Helicone, “Custom LLM Rate Limits”: <https://docs.helicone.ai/features/advanced-usage/custom-rate-limits>
- [HC04] Helicone, “Cost Tracking”: <https://docs.helicone.ai/guides/cookbooks/cost-tracking>
- [HC05] Helicone, “Retries”: <https://docs.helicone.ai/features/advanced-usage/retries>
- [HC06] Helicone, “Caching”: <https://docs.helicone.ai/features/advanced-usage/caching>
- [HC07] Helicone, “Sessions”: <https://docs.helicone.ai/features/sessions>
- [HC08] Helicone, “Pricing”: <https://www.helicone.ai/pricing>

### Agent-finance infrastructure

- [SF01] Skyfire, “Developer Documentation”: <https://docs.skyfire.xyz/docs/developer-documentation>
- [SF02] Skyfire, “KYAPay Tokens”: <https://docs.skyfire.xyz/docs/kyapay-tokens.md>
- [SF03] Skyfire, “KYA Token”: <https://docs.skyfire.xyz/docs/kya-token.md>
- [SF04] Skyfire, “PAY Token”: <https://docs.skyfire.xyz/docs/pay-token.md>
- [NV01] Nevermined, “Payments Infrastructure for AI Agents”: <https://nevermined.ai/>
- [NV02] Nevermined, “Documentation”: <https://nevermined.ai/docs/getting-started/welcome>
- [X401] x402, “Official site”: <https://www.x402.org/>
- [X402] x402, “FAQ”: <https://docs.x402.org/faq.md>
- [X403] x402, “Buyer Quickstart”: <https://docs.x402.org/getting-started/quickstart-for-buyers>
- [X404] x402 Foundation, “x402 repository”: <https://github.com/x402-foundation/x402>
- [ST01] Stripe, “Machine Payments”: <https://docs.stripe.com/payments/machine.md>
- [ST02] Stripe, “Machine Payment Protocol”: <https://docs.stripe.com/payments/machine/mpp.md>
- [AP01] Google, “Agent Payments Protocol”: <https://github.com/google-agentic-commerce/AP2>
- [AP02] Google, “AP2 FAQ”: <https://github.com/google-agentic-commerce/AP2/blob/main/docs/faq.md>
- [MC01] Mastercard, “Mastercard unveils Agent Pay”: <https://www.mastercard.com/us/en/news-and-trends/press/2025/april/mastercard-unveils-agent-pay-pioneering-agentic-payments-technology-to-power-commerce-in-the-age-of-ai.html>
- [MC02] Mastercard, “Agent Pay”: <https://www.mastercard.com/us/en/business/artificial-intelligence/mastercard-agent-pay.html>
- [VS01] Visa, “Trusted Agent Protocol”: <https://developer.visa.com/capabilities/trusted-agent-protocol>
- [VS02] Visa, “Trusted Agent Protocol sample”: <https://github.com/visa/trusted-agent-protocol>

## 15. Các đối thủ cần khảo sát ở vòng tiếp theo

Reviewer xác định tập đối thủ hiện tại vẫn chưa đại diện đầy đủ. Một vòng nghiên cứu tiếp theo nên sàng lọc:

- Azure API Management AI Gateway;
- OpenRouter;
- Kong AI Gateway.

Các sản phẩm này có khả năng làm hẹp thêm differentiation về external Agent registration, organizational quota, managed key, metering và billing. Vì chưa nằm trong yêu cầu “competitors listed” của vòng này, chúng chưa được đưa vào ma trận kết luận.
