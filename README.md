# test-back-golang

เอกสารสรุปสำหรับ Backend API server (Go) ที่ใช้ร่วมกับ Angular frontend

## โครงสร้างโปรเจกต์

```
test-back-golang/
├── main.go              # Entry point และ server setup
├── go.mod               # Go module dependencies
├── models/              # Data models
│   └── models.go
├── handlers/            # API handlers
│   └── handlers.go
├── handlers/            # tests
│   └── handlers_test.go
└── middleware/          # Middleware functions
    └── cors.go
```

## เจตนาเอกสาร

ไฟล์นี้อธิบาย API endpoint หลัก รูปแบบ request/response ตัวอย่างการใช้งาน และขั้นตอนรันเซิร์ฟเวอร์

## การตั้งค่า (Environment)

- PORT: (optional) พอร์ตที่เซิร์ฟเวอร์จะรัน ค่าเริ่มต้นคือ `8080`
- CORS_ALLOWED_ORIGIN: (optional) origin ที่อนุญาตให้เรียก API (เช่น `http://localhost:4200`)

วิธีรัน (development):

```bash
go mod download
go run main.go
```

หรือกำหนดพอร์ตแบบเร็ว:

```bash
PORT=8081 go run main.go
```

Server เริ่มต้นจะรันที่ `http://localhost:8080` (หรือพอร์ตที่กำหนดผ่าน env)

## API Contract — รูปแบบมาตรฐาน

ทุก endpoint ส่งกลับ JSON รูปแบบมาตรฐาน:

{
  "status": "success" | "error",
  "message": "ข้อความอธิบาย",
  "data": <object | array | null>
}

ตัวอย่าง error:

{
  "status": "error",
  "message": "Item not found",
  "data": null
}

## Endpoints

### Items
- GET /api/items
  - คำอธิบาย: ดึงรายการ items ทั้งหมด
  - Response 200:
    {
      "status": "success",
      "message": "Items retrieved",
      "data": [ {"id":1, "name":"Item Name", "description":"..."}, ... ]
    }

- POST /api/items
  - คำอธิบาย: เพิ่ม item ใหม่
  - Request JSON:
    {
      "name": "Item Name",
      "description": "Item Description"
    }
  - Response 201:
    {
      "status":"success",
      "message":"Item created",
      "data": {"id": 1, "name":"Item Name", "description":"Item Description"}
    }

- DELETE /api/items/{id}
  - คำอธิบาย: ลบ item ตาม ID
  - Response 200 (on success):
    {
      "status":"success",
      "message":"Item deleted",
      "data": null
    }

ตัวอย่าง curl (เพิ่ม item):

```bash
curl -X POST http://localhost:8080/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"New item","description":"desc"}'
```

### Comments
- POST /api/addComment
  - Request JSON:
    {
      "author": "Author Name",
      "text": "Comment text",
      "avatar": "A"
    }
  - Response 201:
    {
      "status":"success",
      "message":"Comment added",
      "data": {"id":1, "author":"Author Name", "text":"Comment text","avatar":"A"}
    }

- GET /api/comments
  - Response 200:
    {
      "status":"success",
      "message":"Comments retrieved",
      "data": [ {"id":1, "author":"...","text":"...","avatar":"A"}, ... ]
    }

## การเชื่อมต่อกับ Angular Frontend

ตรวจสอบว่า Angular รันที่ `http://localhost:4200` (หรืออัปเดต `CORS_ALLOWED_ORIGIN`)

ตัวอย่าง service (Angular) — เรียกตาม contract:

```typescript
getItems(): Observable<ApiResponse<Item[]>> {
  return this.http.get<ApiResponse<Item[]>>('http://localhost:8080/api/items');
}

addItem(item: Item): Observable<ApiResponse<Item>> {
  return this.http.post<ApiResponse<Item>>('http://localhost:8080/api/items', item);
}

deleteItem(id: number): Observable<ApiResponse<null>> {
  return this.http.delete<ApiResponse<null>>(`http://localhost:8080/api/items/${id}`);
}
```

หมายเหตุ: `ApiResponse<T>` ควรเป็น interface ที่แมปกับรูปแบบ JSON ข้างต้น

## บันทึกเกี่ยวกับ data storage

- ปัจจุบันข้อมูลเก็บใน memory (in-memory storage). ข้อมูลจะหายเมื่อ server รีสตาร์ท
- สำหรับ production ควรเปลี่ยนมาใช้ database (เช่น PostgreSQL, MySQL) และเพิ่มการเชื่อมต่อใน `datasource/database.go`
- มีไฟล์ `product_codes_schema.sql` สำหรับอ้างอิง schema (ตรวจสอบและแมปกับ `models/models.go` หากต้องการใช้ DB)

## Dependencies

- `github.com/gorilla/mux` — HTTP router และ URL matcher

## การพัฒนาและทดสอบเบื้องต้น

1. ติดตั้ง dependencies:

```bash
go mod download
```

2. build หรือรัน:

```bash
go build ./...
# หรือ
go run main.go
```

3. ถ้ามี unit tests:

```bash
go test ./...
```

## ข้อเสนอแนะเพิ่มเติม (next steps)

- ถ้าต้องการให้ README ระบุ schema ที่ชัดเจน รบกวนอนุมัติให้ผมแมป `product_codes_schema.sql` เข้ากับ `models/models.go` และเพิ่มตัวอย่าง migration/connection string
- ถ้าต้องการรูปแบบ response ต่างจากที่เสนอ โปรดส่งตัวอย่าง (เช่น error format ของทีม)

---

เอกสารฉบับนี้เป็นเวอร์ชันที่ทำให้ API มี contract ชัดเจนขึ้น หากต้องการปรับเปลี่ยนเพิ่มเติมบอกผมได้เลย
