package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/result", func(c *gin.Context) {
		result := []map[string]interface{}{
			{
				"name": "Rafe Uddaraj",
				"roll": "672641",
				"gpa":  []float32{3.64, 2.64, 3.14},
			},
			{
				"name": "Maharaf Hossen",
				"roll": "658843",
				"gpa":  []float32{2.90, 0, 0},
			},
			{
				"name": "X",
				"roll": "658845",
				"gpa":  []float32{2.90, 0, 0},
			},
		} // Correctly initialize the array of maps
		c.JSON(200, gin.H{
			"result": result,
		})
	})
	r.Run(":1211") // Listen and serve on 0.0.0.0:1211
}
