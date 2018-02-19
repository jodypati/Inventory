package object

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"bytes"
	"strconv"
	"os"
    "encoding/csv"
    "strings"
)

type Barangkeluar struct {
	Id        		int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Time			time.Time `gorm:"not null" form:"time" json:"time"`
	SKU 			string `gorm:"not null" form:"sku" json:"sku"`
	ItemName 		string `gorm:"not null" form:"itemname" json:"itemname"`
	StockOut  int `gorm:"not null" form:"stockout" json:"stockout"`
	SellingPrice  	float64 `gorm:"not null" form:"sellingprice" json:"sellingprice"`
	Total  			float64 `gorm:"not null" form:"total" json:"total"`
	ReceiptNumber	string `gorm:"not null" form:"receiptnumber" json:"receiptnumber"`
	Notes			string `gorm:"not null" form:"notes" json:"notes"`

}

func InitDbBarangkeluar() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Barangkeluar{}) {
		db.CreateTable(&Barangkeluar{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Barangkeluar{})
	}

	return db
}

func PostBarangkeluar(c *gin.Context) {
	db := InitDbBarangkeluar()
	defer db.Close()

	var barangkeluar Barangkeluar
	c.Bind(&barangkeluar)

	if barangkeluar.SKU != "" && barangkeluar.ItemName != "" && barangkeluar.StockOut >= 0 && barangkeluar.SellingPrice >= 0 && barangkeluar.ReceiptNumber != "" && barangkeluar.Notes != "" {
		t := time.Now()
		barangkeluar.Time = t
		//var buffer bytes.Buffer
		//buffer.WriteString("Pesanan ID-")
		//buffer.WriteString(t.Format("20060102"))
		//buffer.WriteString("-")
		//buffer.WriteString("terima ")
		//buffer.WriteString(strconv.Itoa(barangkeluar.StockOut))
		
		barangkeluar.Total = float64(barangkeluar.StockOut) * barangkeluar.SellingPrice
		// INSERT INTO "barangkeluars" (name) VALUES (barangkeluar.Name);
		db.Create(&barangkeluar)
		var buffer2 bytes.Buffer
		buffer2.WriteString("ID-")
		buffer2.WriteString(t.Format("20060102"))
		buffer2.WriteString("-")
		buffer2.WriteString(strconv.Itoa(barangkeluar.Id))
		db.Model(&barangkeluar).Update("receiptnumber",buffer2.String())
		// Display error
		c.JSON(201, gin.H{"success": barangkeluar})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/barangkeluars
}

func GetBarangkeluars(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangkeluar()
	// Close connection database
	defer db.Close()

	var barangkeluars []Barangkeluar
	// SELECT * FROM barangkeluars
	db.Find(&barangkeluars)

	// Display JSON result
	c.JSON(200, barangkeluars)

	// curl -i http://localhost:8080/api/v1/barangkeluars
}

func GetBarangkeluar(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangkeluar()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var barangkeluar Barangkeluar
	// SELECT * FROM barangkeluars WHERE id = SSI-D00791015-LL-BWH;
	db.First(&barangkeluar, id)

	if barangkeluar.SKU != "" {
		// Display JSON result
		c.JSON(200, barangkeluar)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Barangkeluar not found"})
	}

	// curl -i http://localhost:8080/api/v1/barangkeluars/1
}

func UpdateBarangkeluar(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangkeluar()
	// Close connection database
	defer db.Close()

	// Get id barangkeluar
	id := c.Params.ByName("id")
	var barangkeluar Barangkeluar
	// SELECT * FROM barangkeluars WHERE id = 1;
	db.First(&barangkeluar, id)

	if barangkeluar.SKU != "" && barangkeluar.ItemName != "" && barangkeluar.Notes != "" && barangkeluar.StockOut >= 0 && barangkeluar.SellingPrice >= 0 && barangkeluar.Total >= 0 && barangkeluar.ReceiptNumber != "" {

		if barangkeluar.ReceiptNumber != "" {
			var newBarangkeluar Barangkeluar
			c.Bind(&newBarangkeluar)
			newBarangkeluar.Total = float64(newBarangkeluar.StockOut) * barangkeluar.SellingPrice
			
			result := Barangkeluar{
				Id:        barangkeluar.Id,
				Time: barangkeluar.Time,
				SKU:	barangkeluar.SKU,
				ItemName: barangkeluar.ItemName,
				SellingPrice: barangkeluar.SellingPrice,
				Total: newBarangkeluar.Total,
				ReceiptNumber:	barangkeluar.ReceiptNumber, 
				StockOut: newBarangkeluar.StockOut,
				Notes: newBarangkeluar.Notes,
			}

			// UPDATE barangkeluars SET firstname='newBarangkeluar.Firstname', lastname='newBarangkeluar.Lastname' WHERE id = barangkeluar.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Barangkeluar not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/barangkeluars/1
}

func DeleteBarangkeluar(c *gin.Context) {
	// Connection to the database
	db := InitDbBarangkeluar()
	// Close connection database
	defer db.Close()

	// Get id barangkeluar
	id := c.Params.ByName("id")
	var barangkeluar Barangkeluar
	// SELECT * FROM barangkeluars WHERE id = 1;
	db.First(&barangkeluar, id)

	if barangkeluar.Id != 0 {
		// DELETE FROM barangkeluars WHERE id = barangkeluar.Id
		db.Delete(&barangkeluar)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Barangkeluar #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Barangkeluar not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/barangkeluars/1
}

func OptionsBarangkeluar(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func ImportBarangKeluar(c *gin.Context){
	filename := "barangkeluar.csv"
	const dateFormat = "02/01/2006 15.04:05"
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

		var barangkeluars []Barangkeluar
		// Loop through lines & turn into object
    	for _, line := range lines {
    		t, _ := time.Parse(dateFormat, line[0])
    		stockout, err := strconv.Atoi(line[3])
    		sellingprice,err := strconv.ParseFloat(line[4], 64)
    		total,err := strconv.ParseFloat(line[5], 64)
    		receiptnumber := ""
    		if sellingprice != 0.0 {
    			notes := strings.Split(line[6], " ")
				if len(notes) > 1 {
					receiptnumber = notes[1]
				}
    		}
    		if err == nil {
	        	barangkeluar := Barangkeluar{
			        Time: t,
			        SKU: line[1],
			        ItemName: line[2],
			        StockOut: stockout,
			        SellingPrice: sellingprice,
			        Total: total,
			        ReceiptNumber: receiptnumber,
			        Notes: line[6],
			        
		        } 
		        db.Create(&barangkeluar)
		        barangkeluars = append(barangkeluars,barangkeluar)
	    	}else{
	    		c.JSON(422, gin.H{"error": err})		
	    	}
	    }
		
		// Display error
		c.JSON(200, gin.H{"success": barangkeluars})
	}
}