<a id="readme-top"></a>

<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://hcmut.edu.vn/img/nhanDienThuongHieu/01_logobachkhoasang.png" alt="Logo" width="30%" height="30%">
  </a>
  <h2 align="center">ĐỒ ÁN TỔNG HỢP HK241 - GRADE PORTAL SERVICE</h2>
  <p align="center">
    Dự án phát triển Backend dành cho ứng dụng website tra cứu và thao tác với điểm dành cho sinh viên và giảng viên đại học.
    <br />
    <a href="https://github.com/dath-241/grade-portal-be-go-2">View Demo</a>
    ·
    <a href="https://github.com/dath-241/grade-portal-be-go-2/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/dath-241/grade-portal-be-go-2/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>

## Thông tin nhóm
- Đề tài : `Grade Portal`
- Giảng viên hướng dẫn: `Lê Đình Thuận`
- Nhiệm vụ : `grade-portal-be-go-2`
- Tên nhóm :  ≽^• GO GO GO • ྀི≼
- Phân công công việc chi tiết tại báo cáo hằng tuần: [tại đây](docs/reports/Weekly)
- Danh sách thành viên:

| STT | Họ và Tên           | MSSV    | Vai Trò       | Github                                                     |
| --- | ------------------- | ------- | ------------- |------------                                                |
| 1   | Chu Minh Tâm        | 2213009 | Product Owner |[inHansZone](https://github.com/inHansZone)                           |
| 2   | Nguyễn Trọng Kha    | 2211417 | Developer     |[KhaLeopard](https://github.com/Nguyentrongkha2k4)          |
| 3   | Nguyễn Tiến Phát    | 2212527 | Developer     |[Nguyen Tien Phat](https://github.com/nguyentienphat2904)   |
| 4   | Nguyễn Ngọc Diệu Hà | 2210846 | Developer     |[HANND04](https://github.com/HANND04)                       |
| 5   | Lê Hoàng Ngọc Hân   | 2210935 | Developer     |[inHansZone](https://github.com/inHansZone)                 |
| 6   | Nguyễn Phúc Hưng    | 2211368 | Developer     |[Richard](https://github.com/Richard1112004)                |
| 7   | Huỳnh Thanh Tâm     | 2213012 | Developer     |[ImNotTam](https://github.com/ImNotTam)                     |

## GIỚI THIỆU DỰ ÁN
### Bối cảnh 
Trước sự gia tăng không ngừng về số lượng sinh viên tại Đại học Bách Khoa - Đại học Quốc gia TP.HCM, việc quản lý thông tin học tập và điểm số đang trở thành một thách thức lớn đối với nhà trường. Song song với đó, chương trình đào tạo không ngừng được cập nhật để theo kịp sự phát triển nhanh chóng của công nghệ và tri thức toàn cầu, đặt ra nhu cầu cấp thiết về một hệ thống hỗ trợ quản lý và tra cứu thông tin học tập hiện đại. Đặc biệt, một hệ thống trực tuyến cho phép sinh viên tra cứu điểm số dễ dàng không chỉ mang lại sự minh bạch và tiện lợi cho người học mà còn hỗ trợ giảng viên trong việc quản lý lớp học và điểm số một cách hiệu quả.

Dự án Grade Portal ra đời nhằm hiện thực hóa mục tiêu xây dựng một hệ thống quản lý điểm số trực tuyến toàn diện, phục vụ sinh viên và giảng viên tại Đại học Bách Khoa. Đây sẽ là một giải pháp công nghệ tiên tiến, đáp ứng nhu cầu hiện đại hóa trong giáo dục, tạo điều kiện thuận lợi cho việc theo dõi, quản lý và cập nhật thông tin học tập một cách minh bạch, chính xác và kịp thời. Với sự hỗ trợ của hệ thống này, cả sinh viên lẫn giảng viên sẽ có trong tay một công cụ hữu ích, góp phần nâng cao trải nghiệm học tập và giảng dạy trong bối cảnh chuyển đổi số mạnh mẽ.
### Mục tiêu
Dự án Grade Portal hướng tới việc phát triển một hệ thống quản lý điểm số trực tuyến hiện đại, giúp sinh viên và giảng viên tại Đại học Bách Khoa tiếp cận và quản lý thông tin học tập dễ dàng, chính xác và hiệu quả hơn. Hệ thống này được thiết kế để đáp ứng nhu cầu ngày càng tăng về sự minh bạch, tiện lợi và tối ưu hóa trong quản lý giáo dục.

- **Đối với sinh viên:**
Grade Portal cung cấp một nền tảng linh hoạt giúp sinh viên dễ dàng tra cứu điểm số, theo dõi tiến trình học tập và cập nhật kết quả nhanh chóng, minh bạch. Thay vì sử dụng các phương pháp truyền thống như liên hệ qua email, gặp trực tiếp giảng viên hay kiểm tra bảng điểm giấy, sinh viên giờ đây có thể truy cập thông tin mọi lúc, mọi nơi. Điều này không chỉ giúp họ nắm bắt tiến độ học tập mà còn hỗ trợ việc điều chỉnh chiến lược học tập một cách phù hợp và hiệu quả.

- **Đối với giảng viên:**
Hệ thống giúp giảm tải công việc quản lý điểm số thông qua các tính năng tự động hóa như đính kèm và xử lý bảng điểm từ các file CSV hoặc Excel. Nhờ đó, giảng viên có thể quản lý dữ liệu học tập một cách nhanh chóng, chính xác và dễ dàng hơn. Bên cạnh đó, tính năng phân quyền linh hoạt giữa giảng viên và quản trị viên sẽ giúp việc quản lý môn học trở nên hiệu quả, đáp ứng tốt nhu cầu đa dạng của nhà trường.

- **Đối Với nhà trường:**
Grade Portal không chỉ hỗ trợ việc quản lý và tra cứu điểm số mà còn đóng góp vào quá trình chuyển đổi số trong giáo dục. Hệ thống này sẽ là một bước tiến quan trọng trong việc cải thiện chất lượng quản lý học tập, thúc đẩy sự hiện đại hóa và nâng cao hiệu quả giảng dạy. Nếu triển khai thành công, Grade Portal sẽ trở thành một phần không thể thiếu trong hệ sinh thái giáo dục hiện đại của nhà trường, tạo tiền đề cho những thành tựu lớn hơn trong tương lai.

Grade Portal không chỉ là một giải pháp công nghệ mà còn là một nỗ lực hướng tới việc nâng cao chất lượng giáo dục tại Đại học Bách Khoa, góp phần vào sự phát triển toàn diện của nhà trường trong kỷ nguyên số hóa.
### TÍNH NĂNG
- `Admin`
  - Quản lý hệ thống, môn học, lớp học và phân quyền cho người dùng.
- `Teacher`
  - Chỉnh sửa lớp học, bảng điểm bằng cách đính kèm file CSV.
- `Student`
  - Tra cứu điểm số các môn học trong từng học kỳ.
- `Hall of Fame`
  - Hệ thống tự động cập nhật và hiển thị danh sách top sinh viên có điểm cao nhất theo kỳ, năm học.
### Funtion, non-function cùng các diagram
Xem chi tiết trong báo cáo tổng hợp: [tại đây](docs/reports/FinalReport/finalReport.md)
### API
- Thông tin chung về API: [API](./docs/api/api_specification.md)
- Thông tin chi tiết về API: [API Specification](./docs/api/api_specification_detail.md).
- Thông tin API định dạng json: [Postman Agent](./docs/api/Grade-Portal-Be-Go-2.json).
### Mô hình hệ thống
![MVC model](./docs/reports/Weekly/img/MVC.png)

- `Route` :
  - Định nghĩa các `URL` và liên kết với các `Controller` tương ứng.
  - `Route` sẽ gọi đến `Controller` khi có một yêu cầu HTTP được gửi đến.
- `Controller` :
  - `Controller` xử lý logic của yêu cầu HTTP (GET, POST, PUT, DELETE). Nó nhận dữ liệu từ `Middleware`, tương tác với các `model` và trả lại phản hồi cho `client`.
  - `Controller` chỉ thực hiện các tác vụ logic như xử lý dữ liệu, gửi phản hồi, hoặc tương tác với cơ sở dữ liệu.
- `Middleware` :
  - `Middleware` là các lớp trung gian xử lý yêu cầu trước khi chúng đến `Controller` hoặc sau khi `Controller` xử lý xong. Nó có thể thực hiện các tác vụ như xác thực, logging, hay sửa đổi yêu cầu.
  - `Middleware` hoạt động như bộ lọc, giúp bảo vệ và kiểm tra các yêu cầu trước khi chúng được định tuyến đến `Controller`.
- `model` : xử lý logic liên quan đến dữ liệu (`CRUD`- Create - Read - Update - Delete). `Controller` gọi đến `model` để thao tác với `Database`.

- `Response` : sau khi `Controller` nhận data từ `model`, nó sẽ gửi phản hồi (responese) lại cho `client` thông qua các `route` ban đầu.
  Quá trình này
## Công nghệ sử dụng
- <a href="https://go.dev/doc/"><img src="https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Golang" /></a> </br>
- <a href="https://pkg.go.dev/github.com/gin-gonic/gin"><img src="https://img.shields.io/badge/Gin_Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Gin" /> </a></br>
- <a href="https://www.mongodb.com/docs/"><img src="https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white" alt="MongoDB" /> </a></br>
- <a href="https://docs.docker.com/"><img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" /> </a></br>
- <a href="https://git-scm.com/docs/gitworkflows"><img src="https://img.shields.io/badge/GitFlow-F05032?style=for-the-badge&logo=git&logoColor=white" alt="GitFlow" /></a></br>
- <a href="https://docs.github.com/en"><img src="https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white" alt="GitHub"/></a></br>
## Hướng dẫn sử dụng
Xem hướng dẫn sử dụng: [tại đây](./docs/Guild/SETUP.md).  
## Báo cáo
- Báo cáo tiến độ làm việc theo tuần: [tại đây](./docs/reports/Weekly).
- Báo cáo tổng hợp: [tại đây](./docs/reports/FinalReport/finalReport.md).

## Liên hệ
Nếu bạn có bất kỳ thắc mắc, góp ý, hoặc cần báo lỗi liên quan đến ứng dụng, vui lòng liên hệ qua email: trongkha08022k4@gmail.com, hoặc ghé thăm GitHub cá nhân của từng thành viên trong nhóm. Chúng tôi rất mong nhận được phản hồi từ bạn để cải thiện ứng dụng tốt hơn!
  </br>
  </br>
  </br>

<p align="right"><a href="#readme-top">Về đầu trang</a></p>

> _Thông tin sẽ được cập nhật thường xuyên khi có chỉnh sửa_
