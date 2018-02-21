# Inventory using Golang

No Description yet

## Getting Started

```
go get github.com/jodypati/Inventory
go run main.go
```
## Authors

* **Sausa Jodypati** - *Initial work* - [jodypati](https://github.com/jodypati)

## API's

```
METHOD  ROUTE                                       Keterangan
POST    /api/v1/barangs                             Insert barang
GET     /api/v1/barangs                             Get seluruh barang
GET     /api/v1/barangs/:sku                        Get barang berdasarkan sku
PUT     /api/v1/barangs/:sku                        Edit barang bedasarkan sku
DELETE  /api/v1/barangs/:sku                        Delete barang berdasarkan sku
GET     /api/v1/laporannilaibarang                  Generate laporan nilai barang
GET     /api/v1/importbarang                        Insert data barang dengan menggunakan file csv
GET     /api/v1/importbarangmasuk                   Insert data barang masuk dengan menggunakan file csv
GET     /api/v1/importbarangkeluar                  Insert data barang keluar dengan menggunakan file csv
GET     /api/v1/exportreport                        Export laporan nilai barang kedalam csv
GET     /api/v1/laporanpenjualan/:datefrom/:dateto  Generate Report laporan penjualan berdasarkan rentang tanggal
GET     /api/v1/truncate                            Drop database barang,brang masuk dan barang keluar

POST    /api/v1/barangmasuks                        Insert barang masuk
GET     /api/v1/barangmasuks                        Get all barang masuk
GET     /api/v1/barangmasuks/:receiptnumber         Get barang masuk berdasarkan inputan nomer kwitansi (receiptnumber)
PUT     /api/v1/barangmasuks/:receiptnumber         Edit data barang masuk berdasarkan inputan nomer kwitansi (receiptnumber)
DELETE  /api/v1/barangmasuks/:receiptnumber         Delete data barang masuk berdasarkan inputan nomer kwitansi (receiptnumber)

POST    /api/v1/barangkeluars                       Insert barang keluar
GET     /api/v1/barangkeluars                       Get all barang keluar
GET     /api/v1/barangkeluars/:receiptnumber        Get barang keluar berdasarkan inputan nomer kwitansi (receiptnumber)
PUT     /api/v1/barangkeluars/:receiptnumber        Edit data barang keluar berdasarkan inputan nomer kwitansi (receiptnumber)
DELETE  /api/v1/barangkeluars/:receiptnumber        Delete data barang keluar berdasarkan inputan nomer kwitansi (receiptnumber)

```

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

