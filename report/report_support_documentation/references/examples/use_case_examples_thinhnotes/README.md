# Hình minh họa Use Case từ Thinhnotes

Nguồn thu thập:

- Thinhnotes, “Use Case Diagram và 5 sai lầm thường gặp”:
  https://thinhnotes.com/chuyen-nghe-ba/use-case-diagram-va-5-sai-lam-thuong-gap/
- Ngày truy cập: 10/06/2026.

Các ảnh được lưu để phục vụ nghiên cứu, đối chiếu và phân tích trong Đồ án 1. Đây không phải bộ mẫu chuẩn để sao chép nguyên trạng. Ngữ nghĩa UML phải được kiểm tra theo OMG UML 2.5.1, Clause 18.

## Danh mục

| File | Nội dung | Cách sử dụng |
|---|---|---|
| [01-simple-blog-use-case.png](./01-simple-blog-use-case.png) | Ví dụ sơ đồ Use Case đơn giản của blog | Nhận diện Actor, Boundary, Use Case và Communication Link |
| [02-actor-notations.png](./02-actor-notations.png) | Các cách biểu diễn Actor | Nhấn mạnh Actor là vai trò, không bắt buộc là con người |
| [03-include-comment-example.png](./03-include-comment-example.png) | Ví dụ nhận xét bài viết với `include` | Phân tích hướng mũi tên và tiêu chí tái sử dụng |
| [04-include-atm-authentication.png](./04-include-atm-authentication.png) | Rút tiền include xác thực tài khoản | Ví dụ trực quan về hướng mũi tên `include` |
| [05-include-book-ride.png](./05-include-book-ride.png) | Đặt xe include các lựa chọn | Phân tích liệu các lựa chọn là Use Case dùng chung hay chỉ là bước trong flow |
| [06-include-rate-trip.png](./06-include-rate-trip.png) | Đánh giá chuyến đi include đặt xe | Ví dụ cần phê bình: quan hệ trước-sau hoặc precondition không mặc nhiên là `include` |
| [07-include-track-delivery.png](./07-include-track-delivery.png) | Theo dõi tiến độ giao hàng | Đối chiếu việc tách hành vi dùng chung |
| [08-include-reuse-example.png](./08-include-reuse-example.png) | Mô hình tái sử dụng X, Y, Z | Minh họa động cơ tái sử dụng của `include` |
| [09-include-selection-guidance.png](./09-include-selection-guidance.png) | Hướng dẫn chọn hành vi để include | Dùng cùng quy tắc giá trị độc lập và tái sử dụng |
| [10-extend-tip-driver.png](./10-extend-tip-driver.png) | Gửi tip extend đánh giá chuyến đi | Nhận diện chiều mũi tên từ Use Case mở rộng về Use Case nền tảng |
| [11-extend-with-extension-point.png](./11-extend-with-extension-point.png) | Extend có extension point và condition | Ví dụ gần nhất với ký pháp chi tiết của OMG |
| [12-extend-ecommerce-example.png](./12-extend-ecommerce-example.png) | Ví dụ `extend` trong thương mại điện tử | Phân tích hành vi tùy chọn và Use Case nền tảng |
| [13-extend-without-comments.png](./13-extend-without-comments.png) | Sơ đồ nhiều `extend` nhưng lược chú thích | Ví dụ về đánh đổi giữa đầy đủ ngữ nghĩa và khả năng đọc |
| [14-generalization-example.png](./14-generalization-example.png) | Generalization của Actor và Use Case | Kiểm tra đầu tam giác rỗng hướng về phần tử tổng quát |
| [15-mistake-naming.png](./15-mistake-naming.png) | Sai lầm đặt tên | So sánh tên kỹ thuật/cụm danh từ với `Động từ + Đối tượng` |
| [16-mistake-functional-decomposition.png](./16-mistake-functional-decomposition.png) | Nhầm Use Case với phân rã chức năng | Nhận diện các oval không thể hiện mục tiêu của Actor |
| [17-mistake-mixed-levels.png](./17-mistake-mixed-levels.png) | Trộn các mức trừu tượng | Kiểm tra tính đồng cấp của Use Case trên cùng sơ đồ |
| [18-mistake-cluttered-diagram.jpg](./18-mistake-cluttered-diagram.jpg) | Sơ đồ quá nhiều phần tử | Phân tích độ rối, tên sai và thiếu Boundary |
| [19-mistake-large-crud-diagram.jpg](./19-mistake-large-crud-diagram.jpg) | Sơ đồ lớn, tập trung CRUD | Nhận diện việc mô hình hóa quá chi tiết |
| [20-crud-consolidation.png](./20-crud-consolidation.png) | So sánh CRUD riêng lẻ và gom nhóm | Chỉ dùng như heuristic trình bày; vẫn phải bảo đảm tên Use Case thể hiện mục tiêu |

## Lưu ý chuẩn hóa

Một số nội dung giải thích trong bài viết mang tính thực hành và giản lược. Khi đánh giá ảnh:

1. `include` không biểu diễn “Use Case B phải xảy ra trước Use Case A” theo nghĩa thời gian.
2. Precondition không phải relationship trên Use Case Diagram.
3. `extend` đi từ Use Case mở rộng về Use Case nền tảng.
4. Extension point thuộc Use Case nền tảng; condition không đồng nghĩa với extension point.
5. Các bước nhập liệu, kiểm tra hoặc xử lý nội bộ không tự động trở thành Use Case.
6. Alternative/Exception Flow không tự động trở thành Use Case mở rộng.
