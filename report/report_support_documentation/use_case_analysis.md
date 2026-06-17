# Hướng dẫn phân tích và đặc tả Use Case cho Đồ án 1

## 1. Mục đích tài liệu

Tài liệu này tổng hợp và hệ thống hóa nội dung từ:

1. Object Management Group (OMG), *Unified Modeling Language, Version 2.5.1*, Clause 18 - UseCases, trang in 638-649;
2. [Use Case Diagram và 5 sai lầm thường gặp](https://thinhnotes.com/chuyen-nghe-ba/use-case-diagram-va-5-sai-lam-thuong-gap/);
3. [Viết đặc tả Use Case sao đơn giản nhưng hiệu quả?](https://thinhnotes.com/chuyen-nghe-ba/viet-dac-ta-use-case-sao-don-gian-nhung-hieu-qua/);
4. Trịnh Thành Trung, *Bài 13 - Tổng quan về UML*, tài liệu bài giảng OOAD, 2016.

Mục tiêu là tạo một tài liệu làm việc có thể sử dụng trực tiếp để:

- xác định đúng Actor và Use Case;
- xây dựng hoặc rà soát Use Case Diagram;
- viết Use Case Specification nhất quán;
- phân biệt luồng chính, luồng thay thế và luồng ngoại lệ;
- kiểm tra sự nhất quán giữa yêu cầu, Use Case, biểu đồ tuần tự, biểu đồ hoạt động và Test Case;
- làm chuẩn phương pháp khi xây dựng hoặc rà soát mô hình Use Case của Đồ án 1.

Thứ tự ưu tiên khi có khác biệt:

1. chuẩn OMG UML 2.5.1;
2. quy định trình bày của giảng viên và Đồ án 1, nếu không làm sai ngữ nghĩa UML;
3. tài liệu bài giảng OOAD và hướng dẫn thực hành từ hai bài viết Thinhnotes;
4. các hướng dẫn thực hành theo phong cách Cockburn/Fowler và kinh nghiệm Business Analysis, dùng để cải thiện chất lượng văn bản đặc tả, không thay thế chuẩn UML.

Tài liệu bài giảng OOAD và hai bài viết Thinhnotes cung cấp cách giải thích và ví dụ thực hành, không phải đặc tả chuẩn UML chính thức. Tài liệu OMG UML 2.5.1 là căn cứ chính cho ký pháp và ngữ nghĩa Use Case Diagram.

---

## 2. Bản chất của Use Case

### 2.1. Định nghĩa thực hành

Use Case là kỹ thuật mô tả:

- sự tương tác giữa người dùng và hệ thống; hoặc
- sự tương tác giữa một hệ thống bên ngoài và hệ thống đang phân tích;
- trong một phạm vi cụ thể;
- nhằm đạt một mục tiêu có giá trị đối với Actor.

Một Use Case tốt phải trả lời được ba câu hỏi:

1. **Ai hoặc hệ thống nào tương tác?**
2. **Tương tác với hệ thống trong phạm vi nào?**
3. **Actor muốn đạt kết quả gì?**

Use Case thể hiện yêu cầu theo góc nhìn bên ngoài hệ thống. Nó tập trung vào **What** - hệ thống cung cấp khả năng gì - thay vì **How** - hệ thống cài đặt bằng thuật toán, cơ sở dữ liệu hay framework nào.

### 2.2. Hai hình thức biểu diễn

Use Case thường tồn tại ở hai dạng bổ sung cho nhau:

- **Use Case Diagram:** cho cái nhìn tổng quan về Actor, mục tiêu và phạm vi hệ thống.
- **Use Case Specification:** mô tả chi tiết điều kiện, trình tự tương tác, kết quả và các nhánh xử lý của từng Use Case.

Mỗi hình oval có ý nghĩa nghiệp vụ trên Use Case Diagram nên có một bản đặc tả tương ứng, trừ các Use Case tổng quát chỉ được dùng để nhóm và đã có quy ước rõ ràng.

### 2.3. Giá trị trong quy trình phân tích

Use Case có các giá trị chính:

- diễn đạt Functional Requirement theo góc nhìn người dùng cuối;
- sử dụng ngôn ngữ tự nhiên, giúp stakeholder dễ xác nhận;
- chia phạm vi hệ thống theo phân hệ hoặc cụm tính năng;
- kiểm tra mức độ bao phủ yêu cầu;
- làm cầu nối từ yêu cầu tổng quát sang Sequence Diagram và Activity Diagram;
- hỗ trợ hình thành Epic và User Story;
- làm đầu vào để thiết kế Test Case và Acceptance Criteria.

### 2.4. Vị trí của Use Case trong OOAD

OOAD chuyển yêu cầu của bài toán thành các mô hình phân tích và thiết kế trước khi lập trình. Các mô hình cung cấp những góc nhìn khác nhau về cùng một hệ thống; vì vậy, không nên buộc Use Case Diagram mô tả đồng thời mục tiêu, thuật toán và tương tác giữa các đối tượng.

Phân vai các tài liệu và biểu đồ như sau:

- **Use Case Diagram** mô tả Actor, mục tiêu chức năng và phạm vi hệ thống, tập trung vào **What**, không mô tả **How**.
- **Use Case Specification** mô tả luồng sự kiện, điều kiện, kết quả và quy tắc nghiệp vụ của một Use Case.
- **Activity Diagram** mô tả luồng hoạt động và các nhánh xử lý được thực hiện trong Use Case.
- **Sequence Diagram** mô tả các đối tượng tham gia và trình tự thông điệp trong một kịch bản của Use Case.
- **Class Diagram** mô tả cấu trúc tĩnh phục vụ hiện thực hóa hệ thống.

Khi một nội dung biểu diễn thứ tự các bước như xác thực, kiểm tra ngân sách, gọi nhà cung cấp và quyết toán, nội dung đó thuộc flow của đặc tả, Activity Diagram hoặc Sequence Diagram. Việc đổi các bước này thành các hình oval không làm chúng trở thành Use Case.

### 2.5. Quy tắc chốt khi phân tích Use Case

Khi rà soát một Use Case hoặc Use Case Diagram, dùng các quy tắc sau để tránh sai mức trừu tượng:

- Use Case Diagram thể hiện **mục tiêu của Actor và phạm vi hệ thống**, không thể hiện thuật toán hoặc trình tự xử lý.
- Use Case Specification mô tả **cuộc trao đổi giữa Actor và hệ thống** bằng văn bản có thể đọc, kiểm thử và đối chiếu với yêu cầu.
- UML chuẩn hóa Actor, Use Case, Boundary, `include`, `extend` và generalization, nhưng không bắt buộc một mẫu đặc tả văn bản duy nhất.
- Một Use Case chỉ nên được tách ra khi nó là mục tiêu có giá trị độc lập đối với Actor hoặc là hành vi dùng chung đủ quan trọng để đặc tả riêng.
- Không biến bước kỹ thuật thành Use Case chỉ vì bước đó phức tạp. Nếu bước đó không phải mục tiêu của Actor, hãy đưa vào flow, Business Rule, Activity Diagram, Sequence Diagram hoặc chương thiết kế.
- Mỗi nhánh trong đặc tả phải cho biết Use Case tiếp tục ở bước nào, kết thúc thành công hay kết thúc không thành công.

### 2.6. Áp dụng riêng cho AI Agent Financial Gateway

Với đề tài này, UC01 -- `Thực hiện yêu cầu AI qua Gateway` là một Use Case cấp actor-goal. Các nội dung như xác thực API key, kiểm tra idempotency, kiểm tra ngân sách, chọn provider/model, streaming, quyết toán chi phí và ghi trace là các bước hoặc quy tắc bên trong UC01, không tự động trở thành các Use Case độc lập.

Chỉ tách một nội dung trong UC01 thành Use Case riêng khi thỏa ít nhất một điều kiện:

- Actor có thể chủ động khởi tạo nội dung đó như một mục tiêu riêng;
- nội dung đó được nhiều Use Case khác tái sử dụng và cần đặc tả riêng;
- nội dung đó tạo ra kết quả có giá trị độc lập, quan sát được từ bên ngoài hệ thống.

Ví dụ:

- `Quản lý API key của Agent` là UC riêng vì lập trình viên hoặc quản trị viên có mục tiêu tạo, xem trạng thái và thu hồi khóa.
- `Kiểm tra API key` không phải UC riêng nếu nó chỉ là một bước bắt buộc trong UC01.
- `Theo dõi giao dịch, usage, cost và trace` là UC riêng vì người dùng có quyền chủ động tra cứu dữ liệu sau khi request phát sinh.
- `Ghi ledger entry` không phải UC riêng nếu không có Actor bên ngoài khởi tạo nó như một mục tiêu độc lập.

---

## 3. Thành phần của Use Case Diagram

### 3.1. Actor

Actor là một **vai trò bên ngoài ranh giới hệ thống** có tương tác với hệ thống. Actor có thể là:

- người dùng;
- vai trò nghiệp vụ;
- hệ thống phần mềm khác;
- dịch vụ bên thứ ba;
- thiết bị hoặc tác nhân tự động.

Hình người trong UML không có nghĩa Actor bắt buộc phải là con người.

Các câu hỏi dùng để tìm Actor:

- Ai sử dụng hệ thống?
- Ai quản trị, cấu hình, vận hành hoặc bảo trì hệ thống?
- Hệ thống nào gửi dữ liệu hoặc yêu cầu vào?
- Hệ thống nào nhận dữ liệu hoặc kết quả từ hệ thống?
- Ai tạo dữ liệu?
- Ai cần dữ liệu đầu ra?
- Ai khởi tạo sự kiện nghiệp vụ?
- Ai nhận giá trị sau khi Use Case hoàn thành?

Quy tắc đặt tên Actor:

- dùng danh từ hoặc cụm danh từ;
- đặt theo vai trò, không đặt theo tên cá nhân;
- tránh dùng động từ;
- phân biệt vai trò nghiệp vụ với chức danh kỹ thuật nếu quyền và mục tiêu khác nhau.

Ví dụ phù hợp:

- `AI Agent`;
- `Lập trình viên`;
- `Quản trị viên tổ chức`;
- `Nhà cung cấp dịch vụ AI`;
- `Nhà cung cấp dịch vụ thanh toán`.

### 3.2. Use Case

Use Case được biểu diễn bằng hình oval và thể hiện một đơn vị chức năng hữu ích mà hệ thống cung cấp. Khi hoàn thành, Use Case phải tạo ra kết quả quan sát được và có giá trị đối với Actor hoặc stakeholder liên quan.

Tên Use Case nên:

- dùng cấu trúc **Động từ + Danh từ/Đối tượng**;
- ngắn, rõ và mang ý nghĩa nghiệp vụ;
- mô tả kết quả có giá trị;
- tránh thuật ngữ cài đặt mà stakeholder không hiểu;
- tránh câu bị động và tên quá dài.

Ví dụ:

- `Thực hiện yêu cầu AI`;
- `Nạp tiền`;
- `Phân bổ ngân sách`;
- `Thiết lập hạn mức tác nhân`;
- `Thu hồi khóa truy cập`;
- `Theo dõi giao dịch`.

Tên như `Dữ liệu`, `Màn hình ngân sách`, `API`, `Single`, `Double` không phải mục tiêu tương tác và thường không phải Use Case.

### 3.3. Communication Link

Communication Link là đường nối Actor với Use Case, cho biết Actor có tham gia tương tác trong Use Case.

Đường nối không tự mô tả trình tự xử lý. Trình tự phải được thể hiện trong đặc tả, Sequence Diagram hoặc Activity Diagram.

### 3.4. System Boundary

System Boundary xác định phần nào thuộc hệ thống đang xây dựng và phần nào là tác nhân bên ngoài.

Boundary có thể đại diện cho:

- toàn bộ hệ thống;
- một phân hệ;
- một module lớn;
- một phạm vi chức năng được chọn để phân tích.

Tên Boundary phải thể hiện đúng phạm vi. Trong Đồ án 1, Boundary tổng quát có thể là `AI Agent Financial Gateway`.

Boundary giúp:

- ngăn việc đưa hành vi của hệ thống ngoài vào như chức năng nội bộ;
- phân nhóm Use Case;
- giảm độ rối của sơ đồ;
- làm rõ phạm vi đồ án.

### 3.5. Relationships

Ba quan hệ thường gặp là `include`, `extend` và `generalization`.

UML không định nghĩa relationship `<<precondition>>` giữa các Use Case. Pre-Condition là thuộc tính văn bản của Use Case Specification. Nếu một Use Case tạo ra trạng thái cần thiết cho Use Case khác, thể hiện sự phụ thuộc đó bằng Pre/Post-Condition và traceability, không tự tạo stereotype mũi tên trên diagram.

#### 3.5.1. Include

`A <<include>> B` có nghĩa hành vi của B được chèn vào hành vi của A tại một vị trí xác định. Khi luồng của A đi tới vị trí này, toàn bộ hành vi được include phải thực hiện trước khi A tiếp tục.

Đây là quan hệ có hướng. Mũi tên nét đứt, đầu tên rỗng hướng từ Use Case bao gồm tới Use Case được bao gồm:

`A --<<include>>--> B`

Theo OMG UML 2.5.1, mục đích chính của `include` là trích phần hành vi chung của từ hai Use Case trở lên để tái sử dụng. Use Case bao gồm phụ thuộc vào phần hành vi được include; Use Case được include không phụ thuộc ngược lại.

Chỉ nên tách một hành vi thành Use Case `include` khi:

- hành vi có ý nghĩa chức năng rõ ràng trong mô hình;
- hành vi có thể được đặc tả như một đơn vị hoàn chỉnh;
- hành vi đủ lớn hoặc đủ phức tạp để cần mô tả riêng; và thường
- hành vi được nhiều Use Case tái sử dụng.

Không nên biến mọi bước nhỏ trong flow thành một Use Case `include`, vì sẽ làm sơ đồ trở thành bản phân rã thuật toán.

Các hành vi như `Xác thực yêu cầu`, `Kiểm tra idempotency`, `Kiểm tra ngân sách`, `Tạm giữ ngân sách`, `Gọi provider` hoặc `Quyết toán giao dịch` không tự động trở thành Use Case con. Nếu chúng chỉ là các bước nội bộ để UC01 đạt mục tiêu, phải đặt chúng trong Basic/Alternative/Exception Flow, Activity Diagram hoặc Sequence Diagram.

Chỉ cân nhắc tách một hành vi dùng chung khi nhiều Use Case thực sự gọi cùng một khả năng có ranh giới, contract và đặc tả độc lập. Việc “được nhiều nơi gọi” chưa đủ nếu hành vi vẫn chỉ là chi tiết kỹ thuật nội bộ.

#### 3.5.2. Extend

`B <<extend>> A` có nghĩa B bổ sung hành vi vào A tại một hoặc nhiều extension point của A. A vẫn có ý nghĩa hoàn chỉnh nếu B không xảy ra; B có thể không có ý nghĩa khi đứng độc lập.

Đây là quan hệ có hướng. Mũi tên nét đứt, đầu tên rỗng hướng từ Use Case mở rộng về Use Case nền tảng:

`B --<<extend>>--> A`

Theo metamodel UML 2.5.1, một quan hệ `extend` phải tham chiếu ít nhất một **extension point** thuộc Use Case nền tảng. Điều kiện kích hoạt là tùy chọn; nếu có, điều kiện được đánh giá tại extension point đầu tiên.

Ví dụ điều kiện kích hoạt:

- khi người dùng yêu cầu xuất báo cáo;
- khi ngân sách không đủ cho mô hình ban đầu nhưng có mô hình thay thế;
- khi giao dịch cần đối soát do chưa nhận được callback.

Không dùng `extend` chỉ vì một bước xảy ra sau Use Case khác. Quan hệ này phải thể hiện hành vi tùy chọn hoặc có điều kiện gắn với Use Case cơ sở.

#### 3.5.3. Generalization

Generalization thể hiện quan hệ cha-con giữa:

- Actor với Actor; hoặc
- Use Case với Use Case.

Thành phần con kế thừa quan hệ và hành vi chung từ thành phần cha.

Ví dụ Actor:

- `Người dùng tổ chức` là Actor tổng quát;
- `Lập trình viên` và `Quản trị viên tổ chức` là các vai trò chuyên biệt.

Ví dụ Use Case:

- `Thanh toán` là Use Case tổng quát;
- `Thanh toán qua thẻ` và `Thanh toán qua ví điện tử` là các Use Case chuyên biệt.

Trong UML, generalization được vẽ bằng **đường liền với đầu tam giác rỗng hướng về thành phần cha**.

### 3.6. Các kết luận trực tiếp từ OMG UML 2.5.1

Các quy tắc sau được rút ra từ Clause 18, trang in 638-649:

- Use Case mô tả hành vi hệ thống cung cấp mà không viện dẫn cấu trúc cài đặt nội bộ.
- Mỗi Use Case là một đơn vị chức năng hữu ích, tạo ra kết quả quan sát được và có giá trị cho Actor hoặc stakeholder.
- Một Use Case có thể chứa luồng cơ bản, biến thể, hành vi ngoại lệ và lỗi trong cùng một đơn vị hành vi.
- Actor biểu diễn một **vai trò**, không phải một cá nhân, thiết bị hay hệ thống vật lý cụ thể.
- Hành vi chi tiết có thể được mô tả bằng văn bản tự nhiên, Activity, Interaction/Sequence, State Machine, hoặc pre/post-condition.
- Hình chữ nhật System Boundary có thể được dùng để biểu diễn Subject; Use Case được đặt trực quan bên trong nhưng việc đặt đó không tự tạo quan hệ sở hữu trong metamodel.
- `include` chèn hành vi được tái sử dụng vào Use Case bao gồm và có mũi tên hướng về Use Case được include.
- `extend` chèn hành vi bổ sung tại extension point và có mũi tên hướng từ Use Case mở rộng về Use Case nền tảng.
- Một Use Case nền tảng trong quan hệ `extend` phải hoàn chỉnh và có ý nghĩa khi không có phần mở rộng.
- UML không định nghĩa khái niệm tùy ý “Use Case con” cho các bước xử lý tuần tự. Nếu nội dung là thuật toán hoặc trình tự nội bộ, phải mô tả trong specification, Activity Diagram hoặc Sequence Diagram.

Hệ quả áp dụng cho Đồ án 1:

- `Xác thực yêu cầu`, `Kiểm soát lũy đẳng`, `Chọn mô hình`, `Tạm giữ ngân sách`, `Gọi provider`, `Quyết toán` và `Ghi execution trace` không phải Use Case riêng nếu chúng chỉ là các bước để hoàn thành một yêu cầu AI.
- Pre-Condition nằm trong Use Case Specification; không vẽ mũi tên `<<precondition>>`.
- Alternative/Exception Flow không tự động trở thành Use Case mở rộng. Chỉ dùng `extend` khi hành vi mở rộng thực sự là một Use Case và được chèn tại extension point của Use Case nền tảng.

---

## 4. Năm sai lầm thường gặp khi vẽ Use Case Diagram

### 4.1. Đặt tên không đúng

Dấu hiệu:

- Actor được đặt bằng động từ;
- Use Case là cụm danh từ;
- tên chứa chi tiết kỹ thuật;
- tên quá dài;
- tên không cho biết mục tiêu của Actor;
- lạm dụng từ `Quản lý`.

Cách sửa:

- Actor dùng danh từ chỉ vai trò;
- Use Case dùng `Động từ + Đối tượng`;
- thay từ kỹ thuật bằng ngôn ngữ nghiệp vụ;
- kiểm tra tên bằng câu: “Actor muốn [tên Use Case] để đạt mục tiêu gì?”.

### 4.2. Biến Use Case Diagram thành sơ đồ phân rã chức năng

Use Case Diagram không phải Function Decomposition Diagram.

Các tên như `Quản lý A`, `Quản lý B` thường quá rộng nếu không làm rõ người dùng đạt gì. Use Case phải xuất phát từ mục tiêu của Actor, không chỉ từ danh sách module hoặc bảng dữ liệu.

Trước khi vẽ cần biết:

- End-user muốn làm gì?
- Vì sao họ cần làm việc đó?
- Hệ thống phải trao đổi dữ liệu với hệ thống ngoài nào?
- Mỗi Use Case có cùng mức trừu tượng hay không?

Không nên đặt cạnh nhau một Use Case cấp nghiệp vụ rất cao như `Quản lý tài chính` với một Use Case chi tiết như `Thu hồi khóa truy cập`.

### 4.3. Sơ đồ quá rối

Nguyên nhân phổ biến:

- xác định sai nhiều thành phần thành Use Case;
- tên Use Case không đúng;
- không có Boundary;
- không ghi điều kiện cho `extend`;
- không sử dụng hợp lý `include`, `extend` hoặc generalization;
- đưa quá nhiều Use Case vào một hình;
- đường nối chồng chéo;
- liệt kê mọi thao tác CRUD.

Khuyến nghị thực hành từ nguồn là một sơ đồ nên có khoảng trên dưới 10 Use Case. Đây là heuristic về khả năng đọc, không phải giới hạn UML bắt buộc.

Khi sơ đồ lớn:

- tạo một sơ đồ tổng quát;
- tách sơ đồ con theo phân hệ;
- giữ UC ID xuyên suốt giữa các sơ đồ;
- dùng Boundary để phân nhóm;
- có thể lặp lại cùng Actor ở hai vị trí để tránh đường nối cắt nhau.

### 4.4. Mô tả CRUD quá chi tiết

Không nên mặc định tạo bốn Use Case `Thêm`, `Xem`, `Sửa`, `Xóa` cho mọi thực thể. Cách này tạo nhiều hình oval nhưng ít giá trị phân tích.

Có hai cách xử lý:

1. Ghi một quy ước chung về các thao tác CRUD và quyền áp dụng.
2. Dùng Use Case tổng hợp `Quản lý X`, đồng thời ghi rõ phạm vi thao tác được hỗ trợ.

Tuy nhiên, nếu từng thao tác có:

- mục tiêu nghiệp vụ khác nhau;
- Actor khác nhau;
- quy tắc bảo mật khác nhau;
- flow phức tạp;
- hậu điều kiện quan trọng;

thì vẫn nên tách thành Use Case riêng.

Đặc biệt phải phân biệt CRUD kỹ thuật với quyền của End-user. Hệ thống có thể tự tạo dữ liệu do đồng bộ từ hệ thống ngoài, nhưng người dùng không nhất thiết có quyền tạo thủ công dữ liệu đó.

### 4.5. Thiếu thẩm mỹ và tính nhất quán

Checklist trình bày:

- kích thước các Use Case tương đối đồng đều;
- mỗi Use Case có UC ID;
- font chữ và cách viết hoa nhất quán;
- đường nối không chồng chéo;
- quan hệ đúng kiểu nét và đúng hướng mũi tên;
- Actor nằm ngoài Boundary;
- Use Case nằm trong Boundary;
- màu sắc vừa đủ và có ý nghĩa;
- ưu tiên khả năng đọc hơn trang trí;
- sơ đồ trả lời câu hỏi **What**, không trình bày chi tiết **How**.

---

## 5. Cách vẽ Use Case Diagram tổng quan và chi tiết

Một hệ thống có nhiều Actor và chức năng không nên cố trình bày mọi thứ trong một sơ đồ duy nhất. Nên tổ chức mô hình Use Case thành ít nhất hai cấp:

1. **Use Case Diagram tổng quan**: thể hiện toàn cảnh hệ thống.
2. **Use Case Diagram chi tiết**: tập trung vào một phân hệ hoặc một nhóm mục tiêu có liên quan.

Hai loại sơ đồ dùng chung một mô hình Actor và Use Case. “Chi tiết” ở đây là chi tiết hơn về phạm vi và quan hệ, không phải tách các bước xử lý của một Use Case thành nhiều oval. Sơ đồ chi tiết không được tạo ra một mô hình nghiệp vụ mâu thuẫn với sơ đồ tổng quan.

### 5.1. Use Case Diagram tổng quan

#### Mục đích

Sơ đồ tổng quan giúp người đọc trả lời nhanh:

- hệ thống đang phân tích là gì;
- hệ thống phục vụ những nhóm Actor nào;
- các khả năng nghiệp vụ lớn của hệ thống là gì;
- hệ thống tích hợp với những hệ thống ngoài nào;
- phạm vi nào thuộc và không thuộc đồ án.

Đối tượng đọc chính:

- giảng viên;
- stakeholder nghiệp vụ;
- thành viên mới của dự án;
- người cần hiểu phạm vi trước khi đọc chi tiết.

#### Mức trừu tượng

Use Case trong sơ đồ tổng quan phải ở mức **mục tiêu người dùng hoặc khả năng nghiệp vụ lớn**, không phải bước xử lý nội bộ.

Ví dụ phù hợp:

- `Thực hiện yêu cầu AI`;
- `Nạp tiền`;
- `Phân bổ ngân sách`;
- `Thiết lập hạn mức tác nhân`;
- `Quản lý khóa truy cập`;
- `Theo dõi giao dịch`;
- `Quản lý thành viên`.

Ví dụ quá chi tiết đối với sơ đồ tổng quan:

- `Kiểm tra idempotency key`;
- `Khóa bản ghi ví`;
- `Tính số token đầu vào`;
- `Ghi cache`;
- `Xác thực chữ ký callback`.

Các hành vi trên thường là bước trong flow hoặc nội dung cần mô tả bằng Activity/Sequence Diagram. Chỉ dùng `include` nếu hành vi thực sự là phần dùng chung được trích ra để tái sử dụng theo quy tắc tại mục 3.5.1.

#### Thành phần nên có

- một System Boundary đại diện toàn bộ hệ thống;
- tất cả nhóm Actor chính;
- các hệ thống ngoài quan trọng;
- các Use Case cấp cao;
- Communication Link chính;
- generalization giữa các Actor nếu giúp giảm lặp;
- một số ít `include`/`extend` thật sự cần để hiểu phạm vi.

#### Thành phần nên hạn chế

- quá nhiều mục tiêu chi tiết;
- mọi thao tác CRUD;
- chi tiết giao diện;
- quyết định kiến trúc;
- trình tự xử lý;
- quan hệ `include`/`extend` dày đặc;
- điều kiện nghiệp vụ dài viết trực tiếp trên hình.

#### Quy trình vẽ

1. Viết một câu mô tả phạm vi hệ thống.
2. Liệt kê Actor bên ngoài.
3. Với từng Actor, liệt kê các mục tiêu có giá trị.
4. Gộp các mục tiêu trùng nghĩa.
5. Chuẩn hóa tên theo `Động từ + Đối tượng`.
6. Chọn các Use Case cùng mức trừu tượng.
7. Vẽ Boundary và đặt Use Case vào trong.
8. Đặt Actor ngoài Boundary và nối Communication Link.
9. Thêm UC ID.
10. Chỉ thêm relationship khi nó làm sơ đồ rõ hơn.
11. Tách sơ đồ nếu số lượng hình oval hoặc đường nối làm giảm khả năng đọc.

#### Quy tắc đọc lại trước khi chốt sơ đồ

Sau khi vẽ, đọc lại từng oval bằng câu:

> Actor dùng hệ thống để `[tên Use Case]` và nhận được `[kết quả quan sát được]`.

Nếu không thể điền được kết quả quan sát được, oval đó có khả năng là bước kỹ thuật hoặc chi tiết thiết kế. Khi đó cần chuyển nó sang đặc tả, Activity Diagram, Sequence Diagram hoặc Business Rule.

Với các Use Case quản trị có tên bắt đầu bằng `Quản lý`, cần ghi rõ phạm vi quản lý trong đặc tả hoặc tách thành các mục tiêu con nếu sơ đồ chi tiết yêu cầu. Không để `Quản lý ...` trở thành nhãn quá rộng nhưng không biết Actor thực hiện được hành động nào.

#### Tiêu chí hoàn thành

Sơ đồ tổng quan đạt yêu cầu khi người chưa đọc tài liệu vẫn có thể giải thích:

- hệ thống làm gì;
- ai sử dụng;
- hệ thống ngoài nào tham gia;
- các nhóm chức năng lớn;
- phạm vi đồ án.

### 5.2. Use Case Diagram chi tiết

#### Mục đích

Sơ đồ chi tiết tập trung vào một:

- phân hệ;
- Actor;
- quy trình nghiệp vụ;
- Use Case cấp cao;
- nhóm chức năng có quan hệ chặt chẽ.

Nó giúp làm rõ:

- các mục tiêu sử dụng có quan hệ chặt chẽ trong phạm vi đã chọn;
- hành vi bắt buộc được tái sử dụng;
- hành vi mở rộng có điều kiện;
- Actor hỗ trợ;
- ranh giới của một module nghiệp vụ.

Ví dụ các sơ đồ chi tiết có thể có:

- quản lý request AI và quyết toán;
- quản lý ngân sách;
- quản lý khóa và tác nhân;
- nạp tiền và đối soát;
- báo cáo và truy vết.

#### Mức trừu tượng

Các Use Case trên cùng sơ đồ chi tiết phải có mức chi tiết tương đương.

Ví dụ với nhóm `Quản lý khóa truy cập`:

- `Tạo khóa truy cập`;
- `Xem danh sách khóa`;
- `Thu hồi khóa truy cập`;
- `Xoay vòng khóa truy cập`.

Không nên đặt cạnh nhóm trên một bước kỹ thuật như `Băm khóa bằng SHA-256`, vì đó là cách cài đặt chứ không phải mục tiêu của Actor.

Tương tự, không phân rã một Use Case theo chuỗi bước xử lý rồi gắn mã `UC01.1`, `UC01.2`, ... nếu các bước đó không phải mục tiêu hoàn chỉnh. Cách biểu diễn này thực chất là Activity Diagram hoặc Sequence Diagram được vẽ bằng ký pháp Use Case.

#### Cách liên kết với sơ đồ tổng quan

Có hai cách tổ chức:

**Cách A - Use Case tổng quan là nhóm điều hướng**

- sơ đồ tổng quan có `UC05 - Quản lý khóa truy cập`;
- sơ đồ chi tiết phân rã thành `UC05.1`, `UC05.2`, `UC05.3`;
- tài liệu ghi rõ UC05 là Use Case tổng hợp.

**Cách B - Giữ nguyên các Use Case nghiệp vụ**

- sơ đồ tổng quan vẫn hiển thị các Use Case quan trọng nhất;
- sơ đồ chi tiết lặp lại các Use Case đó và bổ sung Actor/relationship;
- UC ID và tên không thay đổi.

Cách B thường ít gây nhầm lẫn hơn nếu Use Case tổng quan đã là mục tiêu hoàn chỉnh. Cách A phù hợp khi Use Case cha chỉ dùng để nhóm nhiều mục tiêu con.

#### Sử dụng `include`

Trong sơ đồ chi tiết, dùng `include` khi hành vi:

- bắt buộc đối với Use Case chính;
- có ý nghĩa riêng;
- đủ phức tạp để đặc tả;
- được nhiều Use Case tái sử dụng.

Ví dụ:

- nhiều Use Case quản trị include `Xác thực quyền truy cập`;
- nhiều giao dịch tài chính include `Ghi nhận bút toán`;

Không tách một bước thành `include` chỉ để sơ đồ trông chi tiết hơn.

#### Sử dụng `extend`

Dùng `extend` khi:

- Use Case cơ sở vẫn hoàn chỉnh nếu phần mở rộng không xảy ra;
- phần mở rộng chỉ xảy ra trong điều kiện cụ thể;
- Use Case cơ sở khai báo extension point nơi hành vi được chèn;
- hành vi mở rộng có giá trị hoặc flow riêng.

Ví dụ:

- `Xuất báo cáo` extend `Theo dõi giao dịch` khi người dùng yêu cầu tải dữ liệu;
- `Đối soát giao dịch` extend `Nạp tiền` khi chưa nhận được callback sau thời hạn.

Ghi extension point của Use Case nền tảng và, nếu có, điều kiện kích hoạt gần relationship hoặc trong đặc tả.

#### Dùng Boundary trong sơ đồ chi tiết

Sơ đồ chi tiết có thể:

- giữ Boundary toàn hệ thống; hoặc
- dùng Boundary của phân hệ nếu tên phạm vi được ghi rõ.

Ví dụ:

- `AI Agent Financial Gateway`;
- `Phân hệ quản lý ngân sách`;
- `Phân hệ quản lý truy cập`.

Không vẽ module nội bộ thành Actor chỉ vì nó nằm ngoài Boundary của phân hệ. Nếu module vẫn thuộc hệ thống đang xây dựng, cần ghi rõ đây là sơ đồ ở cấp component/module, hoặc sử dụng loại biểu đồ khác phù hợp hơn.

#### Quy trình vẽ

1. Chọn phạm vi chi tiết.
2. Xác định Use Case cấp cao hoặc nhóm chức năng cần phân tích.
3. Đọc đặc tả và tìm các mục tiêu con có giá trị độc lập.
4. Tìm hành vi bắt buộc có ý nghĩa chức năng và được tái sử dụng để cân nhắc `include`; không tách các bước kỹ thuật nội bộ.
5. Tìm hành vi tùy chọn/có điều kiện để cân nhắc `extend`.
6. Xác định Primary và Supporting Actors.
7. Đặt UC ID theo quy ước của sơ đồ tổng quan.
8. Kiểm tra mọi Use Case trên sơ đồ có đặc tả hoặc được giải thích rõ.
9. Kiểm tra hướng mũi tên và ký hiệu.
10. Đối chiếu lại với diagram tổng quan và requirement.

#### Tiêu chí hoàn thành

Sơ đồ chi tiết đạt yêu cầu khi:

- phạm vi chỉ tập trung vào một nhóm nghiệp vụ;
- các Use Case cùng mức trừu tượng;
- relationship làm rõ sự phụ thuộc, không trang trí;
- Actor và quyền tham gia rõ;
- từng hình oval có thể truy vết tới đặc tả;
- không trình bày trình tự như Sequence Diagram;
- không trình bày thuật toán như Activity Diagram.

### 5.3. So sánh hai loại sơ đồ

| Tiêu chí | Diagram tổng quan | Diagram chi tiết |
|---|---|---|
| Mục tiêu | Hiểu toàn cảnh và phạm vi | Hiểu một phân hệ hoặc nhóm nghiệp vụ |
| Đối tượng đọc | Stakeholder, giảng viên, người mới | BA, developer, tester, reviewer kỹ thuật |
| Số Actor | Các Actor chính | Actor liên quan trực tiếp đến phạm vi |
| Use Case | Mục tiêu cấp cao | Các mục tiêu hoàn chỉnh thuộc phạm vi hẹp hơn |
| Relationship | Hạn chế | Có thể dùng nhiều hơn nhưng phải có ý nghĩa |
| Chi tiết flow | Không | Không; flow vẫn nằm trong Specification |
| UC ID | Bắt buộc để truy vết | Kế thừa hoặc phân cấp từ mô hình tổng quan |
| Boundary | Toàn hệ thống | Toàn hệ thống hoặc một phân hệ được ghi rõ |

### 5.4. Quy tắc nhất quán giữa các sơ đồ

- cùng UC ID phải có cùng tên và ý nghĩa;
- cùng Actor phải có cùng vai trò;
- sơ đồ chi tiết không được thêm quyền trái với sơ đồ tổng quan;
- không biến bước kỹ thuật thành Use Case trên sơ đồ chi tiết;
- mọi relationship phải phù hợp với flow trong đặc tả;
- nếu thay đổi diagram chi tiết, phải kiểm tra diagram tổng quan;
- nếu một Use Case bị tách, phải cập nhật đặc tả, requirement traceability và Test Case.

Sơ đồ chi tiết vẫn là Use Case Diagram, không phải flow chart. Vì vậy không dùng oval để biểu diễn thứ tự xử lý, không đánh số oval theo chuỗi bước `UC01.1`, `UC01.2`, `UC01.3` nếu chúng chỉ là flow của UC01, không đưa điều kiện như `ngân sách đủ`, `provider lỗi`, `callback hợp lệ` thành Use Case, và không vẽ relationship tự tạo như `<<precondition>>`, `<<trigger>>`, `<<success>>`.

Nếu cần thể hiện điều kiện rẽ nhánh hoặc trình tự, dùng Activity Diagram. Nếu cần thể hiện đối tượng và thông điệp theo thời gian, dùng Sequence Diagram.

### 5.5. Cây tài liệu đề xuất

```text
report_support_documentation/
├── use_case_analysis.md
└── use_cases/
    ├── UC01/
    │   ├── specification.md
    │   └── diagram.puml
    ├── UC02/
    │   ├── specification.md
    │   └── diagram.puml
    └── ...
```

Mỗi thư mục Use Case chỉ chứa:

- `specification.md`: đặc tả theo mẫu Summary, Flow và Additional Information;
- `diagram.puml`: mã PlantUML của sơ đồ Use Case tương ứng.

Không gộp đặc tả của nhiều Use Case vào một file review chung. Diagram tổng quan toàn hệ thống có thể được quản lý riêng vì nó không thuộc duy nhất một Use Case.

---

## 6. Cấu trúc đầy đủ của Use Case Specification

Nguồn chia đặc tả thành ba nhóm: `Summary`, `Flow` và `Additional Information`. Mẫu trong đồ án có thể giữ cấu trúc hiện tại, miễn là mỗi Use Case có đủ thông tin để đọc, kiểm thử và truy vết.

UC01 nên được viết đầy đủ hơn vì là Use Case trung tâm của đề tài. UC02--UC09 nên ngắn hơn, tập trung vào mục tiêu của Actor và các nhánh quan trọng nhất để báo cáo không bị phình quá mức.

### 6.1. Summary

| Thành phần | Ý nghĩa | Hướng dẫn viết |
|---|---|---|
| Use Case ID | Mã định danh duy nhất | Ví dụ `UC01`; không đổi giữa diagram, đặc tả và các biểu đồ liên quan |
| Use Case Name | Tên Use Case | Dùng `Động từ + Đối tượng` |
| Description | Mô tả ngắn | Actor muốn làm gì và để đạt mục tiêu gì; nên ngắn gọn |
| Actor(s) | Actor tham gia | Phân biệt Primary Actor và Supporting Actor |
| Priority | Mức ưu tiên | Ví dụ Must/Should/Could hoặc Cao/Trung bình/Thấp |
| Trigger | Sự kiện khởi động | Hành động hoặc sự kiện làm Use Case bắt đầu |
| Pre-Condition | Điều kiện tiên quyết | Trạng thái phải đúng trước khi chạy flow |
| Post-Condition | Trạng thái sau cùng | Nêu trạng thái bảo đảm khi thành công; có thể thêm trạng thái khi thất bại |

Mẫu Description:

> Là `[Actor]`, tôi muốn `[Use Case Name]` để `[mục tiêu/giá trị]`.

Không bắt buộc phải dùng nguyên văn User Story, nhưng mô tả cần giữ đủ ba thành phần: Actor, hành động và giá trị.

### 6.2. Flow

#### Basic Flow

Basic Flow là đường đi chính, đơn giản nhất và dẫn tới Use Case thành công.

Nguyên tắc:

- dùng câu chủ động;
- gọi rõ chủ thể của mỗi bước, ví dụ `AI Agent`, `Gateway`, `Payment Provider`, `Quản trị viên tổ chức`;
- viết theo thứ tự thời gian;
- mỗi bước chỉ mô tả một hành động hoặc phản hồi chính;
- luân phiên rõ Actor và hệ thống;
- tập trung vào hành vi quan sát được;
- không nhúng nhiều logic `nếu/không thì` vào Basic Flow;
- không đưa chi tiết cài đặt không cần thiết;
- kết thúc bằng mục tiêu đã đạt và post-condition tương ứng.

Với UC01, Basic Flow nên nằm khoảng 7--10 bước. Các nội dung như replay request, fallback, ngân sách không đủ, provider lỗi, client ngắt kết nối hoặc quyết toán lỗi nên đưa vào Alternative/Exception Flow thay vì làm Basic Flow quá dài.

#### Alternative Flow

Alternative Flow là một đường đi khác Basic Flow nhưng **Use Case vẫn thành công**.

Ví dụ:

- chọn phương thức xác thực khác;
- định tuyến sang mô hình AI khác nhưng vẫn trả được kết quả;
- chọn cách lọc hoặc xuất dữ liệu khác;
- hệ thống nhận kết quả thanh toán bằng đối soát thay vì callback.

Mỗi Alternative Flow phải:

- tham chiếu bước rẽ nhánh trong Basic Flow;
- nêu điều kiện rẽ nhánh;
- mô tả các bước thay thế;
- chỉ ra quay lại bước nào hoặc kết thúc thành công ở đâu.

Tên nhánh nên bắt đầu bằng số bước gốc, ví dụ `4a - Chọn provider thay thế`. Cách đặt tên này giúp đặc tả nối được với Basic Flow, Activity Diagram và Test Case.

#### Exception Flow

Exception Flow là luồng làm Use Case **không đạt mục tiêu** hoặc kết thúc thất bại.

Ví dụ:

- xác thực khóa thất bại;
- không đủ ngân sách và không có mô hình thay thế;
- chữ ký callback không hợp lệ;
- cập nhật giao dịch thất bại và bị rollback;
- tài khoản không có quyền thực hiện.

Mỗi Exception Flow phải:

- tham chiếu bước phát sinh lỗi;
- nêu nguyên nhân hoặc điều kiện;
- mô tả phản ứng của hệ thống;
- nêu trạng thái dữ liệu sau rollback/compensation;
- chỉ rõ Use Case dừng hay cho phép thử lại.

Đối với Use Case tài chính, Exception Flow không được chỉ viết “thông báo lỗi”. Cần nêu rõ trạng thái tài chính sau lỗi: không phát sinh chi phí, hoàn tiền, giải phóng tạm giữ, giữ trạng thái chờ đối soát hoặc ghi nhận lỗi để xử lý sau.

### 6.3. Additional Information

#### Business Rules

Business Rule là các quy định nghiệp vụ hoặc policy bắt buộc.

Ví dụ trong đồ án:

- không cho phép chi vượt ngân sách khả dụng;
- giao dịch cùng idempotency key không được ghi nhận hai lần;
- mã khóa thô chỉ hiển thị một lần;
- hạn mức của tác nhân không vượt hạn mức khả dụng của lập trình viên;
- số tiền nạp phải nằm trong ngưỡng tối thiểu và tối đa;
- chỉ callback có chữ ký hợp lệ mới được cập nhật số dư.

Nên gắn mã như `BR-UC01-01` để truy vết.

#### Non-Functional Requirements

Use Case chủ yếu thể hiện Functional Requirement. Các yêu cầu phi chức năng liên quan trực tiếp nên được dẫn chiếu trong đặc tả.

Ví dụ:

- độ trễ kiểm tra ngân sách;
- tính nguyên tử;
- tính nhất quán khi xử lý đồng thời;
- thời gian timeout;
- yêu cầu mã hóa hoặc hashing;
- audit log;
- khả năng phục hồi;
- bảo mật dữ liệu nhạy cảm.

Nên tham chiếu mã NFR đã định nghĩa thay vì lặp lại diễn giải dài.

---

## 7. Các khái niệm dễ nhầm

### 7.1. Trigger và Pre-Condition

**Trigger** là sự kiện làm Use Case bắt đầu.

**Pre-Condition** là trạng thái phải tồn tại trước khi Use Case có thể được thực hiện hợp lệ.

Ví dụ UC02 - Nạp tiền:

- Trigger: Quản trị viên chọn chức năng `Nạp tiền`.
- Pre-Condition:
  - quản trị viên đã đăng nhập;
  - tổ chức đang hoạt động;
  - ví tổ chức đã tồn tại.

Trigger có thể là:

- hành động người dùng;
- callback từ hệ thống ngoài;
- sự kiện thời gian;
- job định kỳ;
- sự kiện nội bộ.

Trigger có thể trùng với bước đầu tiên của Basic Flow, nhưng không bắt buộc.

### 7.2. Pre-Condition và bước kiểm tra trong flow

Pre-Condition là điều được giả định đã đúng trước khi Use Case bắt đầu. Nếu hệ thống phải kiểm tra điều kiện như một phần đáng kể của tương tác, kiểm tra đó có thể nằm trong flow.

Ví dụ:

- “Người dùng đã đăng nhập” có thể là Pre-Condition của Use Case quản trị.
- “API key hợp lệ” trong UC01 nên được kiểm tra trong flow nếu chính Gateway chịu trách nhiệm xác thực request.

Không nên vừa giả định một điều là Pre-Condition vừa mô tả hệ thống kiểm tra nó như một bước trọng tâm mà không giải thích rõ phạm vi.

### 7.3. Post-Condition và Kết quả

`Kết quả` thường mô tả giá trị người dùng nhận được.

`Post-Condition` mô tả trạng thái hệ thống có thể kiểm chứng sau Use Case.

Ví dụ UC01:

- Kết quả: AI Agent nhận được kết quả xử lý hoặc lỗi nghiệp vụ rõ ràng.
- Success Post-Condition:
  - giao dịch đã quyết toán;
  - ví phản ánh đúng chi phí thực tế;
  - trace bất biến đã được ghi.
- Failure Post-Condition:
  - không phát sinh khấu trừ trùng;
  - khoản tạm giữ được giải phóng hoặc đánh dấu chờ xử lý có kiểm soát;
  - lỗi được ghi log.

### 7.4. Alternative Flow và Exception Flow

Tiêu chí phân biệt đơn giản nhất:

| Loại flow | Có phải đường chính? | Mục tiêu Use Case có đạt không? |
|---|---:|---:|
| Basic Flow | Có | Có |
| Alternative Flow | Không | Có |
| Exception Flow | Không | Không |

Không phân loại chỉ dựa trên việc có “lỗi” kỹ thuật. Nếu người dùng có thể sửa đầu vào, tiếp tục flow và cuối cùng vẫn đạt mục tiêu, nhánh đó có thể là Alternative Flow. Nếu flow kết thúc mà mục tiêu không đạt, đó là Exception Flow.

### 7.5. Post-Condition của Use Case này và Trigger của Use Case khác

Trạng thái sau khi một Use Case thành công có thể kích hoạt Use Case khác.

Ví dụ:

- UC02 thành công làm số dư ví tăng;
- số dư mới cho phép UC03 phân bổ ngân sách;
- UC04 thiết lập chính sách thành công tạo điều kiện cho UC01 kiểm soát request.

Không vì có quan hệ trước-sau mà bắt buộc phải vẽ `include` hoặc `extend`. Có thể chỉ cần thể hiện bằng pre/post-condition và traceability.

Không nối hai Use Case bằng `<<precondition>>`. Đây không phải relationship chuẩn của Use Case Diagram.

---

## 8. Quy tắc viết flow rõ ràng

### 8.1. Đánh số

- Basic Flow: `1`, `2`, `3`, ...
- Nhánh tại bước 4: `4a`, `4a.1`, `4a.2`, ...
- Nhánh khác tại bước 4: `4b`, `4b.1`, ...
- Kết thúc nhánh bằng một trong các chỉ dẫn:
  - `Use Case tiếp tục tại bước N`;
  - `Use Case kết thúc thành công`;
  - `Use Case dừng, mục tiêu không đạt`;
  - `Chuyển sang UCxx`.

Không nên dùng cùng một mã nhánh cho hai điều kiện khác nhau.

### 8.2. Cấu trúc câu

Ưu tiên:

> `[Actor/Hệ thống] + [động từ] + [đối tượng] + [điều kiện cần thiết]`.

Ví dụ:

- `AI Agent gửi yêu cầu kèm API key và idempotency key.`
- `Hệ thống xác thực API key và xác định chính sách ngân sách tương ứng.`
- `Nhà cung cấp AI trả kết quả và usage metadata.`

Tránh:

- câu không có chủ thể;
- một bước chứa quá nhiều hành động độc lập;
- thuật ngữ cài đặt như tên class, method hoặc bảng cơ sở dữ liệu;
- mô tả giao diện quá chi tiết nếu không ảnh hưởng nghiệp vụ;
- gộp nhiều nhánh logic vào một bước.

### 8.3. Tính đối xứng giữa các flow

Alternative/Exception Flow phải bám đúng bước rẽ nhánh trong Basic Flow.

Ví dụ:

- Basic step 5: Hệ thống kiểm tra ngân sách.
- Alternative `5a`: Ngân sách không đủ cho mô hình yêu cầu nhưng đủ cho mô hình thay thế.
- Exception `5b`: Ngân sách không đủ và không có mô hình thay thế.

Cách tổ chức này giúp:

- nhìn ra điều kiện rẽ nhánh;
- chuyển sang Activity Diagram dễ hơn;
- thiết kế Test Case trực tiếp;
- kiểm tra mọi nhánh có trạng thái kết thúc.

### 8.4. Mức chi tiết

Flow phải đủ chi tiết để stakeholder, developer và tester cùng hiểu, nhưng không biến thành pseudocode.

Nên mô tả:

- Actor gửi gì;
- hệ thống kiểm tra quy tắc nghiệp vụ gì;
- hệ thống trả gì;
- trạng thái nghiệp vụ thay đổi ra sao;
- hệ thống ngoài tham gia tại đâu.

Không cần mô tả trong Use Case:

- tên class hoặc method;
- câu lệnh SQL;
- loại khóa cơ sở dữ liệu cụ thể, trừ khi đây là ràng buộc thiết kế bắt buộc;
- cấu trúc cache nội bộ;
- thuật toán hashing cụ thể nếu đã thuộc Security Design/NFR;
- chi tiết triển khai transaction ở cấp code.

Các chi tiết kỹ thuật quan trọng vẫn phải được lưu ở phần Architecture, Design hoặc NFR và được dẫn chiếu từ Use Case khi cần.

---

## 9. Mẫu Use Case Specification

```markdown
## UCxx - Tên Use Case

| Thuộc tính | Nội dung |
|---|---|
| ID | UCxx |
| Tên | Động từ + Đối tượng |
| Mô tả | Là [Actor], tôi muốn [hành động] để [giá trị] |
| Phạm vi | AI Agent Financial Gateway |
| Primary Actor | Actor khởi tạo và nhận giá trị chính |
| Supporting Actors | Hệ thống/người hỗ trợ |
| Priority | Must / Should / Could |
| Trigger | Sự kiện làm Use Case bắt đầu |
| Pre-Conditions | Các trạng thái phải có trước khi bắt đầu |
| Success Post-Conditions | Trạng thái được bảo đảm khi thành công |
| Failure Post-Conditions | Trạng thái được bảo đảm khi thất bại |

### Basic Flow

1. ...
2. ...
3. ...

### Alternative Flows

#### [Bước]a - Tên nhánh thành công

1. Điều kiện rẽ nhánh...
2. ...
3. Use Case tiếp tục tại bước ...

### Exception Flows

#### [Bước]b - Tên nhánh thất bại

1. Điều kiện lỗi...
2. Hệ thống...
3. Use Case dừng; mục tiêu không đạt.

### Business Rules

- BR-UCxx-01: ...

### Non-Functional Requirements

- NFRxx: ...

### Traceability

- Functional Requirements: FRxx...
- Sequence Diagram: ...
- Activity Diagram: ...
- Test Cases: ...
```

---

## 10. Quy trình phân tích Use Case

### Bước 1 - Xác định phạm vi

Viết một câu boundary:

> AI Agent Financial Gateway chịu trách nhiệm xác thực request, kiểm soát và quyết toán ngân sách, định tuyến dịch vụ AI, quản lý chính sách tài chính và cung cấp khả năng truy vết; hệ thống không triển khai mô hình AI hoặc xử lý thanh toán ngân hàng trực tiếp.

Liệt kê rõ:

- chức năng trong phạm vi;
- chức năng ngoài phạm vi;
- hệ thống ngoài;
- giả định nghiệp vụ.

### Bước 2 - Lập danh sách Actor

Với mỗi Actor, ghi:

- vai trò;
- mục tiêu;
- dữ liệu gửi vào;
- dữ liệu nhận ra;
- quyền;
- Use Case liên quan.

### Bước 3 - Tìm mục tiêu của Actor

Không bắt đầu từ bảng dữ liệu hoặc module code. Với mỗi Actor, hỏi:

- Actor muốn đạt kết quả nghiệp vụ gì?
- Kết quả đó có giá trị và quan sát được không?
- Hệ thống có trách nhiệm gì để tạo kết quả?

### Bước 4 - Chuẩn hóa tên và mức trừu tượng

Kiểm tra:

- mọi tên theo `Động từ + Đối tượng`;
- các Use Case trên cùng sơ đồ có mức chi tiết tương đương;
- các Use Case “Quản lý” được làm rõ phạm vi;
- không có tên trùng nghĩa.

### Bước 5 - Vẽ diagram tổng quát

- đặt Actor ngoài Boundary;
- đặt Use Case trong Boundary;
- gắn UC ID;
- nối Actor đúng vai trò;
- chỉ dùng relationship khi có ý nghĩa;
- giữ sơ đồ dễ đọc.

### Bước 6 - Viết Basic Flow trước

Viết happy path ngắn nhất dẫn tới mục tiêu thành công. Chưa đưa lỗi vào Basic Flow.

### Bước 7 - Tìm Alternative và Exception

Tại mỗi bước Basic Flow, hỏi:

- có lựa chọn hợp lệ khác không?
- dữ liệu có thể không hợp lệ không?
- Actor có thể hủy không?
- hệ thống ngoài có thể timeout hoặc trả lỗi không?
- có vấn đề phân quyền không?
- có vấn đề concurrency/idempotency không?
- nếu lỗi xảy ra thì dữ liệu được rollback hay compensation?
- mục tiêu cuối cùng còn đạt không?

### Bước 8 - Gắn Business Rule và NFR

Mỗi quyết định trong flow phải truy được về:

- Business Rule;
- Functional Requirement;
- Non-Functional Requirement;
- giả định đã được phê duyệt.

### Bước 9 - Kiểm tra traceability

Mỗi Use Case phải liên kết được với:

- yêu cầu chức năng;
- Actor;
- Sequence Diagram nếu có;
- Activity Diagram nếu có;
- lớp điều khiển;
- Test Case.

### Bước 10 - Review với ba góc nhìn

- **Stakeholder:** có hiểu mục tiêu và kết quả không?
- **Developer:** có đủ rõ để triển khai không?
- **Tester:** có xác định được happy path, alternative và failure cases không?

---

## 11. Checklist review Use Case Diagram

- [ ] Nếu dùng Boundary theo quy ước của Đồ án 1, Boundary có tên và đúng phạm vi.
- [ ] Mọi Actor nằm ngoài Boundary.
- [ ] Mọi Use Case nằm trong Boundary.
- [ ] Actor là vai trò hoặc hệ thống ngoài, không phải module nội bộ.
- [ ] Tên Actor là danh từ/cụm danh từ.
- [ ] Tên Use Case theo `Động từ + Đối tượng`.
- [ ] Mỗi Use Case tạo ra mục tiêu có giá trị cho Actor.
- [ ] Các Use Case có cùng mức trừu tượng.
- [ ] Diagram không biến thành sơ đồ module hoặc CRUD.
- [ ] Mỗi Use Case có UC ID.
- [ ] ID và tên khớp với đặc tả.
- [ ] `include` chỉ dùng cho hành vi bắt buộc và hợp lý để tái sử dụng/tách riêng.
- [ ] Không đặt mã UC cho các bước xử lý kỹ thuật trong Basic Flow.
- [ ] Không biến Use Case Diagram chi tiết thành Activity Diagram hoặc Sequence Diagram.
- [ ] `extend` chỉ dùng cho hành vi tùy chọn/có điều kiện.
- [ ] Mỗi `extend` tham chiếu ít nhất một extension point thuộc Use Case nền tảng.
- [ ] Điều kiện kích hoạt của `extend`, nếu có, được mô tả rõ.
- [ ] Mũi tên `extend` đi từ Use Case mở rộng về Use Case nền tảng.
- [ ] Mũi tên `include` đi từ Use Case bao gồm về Use Case được bao gồm.
- [ ] Không sử dụng relationship tự tạo `<<precondition>>`; điều kiện tiên quyết nằm trong đặc tả.
- [ ] Generalization đúng hướng và đúng ký hiệu UML.
- [ ] Không có đường nối chồng chéo khó đọc.
- [ ] Số Use Case trên một hình vẫn đọc được; tách theo phân hệ nếu cần.
- [ ] Mọi hệ thống ngoài trao đổi dữ liệu đều được xem xét như Actor.
- [ ] Không có Use Case trên diagram mà thiếu đặc tả.
- [ ] Không có đặc tả mà thiếu Use Case trên diagram, trừ khi có giải thích.

---

## 12. Checklist review Use Case Specification

- [ ] Có ID duy nhất.
- [ ] Tên khớp diagram.
- [ ] Description nêu Actor, hành động và giá trị.
- [ ] Primary Actor thực sự khởi tạo hoặc nhận giá trị chính.
- [ ] Supporting Actors đầy đủ.
- [ ] Có Priority.
- [ ] Có Trigger rõ ràng.
- [ ] Pre-Condition là trạng thái có trước, không phải flow trá hình.
- [ ] Basic Flow là happy path ngắn nhất.
- [ ] Mỗi bước có chủ thể rõ.
- [ ] Các bước thể hiện tương tác, không phải pseudocode.
- [ ] Alternative Flow vẫn dẫn tới thành công.
- [ ] Exception Flow dẫn tới thất bại hoặc mục tiêu không đạt.
- [ ] Mọi nhánh tham chiếu đúng bước gốc.
- [ ] Mọi nhánh chỉ rõ quay lại, kết thúc hoặc chuyển Use Case.
- [ ] Có Success Post-Conditions.
- [ ] Có Failure Post-Conditions đối với giao dịch quan trọng.
- [ ] Business Rules có mã và truy vết.
- [ ] NFR liên quan được dẫn chiếu.
- [ ] Không mâu thuẫn với yêu cầu chức năng.
- [ ] Không mâu thuẫn với Sequence/Activity Diagram.
- [ ] Có thể suy ra Test Case từ từng flow.

---

## 13. Definition of Done cho mô hình Use Case

Phần Use Case có thể coi là hoàn thành khi:

1. Actor và phạm vi hệ thống được định nghĩa rõ.
2. Diagram tổng quát dễ đọc và đúng ký hiệu.
3. Tất cả Use Case có ID và tên thống nhất.
4. Mọi Use Case trên diagram có đặc tả tương ứng.
5. Mỗi đặc tả có Trigger, Pre-Condition, Basic, Alternative, Exception và Post-Condition.
6. Các flow tài chính có Failure Post-Condition bảo đảm không mất tiền, không ghi trùng và có khả năng đối soát.
7. Business Rule và NFR được truy vết bằng mã.
8. Các Use Case được chọn để mô hình hóa tuần tự phải khớp với Sequence Diagram tương ứng.
9. Các Use Case được mô hình hóa luồng nghiệp vụ phải khớp với Activity Diagram tương ứng.
10. Mỗi nhánh chính/phụ/ngoại lệ có thể chuyển thành ít nhất một Test Case.
11. Thuật ngữ Actor, ví, hạn mức, ngân sách, giao dịch, tác nhân và khóa truy cập được dùng nhất quán.
12. Giảng viên hoặc stakeholder có thể đọc diagram và đặc tả mà không cần biết cấu trúc code.

---

## 14. Nguồn tham khảo

- Object Management Group (OMG), *Unified Modeling Language, Version 2.5.1*, December 2017, Clause 18 - UseCases, trang in 638-649:  
  [formal-17-12-05.pdf](./formal-17-12-05.pdf)
- Thinhnotes, “Use Case Diagram và 5 sai lầm thường gặp”:  
  https://thinhnotes.com/chuyen-nghe-ba/use-case-diagram-va-5-sai-lam-thuong-gap/
- Thinhnotes, “Viết đặc tả Use Case sao đơn giản nhưng hiệu quả?”:  
  https://thinhnotes.com/chuyen-nghe-ba/viet-dac-ta-use-case-sao-don-gian-nhung-hieu-qua/
- Trịnh Thành Trung, *Bài 13 - Tổng quan về UML*, tài liệu bài giảng OOAD, 2016:  
  [OOP_Bai13(vi).pdf](./OOP_Bai13(vi).pdf)

Ngày truy cập website: 10/06/2026.

---

## 15. Kho hình minh họa từ Thinhnotes

Các hình minh họa Use Case Diagram trong bài viết đã được lưu tại:

[images/use_case_examples_thinhnotes/](./images/use_case_examples_thinhnotes/)

Danh mục, chú thích và mục đích sử dụng từng hình:

[README - Use Case examples from Thinhnotes](./images/use_case_examples_thinhnotes/README.md)

Nguyên tắc sử dụng:

- dùng ảnh để nhận diện ký pháp, so sánh cách trình bày và phân tích lỗi;
- luôn ghi nguồn `Thinhnotes.com` khi đưa ảnh vào báo cáo hoặc slide;
- không coi mọi quan hệ trong ảnh là mẫu UML chuẩn;
- nếu nội dung ảnh mâu thuẫn với Clause 18 của OMG UML 2.5.1, áp dụng quy tắc OMG;
- đặc biệt không sao chép quan hệ `include` chỉ để biểu diễn thứ tự, điều kiện tiên quyết hoặc các bước kỹ thuật.
