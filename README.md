# Inventory Management System

Sistem manajemen inventory backend yang dibangun dengan Go, Gin framework, dan MySQL. Sistem ini menyediakan RESTful API untuk mengelola produk, inventaris, dan pesanan, termasuk penanganan file untuk gambar produk.

## 🚀 Fitur

- **Manajemen Produk**: CRUD operasi untuk produk dengan dukungan upload gambar
- **Kontrol Inventaris**: Pelacakan stok dan lokasi produk
- **Sistem Pesanan**: Pencatatan dan manajemen pesanan pelanggan
- **RESTful API**: API endpoints yang terstruktur dengan baik untuk semua operasi
- **Penanganan File**: Upload dan download gambar produk

## 📋 Syarat

Sebelum menjalankan project ini, pastikan telah menginstall:

- Go (versi 1.16 atau lebih baru)
- MySQL
- Git

## 🛠️ Instalasi

1. Clone repositori
```bash
git clone https://github.com/Malvoy/inventory-management.git
cd inventory-management
```

2. Install dependensi
```bash
go mod download
```

3. Konfigurasi database
- Buat database MySQL baru
- Copy `.env.example` ke `.env`
- Sesuaikan konfigurasi database di file `.env`:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=inventory_db
```

4. Jalankan migrasi database
```bash
mysql -u your_username -p your_database < database/schema.sql
```

## 🚀 Menjalankan Aplikasi

1. Jalankan server
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## 📌 API Endpoints

### Produk
- `GET /api/products` - Mendapatkan semua produk
- `GET /api/products/:id` - Mendapatkan produk berdasarkan ID
- `POST /api/products` - Membuat produk baru
- `PUT /api/products/:id` - Memperbarui produk
- `DELETE /api/products/:id` - Menghapus produk
- `POST /api/products/:id/image` - Upload gambar produk
- `GET /api/products/:id/image` - Download gambar produk

### Inventaris
- `GET /api/inventory` - Mendapatkan semua data inventaris
- `GET /api/inventory/:productId` - Mendapatkan inventaris berdasarkan ID produk
- `PUT /api/inventory/:productId` - Memperbarui stok inventaris

### Pesanan
- `POST /api/orders` - Membuat pesanan baru
- `GET /api/orders/:id` - Mendapatkan detail pesanan
- `GET /api/orders` - Mendapatkan semua pesanan

## 🧪 Testing

Untuk menjalankan unit test:
```bash
go test ./...
```

## 📝 Format Request

### Membuat Produk Baru
```json
POST /api/products
{
    "name": "Laptop Gaming",
    "description": "Laptop gaming performa tinggi",
    "price": 15000000,
    "category": "Electronics"
}
```

### Memperbarui Inventaris
```json
PUT /api/inventory/:productId
{
    "quantity": 100,
    "location": "Warehouse A"
}
```

### Membuat Pesanan
```json
POST /api/orders
{
    "product_id": 1,
    "quantity": 2,
    "order_date": "2024-03-24T15:04:05Z"
}
```

## 🤝 Kontribusi

Apabila ingin berkontribusi selalu diterima! Silakan bisa fork repositori ini dan buat pull request untuk setiap perubahan yang ingin kakak-kakak usulkan.


## 📞 Kontak

Jika memiliki pertanyaan atau masukan, silakan buat issue di repositori ini atau hubungi melalui:

- GitHub: [@Malvoy](https://github.com/Malvoy)
