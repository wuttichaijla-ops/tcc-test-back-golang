# test-back-golang

Backend API server สำหรับ Angular frontend project

## โครงสร้างโปรเจกต์

```
test-back-golang/
├── main.go              # Entry point และ server setup
├── go.mod               # Go module dependencies
├── models/              # Data models
│   └── models.go
├── handlers/            # API handlers
│   └── handlers.go
└── middleware/          # Middleware functions
    └── cors.go
```

## API Endpoints

### Items
- `GET /api/items` - ดึงรายการ items ทั้งหมด
- `POST /api/items` - เพิ่ม item ใหม่
  ```json
  {
    "name": "Item Name",
    "description": "Item Description"
  }
  ```
- `DELETE /api/items/{id}` - ลบ item ตาม ID

### Comments
- `POST /api/addComment` - เพิ่ม comment ใหม่
  ```json
  {
    "author": "Author Name",
    "text": "Comment text",
    "avatar": "A"
  }
  ```
- `GET /api/comments` - ดึงรายการ comments ทั้งหมด

## การติดตั้งและรัน

### 1. ติดตั้ง dependencies
```bash
go mod download
```

### 2. รัน server
```bash
go run main.go
```

Server จะรันที่ `http://localhost:8080`

## การเชื่อมต่อกับ Angular Frontend

1. ตรวจสอบว่า Angular frontend รันที่ `http://localhost:4200`
2. CORS middleware ถูกตั้งค่าให้รองรับ Angular frontend แล้ว
3. อัปเดต Angular service (`data.service.ts`) ให้เรียกใช้ API จริงแทน mock data:

```typescript
// แทนที่ mock methods ด้วย HTTP calls
getItems(): Observable<Item[]> {
  return this.http.get<Item[]>('http://localhost:8080/api/items');
}

addItem(item: Item): Observable<Item> {
  return this.http.post<Item>('http://localhost:8080/api/items', item);
}

deleteItem(id: number): Observable<void> {
  return this.http.delete<void>(`http://localhost:8080/api/items/${id}`);
}
```

## Dependencies

- `github.com/gorilla/mux` - HTTP router และ URL matcher

## หมายเหตุ

- ข้อมูลถูกเก็บใน memory (in-memory storage)
- สำหรับ production ควรใช้ database แทน
- Server รันที่ port 8080 (สามารถเปลี่ยนได้ใน `main.go`)
