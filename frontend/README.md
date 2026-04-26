# 📢 Announcement Board

แอปพลิเคชัน กระดานประกาศ แบบ Full Stack สร้างด้วย Go, React และ PostgreSQL

---

## เทคโนโลยีที่ใช้

- **Backend:** Go 
- **Frontend:** React + TypeScript
- **Database:** PostgreSQL
- **Orchestration:** Docker Compose

---

## วิธีรันโปรเจกต์

### สิ่งที่ต้องติดตั้งก่อน

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) ติดตั้งและเปิดทิ้งไว้

### รันแอปพลิเคชัน

```bash
git clone <repo-url>
cd announcement-board
docker compose up --build
```

จากนั้นเปิดเบราว์เซอร์ที่:

```
http://localhost:3000
```

> การ build ครั้งแรกอาจใช้เวลาสักครู่ ครั้งถัดไปสามารถใช้ `docker compose up` โดยไม่ต้องใส่ `--build`

---

## API Endpoints

Base URL: `http://localhost:8080`

| Method | Endpoint | คำอธิบาย |
|--------|----------|----------|
| GET | `/announcements` | ดึงประกาศทั้งหมด (pinned ขึ้นก่อน) |
| POST | `/announcements` | สร้างประกาศใหม่ |
| PUT | `/announcements/:id` | แก้ไข / toggle pin ประกาศ |
| DELETE | `/announcements/:id` | ลบประกาศ |

### Request Body (POST / PUT)

```json
{
  "title": "หัวข้อประกาศ",
  "body": "เนื้อหาประกาศ",
  "author": "ชื่อผู้เขียน",
  "pinned": false
}
```

### Response ตัวอย่าง

```json
{
  "id": 1,
  "title": "หัวข้อประกาศ",
  "body": "เนื้อหาประกาศ",
  "author": "ชื่อผู้เขียน",
  "pinned": false,
  "created_at": "2026-04-26T12:00:00Z"
}
```

### Validation

ทุก field (`title`, `body`, `author`) จำเป็นต้องกรอก ถ้าขาด field จะได้รับ error:

```json
{
  "error": "Key: 'CreateRequest.Title' Error:Field validation for 'Title' failed on the 'required' tag"
}
```

---

## โครงสร้างโปรเจกต์

```
announcement-board/
├── backend/
│   ├── handlers/         # รับ HTTP request และส่ง response
│   ├── services/         # Business logic
│   ├── repositories/     # Query ฐานข้อมูล
│   ├── models/           # โครงสร้างข้อมูล
│   ├── main.go
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── pages/        # หน้าเว็บแต่ละหน้า
│   │   ├── components/   # UI components ที่ใช้ซ้ำได้
│   │   ├── services/     # เรียก API
│   │   ├── types/        # TypeScript types
│   │   └── dialogs/      # Dialog components
│   └── Dockerfile
└── docker-compose.yml
```

---

## รัน Tests

```bash
cd backend
go test ./services/...
```

---

## Tradeoffs

### 1. ดึงข้อมูลใหม่ทุกครั้งหลัง mutation
หลังจากสร้าง, แก้ไข, หรือลบประกาศ frontend จะดึงรายการทั้งหมดจาก API ใหม่ทุกครั้ง วิธีนี้ทำให้ข้อมูลถูกต้องเสมอ แต่ไม่มีประสิทธิภาพเท่า optimistic update ถ้าข้อมูลมีจำนวนมากควรใช้ pagination หรืออัปเดต state ในเครื่องแทน

### 2. ไม่มีระบบ Authentication
ผู้ใช้ทุกคนสามารถสร้าง, pin, หรือลบประกาศได้ ถ้าเพิ่ม ระบบ authentication จะทำให้กำหนดสิทธิ์ได้ว่าใครสามารถจัดการประกาศได้บ้าง

### 3. CORS เปิดรับทุก origin
ปัจจุบัน backend ตั้งค่า `Access-Control-Allow-Origin: *` เหมาะสำหรับ local development แต่ใน production ควรจำกัดให้รับเฉพาะ domain ที่กำหนดเท่านั้น

### 4. ไม่มี pagination
ปัจจุบันดึงประกาศทั้งหมดมาแสดงในครั้งเดียว ถ้าข้อมูลมีจำนวนมากจะทำให้โหลดช้า ควรเพิ่ม pagination หรือ infinite scroll

---

## สิ่งที่อยากเพิ่มถ้ามีเวลา

- **ระบบ Authentication** — Login ด้วย JWT เพื่อกำหนดสิทธิ์การจัดการประกาศ
- **Pagination** — โหลดประกาศทีละหน้าแทนการดึงทั้งหมดมาพร้อมกัน
- **ค้นหาและกรองข้อมูล** — ค้นหาตามหัวข้อหรือกรองตามผู้เขียน
- **แก้ไขประกาศ** — เพิ่มฟีเจอร์แก้ไขเนื้อหาประกาศที่มีอยู่แล้ว
- **Environment config** — แยก `.env` สำหรับ local และ production
- **Integration tests** — ทดสอบ HTTP request/response แบบครบวงจร ไม่ใช่แค่ service layer