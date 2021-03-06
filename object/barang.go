package object

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"os"
    "log"
    "encoding/csv"
    "strconv"
)
type Barangs struct {
	Id        	int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	SKU 		string `gorm:"not null" form:"sku" json:"sku"`
	ItemName 	string `gorm:"not null" form:"itemname" json:"itemname"`
	Quantity  	int `gorm:"not null" form:"quantity" json:"quantity"`
}

type HeaderReport struct {
	Date string `gorm:"not null" form:"date" json:"date"`
	ItemsAmount int `gorm:"not null" form:"itemsamount" json:"itemsamount"`
	GoodsAmount float64 `gorm:"not null" form:"goodsamount" json:"goodsamount"`
	ValueTotal float64 `gorm:"not null" form:"valuetotal" json:"valuetotal"`
	Detail *[]Detail `gorm:"not null" form:"detail" json:"detail"`
}


type Detail struct {
    SKU 		string `gorm:"not null" form:"sku" json:"sku"`
	ItemName 	string `gorm:"not null" form:"itemname" json:"itemname"`
	quantity 	int `gorm:"not null" form:"quantity" json:"quantity"`
	AveragePrice 	float64 `gorm:"not null" form:"averageprice" json:"averageprice"`
	Total float64 `gorm:"not null" form:"total" json:"total"`
}

type SalesHeader struct {
	Date string `gorm:"not null" form:"date" json:"date"`
	DateFrom string `gorm:"not null" form:"datefrom" json:"datefrom"`
	DateTo string `gorm:"not null" form:"dateto" json:"dateto"`
	Turnover float64 `gorm:"not null" form:"turnover" json:"turnover"`
	GrossProfit float64 `gorm:"not null" form:"grossprofit" json:"grossprofit"`
	TotalSales int `gorm:"not null" form:"totalsales" json:"totalsales"`
	TotalGoods int `gorm:"not null" form:"totalgoods" json:"totalgoods"`
	SalesDetail *[]SalesDetail `gorm:"not null" form:"detail" json:"detail"`
}

type SalesDetail struct {
    BookingId 		string `gorm:"not null" form:"bookingid" json:"bookingid"`
    Time			time.Time `gorm:"not null" form:"time" json:"time"`
    SKU 			string `gorm:"not null" form:"sku" json:"sku"`
	ItemName 		string `gorm:"not null" form:"itemname" json:"itemname"`
	Quantity 		int `gorm:"not null" form:"quantity" json:"quantity"`
	SalePrice 		float64 `gorm:"not null" form:"saleprice" json:"saleprice"`
	TotalSale 		float64 `gorm:"not null" form:"totalsale" json:"totalsale"`
	PurchasePrice 	float64 `gorm:"not null" form:"purchaseprice" json:"purchaseprice"`
	Profit			float64 `gorm:"not null" form:"profit" json:"profit"`
}

func InitDbBarang() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Barangs{}) {
		db.CreateTable(&Barangs{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Barangs{})
	}

	return db
}

func PostBarang(c *gin.Context) {
	db := InitDbBarang()
	defer db.Close()

	var barang Barangs
	c.Bind(&barang)

	if barang.SKU != "" && barang.ItemName != "" && barang.Quantity >= 0 {
		// INSERT INTO "barangs" (name) VALUES (barang.Name);
		db.Create(&barang)
		// Display error
		c.JSON(201, gin.H{"success": barang})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/barangs
}

func GetBarangs(c *gin.Context) {
	// Connection to the database
	db := InitDbBarang()
	// Close connection database
	defer db.Close()

	var barangs []Barangs
	// SELECT * FROM barangs
	db.Find(&barangs)

	// Display JSON result
	c.JSON(200, barangs)

	// curl -i http://localhost:8080/api/v1/barangs
}

func GetBarang(c *gin.Context) {
	// Connection to the database
	db := InitDbBarang()
	// Close connection database
	defer db.Close()

	sku := c.Params.ByName("sku")
	var barang Barangs
	// SELECT * FROM barangs WHERE sku = SSI-D00791015-LL-BWH;
	db.First(&barang, sku)

	if barang.SKU != "" {
		// Display JSON result
		c.JSON(200, barang)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Barang not found"})
	}

	// curl -i http://localhost:8080/api/v1/barangs/1
}

func UpdateBarang(c *gin.Context) {
	// Connection to the database
	db := InitDbBarang()
	// Close connection database
	defer db.Close()

	// Get id barang
	sku := c.Params.ByName("sku")
	var barang Barangs
	// SELECT * FROM barangs WHERE sku = 1;
	db.First(&barang, sku)

	if barang.ItemName != "" && barang.Quantity >= 0 {

		if barang.SKU != "" {
			var newBarang Barangs
			c.Bind(&newBarang)

			result := Barangs{
				Id:        barang.Id,
				SKU:	barang.SKU, 
				ItemName: newBarang.ItemName,
				Quantity:  newBarang.Quantity,
			}

			// UPDATE barangs SET firstname='newBarang.Firstname', lastname='newBarang.Lastname' WHERE id = barang.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Barang not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/barangs/1
}

func DeleteBarang(c *gin.Context) {
	// Connection to the database
	db := InitDbBarang()
	// Close connection database
	defer db.Close()

	// Get id barang
	sku := c.Params.ByName("sku")
	var barang Barangs
	// SELECT * FROM barangs WHERE sku = 1;
	db.First(&barang, sku)

	if barang.Id != 0 {
		// DELETE FROM barangs WHERE sku = barang.Id
		db.Delete(&barang)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Barang #" + sku + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Barang not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/barangs/1
}

func OptionsBarang(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
func ImportBarang(c *gin.Context){
	filename := "barang.csv"

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Read File into a Variable
    r := csv.NewReader(f)
    r.Comma = (';')
    lines, err := r.ReadAll()
    if err != nil {
        panic(err)
		// Display error
		c.JSON(422, gin.H{"error": err})
	}else{
		db := InitDbBarang()
		defer db.Close()

		var barangs []Barangs
		// Loop through lines & turn into object
    	for _, line := range lines {
    		
    		jumlah, err := strconv.Atoi(line[2])
    		if err == nil {
	        	barang := Barangs{
			        SKU: line[0],
			        ItemName: line[1],
			        Quantity: jumlah,
		        } 
		        db.Create(&barang)
		        barangs = append(barangs,barang)
	    	}else{
	    		c.JSON(422, gin.H{"error": err})		
	    	}
	    }
		
		// Display error
		c.JSON(200, gin.H{"success": barangs})
	}
}
func ExportValueReport(c *gin.Context){
	db := InitDbBarang()
	// Close connection database
	defer db.Close()
	var details []Detail
	var err error
	db.Table("barangs").Select("barangs.sku,barangs.item_name,barangs.quantity,SUM(total)/SUM(amount_recieved) as average_price, SUM(total)/SUM(amount_recieved)*barangs.quantity as total").Joins("inner join barangmasuks on barangs.SKU=barangmasuks.SKU").Group("barangs.sku,barangs.item_name,barangs.quantity").Scan(&details)

	file, err := os.Create("ValueReport.csv")
    
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
	for _, detail := range details {
        
        data := []string{detail.SKU,detail.ItemName,strconv.Itoa(detail.quantity),strconv.FormatFloat(detail.AveragePrice, 'f', -1, 64),strconv.FormatFloat(detail.Total, 'f', -1, 64)}
        err := writer.Write(data)
        
        if err != nil {
        	log.Fatal("Cannot write file", err)
    	}
	}
	if err == nil {
        c.JSON(200, gin.H{"success": "Report Created"})
    }else{
    	c.JSON(404, gin.H{"error": "Cannot create file"})
    }
    	
}

func GoodsValueReport(c *gin.Context) {
	//Connection to the database
	db := InitDbBarang()
	// Close connection database
	defer db.Close()
	var details []Detail
	var goodsAmount float64
	var valueTotal float64
	
	db.Table("barangs").Select("barangs.sku,barangs.item_name,barangs.quantity,SUM(total)/SUM(amount_recieved) as average_price, SUM(total)/SUM(amount_recieved)*barangs.quantity as total").Joins("inner join barangmasuks on barangs.SKU=barangmasuks.SKU").Group("barangs.sku,barangs.item_name,barangs.quantity").Scan(&details)
	
	itemsAmount := 0;
	for _, detail := range details {
        goodsAmount += detail.AveragePrice
        valueTotal += detail.Total
        itemsAmount++
	}


	t := time.Now()
	result := HeaderReport {
		Date: t.Format("20060102"),
		ItemsAmount: itemsAmount,
		GoodsAmount: goodsAmount,
		ValueTotal: valueTotal,
		Detail: &details,
	}
	if itemsAmount > 0 {
	// Display modified data in JSON message "success"
		c.JSON(200, gin.H{"success": result})
	}else{
		c.JSON(404, gin.H{"error": "Report not found"})
	}
}

func SalesReport(c *gin.Context) {
	//Connection to the database
	db := InitDbBarang()
	// Close connection database
	defer db.Close()
	const dateFormat = "2006-01-02"
	dateFrom, _ := time.Parse(dateFormat, c.Params.ByName("datefrom"))
	dateTo, _ := time.Parse(dateFormat, c.Params.ByName("dateto"))
	var salesDetail []SalesDetail
	// var goodsAmount float64
	// var valueTotal float64
	
	db.Table("barangkeluars").Select("barangkeluars.receipt_number as booking_id,barangkeluars.time,barangkeluars.sku,barangkeluars.item_name,barangkeluars.stock_out as quantity,barangkeluars.selling_price as sale_price,barangkeluars.selling_price*barangkeluars.stock_out as total_sale, SUM(barangmasuks.total)/SUM(amount_recieved) as purchase_price, (barangkeluars.selling_price*barangkeluars.stock_out) - (barangkeluars.stock_out * (SUM(barangmasuks.total)/SUM(amount_recieved))) as profit").Joins("inner join barangmasuks on barangkeluars.SKU=barangmasuks.SKU").Where("barangkeluars.time >= ? and barangkeluars.time <= ?",dateFrom,dateTo).Group("barangkeluars.receipt_number,barangkeluars.time,barangkeluars.sku,barangkeluars.item_name,barangkeluars.stock_out,barangkeluars.selling_price").Scan(&salesDetail)
	
	totalGoods := 0
	totalSales := 0
	turnover := 0.0
	grossProfit := 0.0
	for _, detail := range salesDetail {
        totalGoods += detail.Quantity
        if detail.BookingId != ""{
        	totalSales += detail.Quantity
        }
        turnover += detail.TotalSale
        grossProfit += detail.Profit
	}


	t := time.Now()
	result := SalesHeader {
		Date: t.Format("02 Jan 2006"),
		DateFrom: dateFrom.Format("02 Jan 2006"),
		DateTo: dateTo.Format("02 Jan 2006"),
		Turnover: turnover,
		GrossProfit: grossProfit,
		TotalSales: totalSales,
		TotalGoods: totalGoods,
		SalesDetail: &salesDetail,
	}
	if totalGoods > 0 {
	// Display modified data in JSON message "success"
		c.JSON(200, gin.H{"success": result})
	}else{
		c.JSON(404, gin.H{"error": "Report not found"})
	}
}

func TruncateAll(c *gin.Context) {
	//Connection to the database
	db := InitDbBarang()
	// Close connection database
	defer db.Close()
	db.Exec("DROP TABLE Barangs;")
	db.Exec("DROP TABLE Barangmasuks;")
	db.Exec("DROP TABLE Barangkeluars;")
}