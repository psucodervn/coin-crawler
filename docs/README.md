## Tổng quan

Trong quá trình làm các project của công ty, cụ thể là khi build các service của backend, do chưa có kinh nghiệm nhiều nên mình đã gặp phải một số khó khăn khi lựa chọn công nghệ, thiết kế hệ thống cũng như lúc đi vào hiện thực cụ thể. Một trong số đó là việc xác định phương thức giao tiếp giữa các service với nhau.

Tới thời điểm hiện tại các service của hệ thống backend giao tiếp chủ yếu qua hai phương thức:

- Gọi api trực tiếp (REST hoặc gRPC) đối với các tác vụ đơn giản, có thể trả về kết quả ngay. Cách này mình thấy có hạn chế là mỗi service sẽ có một địa chỉ khác nhau, do đó khi số lượng service tăng lên thì việc quản lý địa chỉ của các service sẽ khá là mệt mỏi.
- Giao tiếp thông qua hệ thống pub/sub hay message queue đối với các tác vụ phức tạp, cần nhiều thời gian thực hiện hơn. Hiện đang dùng [Google Cloud Pub/Sub](https://cloud.google.com/pubsub/) và [RabbitMQ](https://www.rabbitmq.com/).

## NATS

[NATS](https://nats.io/) theo như trang chủ là một [cloud-native](https://stackify.com/cloud-native/) messaging system mã nguồn mở, được sử dụng nhiều nhờ hiệu năng và tính ổn định của nó. Hiện nó là một thành viên trong cộng đồng [CNCF](https://www.cncf.io/), nơi tập hợp và hỗ trợ các dự án cloud-native mã nguồn mở.

## Demo

