package object

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"bytes"
	"strconv"
	"strings"
)

type Barangmasuks struct {
	Id        		int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Time			time.Time `gorm:"not null" form:"time" json:"time"`
	SKU 			string `gorm:"not null" form:"sku" json:"sku"`
	ItemName 		string `gorm:"not null" form:"itemname" json:"itemname"`
	OrderAmount 	int `gorm:"not null" form:"orderamount" json:"orderamount"`
	AmountRecieved  int `gorm:"not null" form:"amountrecieved" json:"amountrecieved"`
	PurchasePrice  	float64 `gorm:"not null" form:"purchaseprice" json:"purchaseprice"`
	Total  			float64 `gorm:"not null" form:"total" json:"total"`
	ReceiptNumber	string `gorm:"not null" form:"receiptnumber" json:"receiptnumber"`
	Notes			string `gorm:"not null" form:"notes" json:"notes"`

}

func InitDbBarangmasuk() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Barangmasuks{}) {
		db.CreateTable(&Barangmasuks{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Barangmasuks{})
	}

	return db
}

func PostBarangmasuk(c *gin.Context) {
	db := InitDbBarangmasuk()
	defer db.Close()

	var barangmasuk Barangmasuks
	c.Bind(&barangmasuk)

	if barangmasuk.SKU != "" && barangmasuk.ItemName != "" && barangmasuk.OrderAmount >= 0 && barangmasuk.AmountRecieved >= 0 && barangmasuk.PurchasePrice >= 0 {
		t := time.Now()
		barangmasuk.Time = t
		var buffer bytes.Buffer
		buffer.WriteString(t.Format("2006-01-02 15:04:05"))
		buffer.WriteString(" ")
		buffer.WriteString("terima ")
		buffer.WriteString(strconv.Itoa(barangmasuk.AmountRecieved))
		if barangmasuk.OrderAmount > barangmasuk.AmountRecieved {
			buffer.WriteString(";Masih Menunggu")
			barangmasuk.Notes = buffer.String()
		}
		barangmasuk.Total = float64(barangmasuk.OrderAmount) * barangmasuk.PurchasePrice
		// INSERT INTO "barangmasuks" (name) VALUES (barangmasuk.Name);
		db.Create(&barangmasuk)
		var buffer2 bytes.Buffer
		buffer2.WriteString(t.Format("20060102"))
		buffer2.WriteString("-")
		buffer2.WriteString(strconv.Itoa(barangmasuk.Id))
		db.Model(&barangmasuk).Update("receiptnumber",buffer2.String())
		// Display error
		c.JSON(201, gin.H{"success": barangmasuk})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/barangmasuks
}

func GetBarangmasuks(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangmasuk()
	// Close connection database
	defer db.Close()

	var barangmasuks []Barangmasuks
	// SELECT * FROM barangmasuks
	db.Find(&barangmasuks)

	// Display JSON result
	c.JSON(200, barangmasuks)

	// curl -i http://localhost:8080/api/v1/barangmasuks
}

func GetBarangmasuk(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangmasuk()
	// Close connection database
	defer db.Close()

	receiptnumber := c.Params.ByName("receiptnumber")
	var barangmasuk Barangmasuks
	// SELECT * FROM barangmasuks WHERE receiptnumber = SSI-D00791015-LL-BWH;
	db.First(&barangmasuk, receiptnumber)

	if barangmasuk.SKU != "" {
		// Display JSON result
		c.JSON(200, barangmasuk)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Barangmasuk not found"})
	}

	// curl -i http://localhost:8080/api/v1/barangmasuks/1
}

func UpdateBarangmasuk(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangmasuk()
	// Close connection database
	defer db.Close()

	// Get id barangmasuk
	receiptnumber := c.Params.ByName("receiptnumber")
	var barangmasuk Barangmasuks
	// SELECT * FROM barangmasuks WHERE receiptnumber = 1;
	db.First(&barangmasuk, receiptnumber)

	if barangmasuk.SKU != "" && barangmasuk.ItemName != "" && barangmasuk.OrderAmount >= 0 && barangmasuk.AmountRecieved >= 0 && barangmasuk.PurchasePrice >= 0 && barangmasuk.Total >= 0 && barangmasuk.ReceiptNumber != "" {

		if barangmasuk.ReceiptNumber != "" || barangmasuk.OrderAmount != barangmasuk.AmountRecieved {
			var newBarangmasuk Barangmasuks
			c.Bind(&newBarangmasuk)
			amountNow := barangmasuk.AmountRecieved + newBarangmasuk.AmountRecieved
			remain := barangmasuk.OrderAmount - amountNow
			var buffer bytes.Buffer
			t := time.Now()
			notes := strings.Split(barangmasuk.Notes, ";")
			if len(notes) > 0 {
				for i := 0; i < len(notes)-1; i++ {
					buffer.WriteString(notes[i])
					buffer.WriteString(";")
				}
			}else {
				buffer.WriteString(barangmasuk.Notes)
				buffer.WriteString(";")
			}
			buffer.WriteString(t.Format("2006-01-02 15:04:05"))
			buffer.WriteString(" ")
			buffer.WriteString("terima ")
			buffer.WriteString(strconv.Itoa(newBarangmasuk.AmountRecieved))
			if remain > 0 {
				buffer.WriteString("; Masih Menunggu")
			}

			result := Barangmasuks{
				Id:        barangmasuk.Id,
				Time: barangmasuk.Time,
				SKU:	barangmasuk.SKU,
				ItemName: barangmasuk.ItemName,
				OrderAmount: barangmasuk.OrderAmount,
				PurchasePrice: barangmasuk.PurchasePrice,
				Total: barangmasuk.Total,
				ReceiptNumber:	barangmasuk.ReceiptNumber, 
				AmountRecieved: amountNow,
				Notes: buffer.String(),
			}

			// UPDATE barangmasuks SET firstname='newBarangmasuk.Firstname', lastname='newBarangmasuk.Lastname' WHERE id = barangmasuk.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Barangmasuk not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/barangmasuks/1
}

func DeleteBarangmasuk(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangmasuk()
	// Close connection database
	defer db.Close()

	// Get id barangmasuk
	receiptnumber := c.Params.ByName("receiptnumber")
	var barangmasuk Barangmasuks
	// SELECT * FROM barangmasuks WHERE receiptnumber = 1;
	db.First(&barangmasuk, receiptnumber)

	if barangmasuk.Id != 0 {
		// DELETE FROM barangmasuks WHERE receiptnumber = barangmasuk.Id
		db.Delete(&barangmasuk)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Barangmasuk #" + receiptnumber + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Barangmasuk not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/barangmasuks/1
}

func OptionsBarangmasuk(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}