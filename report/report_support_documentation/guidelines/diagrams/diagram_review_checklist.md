# Checklist review diagram

Sử dụng checklist này cho từng diagram trước khi đưa vào báo cáo.

## Thông tin diagram

- **Tên diagram:**
- **Loại diagram:**
- **Chương/mục sử dụng:**
- **File source:**
- **File export:**
- **Use case liên quan:**
- **FR/NFR liên quan:**
- **Mức mô hình:** Requirement / Analysis / Design / As-built / Test evidence
- **Loại quy tắc:** UML / Phương pháp OOAD / ERD / Minh chứng không phải UML
- **Góc nhìn:** Overview / Focused

## 1. Kiểm tra nội dung

- [ ] Diagram đúng với [`project_source_of_truth.md`](../../governance/project_source_of_truth.md).
- [ ] Tên actor đúng danh sách actor chuẩn.
- [ ] Tên use case đúng danh sách use case chuẩn.
- [ ] Không có thành phần ngoài phạm vi hệ thống.
- [ ] Không mô tả hệ thống như nơi tạo/huấn luyện/điều phối AI Agent.
- [ ] Không để AI Agent dùng trực tiếp provider API key.
- [ ] Gateway thể hiện đúng vai trò kiểm soát identity/policy/budget/routing/trace nếu liên quan.

## 2. Kiểm tra OOAD

- [ ] Diagram đúng mục đích của loại biểu đồ.
- [ ] Diagram đúng mức trừu tượng của chương đang đặt.
- [ ] Nếu ở Chương 2: không lẫn class/service/repository kỹ thuật.
- [ ] Nếu ở Chương 3: phân tích lớp bám Boundary-Control-Entity.
- [ ] Nếu ở Chương 4: thiết kế có thể trace từ analysis model.
- [ ] Nếu ở Chương 4: Component/Deployment được ghi là as-designed.
- [ ] Nếu ở Chương 5: artifact phản ánh as-built hoặc kết quả kiểm thử thực tế.
- [ ] Quy tắc phương pháp/dự án không bị trình bày như cú pháp bắt buộc của UML.
- [ ] ERD hoặc biểu đồ kết quả kiểm thử không bị gọi nhầm là UML diagram.

## 3. Kiểm tra traceability

- [ ] Diagram liên kết được với ít nhất một use case hoặc requirement.
- [ ] Nếu diagram là use case/specification: FR/NFR mapping đúng source of truth.
- [ ] Nếu diagram là class/sequence/state: trace được về use case tương ứng.
- [ ] Không có use case/class/state xuất hiện lẻ loi mà không được giải thích.
- [ ] Phần tử lặp giữa overview và focused view giữ cùng tên, stereotype và ngữ nghĩa.

## 4. Kiểm tra hình thức

- [ ] Bố cục rõ, ít đường cắt nhau.
- [ ] Không quá nhiều node trong một diagram.
- [ ] Tên node ngắn, rõ, nhất quán.
- [ ] Font/chữ đủ lớn khi đưa vào PDF.
- [ ] Màu sắc/stereotype nhất quán nếu có dùng.
- [ ] Export không bị mờ, méo, vỡ hoặc mất chữ.

## 5. Kiểm tra LaTeX/report

- [ ] Caption đúng tên diagram.
- [ ] Hình từ nguồn bên ngoài có trích dẫn; hình tự thiết kế không cần dòng nguồn.
- [ ] Label LaTeX không trùng.
- [ ] Hình được nhắc đến trong nội dung.
- [ ] Đường dẫn ảnh đúng.
- [ ] Build PDF không lỗi.
- [ ] Hình không tràn trang.
- [ ] Diagram có mục trong `guidelines/diagrams/README.md` hoặc được ghi nhận là artifact mới cần bổ sung.

## 6. Kiểm tra riêng cho Chương 5

- [ ] As-built Component/Deployment Diagram chỉ chứa thành phần đã triển khai.
- [ ] Test case có expected result, actual result và trạng thái.
- [ ] Test trace được tới FR/NFR hoặc use case.
- [ ] Biểu đồ hiệu năng ghi workload, đơn vị, môi trường, cỡ mẫu và phương pháp đo.
- [ ] Nếu đo UC01, Gateway overhead và AI Provider latency được phân biệt khi có thể.

## 7. Kết luận review

- **Điểm đánh giá:** `/5`
- **Kết luận:** Đạt / Cần chỉnh / Vẽ lại
- **Vấn đề cần sửa:**
- **Người review:**
- **Ngày review:**
