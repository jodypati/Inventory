"# Inventory" 
untuk menjalankan aplikasi eksekusi perintah
go run main.go

POST   /api/v1/barangs           --> _/C_/Go/bin/inventory/object.PostBarang (4 handlers)
GET    /api/v1/barangs           --> _/C_/Go/bin/inventory/object.GetBarangs (4 handlers)
GET    /api/v1/barangs/:sku      --> _/C_/Go/bin/inventory/object.GetBarang (4 handlers)
PUT    /api/v1/barangs/:sku      --> _/C_/Go/bin/inventory/object.UpdateBarang (4 handlers)
DELETE /api/v1/barangs/:sku      --> _/C_/Go/bin/inventory/object.DeleteBarang (4 handlers)
POST   /api/v1/barangmasuks      --> _/C_/Go/bin/inventory/object.PostBarangmasuk (4 handlers)
GET    /api/v1/barangmasuks      --> _/C_/Go/bin/inventory/object.GetBarangmasuks (4 handlers)
GET    /api/v1/barangmasuks/:receiptnumber --> _/C_/Go/bin/inventory/object.GetBarangmasuk (4 handlers)
PUT    /api/v1/barangmasuks/:receiptnumber --> _/C_/Go/bin/inventory/object.UpdateBarangmasuk (4 handlers)
DELETE /api/v1/barangmasuks/:receiptnumber --> _/C_/Go/bin/inventory/object.DeleteBarangmasuk (4 handlers)
POST   /api/v1/barangkeluars     --> _/C_/Go/bin/inventory/object.PostBarangkeluar (4 handlers)
GET    /api/v1/barangkeluars     --> _/C_/Go/bin/inventory/object.GetBarangkeluars (4 handlers)
GET    /api/v1/barangkeluars/:receiptnumber --> _/C_/Go/bin/inventory/object.GetBarangkeluar (4 handlers)
PUT    /api/v1/barangkeluars/:receiptnumber --> _/C_/Go/bin/inventory/object.UpdateBarangkeluar (4 handlers)
DELETE /api/v1/barangkeluars/:receiptnumber --> _/C_/Go/bin/inventory/object.DeleteBarangkeluar (4 handlers)
