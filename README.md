# Go CRUD API - Gin, GORM, Swaggo, Docker, Redis

Đây là một ứng dụng RESTful API mẫu được xây dựng bằng [Golang](https://golang.org/) sử dụng:

- ✅ [Gin](https://github.com/gin-gonic/gin): Web Framework mạnh mẽ và dễ sử dụng.
- ✅ [GORM](https://gorm.io/): ORM phổ biến để thao tác cơ sở dữ liệu.
- ✅ [Swaggo](https://github.com/swaggo/swag): Tạo tài liệu API tự động (Swagger).
- ✅ [Docker](https://www.docker.com/): Đóng gói và triển khai nhanh chóng.
- ✅ [Redis](https://redis.io/): Cache dữ liệu để tăng hiệu suất.

## 📁 Cấu trúc thư mục

```
go-crud-api/
│
├── config/         # Cấu hình DB, Redis
├── controllers/    # Xử lý request logic (Users, Products, Orders, ...)
├── docs/           # Swagger docs tự động sinh
├── middlewares/    # Middleware như xác thực JWT
├── models/         # Định nghĩa models GORM
├── routes/         # Định tuyến API
├── utils/          # Hàm tiện ích (JWT, hash, ...)
├── main.go         # Điểm khởi động chính
├── Dockerfile      # Dockerfile cho API
├── docker-compose.yml
└── go.mod
```

## 🚀 Hướng dẫn chạy project

### 1. Clone về máy

```bash
git clone https://github.com/your-username/go-crud-api.git
cd go-crud-api
```

### 2. Cấu hình môi trường `.env`

Tạo file `.env` và cấu hình như sau:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=go_crud

REDIS_ADDR=localhost:6379
JWT_SECRET=your_jwt_secret
```

> Nếu chạy bằng Docker, không cần tạo `.env` mà đã được cấu hình trong `docker-compose.yml`.

### 3. Build & Run với Docker

```bash
docker-compose up --build
```

Sau khi thành công:

- API chạy tại: `http://localhost:8080`
- Swagger: `http://localhost:8080/swagger/index.html`

### 4. Tạo tài liệu Swagger (nếu phát triển local)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

## 🔒 Xác thực & Phân quyền

Hệ thống sử dụng JWT để xác thực người dùng. Các vai trò chính:

- `admin`: truy cập và quản lý toàn bộ hệ thống
- `user`: chỉ được thao tác các chức năng cơ bản

```bash
Authorization: Bearer <your_token>
```

## ✅ Tính năng chính

- [x] CRUD Người dùng
- [x] CRUD Sản phẩm
- [x] Giỏ hàng và Đơn hàng
- [x] Đánh giá sản phẩm
- [x] Gửi email thông báo
- [x] Caching với Redis
- [x] Swagger Docs
- [x] Đóng gói với Docker
- [x] Phân quyền JWT

## 🧪 Test API

Bạn có thể test bằng:

- [Postman](https://www.postman.com/)
- Swagger UI (`/swagger/index.html`)

## 📬 Góp ý & Liên hệ

Nếu bạn có thắc mắc, hãy tạo [Issue](https://github.com/your-username/go-crud-api/issues) hoặc gửi email: `your.email@example.com`.

---

**Happy Coding! 🚀**
