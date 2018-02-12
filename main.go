package main

import (
	"github.com/gin-gonic/gin"
	"../inventory/object"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}


func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/barangs", object.PostBarang)
		v1.GET("/barangs", object.GetBarangs)
		v1.GET("/barangs/:sku", object.GetBarang)
		v1.PUT("/barangs/:sku", object.UpdateBarang)
		v1.DELETE("/barangs/:sku", object.DeleteBarang)

		v1.POST("/barangmasuks", object.PostBarangmasuk)
		v1.GET("/barangmasuks", object.GetBarangmasuks)
		v1.GET("/barangmasuks/:receiptnumber", object.GetBarangmasuk)
		v1.PUT("/barangmasuks/:receiptnumber", object.UpdateBarangmasuk)
		v1.DELETE("/barangmasuks/:receiptnumber", object.DeleteBarangmasuk)

		v1.POST("/barangkeluars", object.PostBarangkeluar)
		v1.GET("/barangkeluars", object.GetBarangkeluars)
		v1.GET("/barangkeluars/:receiptnumber", object.GetBarangkeluar)
		v1.PUT("/barangkeluars/:receiptnumber", object.UpdateBarangkeluar)
		v1.DELETE("/barangkeluars/:receiptnumber", object.DeleteBarangkeluar)
	}

	r.Run(":8080")
}

