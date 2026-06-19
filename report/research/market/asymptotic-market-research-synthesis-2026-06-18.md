# Tổng hợp nghiên cứu thị trường cho Asymptotic -- AI Agent Financial Gateway

**Ngày tổng hợp:** 18/06/2026  
**Mốc thông tin chính:** 16/06/2026  
**Mục đích:** Cung cấp cơ sở nghiên cứu để nâng cấp mục 2.1 “Khảo sát hiện trạng” trong báo cáo.  
**Phạm vi:** Hạ tầng kiểm soát truy cập, chính sách và chi phí khi AI Agent bên ngoài gọi dịch vụ AI trả phí thông qua một Gateway.

## 1. Tóm tắt kết quả

Thị trường liên quan đến Asymptotic đang phát triển theo hai hướng gần nhau nhưng chưa hoàn toàn hợp nhất:

1. **AI Gateway và LLM infrastructure:** cung cấp một điểm truy cập thống nhất tới nhiều AI Provider, virtual API key, định tuyến, fallback, rate limit, theo dõi usage và chi phí.
2. **Hạ tầng tài chính cho AI Agent:** cung cấp định danh Agent, mandate hoặc phạm vi ủy quyền, credential có giới hạn, spending policy, thanh toán máy--máy và dấu vết kiểm toán.

Các sản phẩm AI Gateway như Cloudflare AI Gateway, Portkey và LiteLLM đã hỗ trợ nhiều chức năng mà Asymptotic quan tâm. Vì vậy, không thể định vị Asymptotic chỉ bằng các khả năng “proxy nhiều nhà cung cấp”, “theo dõi chi phí” hoặc “giới hạn ngân sách”.

Khoảng trống phù hợp hơn với Asymptotic là việc kết hợp trong cùng một mô hình:

- đăng ký và định danh riêng AI Agent bên ngoài;
- xác định quan hệ sở hữu và trách nhiệm chi phí theo chuỗi tổ chức → đội ngũ → lập trình viên → Agent;
- sử dụng ví tổ chức làm nguồn tiền và phân bổ hạn mức nội bộ theo cấp;
- kiểm tra định danh, policy, quota và budget trước khi gọi AI Provider;
- tạm giữ hoặc kiểm tra chi phí dự kiến trước thực thi và quyết toán theo usage thực tế;
- kiểm soát lũy đẳng để hạn chế ghi nhận chi phí trùng;
- liên kết request, Agent, API key, provider, model, usage, cost, transaction và trace.

Kết luận trên là một **khoảng trống tích hợp và mô hình hóa**, không phải tuyên bố rằng từng chức năng riêng lẻ chưa tồn tại trên thị trường.

## 2. Phạm vi và ranh giới nghiên cứu

### 2.1. Nội dung thuộc phạm vi

- AI Gateway hoặc LLM Gateway.
- Gateway-issued key hoặc virtual API key.
- Provider credential được quản lý tập trung.
- Rate limit, quota, budget và spending limit.
- Theo dõi usage, token và chi phí.
- Định danh và quản lý AI Agent như một workload riêng.
- Phân bổ trách nhiệm chi phí theo tổ chức và các đơn vị nội bộ.
- Kiểm soát trước thực thi, quyết toán và truy vết tài chính.
- Idempotency đối với request và giao dịch.

### 2.2. Nội dung chỉ dùng để tham khảo xu hướng

- Agent tự mua hàng bằng thẻ.
- Agent thanh toán qua ngân hàng hoặc stablecoin.
- Agentic commerce và checkout.
- Custody, acquiring, card issuing, chargeback và settlement trên mạng thanh toán.
- Thanh toán agent-to-agent.

Những nội dung này giúp nhận diện xu hướng về Agent identity, mandate, credential và audit trail, nhưng không phải phạm vi triển khai chính của Asymptotic.

### 2.3. Nội dung ngoài phạm vi Asymptotic

- Tạo, huấn luyện, triển khai hoặc điều phối AI Agent.
- Payment Gateway production đầy đủ.
- Phát hành thẻ hoặc tài khoản ngân hàng cho Agent.
- Custody tài sản số.
- Thanh toán thương mại tự động cho người tiêu dùng.
- Settlement production với AI Provider.
- Chứng khoán, đầu tư, tín dụng hoặc bảo hiểm tự động.

## 3. Bối cảnh thị trường

AI Agent có thể tạo nhiều request nối tiếp, thử lại hoặc chạy song song trong quá trình hoàn thành một nhiệm vụ. Do số bước và kết quả trung gian không cố định, tổng usage và chi phí thường khó xác định chính xác trước khi thực thi.

Nếu tổ chức cấp trực tiếp provider API key cho từng ứng dụng hoặc Agent, credential bị phân tán và việc xác định chủ thể chịu chi phí trở nên khó khăn. Các AI Gateway giải quyết một phần vấn đề này bằng cách đặt một lớp trung gian giữa ứng dụng và AI Provider.

Thị trường AI Gateway hiện tập trung vào các nhóm chức năng:

- API thống nhất cho nhiều provider;
- định tuyến, retry và fallback;
- cache và rate limit;
- virtual key hoặc gateway credential;
- logging và observability;
- theo dõi token và chi phí;
- budget hoặc spending limit ở một số cấp quản lý.

Song song, thị trường agentic payment đang phát triển:

- Agent identity và Know Your Agent;
- mandate hoặc bằng chứng ủy quyền;
- credential có phạm vi;
- giới hạn số tiền, mục đích hoặc đối tác nhận tiền;
- payment proof và thanh toán theo từng API request;
- audit trail và reconciliation.

Hai hướng phát triển này cho thấy nhu cầu kiểm soát Agent như một chủ thể phần mềm riêng đang tăng lên. Tuy nhiên, Asymptotic cần giữ định vị hẹp hơn: Gateway kiểm soát chi phí sử dụng dịch vụ AI, không phải hạ tầng thanh toán tổng quát cho Agent.

## 4. Tiêu chí khảo sát sản phẩm

Các sản phẩm được đánh giá theo các tiêu chí sau:

| Mã | Tiêu chí | Liên hệ với Asymptotic |
|---|---|---|
| C01 | Proxy hoặc API thống nhất cho nhiều AI Provider | FR04, FR07 |
| C02 | Gateway-issued key hoặc virtual API key | FR03 |
| C03 | Quản lý provider credential tập trung | FR04, NFR02 |
| C04 | Rate limit, quota hoặc policy | FR05, FR10 |
| C05 | Theo dõi usage, token và chi phí | FR09 |
| C06 | Budget hoặc spending limit | FR05, FR06 |
| C07 | Đăng ký AI Agent như một thực thể riêng | FR02 |
| C08 | Gắn Agent với tổ chức và người quản lý | FR01, FR02 |
| C09 | Phân bổ hạn mức tổ chức → đội ngũ → lập trình viên → Agent | FR01, FR05 |
| C10 | Kiểm tra hoặc tạm giữ chi phí trước khi gọi provider | FR06, NFR01 |
| C11 | Quyết toán theo usage thực tế | FR01, FR09, NFR01 |
| C12 | Idempotency để tránh request hoặc chi phí trùng | FR08, NFR01 |
| C13 | Trace thống nhất giữa request, usage, cost và transaction | FR09, NFR03 |

## 5. Nhóm sản phẩm cạnh tranh trực tiếp

### 5.1. Cloudflare AI Gateway

Cloudflare AI Gateway cung cấp một lớp trung gian cho ứng dụng AI với các chức năng như analytics, logging, caching, rate limiting và khả năng làm việc với nhiều AI Provider. Tài liệu sản phẩm cũng thể hiện xu hướng quản lý provider key và kiểm soát mức sử dụng tại Gateway.

**Điểm phù hợp để tham khảo**

- Điểm truy cập tập trung tới AI Provider.
- Logging và analytics cho request AI.
- Rate limiting và các cơ chế nâng cao độ ổn định.
- Giảm việc để ứng dụng tương tác trực tiếp với từng provider.

**Giới hạn so với bài toán Asymptotic**

- Tài liệu công khai được khảo sát chưa thiết lập đầy đủ mô hình đăng ký Agent gắn với một developer chịu trách nhiệm.
- Chưa thiết lập được chuỗi phân bổ hạn mức tổ chức → đội ngũ → lập trình viên → Agent giống phạm vi Asymptotic.
- Không nên kết luận Cloudflare không có budget control; cần mô tả chính xác cấp và cách áp dụng giới hạn theo tài liệu tại thời điểm viết báo cáo.

**Nguồn chính**

- Cloudflare AI Gateway: <https://developers.cloudflare.com/ai-gateway/>

### 5.2. Portkey AI Gateway

Portkey cung cấp AI Gateway với API thống nhất, routing, fallback, retry, rate limit, observability và các cơ chế kiểm soát usage hoặc chi phí. Sản phẩm thể hiện rõ hướng phát triển một control plane cho ứng dụng sử dụng nhiều mô hình.

**Điểm phù hợp để tham khảo**

- Trừu tượng hóa nhiều AI Provider sau một Gateway.
- Virtual key và quản lý truy cập.
- Routing, fallback và retry.
- Observability, usage và cost tracking.
- Giới hạn mức sử dụng hoặc budget theo cấu hình.

**Giới hạn so với bài toán Asymptotic**

- Trọng tâm là hạ tầng AI Gateway cho ứng dụng và nhóm phát triển.
- Chưa đủ bằng chứng công khai để khẳng định có ví tổ chức và chuỗi cấp hạn mức nội bộ như Asymptotic.
- Chưa đủ bằng chứng về mô hình reservation–settlement và idempotency tài chính theo từng request.

**Nguồn chính**

- Portkey AI Gateway: <https://portkey.ai/docs/product/ai-gateway>

### 5.3. LiteLLM Proxy

LiteLLM Proxy là sản phẩm gần với phạm vi kỹ thuật của Asymptotic nhất trong nhóm khảo sát. LiteLLM hỗ trợ proxy tới nhiều model/provider, virtual key, model access control, spend tracking, budget và rate limit ở nhiều cấp.

**Điểm phù hợp để tham khảo**

- API thống nhất cho nhiều provider.
- Virtual key thay cho việc phân phối trực tiếp provider key.
- Theo dõi spend và usage.
- Budget và rate limit.
- Quản lý user, team, key và model access.
- Có khái niệm Agent trong một số chức năng và tài liệu.

**Giới hạn hoặc điểm cần xác minh**

- Cần phân biệt khái niệm Agent của LiteLLM với “AI Agent bên ngoài được đăng ký và gắn trách nhiệm tài chính” trong Asymptotic.
- Cần xác minh cách LiteLLM kế thừa và kiểm tra budget giữa organization, team, user, key và Agent.
- Chưa đủ bằng chứng công khai để khẳng định LiteLLM cung cấp ví tổ chức, top-up, phân bổ hạn mức nội bộ, reservation và settlement theo mô hình của Asymptotic.
- Cần kiểm tra riêng cơ chế idempotency và đảm bảo không tạo giao dịch tài chính trùng.

**Nguồn chính**

- LiteLLM Proxy: <https://docs.litellm.ai/docs/simple_proxy>
- Virtual keys: <https://docs.litellm.ai/docs/proxy/virtual_keys>
- Users, teams, budgets và rate limits: <https://docs.litellm.ai/docs/proxy/users>

### 5.4. Helicone và các nền tảng observability

Helicone và các nền tảng tương tự tập trung vào logging, observability, token usage, cost analytics và đánh giá hoạt động của ứng dụng LLM.

**Giá trị khảo sát**

- Cho thấy request, token, latency và cost trace là nhu cầu phổ biến.
- Cung cấp mô hình tham khảo cho dashboard và phân tích usage.
- Hỗ trợ phân biệt “quan sát sau thực thi” với “kiểm soát tài chính trước thực thi”.

**Giới hạn so với bài toán Asymptotic**

- Observability không đồng nghĩa với authorization hoặc budget reservation.
- Việc ghi nhận chi phí sau khi provider xử lý không tự ngăn được chi tiêu vượt hạn mức.
- Không nên xem sản phẩm observability là đối thủ trực tiếp nếu không đảm nhiệm Gateway policy enforcement.

**Nguồn chính**

- Helicone documentation: <https://docs.helicone.ai/>

## 6. Nhóm hạ tầng tài chính và thanh toán cho AI Agent

### 6.1. Skyfire

Skyfire định vị gần với hạ tầng thanh toán cho AI Agent, bao gồm các khái niệm như Agent identity, Know Your Agent, wallet, payment và audit trail.

**Ý nghĩa đối với Asymptotic**

- Củng cố nhu cầu quản lý Agent như một workload có định danh riêng.
- Cho thấy credential và quyền tài chính của Agent cần có phạm vi.
- Cho thấy audit trail là thành phần cốt lõi của hạ tầng tài chính cho Agent.

**Khác biệt phạm vi**

- Skyfire hướng tới thanh toán và giao dịch của Agent.
- Asymptotic hướng tới kiểm soát ngân sách khi Agent tiêu thụ dịch vụ AI thông qua Gateway.

**Mức độ bằng chứng**

Thông tin công khai chủ yếu đến từ nhà cung cấp. Không nên suy diễn quy mô triển khai hoặc mức độ trưởng thành nếu không có dữ liệu độc lập.

**Nguồn chính**

- Skyfire: <https://skyfire.xyz/>

### 6.2. Nevermined

Nevermined tập trung vào payment protocol, metering và monetization cho AI service hoặc AI Agent.

**Ý nghĩa đối với Asymptotic**

- Usage metering là nền tảng để xác định chi phí.
- Dịch vụ dành cho Agent có xu hướng tính phí theo request hoặc mức sử dụng.
- Payment và access control có thể được tích hợp ở lớp giao thức.

**Khác biệt phạm vi**

- Nevermined hướng tới thanh toán hoặc kiếm tiền từ dịch vụ Agent.
- Asymptotic không xây dựng marketplace hoặc payment protocol tổng quát.

**Nguồn chính**

- Nevermined: <https://nevermined.ai/>

### 6.3. x402 và Stripe Machine Payments

x402 và Machine Payment Protocol sử dụng mô hình HTTP 402 để một dịch vụ thông báo điều kiện thanh toán, client cung cấp bằng chứng thanh toán và request được thực hiện lại.

**Ý nghĩa đối với Asymptotic**

- Cho thấy thanh toán tự động theo từng API request là một hướng phát triển thực tế.
- Nhấn mạnh nhu cầu chống replay, lũy đẳng và liên kết payment proof với request.
- Là ví dụ rõ ràng về việc authorization phải diễn ra trước khi tài nguyên được cung cấp.

**Khác biệt phạm vi**

- Đây là payment rail hoặc payment protocol.
- Asymptotic không yêu cầu Agent thanh toán trực tiếp cho từng AI Provider bằng stablecoin hoặc payment proof.
- Trong Asymptotic, Agent dùng API key do Gateway cấp; Gateway dùng provider credential nội bộ.

**Nguồn chính**

- x402: <https://www.x402.org/>
- Stripe Machine Payments: <https://docs.stripe.com/payments/machine.md>
- Stripe Machine Payment Protocol: <https://docs.stripe.com/payments/machine/mpp.md>

### 6.4. Google AP2, Mastercard Agent Pay và Visa Trusted Agent Protocol

Nhóm sản phẩm và tiêu chuẩn này tập trung vào định danh Agent, mandate, tokenized credential, user control và khả năng nhận biết Agent trong giao dịch.

**Ý nghĩa đối với Asymptotic**

- Agent không nên được coi là chủ thể pháp lý độc lập.
- Hoạt động của Agent cần gắn với cá nhân hoặc tổ chức chịu trách nhiệm.
- Phạm vi ủy quyền cần rõ ràng, có thể thu hồi và truy vết.
- Credential của Agent nên tách khỏi credential nền tảng hoặc nhà cung cấp.

**Khác biệt phạm vi**

- Các giải pháp này tập trung vào commerce hoặc card-network payment.
- Asymptotic áp dụng nguyên tắc tương tự cho quyền tiêu thụ dịch vụ AI, không triển khai card payment.

**Nguồn chính**

- Google AP2: <https://github.com/google-agentic-commerce/AP2>
- Mastercard Agent Pay: <https://www.mastercard.com/news/press/2025/april/mastercard-unveils-agent-pay-pioneering-agentic-payments-technology-to-power-commerce-in-the-age-of-ai/>
- Visa Trusted Agent Protocol: <https://developer.visa.com/capabilities/trusted-agent-protocol>

## 7. Ma trận so sánh tổng hợp

Ký hiệu:

- **Có:** tài liệu công khai thể hiện rõ chức năng.
- **Một phần:** có chức năng gần tương đương nhưng khác phạm vi hoặc mô hình.
- **CTXM:** chưa tìm thấy mô tả đủ rõ trong tài liệu công khai được khảo sát.
- **KPH:** không phù hợp với loại sản phẩm.
- **Mục tiêu:** chức năng thuộc phạm vi thiết kế của Asymptotic, không mặc nhiên có nghĩa đã triển khai production.

| Tiêu chí | Cloudflare AI Gateway | Portkey | LiteLLM | Helicone | Skyfire/Nevermined | Asymptotic |
|---|---|---|---|---|---|---|
| Gateway tới nhiều AI Provider | Có | Có | Có | Một phần | KPH | Mục tiêu |
| Gateway/virtual API key | Một phần | Có | Có | Một phần | Một phần | Mục tiêu |
| Quản lý provider credential nội bộ | Có | Có | Có | Một phần | KPH | Mục tiêu |
| Routing, retry hoặc fallback | Có | Có | Có | Một phần | KPH | Mục tiêu |
| Rate limit/quota/policy | Có | Có | Có | Một phần | Một phần | Mục tiêu |
| Theo dõi usage và cost | Có | Có | Có | Có | Có/Một phần | Mục tiêu |
| Budget hoặc spend limit | Có/Một phần | Có | Có | Một phần | Có/Một phần | Mục tiêu |
| Đăng ký AI Agent bên ngoài | CTXM | CTXM | Một phần | CTXM | Có | Mục tiêu |
| Agent gắn với developer quản lý | CTXM | CTXM | Một phần/CTXM | CTXM | Một phần | Mục tiêu |
| Chuỗi tổ chức → team → developer → Agent | CTXM | CTXM | Một phần | CTXM | CTXM | Mục tiêu |
| Ví tổ chức và top-up | CTXM | CTXM | CTXM | KPH | Có/Một phần | Mục tiêu |
| Phân bổ hạn mức nội bộ theo cấp | CTXM | CTXM | Một phần | KPH | Một phần | Mục tiêu |
| Kiểm tra/tạm giữ chi phí trước provider call | CTXM | Một phần/CTXM | Một phần/CTXM | KPH | Có/Một phần | Mục tiêu |
| Quyết toán theo usage thực tế | CTXM | CTXM | Một phần/CTXM | KPH | Có/Một phần | Mục tiêu |
| Idempotency tài chính | CTXM | CTXM | CTXM | KPH | Một phần/CTXM | Mục tiêu |
| Request–usage–cost–transaction trace | Một phần | Một phần | Một phần | Một phần | Một phần | Mục tiêu |

Ma trận này chỉ phản ánh tài liệu công khai đã khảo sát. “CTXM” không có nghĩa là sản phẩm chắc chắn không có chức năng.

## 8. Xu hướng phát triển của thị trường

### 8.1. Gateway trở thành điểm kiểm soát thay vì chỉ là proxy

AI Gateway đang phát triển từ lớp chuyển tiếp request thành control plane gồm:

- xác thực;
- provider abstraction;
- model access control;
- routing và fallback;
- rate limit và quota;
- logging, usage và cost;
- policy và budget.

Điều này xác nhận hướng kiến trúc Gateway của Asymptotic, nhưng đồng thời làm giảm giá trị khác biệt của các chức năng proxy cơ bản.

### 8.2. Virtual key thay thế việc phân phối provider key

Xu hướng chung là ứng dụng hoặc Agent dùng key do Gateway quản lý. Provider credential được giữ ở phía hạ tầng và không lộ cho workload bên ngoài.

Xu hướng này phù hợp trực tiếp với FR03, FR04, FR07 và NFR02 của Asymptotic.

### 8.3. Kiểm soát chi phí dịch chuyển gần thời điểm thực thi

Cost observability sau thực thi vẫn cần thiết, nhưng các sản phẩm đang bổ sung:

- budget;
- rate limit;
- spending limit;
- policy;
- approval hoặc mandate;
- credential có phạm vi.

Asymptotic nên nhấn mạnh kiểm soát trước provider call và quyết toán sau khi có usage thực tế.

### 8.4. Agent identity trở thành một lớp riêng

Agentic commerce và payment infrastructure đang tách:

- principal identity;
- Agent/workload identity;
- mandate;
- credential;
- transaction evidence.

Đối với Asymptotic, mô hình tương ứng là:

`Organization → Team → Developer → Registered AI Agent → Asymptotic API key → Gateway request`

### 8.5. Audit trail và reconciliation trở thành yêu cầu bắt buộc

Các hệ thống liên quan đều có xu hướng lưu:

- request hoặc intent;
- policy decision;
- credential hoặc key được sử dụng;
- provider/rail response;
- usage;
- cost;
- receipt hoặc transaction;
- trạng thái reconciliation.

Asymptotic cần duy trì liên kết truy vết thống nhất thay vì chỉ lưu log kỹ thuật hoặc số dư tổng hợp.

### 8.6. Idempotency và replay protection quan trọng hơn

Agent có thể tự động retry khi timeout hoặc mất response. Nếu Gateway không phân biệt request mới với request gửi lại, một hành động có thể tạo nhiều provider call hoặc nhiều giao dịch chi phí.

Vì vậy, idempotency là yêu cầu tài chính và khả năng phục hồi, không chỉ là tối ưu API.

## 9. Khoảng trống thị trường phù hợp với Asymptotic

Từ tài liệu công khai được khảo sát, chưa thể xác lập một sản phẩm trung lập thể hiện đầy đủ đồng thời các thành phần sau:

1. đăng ký AI Agent bên ngoài và gắn với developer chịu trách nhiệm;
2. ví tổ chức làm nguồn tiền;
3. chuỗi phân bổ hạn mức tổ chức → đội ngũ → lập trình viên → Agent;
4. Gateway-issued API key cho từng Agent;
5. provider credential nội bộ không lộ cho Agent;
6. kiểm tra identity, policy, quota và budget trước provider call;
7. tạm giữ hoặc ước lượng chi phí dự kiến;
8. quyết toán theo usage thực tế;
9. idempotency để tránh request và chi phí trùng;
10. trace thống nhất giữa request, Agent, usage, cost và transaction.

Khoảng trống này cần được trình bày thận trọng:

> Các sản phẩm hiện có đã giải quyết nhiều thành phần riêng lẻ như AI Gateway, virtual key, định tuyến, rate limit, budget, cost observability, Agent identity và payment credential. Asymptotic không tuyên bố phát minh các thành phần này. Định hướng của hệ thống là tích hợp chúng thành một Gateway kiểm soát trách nhiệm và chi phí sử dụng AI theo cơ cấu nội bộ của tổ chức và theo từng AI Agent bên ngoài.

## 10. Định vị đề xuất cho Asymptotic

### 10.1. Phát biểu định vị

**Asymptotic là cổng tài chính thời gian thực nằm giữa AI Agent bên ngoài và AI Provider trả phí, cho phép tổ chức quản lý Agent, cấp API key, phân bổ hạn mức, kiểm soát từng request trước khi phát sinh chi phí và truy vết usage–cost–transaction sau thực thi.**

### 10.2. Không nên định vị

- “Ví cho AI Agent”.
- “Ngân hàng tự động cho AI”.
- “Payment Gateway cho mọi giao dịch của AI Agent”.
- “Nền tảng tạo và vận hành AI Agent”.
- “Sản phẩm đầu tiên có budget cho AI”.
- “Hệ thống duy nhất theo dõi chi phí LLM”.

Các phát biểu này quá rộng hoặc dễ bị phản bác bởi sản phẩm hiện có.

### 10.3. Điểm khác biệt nên nhấn mạnh

- External Agent registration thay vì chỉ quản lý application/API key.
- Trách nhiệm chi phí theo tổ chức, đội ngũ, developer và Agent.
- Ví tổ chức kết hợp hạn mức nội bộ, không tạo ví độc lập cho từng cấp.
- Pre-execution control kết hợp post-execution settlement.
- Idempotent financial processing.
- Unified financial trace.

## 11. Đề xuất cấu trúc mục 2.1 “Khảo sát hiện trạng”

Mục 2.1 nên được viết theo cấu trúc:

### 2.1.1. Bối cảnh sử dụng API AI trả phí

- Đặc điểm tính phí.
- Request tự động và khó dự đoán của AI Agent.
- Vấn đề khi phân phối trực tiếp provider API key.

### 2.1.2. Tiêu chí khảo sát

- Trình bày rút gọn các tiêu chí C01–C13.
- Giải thích lý do lựa chọn sản phẩm.

### 2.1.3. Khảo sát nhóm AI Gateway

- Cloudflare AI Gateway.
- Portkey.
- LiteLLM.
- Có thể dùng Helicone làm đại diện cho observability.

Mỗi sản phẩm nên có:

- mô tả;
- chức năng chính;
- điểm phù hợp;
- giới hạn so với bài toán;
- ảnh giao diện hoặc kiến trúc từ nguồn chính thức;
- ngày truy cập;
- nguồn trích dẫn.

### 2.1.4. Khảo sát xu hướng tài chính cho AI Agent

- Skyfire hoặc Nevermined.
- x402/Stripe Machine Payments.
- AP2/Mastercard/Visa ở mức xu hướng định danh và ủy quyền.

Không nên dành dung lượng ngang bằng nhóm AI Gateway vì đây là nhóm tham khảo lân cận.

### 2.1.5. Bảng so sánh

Sử dụng bảng rút gọn từ ma trận tại Mục 7. Chỉ giữ các tiêu chí quan trọng nhất:

- multi-provider gateway;
- virtual key;
- usage/cost tracking;
- budget;
- external Agent registration;
- ownership hierarchy;
- organizational wallet;
- pre-request control;
- settlement;
- idempotency;
- unified trace.

### 2.1.6. Khoảng trống và bài toán

Kết luận từ bằng chứng khảo sát trước khi giới thiệu yêu cầu của Asymptotic. Không đưa chi tiết công nghệ triển khai vào phần này.

## 12. Nguyên tắc sử dụng nguồn trong báo cáo

- Ưu tiên documentation, specification, repository hoặc newsroom chính thức.
- Phân biệt chức năng production, preview, beta, pilot và công bố định hướng.
- Ghi ngày truy cập đối với website.
- Không dùng landing page marketing làm bằng chứng duy nhất cho đặc tính kỹ thuật quan trọng.
- Không suy luận “không có chức năng” chỉ vì chưa tìm thấy trong tài liệu.
- Dùng cụm “chưa tìm thấy mô tả trong tài liệu công khai được khảo sát” khi bằng chứng chưa đủ.
- Ảnh giao diện hoặc kiến trúc của sản phẩm phải ghi nguồn.
- Hình tự tổng hợp phải ghi “Nguồn: Tác giả xây dựng”.
- Các số liệu thị trường, lượng giao dịch hoặc mức độ phổ biến cần nguồn độc lập nếu được sử dụng.

## 13. Hạn chế của nghiên cứu

- Sản phẩm AI Gateway thay đổi nhanh; tính năng và gói thương mại có thể thay đổi sau mốc khảo sát.
- Một số chức năng chỉ có trong enterprise plan hoặc tài liệu không công khai.
- Thông tin về Skyfire, Nevermined, Catena và Payman chủ yếu là vendor claim.
- Ma trận chưa thay thế thử nghiệm sản phẩm thực tế.
- Nghiên cứu agentic payment ban đầu có phạm vi rộng hơn Asymptotic và được dùng chủ yếu để nhận diện xu hướng.
- Chưa thực hiện khảo sát độc lập về mức độ sử dụng sản phẩm tại Việt Nam.

## 14. Nguồn tham khảo chính

### AI Gateway và observability

1. Cloudflare, “AI Gateway documentation”: <https://developers.cloudflare.com/ai-gateway/>
2. Portkey, “AI Gateway documentation”: <https://portkey.ai/docs/product/ai-gateway>
3. LiteLLM, “Proxy Server”: <https://docs.litellm.ai/docs/simple_proxy>
4. LiteLLM, “Virtual Keys”: <https://docs.litellm.ai/docs/proxy/virtual_keys>
5. LiteLLM, “Users, Teams, Budgets and Rate Limits”: <https://docs.litellm.ai/docs/proxy/users>
6. Helicone, “Documentation”: <https://docs.helicone.ai/>

### Agent identity, mandate và machine payments

7. Google, “Agent Payments Protocol”: <https://github.com/google-agentic-commerce/AP2>
8. Mastercard, “Mastercard unveils Agent Pay,” 29/04/2025: <https://www.mastercard.com/news/press/2025/april/mastercard-unveils-agent-pay-pioneering-agentic-payments-technology-to-power-commerce-in-the-age-of-ai/>
9. Visa, “Trusted Agent Protocol”: <https://developer.visa.com/capabilities/trusted-agent-protocol>
10. Stripe, “Machine Payments”: <https://docs.stripe.com/payments/machine.md>
11. Stripe, “Machine Payment Protocol”: <https://docs.stripe.com/payments/machine/mpp.md>
12. x402: <https://www.x402.org/>
13. Skyfire: <https://skyfire.xyz/>
14. Nevermined: <https://nevermined.ai/>

### Tiêu chuẩn kiểm soát truy cập và truy vết

15. IETF, “OAuth 2.0 Security Best Current Practice,” RFC 9700, 01/2025: <https://www.rfc-editor.org/rfc/rfc9700.html>
16. IETF, “OAuth 2.0 Rich Authorization Requests,” RFC 9396: <https://www.rfc-editor.org/rfc/rfc9396.html>
17. NIST, “Guide to Attribute Based Access Control,” SP 800-162: <https://csrc.nist.gov/pubs/sp/800/162/upd2/final>
18. NIST, “Zero Trust Architecture,” SP 800-207: <https://csrc.nist.gov/pubs/sp/800/207/final>

## 15. Tài liệu đầu vào

Tài liệu tổng hợp này sử dụng:

- `report/report_support_documentation/governance/project_source_of_truth.md`;
- mục 2.1 hiện tại trong `report/chapter_2.tex`;
- `report/research/market/ai-agent-financial-gateway-research-2026-06-16.md`;
- phương pháp trình bày khảo sát sản phẩm tham khảo từ `report/report_support_documentation/references/sample_reports/DOAN1-NGUYENHUYENTRANG.pdf`;
- tài liệu công khai của các sản phẩm và tiêu chuẩn được liệt kê tại Mục 14.
